package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type PageVariables struct {
	Height float64
	Weight float64
	Bmi    float64
}

func start(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("mainpage.html"))
	height, err := strconv.ParseFloat(r.FormValue("Height"), 64)
	weight, err := strconv.ParseFloat(r.FormValue("Weight"), 64)
	if err != nil {
		log.Fatal(err)
	}
	bmi := weight * 10000 / height * height
	HomePageVars := &PageVariables{height, weight, bmi}
	t.Execute(w, HomePageVars)
}

func main() {
	http.HandleFunc("/", start)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
