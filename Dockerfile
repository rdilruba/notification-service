FROM golang:1.21.6 AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o notification-app

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=build /app/notification-app /app/notification-app

COPY .env /app/.env

EXPOSE 8080

CMD ["./notification-app"]
