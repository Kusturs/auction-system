.PHONY: up down migrate postgres recreate-db build logs

DC=docker compose
DB_USER=auction_user
DB_PASS=auction_password
DB_NAME=auction_db
DB_HOST=localhost
DB_PORT=5432
MIGRATIONS_DIR=migrations
PROJECT_NAME=auction-system
NETWORK=$(PROJECT_NAME)_auction-network

up:
	$(DC) up

down:
	$(DC) down -v

build:
	$(DC) build

logs:
	$(DC) logs -f

postgres:
	$(DC) up -d postgres
	@until docker exec $$(docker ps -q -f name=postgres) pg_isready -U $(DB_USER) -d $(DB_NAME); do \
		echo "Waiting for postgres..."; \
		sleep 1; \
	done

recreate-db: postgres
	docker exec $$(docker ps -q -f name=postgres) dropdb -U $(DB_USER) --if-exists $(DB_NAME)
	docker exec $$(docker ps -q -f name=postgres) createdb -U $(DB_USER) $(DB_NAME)

migrate: recreate-db
	@for file in $(MIGRATIONS_DIR)/*.up.sql; do \
		echo "Applying $$file..."; \
		docker exec -i $$(docker ps -q -f name=postgres) psql -U $(DB_USER) -d $(DB_NAME) < $$file; \
	done

PROTO_DIR=api/proto
GO_OUT_DIR=pkg/api

.PHONY: gen-proto install-proto-deps

install-proto-deps:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest

gen-proto: install-proto-deps
	mkdir -p $(GO_OUT_DIR)
	protoc -I $(PROTO_DIR) \
		--go_out=$(GO_OUT_DIR) --go_opt=paths=source_relative \
		--go-grpc_out=$(GO_OUT_DIR) --go-grpc_opt=paths=source_relative \
		$(PROTO_DIR)/*.proto

restart-app:
	$(DC) restart app

app-logs:
	$(DC) logs -f app

db-logs:
	$(DC) logs -f postgres

start: build up migrate

.DEFAULT_GOAL := start