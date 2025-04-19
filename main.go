package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 用户结构体
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// 模拟数据库（内存存储）
var users = []User{}
var nextID = 1

func main() {
	r := gin.Default() // 创建Gin引擎对象

	// 添加提示：GET /users提示信息
	r.GET("/users", func(c *gin.Context) {
		c.JSON(200, users)
	})

	// 创建用户
	r.POST("/users", func(c *gin.Context) {
		var newUser User
		if err := c.ShouldBindJSON(&newUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		newUser.ID = nextID
		nextID++
		users = append(users, newUser)
		c.JSON(http.StatusOK, newUser)
	})

	// 获取单个用户
	r.GET("/users/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		for _, u := range users {
			if u.ID == id {
				c.JSON(http.StatusOK, u)
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
	})

	// 更新用户
	r.PUT("/users/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		for i, u := range users {
			if u.ID == id {
				var updated User
				if err := c.ShouldBindJSON(&updated); err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
					return
				}
				updated.ID = id
				users[i] = updated
				c.JSON(http.StatusOK, updated)
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
	})

	// 删除用户
	r.DELETE("/users/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		for i, u := range users {
			if u.ID == id {
				users = append(users[:i], users[i+1:]...)
				c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
	})

	r.Run(":8080") // 启动HTTP服务，监听端口8080
}
