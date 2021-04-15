# Specify the version of Go to use
FROM golang:1.16-alpine

# Copy all the files from the host into the container
WORKDIR /src
COPY . .

# Compile the action
RUN go build -o morty

ENTRYPOINT ["morty"]
