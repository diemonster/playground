# Go Playground

Hackweek refactor of Go Playground. 

## Goals

* Allow Third-Party Dependencies in Go Playground
* Disable network/filesystem controls (aka NaCl)
* Adapt deployment for Layer0, refactor data persistence to mongodb
* Add Auth0 Frontend for authentication (not done)
* Use internal domain name (not done)

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
