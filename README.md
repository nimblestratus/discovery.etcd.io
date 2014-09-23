# discovery.etcd.io

This code powers the public service at https://discovery.etcd.io. The API is documented in

https://github.com/coreos/etcd/tree/master/Documentation/cluster-discovery.md
https://github.com/coreos/etcd/tree/master/Documentation/discovery-protocol.md

## Development

discovery.etcd.io uses devweb for easy development. It is simple to get started:

```
./devweb
curl --verbose -X PUT localhost:8087/new
```


## Running

This version of discovery.etcd.io supports the running of a "private"
or local discovery service.  You will need to have an etcd cluster
running, and the discovery server needs to run on an host where etcd
runs.

The url can be specified either via an environment variable or a
commandline parameter.

`DISCOVERY_URL`     ==> the url returned by the server
`DISCOVERY_ADDR`  ==> the address to which to bind

-- or -- 

    Usage of ./discovery:
      -addr="": web service address
      -url="https://discovery.etcd.io": url prefix for discovery, defaults to https://discovery.etcd.io
