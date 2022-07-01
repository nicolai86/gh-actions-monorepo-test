FROM golang:1.18.3-alpine3.16 as builder

WORKDIR /app

COPY ./ ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /build/app main.go

FROM scratch

COPY --from=builder /build/app /bin/app

ENTRYPOINT ["/bin/app"]