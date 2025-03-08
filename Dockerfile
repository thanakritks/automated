# Use the official Go image as the base
FROM golang:1.20 AS builder

WORKDIR /app
COPY . .

# Build the Go app
RUN go mod init myapp && go mod tidy && go build -o myapp

# Use a lightweight image for deployment
FROM alpine:3.14
WORKDIR /home/myapp
COPY --from=builder /app/myapp .

# Create a non-root user and use it
RUN addgroup -S myappgroup && adduser -S myappuser -G myappgroup
USER myappuser

CMD ["./myapp"]
