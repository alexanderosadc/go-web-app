package main

import (
	"log"
	"net/http"

	"github.com/alexanderosadc/go-web-app/pkg/config"
	"github.com/alexanderosadc/go-web-app/pkg/handlers"
	"github.com/alexanderosadc/go-web-app/pkg/render"
)

const portNumber = ":8080"

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	_ = http.ListenAndServe(portNumber, nil)
}
