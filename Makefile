PORT=8081

run:
	@echo "Building and running go"
	go build . && ./memosocket ${PORT}