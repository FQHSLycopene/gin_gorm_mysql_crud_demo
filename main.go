package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_gin_gorm_mysql_crud_demo/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strconv"
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
		//改
		listGroup.PUT("/:id", func(c *gin.Context) {
			var list model.List
			id := c.Param("id")
			DB.Where("id = ?", id).First(&list)
			if list.ID == 0 {
				c.JSON(200, gin.H{
					"code": 400,
					"msg":  "用户id没有找到",
				})
				return
			}
			err := c.ShouldBindJSON(&list)
			if err != nil {
				c.JSON(200, gin.H{
					"code": 400,
					"msg":  "修改失败",
				})
				return
			}
			DB.Save(&list)
			c.JSON(200, gin.H{
				"code": 200,
				"msg":  "修改成功",
			})
		})
		//查
		//条件查询
		listGroup.GET("/:name", func(c *gin.Context) {
			var lists []model.List
			name := c.Param("name")
			var total int64
			DB.Where("name = ?", name).Find(&lists).Count(&total)
			if len(lists) == 0 {
				c.JSON(200, gin.H{
					"msg":  "没有查询到数据",
					"data": nil,
					"code": 400,
				})
			} else {
				c.JSON(200, gin.H{
					"msg": "查询成功",
					"data": gin.H{
						"total": total,
						"lists": lists,
					},
					"code": 200,
				})
			}
		})
		listGroup.GET("/", func(c *gin.Context) {
			var dataList []model.List
			//1,查询全部数据，查询分页数据
			pageSize, _ := strconv.Atoi(c.Query("pageSize"))
			pageNum, _ := strconv.Atoi(c.Query("pageNum"))
			//判断是否要分页
			if pageSize == 0 {
				pageSize = -1
			}
			if pageNum == 0 {
				pageNum = -1
			}
			offsetVal := (pageNum - 1) * pageSize
			if pageSize == -1 {
				offsetVal = -1
			}
			var total int64
			DB.Model(dataList).Count(&total).Limit(pageSize).Offset(offsetVal).Find(&dataList)
			if len(dataList) == 0 {
				c.JSON(200, gin.H{
					"msg":  "没有查询到数据",
					"data": nil,
					"code": 400,
				})
			} else {
				c.JSON(200, gin.H{
					"msg": "查询成功",
					"data": gin.H{
						"list":     dataList,
						"total":    total,
						"pageNum":  pageNum,
						"pageSize": pageSize,
					},
					"code": 200,
				})
			}

		})

	}
	r.Run()

}
