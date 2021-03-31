// +build go1.8

package websocket

import (
	tls "github.com/refraction-networking/utls"
	"net/http/httptrace"
)

func doHandshakeWithTrace(trace *httptrace.ClientTrace, tlsConn *tls.UConn, cfg *tls.Config) error {
	return doHandshake(tlsConn, cfg)
}
