package controllers

import (
	"net/http"
	"zzy2005137/todo/dao"
	"zzy2005137/todo/models"

	"github.com/gin-gonic/gin"
)

func ShowIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func CreateTodo(c *gin.Context) {
	//get data
	var todo models.Todo
	c.ShouldBind(&todo)

	//add into database
	err := dao.DB.Create(&todo).Error

	//return
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, todo)
	}

}

func RetrieveTodo(c *gin.Context) {
	var todos []models.Todo
	if err := dao.DB.Find(&todos).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, todos)
	}
}

func UpdateTodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "id invalid"})
		return
	}

	var todo models.Todo
	if err := dao.DB.Where("id=?", id).First(&todo).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	c.BindJSON(&todo) //修改

	if err := dao.DB.Save(&todo).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func DeleteTodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "id invalid"})
		return
	}

	if err := dao.DB.Where("id = ?", id).Delete(models.Todo{}).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{id: "deleted"})
	}
}
