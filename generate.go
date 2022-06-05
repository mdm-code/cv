package main

import (
	"flag"
	"os"
	"text/template"

	"gopkg.in/yaml.v2"
)

var (
	dataFile  string
	templDir  string
	templBase string
)

func main() {
	fs := flag.NewFlagSet("generate.go", flag.ExitOnError)
	fs.StringVar(&dataFile, "dataFile", "data/@latest", "YAML file with resume data")
	fs.StringVar(&templDir, "templDir", "template/index.html", "template directory root")
	fs.StringVar(&templBase, "templBase", "base.html", "template base HTML file")
	fs.Parse(os.Args[1:])

	buf, err := os.ReadFile(dataFile)
	if err != nil {
		os.Exit(1)
	}

	data := map[string]interface{}{}

	err = yaml.Unmarshal(buf, data)
	if err != nil {
		os.Exit(2)
	}

	files, err := os.ReadDir(templDir)
	if err != nil {
		os.Exit(3)
	}

	fnames := []string{}

	for _, f := range files {
		fnames = append(fnames, templDir+"/"+f.Name())
	}

	templates := template.Must(template.New("").ParseFiles(fnames...))

	templates.ExecuteTemplate(os.Stdout, templBase, data)
}
