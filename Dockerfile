FROM golang:1.14-alpine as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY main.go ./
RUN CGO_ENABLED=0 go build -mod=readonly -v -ldflags="-s -w" -o dumbproxy .

FROM alpine:3.11
RUN apk add --update --no-cache ca-certificates
WORKDIR /app

COPY --from=builder /app/dumbproxy ./
ENTRYPOINT ["/app/dumbproxy"]
