// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeDomainResourceResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribeDomainResourceResponseBody
	GetTotalCount() *int64
	SetWebRules(v []*DescribeDomainResourceResponseBodyWebRules) *DescribeDomainResourceResponseBody
	GetWebRules() []*DescribeDomainResourceResponseBodyWebRules
}

type DescribeDomainResourceResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 39499F01-19D9-4EA4-A0E9-C6014BA5CDBE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of forwarding rules.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The configurations of the forwarding rule.
	WebRules []*DescribeDomainResourceResponseBodyWebRules `json:"WebRules,omitempty" xml:"WebRules,omitempty" type:"Repeated"`
}

func (s DescribeDomainResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainResourceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDomainResourceResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeDomainResourceResponseBody) GetWebRules() []*DescribeDomainResourceResponseBodyWebRules {
	return s.WebRules
}

func (s *DescribeDomainResourceResponseBody) SetRequestId(v string) *DescribeDomainResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainResourceResponseBody) SetTotalCount(v int64) *DescribeDomainResourceResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDomainResourceResponseBody) SetWebRules(v []*DescribeDomainResourceResponseBodyWebRules) *DescribeDomainResourceResponseBody {
	s.WebRules = v
	return s
}

func (s *DescribeDomainResourceResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDomainResourceResponseBodyWebRules struct {
	// The IP addresses that are included in the blacklist of the domain name.
	BlackList []*string `json:"BlackList,omitempty" xml:"BlackList,omitempty" type:"Repeated"`
	// Indicates whether Frequency Control is enabled. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	CcEnabled *bool `json:"CcEnabled,omitempty" xml:"CcEnabled,omitempty"`
	// Indicates whether the Custom Rules switch of Frequency Control is turned on. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	CcRuleEnabled *bool `json:"CcRuleEnabled,omitempty" xml:"CcRuleEnabled,omitempty"`
	// The mode of Frequency Control. Valid values:
	//
	// 	- **default**: the Normal mode
	//
	// 	- **gf_under_attack**: the Emergency mode
	//
	// 	- **gf_sos_verify**: the Strict mode
	//
	// 	- **gf_sos_verify**: the Super Strict mode
	//
	// example:
	//
	// default
	CcTemplate *string `json:"CcTemplate,omitempty" xml:"CcTemplate,omitempty"`
	// The name of the SSL certificate used by the domain name.
	//
	// example:
	//
	// 49944XX.pem
	CertName *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	// The CNAME provided by the instance to which the domain name is added.
	//
	// example:
	//
	// 0ekb69x3j9wvXXXX.aliyunddosXXXX.com
	Cname *string `json:"Cname,omitempty" xml:"Cname,omitempty"`
	// The custom cipher suites.
	CustomCiphers []*string `json:"CustomCiphers,omitempty" xml:"CustomCiphers,omitempty" type:"Repeated"`
	// The domain name of the website.
	//
	// example:
	//
	// www.example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// Indicates whether Enable HTTP/2 is turned on. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	Http2Enable *bool `json:"Http2Enable,omitempty" xml:"Http2Enable,omitempty"`
	// Indicates whether Enable HTTPS Redirection is turned on. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	Http2HttpsEnable *bool `json:"Http2HttpsEnable,omitempty" xml:"Http2HttpsEnable,omitempty"`
	// Indicates whether Enable HTTP Redirection of Back-to-origin Requests is turned on. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	Https2HttpEnable *bool `json:"Https2HttpEnable,omitempty" xml:"Https2HttpEnable,omitempty"`
	// The advanced HTTPS settings. This parameter takes effect only when the value of the **ProxyType*	- parameter includes **https**. The value is a string that consists of a JSON struct. The JSON struct contains the following fields:
	//
	// 	- **Http2https**: indicates whether Enable HTTPS Redirection is turned on. Data type: integer. Valid values: **0*	- and **1**. The value 0 indicates that Enable HTTPS Redirection is turned on. The value 1 indicates that Enable HTTPS Redirection is turned off.
	//
	// 	- **Https2http**: indicates whether Enable HTTP Redirection of Back-to-origin Requests is turned on. Data type: integer. Valid values: **0*	- and **1**. The value 0 indicates that the feature is turned on. The value 1 indicates that the feature is turned off.
	//
	// 	- **Http2**: indicates whether Enable HTTP/2 is turned on. Data type: integer. Valid values: **0*	- and **1**. The value 0 indicates that Enable HTTP/2 is turned off. The value 1 indicates that Enable HTTP/2 is turned on.
	//
	// example:
	//
	// {"Https2http":0,"Http2":0,"Http2https":0}
	HttpsExt *string `json:"HttpsExt,omitempty" xml:"HttpsExt,omitempty"`
	// The IDs of the instances to which the domain name is added.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// Indicates whether the Online Certificate Status Protocol (OCSP) feature is turned on. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	OcspEnabled *bool `json:"OcspEnabled,omitempty" xml:"OcspEnabled,omitempty"`
	// The scheduling algorithm for back-to-origin traffic. Valid values:
	//
	// 	- **ip_hash**: the IP hash algorithm. This algorithm is used to redirect the requests from the same IP address to the same origin server.
	//
	// 	- **rr**: the round-robin algorithm. This algorithm is used to redirect requests to origin servers in turn.
	//
	// 	- **least_time**: the least response time algorithm. This algorithm is used to minimize the latency when requests are forwarded from the instance to origin servers based on the intelligent DNS resolution feature.
	//
	// example:
	//
	// ip_hash
	PolicyMode *string `json:"PolicyMode,omitempty" xml:"PolicyMode,omitempty"`
	// Indicates whether the instance forwards the traffic that is destined for the website. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	ProxyEnabled *bool `json:"ProxyEnabled,omitempty" xml:"ProxyEnabled,omitempty"`
	// The details about the protocol type and port number.
	ProxyTypes []*DescribeDomainResourceResponseBodyWebRulesProxyTypes `json:"ProxyTypes,omitempty" xml:"ProxyTypes,omitempty" type:"Repeated"`
	// The reason why the domain name is invalid. Valid values:
	//
	// 	- **1**: No Content Provider (ICP) filing is completed for the domain name.
	//
	// 	- **2**: The business for which you registered the domain name does not meet regulatory requirements.
	//
	// If the two reasons are both involved, the value **2*	- is returned.
	//
	// example:
	//
	// 1
	PunishReason *int32 `json:"PunishReason,omitempty" xml:"PunishReason,omitempty"`
	// Indicates whether the domain name is invalid. Valid values:
	//
	// 	- **true**: The domain name is invalid. You can view the specific reasons from the **PunishReason*	- parameter.
	//
	// 	- **false**: The domain name is valid.
	//
	// example:
	//
	// false
	PunishStatus *bool `json:"PunishStatus,omitempty" xml:"PunishStatus,omitempty"`
	// The addresses of origin servers.
	RealServers []*string `json:"RealServers,omitempty" xml:"RealServers,omitempty" type:"Repeated"`
	// The address type of the origin server. Valid values:
	//
	// 	- **0**: IP address
	//
	// 	- **1**: domain name
	//
	// example:
	//
	// 0
	RsType *int32 `json:"RsType,omitempty" xml:"RsType,omitempty"`
	// Indicates whether TLS 1.3 is supported. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	Ssl13Enabled *bool `json:"Ssl13Enabled,omitempty" xml:"Ssl13Enabled,omitempty"`
	// The type of the cipher suite. Valid values:
	//
	// 	- **default**: custom cipher suite
	//
	// 	- **all**: all cipher suites
	//
	// 	- **strong**: strong cipher suites
	//
	// example:
	//
	// default
	SslCiphers *string `json:"SslCiphers,omitempty" xml:"SslCiphers,omitempty"`
	// The version of the TLS protocol. Valid values:
	//
	// 	- **tls1.0**: TLS 1.0 or later
	//
	// 	- **tls1.1**: TLS 1.1 or later
	//
	// 	- **tls1.2**: TLS 1.2 or later
	//
	// example:
	//
	// tls1.0
	SslProtocols *string `json:"SslProtocols,omitempty" xml:"SslProtocols,omitempty"`
	// The IP addresses that are included in the whitelist of the domain name.
	WhiteList []*string `json:"WhiteList,omitempty" xml:"WhiteList,omitempty" type:"Repeated"`
}

func (s DescribeDomainResourceResponseBodyWebRules) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainResourceResponseBodyWebRules) GoString() string {
	return s.String()
}

func (s *DescribeDomainResourceResponseBodyWebRules) GetBlackList() []*string {
	return s.BlackList
}

func (s *DescribeDomainResourceResponseBodyWebRules) GetCcEnabled() *bool {
	return s.CcEnabled
}

func (s *DescribeDomainResourceResponseBodyWebRules) GetCcRuleEnabled() *bool {
	return s.CcRuleEnabled
}

func (s *DescribeDomainResourceResponseBodyWebRules) GetCcTemplate() *string {
	return s.CcTemplate
}

func (s *DescribeDomainResourceResponseBodyWebRules) GetCertName() *string {
	return s.CertName
}

func (s *DescribeDomainResourceResponseBodyWebRules) GetCname() *string {
	return s.Cname
}

func (s *DescribeDomainResourceResponseBodyWebRules) GetCustomCiphers() []*string {
	return s.CustomCiphers
}

func (s *DescribeDomainResourceResponseBodyWebRules) GetDomain() *string {
	return s.Domain
}

func (s *DescribeDomainResourceResponseBodyWebRules) GetHttp2Enable() *bool {
	return s.Http2Enable
}

func (s *DescribeDomainResourceResponseBodyWebRules) GetHttp2HttpsEnable() *bool {
	return s.Http2HttpsEnable
}

func (s *DescribeDomainResourceResponseBodyWebRules) GetHttps2HttpEnable() *bool {
	return s.Https2HttpEnable
}

func (s *DescribeDomainResourceResponseBodyWebRules) GetHttpsExt() *string {
	return s.HttpsExt
}

func (s *DescribeDomainResourceResponseBodyWebRules) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *DescribeDomainResourceResponseBodyWebRules) GetOcspEnabled() *bool {
	return s.OcspEnabled
}

func (s *DescribeDomainResourceResponseBodyWebRules) GetPolicyMode() *string {
	return s.PolicyMode
}

func (s *DescribeDomainResourceResponseBodyWebRules) GetProxyEnabled() *bool {
	return s.ProxyEnabled
}

func (s *DescribeDomainResourceResponseBodyWebRules) GetProxyTypes() []*DescribeDomainResourceResponseBodyWebRulesProxyTypes {
	return s.ProxyTypes
}

func (s *DescribeDomainResourceResponseBodyWebRules) GetPunishReason() *int32 {
	return s.PunishReason
}

func (s *DescribeDomainResourceResponseBodyWebRules) GetPunishStatus() *bool {
	return s.PunishStatus
}

func (s *DescribeDomainResourceResponseBodyWebRules) GetRealServers() []*string {
	return s.RealServers
}

func (s *DescribeDomainResourceResponseBodyWebRules) GetRsType() *int32 {
	return s.RsType
}

func (s *DescribeDomainResourceResponseBodyWebRules) GetSsl13Enabled() *bool {
	return s.Ssl13Enabled
}

func (s *DescribeDomainResourceResponseBodyWebRules) GetSslCiphers() *string {
	return s.SslCiphers
}

func (s *DescribeDomainResourceResponseBodyWebRules) GetSslProtocols() *string {
	return s.SslProtocols
}

func (s *DescribeDomainResourceResponseBodyWebRules) GetWhiteList() []*string {
	return s.WhiteList
}

func (s *DescribeDomainResourceResponseBodyWebRules) SetBlackList(v []*string) *DescribeDomainResourceResponseBodyWebRules {
	s.BlackList = v
	return s
}

func (s *DescribeDomainResourceResponseBodyWebRules) SetCcEnabled(v bool) *DescribeDomainResourceResponseBodyWebRules {
	s.CcEnabled = &v
	return s
}

func (s *DescribeDomainResourceResponseBodyWebRules) SetCcRuleEnabled(v bool) *DescribeDomainResourceResponseBodyWebRules {
	s.CcRuleEnabled = &v
	return s
}

func (s *DescribeDomainResourceResponseBodyWebRules) SetCcTemplate(v string) *DescribeDomainResourceResponseBodyWebRules {
	s.CcTemplate = &v
	return s
}

func (s *DescribeDomainResourceResponseBodyWebRules) SetCertName(v string) *DescribeDomainResourceResponseBodyWebRules {
	s.CertName = &v
	return s
}

func (s *DescribeDomainResourceResponseBodyWebRules) SetCname(v string) *DescribeDomainResourceResponseBodyWebRules {
	s.Cname = &v
	return s
}

func (s *DescribeDomainResourceResponseBodyWebRules) SetCustomCiphers(v []*string) *DescribeDomainResourceResponseBodyWebRules {
	s.CustomCiphers = v
	return s
}

func (s *DescribeDomainResourceResponseBodyWebRules) SetDomain(v string) *DescribeDomainResourceResponseBodyWebRules {
	s.Domain = &v
	return s
}

func (s *DescribeDomainResourceResponseBodyWebRules) SetHttp2Enable(v bool) *DescribeDomainResourceResponseBodyWebRules {
	s.Http2Enable = &v
	return s
}

func (s *DescribeDomainResourceResponseBodyWebRules) SetHttp2HttpsEnable(v bool) *DescribeDomainResourceResponseBodyWebRules {
	s.Http2HttpsEnable = &v
	return s
}

func (s *DescribeDomainResourceResponseBodyWebRules) SetHttps2HttpEnable(v bool) *DescribeDomainResourceResponseBodyWebRules {
	s.Https2HttpEnable = &v
	return s
}

func (s *DescribeDomainResourceResponseBodyWebRules) SetHttpsExt(v string) *DescribeDomainResourceResponseBodyWebRules {
	s.HttpsExt = &v
	return s
}

func (s *DescribeDomainResourceResponseBodyWebRules) SetInstanceIds(v []*string) *DescribeDomainResourceResponseBodyWebRules {
	s.InstanceIds = v
	return s
}

func (s *DescribeDomainResourceResponseBodyWebRules) SetOcspEnabled(v bool) *DescribeDomainResourceResponseBodyWebRules {
	s.OcspEnabled = &v
	return s
}

func (s *DescribeDomainResourceResponseBodyWebRules) SetPolicyMode(v string) *DescribeDomainResourceResponseBodyWebRules {
	s.PolicyMode = &v
	return s
}

func (s *DescribeDomainResourceResponseBodyWebRules) SetProxyEnabled(v bool) *DescribeDomainResourceResponseBodyWebRules {
	s.ProxyEnabled = &v
	return s
}

func (s *DescribeDomainResourceResponseBodyWebRules) SetProxyTypes(v []*DescribeDomainResourceResponseBodyWebRulesProxyTypes) *DescribeDomainResourceResponseBodyWebRules {
	s.ProxyTypes = v
	return s
}

func (s *DescribeDomainResourceResponseBodyWebRules) SetPunishReason(v int32) *DescribeDomainResourceResponseBodyWebRules {
	s.PunishReason = &v
	return s
}

func (s *DescribeDomainResourceResponseBodyWebRules) SetPunishStatus(v bool) *DescribeDomainResourceResponseBodyWebRules {
	s.PunishStatus = &v
	return s
}

func (s *DescribeDomainResourceResponseBodyWebRules) SetRealServers(v []*string) *DescribeDomainResourceResponseBodyWebRules {
	s.RealServers = v
	return s
}

func (s *DescribeDomainResourceResponseBodyWebRules) SetRsType(v int32) *DescribeDomainResourceResponseBodyWebRules {
	s.RsType = &v
	return s
}

func (s *DescribeDomainResourceResponseBodyWebRules) SetSsl13Enabled(v bool) *DescribeDomainResourceResponseBodyWebRules {
	s.Ssl13Enabled = &v
	return s
}

func (s *DescribeDomainResourceResponseBodyWebRules) SetSslCiphers(v string) *DescribeDomainResourceResponseBodyWebRules {
	s.SslCiphers = &v
	return s
}

func (s *DescribeDomainResourceResponseBodyWebRules) SetSslProtocols(v string) *DescribeDomainResourceResponseBodyWebRules {
	s.SslProtocols = &v
	return s
}

func (s *DescribeDomainResourceResponseBodyWebRules) SetWhiteList(v []*string) *DescribeDomainResourceResponseBodyWebRules {
	s.WhiteList = v
	return s
}

func (s *DescribeDomainResourceResponseBodyWebRules) Validate() error {
	return dara.Validate(s)
}

type DescribeDomainResourceResponseBodyWebRulesProxyTypes struct {
	// The port numbers.
	ProxyPorts []*string `json:"ProxyPorts,omitempty" xml:"ProxyPorts,omitempty" type:"Repeated"`
	// The type of the protocol. Valid values:
	//
	// 	- **http**
	//
	// 	- **https**
	//
	// 	- **websocket**
	//
	// 	- **websockets**
	//
	// example:
	//
	// http
	ProxyType *string `json:"ProxyType,omitempty" xml:"ProxyType,omitempty"`
}

func (s DescribeDomainResourceResponseBodyWebRulesProxyTypes) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainResourceResponseBodyWebRulesProxyTypes) GoString() string {
	return s.String()
}

func (s *DescribeDomainResourceResponseBodyWebRulesProxyTypes) GetProxyPorts() []*string {
	return s.ProxyPorts
}

func (s *DescribeDomainResourceResponseBodyWebRulesProxyTypes) GetProxyType() *string {
	return s.ProxyType
}

func (s *DescribeDomainResourceResponseBodyWebRulesProxyTypes) SetProxyPorts(v []*string) *DescribeDomainResourceResponseBodyWebRulesProxyTypes {
	s.ProxyPorts = v
	return s
}

func (s *DescribeDomainResourceResponseBodyWebRulesProxyTypes) SetProxyType(v string) *DescribeDomainResourceResponseBodyWebRulesProxyTypes {
	s.ProxyType = &v
	return s
}

func (s *DescribeDomainResourceResponseBodyWebRulesProxyTypes) Validate() error {
	return dara.Validate(s)
}
