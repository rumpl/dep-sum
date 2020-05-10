package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/coreos/go-semver/semver"
	"github.com/dustin/go-humanize"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	version = "dev"
)

type depSize struct {
	name string
	size int64
}

func getDepSize(path string) (int64, error) {
	var size int64
	if fi, err := os.Stat(path); err == nil {
		if !fi.IsDir() {
			return 0, errors.New("not a directory")
		}
	}
	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			size += info.Size()
		}
		return err
	})
	return size, err
}

func parseGoSum(contents string) []string {
	deps := map[string]string{}

	var result []string

	lines := strings.Split(contents, "\n")

	for _, line := range lines {
		parts := strings.Split(line, " ")
		if len(parts) < 2 {
			continue
		}

		name := parts[0]
		version := parts[1][1:]

		// I have no idea what this means in a go.sum file...
		if strings.HasSuffix(version, "go.mod") {
			continue
		}

		v, err := semver.NewVersion(version)
		if err != nil {
			logrus.Warn(err)
			continue
		}

		if dv, ok := deps[name]; ok {
			v1, err := semver.NewVersion(dv)
			if err != nil {
				logrus.Warn(err)
				continue
			}

			if v1.LessThan(*v) {
				deps[name] = version
			}
		} else {
			deps[name] = version
		}
	}

	for k, v := range deps {
		result = append(result, k+"@v"+v)
	}
	return result
}

func runDepSum(opts rootOpts, path string) error {
	f := filepath.Join(path, "go.sum")
	buf, err := ioutil.ReadFile(f)
	if err != nil {
		return err
	}

	deps := parseGoSum(string(buf))
	var total int64

	depsizes := []depSize{}

	for _, dep := range deps {
		size, err := getDepSize(filepath.Join(os.Getenv("GOPATH"), "pkg", "mod", dep))
		if err == nil {
			depsizes = append(depsizes, depSize{
				name: dep,
				size: size,
			})
			total += size
		}
	}

	sort.Slice(depsizes, func(i, j int) bool {
		if opts.reverse {
			return depsizes[i].size > depsizes[j].size
		}
		return depsizes[i].size <= depsizes[j].size
	})

	if opts.quiet {
		fmt.Println(total)
		return nil
	}

	for _, d := range depsizes {
		fmt.Printf("%s\t%s\n", humanize.Bytes(uint64(d.size)), d.name)
	}
	fmt.Printf("\nTotal dependencies size: %s\n", humanize.Bytes(uint64(total)))

	return nil
}

type rootOpts struct {
	reverse bool
	quiet   bool
	verbose bool
}

func main() {
	var opts rootOpts
	root := cobra.Command{
		Use:     "dep-sum",
		Version: version,
		Args:    cobra.ExactArgs(1),
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if opts.verbose {
				logrus.SetLevel(logrus.DebugLevel)
			}
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return runDepSum(opts, args[0])
		},
	}

	root.Flags().BoolVarP(&opts.reverse, "reverse", "r", false, "Sort in reverse order")
	root.Flags().BoolVarP(&opts.quiet, "quiet", "q", false, "Only print the total size in bytes")
	root.Flags().BoolVarP(&opts.verbose, "verbose", "v", false, "Verbose output")

	if err := root.Execute(); err != nil {
		logrus.Fatal(err)
	}
}
