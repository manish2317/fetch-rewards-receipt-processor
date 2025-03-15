// routes/routes.go
package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/manish2317/fetch-rewards-receipt-processor/controllers"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Define your endpoints here
	router.POST("/receipts/process", controllers.ProcessReceipt)
	router.GET("/receipts/:id/points", controllers.GetPoints)

	return router
}
