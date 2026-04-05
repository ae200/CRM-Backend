# CRM-Backend
# CRM Backend API

## Project Overview

This project is the backend for a Customer Relationship Management (CRM) web application. It provides a server that handles customer data and supports full CRUD (Create, Read, Update, Delete) operations.

As users interact with the frontend interface, this backend processes requests and manages customer records accordingly.

---

## Features

The API supports the following functionality:

* Get a list of all customers
* Get details for a single customer
* Add a new customer
* Update an existing customer's information
* Delete a customer

---

## Development Strategy

### 1. Customer Representation

Define a `Customer` structure with typed fields (e.g., ID, name, email, phone number, etc.).

### 2. Data Storage

Store customer data in a structure that supports CRUD operations (e.g., slice, map, or in-memory database). Include some seed data for testing.

### 3. API Endpoints

Design RESTful endpoints and associate them with appropriate HTTP methods:

| Method | Endpoint        | Description                |
| ------ | --------------- | -------------------------- |
| GET    | /customers      | Retrieve all customers     |
| GET    | /customers/{id} | Retrieve a single customer |
| POST   | /customers      | Add a new customer         |
| PUT    | /customers/{id} | Update a customer          |
| DELETE | /customers/{id} | Delete a customer          |

### 4. Routing

Implement a router to direct incoming requests to the correct handlers.

### 5. Request Handlers

Create dedicated handlers for each endpoint:

* Parse incoming request data (e.g., JSON)
* Validate inputs
* Perform appropriate operations on the data store
* Return meaningful HTTP responses and status codes

### 6. Data Handling

Ensure handlers:

* Properly parse request bodies
* Convert data formats where necessary
* Return responses in a consistent format (e.g., JSON)

### 7. Testing During Development

Run the server locally (e.g., `localhost`) and test endpoints using tools like Postman to verify correct behavior.

---

## Unit Testing

Unit tests are provided to help validate your implementation.

### Running Tests

1. Download the `main_test.go` file
2. Place it in the same directory as your main server file
3. Run the following command:

```bash
go test
```

### Expected Output

* If all tests pass:

  ```
  PASS
  ```
* If tests fail, you'll see helpful error messages such as:

  ```
  The addCustomer() handler returned a 200 status code, where it should have returned 201
  ```

These messages indicate how your API should behave and help guide corrections.

> Note: Passing all tests does not guarantee full project completion, as tests may not cover all requirements.

---

## Summary

This project focuses on building a functional RESTful API in Go that manages customer data. By implementing structured data handling, proper routing, and well-defined endpoints, you will create a robust backend service for a CRM application.

