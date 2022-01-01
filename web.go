package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("start!")

	hh := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, This is GO-server!!"))
	}

	http.HandleFunc("/hello", hh)

	http.ListenAndServe(":8080", nil)
}
