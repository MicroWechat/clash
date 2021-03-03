package outboundgroup

import (
	"time"

	"github.com/Dreamacro/clash/adapters/provider"
	C "github.com/Dreamacro/clash/constant"
)

const (
	defaultGetProxiesDuration = time.Second * 5
)

type ProxyGroup interface {
	C.ProxyAdapter
	GetProxyProviders() []provider.ProxyProvider
	Now() string
}

func getProvidersProxies(providers []provider.ProxyProvider, touch bool) []C.Proxy {
	proxies := []C.Proxy{}
	for _, provider := range providers {
		if touch {
			proxies = append(proxies, provider.ProxiesWithTouch()...)
		} else {
			proxies = append(proxies, provider.Proxies()...)
		}
	}
	return proxies
}
