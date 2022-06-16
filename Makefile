makeFileDir := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))

postgres: 
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=qwer -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root simplebank

dropdb:
	docker exec -it postgres12 dropdb simplebank

migrateup: 
	migrate -path db/migration -database "postgresql://root:qwer@172.17.14.20:5432/simplebank?sslmode=disable" -verbose up

migratedown: 
	migrate -path db/migration -database "postgresql://root:qwer@172.17.14.20:5432/simplebank?sslmode=disable" -verbose down

sqlc-generate:
	docker run --rm -v $(makeFileDir):/src -w /src kjconroy/sqlc generate

sqlc-init:
	docker run --rm -v $(makeFileDir):/src -w /src kjconroy/sqlc init

.PHONY: postgres createdb dropdb migrateup migratedown sqlc sqlc-init sqlc-generate