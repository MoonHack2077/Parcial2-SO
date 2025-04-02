package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/gorilla/mux"
	"github.com/MoonHack2077/Parcial2-SO/config"
)

func main() {
	// Cargar .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("❌ Error cargando .env")
	}

	// Conectar a MongoDB
	config.ConectarDB()

	// Crear router
	router := mux.NewRouter()

	// Aquí luego registraremos los endpoints
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Microservicio de Tareas 🚀")
	}).Methods("GET")

	// Correr servidor
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("🟢 Servidor corriendo en el puerto " + port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
