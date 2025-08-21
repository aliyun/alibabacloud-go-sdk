// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWebRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeWebRulesResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribeWebRulesResponseBody
	GetTotalCount() *int64
	SetWebRules(v []*DescribeWebRulesResponseBodyWebRules) *DescribeWebRulesResponseBody
	GetWebRules() []*DescribeWebRulesResponseBodyWebRules
}

type DescribeWebRulesResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0F5B72DD-96F4-423A-B12B-A5151DD746B8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of queried website business forwarding rules.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The configurations of the forwarding rule.
	WebRules []*DescribeWebRulesResponseBodyWebRules `json:"WebRules,omitempty" xml:"WebRules,omitempty" type:"Repeated"`
}

func (s DescribeWebRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeWebRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeWebRulesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeWebRulesResponseBody) GetWebRules() []*DescribeWebRulesResponseBodyWebRules {
	return s.WebRules
}

func (s *DescribeWebRulesResponseBody) SetRequestId(v string) *DescribeWebRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeWebRulesResponseBody) SetTotalCount(v int64) *DescribeWebRulesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeWebRulesResponseBody) SetWebRules(v []*DescribeWebRulesResponseBodyWebRules) *DescribeWebRulesResponseBody {
	s.WebRules = v
	return s
}

func (s *DescribeWebRulesResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeWebRulesResponseBodyWebRules struct {
	// The IP addresses in the blacklist for the domain name.
	BlackList []*string `json:"BlackList,omitempty" xml:"BlackList,omitempty" type:"Repeated"`
	// Indicates whether the Frequency Control policy is enabled. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	CcEnabled *bool `json:"CcEnabled,omitempty" xml:"CcEnabled,omitempty"`
	// Indicates whether the Custom Rule switch of the Frequency Control policy is turned on. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	CcRuleEnabled *bool `json:"CcRuleEnabled,omitempty" xml:"CcRuleEnabled,omitempty"`
	// The mode of the Frequency Control policy. Valid values:
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
	// The name of the SSL certificate.
	//
	// example:
	//
	// testcert
	CertName *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	// The region where the certificate is used. Valid values:
	//
	// 	- cn-hangzhou (default): the Chinese mainland
	//
	// 	- ap-southeast-1: outside the Chinese mainland
	//
	// example:
	//
	// cn-hangzhou
	CertRegion *string `json:"CertRegion,omitempty" xml:"CertRegion,omitempty"`
	// The CNAME provided by the Anti-DDoS Pro or Anti-DDoS Premium instance to which the domain name is added.
	//
	// example:
	//
	// kzmk7b8tt351****.aliyunddos1014****
	Cname *string `json:"Cname,omitempty" xml:"Cname,omitempty"`
	// The custom cipher suites.
	CustomCiphers []*string `json:"CustomCiphers,omitempty" xml:"CustomCiphers,omitempty" type:"Repeated"`
	// The domain name of the website.
	//
	// example:
	//
	// example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The SM certificate settings.
	GmCert *DescribeWebRulesResponseBodyWebRulesGmCert `json:"GmCert,omitempty" xml:"GmCert,omitempty" type:"Struct"`
	// Indicates whether Enable HTTP/2 is turned on. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Http2Enable *bool `json:"Http2Enable,omitempty" xml:"Http2Enable,omitempty"`
	// Indicates whether Enable HTTPS Redirection was turned on. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Http2HttpsEnable *bool `json:"Http2HttpsEnable,omitempty" xml:"Http2HttpsEnable,omitempty"`
	// Indicates whether Enable HTTP Redirection of Back-to-origin Requests is turned on. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Https2HttpEnable *bool `json:"Https2HttpEnable,omitempty" xml:"Https2HttpEnable,omitempty"`
	// Indicates whether the Online Certificate Status Protocol (OCSP) feature is enabled. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	OcspEnabled *bool `json:"OcspEnabled,omitempty" xml:"OcspEnabled,omitempty"`
	// The load balancing algorithm for back-to-origin traffic. Valid values:
	//
	// 	- **ip_hash**: the IP hash algorithm. This algorithm is used to redirect the requests from the same IP address to the same origin server.
	//
	// 	- **rr**: the round-robin algorithm. This algorithm is used to redirect requests to origin servers in turn.
	//
	// 	- **least_time**: the least response time algorithm. This algorithm is used to minimize the latency when requests are forwarded from Anti-DDoS Pro or Anti-DDoS Premium instances to origin servers based on the intelligent DNS resolution feature.
	//
	// example:
	//
	// ip_hash
	PolicyMode *string `json:"PolicyMode,omitempty" xml:"PolicyMode,omitempty"`
	// Indicates whether the forwarding rule is enabled. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	ProxyEnabled *bool `json:"ProxyEnabled,omitempty" xml:"ProxyEnabled,omitempty"`
	// The details of the protocol type and port number.
	ProxyTypes []*DescribeWebRulesResponseBodyWebRulesProxyTypes `json:"ProxyTypes,omitempty" xml:"ProxyTypes,omitempty" type:"Repeated"`
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
	// 	- **true**: You can view the specific reasons from the **PunishReason*	- parameter.
	//
	// 	- **false**
	//
	// example:
	//
	// true
	PunishStatus *bool `json:"PunishStatus,omitempty" xml:"PunishStatus,omitempty"`
	// The details of the origin server address.
	RealServers []*DescribeWebRulesResponseBodyWebRulesRealServers `json:"RealServers,omitempty" xml:"RealServers,omitempty" type:"Repeated"`
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
	// 	- **default**: custom cipher suites
	//
	// 	- **all**: all cipher suites, which contain strong and weak cipher suites
	//
	// 	- **strong**: strong cipher suites
	//
	// example:
	//
	// default
	SslCiphers *string `json:"SslCiphers,omitempty" xml:"SslCiphers,omitempty"`
	// The version of the Transport Layer Security (TLS) protocol. Valid values:
	//
	// 	- **tls1.0**: TLS 1.0 or later
	//
	// 	- **tls1.1**: TLS 1.1 or later
	//
	// 	- **tls1.2**: TLS 1.2 or later
	//
	// example:
	//
	// tls1.1
	SslProtocols *string `json:"SslProtocols,omitempty" xml:"SslProtocols,omitempty"`
	// The name of the certificate uploaded by the user to the certificate center.
	//
	// example:
	//
	// test
	UserCertName *string `json:"UserCertName,omitempty" xml:"UserCertName,omitempty"`
	// The IP addresses in the whitelist for the domain name.
	WhiteList []*string `json:"WhiteList,omitempty" xml:"WhiteList,omitempty" type:"Repeated"`
}

func (s DescribeWebRulesResponseBodyWebRules) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebRulesResponseBodyWebRules) GoString() string {
	return s.String()
}

func (s *DescribeWebRulesResponseBodyWebRules) GetBlackList() []*string {
	return s.BlackList
}

func (s *DescribeWebRulesResponseBodyWebRules) GetCcEnabled() *bool {
	return s.CcEnabled
}

func (s *DescribeWebRulesResponseBodyWebRules) GetCcRuleEnabled() *bool {
	return s.CcRuleEnabled
}

func (s *DescribeWebRulesResponseBodyWebRules) GetCcTemplate() *string {
	return s.CcTemplate
}

func (s *DescribeWebRulesResponseBodyWebRules) GetCertName() *string {
	return s.CertName
}

func (s *DescribeWebRulesResponseBodyWebRules) GetCertRegion() *string {
	return s.CertRegion
}

func (s *DescribeWebRulesResponseBodyWebRules) GetCname() *string {
	return s.Cname
}

func (s *DescribeWebRulesResponseBodyWebRules) GetCustomCiphers() []*string {
	return s.CustomCiphers
}

func (s *DescribeWebRulesResponseBodyWebRules) GetDomain() *string {
	return s.Domain
}

func (s *DescribeWebRulesResponseBodyWebRules) GetGmCert() *DescribeWebRulesResponseBodyWebRulesGmCert {
	return s.GmCert
}

func (s *DescribeWebRulesResponseBodyWebRules) GetHttp2Enable() *bool {
	return s.Http2Enable
}

func (s *DescribeWebRulesResponseBodyWebRules) GetHttp2HttpsEnable() *bool {
	return s.Http2HttpsEnable
}

func (s *DescribeWebRulesResponseBodyWebRules) GetHttps2HttpEnable() *bool {
	return s.Https2HttpEnable
}

func (s *DescribeWebRulesResponseBodyWebRules) GetOcspEnabled() *bool {
	return s.OcspEnabled
}

func (s *DescribeWebRulesResponseBodyWebRules) GetPolicyMode() *string {
	return s.PolicyMode
}

func (s *DescribeWebRulesResponseBodyWebRules) GetProxyEnabled() *bool {
	return s.ProxyEnabled
}

func (s *DescribeWebRulesResponseBodyWebRules) GetProxyTypes() []*DescribeWebRulesResponseBodyWebRulesProxyTypes {
	return s.ProxyTypes
}

func (s *DescribeWebRulesResponseBodyWebRules) GetPunishReason() *int32 {
	return s.PunishReason
}

func (s *DescribeWebRulesResponseBodyWebRules) GetPunishStatus() *bool {
	return s.PunishStatus
}

func (s *DescribeWebRulesResponseBodyWebRules) GetRealServers() []*DescribeWebRulesResponseBodyWebRulesRealServers {
	return s.RealServers
}

func (s *DescribeWebRulesResponseBodyWebRules) GetSsl13Enabled() *bool {
	return s.Ssl13Enabled
}

func (s *DescribeWebRulesResponseBodyWebRules) GetSslCiphers() *string {
	return s.SslCiphers
}

func (s *DescribeWebRulesResponseBodyWebRules) GetSslProtocols() *string {
	return s.SslProtocols
}

func (s *DescribeWebRulesResponseBodyWebRules) GetUserCertName() *string {
	return s.UserCertName
}

func (s *DescribeWebRulesResponseBodyWebRules) GetWhiteList() []*string {
	return s.WhiteList
}

func (s *DescribeWebRulesResponseBodyWebRules) SetBlackList(v []*string) *DescribeWebRulesResponseBodyWebRules {
	s.BlackList = v
	return s
}

func (s *DescribeWebRulesResponseBodyWebRules) SetCcEnabled(v bool) *DescribeWebRulesResponseBodyWebRules {
	s.CcEnabled = &v
	return s
}

func (s *DescribeWebRulesResponseBodyWebRules) SetCcRuleEnabled(v bool) *DescribeWebRulesResponseBodyWebRules {
	s.CcRuleEnabled = &v
	return s
}

func (s *DescribeWebRulesResponseBodyWebRules) SetCcTemplate(v string) *DescribeWebRulesResponseBodyWebRules {
	s.CcTemplate = &v
	return s
}

func (s *DescribeWebRulesResponseBodyWebRules) SetCertName(v string) *DescribeWebRulesResponseBodyWebRules {
	s.CertName = &v
	return s
}

func (s *DescribeWebRulesResponseBodyWebRules) SetCertRegion(v string) *DescribeWebRulesResponseBodyWebRules {
	s.CertRegion = &v
	return s
}

func (s *DescribeWebRulesResponseBodyWebRules) SetCname(v string) *DescribeWebRulesResponseBodyWebRules {
	s.Cname = &v
	return s
}

func (s *DescribeWebRulesResponseBodyWebRules) SetCustomCiphers(v []*string) *DescribeWebRulesResponseBodyWebRules {
	s.CustomCiphers = v
	return s
}

func (s *DescribeWebRulesResponseBodyWebRules) SetDomain(v string) *DescribeWebRulesResponseBodyWebRules {
	s.Domain = &v
	return s
}

func (s *DescribeWebRulesResponseBodyWebRules) SetGmCert(v *DescribeWebRulesResponseBodyWebRulesGmCert) *DescribeWebRulesResponseBodyWebRules {
	s.GmCert = v
	return s
}

func (s *DescribeWebRulesResponseBodyWebRules) SetHttp2Enable(v bool) *DescribeWebRulesResponseBodyWebRules {
	s.Http2Enable = &v
	return s
}

func (s *DescribeWebRulesResponseBodyWebRules) SetHttp2HttpsEnable(v bool) *DescribeWebRulesResponseBodyWebRules {
	s.Http2HttpsEnable = &v
	return s
}

func (s *DescribeWebRulesResponseBodyWebRules) SetHttps2HttpEnable(v bool) *DescribeWebRulesResponseBodyWebRules {
	s.Https2HttpEnable = &v
	return s
}

func (s *DescribeWebRulesResponseBodyWebRules) SetOcspEnabled(v bool) *DescribeWebRulesResponseBodyWebRules {
	s.OcspEnabled = &v
	return s
}

func (s *DescribeWebRulesResponseBodyWebRules) SetPolicyMode(v string) *DescribeWebRulesResponseBodyWebRules {
	s.PolicyMode = &v
	return s
}

func (s *DescribeWebRulesResponseBodyWebRules) SetProxyEnabled(v bool) *DescribeWebRulesResponseBodyWebRules {
	s.ProxyEnabled = &v
	return s
}

func (s *DescribeWebRulesResponseBodyWebRules) SetProxyTypes(v []*DescribeWebRulesResponseBodyWebRulesProxyTypes) *DescribeWebRulesResponseBodyWebRules {
	s.ProxyTypes = v
	return s
}

func (s *DescribeWebRulesResponseBodyWebRules) SetPunishReason(v int32) *DescribeWebRulesResponseBodyWebRules {
	s.PunishReason = &v
	return s
}

func (s *DescribeWebRulesResponseBodyWebRules) SetPunishStatus(v bool) *DescribeWebRulesResponseBodyWebRules {
	s.PunishStatus = &v
	return s
}

func (s *DescribeWebRulesResponseBodyWebRules) SetRealServers(v []*DescribeWebRulesResponseBodyWebRulesRealServers) *DescribeWebRulesResponseBodyWebRules {
	s.RealServers = v
	return s
}

func (s *DescribeWebRulesResponseBodyWebRules) SetSsl13Enabled(v bool) *DescribeWebRulesResponseBodyWebRules {
	s.Ssl13Enabled = &v
	return s
}

func (s *DescribeWebRulesResponseBodyWebRules) SetSslCiphers(v string) *DescribeWebRulesResponseBodyWebRules {
	s.SslCiphers = &v
	return s
}

func (s *DescribeWebRulesResponseBodyWebRules) SetSslProtocols(v string) *DescribeWebRulesResponseBodyWebRules {
	s.SslProtocols = &v
	return s
}

func (s *DescribeWebRulesResponseBodyWebRules) SetUserCertName(v string) *DescribeWebRulesResponseBodyWebRules {
	s.UserCertName = &v
	return s
}

func (s *DescribeWebRulesResponseBodyWebRules) SetWhiteList(v []*string) *DescribeWebRulesResponseBodyWebRules {
	s.WhiteList = v
	return s
}

func (s *DescribeWebRulesResponseBodyWebRules) Validate() error {
	return dara.Validate(s)
}

type DescribeWebRulesResponseBodyWebRulesGmCert struct {
	// The ID of the SM certificate.
	//
	// example:
	//
	// 725****
	CertId *string `json:"CertId,omitempty" xml:"CertId,omitempty"`
	// Indicates whether Enable SM Certificate-based Verification is turned on.
	//
	// 	- 0: no
	//
	// 	- 1: yes
	//
	// example:
	//
	// 1
	GmEnable *int64 `json:"GmEnable,omitempty" xml:"GmEnable,omitempty"`
	// Indicates whether Allow Access Only from SM Certificates-based Clients is turned on.
	//
	// 	- 0: no
	//
	// 	- 1: yes
	//
	// example:
	//
	// 1
	GmOnly *int64 `json:"GmOnly,omitempty" xml:"GmOnly,omitempty"`
}

func (s DescribeWebRulesResponseBodyWebRulesGmCert) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebRulesResponseBodyWebRulesGmCert) GoString() string {
	return s.String()
}

func (s *DescribeWebRulesResponseBodyWebRulesGmCert) GetCertId() *string {
	return s.CertId
}

func (s *DescribeWebRulesResponseBodyWebRulesGmCert) GetGmEnable() *int64 {
	return s.GmEnable
}

func (s *DescribeWebRulesResponseBodyWebRulesGmCert) GetGmOnly() *int64 {
	return s.GmOnly
}

func (s *DescribeWebRulesResponseBodyWebRulesGmCert) SetCertId(v string) *DescribeWebRulesResponseBodyWebRulesGmCert {
	s.CertId = &v
	return s
}

func (s *DescribeWebRulesResponseBodyWebRulesGmCert) SetGmEnable(v int64) *DescribeWebRulesResponseBodyWebRulesGmCert {
	s.GmEnable = &v
	return s
}

func (s *DescribeWebRulesResponseBodyWebRulesGmCert) SetGmOnly(v int64) *DescribeWebRulesResponseBodyWebRulesGmCert {
	s.GmOnly = &v
	return s
}

func (s *DescribeWebRulesResponseBodyWebRulesGmCert) Validate() error {
	return dara.Validate(s)
}

type DescribeWebRulesResponseBodyWebRulesProxyTypes struct {
	// The ports.
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
	// https
	ProxyType *string `json:"ProxyType,omitempty" xml:"ProxyType,omitempty"`
}

func (s DescribeWebRulesResponseBodyWebRulesProxyTypes) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebRulesResponseBodyWebRulesProxyTypes) GoString() string {
	return s.String()
}

func (s *DescribeWebRulesResponseBodyWebRulesProxyTypes) GetProxyPorts() []*string {
	return s.ProxyPorts
}

func (s *DescribeWebRulesResponseBodyWebRulesProxyTypes) GetProxyType() *string {
	return s.ProxyType
}

func (s *DescribeWebRulesResponseBodyWebRulesProxyTypes) SetProxyPorts(v []*string) *DescribeWebRulesResponseBodyWebRulesProxyTypes {
	s.ProxyPorts = v
	return s
}

func (s *DescribeWebRulesResponseBodyWebRulesProxyTypes) SetProxyType(v string) *DescribeWebRulesResponseBodyWebRulesProxyTypes {
	s.ProxyType = &v
	return s
}

func (s *DescribeWebRulesResponseBodyWebRulesProxyTypes) Validate() error {
	return dara.Validate(s)
}

type DescribeWebRulesResponseBodyWebRulesRealServers struct {
	// The address of the origin server.
	//
	// example:
	//
	// 192.0.XX.XX
	RealServer *string `json:"RealServer,omitempty" xml:"RealServer,omitempty"`
	// The type of the origin server address. Valid values:
	//
	// 	- **0**: IP address
	//
	// 	- **1**: domain name The domain name of the origin server is returned if you deploy proxies, such as Web Application Firewall (WAF), between the origin server and the instance. In this case, the address of the proxy, such as the CNAME provided by WAF, is returned.
	//
	// example:
	//
	// 0
	RsType *int32 `json:"RsType,omitempty" xml:"RsType,omitempty"`
}

func (s DescribeWebRulesResponseBodyWebRulesRealServers) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebRulesResponseBodyWebRulesRealServers) GoString() string {
	return s.String()
}

func (s *DescribeWebRulesResponseBodyWebRulesRealServers) GetRealServer() *string {
	return s.RealServer
}

func (s *DescribeWebRulesResponseBodyWebRulesRealServers) GetRsType() *int32 {
	return s.RsType
}

func (s *DescribeWebRulesResponseBodyWebRulesRealServers) SetRealServer(v string) *DescribeWebRulesResponseBodyWebRulesRealServers {
	s.RealServer = &v
	return s
}

func (s *DescribeWebRulesResponseBodyWebRulesRealServers) SetRsType(v int32) *DescribeWebRulesResponseBodyWebRulesRealServers {
	s.RsType = &v
	return s
}

func (s *DescribeWebRulesResponseBodyWebRulesRealServers) Validate() error {
	return dara.Validate(s)
}
