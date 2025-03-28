.PHONY: build clean

build:
    go build -o ../../bin/myapp .

clean:
    rm -f ../../bin/myapp
