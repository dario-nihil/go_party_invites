package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type RSVP struct {
	Name, Email, Phone string
	WillAttend         bool
}

var responses = make([]*RSVP, 0, 10)
var templates = make(map[string]*template.Template, 3)

func main() {
	loadTemplates()

	// specifies the URL path and the handler
	http.HandleFunc("/", welcomeHandler)
	http.HandleFunc("/list", listHandler)

	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		fmt.Println(err)
	}
}

func loadTemplates() {
	templateNames := [5]string{"welcome", "form", "thanks", "sorry", "list"}
	for idx, name := range templateNames {
		t, err := template.ParseFiles("layout.html", name+".html")
		if err == nil {
			templates[name] = t
			fmt.Println("Loaded template", idx, name)
		} else {
			panic(err)
		}
	}
}

func welcomeHandler(writer http.ResponseWriter, request *http.Request) {
	templates["welcome"].Execute(writer, nil)
}

func listHandler(writer http.ResponseWriter, request *http.Request) {
	templates["list"].Execute(writer, responses)
}
