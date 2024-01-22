# 👽Go API

This project is a simple book inventory management API implemented in Go using the Gin web framework. It provides basic functionalities such as retrieving a list of books, retrieving a specific book by ID, creating a new book entry, checking out a book (decreasing its quantity), and returning a book (increasing its quantity). The data is stored in memory as a static list of books.

## Project Structure

- **main.go:** Contains the main application logic, including the setup of the Gin router and the implementation of CRUD operations for the book inventory.

## API Endpoints

1. **GET /books:** Retrieve a list of all books.

2. **GET /books/:id:** Retrieve details of a specific book by ID.

3. **POST /books:** Create a new book entry.

4. **PATCH /checkout:** Check out a book by decreasing its quantity (Query parameter style).

5. **PATCH /return:** Return a book by increasing its quantity (Query parameter style).

## Book Structure

- Each book has an ID, Title, Author, and Quantity.
 
  ```go
  type Book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
  }
  ```

## How to Run

1. Ensure you have Go installed on your machine.
2. Run the following command to install the required dependencies:

```bash
go get -u github.com/gin-gonic/gin
```

Execute the main.go file to start the server:
```bash
go run main.go
```

• Access the API at http://localhost:8080.

## API Usage
Use tools like curl or Postman to interact with the API endpoints.
Example: To check out a book, send a PATCH request to http://localhost:8080/checkout?id=<book_id>.

## Note
This is a simple demonstration project, and it uses in-memory data storage. For a production scenario, consider using a database for persistent data storage.

## Author
&copy; 2024 - Boularbah Ismail; <a href="https://twitter.com/boularbahsmail">Twitter</a>.
