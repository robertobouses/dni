package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/robertobouses/dni/gestor"
)

type DNI struct {
	Numero int    `json:"numero"`
	Letra  string `json:"letra"`
	Nombre string `json:"nombre"`
}

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("postgres", "postgres://usuario:contraseña@localhost/nombre_de_la_base_de_datos?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()

	r.HandleFunc("/dni", gestor.GetAllDNI).Methods("GET")
	r.HandleFunc("/dni", gestor.CreateDNI).Methods("POST")
	r.HandleFunc("/dni/{letra}", gestor.FilterDNI).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}
