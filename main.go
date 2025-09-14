package main

import (
	"fmt"
	"html/template"
	"net/http"
)

const portNumber = ":8000"

func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.page.tmpl")
}

func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.page.tmpl")
}

func renderTemplate(w http.ResponseWriter, path string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + path)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("erro ao executar o template", err)
		return
	}
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	fmt.Println("Aplicação roddando na porta 8000...")
	_ = http.ListenAndServe(portNumber, nil)
}
