// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetLoadBalancerHTTPListenerAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAclId(v string) *SetLoadBalancerHTTPListenerAttributeRequest
	GetAclId() *string
	SetAclStatus(v string) *SetLoadBalancerHTTPListenerAttributeRequest
	GetAclStatus() *string
	SetAclType(v string) *SetLoadBalancerHTTPListenerAttributeRequest
	GetAclType() *string
	SetBandwidth(v int32) *SetLoadBalancerHTTPListenerAttributeRequest
	GetBandwidth() *int32
	SetCookie(v string) *SetLoadBalancerHTTPListenerAttributeRequest
	GetCookie() *string
	SetCookieTimeout(v int32) *SetLoadBalancerHTTPListenerAttributeRequest
	GetCookieTimeout() *int32
	SetDescription(v string) *SetLoadBalancerHTTPListenerAttributeRequest
	GetDescription() *string
	SetGzip(v string) *SetLoadBalancerHTTPListenerAttributeRequest
	GetGzip() *string
	SetHealthCheck(v string) *SetLoadBalancerHTTPListenerAttributeRequest
	GetHealthCheck() *string
	SetHealthCheckConnectPort(v int32) *SetLoadBalancerHTTPListenerAttributeRequest
	GetHealthCheckConnectPort() *int32
	SetHealthCheckDomain(v string) *SetLoadBalancerHTTPListenerAttributeRequest
	GetHealthCheckDomain() *string
	SetHealthCheckHttpCode(v string) *SetLoadBalancerHTTPListenerAttributeRequest
	GetHealthCheckHttpCode() *string
	SetHealthCheckInterval(v int32) *SetLoadBalancerHTTPListenerAttributeRequest
	GetHealthCheckInterval() *int32
	SetHealthCheckMethod(v string) *SetLoadBalancerHTTPListenerAttributeRequest
	GetHealthCheckMethod() *string
	SetHealthCheckTimeout(v int32) *SetLoadBalancerHTTPListenerAttributeRequest
	GetHealthCheckTimeout() *int32
	SetHealthCheckURI(v string) *SetLoadBalancerHTTPListenerAttributeRequest
	GetHealthCheckURI() *string
	SetHealthyThreshold(v int32) *SetLoadBalancerHTTPListenerAttributeRequest
	GetHealthyThreshold() *int32
	SetIdleTimeout(v int32) *SetLoadBalancerHTTPListenerAttributeRequest
	GetIdleTimeout() *int32
	SetListenerPort(v int32) *SetLoadBalancerHTTPListenerAttributeRequest
	GetListenerPort() *int32
	SetLoadBalancerId(v string) *SetLoadBalancerHTTPListenerAttributeRequest
	GetLoadBalancerId() *string
	SetOwnerAccount(v string) *SetLoadBalancerHTTPListenerAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *SetLoadBalancerHTTPListenerAttributeRequest
	GetOwnerId() *int64
	SetRegionId(v string) *SetLoadBalancerHTTPListenerAttributeRequest
	GetRegionId() *string
	SetRequestTimeout(v int32) *SetLoadBalancerHTTPListenerAttributeRequest
	GetRequestTimeout() *int32
	SetResourceOwnerAccount(v string) *SetLoadBalancerHTTPListenerAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *SetLoadBalancerHTTPListenerAttributeRequest
	GetResourceOwnerId() *int64
	SetScheduler(v string) *SetLoadBalancerHTTPListenerAttributeRequest
	GetScheduler() *string
	SetStickySession(v string) *SetLoadBalancerHTTPListenerAttributeRequest
	GetStickySession() *string
	SetStickySessionType(v string) *SetLoadBalancerHTTPListenerAttributeRequest
	GetStickySessionType() *string
	SetUnhealthyThreshold(v int32) *SetLoadBalancerHTTPListenerAttributeRequest
	GetUnhealthyThreshold() *int32
	SetVServerGroup(v string) *SetLoadBalancerHTTPListenerAttributeRequest
	GetVServerGroup() *string
	SetVServerGroupId(v string) *SetLoadBalancerHTTPListenerAttributeRequest
	GetVServerGroupId() *string
	SetXForwardedFor(v string) *SetLoadBalancerHTTPListenerAttributeRequest
	GetXForwardedFor() *string
	SetXForwardedFor_ClientSrcPort(v string) *SetLoadBalancerHTTPListenerAttributeRequest
	GetXForwardedFor_ClientSrcPort() *string
	SetXForwardedFor_SLBID(v string) *SetLoadBalancerHTTPListenerAttributeRequest
	GetXForwardedFor_SLBID() *string
	SetXForwardedFor_SLBIP(v string) *SetLoadBalancerHTTPListenerAttributeRequest
	GetXForwardedFor_SLBIP() *string
	SetXForwardedFor_SLBPORT(v string) *SetLoadBalancerHTTPListenerAttributeRequest
	GetXForwardedFor_SLBPORT() *string
	SetXForwardedFor_proto(v string) *SetLoadBalancerHTTPListenerAttributeRequest
	GetXForwardedFor_proto() *string
}

type SetLoadBalancerHTTPListenerAttributeRequest struct {
	// The ID of the access control list (ACL) that is associated with the listener.
	//
	// > This parameter is required when **AclStatus*	- is set to **on**.
	//
	// example:
	//
	// acl-uf60jw******
	AclId *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
	// Specifies whether to enable access control. Valid values:
	//
	// 	- **on**
	//
	// 	- **off**
	//
	// example:
	//
	// off
	AclStatus *string `json:"AclStatus,omitempty" xml:"AclStatus,omitempty"`
	// The type of the ACL. Valid values:
	//
	// 	- **white**: a whitelist. Only requests from the IP addresses or CIDR blocks in the ACL are forwarded. Whitelists apply to scenarios where you want to allow only specific IP addresses to access an application. Risks may occur if a whitelist is improperly configured. If a whitelist is configured, only requests from IP addresses that are added to the whitelist are forwarded by the listener.
	//
	// If you enable a whitelist but do not add an IP address to the ACL, the listener forwards all requests.
	//
	// 	- **black**: a blacklist. All requests from the IP addresses or CIDR blocks in the ACL are rejected. Blacklists apply to scenarios where you want to block access from specified IP addresses to an application.
	//
	// If a blacklist is configured for a listener but no IP addresses are added to the blacklist, the listener forwards all requests.
	//
	// > This parameter takes effect when the value of **AclStatus*	- is set to **on**.
	//
	// example:
	//
	// white
	AclType *string `json:"AclType,omitempty" xml:"AclType,omitempty"`
	// The maximum bandwidth of the listener. Unit: Mbit/s. Set the value to
	//
	// 	- **-1**: For a pay-by-data-transfer Internet-facing CLB instance, this value specifies that the bandwidth of the listener is unlimited.
	//
	// example:
	//
	// -1
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The cookie that is configured on the server.
	//
	// The cookie must be 1 to 200 characters in length, and can contain ASCII characters and digits. It cannot contain commas (,), semicolons (;), or spaces. It cannot start with a dollar sign ($).
	//
	// > This parameter is required when **StickySession*	- is set to **on*	- and **StickySessionType*	- is set to **server**.
	//
	// example:
	//
	// testCookie
	Cookie *string `json:"Cookie,omitempty" xml:"Cookie,omitempty"`
	// The timeout period of a cookie.
	//
	// Valid values: **1*	- to **86400**. Unit: seconds.
	//
	// > This parameter is required when **StickySession*	- is set to **on*	- and **StickySessionType*	- is set to **insert**.
	//
	// example:
	//
	// 500
	CookieTimeout *int32 `json:"CookieTimeout,omitempty" xml:"CookieTimeout,omitempty"`
	// The description of the listener.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
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
	// Valid values: **1*	- to **65535**.
	//
	// > This parameter takes effect only if you set **HealthCheck*	- to **on**.
	//
	// example:
	//
	// 8080
	HealthCheckConnectPort *int32 `json:"HealthCheckConnectPort,omitempty" xml:"HealthCheckConnectPort,omitempty"`
	// The domain name that is used for health checks. Valid values:
	//
	// 	- **$_ip**: the private IP address of a backend server. If you specify \\*\\*$_ip **or*	- ignore HealthCheckDomain\\*\\*, CLB uses the private IP addresses of backend servers as the health check domain names.
	//
	// 	- **domain**: The domain name must be 1 to 80 characters in length, and can contain letters, digits, periods (.), and hyphens (-).
	//
	// > The parameter takes effect only if you set **HealthCheck*	- to **on**.
	//
	// example:
	//
	// 172.XX.XX.16
	HealthCheckDomain *string `json:"HealthCheckDomain,omitempty" xml:"HealthCheckDomain,omitempty"`
	// The HTTP status code for a successful health check. Separate multiple HTTP status codes with commas (,).
	//
	// Valid values: **http_2xx**, **http_3xx**, **http_4xx**, and **http_5xx**.
	//
	// >  The parameter takes effect only if you set **HealthCheck*	- to **on**.
	//
	// example:
	//
	// http_2xx,http_3xx
	HealthCheckHttpCode *string `json:"HealthCheckHttpCode,omitempty" xml:"HealthCheckHttpCode,omitempty"`
	// The interval at which health checks are performed.
	//
	// Valid values: **1*	- to **50**. Unit: seconds.
	//
	// >  The parameter takes effect only if you set **HealthCheck*	- to **on**.
	//
	// example:
	//
	// 5
	HealthCheckInterval *int32 `json:"HealthCheckInterval,omitempty" xml:"HealthCheckInterval,omitempty"`
	// The health check method that is used in HTTP health checks. Valid values: **head*	- and **get**.
	//
	// > The parameter takes effect only if you set **HealthCheck*	- to **on**.
	//
	// example:
	//
	// get
	HealthCheckMethod *string `json:"HealthCheckMethod,omitempty" xml:"HealthCheckMethod,omitempty"`
	// The timeout period of a health check response. If a backend ECS instance does not respond within the specified timeout period, the ECS instance fails the health check. This parameter takes effect only if the **HealthCheck*	- parameter is set to **on**.
	//
	// Valid values: **1*	- to **300**. Unit: seconds.
	//
	// example:
	//
	// 3
	HealthCheckTimeout *int32 `json:"HealthCheckTimeout,omitempty" xml:"HealthCheckTimeout,omitempty"`
	// The Uniform Resource Identifier (URI) that you want to use for health checks.
	//
	// The URI must be 1 to 80 characters in length, and can contain letters, digits, and the following characters: - / . % ? # & The URI must start with a forward slash (/) but cannot be a single forward slash (/).
	//
	// > The parameter takes effect only if you set **HealthCheck*	- to **on**.
	//
	// example:
	//
	// /test/index.html
	HealthCheckURI *string `json:"HealthCheckURI,omitempty" xml:"HealthCheckURI,omitempty"`
	// The number of times that an unhealthy backend server must consecutively pass health checks before it is declared healthy. In this case, the health status is changed from **fail*	- to **success**.
	//
	// Valid values: **2*	- to **10**.
	//
	// > The parameter takes effect only if you set **HealthCheck*	- to **on**.
	//
	// example:
	//
	// 4
	HealthyThreshold *int32 `json:"HealthyThreshold,omitempty" xml:"HealthyThreshold,omitempty"`
	// The timeout period of an idle connection. Unit: seconds. Valid values: **1 to 60**. Default value: **15**.
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
	// The CLB instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// lb-bp1qjwo61pqz3ah*****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	OwnerAccount   *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the CLB instance.
	//
	// You can query the region ID from the [Regions and zones](https://help.aliyun.com/document_detail/27585.html) list or by calling the [DescribeRegions](https://help.aliyun.com/document_detail/27584.html) operation.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The timeout period of a request. Unit: seconds. Valid values: **1 to 180**. Default value: **60**.
	//
	// If no response is received from the backend server within the request timeout period, CLB returns an HTTP 504 error code to the client.
	//
	// example:
	//
	// 3
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
	// The method that is used to handle a cookie. Valid values:
	//
	// 	- **insert**: inserts a cookie.
	//
	// CLB inserts a cookie (SERVERID) into the first HTTP or HTTPS response packet that is sent to a client. The next request from the client contains this cookie, and the listener distributes the request to the recorded backend server.
	//
	// 	- **server**: rewrites a cookie.
	//
	// When CLB detects a user-defined cookie, CLB overwrites the original cookie with the user-defined cookie. The next request from the client carries the user-defined cookie, and the listener forwards the request to the recorded backend server.
	//
	// > This parameter is required when **StickySession*	- is set to **on**.
	//
	// example:
	//
	// insert
	StickySessionType *string `json:"StickySessionType,omitempty" xml:"StickySessionType,omitempty"`
	// The number of times that a healthy backend server must consecutively fail health checks before it is declared unhealthy. In this case, the health status is changed from **success*	- to **fail**.
	//
	// Valid values: **2*	- to **10**.
	//
	// > The parameter takes effect only if you set **HealthCheck*	- to **on**.
	//
	// example:
	//
	// 4
	UnhealthyThreshold *int32 `json:"UnhealthyThreshold,omitempty" xml:"UnhealthyThreshold,omitempty"`
	// Specifies whether to use a vServer group. Valid values:
	//
	// 	- **on**
	//
	// 	- **off**
	//
	// example:
	//
	// on
	VServerGroup *string `json:"VServerGroup,omitempty" xml:"VServerGroup,omitempty"`
	// The ID of the vServer group.
	//
	// example:
	//
	// rsp-cige6j*****
	VServerGroupId *string `json:"VServerGroupId,omitempty" xml:"VServerGroupId,omitempty"`
	// Specifies whether to use the `X-Forwarded-For` header to preserve client IP addresses. Valid values:
	//
	// 	- **on*	- (default)
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
	// Specifies whether to use the `SLB-IP` header to retrieve the virtual IP address (VIP) requested by the client. Valid values:
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

func (s SetLoadBalancerHTTPListenerAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s SetLoadBalancerHTTPListenerAttributeRequest) GoString() string {
	return s.String()
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) GetAclId() *string {
	return s.AclId
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) GetAclStatus() *string {
	return s.AclStatus
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) GetAclType() *string {
	return s.AclType
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) GetCookie() *string {
	return s.Cookie
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) GetCookieTimeout() *int32 {
	return s.CookieTimeout
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) GetDescription() *string {
	return s.Description
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) GetGzip() *string {
	return s.Gzip
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) GetHealthCheck() *string {
	return s.HealthCheck
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) GetHealthCheckConnectPort() *int32 {
	return s.HealthCheckConnectPort
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) GetHealthCheckDomain() *string {
	return s.HealthCheckDomain
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) GetHealthCheckHttpCode() *string {
	return s.HealthCheckHttpCode
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) GetHealthCheckInterval() *int32 {
	return s.HealthCheckInterval
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) GetHealthCheckMethod() *string {
	return s.HealthCheckMethod
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) GetHealthCheckTimeout() *int32 {
	return s.HealthCheckTimeout
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) GetHealthCheckURI() *string {
	return s.HealthCheckURI
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) GetHealthyThreshold() *int32 {
	return s.HealthyThreshold
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) GetIdleTimeout() *int32 {
	return s.IdleTimeout
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) GetListenerPort() *int32 {
	return s.ListenerPort
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) GetRequestTimeout() *int32 {
	return s.RequestTimeout
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) GetScheduler() *string {
	return s.Scheduler
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) GetStickySession() *string {
	return s.StickySession
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) GetStickySessionType() *string {
	return s.StickySessionType
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) GetUnhealthyThreshold() *int32 {
	return s.UnhealthyThreshold
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) GetVServerGroup() *string {
	return s.VServerGroup
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) GetVServerGroupId() *string {
	return s.VServerGroupId
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) GetXForwardedFor() *string {
	return s.XForwardedFor
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) GetXForwardedFor_ClientSrcPort() *string {
	return s.XForwardedFor_ClientSrcPort
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) GetXForwardedFor_SLBID() *string {
	return s.XForwardedFor_SLBID
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) GetXForwardedFor_SLBIP() *string {
	return s.XForwardedFor_SLBIP
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) GetXForwardedFor_SLBPORT() *string {
	return s.XForwardedFor_SLBPORT
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) GetXForwardedFor_proto() *string {
	return s.XForwardedFor_proto
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) SetAclId(v string) *SetLoadBalancerHTTPListenerAttributeRequest {
	s.AclId = &v
	return s
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) SetAclStatus(v string) *SetLoadBalancerHTTPListenerAttributeRequest {
	s.AclStatus = &v
	return s
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) SetAclType(v string) *SetLoadBalancerHTTPListenerAttributeRequest {
	s.AclType = &v
	return s
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) SetBandwidth(v int32) *SetLoadBalancerHTTPListenerAttributeRequest {
	s.Bandwidth = &v
	return s
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) SetCookie(v string) *SetLoadBalancerHTTPListenerAttributeRequest {
	s.Cookie = &v
	return s
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) SetCookieTimeout(v int32) *SetLoadBalancerHTTPListenerAttributeRequest {
	s.CookieTimeout = &v
	return s
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) SetDescription(v string) *SetLoadBalancerHTTPListenerAttributeRequest {
	s.Description = &v
	return s
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) SetGzip(v string) *SetLoadBalancerHTTPListenerAttributeRequest {
	s.Gzip = &v
	return s
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) SetHealthCheck(v string) *SetLoadBalancerHTTPListenerAttributeRequest {
	s.HealthCheck = &v
	return s
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) SetHealthCheckConnectPort(v int32) *SetLoadBalancerHTTPListenerAttributeRequest {
	s.HealthCheckConnectPort = &v
	return s
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) SetHealthCheckDomain(v string) *SetLoadBalancerHTTPListenerAttributeRequest {
	s.HealthCheckDomain = &v
	return s
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) SetHealthCheckHttpCode(v string) *SetLoadBalancerHTTPListenerAttributeRequest {
	s.HealthCheckHttpCode = &v
	return s
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) SetHealthCheckInterval(v int32) *SetLoadBalancerHTTPListenerAttributeRequest {
	s.HealthCheckInterval = &v
	return s
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) SetHealthCheckMethod(v string) *SetLoadBalancerHTTPListenerAttributeRequest {
	s.HealthCheckMethod = &v
	return s
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) SetHealthCheckTimeout(v int32) *SetLoadBalancerHTTPListenerAttributeRequest {
	s.HealthCheckTimeout = &v
	return s
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) SetHealthCheckURI(v string) *SetLoadBalancerHTTPListenerAttributeRequest {
	s.HealthCheckURI = &v
	return s
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) SetHealthyThreshold(v int32) *SetLoadBalancerHTTPListenerAttributeRequest {
	s.HealthyThreshold = &v
	return s
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) SetIdleTimeout(v int32) *SetLoadBalancerHTTPListenerAttributeRequest {
	s.IdleTimeout = &v
	return s
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) SetListenerPort(v int32) *SetLoadBalancerHTTPListenerAttributeRequest {
	s.ListenerPort = &v
	return s
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) SetLoadBalancerId(v string) *SetLoadBalancerHTTPListenerAttributeRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) SetOwnerAccount(v string) *SetLoadBalancerHTTPListenerAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) SetOwnerId(v int64) *SetLoadBalancerHTTPListenerAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) SetRegionId(v string) *SetLoadBalancerHTTPListenerAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) SetRequestTimeout(v int32) *SetLoadBalancerHTTPListenerAttributeRequest {
	s.RequestTimeout = &v
	return s
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) SetResourceOwnerAccount(v string) *SetLoadBalancerHTTPListenerAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) SetResourceOwnerId(v int64) *SetLoadBalancerHTTPListenerAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) SetScheduler(v string) *SetLoadBalancerHTTPListenerAttributeRequest {
	s.Scheduler = &v
	return s
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) SetStickySession(v string) *SetLoadBalancerHTTPListenerAttributeRequest {
	s.StickySession = &v
	return s
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) SetStickySessionType(v string) *SetLoadBalancerHTTPListenerAttributeRequest {
	s.StickySessionType = &v
	return s
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) SetUnhealthyThreshold(v int32) *SetLoadBalancerHTTPListenerAttributeRequest {
	s.UnhealthyThreshold = &v
	return s
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) SetVServerGroup(v string) *SetLoadBalancerHTTPListenerAttributeRequest {
	s.VServerGroup = &v
	return s
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) SetVServerGroupId(v string) *SetLoadBalancerHTTPListenerAttributeRequest {
	s.VServerGroupId = &v
	return s
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) SetXForwardedFor(v string) *SetLoadBalancerHTTPListenerAttributeRequest {
	s.XForwardedFor = &v
	return s
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) SetXForwardedFor_ClientSrcPort(v string) *SetLoadBalancerHTTPListenerAttributeRequest {
	s.XForwardedFor_ClientSrcPort = &v
	return s
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) SetXForwardedFor_SLBID(v string) *SetLoadBalancerHTTPListenerAttributeRequest {
	s.XForwardedFor_SLBID = &v
	return s
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) SetXForwardedFor_SLBIP(v string) *SetLoadBalancerHTTPListenerAttributeRequest {
	s.XForwardedFor_SLBIP = &v
	return s
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) SetXForwardedFor_SLBPORT(v string) *SetLoadBalancerHTTPListenerAttributeRequest {
	s.XForwardedFor_SLBPORT = &v
	return s
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) SetXForwardedFor_proto(v string) *SetLoadBalancerHTTPListenerAttributeRequest {
	s.XForwardedFor_proto = &v
	return s
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) Validate() error {
	return dara.Validate(s)
}
