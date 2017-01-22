package main

import "net/http"
import "fmt"

func main() {
	healthcheck, err := http.Get("http://slash.service.dev.rewe-digital.com/metrics")
	fmt.Println(healthcheck, err)
	/*healthcheck, err := http.Get("http://slash.service.dev.rewe-digital.com/admin/healthcheck")
	fmt.Println(healthcheck, err)
	health, err := http.Get("http://slash.service.dev.rewe-digital.com/admin/health")
	fmt.Println(health, err)*/
}
