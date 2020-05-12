package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
	"time"
)

var (
	db  *gorm.DB
	err error
)

type User struct {
	gorm.Model
	Name  string `gorm:"size:60"`
	Email string `gorm:"unique"`
}

func dbConnect() {
	db, err = gorm.Open("mysql", "homestead:secret@tcp(127.0.0.1:33060)/go_web_blog?charset=utf8mb4&parseTime=true&loc=Local")
	if err != nil {
		panic(err)
	}
}

func main() {
	server := gin.Default()

	dbConnect()
	defer db.Close()

	db.AutoMigrate(&User{})

	db.Create(&User{
		Model: gorm.Model{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Name:  "高杉",
		Email: "931692760@qq.com",
	})

	//用户列表
	server.GET("/users", func(context *gin.Context) {
		var users []*User
		if err = db.Find(&users).Error; err != nil {
			panic(err)
		}
		context.JSON(http.StatusOK, gin.H{
			"data": users,
		})
	})

	if err := server.Run(":80"); err != nil {
		panic(err)
	}
}
