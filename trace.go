// +build go1.8

package websocket

import (
	"net/http/httptrace"

	tls "github.com/refraction-networking/utls"
)

func doHandshakeWithTrace(trace *httptrace.ClientTrace, tlsConn *tls.UConn, cfg *tls.Config) error {
	return doHandshake(tlsConn, cfg)
}
