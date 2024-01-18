package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func DBConnection() (*sql.DB,error) {
	const conSrt = "user=postgres password=toor dbname=go_exel sslmode=disable"
	DBCon, err := sql.Open("postgres",conSrt)
	if err != nil {
		return nil, err
	}
	return DBCon,nil
}

func DBClose(db *sql.DB) {
	err := db.Close()
	if err != nil {
		fmt.Println("Error al cerrar la coneccion: ", err.Error())
	}
}