package main

import (
	"html/template"
	"log"
	"net/http"
)

type Todo struct {
	Id      int
	Message string
}

func main() {

	data := map[string][]Todo{
		"Todos": {
			Todo{Id: 1, Message: "Buy Ipad"},
		},
	}

	todosHandler := func(w http.ResponseWriter, r *http.Request) {
		temp1 := template.Must((template.ParseFiles("index.html")))

		temp1.Execute(w, data)

	}
	http.HandleFunc("/", todosHandler)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
