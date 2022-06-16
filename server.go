package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

type Petrol struct {
	ID     string `json:"id"`
	Amount string `json:"amount"`
}

var petrols []Petrol

func getPetrol(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(petrols)
}

func formHandler(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful")
	name := r.FormValue("name")
	address := r.FormValue("address")

	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)

}

func main() {
	//hardcode for now
	petrols = append(petrols, Petrol{ID: "1", Amount: "60L"})

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Failed to load .env")
	}
	fmt.Printf("Server started at port 8080\n")

	fileServer := http.FileServer((http.Dir(os.Getenv("STATIC_DIR"))))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)

	//router
	router := mux.NewRouter()
	router.HandleFunc("/form/petrol", getPetrol).Methods("Get")
	log.Fatal(http.ListenAndServe(":8080", router))
}
