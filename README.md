# Serve - Quick Single EndPoint HTTP Server

Purpose
-------------
Allowing for putting up of a quick endpoint with a set body, at a set port on localhost.
This can be used to test if server is externally accessible.

E.g.
```shell script
serve
```

Locally:
```shell script
curl http://localhost #Returns OK
```

Remotely:
```shell script
curl http://DNS_ADDRESS #Returns OK if port is open
```

<a id="section"></a>

Options
-------------
body: String — default: "OK"

path: String — default: "/"

port: Uint16 — default: 80


<a id="running"></a>

Binaries
-------------
There are builds for the main binaries needed:
```
    build/
        linux/amd64/serve
        linux/arm64/serve
        macos/amd64/serve
        macos/arm64/serve
        windows/amd64/serve.exe
```
This prevents the need to compile to use this utility


<a id="running"></a>
