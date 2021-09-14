package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"path"
)

var (
	basePath = "D:\\Arkham horror"
	trPath   = path.Join(basePath, "static", "translated", "archive.json")
	tmplPath = path.Join(basePath, "static", "template", "archive-karten.html")
)

type archiveKarte struct {
	Id    int      `json:"id"`
	Title string   `json:"title"`
	Front []string `json:"front"`
	Back  []string `json:"back"`
}

type archiveKarten struct {
	ArchiveKarten []archiveKarte `json:"archiveKarten"`
}

func jsonToStruct() (data archiveKarten) {
	file, err := ioutil.ReadFile(trPath)
	if err != nil {
		fmt.Printf("%v\n", "file error")
		panic(err)
	}
	data = archiveKarten{}
	err = json.Unmarshal([]byte(file), &data)
	if err != nil {
		fmt.Printf("%v\n", "json error")
		panic(err)
	}
	return data
}

func genHtml() string {
	var res bytes.Buffer
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(&res, jsonToStruct())
	return res.String()
}

func write(content, path string) {
	err := os.WriteFile(path, []byte(content), 0644)
	if err != nil {
		panic(err)
	}
}

func main() {
	write(genHtml(), "output.html")
}
