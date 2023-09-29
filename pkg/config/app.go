package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB // This will help the other files to interact with the database
)

func Connect() {

	// Gorm always let us using postgresql, sqlite or sqlserver

	dsn := "jose:1234@tcp(127.0.0.1:3306)/simplerest?charset=utf8mb4&parseTime=True&loc=Local"

	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db = d

}

func GetDB() *gorm.DB {
	return db
}
