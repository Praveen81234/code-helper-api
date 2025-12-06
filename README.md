# Code Helper API (Go Fiber)

This project is a simple backend API built using Go Fiber for the Coding Junior assignment.  
It provides three main APIs:

## 1. Run API
- POST `/run`
- Accepts raw code
- Returns simulated output

## 2. Auto Fix API
- POST `/autofix`
- Adds missing semicolons
- Fix indentation
- Removes extra spaces
- Returns fixed code

## 3. Help API
- POST `/help`
- Returns predefined tips based on keywords like "syntax", "error", "loop", etc.

## How to run

1. Install Go
2. Open terminal
3. Run the following commands:

```
go mod tidy
go run main.go
```

Server will start on:
```
http://localhost:3000
```
## Example requests (using Postman)

### 1. Run Code API

- Method: `POST`
- URL: `http://localhost:3000/run`
- Body (JSON):

```json
{
  "code": "print('hello')"
}
```

---

### 2. Auto-fix API  

- Method: `POST`
- URL: `http://localhost:3000/auto-fix`
- Body (JSON):

```json
{
  "code": "print('hello'"
}
```

**Output Example**
```json
{
  "fixedCode": "print('hello')"
}
```

---

### 3. Help API  

- Method: `POST`
- URL: `http://localhost:3000/help`
- Body:

```json
{
  "query": "syntax"
}
```

**Output Example**
```json
{
  "tip": "Tip about missing brackets, commas, semicolons."
}
```

