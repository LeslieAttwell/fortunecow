# Build Stage
FROM golang:alpine AS build-env
COPY app/main.go /src/main.go
RUN cd /src && \
    CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o fortunecow . 

# Create Image
FROM ubuntu:18.10
RUN apt-get update && apt-get install -y \ 
    fortune fortunes fortunes-min fortunes-bofh-excuses fortunes-ubuntu-server cowsay
WORKDIR /app
COPY --from=build-env /src/fortunecow /app/fortunecow
EXPOSE 8081
ENTRYPOINT ./fortunecow
