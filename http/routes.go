package http

import (
	_"fmt"
	"strconv"
	"github.com/banwire/api-exam/handlers"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"database/sql"
)



// Router define all routes http
func Router(router *gin.Engine) {

	db, err := sql.Open("mysql", "root:0987654321@tcp(127.0.0.1:3306)/golandapiexam")
			if err != nil {
				panic(err)
			}

	v1 := router.Group("/v1")
	{
		v1.POST("/example", handlers.ExampleHandler)
		//Save comercio
		v1.POST("/save-comercio", func (c *gin.Context) {

			merchant_name := c.PostForm("merchant_name")
			commission := c.PostForm("commission")
			created_at := time.Now()
			updated_at := time.Now()

			i,_ := strconv.Atoi(commission)

			if (i >= 1 && i <= 100) {

				var query string  = "INSERT INTO merchants (name, commission, created_at, updated_at) VALUES('"+merchant_name+"','"+commission+"','"+created_at.String()+"','"+updated_at.String()+"')"
				
				insert, err := db.Query(query)
	
				if err != nil {
					panic(err)
				}
				defer insert.Close()
				c.JSON(http.StatusOK, gin.H{
					"save": "OK",
				})
				
			}else{
				c.JSON(http.StatusOK, gin.H{
					"save": "NO",
					"messageError": "Comision no valida",
				})
			}
		})
		//Editar comercio
		v1.POST("/edit-comercio", func (c *gin.Context) {

			merchant_id := c.PostForm("merchant_id")
			merchant_name := c.PostForm("merchant_name")
			commission := c.PostForm("commission")
			updated_at := time.Now()

			i,_ := strconv.Atoi(commission)

			if (i >= 1 && i <= 100) {

				var query string  = "UPDATE merchants SET name='"+merchant_name+"', commission='"+commission+"', updated_at='"+updated_at.String()+"' WHERE id="+merchant_id
				
				insert, err := db.Query(query)
	
				if err != nil {
					panic(err)
				}
				defer insert.Close()
				c.JSON(http.StatusOK, gin.H{
					"save": "OK",
				})
				
			}else{
				c.JSON(http.StatusOK, gin.H{
					"save": "NO",
					"messageError": "Comision no valida",
				})
			}
		})
	}
}
