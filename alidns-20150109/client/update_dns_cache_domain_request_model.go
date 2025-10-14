// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDnsCacheDomainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCacheTtlMax(v int32) *UpdateDnsCacheDomainRequest
	GetCacheTtlMax() *int32
	SetCacheTtlMin(v int32) *UpdateDnsCacheDomainRequest
	GetCacheTtlMin() *int32
	SetDomainName(v string) *UpdateDnsCacheDomainRequest
	GetDomainName() *string
	SetInstanceId(v string) *UpdateDnsCacheDomainRequest
	GetInstanceId() *string
	SetLang(v string) *UpdateDnsCacheDomainRequest
	GetLang() *string
	SetSourceDnsServer(v []*UpdateDnsCacheDomainRequestSourceDnsServer) *UpdateDnsCacheDomainRequest
	GetSourceDnsServer() []*UpdateDnsCacheDomainRequestSourceDnsServer
	SetSourceEdns(v string) *UpdateDnsCacheDomainRequest
	GetSourceEdns() *string
	SetSourceProtocol(v string) *UpdateDnsCacheDomainRequest
	GetSourceProtocol() *string
}

type UpdateDnsCacheDomainRequest struct {
	// The maximum TTL period of the cached data retrieved from the origin DNS server. Unit: seconds. Valid values: 30 to 86400.
	//
	// example:
	//
	// 86400
	CacheTtlMax *int32 `json:"CacheTtlMax,omitempty" xml:"CacheTtlMax,omitempty"`
	// The minimum time-to-live (TTL) period of the cached data retrieved from the origin Domain Name System (DNS) server. Unit: seconds. Valid values: 30 to 86400.
	//
	// example:
	//
	// 30
	CacheTtlMin *int32 `json:"CacheTtlMin,omitempty" xml:"CacheTtlMin,omitempty"`
	// The domain name. You can call the [DescribeDomains](https://www.alibabacloud.com/help/zh/dns/api-alidns-2015-01-09-describedomains?spm=a2c63.p38356.help-menu-search-29697.d_0) operation to obtian the domain name.
	//
	// This parameter is required.
	//
	// example:
	//
	// dns.example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The instance ID of the cache-accelerated domain name. You can call the [ListCloudGtmInstances](https://www.alibabacloud.com/help/zh/dns/api-alidns-2015-01-09-listcloudgtminstances?spm=a2c63.p38356.help-menu-search-29697.d_0) operation to obtain the ID.
	//
	// example:
	//
	// dns-sg-l9u2ux1fw01
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The language of the content within the request and response. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English Default: **zh**
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The origin DNS servers. A maximum of 10 origin DNS servers are supported.
	SourceDnsServer []*UpdateDnsCacheDomainRequestSourceDnsServer `json:"SourceDnsServer,omitempty" xml:"SourceDnsServer,omitempty" type:"Repeated"`
	// Specifies whether the origin DNS server supports Extension Mechanisms for DNS (EDNS). Valid values: NOT_SUPPORT and SUPPORT.
	//
	// example:
	//
	// SUPPORT
	SourceEdns *string `json:"SourceEdns,omitempty" xml:"SourceEdns,omitempty"`
	// The origin protocol policy. Valid values: TCP and UDP. Default value: UDP.
	//
	// example:
	//
	// UDP
	SourceProtocol *string `json:"SourceProtocol,omitempty" xml:"SourceProtocol,omitempty"`
}

func (s UpdateDnsCacheDomainRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDnsCacheDomainRequest) GoString() string {
	return s.String()
}

func (s *UpdateDnsCacheDomainRequest) GetCacheTtlMax() *int32 {
	return s.CacheTtlMax
}

func (s *UpdateDnsCacheDomainRequest) GetCacheTtlMin() *int32 {
	return s.CacheTtlMin
}

func (s *UpdateDnsCacheDomainRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *UpdateDnsCacheDomainRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateDnsCacheDomainRequest) GetLang() *string {
	return s.Lang
}

func (s *UpdateDnsCacheDomainRequest) GetSourceDnsServer() []*UpdateDnsCacheDomainRequestSourceDnsServer {
	return s.SourceDnsServer
}

func (s *UpdateDnsCacheDomainRequest) GetSourceEdns() *string {
	return s.SourceEdns
}

func (s *UpdateDnsCacheDomainRequest) GetSourceProtocol() *string {
	return s.SourceProtocol
}

func (s *UpdateDnsCacheDomainRequest) SetCacheTtlMax(v int32) *UpdateDnsCacheDomainRequest {
	s.CacheTtlMax = &v
	return s
}

func (s *UpdateDnsCacheDomainRequest) SetCacheTtlMin(v int32) *UpdateDnsCacheDomainRequest {
	s.CacheTtlMin = &v
	return s
}

func (s *UpdateDnsCacheDomainRequest) SetDomainName(v string) *UpdateDnsCacheDomainRequest {
	s.DomainName = &v
	return s
}

func (s *UpdateDnsCacheDomainRequest) SetInstanceId(v string) *UpdateDnsCacheDomainRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateDnsCacheDomainRequest) SetLang(v string) *UpdateDnsCacheDomainRequest {
	s.Lang = &v
	return s
}

func (s *UpdateDnsCacheDomainRequest) SetSourceDnsServer(v []*UpdateDnsCacheDomainRequestSourceDnsServer) *UpdateDnsCacheDomainRequest {
	s.SourceDnsServer = v
	return s
}

func (s *UpdateDnsCacheDomainRequest) SetSourceEdns(v string) *UpdateDnsCacheDomainRequest {
	s.SourceEdns = &v
	return s
}

func (s *UpdateDnsCacheDomainRequest) SetSourceProtocol(v string) *UpdateDnsCacheDomainRequest {
	s.SourceProtocol = &v
	return s
}

func (s *UpdateDnsCacheDomainRequest) Validate() error {
	if s.SourceDnsServer != nil {
		for _, item := range s.SourceDnsServer {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateDnsCacheDomainRequestSourceDnsServer struct {
	// The domain name or IP address of the origin DNS server.
	//
	// example:
	//
	// 223.5.5.5
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// The port of the origin DNS server.
	//
	// example:
	//
	// 53
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s UpdateDnsCacheDomainRequestSourceDnsServer) String() string {
	return dara.Prettify(s)
}

func (s UpdateDnsCacheDomainRequestSourceDnsServer) GoString() string {
	return s.String()
}

func (s *UpdateDnsCacheDomainRequestSourceDnsServer) GetHost() *string {
	return s.Host
}

func (s *UpdateDnsCacheDomainRequestSourceDnsServer) GetPort() *string {
	return s.Port
}

func (s *UpdateDnsCacheDomainRequestSourceDnsServer) SetHost(v string) *UpdateDnsCacheDomainRequestSourceDnsServer {
	s.Host = &v
	return s
}

func (s *UpdateDnsCacheDomainRequestSourceDnsServer) SetPort(v string) *UpdateDnsCacheDomainRequestSourceDnsServer {
	s.Port = &v
	return s
}

func (s *UpdateDnsCacheDomainRequestSourceDnsServer) Validate() error {
	return dara.Validate(s)
}
