package main

import (
	"fmt"
	"net/http"
)

var portNumber = ":8000"

func AboutPage(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "about page")
	fmt.Println("Endpoint Hit: AboutPage")
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "home page")
	fmt.Println("Endpoint Hit: HomePage")
}

func main() {
	fmt.Println("the port number is ", portNumber)
	http.HandleFunc("/", HomePage)
	http.HandleFunc("/about", AboutPage)
	_ = http.ListenAndServe(portNumber, nil)
}
