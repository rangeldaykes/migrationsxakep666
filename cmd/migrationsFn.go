package main

import (
	"context"
	migrate "github.com/xakep666/mongo-migrate"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MigAddMyIndex() migrate.Migration {
	m := migrate.Migration{
		Version: 1,
		Description: "add my-index",

		Up: func(db *mongo.Database) error {
			opt := options.Index().SetName("my-index")
			keys := bson.D{{"my-key", 1}}
			model := mongo.IndexModel{Keys: keys, Options: opt}
			_, err := db.Collection("my-coll").Indexes().CreateOne(context.TODO(), model)
			if err != nil {
				return err
			}

			return nil
		},
		Down: func(db *mongo.Database) error {
			_, err := db.Collection("my-coll").Indexes().DropOne(context.TODO(), "my-index")
			if err != nil {
				return err
			}
			return nil
		},
	}

	return m
}

func MigAddMyIndex2() migrate.Migration {
	m := migrate.Migration{
		Version: 2,
		Description: "add my-index2",

		Up: func(db *mongo.Database) error {
			opt := options.Index().SetName("my-index2")
			keys := bson.D{{"my-key2", 1}}
			model := mongo.IndexModel{Keys: keys, Options: opt}
			_, err := db.Collection("my-coll2").Indexes().CreateOne(context.TODO(), model)
			if err != nil {
				return err
			}

			return nil
		},
		Down: func(db *mongo.Database) error {
			_, err := db.Collection("my-coll").Indexes().DropOne(context.TODO(), "my-index")
			if err != nil {
				return err
			}
			return nil
		},
	}

	return m
}