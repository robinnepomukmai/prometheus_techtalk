package main

import (
	"fmt"
	"net/http"

	"github.com/Nepooomuk/weatherservice-go/weather"
)

func main() {
	http.HandleFunc("/weather", func(w http.ResponseWriter, r *http.Request) {
		report, err := weather.CreateForecast()
		fmt.Fprintln(w, report, err)
	})

	http.ListenAndServe(":8080", nil)
}
