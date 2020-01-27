# go-small-fasthttp

Minimal HTTP server for stress tests (based fasthttp)

For build:
> go build

For start:

> ./go-small-fasthttp

# result 

```
wrk -t100 -c500 -d10s http://127.0.0.1:4000/123

Running 10s test @ http://127.0.0.1:4000/123
  100 threads and 500 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     3.86ms  826.70us  13.51ms   74.89%
    Req/Sec     1.20k   322.44     7.69k    85.04%
  1198943 requests in 10.10s, 179.51MB read
  Socket errors: connect 0, read 205, write 0, timeout 0
Requests/sec: 118706.98
Transfer/sec:     17.77MB
```