package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	// tpl is scontainer which contain files
	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	nf, err := os.Create("index.html")
	if err != nil {
		log.Println("error creating the desired file", err)
	}
	defer nf.Close()
	err = tpl.Execute(nf, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
