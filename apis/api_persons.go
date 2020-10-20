package apis

import (
	"net/http"

	"github.com/202lp2/go2/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

//CRUD for items table
func PersonsIndex(c *gin.Context) {
	var lis []models.Person

	db, _ := c.Get("db")
	conn := db.(gorm.DB)

	// Migrate the schema
	conn.AutoMigrate(&models.Person{})

	conn.Find(&lis)
	c.JSON(http.StatusOK, lis)
}

func PersonsCreate(c *gin.Context) {
	db, _ := c.Get("db")
	conn := db.(gorm.DB)
	d := models.Person{Name: c.PostForm("name"), Age: c.PostForm("age"), ApPaterno: c.PostForm("apPaterno"),
		ApMaterno: c.PostForm("apMaterno"), EstadoCivil: c.PostForm("estadoCivil")}
	conn.Create(&d)

	c.JSON(200, gin.H{ // serializador de gin
		"name":        d.Name,
		"age":         d.Age,
		"apPaterno":   d.ApPaterno,
		"apMaterno":   d.ApMaterno,
		"estadoCivil": d.EstadoCivil,
	})
}

func PersonsDelete(c *gin.Context) {
	db, _ := c.Get("db")
	conn := db.(gorm.DB)

	id := c.Param("id")
	var d models.Person
	if err := conn.Where("id = ?", id).First(&d).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message":     "Eliminado",
		"name":        d.Name,
		"age":         d.Age,
		"apPaterno":   d.ApPaterno,
		"apMaterno":   d.ApMaterno,
		"estadoCivil": d.EstadoCivil,
	})
	conn.Unscoped().Delete(&d)
}

func PersonsUpdate(c *gin.Context) {
	db, _ := c.Get("db")
	conn := db.(gorm.DB)
	id := c.Param("id")
	var d models.Person
	if err := conn.First(&d, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	d.Name = c.PostForm("name")
	d.Age = c.PostForm("age")
	d.ApPaterno = c.PostForm("apPaterno")
	d.ApMaterno = c.PostForm("apMaterno")
	d.EstadoCivil = c.PostForm("estadoCivil")
	//c.BindJSON(&d)
	conn.Save(&d)
	c.JSON(http.StatusOK, &d)
}
