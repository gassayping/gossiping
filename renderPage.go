package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"strings"
	"text/template"
)

var templatesDir = os.DirFS("./templates")
var serveDir = os.DirFS("./srv")

type PageData struct {
	Title string
	Body  string
}

func renderPage(path string) (string, error) {
	templ, err := template.New("container").ParseFS(templatesDir, "container.html")
	if err != nil {
		return "", errors.New("Error reading container file: " + err.Error())
	}
	file, err := fs.ReadFile(serveDir, path)
	if err != nil {
		return "", err
	}
	b := new(strings.Builder)
	fmt.Println(templ.Templates()[0].Name())
	err = templ.ExecuteTemplate(b, "container.html", PageData{
		Title: "Gassayping",
		Body:  string(file),
	})
	if err != nil {
		return "", err
	}
	return b.String(), nil
}
