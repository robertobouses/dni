package gestor

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

func GetAllDNI(w http.ResponseWriter, r *http.Request) {
	var dniList []DNI

	db, err := sql.Open("postgres", "postgres://usuario:contraseña@localhost/nombre_de_la_base_de_datos?sslmode=disable")

	rows, err := db.Query("SELECT numero, letra, nombre FROM dni")
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
