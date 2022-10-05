package main

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func main() {
	username := "admin"
	password := "admin"
	token := strconv.FormatInt(time.Now().Unix(), 10)

	r := gin.Default()

	r.GET("/login", func(c *gin.Context) {
		if c.Query("username") == username && c.Query("password") == password {
			c.JSON(200, gin.H{
				"token": token,
			})
		} else {
			c.JSON(401, gin.H{
				"error": "Invalid credentials",
			})
		}
	})

	name := "NX"

	r.GET("/getUser", func(c *gin.Context) {
		if c.Query("token") == token {
			c.JSON(200, gin.H{
				"name": name,
			})
		} else {
			c.JSON(401, gin.H{
				"error": "Invalid token",
			})
		}
	})

	r.GET("/logout", func(c *gin.Context) {
		if c.Query("token") == token {
			token = strconv.FormatInt(time.Now().Unix(), 10)
			c.JSON(200, gin.H{
				"message": "Logout successful",
			})
		} else {
			c.JSON(401, gin.H{
				"error": "Invalid token",
			})
		}
	})
	r.Run()
}
