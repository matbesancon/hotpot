
build: clean
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o main ./cmd/hotpot/main.go

container: build
	docker build . --tag=mbesancon/hotpot

run: clean build container
	docker run -p 8080:8080 mbesancon/hotpot

clean:
	rm ./main

publish: clean build container
	docker push mbesancon/hotpot
