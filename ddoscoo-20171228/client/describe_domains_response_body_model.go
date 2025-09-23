// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomains(v []*DescribeDomainsResponseBodyDomains) *DescribeDomainsResponseBody
	GetDomains() []*DescribeDomainsResponseBodyDomains
	SetRequestId(v string) *DescribeDomainsResponseBody
	GetRequestId() *string
	SetTotal(v int64) *DescribeDomainsResponseBody
	GetTotal() *int64
}

type DescribeDomainsResponseBody struct {
	Domains []*DescribeDomainsResponseBodyDomains `json:"Domains,omitempty" xml:"Domains,omitempty" type:"Repeated"`
	// example:
	//
	// CF33B4C3-196E-4015-AADD-5CAD00057B80
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 10
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeDomainsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainsResponseBody) GetDomains() []*DescribeDomainsResponseBodyDomains {
	return s.Domains
}

func (s *DescribeDomainsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDomainsResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *DescribeDomainsResponseBody) SetDomains(v []*DescribeDomainsResponseBodyDomains) *DescribeDomainsResponseBody {
	s.Domains = v
	return s
}

func (s *DescribeDomainsResponseBody) SetRequestId(v string) *DescribeDomainsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainsResponseBody) SetTotal(v int64) *DescribeDomainsResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeDomainsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDomainsResponseBodyDomains struct {
	BlackList []*string `json:"BlackList,omitempty" xml:"BlackList,omitempty" type:"Repeated"`
	// example:
	//
	// true
	CcEnabled *bool `json:"CcEnabled,omitempty" xml:"CcEnabled,omitempty"`
	// example:
	//
	// true
	CcRuleEnabled *bool `json:"CcRuleEnabled,omitempty" xml:"CcRuleEnabled,omitempty"`
	// example:
	//
	// normal
	CcTemplate *string `json:"CcTemplate,omitempty" xml:"CcTemplate,omitempty"`
	// example:
	//
	// testCertName
	CertName   *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	CertRegion *string `json:"CertRegion,omitempty" xml:"CertRegion,omitempty"`
	// example:
	//
	// xxxxxxx.aliyunddos1006.com
	Cname *string `json:"Cname,omitempty" xml:"Cname,omitempty"`
	// example:
	//
	// www.aliyun.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// example:
	//
	// true
	Http2Enable   *bool                                              `json:"Http2Enable,omitempty" xml:"Http2Enable,omitempty"`
	ProxyTypeList []*DescribeDomainsResponseBodyDomainsProxyTypeList `json:"ProxyTypeList,omitempty" xml:"ProxyTypeList,omitempty" type:"Repeated"`
	RealServers   []*DescribeDomainsResponseBodyDomainsRealServers   `json:"RealServers,omitempty" xml:"RealServers,omitempty" type:"Repeated"`
	// example:
	//
	// xx
	SslCiphers *string `json:"SslCiphers,omitempty" xml:"SslCiphers,omitempty"`
	// example:
	//
	// xx
	SslProtocols *string   `json:"SslProtocols,omitempty" xml:"SslProtocols,omitempty"`
	WhiteList    []*string `json:"WhiteList,omitempty" xml:"WhiteList,omitempty" type:"Repeated"`
}

func (s DescribeDomainsResponseBodyDomains) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainsResponseBodyDomains) GoString() string {
	return s.String()
}

func (s *DescribeDomainsResponseBodyDomains) GetBlackList() []*string {
	return s.BlackList
}

func (s *DescribeDomainsResponseBodyDomains) GetCcEnabled() *bool {
	return s.CcEnabled
}

func (s *DescribeDomainsResponseBodyDomains) GetCcRuleEnabled() *bool {
	return s.CcRuleEnabled
}

func (s *DescribeDomainsResponseBodyDomains) GetCcTemplate() *string {
	return s.CcTemplate
}

func (s *DescribeDomainsResponseBodyDomains) GetCertName() *string {
	return s.CertName
}

func (s *DescribeDomainsResponseBodyDomains) GetCertRegion() *string {
	return s.CertRegion
}

func (s *DescribeDomainsResponseBodyDomains) GetCname() *string {
	return s.Cname
}

func (s *DescribeDomainsResponseBodyDomains) GetDomain() *string {
	return s.Domain
}

func (s *DescribeDomainsResponseBodyDomains) GetHttp2Enable() *bool {
	return s.Http2Enable
}

func (s *DescribeDomainsResponseBodyDomains) GetProxyTypeList() []*DescribeDomainsResponseBodyDomainsProxyTypeList {
	return s.ProxyTypeList
}

func (s *DescribeDomainsResponseBodyDomains) GetRealServers() []*DescribeDomainsResponseBodyDomainsRealServers {
	return s.RealServers
}

func (s *DescribeDomainsResponseBodyDomains) GetSslCiphers() *string {
	return s.SslCiphers
}

func (s *DescribeDomainsResponseBodyDomains) GetSslProtocols() *string {
	return s.SslProtocols
}

func (s *DescribeDomainsResponseBodyDomains) GetWhiteList() []*string {
	return s.WhiteList
}

func (s *DescribeDomainsResponseBodyDomains) SetBlackList(v []*string) *DescribeDomainsResponseBodyDomains {
	s.BlackList = v
	return s
}

func (s *DescribeDomainsResponseBodyDomains) SetCcEnabled(v bool) *DescribeDomainsResponseBodyDomains {
	s.CcEnabled = &v
	return s
}

func (s *DescribeDomainsResponseBodyDomains) SetCcRuleEnabled(v bool) *DescribeDomainsResponseBodyDomains {
	s.CcRuleEnabled = &v
	return s
}

func (s *DescribeDomainsResponseBodyDomains) SetCcTemplate(v string) *DescribeDomainsResponseBodyDomains {
	s.CcTemplate = &v
	return s
}

func (s *DescribeDomainsResponseBodyDomains) SetCertName(v string) *DescribeDomainsResponseBodyDomains {
	s.CertName = &v
	return s
}

func (s *DescribeDomainsResponseBodyDomains) SetCertRegion(v string) *DescribeDomainsResponseBodyDomains {
	s.CertRegion = &v
	return s
}

func (s *DescribeDomainsResponseBodyDomains) SetCname(v string) *DescribeDomainsResponseBodyDomains {
	s.Cname = &v
	return s
}

func (s *DescribeDomainsResponseBodyDomains) SetDomain(v string) *DescribeDomainsResponseBodyDomains {
	s.Domain = &v
	return s
}

func (s *DescribeDomainsResponseBodyDomains) SetHttp2Enable(v bool) *DescribeDomainsResponseBodyDomains {
	s.Http2Enable = &v
	return s
}

func (s *DescribeDomainsResponseBodyDomains) SetProxyTypeList(v []*DescribeDomainsResponseBodyDomainsProxyTypeList) *DescribeDomainsResponseBodyDomains {
	s.ProxyTypeList = v
	return s
}

func (s *DescribeDomainsResponseBodyDomains) SetRealServers(v []*DescribeDomainsResponseBodyDomainsRealServers) *DescribeDomainsResponseBodyDomains {
	s.RealServers = v
	return s
}

func (s *DescribeDomainsResponseBodyDomains) SetSslCiphers(v string) *DescribeDomainsResponseBodyDomains {
	s.SslCiphers = &v
	return s
}

func (s *DescribeDomainsResponseBodyDomains) SetSslProtocols(v string) *DescribeDomainsResponseBodyDomains {
	s.SslProtocols = &v
	return s
}

func (s *DescribeDomainsResponseBodyDomains) SetWhiteList(v []*string) *DescribeDomainsResponseBodyDomains {
	s.WhiteList = v
	return s
}

func (s *DescribeDomainsResponseBodyDomains) Validate() error {
	return dara.Validate(s)
}

type DescribeDomainsResponseBodyDomainsProxyTypeList struct {
	ProxyPorts []*string `json:"ProxyPorts,omitempty" xml:"ProxyPorts,omitempty" type:"Repeated"`
	// example:
	//
	// http
	ProxyType *string `json:"ProxyType,omitempty" xml:"ProxyType,omitempty"`
}

func (s DescribeDomainsResponseBodyDomainsProxyTypeList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainsResponseBodyDomainsProxyTypeList) GoString() string {
	return s.String()
}

func (s *DescribeDomainsResponseBodyDomainsProxyTypeList) GetProxyPorts() []*string {
	return s.ProxyPorts
}

func (s *DescribeDomainsResponseBodyDomainsProxyTypeList) GetProxyType() *string {
	return s.ProxyType
}

func (s *DescribeDomainsResponseBodyDomainsProxyTypeList) SetProxyPorts(v []*string) *DescribeDomainsResponseBodyDomainsProxyTypeList {
	s.ProxyPorts = v
	return s
}

func (s *DescribeDomainsResponseBodyDomainsProxyTypeList) SetProxyType(v string) *DescribeDomainsResponseBodyDomainsProxyTypeList {
	s.ProxyType = &v
	return s
}

func (s *DescribeDomainsResponseBodyDomainsProxyTypeList) Validate() error {
	return dara.Validate(s)
}

type DescribeDomainsResponseBodyDomainsRealServers struct {
	// example:
	//
	// 1.1.1.1
	RealServer *string `json:"RealServer,omitempty" xml:"RealServer,omitempty"`
	// example:
	//
	// 0
	RsType *int32 `json:"RsType,omitempty" xml:"RsType,omitempty"`
}

func (s DescribeDomainsResponseBodyDomainsRealServers) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainsResponseBodyDomainsRealServers) GoString() string {
	return s.String()
}

func (s *DescribeDomainsResponseBodyDomainsRealServers) GetRealServer() *string {
	return s.RealServer
}

func (s *DescribeDomainsResponseBodyDomainsRealServers) GetRsType() *int32 {
	return s.RsType
}

func (s *DescribeDomainsResponseBodyDomainsRealServers) SetRealServer(v string) *DescribeDomainsResponseBodyDomainsRealServers {
	s.RealServer = &v
	return s
}

func (s *DescribeDomainsResponseBodyDomainsRealServers) SetRsType(v int32) *DescribeDomainsResponseBodyDomainsRealServers {
	s.RsType = &v
	return s
}

func (s *DescribeDomainsResponseBodyDomainsRealServers) Validate() error {
	return dara.Validate(s)
}
