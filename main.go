package main

import (
    "fmt"
    "log"
    "net/http"
    "os"

    "github.com/joho/godotenv"
    "reset/config"
    "reset/routes"
    "reset/middleware"
)

func main() {
    // Load env
    errEnv := godotenv.Load()
    if errEnv != nil {
        log.Println("Warning: .env file not found, using environment variables")
    }

    // Gunakan PORT dari Railway, fallback ke 8080
    port := os.Getenv("PORT")
    if port == "" {
        port = os.Getenv("APP_PORT")
        if port == "" {
            port = "8080"
        }
    }
    fmt.Println("Server running on port " + port)

    // Connect ke database
    db, err := config.ConnectToDatabase()
    if err != nil {
        log.Fatal("Database connection failed:", err)
    }

    // Setup router + middleware CORS
    router := routes.SetupRoutes(db) // Pastikan routes.Routes/SetupRoutes mengembalikan *http.ServeMux atau router
    handler := middleware.CorsMiddleware(router)

    // Jalankan server
    log.Fatal(http.ListenAndServe(":"+port, handler))
}
