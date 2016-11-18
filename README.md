Proxy for OCSP stapling
=======================

OCSP stapling means that the SSL server (rather than client) has to make requests to CA servers
for revoked certificates lists, making the check faster and more reliable for clients.

If:
* you're not allowed to connect from your SSL servers to the CA server because of a firewall,
* and your SSL server allows you to force the URL of the OCSP server
* but not of a HTTP proxy

then this tool may help you.

Usage
-----

   HTTP_PROXY=http://proxy:8888 ./ocsp-proxy -ocsphost ocspserver.com -http :8080

Il will listen on port 8080 for HTTP request and will forward the request to the ocsphost,
using the generic http proxy supplied the Go stdlib way.

In your nginx OCSP stapling configuration, add the line:

   ssl_stapling_responder http://127.0.0.1:8080;

(assuming the ocsp-proxy is running on 127.0.0.1 port 8080)

To find out your ocsphost, as far as I know:

   openssl x509 -in certificate.crt -noout -text | grep OCSP

(use the domain without scheme)
