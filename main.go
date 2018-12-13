package main

import (
	"fmt"
	"net/http"
	"text/template"
)

type Inventory struct {
	Material string
	Count int
}

func main() {
	thing := Inventory{
		Material: "wood",
		Count: 4,
	}

	tmpl, err := template.ParseFiles("niks.template")
	if err != nil{
		fmt.Println(err.Error())
	}

	 http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		 tmpl.Execute(w, thing)
	 })

	http.ListenAndServe(":8080", nil)
}