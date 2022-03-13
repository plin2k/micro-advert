FROM golang:latest AS builder

RUN mkdir /app
ADD . /app
WORKDIR /app

RUN CGO_ENABLED=1 GOOS=linux go build -o main main.go

FROM alpine:latest AS production

WORKDIR /app

COPY --from=builder /app .

EXPOSE 8080

CMD ["./main"]
