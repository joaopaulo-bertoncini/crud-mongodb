curl -X POST "http://localhost:8080/people" \
     -H "Content-Type: application/json" \
     -d '{"name": "João Bertoncini", "email": "joao@example.com", "cpf": "12345678900"}'

curl -X GET "http://localhost:8080/people"

curl -X GET "http://localhost:8080/people/{ID_}"

curl -X PUT "http://localhost:8080/people/{ID_AQUI}" \
     -H "Content-Type: application/json" \
     -d '{"name": "João B.", "email": "joao.b@example.com", "cpf": "12345678900"}'

curl -X DELETE "http://localhost:8080/people/{ID_}"
