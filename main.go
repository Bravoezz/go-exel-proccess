package main

import (
	"fmt"
	"log"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func main() {
	startTime := time.Now()

	const exelFileName = "lista_clientes.xlsx"
	file, err := excelize.OpenFile(exelFileName)
	if err != nil {
		fmt.Println("Error")
		log.Fatal(err)
	}

	db, err := DBConnection()
	errorPing := db.Ping()
	if errorPing != nil || err != nil  {
		fmt.Println("Error al iniciar la coneccion", err.Error())
	}
	defer DBClose(db)

	sheetName := file.GetSheetMap()[1]

	query := `
		INSERT INTO clientes
		(nombre_completo,fecha_nacimiento,direccion,localidad,telefono,email,fecha_alta,grupo)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8)
	`
	for i, row := range file.GetRows(sheetName) {
		if i == 0 {
			continue
		}
		// if i == 3 {
		// 	break
		// }
		
		fmt.Println("data: ",row, " index: ", i)
		_, err := db.Exec(query,row[1],row[2],row[3],row[4],row[5],row[6],row[7],row[8])
		if err != nil {
			fmt.Println("error en el query", err.Error())
		}
	}


	endTime := time.Now()
	totalDuration := endTime.Sub(startTime)
	fmt.Printf("Total duration: %v \n", totalDuration)
}
