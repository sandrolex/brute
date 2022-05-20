bin_dir:
	mkdir -p bin/

build: bin_dir
	go build -o bin/brute cmd/main.go

clean:
	rm -rf bin
