// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetLoadBalancerHTTPSListenerAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAclId(v string) *SetLoadBalancerHTTPSListenerAttributeRequest
	GetAclId() *string
	SetAclStatus(v string) *SetLoadBalancerHTTPSListenerAttributeRequest
	GetAclStatus() *string
	SetAclType(v string) *SetLoadBalancerHTTPSListenerAttributeRequest
	GetAclType() *string
	SetBandwidth(v int32) *SetLoadBalancerHTTPSListenerAttributeRequest
	GetBandwidth() *int32
	SetCACertificateId(v string) *SetLoadBalancerHTTPSListenerAttributeRequest
	GetCACertificateId() *string
	SetCookie(v string) *SetLoadBalancerHTTPSListenerAttributeRequest
	GetCookie() *string
	SetCookieTimeout(v int32) *SetLoadBalancerHTTPSListenerAttributeRequest
	GetCookieTimeout() *int32
	SetDescription(v string) *SetLoadBalancerHTTPSListenerAttributeRequest
	GetDescription() *string
	SetEnableHttp2(v string) *SetLoadBalancerHTTPSListenerAttributeRequest
	GetEnableHttp2() *string
	SetGzip(v string) *SetLoadBalancerHTTPSListenerAttributeRequest
	GetGzip() *string
	SetHealthCheck(v string) *SetLoadBalancerHTTPSListenerAttributeRequest
	GetHealthCheck() *string
	SetHealthCheckConnectPort(v int32) *SetLoadBalancerHTTPSListenerAttributeRequest
	GetHealthCheckConnectPort() *int32
	SetHealthCheckDomain(v string) *SetLoadBalancerHTTPSListenerAttributeRequest
	GetHealthCheckDomain() *string
	SetHealthCheckHttpCode(v string) *SetLoadBalancerHTTPSListenerAttributeRequest
	GetHealthCheckHttpCode() *string
	SetHealthCheckInterval(v int32) *SetLoadBalancerHTTPSListenerAttributeRequest
	GetHealthCheckInterval() *int32
	SetHealthCheckMethod(v string) *SetLoadBalancerHTTPSListenerAttributeRequest
	GetHealthCheckMethod() *string
	SetHealthCheckTimeout(v int32) *SetLoadBalancerHTTPSListenerAttributeRequest
	GetHealthCheckTimeout() *int32
	SetHealthCheckURI(v string) *SetLoadBalancerHTTPSListenerAttributeRequest
	GetHealthCheckURI() *string
	SetHealthyThreshold(v int32) *SetLoadBalancerHTTPSListenerAttributeRequest
	GetHealthyThreshold() *int32
	SetIdleTimeout(v int32) *SetLoadBalancerHTTPSListenerAttributeRequest
	GetIdleTimeout() *int32
	SetListenerPort(v int32) *SetLoadBalancerHTTPSListenerAttributeRequest
	GetListenerPort() *int32
	SetLoadBalancerId(v string) *SetLoadBalancerHTTPSListenerAttributeRequest
	GetLoadBalancerId() *string
	SetOwnerAccount(v string) *SetLoadBalancerHTTPSListenerAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *SetLoadBalancerHTTPSListenerAttributeRequest
	GetOwnerId() *int64
	SetRegionId(v string) *SetLoadBalancerHTTPSListenerAttributeRequest
	GetRegionId() *string
	SetRequestTimeout(v int32) *SetLoadBalancerHTTPSListenerAttributeRequest
	GetRequestTimeout() *int32
	SetResourceOwnerAccount(v string) *SetLoadBalancerHTTPSListenerAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *SetLoadBalancerHTTPSListenerAttributeRequest
	GetResourceOwnerId() *int64
	SetScheduler(v string) *SetLoadBalancerHTTPSListenerAttributeRequest
	GetScheduler() *string
	SetServerCertificateId(v string) *SetLoadBalancerHTTPSListenerAttributeRequest
	GetServerCertificateId() *string
	SetStickySession(v string) *SetLoadBalancerHTTPSListenerAttributeRequest
	GetStickySession() *string
	SetStickySessionType(v string) *SetLoadBalancerHTTPSListenerAttributeRequest
	GetStickySessionType() *string
	SetTLSCipherPolicy(v string) *SetLoadBalancerHTTPSListenerAttributeRequest
	GetTLSCipherPolicy() *string
	SetUnhealthyThreshold(v int32) *SetLoadBalancerHTTPSListenerAttributeRequest
	GetUnhealthyThreshold() *int32
	SetVServerGroup(v string) *SetLoadBalancerHTTPSListenerAttributeRequest
	GetVServerGroup() *string
	SetVServerGroupId(v string) *SetLoadBalancerHTTPSListenerAttributeRequest
	GetVServerGroupId() *string
	SetXForwardedFor(v string) *SetLoadBalancerHTTPSListenerAttributeRequest
	GetXForwardedFor() *string
	SetXForwardedFor_ClientSrcPort(v string) *SetLoadBalancerHTTPSListenerAttributeRequest
	GetXForwardedFor_ClientSrcPort() *string
	SetXForwardedFor_SLBID(v string) *SetLoadBalancerHTTPSListenerAttributeRequest
	GetXForwardedFor_SLBID() *string
	SetXForwardedFor_SLBIP(v string) *SetLoadBalancerHTTPSListenerAttributeRequest
	GetXForwardedFor_SLBIP() *string
	SetXForwardedFor_SLBPORT(v string) *SetLoadBalancerHTTPSListenerAttributeRequest
	GetXForwardedFor_SLBPORT() *string
	SetXForwardedFor_proto(v string) *SetLoadBalancerHTTPSListenerAttributeRequest
	GetXForwardedFor_proto() *string
}

type SetLoadBalancerHTTPSListenerAttributeRequest struct {
	// The ID of the network access control list (ACL) that is associated with the listener.
	//
	// This parameter is required if **AclStatus*	- is set to **on**.
	//
	// example:
	//
	// nacl-a2do9e413e0spzasx****
	AclId *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
	// Specifies whether to enable access control. Valid values:
	//
	// 	- **on**: enables access control
	//
	// 	- **off**: disables access control
	//
	// example:
	//
	// off
	AclStatus *string `json:"AclStatus,omitempty" xml:"AclStatus,omitempty"`
	// The type of network ACL. Valid values:
	//
	// 	- **white**: a whitelist. Only requests from the IP addresses or CIDR blocks in the network ACL are forwarded. Whitelists apply to scenarios where you want to allow only specific IP addresses to access an application. Your business may be adversely affected if the whitelist is not set properly. After a whitelist is configured, only IP addresses in the whitelist can access the CLB listener.
	//
	//     If no IP address is added to the whitelist, the CLB listener forwards all requests.
	//
	// 	- **black**: a blacklist. All requests from the IP addresses or CIDR blocks in the network ACL are denied. Blacklists apply to scenarios where you want to deny access from specified IP addresses to an application.
	//
	//     If no IP address is added to the blacklist, the listener forwards all requests.
	//
	// >  This parameter takes effect only when **AclStatus*	- is set to **on**.
	//
	// example:
	//
	// white
	AclType *string `json:"AclType,omitempty" xml:"AclType,omitempty"`
	// The maximum bandwidth of the listener. Unit: Mbit/s.
	//
	// Valid values:
	//
	// 	- **-1**: If you set the value to -1, the bandwidth of the listener is unlimited.
	//
	// example:
	//
	// -1
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The ID of the CA certificate.
	//
	// 	- If both the CA certificate and the server certificate are uploaded, mutual authentication is used.
	//
	// 	- If you upload only the server certificate, one-way authentication is used.
	//
	// example:
	//
	// 139a00604ad-cn-east-****
	CACertificateId *string `json:"CACertificateId,omitempty" xml:"CACertificateId,omitempty"`
	// The cookie that you want to configure for the server.
	//
	// The cookie must be 1 to 200 characters in length, and can contain only ASCII letters and digits. It cannot contain commas (,), semicolons (;), or space characters. It cannot start with a dollar sign ($).
	//
	// >  This parameter is required when you set the **StickySession*	- parameter to **on*	- and the **StickySessionType*	- parameter to **server**.
	//
	// example:
	//
	// testCookie
	Cookie *string `json:"Cookie,omitempty" xml:"Cookie,omitempty"`
	// The timeout period of the cookie. Unit: seconds.
	//
	// Valid values: **1*	- to **86400**.
	//
	// >  This parameter is required if the **StickySession*	- parameter is set to **on*	- and the **StickySessionType*	- parameter is set to **insert**.
	//
	// example:
	//
	// 500
	CookieTimeout *int32 `json:"CookieTimeout,omitempty" xml:"CookieTimeout,omitempty"`
	// The description of the listener.
	//
	// The name must be 1 to 256 characters in length and can contain letters, digits, hyphens (-), forward slashes (/), periods (.), and underscores (_).
	//
	// example:
	//
	// https_80
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to use `HTTP 2.0`. Valid values:
	//
	// 	- **on**: yes
	//
	// 	- **off**: no
	//
	// example:
	//
	// off
	EnableHttp2 *string `json:"EnableHttp2,omitempty" xml:"EnableHttp2,omitempty"`
	// Specifies whether to enable `Gzip` compression to compress specific types of files. Valid values:
	//
	// 	- **on**: yes
	//
	// 	- **off**: no
	//
	// example:
	//
	// on
	Gzip *string `json:"Gzip,omitempty" xml:"Gzip,omitempty"`
	// Specifies whether to enable health checks. Valid values:
	//
	// 	- **on**: yes
	//
	// 	- **off**: no
	//
	// example:
	//
	// on
	HealthCheck *string `json:"HealthCheck,omitempty" xml:"HealthCheck,omitempty"`
	// The port that is used for health checks.
	//
	// Valid values: **1*	- to **65535**.
	//
	// >  This parameter takes effect only when the **HealthCheck*	- parameter is set to **on**.
	//
	// example:
	//
	// 8080
	HealthCheckConnectPort *int32 `json:"HealthCheckConnectPort,omitempty" xml:"HealthCheckConnectPort,omitempty"`
	// The domain name that is used for health checks. Valid values:
	//
	// 	- **$_ip**: the private IP address of a backend server. If you do not set this parameter or set the parameter to $_ip, the CLB instance uses the private IP address of each backend server as the domain name for health checks.
	//
	// 	- **domain**: The domain name must be 1 to 80 characters in length, and can contain letters, digits, periods (.),and hyphens (-).
	//
	// >  This parameter takes effect only when the **HealthCheck*	- parameter is set to **on**.
	//
	// example:
	//
	// 172.XX.XX.16
	HealthCheckDomain *string `json:"HealthCheckDomain,omitempty" xml:"HealthCheckDomain,omitempty"`
	// The HTTP status code of a successful health check. Separate multiple HTTP status codes with commas (,).
	//
	// Valid values: **http_2xx**, **http_3xx**, **http_4xx**, and **http_5xx**.
	//
	// >  This parameter takes effect only when the **HealthCheck*	- parameter is set to **on**.
	//
	// example:
	//
	// http_2xx,http_3xx
	HealthCheckHttpCode *string `json:"HealthCheckHttpCode,omitempty" xml:"HealthCheckHttpCode,omitempty"`
	// The interval between two consecutive health checks. Unit: seconds.
	//
	// Valid values: **1*	- to **50**.
	//
	// >  This parameter takes effect only when the **HealthCheck*	- parameter is set to **on**.
	//
	// example:
	//
	// 5
	HealthCheckInterval *int32 `json:"HealthCheckInterval,omitempty" xml:"HealthCheckInterval,omitempty"`
	// The HTTP method that is used for health checks. Valid values: **head*	- and **get**.
	//
	// >  This parameter takes effect only when the **HealthCheck*	- parameter is set to **on**.
	//
	// example:
	//
	// get
	HealthCheckMethod *string `json:"HealthCheckMethod,omitempty" xml:"HealthCheckMethod,omitempty"`
	// The timeout period of a health check response. If a backend ECS instance does not respond within the specified timeout period, the ECS instance fails the health check. Unit: seconds Valid values: **1*	- to **300**.
	//
	// >  This parameter takes effect only if the **HealthCheck*	- parameter is set to **on**.
	//
	// example:
	//
	// 3
	HealthCheckTimeout *int32 `json:"HealthCheckTimeout,omitempty" xml:"HealthCheckTimeout,omitempty"`
	// The URL that is used for health checks.
	//
	// The URL must be 1 to 80 characters in length and can contain letters, digits, and the following characters: - / . % ? # &. The URL must not be a single forward slash (/) but it must start with a forward slash (/).
	//
	// >  This parameter takes effect only when the **HealthCheck*	- parameter is set to **on**.
	//
	// example:
	//
	// /test/index.html
	HealthCheckURI *string `json:"HealthCheckURI,omitempty" xml:"HealthCheckURI,omitempty"`
	// The number of health checks that an unhealthy backend server must consecutively pass before it can be declared healthy (from **fail*	- to **success**).
	//
	// Valid values: **2*	- to **10**.
	//
	// >  This parameter takes effect only when the **HealthCheck*	- parameter is set to **on**.
	//
	// example:
	//
	// 4
	HealthyThreshold *int32 `json:"HealthyThreshold,omitempty" xml:"HealthyThreshold,omitempty"`
	// The timeout period of an idle connection. Unit: seconds. Valid values: **1 to 60**. Default value: **15**.
	//
	// If no request is received within the specified timeout period, CLB closes the connection. When another request is received, CLB establishes a new connection.
	//
	// example:
	//
	// 23
	IdleTimeout *int32 `json:"IdleTimeout,omitempty" xml:"IdleTimeout,omitempty"`
	// The frontend port that is used by the CLB instance.
	//
	// Valid values: **1*	- to **65535**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 80
	ListenerPort *int32 `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	// The ID of the CLB instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// lb-sjhfdji****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	OwnerAccount   *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the CLB instance is deployed.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The timeout period of a request. Unit: seconds. Valid values: **1 to 180**. Default value: **60**.
	//
	// If no response is received from the backend server during the request timeout period, CLB sends an HTTP 504 error code to the client.
	//
	// example:
	//
	// 223
	RequestTimeout       *int32  `json:"RequestTimeout,omitempty" xml:"RequestTimeout,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The scheduling algorithm. Valid values:
	//
	// 	- **wrr**: Backend servers that have higher weights receive more requests than backend servers that have lower weights.
	//
	// 	- **rr**: Requests are distributed to backend servers in sequence.
	//
	// example:
	//
	// wrr
	Scheduler *string `json:"Scheduler,omitempty" xml:"Scheduler,omitempty"`
	// The ID of the server certificate.
	//
	// example:
	//
	// idkp-123-cn-te****
	ServerCertificateId *string `json:"ServerCertificateId,omitempty" xml:"ServerCertificateId,omitempty"`
	// Specifies whether to enable session persistence. Valid values:
	//
	// 	- **on**: yes
	//
	// 	- **off**: no
	//
	// example:
	//
	// on
	StickySession *string `json:"StickySession,omitempty" xml:"StickySession,omitempty"`
	// The method that is used to handle a cookie. Valid values:
	//
	// 	- **insert**: inserts a cookie.
	//
	//     CLB inserts a cookie (SERVERID) into the first HTTP or HTTPS response that is sent to a client. The next request from the client will contain this cookie, and the listener will distribute this request to the recorded backend server.
	//
	// 	- **server**: rewrites a cookie.
	//
	//     When CLB detects a user-defined cookie, it overwrites the original cookie with the user-defined cookie. The next request from the client will contain the user-defined cookie, and the listener will distribute this request to the recorded backend server.
	//
	// >  This parameter is required if the **StickySession*	- parameter is set to **on**.
	//
	// example:
	//
	// insert
	StickySessionType *string `json:"StickySessionType,omitempty" xml:"StickySessionType,omitempty"`
	// The Transport Layer Security (TLS) security policy. Each security policy contains TLS protocol versions and cipher suites available for HTTPS.
	//
	// 	- **tls_cipher_policy_1_0**:
	//
	//     Supported TLS versions: TLS 1.0, TLS 1.1, and TLS 1.2
	//
	//     Supported cipher suites: ECDHE-RSA-AES128-GCM-SHA256, ECDHE-RSA-AES256-GCM-SHA384, ECDHE-RSA-AES128-SHA256, ECDHE-RSA-AES256-SHA384, AES128-GCM-SHA256, AES256-GCM-SHA384, AES128-SHA256, AES256-SHA256, ECDHE-RSA-AES128-SHA, ECDHE-RSA-AES256-SHA, AES128-SHA, AES256-SHA, and DES-CBC3-SHA
	//
	// 	- **tls_cipher_policy_1_1**:
	//
	//     Supported TLS versions: TLS 1.1 and TLS 1.2
	//
	//     Supported cipher suites: ECDHE-RSA-AES128-GCM-SHA256, ECDHE-RSA-AES256-GCM-SHA384, ECDHE-RSA-AES128-SHA256, ECDHE-RSA-AES256-SHA384, AES128-GCM-SHA256, AES256-GCM-SHA384, AES128-SHA256, AES256-SHA256, ECDHE-RSA-AES128-SHA, ECDHE-RSA-AES256-SHA, AES128-SHA, AES256-SHA, and DES-CBC3-SHA
	//
	// 	- **tls_cipher_policy_1_2**
	//
	//     Supported TLS version: TLS 1.2
	//
	//     Supported cipher suites: ECDHE-RSA-AES128-GCM-SHA256, ECDHE-RSA-AES256-GCM-SHA384, ECDHE-RSA-AES128-SHA256, ECDHE-RSA-AES256-SHA384, AES128-GCM-SHA256, AES256-GCM-SHA384, AES128-SHA256, AES256-SHA256, ECDHE-RSA-AES128-SHA, ECDHE-RSA-AES256-SHA, AES128-SHA, AES256-SHA, and DES-CBC3-SHA
	//
	// 	- **tls_cipher_policy_1_2_strict**
	//
	//     Supported TLS version: TLS 1.2
	//
	//     Supported cipher suites: ECDHE-RSA-AES128-GCM-SHA256, ECDHE-RSA-AES256-GCM-SHA384, ECDHE-RSA-AES128-SHA256, ECDHE-RSA-AES256-SHA384, ECDHE-RSA-AES128-SHA, and ECDHE-RSA-AES256-SHA
	//
	// 	- **tls_cipher_policy_1_2_strict_with_1_3**
	//
	//     Supported TLS versions: TLS 1.2 and TLS 1.3
	//
	//     Supported cipher suites: TLS_AES_128_GCM_SHA256, TLS_AES_256_GCM_SHA384, TLS_CHACHA20_POLY1305_SHA256, TLS_AES_128_CCM_SHA256, TLS_AES_128_CCM_8_SHA256, ECDHE-ECDSA-AES128-GCM-SHA256, ECDHE-ECDSA-AES256-GCM-SHA384, ECDHE-ECDSA-AES128-SHA256, ECDHE-ECDSA-AES256-SHA384, ECDHE-RSA-AES128-GCM-SHA256, ECDHE-RSA-AES256-GCM-SHA384, ECDHE-RSA-AES128-SHA256, ECDHE-RSA-AES256-SHA384, ECDHE-ECDSA-AES128-SHA, ECDHE-ECDSA-AES256-SHA, ECDHE-RSA-AES128-SHA, and ECDHE-RSA-AES256-SHA
	//
	// example:
	//
	// tls_cipher_policy_1_2
	TLSCipherPolicy *string `json:"TLSCipherPolicy,omitempty" xml:"TLSCipherPolicy,omitempty"`
	// The number of health checks that a healthy backend server must consecutively fail before it can be declared unhealthy (from **success*	- to **fail**).
	//
	// Valid values: **2*	- to **10**.
	//
	// >  This parameter takes effect only when the **HealthCheck*	- parameter is set to **on**.
	//
	// example:
	//
	// 4
	UnhealthyThreshold *int32 `json:"UnhealthyThreshold,omitempty" xml:"UnhealthyThreshold,omitempty"`
	// Specifies whether to use a vServer group. Valid values:
	//
	// 	- **on**: yes
	//
	// 	- **off**: no
	//
	// example:
	//
	// on
	VServerGroup *string `json:"VServerGroup,omitempty" xml:"VServerGroup,omitempty"`
	// The ID of the vServer group.
	//
	// example:
	//
	// rsp-cige6j****
	VServerGroupId *string `json:"VServerGroupId,omitempty" xml:"VServerGroupId,omitempty"`
	// Specifies whether to use the `X-Forwarded-For` header to retrieve client IP addresses. Valid values:
	//
	// 	- **on**: yes
	//
	// 	- **off**: no
	//
	// example:
	//
	// on
	XForwardedFor *string `json:"XForwardedFor,omitempty" xml:"XForwardedFor,omitempty"`
	// Specifies whether to use the `XForwardedFor_ClientSrcPort` header to retrieve the client port. Valid values:
	//
	// 	- **on**
	//
	// 	- **off**
	//
	// example:
	//
	// off
	XForwardedFor_ClientSrcPort *string `json:"XForwardedFor_ClientSrcPort,omitempty" xml:"XForwardedFor_ClientSrcPort,omitempty"`
	// Specifies whether to use the `SLB-ID` header to retrieve the ID of the CLB instance. Valid values:
	//
	// 	- **on**: yes
	//
	// 	- **off**: no
	//
	// example:
	//
	// on
	XForwardedFor_SLBID *string `json:"XForwardedFor_SLBID,omitempty" xml:"XForwardedFor_SLBID,omitempty"`
	// Specifies whether to use the `SLB-IP` header to obtain the virtual IP address (VIP) requested by the client. Valid values:
	//
	// 	- **on**: yes
	//
	// 	- **off**: no
	//
	// example:
	//
	// on
	XForwardedFor_SLBIP *string `json:"XForwardedFor_SLBIP,omitempty" xml:"XForwardedFor_SLBIP,omitempty"`
	// Specifies whether to use the `XForwardedFor_SLBPORT` header to retrieve the listener port of the CLB instance. Valid values:
	//
	// 	- **on**
	//
	// 	- **off**
	//
	// example:
	//
	// off
	XForwardedFor_SLBPORT *string `json:"XForwardedFor_SLBPORT,omitempty" xml:"XForwardedFor_SLBPORT,omitempty"`
	// Specifies whether to use the `X-Forwarded-Proto` header to retrieve the listener protocol. Valid values:
	//
	// 	- **on**: yes
	//
	// 	- **off**: no
	//
	// example:
	//
	// on
	XForwardedFor_proto *string `json:"XForwardedFor_proto,omitempty" xml:"XForwardedFor_proto,omitempty"`
}

func (s SetLoadBalancerHTTPSListenerAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s SetLoadBalancerHTTPSListenerAttributeRequest) GoString() string {
	return s.String()
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) GetAclId() *string {
	return s.AclId
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) GetAclStatus() *string {
	return s.AclStatus
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) GetAclType() *string {
	return s.AclType
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) GetCACertificateId() *string {
	return s.CACertificateId
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) GetCookie() *string {
	return s.Cookie
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) GetCookieTimeout() *int32 {
	return s.CookieTimeout
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) GetDescription() *string {
	return s.Description
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) GetEnableHttp2() *string {
	return s.EnableHttp2
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) GetGzip() *string {
	return s.Gzip
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) GetHealthCheck() *string {
	return s.HealthCheck
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) GetHealthCheckConnectPort() *int32 {
	return s.HealthCheckConnectPort
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) GetHealthCheckDomain() *string {
	return s.HealthCheckDomain
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) GetHealthCheckHttpCode() *string {
	return s.HealthCheckHttpCode
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) GetHealthCheckInterval() *int32 {
	return s.HealthCheckInterval
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) GetHealthCheckMethod() *string {
	return s.HealthCheckMethod
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) GetHealthCheckTimeout() *int32 {
	return s.HealthCheckTimeout
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) GetHealthCheckURI() *string {
	return s.HealthCheckURI
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) GetHealthyThreshold() *int32 {
	return s.HealthyThreshold
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) GetIdleTimeout() *int32 {
	return s.IdleTimeout
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) GetListenerPort() *int32 {
	return s.ListenerPort
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) GetRequestTimeout() *int32 {
	return s.RequestTimeout
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) GetScheduler() *string {
	return s.Scheduler
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) GetServerCertificateId() *string {
	return s.ServerCertificateId
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) GetStickySession() *string {
	return s.StickySession
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) GetStickySessionType() *string {
	return s.StickySessionType
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) GetTLSCipherPolicy() *string {
	return s.TLSCipherPolicy
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) GetUnhealthyThreshold() *int32 {
	return s.UnhealthyThreshold
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) GetVServerGroup() *string {
	return s.VServerGroup
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) GetVServerGroupId() *string {
	return s.VServerGroupId
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) GetXForwardedFor() *string {
	return s.XForwardedFor
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) GetXForwardedFor_ClientSrcPort() *string {
	return s.XForwardedFor_ClientSrcPort
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) GetXForwardedFor_SLBID() *string {
	return s.XForwardedFor_SLBID
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) GetXForwardedFor_SLBIP() *string {
	return s.XForwardedFor_SLBIP
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) GetXForwardedFor_SLBPORT() *string {
	return s.XForwardedFor_SLBPORT
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) GetXForwardedFor_proto() *string {
	return s.XForwardedFor_proto
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) SetAclId(v string) *SetLoadBalancerHTTPSListenerAttributeRequest {
	s.AclId = &v
	return s
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) SetAclStatus(v string) *SetLoadBalancerHTTPSListenerAttributeRequest {
	s.AclStatus = &v
	return s
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) SetAclType(v string) *SetLoadBalancerHTTPSListenerAttributeRequest {
	s.AclType = &v
	return s
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) SetBandwidth(v int32) *SetLoadBalancerHTTPSListenerAttributeRequest {
	s.Bandwidth = &v
	return s
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) SetCACertificateId(v string) *SetLoadBalancerHTTPSListenerAttributeRequest {
	s.CACertificateId = &v
	return s
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) SetCookie(v string) *SetLoadBalancerHTTPSListenerAttributeRequest {
	s.Cookie = &v
	return s
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) SetCookieTimeout(v int32) *SetLoadBalancerHTTPSListenerAttributeRequest {
	s.CookieTimeout = &v
	return s
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) SetDescription(v string) *SetLoadBalancerHTTPSListenerAttributeRequest {
	s.Description = &v
	return s
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) SetEnableHttp2(v string) *SetLoadBalancerHTTPSListenerAttributeRequest {
	s.EnableHttp2 = &v
	return s
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) SetGzip(v string) *SetLoadBalancerHTTPSListenerAttributeRequest {
	s.Gzip = &v
	return s
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) SetHealthCheck(v string) *SetLoadBalancerHTTPSListenerAttributeRequest {
	s.HealthCheck = &v
	return s
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) SetHealthCheckConnectPort(v int32) *SetLoadBalancerHTTPSListenerAttributeRequest {
	s.HealthCheckConnectPort = &v
	return s
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) SetHealthCheckDomain(v string) *SetLoadBalancerHTTPSListenerAttributeRequest {
	s.HealthCheckDomain = &v
	return s
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) SetHealthCheckHttpCode(v string) *SetLoadBalancerHTTPSListenerAttributeRequest {
	s.HealthCheckHttpCode = &v
	return s
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) SetHealthCheckInterval(v int32) *SetLoadBalancerHTTPSListenerAttributeRequest {
	s.HealthCheckInterval = &v
	return s
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) SetHealthCheckMethod(v string) *SetLoadBalancerHTTPSListenerAttributeRequest {
	s.HealthCheckMethod = &v
	return s
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) SetHealthCheckTimeout(v int32) *SetLoadBalancerHTTPSListenerAttributeRequest {
	s.HealthCheckTimeout = &v
	return s
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) SetHealthCheckURI(v string) *SetLoadBalancerHTTPSListenerAttributeRequest {
	s.HealthCheckURI = &v
	return s
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) SetHealthyThreshold(v int32) *SetLoadBalancerHTTPSListenerAttributeRequest {
	s.HealthyThreshold = &v
	return s
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) SetIdleTimeout(v int32) *SetLoadBalancerHTTPSListenerAttributeRequest {
	s.IdleTimeout = &v
	return s
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) SetListenerPort(v int32) *SetLoadBalancerHTTPSListenerAttributeRequest {
	s.ListenerPort = &v
	return s
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) SetLoadBalancerId(v string) *SetLoadBalancerHTTPSListenerAttributeRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) SetOwnerAccount(v string) *SetLoadBalancerHTTPSListenerAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) SetOwnerId(v int64) *SetLoadBalancerHTTPSListenerAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) SetRegionId(v string) *SetLoadBalancerHTTPSListenerAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) SetRequestTimeout(v int32) *SetLoadBalancerHTTPSListenerAttributeRequest {
	s.RequestTimeout = &v
	return s
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) SetResourceOwnerAccount(v string) *SetLoadBalancerHTTPSListenerAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) SetResourceOwnerId(v int64) *SetLoadBalancerHTTPSListenerAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) SetScheduler(v string) *SetLoadBalancerHTTPSListenerAttributeRequest {
	s.Scheduler = &v
	return s
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) SetServerCertificateId(v string) *SetLoadBalancerHTTPSListenerAttributeRequest {
	s.ServerCertificateId = &v
	return s
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) SetStickySession(v string) *SetLoadBalancerHTTPSListenerAttributeRequest {
	s.StickySession = &v
	return s
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) SetStickySessionType(v string) *SetLoadBalancerHTTPSListenerAttributeRequest {
	s.StickySessionType = &v
	return s
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) SetTLSCipherPolicy(v string) *SetLoadBalancerHTTPSListenerAttributeRequest {
	s.TLSCipherPolicy = &v
	return s
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) SetUnhealthyThreshold(v int32) *SetLoadBalancerHTTPSListenerAttributeRequest {
	s.UnhealthyThreshold = &v
	return s
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) SetVServerGroup(v string) *SetLoadBalancerHTTPSListenerAttributeRequest {
	s.VServerGroup = &v
	return s
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) SetVServerGroupId(v string) *SetLoadBalancerHTTPSListenerAttributeRequest {
	s.VServerGroupId = &v
	return s
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) SetXForwardedFor(v string) *SetLoadBalancerHTTPSListenerAttributeRequest {
	s.XForwardedFor = &v
	return s
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) SetXForwardedFor_ClientSrcPort(v string) *SetLoadBalancerHTTPSListenerAttributeRequest {
	s.XForwardedFor_ClientSrcPort = &v
	return s
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) SetXForwardedFor_SLBID(v string) *SetLoadBalancerHTTPSListenerAttributeRequest {
	s.XForwardedFor_SLBID = &v
	return s
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) SetXForwardedFor_SLBIP(v string) *SetLoadBalancerHTTPSListenerAttributeRequest {
	s.XForwardedFor_SLBIP = &v
	return s
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) SetXForwardedFor_SLBPORT(v string) *SetLoadBalancerHTTPSListenerAttributeRequest {
	s.XForwardedFor_SLBPORT = &v
	return s
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) SetXForwardedFor_proto(v string) *SetLoadBalancerHTTPSListenerAttributeRequest {
	s.XForwardedFor_proto = &v
	return s
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) Validate() error {
	return dara.Validate(s)
}
