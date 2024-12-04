package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(mysql.Open("root:05072001@tcp(localhost:3306)/gin_crud_mysql"))
	if err != nil {
		panic(err)
	}

	//auto migrate
	database.AutoMigrate(&Product{})

	DB = database
}
