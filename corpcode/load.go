package corpcode

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func inserMany(corps []Corp) {

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://94.237.67.29:27017"))
	if err != nil {
		fmt.Println(err)
	}

	ctx, _ = context.WithTimeout(context.Background(), 2*time.Second)
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		fmt.Println(err)
	}

	collection := client.Database("disclosur").Collection("corpcode")

	ctx, _ = context.WithTimeout(context.Background(), 600*time.Second)
	/*
		https://github.com/golang/go/wiki/InterfaceSlice#what-can-i-do-instead
	*/
	var interfaceSlice []interface{} = make([]interface{}, len(corps))
	for i, d := range corps {
		interfaceSlice[i] = d
	}

	_, err = collection.InsertMany(ctx, interfaceSlice)
	if err != nil {
		fmt.Println(err)
	}

}

func cumulate(corps []Corp) {
	fmt.Println(len(corps))
	// inserMany(corps)
}
