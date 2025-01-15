FROM golang:1.23

#set working directory
WORKDIR /go/src/app

#copy the source code
COPY . .

#EXPOSE the port
EXPOSE 8000

#Build the Go app
RUN go build -o main cmd/main.go

#Run the executable
CMD ["./main"]