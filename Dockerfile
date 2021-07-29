FROM golang:latest

RUN mkdir /src
WORKDIR /src
Add . /src/

RUN go mod init src
RUN go get -u github.com/gin-gonic/gin