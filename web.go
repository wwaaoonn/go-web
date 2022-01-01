package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func main() {
	fmt.Println("start!")

	tf, er := template.ParseFiles("templates/hello.html")
	if er != nil {
		tf, _ = template.New("index").Parse("<html><body><h1>Sorry, Not found...</h1></body></html>")
	}
	hh := func(w http.ResponseWriter, r *http.Request) {
		er = tf.Execute(w, nil)
		if er != nil {
			log.Fatal(er)
		}
	}

	http.HandleFunc("/hello", hh)

	http.ListenAndServe(":8080", nil)
}
