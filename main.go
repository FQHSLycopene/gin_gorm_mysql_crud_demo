package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_gin_gorm_mysql_crud_demo/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func initDB() {
	dsn := "Lycopene:MiMaJiuShi123321!@tcp(127.0.0.1:3306)/vue_go_crud_demo?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
}

func init() {
	initDB()
}

func main() {
	DB.AutoMigrate(&model.List{})
	r := gin.Default()
	listGroup := r.Group("/list")
	{
		//增
		listGroup.POST("/", func(c *gin.Context) {
			list := model.List{}
			err := c.ShouldBindJSON(&list)
			if err != nil {
				fmt.Println(err)
				c.JSON(200, gin.H{
					"msg":  "添加失败",
					"data": nil,
					"code": 400,
				})
			} else {
				DB.Create(&list)
				c.JSON(200, gin.H{
					"msg":  "添加成功",
					"data": list,
					"code": 200,
				})
			}
		})
		//删
		listGroup.DELETE("/:id", func(c *gin.Context) {
			var lists []model.List
			//接受id
			id := c.Param("id")
			//	判断id是否存在
			DB.Where("id = ?", id).Find(&lists)
			//id存在则删除，不存在报错
			if len(lists) == 0 {
				c.JSON(200, gin.H{
					"msg":  "id没有找到，删除失败",
					"data": nil,
					"code": 400,
				})
			} else {
				DB.Delete(&lists)
				c.JSON(200, gin.H{
					"msg":  "删除成功",
					"data": lists,
					"code": 200,
				})
			}
		})
		listGroup.GET("/")
		listGroup.GET("/:id")
		listGroup.PUT("/")

	}
	r.Run()

}
