package gestor

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

type DNI struct {
	Numero int    `json:"numero"`
	Letra  string `json:"letra"`
	Nombre string `json:"nombre"`
}

func CreateDNI(w http.ResponseWriter, r *http.Request) {

	var dni DNI

	db, err := sql.Open("postgres", "postgres://usuario:contraseña@localhost/nombre_de_la_base_de_datos?sslmode=disable")

	err = json.NewDecoder(r.Body).Decode(&dni)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if !isValidDNI(dni.Numero, dni.Letra) {
		http.Error(w, "La letra del DNI no es válida", http.StatusBadRequest)
		return
	}

	_, err = db.Exec("INSERT INTO dni (numero, letra, nombre) VALUES ($1, $2, $3)", dni.Numero, dni.Letra, dni.Nombre)
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusCreated)
}
