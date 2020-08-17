package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/veremchukvv/render"
)

var cfg = struct {
	Port     string
	UserAddr string
	WebAddr  string
}{
	Port:     "8083",
	UserAddr: "http://localhost:8080",
	WebAddr:  "http://localhost:8082",
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/checkout", checkoutFormHandler).Methods("GET")
	r.HandleFunc("/checkout", checkoutHandler).Methods("POST")
	render.SetTemplateDir(".")
	render.AddTemplate("payform", "payform.html")
	render.AddTemplate("msg", "msg.html")
	err := render.ParseTemplates()
	if err != nil {
		log.Fatalf("Init template error %v", err)
	}
	log.Printf("Starting on port %s", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, r))
}
