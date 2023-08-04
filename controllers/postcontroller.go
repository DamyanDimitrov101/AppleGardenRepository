package controllers

import (
	"github.com/damyan.dimitrov/gogardenofapples/initializers"
	"github.com/damyan.dimitrov/gogardenofapples/models"
	"github.com/gin-gonic/gin"
)

func PostCreate(c *gin.Context) {
	//get data from request body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	//create a post
	post := models.Post{Title: body.Title, Body: body.Body}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{"post": post})
}

func PostIndex(c *gin.Context) {
	// GET the posts
	var posts []models.Post
	initializers.DB.Find(&posts)

	//Respond
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostShow(c *gin.Context) {
	//Get id of url
	id := c.Param("id")

	// GET the post
	var post models.Post
	initializers.DB.First(&post, id)

	//Respond
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostUpdate(c *gin.Context) {
	//Get id of url
	id := c.Param("id")

	//get data from request body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	// GET the post
	var post models.Post
	initializers.DB.First(&post, id)

	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	//Respond
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostDelete(c *gin.Context) {
	//Get id of url
	id := c.Param("id")

	//Delete the post
	initializers.DB.Delete(&models.Post{}, id)

	//Respond
	c.Status(200)
}
