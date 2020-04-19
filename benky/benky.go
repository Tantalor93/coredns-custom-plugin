package benky

import (
	"context"
	"github.com/coredns/coredns/plugin"
	clog "github.com/coredns/coredns/plugin/pkg/log"
	"github.com/miekg/dns"
)

var log = clog.NewWithPlugin("benky")

type Benky struct {
	Next plugin.Handler
}

// ServeDNS implements the plugin.Handler interface. This method gets called when example is used
// in a Server. Return type int represents DNS response code https://www.iana.org/assignments/dns-parameters/dns-parameters.xhtml
// for example you can use constants dns.RcodeNameError, but usually you just want to call next plugin
// r *dns.Msg is incoming DNS query , w returns response to the client using w.WriteMsg(&dns.Msg{...})
func (b Benky) ServeDNS(ctx context.Context, w dns.ResponseWriter, r *dns.Msg) (int, error) {
	log.Info("jedeem")
	return plugin.NextOrFailure(b.Name(), b.Next, ctx, w, r)
}

// Name implements the Handler interface.
func (b Benky) Name() string { return "benky" }
