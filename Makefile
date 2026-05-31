# postgres:
# 	docker run --name db -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=secret -d postgres
docker:
	docker-compose up -d
createdb:
	docker exec -it db createdb --username=root --owner=postgres simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down

dropdb:
	docker exec -it db dropdb simple_bank --username=root

test:
	go test -v -cover ./...

sqlc:
	sqlc generate


.PHONY: docker createdb dropdb migrateup migratedown test sqlc