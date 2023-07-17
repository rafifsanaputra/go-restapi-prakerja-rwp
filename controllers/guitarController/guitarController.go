package guitarcontroller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rafifsanaputra/go-restapi-prakerja-rwp/models"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {
	var guitars []models.Guitar
	models.DB.Find(&guitars)

	c.JSON(http.StatusOK, gin.H{"guitars": guitars})
}

// GET /api/guitar/:id
func Show(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	var guitar models.Guitar

	if err := models.DB.First(&guitar, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Guitar tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"guitar": guitar})
}

// POST /api/guitar
func Create(c *gin.Context) {
	var guitar models.Guitar

	if err := c.ShouldBindJSON(&guitar); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Create(&guitar)
	c.JSON(http.StatusCreated, gin.H{"guitar": guitar})
}

// PUT /api/guitar/:id
func Update(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	var guitar models.Guitar

	if err := models.DB.First(&guitar, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.JSON(http.StatusNotFound, gin.H{"error": "Guitar not found"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	err := c.ShouldBindJSON(&guitar)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Save(&guitar)
	c.JSON(http.StatusOK, gin.H{"guitar": guitar})
}

// DELETE /api/guitar
func Delete(c *gin.Context) {
	var requestBody struct {
		ID int64 `json:"id"`
	}

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var guitar models.Guitar
	result := models.DB.First(&guitar, requestBody.ID)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Guitar not found"})
		return
	}

	models.DB.Delete(&guitar)

	c.JSON(http.StatusOK, gin.H{"message": "Guitar deleted"})
}
