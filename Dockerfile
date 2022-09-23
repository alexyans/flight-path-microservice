FROM golang:alpine

WORKDIR /go/src/github.com/volume/flight-path-microservice
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -v -o /go/bin/flight-path-microservice

CMD ["/go/bin/flight-path-microservice"]

