package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectToDatabase() (db *sql.DB, err error) {

	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	mysql := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local",
		dbUser,
		dbPass,
		dbHost,
		dbPort,
		dbName,
	)

	db, err = sql.Open("mysql", mysql)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Println("Error pinging DB:", err)
		return nil, err
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db, nil
}
