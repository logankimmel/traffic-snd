FROM golang:1.13 as golang
COPY main.go /go/src/main/
WORKDIR /go/src/main
RUN cd /go/src/main && go get ./...
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM scratch
COPY --from=0 /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=golang /go/src/main/main /
CMD ["/main"]
