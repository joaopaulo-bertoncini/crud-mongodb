# Building a CRUD API with Go and MongoDB

Developing a RESTful API using Go and MongoDB is a great way to build scalable and efficient applications. In this tutorial, we will create a simple CRUD (Create, Read, Update, Delete) API using the Gin web framework and MongoDB as the database.

### Prerequisites
Before getting started, ensure you have the following installed:

- Go (1.18 or later)
- Docker

## Run the API with make

jp@jp-pc:~/crud-mongodb$ make

Usage:\n
  help          print this help message\\
  run/api       run the cmd/api application\\
  build/api     build the cmd/api application\\
  clean/apps    clear generated bin files\\
  docker/up     start the local stack in background\\
  docker/down   shutdown the running containers\\
  audit         tidy dependencies, format and vet all code\\
  tidy          tidy dependencies\\

## Testing the API with cURL

### Create a New Person

curl -X POST "http://localhost:8080/people" \
     -H "Content-Type: application/json" \
     -d '{"name": "João Bertoncini", "email": "joao@example.com", "cpf": "12345678900"}'

### List All People

curl -X GET "http://localhost:8080/people"

### Get a Specific Person by ID

curl -X GET "http://localhost:8080/people/{ID_}"

Replace {ID_} with the actual ID returned when creating a person.

### Update a Person

curl -X PUT "http://localhost:8080/people/{ID_AQUI}" \
     -H "Content-Type: application/json" \
     -d '{"name": "João B.", "email": "joao.b@example.com", "cpf": "12345678900"}'

### Delete a Person

curl -X DELETE "http://localhost:8080/people/{ID_}"

