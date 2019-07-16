# serv
A simple static file server. Can serve up multiple static directories, as defined by a simple JSON argument.

## Motivations
This was written to avoid writing tons of one off servers in golang to host
local directories for testing. Instead of creating a one-off golang
script each time I need to serve something, I created this.

It's also possible to do a similar task like this:
```bash
$ python -m SimpleHTTPServer
```

But the golang variant seems to perform better for large files. It also
allows for easily mapping URLs to directories.

## Install
### From Binary
On the [GitHub releases page](https://github.com/markwatson/serv/releases), download the latest binary for your platform from the "Assets" dropdown. Rename the file to `serv` if you like, and add to a folder in your path.

### From Source
- Install golang: <https://golang.org>
- Setup a go environment: <https://golang.org/doc/code.html>
- Use go get to install
```bash
$ go get github.com/markwatson/serv
```

## Examples

### Get help
```bash
$ serv -help
Usage of serv:
-listen=":8080": Interface/port to listen on. eg. :8080 or 127.0.0.1:8080
-paths="{\"/\": \".\"}": Paths to serve. A json object with the keys as the url pattern, and the value as the root. Default serves current folder.
```

### Serve up current directory on port 8080
```bash
$ serv
```

### Serve up current directory on port 6666
```bash
$ serv -listen :6666
```

### Serve up current directory on port 8080, listening only on localhost
```bash
$ serv -listen localhost:8080
```

### Serve up multiple directories
```bash
serv -paths '{
"/components/": "./web/web-app-dashboard/src/main/webapp/components",
"/elements/": "./web/web-app-dashboard/src/main/webapp/elements",
"/gen/": "./web/web-app-dashboard/src/main/webapp/gen",
"/resources/": "./web/web-app-dashboard/src/main/webapp/resources",
"/": "./web/web-app-dashboard/src/main/webapp/WEB-INF/templates"
}'
```
