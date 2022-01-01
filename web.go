package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("start!")

	http.ListenAndServe(":8080", http.NotFoundHandler())
}
