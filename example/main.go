package main

import (
	"context"
	"log"

	"github.com/arikfr/godbsql"
)

func main() {
	ctx := context.Background()
	configuration := godbsql.NewConnectConfiguration()
	configuration.HTTPPath = "sql/1.0/endpoints/5e89f447c123a5s8" // this is the default path in hive configuration.
	configuration.Token = "dapi..."

	connection, errConn := godbsql.Connect("demo.cloud.databricks.com", configuration)

	log.Println("Connecting")
	if errConn != nil {
		log.Fatal(errConn)
	}
	log.Println("Cursor")
	cursor := connection.Cursor()

	cursor.Exec(ctx, "SELECT 1")
	if cursor.Err != nil {
		log.Fatal(cursor.Err)
	}

	var i int32
	for cursor.HasMore(ctx) {
		cursor.FetchOne(ctx, &i)
		if cursor.Err != nil {
			log.Fatal(cursor.Err)
		}
		log.Println(i)
	}

	cursor.Close()
	connection.Close()
}
