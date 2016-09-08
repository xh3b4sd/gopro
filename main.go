package main

//go:generate ${GOPATH}/bin/loader generate -i ./template/

import (
	"bytes"
	"go/format"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type TemplateContext struct {
	BinaryName string
}

func main() {
	if len(os.Args) != 2 {
		panic("expect project name as first argument")
	}

	newTmplCtx := TemplateContext{
		BinaryName: os.Args[1],
	}

	err := os.Mkdir("./"+newTmplCtx.BinaryName+"/", os.FileMode(0744))
	if err != nil {
		panic(err)
	}

	for _, fileName := range LoaderFileNames() {
		rawTemplate, err := LoaderReadFile(fileName)
		if err != nil {
			panic(err)
		}

		tmpl, err := template.New(filepath.Base(fileName)).Parse(string(rawTemplate))
		if err != nil {
			panic(err)
		}
		var b bytes.Buffer
		err = tmpl.Execute(&b, newTmplCtx)
		if err != nil {
			panic(err)
		}

		// format
		rawFile := b.Bytes()
		if strings.HasSuffix(filepath.Base(fileName), ".go.tmpl") {
			rawFile, err = format.Source(rawFile)
			if err != nil {
				panic(err)
			}
		}

		// write
		newFileName := "./" + newTmplCtx.BinaryName + "/" + strings.Replace(filepath.Base(fileName), ".tmpl", "", -1)
		err = ioutil.WriteFile(newFileName, rawFile, os.FileMode(0644))
		if err != nil {
			panic(err)
		}
	}
}
