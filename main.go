// main.go
package main

import (
	"github.com/manish2317/fetch-rewards-receipt-processor/routes"
)

func main() {
	router := routes.SetupRouter()
	router.Run(":8080")
}
