# dumbproxy

This is a dumb reverse proxy for the absolute simplest use cases imaginable.

Don't want to depend on this sketchy repo? Fork it or push the docker image to your own registry.

## Usage

Docker:

```
docker run --rm -p 8080:8080 \
  forestgagnon/dumbproxy:1 \
  --listen=:8080 --upstream=https://example.com
```

Go:

```
go get -u github.com/forestgagnon/dumbproxy
```

## Contributing

No features will ever be added.

If the docker image becomes so stale someday that it's missing essential CA certificates, feel free to file an issue. If I'm still alive, I'll fix it.
