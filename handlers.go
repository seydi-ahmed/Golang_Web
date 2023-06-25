package main

import (
	"net/http"
	"text/template"
)

func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home")
}

func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "About")
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	t, err := template.ParseFiles("./template/" + tmpl + ".page.tmpl")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, nil)
}

func createtemplateCache() (map[string]*template.Template, error) {

}
