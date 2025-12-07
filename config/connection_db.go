package config

import (
    "database/sql"
    "fmt"
    "log"
    "os"
    "time"

    _ "github.com/go-sql-driver/mysql"
    "github.com/joho/godotenv"
)

func ConnectToDatabase() (*sql.DB, error) {
    // Load .env (ignored in Railway, but needed locally)
    err := godotenv.Load()
    if err != nil {
        log.Println("Warning: .env not found or failed to load:", err)
    }

    dbName := os.Getenv("DB_NAME")
    dbUser := os.Getenv("DB_USER")
    dbPass := os.Getenv("DB_PASS")
    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")

    mysqlConn := fmt.Sprintf(
        "%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local",
        dbUser, dbPass, dbHost, dbPort, dbName,
    )

    db, err := sql.Open("mysql", mysqlConn)
    if err != nil {
        log.Println("Error opening DB connection:", err)
        return nil, err
    }

    // Test connection
    err = db.Ping()
    if err != nil {
        log.Println("Error pinging DB:", err)
        return nil, err
    }

    db.SetConnMaxLifetime(time.Minute * 3)
    db.SetMaxOpenConns(10)
    db.SetMaxIdleConns(10)

    log.Println("Database connected successfully")
    return db, nil
}
