package main

import (
	"go-api-sonar/api"
	"net/http"
)

func main() {
	http.HandleFunc("/add", api.MathHandler)
	http.ListenAndServe(":8080", nil)
}
