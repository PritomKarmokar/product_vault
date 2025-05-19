package main

import (
	"github.com/spf13/viper"
	"log"
)

func LoadConfig(path string) {
	viper.SetConfigFile(path)
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v\n", err)
	}
}
