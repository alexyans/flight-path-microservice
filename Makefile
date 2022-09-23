.PHONY: build up

build:
	docker build --tag flight-path-microservice .

up:
	docker run -it flight-path-microservice

