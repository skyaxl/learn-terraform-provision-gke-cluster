FROM golang:1.15.14-alpine AS hello-1

WORKDIR /src

COPY . .
RUN go build ./...
EXPOSE 8080
CMD ["./hello"]
