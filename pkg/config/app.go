package config

import (
	//"gorm.io/driver/mysql"
	//"gorm.io/gorm"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	//dsn := "root:root@tcp(127.0.0.1:3309)/simplerest?charset=utf8mb4&parseTime=True&loc=Local"
	//d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	d, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3309)/simplerest?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
