package main

import (
    "fmt"
    "log"
    "net/http"

    "github.com/gorilla/mux"
    "backend/config"
    "backend/handlers"
)

func main() {
    // Load application configuration
    err := config.LoadConfig()
    if err != nil {
        log.Fatalf("Failed to load config: %v", err)
    }

    // Initialize database connection
    err = config.InitDB()
    if err != nil {
        log.Fatalf("Failed to initialize database: %v", err)
    }
    defer config.DB.Close()

    // Initialize router
    router := mux.NewRouter()

    // Define routes
    router.HandleFunc("/api/auth/register", handlers.Register).Methods("POST")
    router.HandleFunc("/api/auth/login", handlers.Login).Methods("POST")
    router.HandleFunc("/api/auth/logout", handlers.Logout).Methods("GET")

    router.HandleFunc("/api/tasks", handlers.GetTasks).Methods("GET")
    router.HandleFunc("/api/tasks/{id}", handlers.GetTask).Methods("GET")
    router.HandleFunc("/api/tasks", handlers.CreateTask).Methods("POST")
    router.HandleFunc("/api/tasks/{id}", handlers.UpdateTask).Methods("PUT")
    router.HandleFunc("/api/tasks/{id}", handlers.DeleteTask).Methods("DELETE")

    // Start the server
    port := config.AppConfig.Port
    fmt.Printf("Server is running on port %s...\n", port)
    log.Fatal(http.ListenAndServe(":"+port, router))
}
