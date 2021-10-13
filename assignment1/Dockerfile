FROM golang:1.17-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY /httpserver/ ./

RUN go build -o /http-server

EXPOSE 8080

ENTRYPOINT ["/http-server"]
