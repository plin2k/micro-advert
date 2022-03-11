FROM golang:latest AS builder

RUN mkdir /app
ADD . /app
WORKDIR /app

RUN CGO_ENABLED=0 GOOS=linux go build -o main main.go

FROM alpine:latest AS production

WORKDIR /app

COPY --from=builder /app/main main
COPY --from=builder /app/resources ./resources
COPY --from=builder /app/templates ./templates


EXPOSE 8080

CMD ["./main"]
