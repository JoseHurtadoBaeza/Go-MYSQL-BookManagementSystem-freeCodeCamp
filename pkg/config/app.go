package config

import "gorm.io/gorm"

var (
	db *gorm.DB // This will help the other files to interact with the database
)

func Connect() {

	d, err := gorm.Open("mysql", "root:1234/simplerest?charset=utf8&parseTime=True&loc=Local") // Gorm always let us using postgresql or sqlite
	if err != nil {
		panic(err)
	}
	db = d

}

func GetDB() *gorm.DB {
	return db
}
