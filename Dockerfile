FROM golang:1.8 as builder
WORKDIR /go/src/github.com/tomhibbers/lnlamas-api-go/
RUN go get -d -v golang.org/x/net/html  
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/github.com/tomhibbers/lnlamas-api-go/main .
CMD ["./main"] 