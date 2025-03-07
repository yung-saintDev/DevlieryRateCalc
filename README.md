# 🚚 Delivery Price Calculator (Golang + Echo)

A simple **Golang-based REST API** that calculates the delivery price based on:
- 📦 **Restaurant total price**
- 📍 **Distance-based fee** (RUB per km)
- ⚖ **Weight-based fee** (RUB per kg)
- ⏩ **Priority surcharge** (for fast delivery)

Built using **Echo**, a fast and minimalistic web framework for Golang.

---

## 📌 Features
✅ Lightweight and fast REST API  
✅ Built using the **Echo framework**  
✅ Supports **JSON input/output**  
✅ Validates input data  
✅ Configurable delivery fee structure  

---

## 🚀 Installation & Usage

### Prerequisites
- [Golang](https://go.dev/) installed (Go 1.18 or later)
- [Echo](https://echo.labstack.com/) package

### Clone the Repository
```sh
git clone https://github.com/yourusername/delivery-price-calculator.git
cd delivery-price-calculator
```

### Install Dependencies
```sh
go mod tidy
```

### Run the Server
```sh
go run main.go
```
Server will start on **http://localhost:8080** 🚀

---

## 📡 API Endpoints

### 1️⃣ Calculate Delivery Price
**Endpoint:** `POST /calculate`  
**Description:** Returns the delivery price based on the provided parameters.  

#### Request Body (JSON)
```json
{
  "total_rest": 500,
  "distance_km": 10,
  "weight_kg": 2,
  "priority": "fast"
}
```

#### Response (JSON)
```json
{
  "price": 850
}
```

---

## 🛠 Configuration
You can modify the delivery pricing logic inside the `calculatePrice` function in **main.go**:
```go
const (
    distanceFee  = 20.0 // RUB per km
    weightFee    = 10.0 // RUB per kg
    fastMultiplier = 1.5
)
```

---

## 🧪 Testing with Postman
You can test the API using **Postman**:
1. Open Postman and create a new request.
2. Set the request type to **POST**.
3. Enter the URL: `http://localhost:8080/calculate`.
4. Navigate to the **Body** tab and select **raw**.
5. Choose **JSON** as the format.
6. Enter the following JSON payload:
   ```json
   {
     "total_rest": 500,
     "distance_km": 10,
     "weight_kg": 2,
     "priority": "fast"
   }
   ```
7. Click **Send** and verify the response.

Postman will return the calculated price as a JSON response. 🚀

---
