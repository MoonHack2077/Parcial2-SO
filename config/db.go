package config

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func ConectarDB() {
	mongoURI := os.Getenv("MONGO_URI")
	dbName := os.Getenv("MONGO_DB")
	fmt.Println("MONGO_URI:", mongoURI)
fmt.Println("DB Name:", dbName)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		panic(err)
	}

	DB = client.Database(dbName)
	fmt.Println("✅ Conectado a MongoDB")
}

func GetCollection(nombre string) *mongo.Collection {
	return DB.Collection(nombre)
}
