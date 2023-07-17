package main

import (
	"github.com/gin-gonic/gin"
	guitarcontroller "github.com/rafifsanaputra/go-restapi-prakerja-rwp/controllers/guitarController"
	"github.com/rafifsanaputra/go-restapi-prakerja-rwp/models"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()

	r.GET("/api/guitars", guitarcontroller.Index)
	r.GET("/api/guitar/:id", guitarcontroller.Show)
	r.POST("/api/guitar", guitarcontroller.Create)
	r.PUT("/api/guitar/:id", guitarcontroller.Update)
	r.DELETE("/api/guitar/", guitarcontroller.Delete)

	r.Run()
}
