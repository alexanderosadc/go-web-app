package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

// RenderTemplate renders a template
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	// creates template cache
	tc, err := CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}
	// gets requested template from cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal(err)
	}

	buf := new(bytes.Buffer)

	if err = t.Execute(buf, nil); err != nil {
		log.Println(err)
	}

	//renders the template
	if _, err = buf.WriteTo(w); err != nil {
		log.Println(err)
	}

}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	matches, err := filepath.Glob("./templates/*.layout.tmpl")
	if err != nil {
		return myCache, nil
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, nil
			}
		}

		myCache[name] = ts
	}

	return myCache, nil
}
