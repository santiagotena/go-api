## Run Application
```bash
go run main.go
```

## Available Methods
Display all Books - GET
```bash
curl localhost:8080/books
```

Display a book based on its id - GET
```bash
curl localhost:8080/books/3
```

Add a book from a .json file - POST
```bash
curl localhost:8080/books --include --header "Content-Type: application/json" -d @body.json --request "POST"
```

Checkout a book from the library, if stock allows it - PATCH
```bash
curl localhost:8080/checkout?id=2 --request "PATCH"
```

Return a book to the library, increase book quantity - PATCH
```bash
curl localhost:8080/return?id=2 --request "PATCH"
```