// controllers/controllers.go
package controllers

import (
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/manish2317/fetch-rewards-receipt-processor/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var receiptPoints = make(map[string]int)

func ProcessReceipt(c *gin.Context) {
	var receipt models.Receipt
	if err := c.BindJSON(&receipt); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	id := uuid.New().String()
	points := calculatePoints(receipt)
	receiptPoints[id] = points

	c.JSON(200, gin.H{"id": id})
}

func GetPoints(c *gin.Context) {
	id := c.Param("id")
	points, exists := receiptPoints[id]
	if !exists {
		c.JSON(404, gin.H{"error": "Receipt not found"})
		return
	}

	c.JSON(200, gin.H{"points": points})
}

func calculatePoints(receipt models.Receipt) int {
	points := 0

	// Rule 1: One point per alphanumeric character in retailer name
	regex := regexp.MustCompile(`[A-Za-z0-9]`)
	points += len(regex.FindAllString(receipt.Retailer, -1))

	totalFloat, err := strconv.ParseFloat(receipt.Total, 64)
	if err != nil {
		totalFloat = 0
	}

	// 50 points if total is a round dollar amount
	if totalFloat == math.Trunc(totalFloat) {
		points += 50
	}

	// 25 points if total is multiple of 0.25
	if math.Mod(totalFloat, 0.25) == 0 {
		points += 25
	}

	// 5 points for every two items
	points += (len(receipt.Items) / 2) * 5

	// Item description length rule
	for _, item := range receipt.Items {
		desc := strings.TrimSpace(item.ShortDescription)
		if len(desc)%3 == 0 {
			priceFloat, err := strconv.ParseFloat(item.Price, 64)
			if err == nil {
				points += int(math.Ceil(priceFloat * 0.2))
			}
		}
	} // <-- This closing brace was missing or misplaced earlier

	// 6 points if purchase day is odd
	purchaseDate, err := time.Parse("2006-01-02", receipt.PurchaseDate)
	if err == nil && purchaseDate.Day()%2 == 1 {
		points += 6
	}

	// 10 points if purchase time between 2:00 PM - 4:00 PM
	purchaseTime, err := time.Parse("15:04", receipt.PurchaseTime)
	if err == nil {
		startTime, _ := time.Parse("15:04", "14:00")
		endTime, _ := time.Parse("15:04", "16:00")
		if purchaseTime.After(startTime) && purchaseTime.Before(endTime) {
			points += 10
		}
	}

	return points
}
