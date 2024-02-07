
gen:
	oapi-codegen --generate types,server,spec -package api api/swagger.yaml > cmd/api/api.gen.go
compose:
	docker compose up -d
upmigrate:
	migrate -path internal/db/migrations -database 'postgres://vallkof:password@localhost:15432/enrichdb?sslmode=disable' up
downmigrate:
	migrate -path internal/db/migrations -database 'postgres://vallkof:password@localhost:15432/enrichdb?sslmode=disable' down
run:
	go run cmd/enrich-service/main.go
stop:
	docker stop enrich-service
clean:
	docker rm enrich-service

.PHONY: gen compose upmigrate downmigrate run stop clean