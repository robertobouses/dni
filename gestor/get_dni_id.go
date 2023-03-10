package gestor

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func FilterDNI(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	letra := strings.ToUpper(params["letra"])

	var dniList []DNI

	db, err := sql.Open("postgres", "postgres://usuario:contraseña@localhost/nombre_de_la_base_de_datos?sslmode=disable")

	rows, err := db.Query("SELECT numero, letra, nombre FROM dni WHERE letra = $1", letra)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var dni DNI
		err := rows.Scan(&dni.Numero, &dni.Letra, &dni.Nombre)
		if err != nil {
			log.Fatal(err)
		}
		dniList = append(dniList, dni)
	}

	json.NewEncoder(w).Encode(dniList)
}

func isValidDNI(numero int, letra string) bool {
	const letras = "TRWAGMYFPDXBNJZSQVHLCKE"
	calculoLetra := letras[numero%23]
	return strings.ToUpper(letra) == string(calculoLetra)
}
