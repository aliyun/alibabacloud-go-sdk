// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLoadBalancerHTTPSListenerAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAclId(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody
	GetAclId() *string
	SetAclIds(v *DescribeLoadBalancerHTTPSListenerAttributeResponseBodyAclIds) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody
	GetAclIds() *DescribeLoadBalancerHTTPSListenerAttributeResponseBodyAclIds
	SetAclStatus(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody
	GetAclStatus() *string
	SetAclType(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody
	GetAclType() *string
	SetBackendServerPort(v int32) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody
	GetBackendServerPort() *int32
	SetBandwidth(v int32) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody
	GetBandwidth() *int32
	SetCACertificateId(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody
	GetCACertificateId() *string
	SetCookie(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody
	GetCookie() *string
	SetCookieTimeout(v int32) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody
	GetCookieTimeout() *int32
	SetDescription(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody
	GetDescription() *string
	SetDomainExtensions(v *DescribeLoadBalancerHTTPSListenerAttributeResponseBodyDomainExtensions) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody
	GetDomainExtensions() *DescribeLoadBalancerHTTPSListenerAttributeResponseBodyDomainExtensions
	SetEnableHttp2(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody
	GetEnableHttp2() *string
	SetGzip(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody
	GetGzip() *string
	SetHealthCheck(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody
	GetHealthCheck() *string
	SetHealthCheckConnectPort(v int32) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody
	GetHealthCheckConnectPort() *int32
	SetHealthCheckDomain(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody
	GetHealthCheckDomain() *string
	SetHealthCheckHttpCode(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody
	GetHealthCheckHttpCode() *string
	SetHealthCheckInterval(v int32) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody
	GetHealthCheckInterval() *int32
	SetHealthCheckMethod(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody
	GetHealthCheckMethod() *string
	SetHealthCheckTimeout(v int32) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody
	GetHealthCheckTimeout() *int32
	SetHealthCheckURI(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody
	GetHealthCheckURI() *string
	SetHealthyThreshold(v int32) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody
	GetHealthyThreshold() *int32
	SetIdleTimeout(v int32) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody
	GetIdleTimeout() *int32
	SetListenerPort(v int32) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody
	GetListenerPort() *int32
	SetLoadBalancerId(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody
	GetLoadBalancerId() *string
	SetRequestId(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody
	GetRequestId() *string
	SetRequestTimeout(v int32) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody
	GetRequestTimeout() *int32
	SetRules(v *DescribeLoadBalancerHTTPSListenerAttributeResponseBodyRules) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody
	GetRules() *DescribeLoadBalancerHTTPSListenerAttributeResponseBodyRules
	SetScheduler(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody
	GetScheduler() *string
	SetSecurityStatus(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody
	GetSecurityStatus() *string
	SetServerCertificateId(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody
	GetServerCertificateId() *string
	SetStatus(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody
	GetStatus() *string
	SetStickySession(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody
	GetStickySession() *string
	SetStickySessionType(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody
	GetStickySessionType() *string
	SetTLSCipherPolicy(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody
	GetTLSCipherPolicy() *string
	SetTags(v *DescribeLoadBalancerHTTPSListenerAttributeResponseBodyTags) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody
	GetTags() *DescribeLoadBalancerHTTPSListenerAttributeResponseBodyTags
	SetUnhealthyThreshold(v int32) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody
	GetUnhealthyThreshold() *int32
	SetVServerGroupId(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody
	GetVServerGroupId() *string
	SetXForwardedFor(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody
	GetXForwardedFor() *string
	SetXForwardedFor_ClientCertClientVerify(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody
	GetXForwardedFor_ClientCertClientVerify() *string
	SetXForwardedFor_ClientCertFingerprint(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody
	GetXForwardedFor_ClientCertFingerprint() *string
	SetXForwardedFor_ClientCertIssuerDN(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody
	GetXForwardedFor_ClientCertIssuerDN() *string
	SetXForwardedFor_ClientCertSubjectDN(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody
	GetXForwardedFor_ClientCertSubjectDN() *string
	SetXForwardedFor_ClientSrcPort(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody
	GetXForwardedFor_ClientSrcPort() *string
	SetXForwardedFor_SLBID(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody
	GetXForwardedFor_SLBID() *string
	SetXForwardedFor_SLBIP(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody
	GetXForwardedFor_SLBIP() *string
	SetXForwardedFor_SLBPORT(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody
	GetXForwardedFor_SLBPORT() *string
	SetXForwardedFor_proto(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody
	GetXForwardedFor_proto() *string
}

type DescribeLoadBalancerHTTPSListenerAttributeResponseBody struct {
	// The ID of the network ACL that is associated with a listener.
	//
	// > This parameter is required when **AclStatus*	- is set to **on**.
	//
	// example:
	//
	// acl-a2do9e413e0spzasx****
	AclId *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
	// The ID of the network access control list (ACL) that is associated with the listener.
	AclIds *DescribeLoadBalancerHTTPSListenerAttributeResponseBodyAclIds `json:"AclIds,omitempty" xml:"AclIds,omitempty" type:"Struct"`
	// Indicates whether access control is enabled. Valid values:
	//
	// 	- **on**
	//
	// 	- **off**
	//
	// example:
	//
	// off
	AclStatus *string `json:"AclStatus,omitempty" xml:"AclStatus,omitempty"`
	// The type of the access control list (ACL). Valid values:
	//
	// 	- **white**: a whitelist. Only requests from the IP addresses or CIDR blocks in the ACL are forwarded. Whitelists apply to scenarios where you want to allow only specific IP addresses to access an application. Your service may be adversely affected if the whitelist is not properly configured. If a whitelist is configured, only requests from IP addresses that are added to the whitelist are forwarded by the listener.
	//
	// If you enable a whitelist but do not add an IP address to the ACL, the listener forwards all requests.
	//
	// 	- **black**: a blacklist. All requests from the IP addresses or CIDR blocks in the network ACL are rejected. Blacklists apply to scenarios where you want to block access from specified IP addresses to an application.
	//
	// If a blacklist is configured for a listener but no IP address is added to the blacklist, the listener forwards all requests.
	//
	// > This parameter is required when **AclStatus*	- is set to **on**.
	//
	// example:
	//
	// white
	AclType *string `json:"AclType,omitempty" xml:"AclType,omitempty"`
	// The backend port that is used by the CLB instance.
	//
	// example:
	//
	// 8080
	BackendServerPort *int32 `json:"BackendServerPort,omitempty" xml:"BackendServerPort,omitempty"`
	// The maximum bandwidth of the listener. Unit: Mbit/s.
	//
	// example:
	//
	// 10
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The ID of the certification authority (CA) certificate.
	//
	// example:
	//
	// idkp-234-cn-test-0**
	CACertificateId *string `json:"CACertificateId,omitempty" xml:"CACertificateId,omitempty"`
	// The cookie that is configured on the server.
	//
	// example:
	//
	// testCookie
	Cookie *string `json:"Cookie,omitempty" xml:"Cookie,omitempty"`
	// The timeout period of a cookie.
	//
	// example:
	//
	// 500
	CookieTimeout *int32 `json:"CookieTimeout,omitempty" xml:"CookieTimeout,omitempty"`
	// The name of the listener.
	//
	// example:
	//
	// HTTPS_443
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// A list of additional certificates.
	DomainExtensions *DescribeLoadBalancerHTTPSListenerAttributeResponseBodyDomainExtensions `json:"DomainExtensions,omitempty" xml:"DomainExtensions,omitempty" type:"Struct"`
	// Indicates whether `HTTP/2` is used. Valid values:
	//
	// 	- **on**
	//
	// 	- **off**
	//
	// example:
	//
	// off
	EnableHttp2 *string `json:"EnableHttp2,omitempty" xml:"EnableHttp2,omitempty"`
	// Indicates whether `Gzip` compression is enabled. Valid values:
	//
	// 	- **on**
	//
	// 	- **off**
	//
	// example:
	//
	// on
	Gzip *string `json:"Gzip,omitempty" xml:"Gzip,omitempty"`
	// Indicates whether the health check feature is enabled. Valid values:
	//
	// 	- **on**
	//
	// 	- **off**
	//
	// example:
	//
	// on
	HealthCheck *string `json:"HealthCheck,omitempty" xml:"HealthCheck,omitempty"`
	// The port that is used for health checks.
	//
	// > This parameter is required when **HealthCheck*	- is set to **on**.
	//
	// example:
	//
	// 8080
	HealthCheckConnectPort *int32 `json:"HealthCheckConnectPort,omitempty" xml:"HealthCheckConnectPort,omitempty"`
	// The domain name that you want to use for health checks.
	//
	// example:
	//
	// www.test.com
	HealthCheckDomain *string `json:"HealthCheckDomain,omitempty" xml:"HealthCheckDomain,omitempty"`
	// The HTTP status code for a successful health check.
	//
	// example:
	//
	// http_2xx,http_3xx
	HealthCheckHttpCode *string `json:"HealthCheckHttpCode,omitempty" xml:"HealthCheckHttpCode,omitempty"`
	// The interval at which health checks are performed. Unit: seconds.
	//
	// example:
	//
	// 5
	HealthCheckInterval *int32 `json:"HealthCheckInterval,omitempty" xml:"HealthCheckInterval,omitempty"`
	// The health check method used by HTTP listeners. Valid values: **head*	- and **get**.
	//
	// > This parameter is available only when **HealthCheck*	- is set to **on**.
	//
	// example:
	//
	// get
	HealthCheckMethod *string `json:"HealthCheckMethod,omitempty" xml:"HealthCheckMethod,omitempty"`
	// The maximum timeout period of a health check. Unit: seconds.
	//
	// example:
	//
	// 3
	HealthCheckTimeout *int32 `json:"HealthCheckTimeout,omitempty" xml:"HealthCheckTimeout,omitempty"`
	// The URL path that is used for health checks.
	//
	// example:
	//
	// /test/index.html
	HealthCheckURI *string `json:"HealthCheckURI,omitempty" xml:"HealthCheckURI,omitempty"`
	// The healthy threshold.
	//
	// example:
	//
	// 4
	HealthyThreshold *int32 `json:"HealthyThreshold,omitempty" xml:"HealthyThreshold,omitempty"`
	// The timeout period of an idle connection. Valid values: **1*	- to **60**. Default value: **15**. Unit: seconds.
	//
	// If no request is received within the specified timeout period, CLB closes the connection. When a request is received, CLB establishes a new connection.
	//
	// example:
	//
	// 23
	IdleTimeout *int32 `json:"IdleTimeout,omitempty" xml:"IdleTimeout,omitempty"`
	// The frontend port that is used by the CLB instance.
	//
	// example:
	//
	// 80
	ListenerPort *int32 `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	// The CLB instance ID.
	//
	// example:
	//
	// lb-bp1mxu5r8lau****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 365F4154-92F6-4AE4-92F8-7FF3********
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The timeout period of a request. Valid values: **1*	- to **180**. Default value: **60**. Unit: seconds.
	//
	// If no response is received from a backend server within the specified timeout period, CLB returns the HTTP 504 status code to the client.
	//
	// example:
	//
	// 43
	RequestTimeout *int32 `json:"RequestTimeout,omitempty" xml:"RequestTimeout,omitempty"`
	// The list of forwarding rules that are associated with the listener.
	Rules *DescribeLoadBalancerHTTPSListenerAttributeResponseBodyRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Struct"`
	// The routing algorithm. Valid values: **wrr*	- and **rr**.
	//
	// 	- **wrr**: Backend servers that have higher weights receive more requests than backend servers that have lower weights.
	//
	// 	- **rr**: Requests are distributed to backend servers in sequence.
	//
	// example:
	//
	// wrr
	Scheduler *string `json:"Scheduler,omitempty" xml:"Scheduler,omitempty"`
	// Indicates whether the listener is in the Secure state. Valid values:
	//
	// 	- **on**
	//
	// 	- **off**
	//
	// example:
	//
	// on
	SecurityStatus *string `json:"SecurityStatus,omitempty" xml:"SecurityStatus,omitempty"`
	// The ID of the server certificate.
	//
	// example:
	//
	// idkp-123-cn-test-0**
	ServerCertificateId *string `json:"ServerCertificateId,omitempty" xml:"ServerCertificateId,omitempty"`
	// The status of the listener. Valid values:
	//
	// 	- **running**
	//
	// 	- **stopped**
	//
	// example:
	//
	// stopped
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates whether session persistence is enabled. Valid values:
	//
	// 	- **on**
	//
	// 	- **off**
	//
	// example:
	//
	// on
	StickySession *string `json:"StickySession,omitempty" xml:"StickySession,omitempty"`
	// The method that is used to handle a cookie.
	//
	// Valid values: **insert*	- and **server**.
	//
	// 	- **insert**: inserts a cookie.
	//
	//     CLB inserts a cookie (SERVERID) into the first HTTP or HTTPS response packet that is sent to a client. The next request from the client will contain this cookie, and the listener will distribute this request to the recorded backend server.
	//
	// 	- **server**: rewrites a cookie.
	//
	//     When CLB detects a user-defined cookie, it overwrites the original cookie with the user-defined cookie. The next request from the client carries the user-defined cookie, and the listener will distribute the request to the recorded backend server.
	//
	// example:
	//
	// insert
	StickySessionType *string `json:"StickySessionType,omitempty" xml:"StickySessionType,omitempty"`
	// The Transport Layer Security (TLS) security policy for a high-performance CLB instance.
	//
	// Each security policy contains TLS protocol versions and cipher suites available for HTTPS. Valid values:
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
	// tls_cipher_policy_1_0
	TLSCipherPolicy *string `json:"TLSCipherPolicy,omitempty" xml:"TLSCipherPolicy,omitempty"`
	// The tags.
	Tags *DescribeLoadBalancerHTTPSListenerAttributeResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	// The unhealthy threshold.
	//
	// example:
	//
	// 4
	UnhealthyThreshold *int32 `json:"UnhealthyThreshold,omitempty" xml:"UnhealthyThreshold,omitempty"`
	// The ID of the associated server group.
	//
	// example:
	//
	// rsp-cige6j5e********
	VServerGroupId *string `json:"VServerGroupId,omitempty" xml:"VServerGroupId,omitempty"`
	// Indicates whether the `X-Forwarded-For` header is used to retrieve client IP addresses. Valid values:
	//
	// 	- **on**
	//
	// 	- **off**
	//
	// example:
	//
	// on
	XForwardedFor *string `json:"XForwardedFor,omitempty" xml:"XForwardedFor,omitempty"`
	// Indicates whether the `XForwardedFor_ClientCertClientVerify` header is used to retrieve the verification result of the client certificate. Valid values:
	//
	// 	- **on**
	//
	// 	- **off**
	//
	// example:
	//
	// off
	XForwardedFor_ClientCertClientVerify *string `json:"XForwardedFor_ClientCertClientVerify,omitempty" xml:"XForwardedFor_ClientCertClientVerify,omitempty"`
	// Indicates whether the `XForwardedFor_ClientCertFingerprint` header is used to retrieve the fingerprint of the client certificate. Valid values:
	//
	// 	- **on**
	//
	// 	- **off**
	//
	// example:
	//
	// off
	XForwardedFor_ClientCertFingerprint *string `json:"XForwardedFor_ClientCertFingerprint,omitempty" xml:"XForwardedFor_ClientCertFingerprint,omitempty"`
	// Indicates whether the `XForwardedFor_ClientCertIssuerDN` header is used to retrieve information about the authority that issues the client certificate. Valid values:
	//
	// 	- **on**
	//
	// 	- **off**
	//
	// example:
	//
	// off
	XForwardedFor_ClientCertIssuerDN *string `json:"XForwardedFor_ClientCertIssuerDN,omitempty" xml:"XForwardedFor_ClientCertIssuerDN,omitempty"`
	// Indicates whether the `XForwardedFor_ClientCertSubjectDN` header is used to retrieve information about the owner of the client certificate. Valid values:
	//
	// 	- **on**
	//
	// 	- **off**
	//
	// example:
	//
	// off
	XForwardedFor_ClientCertSubjectDN *string `json:"XForwardedFor_ClientCertSubjectDN,omitempty" xml:"XForwardedFor_ClientCertSubjectDN,omitempty"`
	// Indicates whether the `XForwardedFor_ClientSrcPort` header is used to retrieve the client port. Valid values:
	//
	// 	- **on**
	//
	// 	- **off**
	//
	// example:
	//
	// off
	XForwardedFor_ClientSrcPort *string `json:"XForwardedFor_ClientSrcPort,omitempty" xml:"XForwardedFor_ClientSrcPort,omitempty"`
	// Indicates whether the `SLB-ID` header is used to retrieve the ID of the ALB instance. Valid values:
	//
	// 	- **on**
	//
	// 	- **off**
	//
	// example:
	//
	// on
	XForwardedFor_SLBID *string `json:"XForwardedFor_SLBID,omitempty" xml:"XForwardedFor_SLBID,omitempty"`
	// Indicates whether the `SLB-IP` header is used to retrieve the virtual IP address requested by the client. Valid values:
	//
	// 	- **on**
	//
	// 	- **off**
	//
	// example:
	//
	// on
	XForwardedFor_SLBIP *string `json:"XForwardedFor_SLBIP,omitempty" xml:"XForwardedFor_SLBIP,omitempty"`
	// Indicates whether the `XForwardedFor_SLBPORT` header is used to retrieve the listening port. Valid values:
	//
	// 	- **on**
	//
	// 	- **off**
	//
	// example:
	//
	// off
	XForwardedFor_SLBPORT *string `json:"XForwardedFor_SLBPORT,omitempty" xml:"XForwardedFor_SLBPORT,omitempty"`
	// Indicates whether the `X-Forwarded-Proto` header is used to retrieve the listener protocol. Valid values:
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

func (s DescribeLoadBalancerHTTPSListenerAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLoadBalancerHTTPSListenerAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) GetAclId() *string {
	return s.AclId
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) GetAclIds() *DescribeLoadBalancerHTTPSListenerAttributeResponseBodyAclIds {
	return s.AclIds
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) GetAclStatus() *string {
	return s.AclStatus
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) GetAclType() *string {
	return s.AclType
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) GetBackendServerPort() *int32 {
	return s.BackendServerPort
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) GetCACertificateId() *string {
	return s.CACertificateId
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) GetCookie() *string {
	return s.Cookie
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) GetCookieTimeout() *int32 {
	return s.CookieTimeout
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) GetDescription() *string {
	return s.Description
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) GetDomainExtensions() *DescribeLoadBalancerHTTPSListenerAttributeResponseBodyDomainExtensions {
	return s.DomainExtensions
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) GetEnableHttp2() *string {
	return s.EnableHttp2
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) GetGzip() *string {
	return s.Gzip
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) GetHealthCheck() *string {
	return s.HealthCheck
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) GetHealthCheckConnectPort() *int32 {
	return s.HealthCheckConnectPort
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) GetHealthCheckDomain() *string {
	return s.HealthCheckDomain
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) GetHealthCheckHttpCode() *string {
	return s.HealthCheckHttpCode
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) GetHealthCheckInterval() *int32 {
	return s.HealthCheckInterval
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) GetHealthCheckMethod() *string {
	return s.HealthCheckMethod
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) GetHealthCheckTimeout() *int32 {
	return s.HealthCheckTimeout
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) GetHealthCheckURI() *string {
	return s.HealthCheckURI
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) GetHealthyThreshold() *int32 {
	return s.HealthyThreshold
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) GetIdleTimeout() *int32 {
	return s.IdleTimeout
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) GetListenerPort() *int32 {
	return s.ListenerPort
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) GetRequestTimeout() *int32 {
	return s.RequestTimeout
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) GetRules() *DescribeLoadBalancerHTTPSListenerAttributeResponseBodyRules {
	return s.Rules
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) GetScheduler() *string {
	return s.Scheduler
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) GetSecurityStatus() *string {
	return s.SecurityStatus
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) GetServerCertificateId() *string {
	return s.ServerCertificateId
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) GetStickySession() *string {
	return s.StickySession
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) GetStickySessionType() *string {
	return s.StickySessionType
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) GetTLSCipherPolicy() *string {
	return s.TLSCipherPolicy
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) GetTags() *DescribeLoadBalancerHTTPSListenerAttributeResponseBodyTags {
	return s.Tags
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) GetUnhealthyThreshold() *int32 {
	return s.UnhealthyThreshold
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) GetVServerGroupId() *string {
	return s.VServerGroupId
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) GetXForwardedFor() *string {
	return s.XForwardedFor
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) GetXForwardedFor_ClientCertClientVerify() *string {
	return s.XForwardedFor_ClientCertClientVerify
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) GetXForwardedFor_ClientCertFingerprint() *string {
	return s.XForwardedFor_ClientCertFingerprint
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) GetXForwardedFor_ClientCertIssuerDN() *string {
	return s.XForwardedFor_ClientCertIssuerDN
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) GetXForwardedFor_ClientCertSubjectDN() *string {
	return s.XForwardedFor_ClientCertSubjectDN
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) GetXForwardedFor_ClientSrcPort() *string {
	return s.XForwardedFor_ClientSrcPort
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) GetXForwardedFor_SLBID() *string {
	return s.XForwardedFor_SLBID
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) GetXForwardedFor_SLBIP() *string {
	return s.XForwardedFor_SLBIP
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) GetXForwardedFor_SLBPORT() *string {
	return s.XForwardedFor_SLBPORT
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) GetXForwardedFor_proto() *string {
	return s.XForwardedFor_proto
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) SetAclId(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody {
	s.AclId = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) SetAclIds(v *DescribeLoadBalancerHTTPSListenerAttributeResponseBodyAclIds) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody {
	s.AclIds = v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) SetAclStatus(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody {
	s.AclStatus = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) SetAclType(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody {
	s.AclType = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) SetBackendServerPort(v int32) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody {
	s.BackendServerPort = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) SetBandwidth(v int32) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody {
	s.Bandwidth = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) SetCACertificateId(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody {
	s.CACertificateId = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) SetCookie(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody {
	s.Cookie = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) SetCookieTimeout(v int32) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody {
	s.CookieTimeout = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) SetDescription(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) SetDomainExtensions(v *DescribeLoadBalancerHTTPSListenerAttributeResponseBodyDomainExtensions) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody {
	s.DomainExtensions = v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) SetEnableHttp2(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody {
	s.EnableHttp2 = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) SetGzip(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody {
	s.Gzip = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) SetHealthCheck(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody {
	s.HealthCheck = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) SetHealthCheckConnectPort(v int32) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody {
	s.HealthCheckConnectPort = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) SetHealthCheckDomain(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody {
	s.HealthCheckDomain = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) SetHealthCheckHttpCode(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody {
	s.HealthCheckHttpCode = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) SetHealthCheckInterval(v int32) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody {
	s.HealthCheckInterval = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) SetHealthCheckMethod(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody {
	s.HealthCheckMethod = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) SetHealthCheckTimeout(v int32) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody {
	s.HealthCheckTimeout = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) SetHealthCheckURI(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody {
	s.HealthCheckURI = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) SetHealthyThreshold(v int32) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody {
	s.HealthyThreshold = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) SetIdleTimeout(v int32) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody {
	s.IdleTimeout = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) SetListenerPort(v int32) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody {
	s.ListenerPort = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) SetLoadBalancerId(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody {
	s.LoadBalancerId = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) SetRequestId(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) SetRequestTimeout(v int32) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody {
	s.RequestTimeout = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) SetRules(v *DescribeLoadBalancerHTTPSListenerAttributeResponseBodyRules) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody {
	s.Rules = v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) SetScheduler(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody {
	s.Scheduler = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) SetSecurityStatus(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody {
	s.SecurityStatus = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) SetServerCertificateId(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody {
	s.ServerCertificateId = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) SetStatus(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) SetStickySession(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody {
	s.StickySession = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) SetStickySessionType(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody {
	s.StickySessionType = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) SetTLSCipherPolicy(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody {
	s.TLSCipherPolicy = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) SetTags(v *DescribeLoadBalancerHTTPSListenerAttributeResponseBodyTags) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody {
	s.Tags = v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) SetUnhealthyThreshold(v int32) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody {
	s.UnhealthyThreshold = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) SetVServerGroupId(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody {
	s.VServerGroupId = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) SetXForwardedFor(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody {
	s.XForwardedFor = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) SetXForwardedFor_ClientCertClientVerify(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody {
	s.XForwardedFor_ClientCertClientVerify = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) SetXForwardedFor_ClientCertFingerprint(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody {
	s.XForwardedFor_ClientCertFingerprint = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) SetXForwardedFor_ClientCertIssuerDN(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody {
	s.XForwardedFor_ClientCertIssuerDN = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) SetXForwardedFor_ClientCertSubjectDN(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody {
	s.XForwardedFor_ClientCertSubjectDN = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) SetXForwardedFor_ClientSrcPort(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody {
	s.XForwardedFor_ClientSrcPort = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) SetXForwardedFor_SLBID(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody {
	s.XForwardedFor_SLBID = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) SetXForwardedFor_SLBIP(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody {
	s.XForwardedFor_SLBIP = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) SetXForwardedFor_SLBPORT(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody {
	s.XForwardedFor_SLBPORT = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) SetXForwardedFor_proto(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody {
	s.XForwardedFor_proto = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) Validate() error {
	if s.AclIds != nil {
		if err := s.AclIds.Validate(); err != nil {
			return err
		}
	}
	if s.DomainExtensions != nil {
		if err := s.DomainExtensions.Validate(); err != nil {
			return err
		}
	}
	if s.Rules != nil {
		if err := s.Rules.Validate(); err != nil {
			return err
		}
	}
	if s.Tags != nil {
		if err := s.Tags.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeLoadBalancerHTTPSListenerAttributeResponseBodyAclIds struct {
	AclId []*string `json:"AclId,omitempty" xml:"AclId,omitempty" type:"Repeated"`
}

func (s DescribeLoadBalancerHTTPSListenerAttributeResponseBodyAclIds) String() string {
	return dara.Prettify(s)
}

func (s DescribeLoadBalancerHTTPSListenerAttributeResponseBodyAclIds) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBodyAclIds) GetAclId() []*string {
	return s.AclId
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBodyAclIds) SetAclId(v []*string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBodyAclIds {
	s.AclId = v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBodyAclIds) Validate() error {
	return dara.Validate(s)
}

type DescribeLoadBalancerHTTPSListenerAttributeResponseBodyDomainExtensions struct {
	DomainExtension []*DescribeLoadBalancerHTTPSListenerAttributeResponseBodyDomainExtensionsDomainExtension `json:"DomainExtension,omitempty" xml:"DomainExtension,omitempty" type:"Repeated"`
}

func (s DescribeLoadBalancerHTTPSListenerAttributeResponseBodyDomainExtensions) String() string {
	return dara.Prettify(s)
}

func (s DescribeLoadBalancerHTTPSListenerAttributeResponseBodyDomainExtensions) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBodyDomainExtensions) GetDomainExtension() []*DescribeLoadBalancerHTTPSListenerAttributeResponseBodyDomainExtensionsDomainExtension {
	return s.DomainExtension
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBodyDomainExtensions) SetDomainExtension(v []*DescribeLoadBalancerHTTPSListenerAttributeResponseBodyDomainExtensionsDomainExtension) *DescribeLoadBalancerHTTPSListenerAttributeResponseBodyDomainExtensions {
	s.DomainExtension = v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBodyDomainExtensions) Validate() error {
	if s.DomainExtension != nil {
		for _, item := range s.DomainExtension {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeLoadBalancerHTTPSListenerAttributeResponseBodyDomainExtensionsDomainExtension struct {
	// The domain name.
	//
	// example:
	//
	// www.example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The ID of the additional certificate.
	//
	// example:
	//
	// 12
	DomainExtensionId *string `json:"DomainExtensionId,omitempty" xml:"DomainExtensionId,omitempty"`
	// The ID of the certificate used by the domain name.
	//
	// example:
	//
	// 13344444****
	ServerCertificateId *string `json:"ServerCertificateId,omitempty" xml:"ServerCertificateId,omitempty"`
}

func (s DescribeLoadBalancerHTTPSListenerAttributeResponseBodyDomainExtensionsDomainExtension) String() string {
	return dara.Prettify(s)
}

func (s DescribeLoadBalancerHTTPSListenerAttributeResponseBodyDomainExtensionsDomainExtension) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBodyDomainExtensionsDomainExtension) GetDomain() *string {
	return s.Domain
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBodyDomainExtensionsDomainExtension) GetDomainExtensionId() *string {
	return s.DomainExtensionId
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBodyDomainExtensionsDomainExtension) GetServerCertificateId() *string {
	return s.ServerCertificateId
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBodyDomainExtensionsDomainExtension) SetDomain(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBodyDomainExtensionsDomainExtension {
	s.Domain = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBodyDomainExtensionsDomainExtension) SetDomainExtensionId(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBodyDomainExtensionsDomainExtension {
	s.DomainExtensionId = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBodyDomainExtensionsDomainExtension) SetServerCertificateId(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBodyDomainExtensionsDomainExtension {
	s.ServerCertificateId = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBodyDomainExtensionsDomainExtension) Validate() error {
	return dara.Validate(s)
}

type DescribeLoadBalancerHTTPSListenerAttributeResponseBodyRules struct {
	Rule []*DescribeLoadBalancerHTTPSListenerAttributeResponseBodyRulesRule `json:"Rule,omitempty" xml:"Rule,omitempty" type:"Repeated"`
}

func (s DescribeLoadBalancerHTTPSListenerAttributeResponseBodyRules) String() string {
	return dara.Prettify(s)
}

func (s DescribeLoadBalancerHTTPSListenerAttributeResponseBodyRules) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBodyRules) GetRule() []*DescribeLoadBalancerHTTPSListenerAttributeResponseBodyRulesRule {
	return s.Rule
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBodyRules) SetRule(v []*DescribeLoadBalancerHTTPSListenerAttributeResponseBodyRulesRule) *DescribeLoadBalancerHTTPSListenerAttributeResponseBodyRules {
	s.Rule = v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBodyRules) Validate() error {
	if s.Rule != nil {
		for _, item := range s.Rule {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeLoadBalancerHTTPSListenerAttributeResponseBodyRulesRule struct {
	// The domain name.
	//
	// example:
	//
	// www.example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The ID of the forwarding rule.
	//
	// example:
	//
	// 23
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The name of the forwarding rule.
	//
	// example:
	//
	// example
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The request URL.
	//
	// example:
	//
	// /example
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// The ID of the server group specified in the forwarding rule.
	//
	// example:
	//
	// rsp-cige6j5e********
	VServerGroupId *string `json:"VServerGroupId,omitempty" xml:"VServerGroupId,omitempty"`
}

func (s DescribeLoadBalancerHTTPSListenerAttributeResponseBodyRulesRule) String() string {
	return dara.Prettify(s)
}

func (s DescribeLoadBalancerHTTPSListenerAttributeResponseBodyRulesRule) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBodyRulesRule) GetDomain() *string {
	return s.Domain
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBodyRulesRule) GetRuleId() *string {
	return s.RuleId
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBodyRulesRule) GetRuleName() *string {
	return s.RuleName
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBodyRulesRule) GetUrl() *string {
	return s.Url
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBodyRulesRule) GetVServerGroupId() *string {
	return s.VServerGroupId
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBodyRulesRule) SetDomain(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBodyRulesRule {
	s.Domain = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBodyRulesRule) SetRuleId(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBodyRulesRule {
	s.RuleId = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBodyRulesRule) SetRuleName(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBodyRulesRule {
	s.RuleName = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBodyRulesRule) SetUrl(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBodyRulesRule {
	s.Url = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBodyRulesRule) SetVServerGroupId(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBodyRulesRule {
	s.VServerGroupId = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBodyRulesRule) Validate() error {
	return dara.Validate(s)
}

type DescribeLoadBalancerHTTPSListenerAttributeResponseBodyTags struct {
	Tag []*DescribeLoadBalancerHTTPSListenerAttributeResponseBodyTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeLoadBalancerHTTPSListenerAttributeResponseBodyTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeLoadBalancerHTTPSListenerAttributeResponseBodyTags) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBodyTags) GetTag() []*DescribeLoadBalancerHTTPSListenerAttributeResponseBodyTagsTag {
	return s.Tag
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBodyTags) SetTag(v []*DescribeLoadBalancerHTTPSListenerAttributeResponseBodyTagsTag) *DescribeLoadBalancerHTTPSListenerAttributeResponseBodyTags {
	s.Tag = v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBodyTags) Validate() error {
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

type DescribeLoadBalancerHTTPSListenerAttributeResponseBodyTagsTag struct {
	// The key of tag N. Valid values of N: **1*	- to **20**. The tag value cannot be an empty string. The tag key can be up to 64 characters in length. The key cannot start with `acs:` or `aliyun` or contain `http://` or `https://`.
	//
	// example:
	//
	// TestKey
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The value of tag N. Valid values of N: **1*	- to **20**. The tag value can be an empty string. The tag value can be up to 128 characters in length, and cannot start with `acs:`. It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TestValue
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s DescribeLoadBalancerHTTPSListenerAttributeResponseBodyTagsTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeLoadBalancerHTTPSListenerAttributeResponseBodyTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBodyTagsTag) GetTagKey() *string {
	return s.TagKey
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBodyTagsTag) GetTagValue() *string {
	return s.TagValue
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBodyTagsTag) SetTagKey(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBodyTagsTag {
	s.TagKey = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBodyTagsTag) SetTagValue(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBodyTagsTag {
	s.TagValue = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBodyTagsTag) Validate() error {
	return dara.Validate(s)
}
