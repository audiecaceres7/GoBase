include .env

build:
	@echo go build -o app main.go  && ./app

db_status:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(DB_CONN_STRING) goose -dir=$(GOOSE_MIGRATION_DIR) status

db_up:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(DB_CONN_STRING) goose -dir=$(GOOSE_MIGRATION_DIR) up

db_down:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(DB_CONN_STRING) goose -dir=$(GOOSE_MIGRATION_DIR) down

db_reset: 
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(DB_CONN_STRING) goose -dir=$(GOOSE_MIGRATION_DIR) reset
