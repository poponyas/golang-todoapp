FROM golang:1.25-bookworm AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o /app/exe ./cmd/todoapp

FROM alpine:3.21
WORKDIR /app
COPY --from=builder /app/exe /app/exe
COPY --from=builder /app/pub /app/pub
ENV PROJECT_ROOT=/app
CMD ["/app/exe"]