dockerBuild:
	docker-compose up --build

dockerUp:
	docker-compose up

dockerDown:
	docker-compose down

dockerLog:
	docker-compose logs

goBuildR:
	go build cmd/go-auth-api/main.go && ./main