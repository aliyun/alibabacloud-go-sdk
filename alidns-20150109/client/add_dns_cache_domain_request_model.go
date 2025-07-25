// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDnsCacheDomainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCacheTtlMax(v int32) *AddDnsCacheDomainRequest
	GetCacheTtlMax() *int32
	SetCacheTtlMin(v int32) *AddDnsCacheDomainRequest
	GetCacheTtlMin() *int32
	SetDomainName(v string) *AddDnsCacheDomainRequest
	GetDomainName() *string
	SetInstanceId(v string) *AddDnsCacheDomainRequest
	GetInstanceId() *string
	SetLang(v string) *AddDnsCacheDomainRequest
	GetLang() *string
	SetRemark(v string) *AddDnsCacheDomainRequest
	GetRemark() *string
	SetSourceDnsServer(v []*AddDnsCacheDomainRequestSourceDnsServer) *AddDnsCacheDomainRequest
	GetSourceDnsServer() []*AddDnsCacheDomainRequestSourceDnsServer
	SetSourceEdns(v string) *AddDnsCacheDomainRequest
	GetSourceEdns() *string
	SetSourceProtocol(v string) *AddDnsCacheDomainRequest
	GetSourceProtocol() *string
}

type AddDnsCacheDomainRequest struct {
	// The maximum TTL period of the cached data retrieved from the origin DNS server. Unit: seconds. Valid values: 30 to 86400.
	//
	// This parameter is required.
	//
	// example:
	//
	// 86400
	CacheTtlMax *int32 `json:"CacheTtlMax,omitempty" xml:"CacheTtlMax,omitempty"`
	// The minimum time-to-live (TTL) period of the cached data retrieved from the origin Domain Name System (DNS) server. Unit: seconds. Valid values: 30 to 86400.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30
	CacheTtlMin *int32 `json:"CacheTtlMin,omitempty" xml:"CacheTtlMin,omitempty"`
	// The domain name. You can call the [DescribeDomains](https://www.alibabacloud.com/help/zh/dns/api-alidns-2015-01-09-describedomains?spm=a2c63.p38356.help-menu-search-29697.d_0) operation to obtain the domain name.
	//
	// This parameter is required.
	//
	// example:
	//
	// dns.example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The instance ID of the cache-accelerated domain name. You can call the [ListCloudGtmInstances](https://www.alibabacloud.com/help/zh/dns/api-alidns-2015-01-09-listcloudgtminstances?spm=a2c63.p38356.help-menu-search-29697.d_0) operation to obtain the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// dns-cn-j6666
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
	// The remarks.
	//
	// example:
	//
	// test
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The origin DNS servers. A maximum of 10 origin DNS servers are supported.
	//
	// This parameter is required.
	SourceDnsServer []*AddDnsCacheDomainRequestSourceDnsServer `json:"SourceDnsServer,omitempty" xml:"SourceDnsServer,omitempty" type:"Repeated"`
	// Specifies whether the origin DNS server supports Extension Mechanisms for DNS (EDNS). Valid values: NOT_SUPPORT and SUPPORT.
	//
	// This parameter is required.
	//
	// example:
	//
	// SUPPORT
	SourceEdns *string `json:"SourceEdns,omitempty" xml:"SourceEdns,omitempty"`
	// The origin protocol policy. Valid values: TCP and UDP. Default value: UDP.
	//
	// This parameter is required.
	//
	// example:
	//
	// UDP
	SourceProtocol *string `json:"SourceProtocol,omitempty" xml:"SourceProtocol,omitempty"`
}

func (s AddDnsCacheDomainRequest) String() string {
	return dara.Prettify(s)
}

func (s AddDnsCacheDomainRequest) GoString() string {
	return s.String()
}

func (s *AddDnsCacheDomainRequest) GetCacheTtlMax() *int32 {
	return s.CacheTtlMax
}

func (s *AddDnsCacheDomainRequest) GetCacheTtlMin() *int32 {
	return s.CacheTtlMin
}

func (s *AddDnsCacheDomainRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *AddDnsCacheDomainRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AddDnsCacheDomainRequest) GetLang() *string {
	return s.Lang
}

func (s *AddDnsCacheDomainRequest) GetRemark() *string {
	return s.Remark
}

func (s *AddDnsCacheDomainRequest) GetSourceDnsServer() []*AddDnsCacheDomainRequestSourceDnsServer {
	return s.SourceDnsServer
}

func (s *AddDnsCacheDomainRequest) GetSourceEdns() *string {
	return s.SourceEdns
}

func (s *AddDnsCacheDomainRequest) GetSourceProtocol() *string {
	return s.SourceProtocol
}

func (s *AddDnsCacheDomainRequest) SetCacheTtlMax(v int32) *AddDnsCacheDomainRequest {
	s.CacheTtlMax = &v
	return s
}

func (s *AddDnsCacheDomainRequest) SetCacheTtlMin(v int32) *AddDnsCacheDomainRequest {
	s.CacheTtlMin = &v
	return s
}

func (s *AddDnsCacheDomainRequest) SetDomainName(v string) *AddDnsCacheDomainRequest {
	s.DomainName = &v
	return s
}

func (s *AddDnsCacheDomainRequest) SetInstanceId(v string) *AddDnsCacheDomainRequest {
	s.InstanceId = &v
	return s
}

func (s *AddDnsCacheDomainRequest) SetLang(v string) *AddDnsCacheDomainRequest {
	s.Lang = &v
	return s
}

func (s *AddDnsCacheDomainRequest) SetRemark(v string) *AddDnsCacheDomainRequest {
	s.Remark = &v
	return s
}

func (s *AddDnsCacheDomainRequest) SetSourceDnsServer(v []*AddDnsCacheDomainRequestSourceDnsServer) *AddDnsCacheDomainRequest {
	s.SourceDnsServer = v
	return s
}

func (s *AddDnsCacheDomainRequest) SetSourceEdns(v string) *AddDnsCacheDomainRequest {
	s.SourceEdns = &v
	return s
}

func (s *AddDnsCacheDomainRequest) SetSourceProtocol(v string) *AddDnsCacheDomainRequest {
	s.SourceProtocol = &v
	return s
}

func (s *AddDnsCacheDomainRequest) Validate() error {
	return dara.Validate(s)
}

type AddDnsCacheDomainRequestSourceDnsServer struct {
	// The domain name or IP address of the origin DNS server.
	//
	// This parameter is required.
	//
	// example:
	//
	// 192.0.0.0
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// The port of the origin DNS server.
	//
	// This parameter is required.
	//
	// example:
	//
	// 53
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s AddDnsCacheDomainRequestSourceDnsServer) String() string {
	return dara.Prettify(s)
}

func (s AddDnsCacheDomainRequestSourceDnsServer) GoString() string {
	return s.String()
}

func (s *AddDnsCacheDomainRequestSourceDnsServer) GetHost() *string {
	return s.Host
}

func (s *AddDnsCacheDomainRequestSourceDnsServer) GetPort() *string {
	return s.Port
}

func (s *AddDnsCacheDomainRequestSourceDnsServer) SetHost(v string) *AddDnsCacheDomainRequestSourceDnsServer {
	s.Host = &v
	return s
}

func (s *AddDnsCacheDomainRequestSourceDnsServer) SetPort(v string) *AddDnsCacheDomainRequestSourceDnsServer {
	s.Port = &v
	return s
}

func (s *AddDnsCacheDomainRequestSourceDnsServer) Validate() error {
	return dara.Validate(s)
}
