# playground

WIP refactor of Go Playground

## Building

```
# build the image
docker build -t playground .
```

## Running

```
docker run --rm -d -p 8080:8080 playground
# run go some code
cat /path/to/code.go | go run client.go | curl --data @- localhost:8080/compile
```

* If you're using `docker-machine`, replace `localhost` with the IP of your Docker server (typically `192.168.99.100`)

# Contributing

To submit changes to this repository, see http://golang.org/doc/contribute.html.