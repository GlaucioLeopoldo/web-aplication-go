package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/GlaucioLeopoldo/web-aplication-go/pkg/config"
	"github.com/GlaucioLeopoldo/web-aplication-go/pkg/handler"
	"github.com/GlaucioLeopoldo/web-aplication-go/pkg/render"
)

const portNumber = ":8000"

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Error")
	}

	app.TemplateCache = tc
	repo := handler.NewRepo(&app)

	handler.NewHandlers(repo)
	app.UseCache = false
	render.NewTemplate(&app)

	http.HandleFunc("/", handler.Repo.Home)
	http.HandleFunc("/about", handler.Repo.About)
	fmt.Println("Aplicação roddando na porta 8000...")
	_ = http.ListenAndServe(portNumber, nil)
}
