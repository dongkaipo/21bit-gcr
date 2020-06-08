.PHONY: build

build:
	go build -o ./build/21bit-gcr .

build-linux:
	docker run -it -v $(pwd):/tmp/21bit-gcr golang:1.12.9-stretch bash -c "cd /tmp/21bit-gcr && make build"

clean:
	rm -rvf ./build
