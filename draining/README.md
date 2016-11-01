# HTTP Graceful Shutdown

This is an example web server that shuts down gracefully. Upon receiving the SIGINT signal it stops accepting connections, and waits until all outstanding connections are complete before ending.

See [https://blog.tylerchr.com/go-18-changes](https://blog.tylerchr.com/golang-18-changes).