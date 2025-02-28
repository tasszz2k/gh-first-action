# Stage 1: Build the Go binary
FROM golang:1.20-alpine AS builder

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum if available to cache dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the binary, naming it "action"
RUN go build -o action .

# Stage 2: Create a lightweight runtime image
FROM alpine:latest

# Set the working directory
WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/action .

# Set the container entrypoint to the action binary
ENTRYPOINT ["/app/action"]