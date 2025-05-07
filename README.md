# Go-Movie-Crud
# 🎬 Go Movie API

A simple RESTful API built with **Go** and **Gorilla Mux** to manage a collection of movies. This API allows you to create, read, update, and delete movie records using in-memory storage.

---

## 🛠 Technologies Used

- Go (Golang)
- Gorilla Mux (HTTP router)
- net/http (Go standard library)

---

## 🚀 Getting Started

### ✅ Prerequisites

- Go 1.18+ installed
- Git installed

### 📦 Clone the Repository

```bash
git clone https://github.com/your-username/go-movie-api.git
cd go-movie-api
go run main.go
The server will start at:


http://localhost:8000

POST /movie
GET /movies
GET /movies/{id}

Example: /movies/1
PUT /movies/{id}

Example: /movies/1
DELETE /movies/{id}

Example: /movies/1