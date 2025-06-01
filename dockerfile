FROM golang:1.21

# Install tzdata for timezones
RUN apt-get update && apt-get install -y tzdata

# Set timezone to IST
ENV TZ=Asia/Kolkata

WORKDIR /app

# Copy all files into container
COPY . .

# Build the Go app
RUN go build -o main main.go

# Expose port
EXPOSE 8080

# Start app
CMD ["./main"]