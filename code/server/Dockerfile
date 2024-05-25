FROM golang:alpine AS builder
# disable C link for smaller build
ENV CGO_ENABLED 0
ENV GOOS linux
WORKDIR /opt/app
COPY . .
# flags for smaller build
RUN go build -ldflags="-s -w"

FROM alpine
WORKDIR /opt/app
COPY --from=builder /opt/app/server .
CMD ["/bin/sh", "-c", "./server"]
