# Michał Adamczyk's Curriculum Vitae

Here is my ![@latest](github.com/mdm-code/cv/michal-adamczyk.pdf) resume. It
is derived from data stored in the ![@data](github.com/mdm-code/cv/data/@latest)
file that keeps all relevant information about my professional life in one place.
The structured data format of a `yaml` flat file makes it easy to browse through
and filter for relevant information.


## Build

1. Use `--help` to check available command-line arguments for `generate.go`.

```sh
go run generage.go --help
```

2. Utilize `Makefile` default goal to compile `index.html` with `make`.


3. Use `generate.go` explicitly:

```sh
go run generate.go -dataFile data/@latest -templBase base.html --templDir template/index.html > index.html
```

4. Use `Firefox` to open `index.html` and save the page as a PDF file.
