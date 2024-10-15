.PHONY: build
build:
	@echo "Building"
	go build -o bin/echo-server

.PHONY: package
package:
	@echo "Packaging to Docker"
	docker build --tag pushm0v/echo-server:latest .
	docker push pushm0v/echo-server:latest
