// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLoadBalancerHTTPSListenerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAclId(v string) *CreateLoadBalancerHTTPSListenerRequest
	GetAclId() *string
	SetAclStatus(v string) *CreateLoadBalancerHTTPSListenerRequest
	GetAclStatus() *string
	SetAclType(v string) *CreateLoadBalancerHTTPSListenerRequest
	GetAclType() *string
	SetBackendServerPort(v int32) *CreateLoadBalancerHTTPSListenerRequest
	GetBackendServerPort() *int32
	SetBandwidth(v int32) *CreateLoadBalancerHTTPSListenerRequest
	GetBandwidth() *int32
	SetCACertificateId(v string) *CreateLoadBalancerHTTPSListenerRequest
	GetCACertificateId() *string
	SetCookie(v string) *CreateLoadBalancerHTTPSListenerRequest
	GetCookie() *string
	SetCookieTimeout(v int32) *CreateLoadBalancerHTTPSListenerRequest
	GetCookieTimeout() *int32
	SetDescription(v string) *CreateLoadBalancerHTTPSListenerRequest
	GetDescription() *string
	SetEnableHttp2(v string) *CreateLoadBalancerHTTPSListenerRequest
	GetEnableHttp2() *string
	SetGzip(v string) *CreateLoadBalancerHTTPSListenerRequest
	GetGzip() *string
	SetHealthCheck(v string) *CreateLoadBalancerHTTPSListenerRequest
	GetHealthCheck() *string
	SetHealthCheckConnectPort(v int32) *CreateLoadBalancerHTTPSListenerRequest
	GetHealthCheckConnectPort() *int32
	SetHealthCheckDomain(v string) *CreateLoadBalancerHTTPSListenerRequest
	GetHealthCheckDomain() *string
	SetHealthCheckHttpCode(v string) *CreateLoadBalancerHTTPSListenerRequest
	GetHealthCheckHttpCode() *string
	SetHealthCheckInterval(v int32) *CreateLoadBalancerHTTPSListenerRequest
	GetHealthCheckInterval() *int32
	SetHealthCheckMethod(v string) *CreateLoadBalancerHTTPSListenerRequest
	GetHealthCheckMethod() *string
	SetHealthCheckTimeout(v int32) *CreateLoadBalancerHTTPSListenerRequest
	GetHealthCheckTimeout() *int32
	SetHealthCheckURI(v string) *CreateLoadBalancerHTTPSListenerRequest
	GetHealthCheckURI() *string
	SetHealthyThreshold(v int32) *CreateLoadBalancerHTTPSListenerRequest
	GetHealthyThreshold() *int32
	SetIdleTimeout(v int32) *CreateLoadBalancerHTTPSListenerRequest
	GetIdleTimeout() *int32
	SetListenerPort(v int32) *CreateLoadBalancerHTTPSListenerRequest
	GetListenerPort() *int32
	SetLoadBalancerId(v string) *CreateLoadBalancerHTTPSListenerRequest
	GetLoadBalancerId() *string
	SetOwnerAccount(v string) *CreateLoadBalancerHTTPSListenerRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateLoadBalancerHTTPSListenerRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CreateLoadBalancerHTTPSListenerRequest
	GetRegionId() *string
	SetRequestTimeout(v int32) *CreateLoadBalancerHTTPSListenerRequest
	GetRequestTimeout() *int32
	SetResourceOwnerAccount(v string) *CreateLoadBalancerHTTPSListenerRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateLoadBalancerHTTPSListenerRequest
	GetResourceOwnerId() *int64
	SetScheduler(v string) *CreateLoadBalancerHTTPSListenerRequest
	GetScheduler() *string
	SetServerCertificateId(v string) *CreateLoadBalancerHTTPSListenerRequest
	GetServerCertificateId() *string
	SetStickySession(v string) *CreateLoadBalancerHTTPSListenerRequest
	GetStickySession() *string
	SetStickySessionType(v string) *CreateLoadBalancerHTTPSListenerRequest
	GetStickySessionType() *string
	SetTLSCipherPolicy(v string) *CreateLoadBalancerHTTPSListenerRequest
	GetTLSCipherPolicy() *string
	SetTag(v []*CreateLoadBalancerHTTPSListenerRequestTag) *CreateLoadBalancerHTTPSListenerRequest
	GetTag() []*CreateLoadBalancerHTTPSListenerRequestTag
	SetUnhealthyThreshold(v int32) *CreateLoadBalancerHTTPSListenerRequest
	GetUnhealthyThreshold() *int32
	SetVServerGroupId(v string) *CreateLoadBalancerHTTPSListenerRequest
	GetVServerGroupId() *string
	SetXForwardedFor(v string) *CreateLoadBalancerHTTPSListenerRequest
	GetXForwardedFor() *string
	SetXForwardedFor_ClientSrcPort(v string) *CreateLoadBalancerHTTPSListenerRequest
	GetXForwardedFor_ClientSrcPort() *string
	SetXForwardedFor_SLBID(v string) *CreateLoadBalancerHTTPSListenerRequest
	GetXForwardedFor_SLBID() *string
	SetXForwardedFor_SLBIP(v string) *CreateLoadBalancerHTTPSListenerRequest
	GetXForwardedFor_SLBIP() *string
	SetXForwardedFor_SLBPORT(v string) *CreateLoadBalancerHTTPSListenerRequest
	GetXForwardedFor_SLBPORT() *string
	SetXForwardedFor_proto(v string) *CreateLoadBalancerHTTPSListenerRequest
	GetXForwardedFor_proto() *string
}

type CreateLoadBalancerHTTPSListenerRequest struct {
	// The ID of the network access control list (ACL) that is associated with the listener.
	//
	// >  This parameter is required if **AclStatus*	- is set to **on**.
	//
	// example:
	//
	// nacl-a2do9e413e0spzasx****
	AclId *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
	// Specifies whether to enable access control. Valid values:
	//
	// 	- **on**: yes
	//
	// 	- **off**: no
	//
	// example:
	//
	// off
	AclStatus *string `json:"AclStatus,omitempty" xml:"AclStatus,omitempty"`
	// The type of the network ACL. Valid values:
	//
	// 	- **white**: a whitelist. Only requests from the IP addresses or CIDR blocks in the network ACL are forwarded. Whitelists apply to scenarios in which you want to allow only specific IP addresses to access an application. Your service may be adversely affected if the allowlist is not properly configured. After a whitelist is configured, only requests from IP addresses that are added to the whitelist are forwarded by the listener.
	//
	//     If you enable a whitelist but do not add an IP address to the ACL, the listener forwards all requests.
	//
	// 	- **black**: a blacklist. All requests from the IP addresses or CIDR blocks in the network ACL are denied. The blacklist applies to scenarios in which you want to deny access from specific IP addresses to an application.
	//
	//     If a blacklist is configured for a listener but no IP address is added to the blacklist, the listener forwards all requests.
	//
	// >  If **AclStatus*	- is set to **on**, this parameter is required.
	//
	// example:
	//
	// white
	AclType *string `json:"AclType,omitempty" xml:"AclType,omitempty"`
	// The backend port that is used by the CLB instance. Valid values: **1*	- to **65535**.
	//
	// If the VServerGroupId parameter is not set, this parameter is required.
	//
	// example:
	//
	// 80
	BackendServerPort *int32 `json:"BackendServerPort,omitempty" xml:"BackendServerPort,omitempty"`
	// The maximum bandwidth of the listener. Unit: Mbit/s.
	//
	// The URL must meet the following requirements:
	//
	// 	- **-1**: For a pay-by-data-transfer Internet-facing CLB instance, this parameter is set to -1. This indicates that the bandwidth of the listener is unlimited.
	//
	// This parameter is required.
	//
	// example:
	//
	// -1
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The ID of the certification authority (CA) certificate.
	//
	// If both the CA certificate and the server certificate are uploaded, mutual authentication is used.
	//
	// If you upload only the server certificate, one-way authentication is used.
	//
	// example:
	//
	// 139a00604ad-cn-east-hangzh****
	CACertificateId *string `json:"CACertificateId,omitempty" xml:"CACertificateId,omitempty"`
	// The cookie that you configure for the server.
	//
	// The cookie must be 1 to 200 characters in length, and can contain only ASCII letters and digits. It cannot contain commas (,), semicolons (;), spaces, or start with a dollar sign ($).
	//
	// >  This parameter is required when the **StickySession*	- parameter is set to **on*	- and the **StickySessionType*	- parameter is set to **server**.
	//
	// example:
	//
	// B490B5EBF6F3CD402E515D22BCDA****
	Cookie *string `json:"Cookie,omitempty" xml:"Cookie,omitempty"`
	// The timeout period of a cookie. Unit: seconds.
	//
	// Valid values: **1*	- to **86400**.
	//
	// >  If **StickySession*	- is set to **on*	- and **StickySessionType*	- is set to **insert**, this parameter is required.
	//
	// example:
	//
	// 500
	CookieTimeout *int32 `json:"CookieTimeout,omitempty" xml:"CookieTimeout,omitempty"`
	// The name of the listener.
	//
	// The name must be 1 to 256 characters in length and can contain letters, digits, hyphens (-), forward slashes (/), periods (.), and underscores (_).
	//
	// example:
	//
	// HTTPS_443
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to enable HTTP/2. Valid values:
	//
	// 	- **on**: yes
	//
	// 	- **off**: no
	//
	// example:
	//
	// off
	EnableHttp2 *string `json:"EnableHttp2,omitempty" xml:"EnableHttp2,omitempty"`
	// Specifies whether to enable `GZIP` compression to compress specific types of files. Valid values:
	//
	// 	- **on**
	//
	// 	- **off**
	//
	// example:
	//
	// on
	Gzip *string `json:"Gzip,omitempty" xml:"Gzip,omitempty"`
	// Specifies whether to enable the health check feature. Valid values:
	//
	// 	- **on**: yes
	//
	// 	- **off**: no
	//
	// This parameter is required.
	//
	// example:
	//
	// on
	HealthCheck *string `json:"HealthCheck,omitempty" xml:"HealthCheck,omitempty"`
	// The port that is used for health checks.
	//
	// Valid values: **1*	- to **65535**.
	//
	// >  This parameter takes effect only if the **HealthCheck*	- parameter is set to **on**.
	//
	// example:
	//
	// 8080
	HealthCheckConnectPort *int32 `json:"HealthCheckConnectPort,omitempty" xml:"HealthCheckConnectPort,omitempty"`
	// The domain name that is used for health checks. Valid values:
	//
	// 	- **$_ip**: the private IP address of a backend server. If you do not set the HealthCheckDomain parameter or set the parameter to $_ip, the CLB instance uses the private IP address of each backend server for health checks.
	//
	// 	- **domain**: The domain name must be 1 to 80 characters in length and can contain letters, digits, periods (.), and hyphens (-).
	//
	// >  This parameter takes effect only if the **HealthCheck*	- parameter is set to **on**.
	//
	// example:
	//
	// 172.XX.XX.16
	HealthCheckDomain *string `json:"HealthCheckDomain,omitempty" xml:"HealthCheckDomain,omitempty"`
	// The HTTP status code for a successful health check. Separate multiple HTTP status codes with commas (,).
	//
	// Valid values: **http_2xx**, **http_3xx**, **http_4xx**, and **http_5xx**.
	//
	// >  This parameter takes effect only if the **HealthCheck*	- parameter is set to **on**.
	//
	// example:
	//
	// http_2xx,http_3xx
	HealthCheckHttpCode *string `json:"HealthCheckHttpCode,omitempty" xml:"HealthCheckHttpCode,omitempty"`
	// The interval between two consecutive health checks. Unit: seconds.
	//
	// Valid values: **1*	- to **50**.
	//
	// >  This parameter takes effect only if the **HealthCheck*	- parameter is set to **on**.
	//
	// example:
	//
	// 5
	HealthCheckInterval *int32 `json:"HealthCheckInterval,omitempty" xml:"HealthCheckInterval,omitempty"`
	// The health check method used in HTTP health checks. Valid values: **head*	- and **get**.
	//
	// >  This parameter takes effect only if the **HealthCheck*	- parameter is set to **on**.
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
	// The URI that is used for health checks.
	//
	// The URI must be 1 to 80 characters in length, and can contain letters, digits, and the following special characters: `-/.%?#&`. The URI must start with a forward slash (`/`), but cannot be a single forward slash (`/`).
	//
	// >  This parameter takes effect only if the **HealthCheck*	- parameter is set to **on**.
	//
	// example:
	//
	// /test/index.html
	HealthCheckURI *string `json:"HealthCheckURI,omitempty" xml:"HealthCheckURI,omitempty"`
	// The number of times that an unhealthy backend server must consecutively pass health checks before it is declared healthy. In this case, the health status is changed from **fail*	- to **success**.
	//
	// Valid values: **2*	- to **10**.
	//
	// >  This parameter takes effect only if the **HealthCheck*	- parameter is set to **on**.
	//
	// example:
	//
	// 4
	HealthyThreshold *int32 `json:"HealthyThreshold,omitempty" xml:"HealthyThreshold,omitempty"`
	// The timeout period of an idle connection. Valid values: **1 to 60**. Default value: **15**. Unit: seconds.
	//
	// If no request is received within the specified timeout period, CLB closes the connection. When a request is received, CLB establishes a new connection.
	//
	// example:
	//
	// 12
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
	// lb-bp1o94dp5i6earr****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	OwnerAccount   *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the CLB instance.
	//
	// You can query the region ID from the [Regions and zones](https://help.aliyun.com/document_detail/40654.html) list or by calling the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The timeout period of a request. Valid values: **1 to 180**. Default value: **60**. Unit: seconds.
	//
	// If no response is received from a backend server within the specified timeout period, CLB returns the HTTP 504 status code to the client.
	//
	// example:
	//
	// 23
	RequestTimeout       *int32  `json:"RequestTimeout,omitempty" xml:"RequestTimeout,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The scheduling algorithm. Valid values:
	//
	// 	- **wrr**: Backend servers with higher weights receive more requests than those with lower weights.
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
	// idkp-123-cn-test-****
	ServerCertificateId *string `json:"ServerCertificateId,omitempty" xml:"ServerCertificateId,omitempty"`
	// Specifies whether to enable session persistence. Valid values:
	//
	// 	- **on**
	//
	// 	- **off**
	//
	// example:
	//
	// on
	StickySession *string `json:"StickySession,omitempty" xml:"StickySession,omitempty"`
	// The method that is used to handle a cookie. Valid values: **insert*	- and **server**.
	//
	// 	- **insert**: inserts a cookie.
	//
	//     CLB inserts a cookie (SERVERID) into the first HTTP or HTTPS response packet that is sent to a client. The next request from the client will contain this cookie, and the listener will distribute this request to the recorded backend server.
	//
	// 	- **server**: rewrites a cookie.
	//
	//     When CLB detects a user-defined cookie, it overwrites the original cookie with the user-defined cookie. The next request from the client carries the user-defined cookie, and the listener will distribute the request to the recorded backend server.
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
	// tls_cipher_policy_1_1
	TLSCipherPolicy *string `json:"TLSCipherPolicy,omitempty" xml:"TLSCipherPolicy,omitempty"`
	// The tags.
	Tag []*CreateLoadBalancerHTTPSListenerRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The number of times that a healthy backend server must consecutively fail health checks before it is declared unhealthy. In this case, the health status is changed from **success*	- to **fail**.
	//
	// Valid values: **2*	- to **10**.
	//
	// >  This parameter takes effect only if the **HealthCheck*	- parameter is set to **on**.
	//
	// example:
	//
	// 4
	UnhealthyThreshold *int32 `json:"UnhealthyThreshold,omitempty" xml:"UnhealthyThreshold,omitempty"`
	// The ID of the server group.
	//
	// example:
	//
	// rsp-cige6j5e7p****
	VServerGroupId *string `json:"VServerGroupId,omitempty" xml:"VServerGroupId,omitempty"`
	// Specifies whether to use the `X-Forwarded-For` header to retrieve client IP addresses. Valid values:
	//
	// 	- **on**
	//
	// 	- **off**
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
	// 	- **on**
	//
	// 	- **off**
	//
	// example:
	//
	// on
	XForwardedFor_SLBID *string `json:"XForwardedFor_SLBID,omitempty" xml:"XForwardedFor_SLBID,omitempty"`
	// Specifies whether to use the `SLB-IP` header to retrieve the virtual IP address of the client. Valid values:
	//
	// 	- **on**
	//
	// 	- **off**
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
	// 	- **on**
	//
	// 	- **off**
	//
	// example:
	//
	// on
	XForwardedFor_proto *string `json:"XForwardedFor_proto,omitempty" xml:"XForwardedFor_proto,omitempty"`
}

func (s CreateLoadBalancerHTTPSListenerRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateLoadBalancerHTTPSListenerRequest) GoString() string {
	return s.String()
}

func (s *CreateLoadBalancerHTTPSListenerRequest) GetAclId() *string {
	return s.AclId
}

func (s *CreateLoadBalancerHTTPSListenerRequest) GetAclStatus() *string {
	return s.AclStatus
}

func (s *CreateLoadBalancerHTTPSListenerRequest) GetAclType() *string {
	return s.AclType
}

func (s *CreateLoadBalancerHTTPSListenerRequest) GetBackendServerPort() *int32 {
	return s.BackendServerPort
}

func (s *CreateLoadBalancerHTTPSListenerRequest) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *CreateLoadBalancerHTTPSListenerRequest) GetCACertificateId() *string {
	return s.CACertificateId
}

func (s *CreateLoadBalancerHTTPSListenerRequest) GetCookie() *string {
	return s.Cookie
}

func (s *CreateLoadBalancerHTTPSListenerRequest) GetCookieTimeout() *int32 {
	return s.CookieTimeout
}

func (s *CreateLoadBalancerHTTPSListenerRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateLoadBalancerHTTPSListenerRequest) GetEnableHttp2() *string {
	return s.EnableHttp2
}

func (s *CreateLoadBalancerHTTPSListenerRequest) GetGzip() *string {
	return s.Gzip
}

func (s *CreateLoadBalancerHTTPSListenerRequest) GetHealthCheck() *string {
	return s.HealthCheck
}

func (s *CreateLoadBalancerHTTPSListenerRequest) GetHealthCheckConnectPort() *int32 {
	return s.HealthCheckConnectPort
}

func (s *CreateLoadBalancerHTTPSListenerRequest) GetHealthCheckDomain() *string {
	return s.HealthCheckDomain
}

func (s *CreateLoadBalancerHTTPSListenerRequest) GetHealthCheckHttpCode() *string {
	return s.HealthCheckHttpCode
}

func (s *CreateLoadBalancerHTTPSListenerRequest) GetHealthCheckInterval() *int32 {
	return s.HealthCheckInterval
}

func (s *CreateLoadBalancerHTTPSListenerRequest) GetHealthCheckMethod() *string {
	return s.HealthCheckMethod
}

func (s *CreateLoadBalancerHTTPSListenerRequest) GetHealthCheckTimeout() *int32 {
	return s.HealthCheckTimeout
}

func (s *CreateLoadBalancerHTTPSListenerRequest) GetHealthCheckURI() *string {
	return s.HealthCheckURI
}

func (s *CreateLoadBalancerHTTPSListenerRequest) GetHealthyThreshold() *int32 {
	return s.HealthyThreshold
}

func (s *CreateLoadBalancerHTTPSListenerRequest) GetIdleTimeout() *int32 {
	return s.IdleTimeout
}

func (s *CreateLoadBalancerHTTPSListenerRequest) GetListenerPort() *int32 {
	return s.ListenerPort
}

func (s *CreateLoadBalancerHTTPSListenerRequest) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *CreateLoadBalancerHTTPSListenerRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateLoadBalancerHTTPSListenerRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateLoadBalancerHTTPSListenerRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateLoadBalancerHTTPSListenerRequest) GetRequestTimeout() *int32 {
	return s.RequestTimeout
}

func (s *CreateLoadBalancerHTTPSListenerRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateLoadBalancerHTTPSListenerRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateLoadBalancerHTTPSListenerRequest) GetScheduler() *string {
	return s.Scheduler
}

func (s *CreateLoadBalancerHTTPSListenerRequest) GetServerCertificateId() *string {
	return s.ServerCertificateId
}

func (s *CreateLoadBalancerHTTPSListenerRequest) GetStickySession() *string {
	return s.StickySession
}

func (s *CreateLoadBalancerHTTPSListenerRequest) GetStickySessionType() *string {
	return s.StickySessionType
}

func (s *CreateLoadBalancerHTTPSListenerRequest) GetTLSCipherPolicy() *string {
	return s.TLSCipherPolicy
}

func (s *CreateLoadBalancerHTTPSListenerRequest) GetTag() []*CreateLoadBalancerHTTPSListenerRequestTag {
	return s.Tag
}

func (s *CreateLoadBalancerHTTPSListenerRequest) GetUnhealthyThreshold() *int32 {
	return s.UnhealthyThreshold
}

func (s *CreateLoadBalancerHTTPSListenerRequest) GetVServerGroupId() *string {
	return s.VServerGroupId
}

func (s *CreateLoadBalancerHTTPSListenerRequest) GetXForwardedFor() *string {
	return s.XForwardedFor
}

func (s *CreateLoadBalancerHTTPSListenerRequest) GetXForwardedFor_ClientSrcPort() *string {
	return s.XForwardedFor_ClientSrcPort
}

func (s *CreateLoadBalancerHTTPSListenerRequest) GetXForwardedFor_SLBID() *string {
	return s.XForwardedFor_SLBID
}

func (s *CreateLoadBalancerHTTPSListenerRequest) GetXForwardedFor_SLBIP() *string {
	return s.XForwardedFor_SLBIP
}

func (s *CreateLoadBalancerHTTPSListenerRequest) GetXForwardedFor_SLBPORT() *string {
	return s.XForwardedFor_SLBPORT
}

func (s *CreateLoadBalancerHTTPSListenerRequest) GetXForwardedFor_proto() *string {
	return s.XForwardedFor_proto
}

func (s *CreateLoadBalancerHTTPSListenerRequest) SetAclId(v string) *CreateLoadBalancerHTTPSListenerRequest {
	s.AclId = &v
	return s
}

func (s *CreateLoadBalancerHTTPSListenerRequest) SetAclStatus(v string) *CreateLoadBalancerHTTPSListenerRequest {
	s.AclStatus = &v
	return s
}

func (s *CreateLoadBalancerHTTPSListenerRequest) SetAclType(v string) *CreateLoadBalancerHTTPSListenerRequest {
	s.AclType = &v
	return s
}

func (s *CreateLoadBalancerHTTPSListenerRequest) SetBackendServerPort(v int32) *CreateLoadBalancerHTTPSListenerRequest {
	s.BackendServerPort = &v
	return s
}

func (s *CreateLoadBalancerHTTPSListenerRequest) SetBandwidth(v int32) *CreateLoadBalancerHTTPSListenerRequest {
	s.Bandwidth = &v
	return s
}

func (s *CreateLoadBalancerHTTPSListenerRequest) SetCACertificateId(v string) *CreateLoadBalancerHTTPSListenerRequest {
	s.CACertificateId = &v
	return s
}

func (s *CreateLoadBalancerHTTPSListenerRequest) SetCookie(v string) *CreateLoadBalancerHTTPSListenerRequest {
	s.Cookie = &v
	return s
}

func (s *CreateLoadBalancerHTTPSListenerRequest) SetCookieTimeout(v int32) *CreateLoadBalancerHTTPSListenerRequest {
	s.CookieTimeout = &v
	return s
}

func (s *CreateLoadBalancerHTTPSListenerRequest) SetDescription(v string) *CreateLoadBalancerHTTPSListenerRequest {
	s.Description = &v
	return s
}

func (s *CreateLoadBalancerHTTPSListenerRequest) SetEnableHttp2(v string) *CreateLoadBalancerHTTPSListenerRequest {
	s.EnableHttp2 = &v
	return s
}

func (s *CreateLoadBalancerHTTPSListenerRequest) SetGzip(v string) *CreateLoadBalancerHTTPSListenerRequest {
	s.Gzip = &v
	return s
}

func (s *CreateLoadBalancerHTTPSListenerRequest) SetHealthCheck(v string) *CreateLoadBalancerHTTPSListenerRequest {
	s.HealthCheck = &v
	return s
}

func (s *CreateLoadBalancerHTTPSListenerRequest) SetHealthCheckConnectPort(v int32) *CreateLoadBalancerHTTPSListenerRequest {
	s.HealthCheckConnectPort = &v
	return s
}

func (s *CreateLoadBalancerHTTPSListenerRequest) SetHealthCheckDomain(v string) *CreateLoadBalancerHTTPSListenerRequest {
	s.HealthCheckDomain = &v
	return s
}

func (s *CreateLoadBalancerHTTPSListenerRequest) SetHealthCheckHttpCode(v string) *CreateLoadBalancerHTTPSListenerRequest {
	s.HealthCheckHttpCode = &v
	return s
}

func (s *CreateLoadBalancerHTTPSListenerRequest) SetHealthCheckInterval(v int32) *CreateLoadBalancerHTTPSListenerRequest {
	s.HealthCheckInterval = &v
	return s
}

func (s *CreateLoadBalancerHTTPSListenerRequest) SetHealthCheckMethod(v string) *CreateLoadBalancerHTTPSListenerRequest {
	s.HealthCheckMethod = &v
	return s
}

func (s *CreateLoadBalancerHTTPSListenerRequest) SetHealthCheckTimeout(v int32) *CreateLoadBalancerHTTPSListenerRequest {
	s.HealthCheckTimeout = &v
	return s
}

func (s *CreateLoadBalancerHTTPSListenerRequest) SetHealthCheckURI(v string) *CreateLoadBalancerHTTPSListenerRequest {
	s.HealthCheckURI = &v
	return s
}

func (s *CreateLoadBalancerHTTPSListenerRequest) SetHealthyThreshold(v int32) *CreateLoadBalancerHTTPSListenerRequest {
	s.HealthyThreshold = &v
	return s
}

func (s *CreateLoadBalancerHTTPSListenerRequest) SetIdleTimeout(v int32) *CreateLoadBalancerHTTPSListenerRequest {
	s.IdleTimeout = &v
	return s
}

func (s *CreateLoadBalancerHTTPSListenerRequest) SetListenerPort(v int32) *CreateLoadBalancerHTTPSListenerRequest {
	s.ListenerPort = &v
	return s
}

func (s *CreateLoadBalancerHTTPSListenerRequest) SetLoadBalancerId(v string) *CreateLoadBalancerHTTPSListenerRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *CreateLoadBalancerHTTPSListenerRequest) SetOwnerAccount(v string) *CreateLoadBalancerHTTPSListenerRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateLoadBalancerHTTPSListenerRequest) SetOwnerId(v int64) *CreateLoadBalancerHTTPSListenerRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateLoadBalancerHTTPSListenerRequest) SetRegionId(v string) *CreateLoadBalancerHTTPSListenerRequest {
	s.RegionId = &v
	return s
}

func (s *CreateLoadBalancerHTTPSListenerRequest) SetRequestTimeout(v int32) *CreateLoadBalancerHTTPSListenerRequest {
	s.RequestTimeout = &v
	return s
}

func (s *CreateLoadBalancerHTTPSListenerRequest) SetResourceOwnerAccount(v string) *CreateLoadBalancerHTTPSListenerRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateLoadBalancerHTTPSListenerRequest) SetResourceOwnerId(v int64) *CreateLoadBalancerHTTPSListenerRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateLoadBalancerHTTPSListenerRequest) SetScheduler(v string) *CreateLoadBalancerHTTPSListenerRequest {
	s.Scheduler = &v
	return s
}

func (s *CreateLoadBalancerHTTPSListenerRequest) SetServerCertificateId(v string) *CreateLoadBalancerHTTPSListenerRequest {
	s.ServerCertificateId = &v
	return s
}

func (s *CreateLoadBalancerHTTPSListenerRequest) SetStickySession(v string) *CreateLoadBalancerHTTPSListenerRequest {
	s.StickySession = &v
	return s
}

func (s *CreateLoadBalancerHTTPSListenerRequest) SetStickySessionType(v string) *CreateLoadBalancerHTTPSListenerRequest {
	s.StickySessionType = &v
	return s
}

func (s *CreateLoadBalancerHTTPSListenerRequest) SetTLSCipherPolicy(v string) *CreateLoadBalancerHTTPSListenerRequest {
	s.TLSCipherPolicy = &v
	return s
}

func (s *CreateLoadBalancerHTTPSListenerRequest) SetTag(v []*CreateLoadBalancerHTTPSListenerRequestTag) *CreateLoadBalancerHTTPSListenerRequest {
	s.Tag = v
	return s
}

func (s *CreateLoadBalancerHTTPSListenerRequest) SetUnhealthyThreshold(v int32) *CreateLoadBalancerHTTPSListenerRequest {
	s.UnhealthyThreshold = &v
	return s
}

func (s *CreateLoadBalancerHTTPSListenerRequest) SetVServerGroupId(v string) *CreateLoadBalancerHTTPSListenerRequest {
	s.VServerGroupId = &v
	return s
}

func (s *CreateLoadBalancerHTTPSListenerRequest) SetXForwardedFor(v string) *CreateLoadBalancerHTTPSListenerRequest {
	s.XForwardedFor = &v
	return s
}

func (s *CreateLoadBalancerHTTPSListenerRequest) SetXForwardedFor_ClientSrcPort(v string) *CreateLoadBalancerHTTPSListenerRequest {
	s.XForwardedFor_ClientSrcPort = &v
	return s
}

func (s *CreateLoadBalancerHTTPSListenerRequest) SetXForwardedFor_SLBID(v string) *CreateLoadBalancerHTTPSListenerRequest {
	s.XForwardedFor_SLBID = &v
	return s
}

func (s *CreateLoadBalancerHTTPSListenerRequest) SetXForwardedFor_SLBIP(v string) *CreateLoadBalancerHTTPSListenerRequest {
	s.XForwardedFor_SLBIP = &v
	return s
}

func (s *CreateLoadBalancerHTTPSListenerRequest) SetXForwardedFor_SLBPORT(v string) *CreateLoadBalancerHTTPSListenerRequest {
	s.XForwardedFor_SLBPORT = &v
	return s
}

func (s *CreateLoadBalancerHTTPSListenerRequest) SetXForwardedFor_proto(v string) *CreateLoadBalancerHTTPSListenerRequest {
	s.XForwardedFor_proto = &v
	return s
}

func (s *CreateLoadBalancerHTTPSListenerRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateLoadBalancerHTTPSListenerRequestTag struct {
	// The tag key of the resource. You can specify up to 20 tag keys.
	//
	// The tag key cannot be an empty string. The tag key must be 1 to 64 characters in length and cannot start with `aliyun` or `acs:`. The tag key cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the resource. You can specify up to 20 tag values. The tag value cannot be an empty string.
	//
	// The tag value can be up to 128 characters in length and cannot start with `acs:` or `aliyun`. The tag value cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateLoadBalancerHTTPSListenerRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateLoadBalancerHTTPSListenerRequestTag) GoString() string {
	return s.String()
}

func (s *CreateLoadBalancerHTTPSListenerRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateLoadBalancerHTTPSListenerRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateLoadBalancerHTTPSListenerRequestTag) SetKey(v string) *CreateLoadBalancerHTTPSListenerRequestTag {
	s.Key = &v
	return s
}

func (s *CreateLoadBalancerHTTPSListenerRequestTag) SetValue(v string) *CreateLoadBalancerHTTPSListenerRequestTag {
	s.Value = &v
	return s
}

func (s *CreateLoadBalancerHTTPSListenerRequestTag) Validate() error {
	return dara.Validate(s)
}
