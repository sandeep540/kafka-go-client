# Start by building the application.
FROM golang:1.19 as build

WORKDIR /go/src/app
COPY . .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -o /go/bin/app .

# Now copy it into our base image.
FROM gcr.io/distroless/static-debian11
COPY --from=build /go/bin/app /
ENTRYPOINT [ "/app" ]


# docker build -t kafka-client-go .
# docker run -it kafka-client-go 