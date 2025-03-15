# Fetch Rewards Receipt Processor

This repository contains a web service designed to process receipts and calculate points based on specific rules. It is implemented using Go and the Gin framework.

## Project Structure

```
fetch-rewards-receipt-processor
├── controllers
│   └── controllers.go
├── models
│   └── receipt.go
├── routes
│   └── routes.go
├── go.mod
├── go.sum
└── main.go
```

## Requirements

- Go (1.20+ recommended)
- Gin framework
- Git

## Installation

Clone the repository:

```bash
git clone <repository-url>
cd fetch-rewards-receipt-processor
```

Install the necessary Go packages:

```bash
go mod tidy
```

## Running the Application

Start the server using the following command:

```bash
go run main.go
```

Your application will be running at `http://localhost:8080`.

## API Endpoints

### Process Receipt

- **Endpoint**: `/receipts/process`
- **Method**: `POST`
- **Payload**: Receipt JSON
- **Response**: Receipt ID

Example Request:

```bash
curl -X POST http://localhost:8080/receipts/process \
-H "Content-Type: application/json" \
-d '{
  "retailer": "Walmart123",
  "purchaseDate": "2022-03-17",
  "purchaseTime": "15:00",
  "items": [
    {"shortDescription": "Milk", "price": "3.00"},
    {"shortDescription": "Bread", "price": "2.00"}
  ],
  "total": "5.00"
}'
```

Example Response:

```json
{
  "id": "some-uuid-generated-id"
}
```

### Get Points

- **Endpoint**: `/receipts/{id}/points`
- **Method**: `GET`
- **Response**: Points calculated for the given receipt ID

Example Request:

```bash
curl http://localhost:8080/receipts/some-uuid-generated-id/points
```

Example Response:

```json
{
  "points": 106
}
```

## Points Calculation Rules

Points are calculated based on the following rules:

1. One point for every alphanumeric character in the retailer name.
2. 50 points if the total is a round dollar amount with no cents.
3. 25 points if the total is a multiple of 0.25.
4. 5 points for every two items on the receipt.
5. If the trimmed length of the item description is a multiple of 3, multiply the price by 0.2 and round up to the nearest integer.
6. 6 points if the purchase day is odd.
7. 10 points if the purchase time is between 2:00 PM and 4:00 PM.

## Testing

You can use `curl` or Postman to test the API endpoints with the provided example payloads above.

