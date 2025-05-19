package main

import (
	"github.com/spf13/viper"
)

func main() {
	a := App{}
	LoadConfig(".env")
	a.Initialize(
		viper.GetString("DB_USERNAME"),
		viper.GetString("DB_PASSWORD"),
		viper.GetString("DB_NAME"),
	)
	a.Run(":8080")
}

//https://semaphore.io/community/tutorials/building-and-testing-a-rest-api-in-go-with-gorilla-mux-and-postgresql
