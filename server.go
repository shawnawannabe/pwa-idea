package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful")
	// name := r.FormValue("name")
	// address := r.FormValue("address")

	// fmt.Fprintf(w, "Name = %s\n", name)
	// fmt.Fprintf(w, "Address = %s\n", address)

}

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Failed to load .env")
	}
	fmt.Printf("Server started at port 8080\n")

	fileServer := http.FileServer((http.Dir(os.Getenv("STATIC_DIR"))))
	http.Handle("/", fileServer)
	http.HandleFunc("/", formHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
