FROM golang:1.18-rc-alpine
MAINTAINER youmingsama
WORKDIR /ToDoList
COPY go.mod ./
COPY go.sum ./
RUN  go mod download
COPY . .
RUN go  build  -o  todolist
EXPOSE 8086
CMD  ["./todolist"]
