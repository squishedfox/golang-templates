package main

import (
	"net/http"
	"text/template"
)

var (
	t *template.Template
)

func defaultHandler(w http.ResponseWriter, req *http.Request) {
	templateClone, err := t.Clone()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	homeTemplate, err := templateClone.ParseFiles("home.gotmpl")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	w.WriteHeader(http.StatusOK)
	homeTemplate.ExecuteTemplate(w, "layout", &struct{}{})
}

func aboutHandler(w http.ResponseWriter, req *http.Request) {
	templateClone, err := t.Clone()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	homeTemplate, err := templateClone.ParseFiles("about.gotmpl")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	w.WriteHeader(http.StatusOK)
	homeTemplate.ExecuteTemplate(w, "layout", &struct{}{})
}

func init() {
	t = template.Must(template.ParseFiles("layout.gotmpl"))
}

func main() {
	webApp := http.NewServeMux()
	webApp.HandleFunc("/", defaultHandler)
	webApp.HandleFunc("/about", aboutHandler)

	http.ListenAndServe(":8080", webApp)
}
