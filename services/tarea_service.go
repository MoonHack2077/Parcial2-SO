package services

import (
	"github.com/MoonHack2077/Parcial2-SO/models"
	"github.com/MoonHack2077/Parcial2-SO/repositories"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CrearTarea(tarea models.Tarea) (primitive.ObjectID, error) {
	return repositories.CrearTarea(tarea)
}

func ObtenerTareas() ([]models.Tarea, error) {
	return repositories.ObtenerTodasTareas()
}
