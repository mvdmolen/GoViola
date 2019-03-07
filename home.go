package main

import (
	"net/http"
)

//handler for / renders the home.html
func Home(w http.ResponseWriter, req *http.Request) {
	pageVars := PageVars{
		Title: "GoViola",
	}
	render(w, "home.html", pageVars)
}
