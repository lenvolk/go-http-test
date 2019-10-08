# go-http-test
docker build --tag go-http-test:0.8 .
docker run --rm -d -p 80:8090/tcp go-http-test:0.8
http://localhost:8090/time
