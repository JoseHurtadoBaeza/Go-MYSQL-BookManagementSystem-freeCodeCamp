package config

import "gorm.io/gorm"

var (
	db *gorm.DB // This will help the other files to interact with the database
)

func Connect() {

	d, err := gorm.Open("mysql", "root:1234/simplerest?") // Gorm always let us using postgresql or sqlite
	if err != nil {
		panic(err)
	}
	db = d

}
