package main

import (
	"fmt"
	"html/template"
)

type RSVP struct {
	Name, Email, Phone string
	WillAttend         bool
}

var responses = make([]*RSVP, 0, 10)
var templates = make(map[string]*template.Template, 3)

func main() {
	fmt.Println("TODO: add some features")
	loadTemplates()
}

func loadTemplates() {
	//TODO - load templates here
}
