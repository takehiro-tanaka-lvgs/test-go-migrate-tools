FROM golang:1.22

WORKDIR /go/src/github.com/takehiro-tanaka-lvgs/sql-migrate-sample

RUN apt-get update && apt-get install git
RUN go install github.com/go-delve/delve/cmd/dlv@latest
RUN go install golang.org/x/tools/gopls@latest
RUN go install github.com/go-task/task/v3/cmd/task@latest

RUN go install github.com/rubenv/sql-migrate/...@latest

COPY ./go.mod ./
COPY ./go.sum ./

RUN go mod download
RUN go mod tidy
