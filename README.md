# rtorrent-size
Command line utility shows how much storage space (in bytes) all rTorrent instance 
torrents occupy when fully downloaded. It sums the sizes of all the 
torrents (still downloading and already downloaded) and prints it to the stdout.

Installation

```shell
go install github.com/smaugfm/rtorrent-size@latest
```

Usage:
```shell
rtorrent-size http:://localhost:8000/RPC2
```

With HTTP Basic Auth:
```shell
rtorrent-size --username someUser --password somePassword http:://localhost:8000/RPC2
```
