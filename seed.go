package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
	"strconv"
)

var connStr = "postgresql://postgres:root@localhost/postgres?sslmode=disable"

type FifaRecord struct {
	name     string
	pos      string
	club     string
	league   string
	nation   string
	height   int
	weight   int
	age      int
	foot     string
	best_pos string
	value    int
	wage     int
}

func getDataFromCSV() []FifaRecord {
	file, err := os.Open("./dataset/datafm20.csv")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	// read csv values using csv.Reader
	csvReader := csv.NewReader(file)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	var allRecords []FifaRecord

	for i, line := range data {
		if i > 0 {
			int_height, _ := strconv.Atoi(line[5])
			int_weight, _ := strconv.Atoi(line[6])
			int_age, _ := strconv.Atoi(line[7])
			int_value, _ := strconv.Atoi(line[10])
			int_wage, _ := strconv.Atoi(line[11])

			entry := FifaRecord{
				line[0],
				line[1],
				line[2],
				line[3],
				line[4],
				int_height,
				int_weight,
				int_age,
				line[8],
				line[9],
				int_value,
				int_wage,
			}
			allRecords = append(allRecords, entry)
		}
	}
	return allRecords
}

func seedData() {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	data := getDataFromCSV()
	fmt.Println("SQL ----- Migrating Data\n")
	for _, entry := range data {
		_, err := db.Exec("INSERT INTO fifa VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)",
			entry.name,
			entry.pos,
			entry.club,
			entry.nation,
			entry.height,
			entry.weight,
			entry.age,
			entry.foot,
			entry.best_pos,
			entry.value,
			entry.wage,
		)

		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("SQL ----- Migration Done")
}

func main() {
	seedData()
}
