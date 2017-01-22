package main

import (
	"fmt"
	"net/http"

	"github.com/Nepooomuk/weatherservice-go/weather"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	http.HandleFunc("/weather", func(w http.ResponseWriter, r *http.Request) {
		report, err := weather.CreateForecast()
		fmt.Fprintln(w, report, err)
	})
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8080", nil)
}
