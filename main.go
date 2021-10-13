package main
import (
	"log"
	"context"
	"time"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson"
)
func main() {
	clientOptions := options.Client().ApplyURI("mongodb+srv://seke:upsi1234@cluster0.zbbel.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
    	log.Fatal(err)
	}
	databases, err := client.ListDatabaseNames(ctx, bson.D{})
	if err != nil {
    	log.Fatal(err)
	}
	fmt.Println(databases)
}



