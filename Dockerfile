FROM golang:1.17.8-alpine3.15 AS builder 

WORKDIR /app

COPY . . 

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" . 

FROM amd64/alpine:3.15

RUN apk update 
RUN apk add postgresql-client
RUN apk add net-tools
RUN apk add busybox-extras
RUN apk add curl

COPY --from=builder /app/podtato /usr/bin/ 

ENTRYPOINT ["podtato"]



