FROM golang:alpine3.14
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY *.go ./
RUN go build -o /go-keycloack-auth
EXPOSE 8081
CMD [ "/go-keycloack-auth" "--host", "0.0.0.0"]