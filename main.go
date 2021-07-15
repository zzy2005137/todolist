package main

import (
	"zzy2005137/todo/dao"
	"zzy2005137/todo/models"
	"zzy2005137/todo/routers"

	_ "gorm.io/driver/mysql"
)

func main() {

	//创建数据库

	//连接数据库

	dao.Connect()

	//模型绑定
	dao.DB.AutoMigrate(&models.Todo{})

	//启动router
	routers.SetupRouter()

}
