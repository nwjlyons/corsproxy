<h1 align="center">Corsproxy</h1>

Bypass CORS issues in local development

## Install

```
go install github.com/nwjlyons/corsproxy
``

## Example

Proxy http://localhost:8002 to http://example.com and add `Access-Control-Allow-Origin: *` to the response headers:

```
./corsproxy -port 8002 http://example.com
```

