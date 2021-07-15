package routers

import (
	"zzy2005137/todo/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() {

	r := gin.Default()
	r.Static("/static", "static")
	r.LoadHTMLGlob("template/*")

	r.GET("/", controllers.ShowIndex)

	v1Group := r.Group("v1")
	{
		//添加
		v1Group.POST("todo", controllers.CreateTodo)

		//查看
		v1Group.GET("/todo", controllers.RetrieveTodo)

		//修改
		v1Group.PUT("/todo/:id", controllers.UpdateTodo)

		//删除
		v1Group.DELETE("/todo/:id", controllers.DeleteTodo)
	}

	r.Run(":9090")

}
