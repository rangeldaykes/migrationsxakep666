package main

import (
	"context"
	"fmt"
	migrate "github.com/xakep666/mongo-migrate"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

const migrationsName = "migrations"

func main() {
	//MongoConnect("localhost", "", "", "testgeo")
	connect, err := MongoConnect("localhost", "testgeo")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(connect.Name())
}

func MongoConnect(host, database string) (*mongo.Database, error) {
	//uri := fmt.Sprintf("mongodb://%s:%s@%s:27017", user, password, host)
	uri := fmt.Sprintf("mongodb://%s:27017", host)
	opt := options.Client().ApplyURI(uri)
	client, err := mongo.NewClient(opt)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}
	db := client.Database(database)

	m := migrate.NewMigrate(
		db,
		MigAddMyIndex(),
		MigAddMyIndex2(),
	)

	////////////////////////////////////////////////////////////////////////////////////

	// use a filter to only select capped collections
	exist, err := isCollectionExist(migrationsName, *db, ctx)
	if err != nil {
		return nil, err
	}
	if !exist {
		m.SetVersion(0, "initial version")
	}

	////////////////////////////////////////////////////////////////////////////////////

	//migrate.Up(migrate.AllAvailable)

	if err := m.Up(migrate.AllAvailable); err != nil {
		return nil, err
	}

	return db, nil
}

func isCollectionExist(name string, db mongo.Database, ctx context.Context) (isExist bool, err error) {
	collections, err := db.ListCollectionNames(ctx, bson.D{})
	if err != nil {
		return false, err
	}

	for _, c := range collections {
		if name == c {
			return true, nil
		}
	}
	return false, nil
}

//func MigAddMyIndex() *migrate.Migration {
//
//	m := migrate.Migration{
//		Version: 1,
//		Description: "add my-index",
//
//		Up: func(db *mongo.Database) error {
//			opt := options.Index().SetName("my-index")
//			keys := bson.D{{"my-key", 1}}
//			model := mongo.IndexModel{Keys: keys, Options: opt}
//			_, err := db.Collection("my-coll").Indexes().CreateOne(context.TODO(), model)
//			if err != nil {
//				return err
//			}
//
//			return nil
//		},
//		Down: func(db *mongo.Database) error {
//			_, err := db.Collection("my-coll").Indexes().DropOne(context.TODO(), "my-index")
//			if err != nil {
//				return err
//			}
//			return nil
//		},
//	}
//
//	return &m
//}