# syntax=docker/dockerfile:1

FROM golang:1.23-alpine

COPY --from=ghcr.io/typst/typst:latest /bin/typst /usr/local/bin/typst

RUN apk add --no-cache libgcc libstdc++

WORKDIR $GOPATH/src/github.com/gresio/print-server

# Copy everything from the current directory to the PWD (Present Working Directory) inside the container
COPY . .

# Download all the dependencies
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...

# This container exposes port 8631 to the outside world
EXPOSE 8631

# Run the executable
CMD ["print-server"]