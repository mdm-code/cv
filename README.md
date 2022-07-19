# Michał Adamczyk's Curriculum Vitae

Here is my
![@latest](https://github.com/mdm-code/cv/blob/main/michal-adamczyk.pdf)
resume. It is derived from data stored in the
![@data](https://github.com/mdm-code/cv/blob/main/data/data.yaml) file that
keeps all relevant information about my professional life in one place. The
structured data format of a `yaml` flat file makes it easy to browse through
and filter for relevant information.


## Build

1. Use `--help` to check available command-line arguments for `generate.go`.

```sh
go run generate.go --help
```

2. Utilize `Makefile` default goal to compile `index.html` with `make`.


3. Use `generate.go` explicitly:

```sh
go run generate.go -dataFile data/@latest -templBase base.html --templDir template/index.html > index.html
```

4. Specify `-firm` in case you want RODO clause to be added at the end of your CV.

```sh
go run generage -firm FooBar > index.html
```

The templating engine will check if the `.Firm` variable is not empty and in
case it isn't, It will pop the whole RODO clause alongside the company name at
the bottom of the document.

5. Use `Firefox` to open `index.html` and save the page as a PDF file.
