package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func initDB() {
	dsn := "Lycopene:MiMaJiuShi123321!@tcp(127.0.0.1:3306)/vue_go_crud_demo?charset=utf8mb4&parseTime=True&loc=Local"
	DB, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})
}

func main() {
	r := gin.Default()

	r.Run()

}
