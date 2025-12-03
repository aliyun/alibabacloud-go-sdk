// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLoadBalancerHTTPListenerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAclId(v string) *CreateLoadBalancerHTTPListenerRequest
	GetAclId() *string
	SetAclStatus(v string) *CreateLoadBalancerHTTPListenerRequest
	GetAclStatus() *string
	SetAclType(v string) *CreateLoadBalancerHTTPListenerRequest
	GetAclType() *string
	SetBackendServerPort(v int32) *CreateLoadBalancerHTTPListenerRequest
	GetBackendServerPort() *int32
	SetBandwidth(v int32) *CreateLoadBalancerHTTPListenerRequest
	GetBandwidth() *int32
	SetCookie(v string) *CreateLoadBalancerHTTPListenerRequest
	GetCookie() *string
	SetCookieTimeout(v int32) *CreateLoadBalancerHTTPListenerRequest
	GetCookieTimeout() *int32
	SetDescription(v string) *CreateLoadBalancerHTTPListenerRequest
	GetDescription() *string
	SetForwardPort(v int32) *CreateLoadBalancerHTTPListenerRequest
	GetForwardPort() *int32
	SetGzip(v string) *CreateLoadBalancerHTTPListenerRequest
	GetGzip() *string
	SetHealthCheck(v string) *CreateLoadBalancerHTTPListenerRequest
	GetHealthCheck() *string
	SetHealthCheckConnectPort(v int32) *CreateLoadBalancerHTTPListenerRequest
	GetHealthCheckConnectPort() *int32
	SetHealthCheckDomain(v string) *CreateLoadBalancerHTTPListenerRequest
	GetHealthCheckDomain() *string
	SetHealthCheckHttpCode(v string) *CreateLoadBalancerHTTPListenerRequest
	GetHealthCheckHttpCode() *string
	SetHealthCheckInterval(v int32) *CreateLoadBalancerHTTPListenerRequest
	GetHealthCheckInterval() *int32
	SetHealthCheckMethod(v string) *CreateLoadBalancerHTTPListenerRequest
	GetHealthCheckMethod() *string
	SetHealthCheckTimeout(v int32) *CreateLoadBalancerHTTPListenerRequest
	GetHealthCheckTimeout() *int32
	SetHealthCheckURI(v string) *CreateLoadBalancerHTTPListenerRequest
	GetHealthCheckURI() *string
	SetHealthyThreshold(v int32) *CreateLoadBalancerHTTPListenerRequest
	GetHealthyThreshold() *int32
	SetIdleTimeout(v int32) *CreateLoadBalancerHTTPListenerRequest
	GetIdleTimeout() *int32
	SetListenerForward(v string) *CreateLoadBalancerHTTPListenerRequest
	GetListenerForward() *string
	SetListenerPort(v int32) *CreateLoadBalancerHTTPListenerRequest
	GetListenerPort() *int32
	SetLoadBalancerId(v string) *CreateLoadBalancerHTTPListenerRequest
	GetLoadBalancerId() *string
	SetOwnerAccount(v string) *CreateLoadBalancerHTTPListenerRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateLoadBalancerHTTPListenerRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CreateLoadBalancerHTTPListenerRequest
	GetRegionId() *string
	SetRequestTimeout(v int32) *CreateLoadBalancerHTTPListenerRequest
	GetRequestTimeout() *int32
	SetResourceOwnerAccount(v string) *CreateLoadBalancerHTTPListenerRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateLoadBalancerHTTPListenerRequest
	GetResourceOwnerId() *int64
	SetScheduler(v string) *CreateLoadBalancerHTTPListenerRequest
	GetScheduler() *string
	SetStickySession(v string) *CreateLoadBalancerHTTPListenerRequest
	GetStickySession() *string
	SetStickySessionType(v string) *CreateLoadBalancerHTTPListenerRequest
	GetStickySessionType() *string
	SetTag(v []*CreateLoadBalancerHTTPListenerRequestTag) *CreateLoadBalancerHTTPListenerRequest
	GetTag() []*CreateLoadBalancerHTTPListenerRequestTag
	SetUnhealthyThreshold(v int32) *CreateLoadBalancerHTTPListenerRequest
	GetUnhealthyThreshold() *int32
	SetVServerGroupId(v string) *CreateLoadBalancerHTTPListenerRequest
	GetVServerGroupId() *string
	SetXForwardedFor(v string) *CreateLoadBalancerHTTPListenerRequest
	GetXForwardedFor() *string
	SetXForwardedFor_ClientSrcPort(v string) *CreateLoadBalancerHTTPListenerRequest
	GetXForwardedFor_ClientSrcPort() *string
	SetXForwardedFor_SLBID(v string) *CreateLoadBalancerHTTPListenerRequest
	GetXForwardedFor_SLBID() *string
	SetXForwardedFor_SLBIP(v string) *CreateLoadBalancerHTTPListenerRequest
	GetXForwardedFor_SLBIP() *string
	SetXForwardedFor_SLBPORT(v string) *CreateLoadBalancerHTTPListenerRequest
	GetXForwardedFor_SLBPORT() *string
	SetXForwardedFor_proto(v string) *CreateLoadBalancerHTTPListenerRequest
	GetXForwardedFor_proto() *string
}

type CreateLoadBalancerHTTPListenerRequest struct {
	// The ID of the network access control list (ACL) that is associated with the listener.
	//
	// >  If **AclStatus*	- is set to **on**, this parameter is required.
	//
	// example:
	//
	// acl-uf60jw******
	AclId *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
	// Specifies whether to enable access control. Valid values:
	//
	// 	- **on**: yes
	//
	// 	- **off*	- (default): no
	//
	// example:
	//
	// off
	AclStatus *string `json:"AclStatus,omitempty" xml:"AclStatus,omitempty"`
	// The type of access control. Valid values:
	//
	// 	- **white**: Only requests from IP addresses and CIDR blocks on the whitelist are forwarded by the listener. Your service may be adversely affected if the whitelist is not properly configured. If a whitelist is configured, the listener forwards only requests from IP addresses that are added to the whitelist.
	//
	//     If you configure a whitelist but do not add an IP address to the whitelist, the listener forwards all requests.
	//
	// 	- **black**: Requests from the IP addresses and CIDR blocks on the blacklist are blocked.
	//
	//     If you configure a blacklist but do not add an IP address to the blacklist, the listener forwards all requests.
	//
	// >  When **AclStatus*	- is set to **on**, this parameter takes effect.
	//
	// example:
	//
	// white
	AclType *string `json:"AclType,omitempty" xml:"AclType,omitempty"`
	// The backend port that is used by the CLB instance.
	//
	// Valid values: **1*	- to **65535**.
	//
	// >  If the VServerGroupId parameter is not specified, this parameter is required.
	//
	// example:
	//
	// 80
	BackendServerPort *int32 `json:"BackendServerPort,omitempty" xml:"BackendServerPort,omitempty"`
	// The maximum bandwidth of the listener. Unit: Mbit/s. Valid values:
	//
	// **-1**: specifies that the bandwidth of the listener is unlimited.
	//
	// example:
	//
	// -1
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The cookie configured for the server.
	//
	// The cookie must be 1 to 200 characters in length, and can contain only ASCII letters and digits. It cannot contain commas (,), semicolons (;), space characters, or start with a dollar sign ($).
	//
	// >  This parameter is required when the **StickySession*	- parameter is set to **on*	- and the **StickySessionType*	- parameter is set to **server**.
	//
	// example:
	//
	// testCookie
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
	// The name must be 1 to 256 characters in length, and can contain letters, digits, hyphens (-), forward slashes (/), periods (.), and underscores (_).
	//
	// example:
	//
	// HTTP_443
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The listening port that is used to redirect HTTP requests to HTTPS.
	//
	// example:
	//
	// 443
	ForwardPort *int32 `json:"ForwardPort,omitempty" xml:"ForwardPort,omitempty"`
	// Specifies whether to enable `GZIP` compression to compress specific types of files. Valid values:
	//
	// 	- **on*	- (default)
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
	// The backend port that is used for health checks.
	//
	// Valid values: **1*	- to **65535**.
	//
	// >  This parameter takes effect only if the **HealthCheck*	- parameter is set to **on**.
	//
	// example:
	//
	// 80
	HealthCheckConnectPort *int32 `json:"HealthCheckConnectPort,omitempty" xml:"HealthCheckConnectPort,omitempty"`
	// The domain name that is used for health checks. Valid values:
	//
	// 	- **$_ip**: the private IP address of a backend server. If an IP address is specified, or this parameter is not specified, the CLB instance uses the private IP address of each backend server as the domain name for health checks.
	//
	// 	- **domain**: The domain name must be 1 to 80 characters in length, and can contain letters, digits, periods (.), and hyphens (-).
	//
	// >  This parameter takes effect only if the **HealthCheck*	- parameter is set to **on**.
	//
	// example:
	//
	// 172.16.0.0/12
	HealthCheckDomain *string `json:"HealthCheckDomain,omitempty" xml:"HealthCheckDomain,omitempty"`
	// The HTTP status code for a successful health check. Separate multiple HTTP status codes with commas (,).
	//
	// Valid values: **http_2xx*	- (default), **http_3xx**, **http_4xx**, and **http_5xx**.
	//
	// >  This parameter takes effect only if the **HealthCheck*	- parameter is set to **on**.
	//
	// example:
	//
	// http_2xx
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
	// The timeout period of a health check response. If a backend ECS instance does not respond within the specified timeout period, the ECS instance fails the health check. Unit: seconds
	//
	// Valid values: **1*	- to **300**.
	//
	// >  This parameter takes effect only if the **HealthCheck*	- parameter is set to **on**.
	//
	// example:
	//
	// 3
	HealthCheckTimeout *int32 `json:"HealthCheckTimeout,omitempty" xml:"HealthCheckTimeout,omitempty"`
	// The URI that is used for health checks.
	//
	// The URI must be 1 to 80 characters in length, and can contain letters, digits, hyphens (-), forward slashes (/), periods (.), percent signs (%), question marks (?), number signs (#), and ampersands (&). The URI must start with a forward slash (/) but cannot be a single forward slash (/).
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
	// The timeout period of an idle connection. Unit: seconds.
	//
	// Default value: **15**. Valid values: **1*	- to **60**.
	//
	// If no request is received within the specified timeout period, SLB closes the connection. When a request is received, SLB establishes a new connection.
	//
	// example:
	//
	// 3
	IdleTimeout *int32 `json:"IdleTimeout,omitempty" xml:"IdleTimeout,omitempty"`
	// Specifies whether to enable HTTP-to-HTTPS redirection. Valid values:
	//
	// 	- **on**: yes
	//
	// 	- **off*	- (default): no
	//
	// example:
	//
	// off
	ListenerForward *string `json:"ListenerForward,omitempty" xml:"ListenerForward,omitempty"`
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
	// lb-bp1c9vixxjh92q83tw*****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	OwnerAccount   *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the CLB instance is deployed.
	//
	// You can query the region ID from the [Regions and zones](https://help.aliyun.com/document_detail/40654.html) list or by calling the [DescribeRegions](https://help.aliyun.com/document_detail/27584.html) operation.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The timeout period of a request. Unit: seconds.
	//
	// Default value: **60**. Valid values: **1*	- to **180**.
	//
	// If no response is received from the backend server within the specified timeout period, CLB sends an `HTTP 504` error code to the client.
	//
	// example:
	//
	// 6
	RequestTimeout       *int32  `json:"RequestTimeout,omitempty" xml:"RequestTimeout,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The scheduling algorithm. Valid values:
	//
	// 	- **wrr*	- (default): Backend servers with higher weights receive more requests than those with lower weights.
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
	// 	- **off*	- (default)
	//
	// example:
	//
	// off
	StickySession *string `json:"StickySession,omitempty" xml:"StickySession,omitempty"`
	// The method that is used to handle cookies. Valid values:
	//
	// 	- **insert**: inserts a cookie.
	//
	//     The first time a client accesses CLB, CLB inserts a cookie into the response packet. Subsequent requests from the client that carry the cookie are distributed to the same backend server as the first request.
	//
	// 	- **server**: rewrites the original cookie.
	//
	//     CLB rewrites the custom cookies in requests from a client. Subsequent requests from the client that carry the new cookie are forwarded to the same backend server as the first request.
	//
	// >  This parameter is required if the **StickySession*	- parameter is set to **on**.
	//
	// example:
	//
	// insert
	StickySessionType *string `json:"StickySessionType,omitempty" xml:"StickySessionType,omitempty"`
	// The tags.
	Tag []*CreateLoadBalancerHTTPListenerRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
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
	// The ID of the vServer group.
	//
	// example:
	//
	// rsp-cige6j*****
	VServerGroupId *string `json:"VServerGroupId,omitempty" xml:"VServerGroupId,omitempty"`
	// Specifies whether to use the `X-Forwarded-For` header to retrieve client IP addresses. Valid values:
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
	// 	- **off*	- (default)
	//
	// example:
	//
	// on
	XForwardedFor_SLBID *string `json:"XForwardedFor_SLBID,omitempty" xml:"XForwardedFor_SLBID,omitempty"`
	// Specifies whether to use the `SLB-IP` header to retrieve the virtual IP address of the client. Valid values:
	//
	// 	- **on**
	//
	// 	- **off*	- (default)
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
	// 	- **off*	- (default)
	//
	// example:
	//
	// on
	XForwardedFor_proto *string `json:"XForwardedFor_proto,omitempty" xml:"XForwardedFor_proto,omitempty"`
}

func (s CreateLoadBalancerHTTPListenerRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateLoadBalancerHTTPListenerRequest) GoString() string {
	return s.String()
}

func (s *CreateLoadBalancerHTTPListenerRequest) GetAclId() *string {
	return s.AclId
}

func (s *CreateLoadBalancerHTTPListenerRequest) GetAclStatus() *string {
	return s.AclStatus
}

func (s *CreateLoadBalancerHTTPListenerRequest) GetAclType() *string {
	return s.AclType
}

func (s *CreateLoadBalancerHTTPListenerRequest) GetBackendServerPort() *int32 {
	return s.BackendServerPort
}

func (s *CreateLoadBalancerHTTPListenerRequest) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *CreateLoadBalancerHTTPListenerRequest) GetCookie() *string {
	return s.Cookie
}

func (s *CreateLoadBalancerHTTPListenerRequest) GetCookieTimeout() *int32 {
	return s.CookieTimeout
}

func (s *CreateLoadBalancerHTTPListenerRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateLoadBalancerHTTPListenerRequest) GetForwardPort() *int32 {
	return s.ForwardPort
}

func (s *CreateLoadBalancerHTTPListenerRequest) GetGzip() *string {
	return s.Gzip
}

func (s *CreateLoadBalancerHTTPListenerRequest) GetHealthCheck() *string {
	return s.HealthCheck
}

func (s *CreateLoadBalancerHTTPListenerRequest) GetHealthCheckConnectPort() *int32 {
	return s.HealthCheckConnectPort
}

func (s *CreateLoadBalancerHTTPListenerRequest) GetHealthCheckDomain() *string {
	return s.HealthCheckDomain
}

func (s *CreateLoadBalancerHTTPListenerRequest) GetHealthCheckHttpCode() *string {
	return s.HealthCheckHttpCode
}

func (s *CreateLoadBalancerHTTPListenerRequest) GetHealthCheckInterval() *int32 {
	return s.HealthCheckInterval
}

func (s *CreateLoadBalancerHTTPListenerRequest) GetHealthCheckMethod() *string {
	return s.HealthCheckMethod
}

func (s *CreateLoadBalancerHTTPListenerRequest) GetHealthCheckTimeout() *int32 {
	return s.HealthCheckTimeout
}

func (s *CreateLoadBalancerHTTPListenerRequest) GetHealthCheckURI() *string {
	return s.HealthCheckURI
}

func (s *CreateLoadBalancerHTTPListenerRequest) GetHealthyThreshold() *int32 {
	return s.HealthyThreshold
}

func (s *CreateLoadBalancerHTTPListenerRequest) GetIdleTimeout() *int32 {
	return s.IdleTimeout
}

func (s *CreateLoadBalancerHTTPListenerRequest) GetListenerForward() *string {
	return s.ListenerForward
}

func (s *CreateLoadBalancerHTTPListenerRequest) GetListenerPort() *int32 {
	return s.ListenerPort
}

func (s *CreateLoadBalancerHTTPListenerRequest) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *CreateLoadBalancerHTTPListenerRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateLoadBalancerHTTPListenerRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateLoadBalancerHTTPListenerRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateLoadBalancerHTTPListenerRequest) GetRequestTimeout() *int32 {
	return s.RequestTimeout
}

func (s *CreateLoadBalancerHTTPListenerRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateLoadBalancerHTTPListenerRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateLoadBalancerHTTPListenerRequest) GetScheduler() *string {
	return s.Scheduler
}

func (s *CreateLoadBalancerHTTPListenerRequest) GetStickySession() *string {
	return s.StickySession
}

func (s *CreateLoadBalancerHTTPListenerRequest) GetStickySessionType() *string {
	return s.StickySessionType
}

func (s *CreateLoadBalancerHTTPListenerRequest) GetTag() []*CreateLoadBalancerHTTPListenerRequestTag {
	return s.Tag
}

func (s *CreateLoadBalancerHTTPListenerRequest) GetUnhealthyThreshold() *int32 {
	return s.UnhealthyThreshold
}

func (s *CreateLoadBalancerHTTPListenerRequest) GetVServerGroupId() *string {
	return s.VServerGroupId
}

func (s *CreateLoadBalancerHTTPListenerRequest) GetXForwardedFor() *string {
	return s.XForwardedFor
}

func (s *CreateLoadBalancerHTTPListenerRequest) GetXForwardedFor_ClientSrcPort() *string {
	return s.XForwardedFor_ClientSrcPort
}

func (s *CreateLoadBalancerHTTPListenerRequest) GetXForwardedFor_SLBID() *string {
	return s.XForwardedFor_SLBID
}

func (s *CreateLoadBalancerHTTPListenerRequest) GetXForwardedFor_SLBIP() *string {
	return s.XForwardedFor_SLBIP
}

func (s *CreateLoadBalancerHTTPListenerRequest) GetXForwardedFor_SLBPORT() *string {
	return s.XForwardedFor_SLBPORT
}

func (s *CreateLoadBalancerHTTPListenerRequest) GetXForwardedFor_proto() *string {
	return s.XForwardedFor_proto
}

func (s *CreateLoadBalancerHTTPListenerRequest) SetAclId(v string) *CreateLoadBalancerHTTPListenerRequest {
	s.AclId = &v
	return s
}

func (s *CreateLoadBalancerHTTPListenerRequest) SetAclStatus(v string) *CreateLoadBalancerHTTPListenerRequest {
	s.AclStatus = &v
	return s
}

func (s *CreateLoadBalancerHTTPListenerRequest) SetAclType(v string) *CreateLoadBalancerHTTPListenerRequest {
	s.AclType = &v
	return s
}

func (s *CreateLoadBalancerHTTPListenerRequest) SetBackendServerPort(v int32) *CreateLoadBalancerHTTPListenerRequest {
	s.BackendServerPort = &v
	return s
}

func (s *CreateLoadBalancerHTTPListenerRequest) SetBandwidth(v int32) *CreateLoadBalancerHTTPListenerRequest {
	s.Bandwidth = &v
	return s
}

func (s *CreateLoadBalancerHTTPListenerRequest) SetCookie(v string) *CreateLoadBalancerHTTPListenerRequest {
	s.Cookie = &v
	return s
}

func (s *CreateLoadBalancerHTTPListenerRequest) SetCookieTimeout(v int32) *CreateLoadBalancerHTTPListenerRequest {
	s.CookieTimeout = &v
	return s
}

func (s *CreateLoadBalancerHTTPListenerRequest) SetDescription(v string) *CreateLoadBalancerHTTPListenerRequest {
	s.Description = &v
	return s
}

func (s *CreateLoadBalancerHTTPListenerRequest) SetForwardPort(v int32) *CreateLoadBalancerHTTPListenerRequest {
	s.ForwardPort = &v
	return s
}

func (s *CreateLoadBalancerHTTPListenerRequest) SetGzip(v string) *CreateLoadBalancerHTTPListenerRequest {
	s.Gzip = &v
	return s
}

func (s *CreateLoadBalancerHTTPListenerRequest) SetHealthCheck(v string) *CreateLoadBalancerHTTPListenerRequest {
	s.HealthCheck = &v
	return s
}

func (s *CreateLoadBalancerHTTPListenerRequest) SetHealthCheckConnectPort(v int32) *CreateLoadBalancerHTTPListenerRequest {
	s.HealthCheckConnectPort = &v
	return s
}

func (s *CreateLoadBalancerHTTPListenerRequest) SetHealthCheckDomain(v string) *CreateLoadBalancerHTTPListenerRequest {
	s.HealthCheckDomain = &v
	return s
}

func (s *CreateLoadBalancerHTTPListenerRequest) SetHealthCheckHttpCode(v string) *CreateLoadBalancerHTTPListenerRequest {
	s.HealthCheckHttpCode = &v
	return s
}

func (s *CreateLoadBalancerHTTPListenerRequest) SetHealthCheckInterval(v int32) *CreateLoadBalancerHTTPListenerRequest {
	s.HealthCheckInterval = &v
	return s
}

func (s *CreateLoadBalancerHTTPListenerRequest) SetHealthCheckMethod(v string) *CreateLoadBalancerHTTPListenerRequest {
	s.HealthCheckMethod = &v
	return s
}

func (s *CreateLoadBalancerHTTPListenerRequest) SetHealthCheckTimeout(v int32) *CreateLoadBalancerHTTPListenerRequest {
	s.HealthCheckTimeout = &v
	return s
}

func (s *CreateLoadBalancerHTTPListenerRequest) SetHealthCheckURI(v string) *CreateLoadBalancerHTTPListenerRequest {
	s.HealthCheckURI = &v
	return s
}

func (s *CreateLoadBalancerHTTPListenerRequest) SetHealthyThreshold(v int32) *CreateLoadBalancerHTTPListenerRequest {
	s.HealthyThreshold = &v
	return s
}

func (s *CreateLoadBalancerHTTPListenerRequest) SetIdleTimeout(v int32) *CreateLoadBalancerHTTPListenerRequest {
	s.IdleTimeout = &v
	return s
}

func (s *CreateLoadBalancerHTTPListenerRequest) SetListenerForward(v string) *CreateLoadBalancerHTTPListenerRequest {
	s.ListenerForward = &v
	return s
}

func (s *CreateLoadBalancerHTTPListenerRequest) SetListenerPort(v int32) *CreateLoadBalancerHTTPListenerRequest {
	s.ListenerPort = &v
	return s
}

func (s *CreateLoadBalancerHTTPListenerRequest) SetLoadBalancerId(v string) *CreateLoadBalancerHTTPListenerRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *CreateLoadBalancerHTTPListenerRequest) SetOwnerAccount(v string) *CreateLoadBalancerHTTPListenerRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateLoadBalancerHTTPListenerRequest) SetOwnerId(v int64) *CreateLoadBalancerHTTPListenerRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateLoadBalancerHTTPListenerRequest) SetRegionId(v string) *CreateLoadBalancerHTTPListenerRequest {
	s.RegionId = &v
	return s
}

func (s *CreateLoadBalancerHTTPListenerRequest) SetRequestTimeout(v int32) *CreateLoadBalancerHTTPListenerRequest {
	s.RequestTimeout = &v
	return s
}

func (s *CreateLoadBalancerHTTPListenerRequest) SetResourceOwnerAccount(v string) *CreateLoadBalancerHTTPListenerRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateLoadBalancerHTTPListenerRequest) SetResourceOwnerId(v int64) *CreateLoadBalancerHTTPListenerRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateLoadBalancerHTTPListenerRequest) SetScheduler(v string) *CreateLoadBalancerHTTPListenerRequest {
	s.Scheduler = &v
	return s
}

func (s *CreateLoadBalancerHTTPListenerRequest) SetStickySession(v string) *CreateLoadBalancerHTTPListenerRequest {
	s.StickySession = &v
	return s
}

func (s *CreateLoadBalancerHTTPListenerRequest) SetStickySessionType(v string) *CreateLoadBalancerHTTPListenerRequest {
	s.StickySessionType = &v
	return s
}

func (s *CreateLoadBalancerHTTPListenerRequest) SetTag(v []*CreateLoadBalancerHTTPListenerRequestTag) *CreateLoadBalancerHTTPListenerRequest {
	s.Tag = v
	return s
}

func (s *CreateLoadBalancerHTTPListenerRequest) SetUnhealthyThreshold(v int32) *CreateLoadBalancerHTTPListenerRequest {
	s.UnhealthyThreshold = &v
	return s
}

func (s *CreateLoadBalancerHTTPListenerRequest) SetVServerGroupId(v string) *CreateLoadBalancerHTTPListenerRequest {
	s.VServerGroupId = &v
	return s
}

func (s *CreateLoadBalancerHTTPListenerRequest) SetXForwardedFor(v string) *CreateLoadBalancerHTTPListenerRequest {
	s.XForwardedFor = &v
	return s
}

func (s *CreateLoadBalancerHTTPListenerRequest) SetXForwardedFor_ClientSrcPort(v string) *CreateLoadBalancerHTTPListenerRequest {
	s.XForwardedFor_ClientSrcPort = &v
	return s
}

func (s *CreateLoadBalancerHTTPListenerRequest) SetXForwardedFor_SLBID(v string) *CreateLoadBalancerHTTPListenerRequest {
	s.XForwardedFor_SLBID = &v
	return s
}

func (s *CreateLoadBalancerHTTPListenerRequest) SetXForwardedFor_SLBIP(v string) *CreateLoadBalancerHTTPListenerRequest {
	s.XForwardedFor_SLBIP = &v
	return s
}

func (s *CreateLoadBalancerHTTPListenerRequest) SetXForwardedFor_SLBPORT(v string) *CreateLoadBalancerHTTPListenerRequest {
	s.XForwardedFor_SLBPORT = &v
	return s
}

func (s *CreateLoadBalancerHTTPListenerRequest) SetXForwardedFor_proto(v string) *CreateLoadBalancerHTTPListenerRequest {
	s.XForwardedFor_proto = &v
	return s
}

func (s *CreateLoadBalancerHTTPListenerRequest) Validate() error {
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

type CreateLoadBalancerHTTPListenerRequestTag struct {
	// The tag key of the bastion host. Valid values of N: **1 to 20**. The tag key cannot be an empty string.
	//
	// The tag key can be at most 64 characters in length, and cannot contain `http://` or `https://`. It must not start with `aliyun` or `acs:`.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. Valid values of N: **1 to 20**. The tag value can be an empty string.
	//
	// The tag value can be up to 128 characters in length and cannot start with `acs:` or `aliyun`. The tag value cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateLoadBalancerHTTPListenerRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateLoadBalancerHTTPListenerRequestTag) GoString() string {
	return s.String()
}

func (s *CreateLoadBalancerHTTPListenerRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateLoadBalancerHTTPListenerRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateLoadBalancerHTTPListenerRequestTag) SetKey(v string) *CreateLoadBalancerHTTPListenerRequestTag {
	s.Key = &v
	return s
}

func (s *CreateLoadBalancerHTTPListenerRequestTag) SetValue(v string) *CreateLoadBalancerHTTPListenerRequestTag {
	s.Value = &v
	return s
}

func (s *CreateLoadBalancerHTTPListenerRequestTag) Validate() error {
	return dara.Validate(s)
}
