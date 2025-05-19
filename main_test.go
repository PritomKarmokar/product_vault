package main_test

import (
	"github.com/PritomKarmokar/product_vault"
	"github.com/spf13/viper"
	"log"
	"os"
	"testing"
)

var a main.App

func TestMain(m *testing.M) {
	loadEnv()

	a.Initialize(
		viper.GetString("DB_USERNAME"),
		viper.GetString("DB_PASSWORD"),
		viper.GetString("DB_NAME"),
	)

	ensureTableExists()
	code := m.Run()
	clearTable()
	os.Exit(code)
}

func loadEnv() {
	viper.SetConfigFile(".env")
	//viper.AddConfigPath(".")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v\n", err)
	}
}

func ensureTableExists() {
	if _, err := a.DB.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}

func clearTable() {
	a.DB.Exec("DELETE FROM products")
	a.DB.Exec("ALTER SEQUENCE products_id_seq RESTART WITH 1")
}

const tableCreationQuery = `CREATE TABLE IF NOT EXISTS products 
(
    id SERIAL,
    name TEXT NOT NULL,
    price NUMERIC(10,2) NOT NULL DEFAULT 0.00,
    CONSTRAINT products_pkey PRIMARY KEY (id)
)`
