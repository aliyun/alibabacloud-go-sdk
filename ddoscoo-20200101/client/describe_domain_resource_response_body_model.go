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
	// The request ID.
	//
	// example:
	//
	// 39499F01-19D9-4EA4-A0E9-C6014BA5CDBE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of web rules returned.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// An array of web rule configurations.
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
	if s.WebRules != nil {
		for _, item := range s.WebRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDomainResourceResponseBodyWebRules struct {
	// An array of IP addresses on the blocklist for the domain.
	BlackList []*string `json:"BlackList,omitempty" xml:"BlackList,omitempty" type:"Repeated"`
	// Whether CC protection is enabled. Valid values:
	//
	// - **true**: Enabled.
	//
	// - **false**: Disabled.
	//
	// example:
	//
	// true
	CcEnabled *bool `json:"CcEnabled,omitempty" xml:"CcEnabled,omitempty"`
	// Whether custom CC protection rules are enabled. Valid values:
	//
	// - **true**: Enabled.
	//
	// - **false**: Disabled.
	//
	// example:
	//
	// true
	CcRuleEnabled *bool `json:"CcRuleEnabled,omitempty" xml:"CcRuleEnabled,omitempty"`
	// The CC protection mode. Valid values:
	//
	// - **default**: normal mode
	//
	// - **gf_under_attack**: Emergency mode
	//
	// - **gf_sos_verify**: Strict mode
	//
	// - **gf_sos_enhance**: Super Strict mode
	//
	// example:
	//
	// default
	CcTemplate *string `json:"CcTemplate,omitempty" xml:"CcTemplate,omitempty"`
	// The name of the SSL certificate used by the domain.
	//
	// example:
	//
	// 49944XX.pem
	CertName *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	// The CNAME assigned to the domain by Anti-DDoS Pro.
	//
	// example:
	//
	// 0ekb69x3j9wvXXXX.aliyunddosXXXX.com
	Cname *string `json:"Cname,omitempty" xml:"Cname,omitempty"`
	// An array of custom cipher suites.
	CustomCiphers []*string `json:"CustomCiphers,omitempty" xml:"CustomCiphers,omitempty" type:"Repeated"`
	// The domain name.
	//
	// example:
	//
	// www.example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// Whether HTTP/2 is enabled. Valid values:
	//
	// - **true**: Enabled.
	//
	// - **false**: Disabled.
	//
	// example:
	//
	// false
	Http2Enable *bool `json:"Http2Enable,omitempty" xml:"Http2Enable,omitempty"`
	// Whether forced HTTPS redirect is enabled. Valid values:
	//
	// - **true**: Enabled.
	//
	// - **false**: Disabled.
	//
	// example:
	//
	// false
	Http2HttpsEnable *bool `json:"Http2HttpsEnable,omitempty" xml:"Http2HttpsEnable,omitempty"`
	// Whether HTTP back-to-source is enabled. Valid values:
	//
	// - **true**: Enabled.
	//
	// - **false**: Disabled.
	//
	// example:
	//
	// false
	Https2HttpEnable *bool `json:"Https2HttpEnable,omitempty" xml:"Https2HttpEnable,omitempty"`
	// The advanced HTTPS settings, formatted as a JSON string. These settings apply only when the domain uses HTTPS (the `ProxyType` value includes `https`). The string contains the following fields:
	//
	// - **Http2https**: integer. Specifies whether to enable forced HTTPS redirect. Valid values: **0*	- (disabled) and **1*	- (enabled).
	//
	// - **Https2http**: integer. Specifies whether to enable HTTP back-to-source. Valid values: **0*	- (disabled) and **1*	- (enabled).
	//
	// - **Http2**: integer. Specifies whether to enable HTTP/2. Valid values: **0*	- (disabled) and **1*	- (enabled).
	//
	// example:
	//
	// {"Https2http":0,"Http2":0,"Http2https":0}
	HttpsExt *string `json:"HttpsExt,omitempty" xml:"HttpsExt,omitempty"`
	// An array of IDs for the Anti-DDoS Pro instances associated with the domain.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// Whether Online Certificate Status Protocol (OCSP) is enabled. Valid values:
	//
	// - **true**: Enabled.
	//
	// - **false**: Disabled.
	//
	// example:
	//
	// false
	OcspEnabled *bool `json:"OcspEnabled,omitempty" xml:"OcspEnabled,omitempty"`
	// The back-to-source load balancing algorithm. This algorithm determines how traffic is distributed to origin servers. Valid values:
	//
	// - **ip_hash**: IP Hash. This algorithm hashes the source IP address and routes all requests from the same IP address to the same origin server.
	//
	// - **rr**: Round Robin. This algorithm distributes requests across origin servers in turn.
	//
	// - **least_time**: Least Time. This algorithm uses intelligent DNS resolution to minimize end-to-end latency from the protection node to the origin server.
	//
	// example:
	//
	// ip_hash
	PolicyMode *string `json:"PolicyMode,omitempty" xml:"PolicyMode,omitempty"`
	// Whether Anti-DDoS Pro forwards traffic for the domain. Valid values:
	//
	// - **true**: Forwarding is enabled.
	//
	// - **false**: Forwarding is disabled.
	//
	// example:
	//
	// true
	ProxyEnabled *bool `json:"ProxyEnabled,omitempty" xml:"ProxyEnabled,omitempty"`
	// An array of proxy configurations for the domain, each specifying a protocol and its associated ports.
	ProxyTypes []*DescribeDomainResourceResponseBodyWebRulesProxyTypes `json:"ProxyTypes,omitempty" xml:"ProxyTypes,omitempty" type:"Repeated"`
	// The reason the domain is penalized. Valid values:
	//
	// - **1**: The domain does not have an ICP filing.
	//
	// - **2**: The business hosted on the domain does not meet regulatory requirements.
	//
	// If both reasons apply, this parameter returns **2**.
	//
	// example:
	//
	// 1
	PunishReason *int32 `json:"PunishReason,omitempty" xml:"PunishReason,omitempty"`
	// Whether the domain is penalized. Valid values:
	//
	// - **true**: The domain is penalized. For more information, see the `PunishReason` parameter.
	//
	// - **false**: The domain is not penalized.
	//
	// example:
	//
	// false
	PunishStatus *bool `json:"PunishStatus,omitempty" xml:"PunishStatus,omitempty"`
	// An array of origin server addresses.
	RealServers []*string `json:"RealServers,omitempty" xml:"RealServers,omitempty" type:"Repeated"`
	// The address type of the origin server. Valid values:
	//
	// - **0**: IP address
	//
	// - **1**: domain
	//
	// example:
	//
	// 0
	RsType *int32 `json:"RsType,omitempty" xml:"RsType,omitempty"`
	// Whether TLS 1.3 is enabled. Valid values:
	//
	// - **true**: Enabled.
	//
	// - **false**: Disabled.
	//
	// example:
	//
	// false
	Ssl13Enabled *bool `json:"Ssl13Enabled,omitempty" xml:"Ssl13Enabled,omitempty"`
	// The cipher suite configuration type. Valid values:
	//
	// - **default**: Custom
	//
	// - **all**: All
	//
	// - **strong**: Strong
	//
	// - **improved**: Improved
	//
	// example:
	//
	// default
	SslCiphers *string `json:"SslCiphers,omitempty" xml:"SslCiphers,omitempty"`
	// The minimum TLS protocol version. Valid values:
	//
	// - **tls1.0**: TLS 1.0 and later versions.
	//
	// - **tls1.1**: TLS 1.1 and later versions.
	//
	// - **tls1.2**: TLS 1.2 and later versions.
	//
	// example:
	//
	// tls1.0
	SslProtocols *string `json:"SslProtocols,omitempty" xml:"SslProtocols,omitempty"`
	// An array of IP addresses on the allowlist for the domain.
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
	if s.ProxyTypes != nil {
		for _, item := range s.ProxyTypes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDomainResourceResponseBodyWebRulesProxyTypes struct {
	// The array of ports associated with the protocol.
	ProxyPorts []*string `json:"ProxyPorts,omitempty" xml:"ProxyPorts,omitempty" type:"Repeated"`
	// The proxy protocol. Valid values:
	//
	// - **http**: HTTP
	//
	// - **https**: HTTPS
	//
	// - **websocket**: WebSocket
	//
	// - **websockets**: WebSockets
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
