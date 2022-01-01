package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func main() {
	fmt.Println("start!")

	html := `<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta http-equiv="X-UA-Compatible" content="IE=edge">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Document</title>
	</head>
	<body>
		<h1>Index</h1>
		<p>This is sample web page.</p>
	</body>
	</html>`

	tf, er := template.New("index").Parse(html)
	if er != nil {
		log.Fatal(er)
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
