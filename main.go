package main

import (
	"countryAPI/handlers"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("server running on port :8080")

	http.HandleFunc("/viloyatlar", handlers.RegionsHandler)
	http.HandleFunc("/shaxarlar", handlers.CitiesHandler)
	http.HandleFunc("/joylar", handlers.PlacesHandler)

	http.ListenAndServe(":8080", nil)
}
