package render

import (
	"fmt"
	"html/template"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")
	if err := parsedTemplate.Execute(w, nil); err != nil {
		fmt.Println("error parsing template:", err)
	}
	return
}
