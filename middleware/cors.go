package middleware

import (
    "net/http"
)

// CorsMiddleware versi sementara: izinkan semua origin
func CorsMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

        // Izinkan semua origin
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, ngrok-skip-browser-warning")
        w.Header().Set("Access-Control-Allow-Credentials", "true")

        // Preflight OPTIONS langsung OK
        if r.Method == http.MethodOptions {
            w.WriteHeader(http.StatusOK)
            return
        }

        // Lanjut ke handler berikutnya
        next.ServeHTTP(w, r)
    })
}
