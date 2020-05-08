package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseGoSum(t *testing.T) {
	contents := `
github.com/beorn7/perks v1.0.0 h1:HWo1m869IqiPhD389kmkxeTalrjNbbJTC8LXupb+sl0=
github.com/beorn7/perks v1.0.0/go.mod h1:KWe93zE9D1o94FZ5RNwFwVgaQK1VOXiVxmqh+CedLV8=
`

	result := parseGoSum(contents)

	assert.Equal(t, result, []string{"github.com/beorn7/perks@v1.0.0"})
}

func TestParseGoSumDuplicate(t *testing.T) {
	contents := `
github.com/konsorten/go-windows-terminal-sequences v1.0.1 h1:mweAR1A6xJ3oS2pRaGiHgQ4OO8tzTaLawm8vnODuwDk=
github.com/konsorten/go-windows-terminal-sequences v1.0.1/go.mod h1:T0+1ngSBFLxvqU3pZ+m/2kptfBszLMUkC4ZK/EgS/cQ=
github.com/konsorten/go-windows-terminal-sequences v1.0.3 h1:CE8S1cTafDpPvMhIxNJKvHsGVBgn1xWYf1NbHQhywc8=
github.com/konsorten/go-windows-terminal-sequences v1.0.3/go.mod h1:T0+1ngSBFLxvqU3pZ+m/2kptfBszLMUkC4ZK/EgS/cQ=
`

	result := parseGoSum(contents)

	assert.Equal(t, result, []string{"github.com/konsorten/go-windows-terminal-sequences@v1.0.3"})
}
