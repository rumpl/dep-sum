all: cli

cli:
	go build .

test:
	go test ./...

.PHONY: cli test
