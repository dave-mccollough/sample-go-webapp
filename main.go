package main

import (
	"fmt"
	"net/http"
)

func funcHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Sample Go webapp on Azure App Service!")
}

func main() {
	http.HandleFunc("/", funcHandler)
	http.ListenAndServe(":3000", nil)
}
