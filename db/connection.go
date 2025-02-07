package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dsn = "host=localhost user=fede password=mysecretpassword dbname=gorm port=5432"
var DB *gorm.DB

func DBConnection() {
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("DB connected")
	}
}
