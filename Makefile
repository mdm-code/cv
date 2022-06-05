GO=go
OUTPUT=index.html

DEFAULT_GOAL := run

run:
	$(GO) run generate.go > $(OUTPUT)
.PHONY: run
