package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8000"

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the home page")
}

func About(w http.ResponseWriter, r *http.Request) {
	sum := sumatori(2, 2)
	_, _ = fmt.Fprintf(w, "%s", fmt.Sprintf("This is the about page and 2 + 2 = %d", sum))
}

func sumatori(a, b int) int {
	return a + b
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	fmt.Println("Aplicação roddando na porta 8000...")
	_ = http.ListenAndServe(portNumber, nil)
}
