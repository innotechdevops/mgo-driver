package mgodriver

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"log"
	"os"
	"time"
)

const DefaultPort = "27017"

// Config is a model for connect MongoDB
type Config struct {
	User         string
	Pass         string
	Host         string
	DatabaseName string
	Port         string
}

// MongoDriver is the interface
type MongoDriver interface {
	Connect() *mongo.Database
}

type mongoDB struct {
	Conf Config
}

func (m *mongoDB) Connect() *mongo.Database {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s@%s:%s", m.Conf.User, m.Conf.Pass, m.Conf.Host, m.Conf.Port)))
	if err != nil {
		log.Fatalf("MongoDb Connection fail : %s", err)
	}
	log.Print("MongoDb Connected.")
	return client.Database(m.Conf.DatabaseName)
}

// New for create mongodb driver
func New(config Config) MongoDriver {
	return &mongoDB{
		Conf: config,
	}
}

// ConfigEnv for create mongodb driver
func ConfigEnv() Config {
	return Config{
		User:         os.Getenv("MONGO_USER"),
		Pass:         os.Getenv("MONGO_PASS"),
		Host:         os.Getenv("MONGO_HOST"),
		DatabaseName: os.Getenv("MONGO_DATABASE"),
		Port:         os.Getenv("MONGO_PORT"),
	}
}
