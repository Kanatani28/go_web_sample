package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // Use MySQL in gorm

	"../entity"
)

var (
	db  *gorm.DB
	err error
)

// Init is initialize db from main function
func Init() {
	// db, err = gorm.Open("postgres", "host=0.0.0.0 port=5432 user=gorm dbname=gorm password=gorm sslmode=disable")
	db, err = gorm.Open("mysql", "root:root@/portal?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	autoMigration()
}

// GetDB is called in models
func GetDB() *gorm.DB {
	return db
}

func autoMigration() {
	db.AutoMigrate(&entity.User{})
}
