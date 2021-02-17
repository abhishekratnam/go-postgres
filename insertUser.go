package data

import (
	"database/sql"
	"fmt"
	"log"
	"microservice/models"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "jarvis"
	dbname   = "abhi" //ALTER USER postgres PASSWORD 'jarvis';
)

func createConnection() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")
	return db
}

func Insertuser(user models.Product) int64 {
	db := createConnection()
	defer db.Close()
	sqlStatement := `INSERT INTO Products (        
		Name,        
		Description ,
		Price,       
		SKU) VALUES ($1, $2, $3, $4) RETURNING id`
	var id int64
	err := db.QueryRow(sqlStatement, user.Name, user.Description, user.Price, user.SKU).Scan(&id)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}
	fmt.Printf("Inserted a single record %v", id)
	return id
}
