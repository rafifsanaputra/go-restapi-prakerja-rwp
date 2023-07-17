# go-restapi-prakerja-rwp

This repository contains a Go project built with the Gin framework, using GORM as an ORM library and MySQL as the database. It provides a simple API for managing a collection of guitars. This project is developed as part of an assignment for the Prakerja Course.

## Features

The API provides the following endpoints:

- `GET /api/guitars`: Retrieves a list of all guitars.
- `GET /api/guitar/:id`: Retrieves details of a specific guitar by ID.
- `POST /api/guitar`: Creates a new guitar (accepts raw JSON data in the request body).
- `PUT /api/guitar/:id`: Updates an existing guitar by ID (accepts raw JSON data in the request body).
- `DELETE /api/guitar/:id`: Deletes a specific guitar by ID.

## Technologies Used

- Go programming language
- Gin web framework
- GORM ORM library
- MySQL database

## Usage

To use the project, follow the instructions below:

1. Install Go and MySQL on your machine.
2. Clone the repository:
  ```shell
  git clone https://github.com/rafifsanaputra/go-restapi-prakerja-rwp.git
3. Set up your MySQL database and update the connection details in the code.

4. Run the application:
  ```shell
  go run main.go
5. Access the API endpoints using an HTTP client or web browser.
For POST /api/guitar and PUT /api/guitar/:id endpoints, you can send raw JSON data in the request body to create or update a guitar. The JSON data should have the following structure:
For creating a new guitar (POST /api/guitar):
  ```json
  {
    "brand": "Fender",
    "model": "Stratocaster",
    "description": "A classic electric guitar known for its versatile sound."
  }
For updating an existing guitar (PUT /api/guitar/:id):
  ```json
  {
    "brand": "Gibson",
    "model": "Les Paul",
    "description": "A solid-body electric guitar with a warm tone."
  }
Please note that you can customize the values of the brand, model, and description fields according to your specific requirements.
Make sure to include the appropriate Content-Type header in your request (e.g., application/json) when sending the raw JSON data to the mentioned endpoints.