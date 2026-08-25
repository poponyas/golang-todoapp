include .env
export

export PROJECT_ROOT=$(shell pwd)

env-up:
	@docker compose up -d todoapp-postgres

env-down:
	@docker compose down todoapp-postgres

env-cleanup:
	@read -p "Are you sure you want to remove the database volume? (y/N): " ans; \
	if [ "$$ans" = "y" ] || [ "$$ans" = "Y" ]; then \
		docker compose down todoapp-postgres port-forwarder && \
		sudo rm -rf out/pgdata && \
		echo "Cleanup completed."; \
	else \
		echo "Aborting cleanup."; \
	fi

migrate-create:
	@if [ -z "$(seq)" ]; then \
		echo "Error: Please provide a sequence number using 'make migrate-create seq=<number>'"; \
		exit 1; \
	fi; \
	docker compose run --rm todoapp-postgres-migrate \
		create \
		-ext sql \
		-dir migrations \
		-seq "$(seq)"

migrate-up:
	@make migrate-action action=up

migrate-down:
	@make migrate-action action=down

migrate-force:
	@make migrate-action action="force $(v)"

migrate-action:
	@if [ -z "$(action)" ]; then \
		echo "Error: Please provide an action using 'make migrate-action action=<up|down>'"; \
		exit 1; \
	fi; \
	docker compose run --rm todoapp-postgres-migrate \
		-path /migrations \
		-database postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@todoapp-postgres:5432/$(POSTGRES_DB)?sslmode=disable \
		"$(action)"

env-port-forward:
	@docker compose up -d port-forwarder
	
env-port-close:
	@docker compose down port-forwarder

todoapp-run:
	@export LOGGER_FOLDER=$(PROJECT_ROOT)/out/logs && \
	export POSTGRES_HOST=localhost && \
	go mod tidy && \
	go run cmd/todoapp/main.go