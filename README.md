# dep-sum

Get an aproximation of the download size of your go dependencies.

_Note_: This works only for projects that use go modules.

## Usage

```
$ dep-sum <PATH>
```

Example, for this project a build would need to download ~10MB of dependencies.

```
❯ ./dep-sum  .
5.0 kB	github.com/konsorten/go-windows-terminal-sequences@v1.0.3
5.5 kB	github.com/inconshreveable/mousetrap@v1.0.0
22 kB	github.com/kr/text@v0.1.0
32 kB	github.com/kr/pretty@v0.1.0
34 kB	github.com/coreos/go-semver@v0.3.0
36 kB	github.com/pmezard/go-difflib@v1.0.0
65 kB	github.com/dustin/go-humanize@v1.0.0
144 kB	gopkg.in/check.v1@v1.0.0-20180628173108-788fd7840127
170 kB	github.com/sirupsen/logrus@v1.6.0
212 kB	github.com/davecgh/go-spew@v1.1.1
255 kB	github.com/spf13/pflag@v1.0.3
334 kB	gopkg.in/yaml.v2@v2.2.2
371 kB	github.com/stretchr/testify@v1.2.2
504 kB	github.com/spf13/cobra@v1.0.0
8.3 MB	golang.org/x/sys@v0.0.0-20190422165155-953cdadca894

Total dependencies size: 10 MB
```

# License

[MIT](https://rumpl.mit-license.org/)
