FROM golang:1.23-alpine

WORKDIR /app

COPY . .
COPY .env-exemple .env

EXPOSE 8000

RUN go build -o main cmd/main.go

CMD ["./main"]