package store

import (
	"sync"

	"fntv-proxy/model"
)

var (
	proxyInfo     model.ProxyInfo
	proxyInfoLock sync.RWMutex
)

// SetProxyInfo 更新代理信息
func SetProxyInfo(url, cookie string) {
	proxyInfoLock.Lock()
	defer proxyInfoLock.Unlock()
	proxyInfo.URL = url
	proxyInfo.Cookie = cookie
}

// GetProxyInfo 获取代理信息
func GetProxyInfo() model.ProxyInfo {
	proxyInfoLock.RLock()
	defer proxyInfoLock.RUnlock()
	return proxyInfo
}
