// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLoadBalancerHTTPListenerAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAclId(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBody
	GetAclId() *string
	SetAclIds(v *DescribeLoadBalancerHTTPListenerAttributeResponseBodyAclIds) *DescribeLoadBalancerHTTPListenerAttributeResponseBody
	GetAclIds() *DescribeLoadBalancerHTTPListenerAttributeResponseBodyAclIds
	SetAclStatus(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBody
	GetAclStatus() *string
	SetAclType(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBody
	GetAclType() *string
	SetBackendServerPort(v int32) *DescribeLoadBalancerHTTPListenerAttributeResponseBody
	GetBackendServerPort() *int32
	SetBandwidth(v int32) *DescribeLoadBalancerHTTPListenerAttributeResponseBody
	GetBandwidth() *int32
	SetCookie(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBody
	GetCookie() *string
	SetCookieTimeout(v int32) *DescribeLoadBalancerHTTPListenerAttributeResponseBody
	GetCookieTimeout() *int32
	SetDescription(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBody
	GetDescription() *string
	SetForwardPort(v int32) *DescribeLoadBalancerHTTPListenerAttributeResponseBody
	GetForwardPort() *int32
	SetGzip(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBody
	GetGzip() *string
	SetHealthCheck(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBody
	GetHealthCheck() *string
	SetHealthCheckConnectPort(v int32) *DescribeLoadBalancerHTTPListenerAttributeResponseBody
	GetHealthCheckConnectPort() *int32
	SetHealthCheckDomain(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBody
	GetHealthCheckDomain() *string
	SetHealthCheckHttpCode(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBody
	GetHealthCheckHttpCode() *string
	SetHealthCheckInterval(v int32) *DescribeLoadBalancerHTTPListenerAttributeResponseBody
	GetHealthCheckInterval() *int32
	SetHealthCheckMethod(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBody
	GetHealthCheckMethod() *string
	SetHealthCheckTimeout(v int32) *DescribeLoadBalancerHTTPListenerAttributeResponseBody
	GetHealthCheckTimeout() *int32
	SetHealthCheckURI(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBody
	GetHealthCheckURI() *string
	SetHealthyThreshold(v int32) *DescribeLoadBalancerHTTPListenerAttributeResponseBody
	GetHealthyThreshold() *int32
	SetIdleTimeout(v int32) *DescribeLoadBalancerHTTPListenerAttributeResponseBody
	GetIdleTimeout() *int32
	SetListenerForward(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBody
	GetListenerForward() *string
	SetListenerPort(v int32) *DescribeLoadBalancerHTTPListenerAttributeResponseBody
	GetListenerPort() *int32
	SetLoadBalancerId(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBody
	GetLoadBalancerId() *string
	SetRequestId(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBody
	GetRequestId() *string
	SetRequestTimeout(v int32) *DescribeLoadBalancerHTTPListenerAttributeResponseBody
	GetRequestTimeout() *int32
	SetRules(v *DescribeLoadBalancerHTTPListenerAttributeResponseBodyRules) *DescribeLoadBalancerHTTPListenerAttributeResponseBody
	GetRules() *DescribeLoadBalancerHTTPListenerAttributeResponseBodyRules
	SetScheduler(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBody
	GetScheduler() *string
	SetSecurityStatus(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBody
	GetSecurityStatus() *string
	SetStatus(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBody
	GetStatus() *string
	SetStickySession(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBody
	GetStickySession() *string
	SetStickySessionType(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBody
	GetStickySessionType() *string
	SetTags(v *DescribeLoadBalancerHTTPListenerAttributeResponseBodyTags) *DescribeLoadBalancerHTTPListenerAttributeResponseBody
	GetTags() *DescribeLoadBalancerHTTPListenerAttributeResponseBodyTags
	SetUnhealthyThreshold(v int32) *DescribeLoadBalancerHTTPListenerAttributeResponseBody
	GetUnhealthyThreshold() *int32
	SetVServerGroupId(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBody
	GetVServerGroupId() *string
	SetXForwardedFor(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBody
	GetXForwardedFor() *string
	SetXForwardedFor_ClientSrcPort(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBody
	GetXForwardedFor_ClientSrcPort() *string
	SetXForwardedFor_SLBID(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBody
	GetXForwardedFor_SLBID() *string
	SetXForwardedFor_SLBIP(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBody
	GetXForwardedFor_SLBIP() *string
	SetXForwardedFor_SLBPORT(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBody
	GetXForwardedFor_SLBPORT() *string
	SetXForwardedFor_proto(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBody
	GetXForwardedFor_proto() *string
}

type DescribeLoadBalancerHTTPListenerAttributeResponseBody struct {
	// The ID of the network ACL that is associated with a listener.
	//
	// > This parameter is returned when **AclStatus*	- is set to **on**.
	//
	// example:
	//
	// acl-uf60jw******
	AclId *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
	// The ID of the access control list (ACL).
	AclIds *DescribeLoadBalancerHTTPListenerAttributeResponseBodyAclIds `json:"AclIds,omitempty" xml:"AclIds,omitempty" type:"Struct"`
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
	// The type of the ACL. Valid values:
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
	// 80
	BackendServerPort *int32 `json:"BackendServerPort,omitempty" xml:"BackendServerPort,omitempty"`
	// The maximum bandwidth of the listener. Unit: Mbit/s.
	//
	// 	- **-1**: If -1 is returned, it indicates that the bandwidth of the listener is unlimited.
	//
	// 	- **1 to 5120**: If a value from 1 to 5120 is returned, the value indicates the maximum bandwidth of the listener. The sum of the maximum bandwidth of all listeners added to a CLB instance does not exceed the maximum bandwidth of the CLB instance.
	//
	// example:
	//
	// -1
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The cookie that is configured on the server.
	//
	// example:
	//
	// testCookie
	Cookie *string `json:"Cookie,omitempty" xml:"Cookie,omitempty"`
	// The timeout period of a cookie. Unit: seconds.
	//
	// example:
	//
	// 500
	CookieTimeout *int32 `json:"CookieTimeout,omitempty" xml:"CookieTimeout,omitempty"`
	// The name of the listener.
	//
	// example:
	//
	// HTTP_80
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The listener port that is used to redirect HTTP requests to HTTPS.
	//
	// >  If the **ListenerForward*	- parameter is set to **off**, this parameter is not displayed.
	//
	// example:
	//
	// 80
	ForwardPort *int32 `json:"ForwardPort,omitempty" xml:"ForwardPort,omitempty"`
	// Indicates whether `Gzip` compression is enabled to compress specific types of files. Valid values:
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
	// > This parameter takes effect only when **HealthCheck*	- is set to **on**.
	//
	// example:
	//
	// 8080
	HealthCheckConnectPort *int32 `json:"HealthCheckConnectPort,omitempty" xml:"HealthCheckConnectPort,omitempty"`
	// The domain name that you want to use for health checks.
	//
	// example:
	//
	// www.domain.com
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
	// > This parameter is returned when **HealthCheck*	- is set to **on**.
	//
	// example:
	//
	// get
	HealthCheckMethod *string `json:"HealthCheckMethod,omitempty" xml:"HealthCheckMethod,omitempty"`
	// The timeout period of each health check. Unit: seconds.
	//
	// example:
	//
	// 3
	HealthCheckTimeout *int32 `json:"HealthCheckTimeout,omitempty" xml:"HealthCheckTimeout,omitempty"`
	// The URL path that is used for health checks.
	//
	// The URI must be 1 to 80 characters in length, and can contain only digits, letters, hyphens (-), forward slashes (/), periods (.), percent signs (%), question marks (?), number signs (#), and ampersands (&). The URI must start with a forward slash (/) but cannot be a single forward slash (/).
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
	// The timeout period of an idle connection. Unit: seconds.
	//
	// Default value: **15**. Valid values: **1 to 60**.
	//
	// If no request is received within the specified timeout period, CLB closes the connection. When a request is received, CLB establishes a new connection.
	//
	// example:
	//
	// 2
	IdleTimeout *int32 `json:"IdleTimeout,omitempty" xml:"IdleTimeout,omitempty"`
	// Indicates whether HTTP-to-HTTPS redirection is enabled. Valid values:
	//
	// 	- **on**
	//
	// 	- **off**
	//
	// example:
	//
	// on
	ListenerForward *string `json:"ListenerForward,omitempty" xml:"ListenerForward,omitempty"`
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
	// lb-bp1uaunez0uho0zf0****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 365F4154-92F6-4AE4-92F8-7FF34B540710
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The timeout period of a request. Unit: seconds.
	//
	// Default value: **60**. Valid values: **1 to 180**.
	//
	// If no response is received from a backend server within the specified timeout period, CLB returns the HTTP 504 status code to the client.
	//
	// example:
	//
	// 34
	RequestTimeout *int32 `json:"RequestTimeout,omitempty" xml:"RequestTimeout,omitempty"`
	// The list of forwarding rules.
	Rules *DescribeLoadBalancerHTTPListenerAttributeResponseBodyRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Struct"`
	// The routing algorithm. Valid values:
	//
	// 	- **wrr**: Backend servers that have higher weights receive more requests than backend servers that have lower weights.
	//
	// 	- \\*\\	- rr\\*\\*: Requests are sequentially distributed to backend servers.
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
	//     CLB inserts a session cookie (SERVERID) into the first HTTP or HTTPS response that is sent to a client. Subsequent requests to CLB carry this cookie, and CLB determines the destination servers of the requests based on the cookies.
	//
	// 	- **server**: rewrites a cookie.
	//
	//     When CLB detects a user-defined cookie, it overwrites the original cookie with the user-defined cookie. The next request from the client carries the user-defined cookie, and the listener forwards this request to the recorded backend server.
	//
	// > This parameter is required when **StickySession*	- is set to **on**.
	//
	// example:
	//
	// insert
	StickySessionType *string `json:"StickySessionType,omitempty" xml:"StickySessionType,omitempty"`
	// The tags.
	Tags *DescribeLoadBalancerHTTPListenerAttributeResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
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
	// rsp-cige6j****
	VServerGroupId *string `json:"VServerGroupId,omitempty" xml:"VServerGroupId,omitempty"`
	// Indicates whether the `X-Forwarded-For` header is used to preserve the real IP address of the client. Valid values:
	//
	// 	- **on**
	//
	// 	- **off**
	//
	// example:
	//
	// on
	XForwardedFor *string `json:"XForwardedFor,omitempty" xml:"XForwardedFor,omitempty"`
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
	// Indicates whether the `SLB-ID` header is used to retrieve the ID of the CLB instance. Valid values:
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
	// Indicates whether the `XForwardedFor_SLBPORT` header is used to retrieve the listener port of the CLB instance. Valid values:
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

func (s DescribeLoadBalancerHTTPListenerAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLoadBalancerHTTPListenerAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) GetAclId() *string {
	return s.AclId
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) GetAclIds() *DescribeLoadBalancerHTTPListenerAttributeResponseBodyAclIds {
	return s.AclIds
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) GetAclStatus() *string {
	return s.AclStatus
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) GetAclType() *string {
	return s.AclType
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) GetBackendServerPort() *int32 {
	return s.BackendServerPort
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) GetCookie() *string {
	return s.Cookie
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) GetCookieTimeout() *int32 {
	return s.CookieTimeout
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) GetDescription() *string {
	return s.Description
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) GetForwardPort() *int32 {
	return s.ForwardPort
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) GetGzip() *string {
	return s.Gzip
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) GetHealthCheck() *string {
	return s.HealthCheck
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) GetHealthCheckConnectPort() *int32 {
	return s.HealthCheckConnectPort
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) GetHealthCheckDomain() *string {
	return s.HealthCheckDomain
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) GetHealthCheckHttpCode() *string {
	return s.HealthCheckHttpCode
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) GetHealthCheckInterval() *int32 {
	return s.HealthCheckInterval
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) GetHealthCheckMethod() *string {
	return s.HealthCheckMethod
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) GetHealthCheckTimeout() *int32 {
	return s.HealthCheckTimeout
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) GetHealthCheckURI() *string {
	return s.HealthCheckURI
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) GetHealthyThreshold() *int32 {
	return s.HealthyThreshold
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) GetIdleTimeout() *int32 {
	return s.IdleTimeout
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) GetListenerForward() *string {
	return s.ListenerForward
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) GetListenerPort() *int32 {
	return s.ListenerPort
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) GetRequestTimeout() *int32 {
	return s.RequestTimeout
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) GetRules() *DescribeLoadBalancerHTTPListenerAttributeResponseBodyRules {
	return s.Rules
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) GetScheduler() *string {
	return s.Scheduler
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) GetSecurityStatus() *string {
	return s.SecurityStatus
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) GetStickySession() *string {
	return s.StickySession
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) GetStickySessionType() *string {
	return s.StickySessionType
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) GetTags() *DescribeLoadBalancerHTTPListenerAttributeResponseBodyTags {
	return s.Tags
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) GetUnhealthyThreshold() *int32 {
	return s.UnhealthyThreshold
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) GetVServerGroupId() *string {
	return s.VServerGroupId
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) GetXForwardedFor() *string {
	return s.XForwardedFor
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) GetXForwardedFor_ClientSrcPort() *string {
	return s.XForwardedFor_ClientSrcPort
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) GetXForwardedFor_SLBID() *string {
	return s.XForwardedFor_SLBID
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) GetXForwardedFor_SLBIP() *string {
	return s.XForwardedFor_SLBIP
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) GetXForwardedFor_SLBPORT() *string {
	return s.XForwardedFor_SLBPORT
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) GetXForwardedFor_proto() *string {
	return s.XForwardedFor_proto
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) SetAclId(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBody {
	s.AclId = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) SetAclIds(v *DescribeLoadBalancerHTTPListenerAttributeResponseBodyAclIds) *DescribeLoadBalancerHTTPListenerAttributeResponseBody {
	s.AclIds = v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) SetAclStatus(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBody {
	s.AclStatus = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) SetAclType(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBody {
	s.AclType = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) SetBackendServerPort(v int32) *DescribeLoadBalancerHTTPListenerAttributeResponseBody {
	s.BackendServerPort = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) SetBandwidth(v int32) *DescribeLoadBalancerHTTPListenerAttributeResponseBody {
	s.Bandwidth = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) SetCookie(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBody {
	s.Cookie = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) SetCookieTimeout(v int32) *DescribeLoadBalancerHTTPListenerAttributeResponseBody {
	s.CookieTimeout = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) SetDescription(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) SetForwardPort(v int32) *DescribeLoadBalancerHTTPListenerAttributeResponseBody {
	s.ForwardPort = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) SetGzip(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBody {
	s.Gzip = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) SetHealthCheck(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBody {
	s.HealthCheck = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) SetHealthCheckConnectPort(v int32) *DescribeLoadBalancerHTTPListenerAttributeResponseBody {
	s.HealthCheckConnectPort = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) SetHealthCheckDomain(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBody {
	s.HealthCheckDomain = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) SetHealthCheckHttpCode(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBody {
	s.HealthCheckHttpCode = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) SetHealthCheckInterval(v int32) *DescribeLoadBalancerHTTPListenerAttributeResponseBody {
	s.HealthCheckInterval = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) SetHealthCheckMethod(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBody {
	s.HealthCheckMethod = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) SetHealthCheckTimeout(v int32) *DescribeLoadBalancerHTTPListenerAttributeResponseBody {
	s.HealthCheckTimeout = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) SetHealthCheckURI(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBody {
	s.HealthCheckURI = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) SetHealthyThreshold(v int32) *DescribeLoadBalancerHTTPListenerAttributeResponseBody {
	s.HealthyThreshold = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) SetIdleTimeout(v int32) *DescribeLoadBalancerHTTPListenerAttributeResponseBody {
	s.IdleTimeout = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) SetListenerForward(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBody {
	s.ListenerForward = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) SetListenerPort(v int32) *DescribeLoadBalancerHTTPListenerAttributeResponseBody {
	s.ListenerPort = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) SetLoadBalancerId(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBody {
	s.LoadBalancerId = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) SetRequestId(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) SetRequestTimeout(v int32) *DescribeLoadBalancerHTTPListenerAttributeResponseBody {
	s.RequestTimeout = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) SetRules(v *DescribeLoadBalancerHTTPListenerAttributeResponseBodyRules) *DescribeLoadBalancerHTTPListenerAttributeResponseBody {
	s.Rules = v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) SetScheduler(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBody {
	s.Scheduler = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) SetSecurityStatus(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBody {
	s.SecurityStatus = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) SetStatus(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) SetStickySession(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBody {
	s.StickySession = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) SetStickySessionType(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBody {
	s.StickySessionType = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) SetTags(v *DescribeLoadBalancerHTTPListenerAttributeResponseBodyTags) *DescribeLoadBalancerHTTPListenerAttributeResponseBody {
	s.Tags = v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) SetUnhealthyThreshold(v int32) *DescribeLoadBalancerHTTPListenerAttributeResponseBody {
	s.UnhealthyThreshold = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) SetVServerGroupId(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBody {
	s.VServerGroupId = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) SetXForwardedFor(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBody {
	s.XForwardedFor = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) SetXForwardedFor_ClientSrcPort(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBody {
	s.XForwardedFor_ClientSrcPort = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) SetXForwardedFor_SLBID(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBody {
	s.XForwardedFor_SLBID = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) SetXForwardedFor_SLBIP(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBody {
	s.XForwardedFor_SLBIP = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) SetXForwardedFor_SLBPORT(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBody {
	s.XForwardedFor_SLBPORT = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) SetXForwardedFor_proto(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBody {
	s.XForwardedFor_proto = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) Validate() error {
	if s.AclIds != nil {
		if err := s.AclIds.Validate(); err != nil {
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

type DescribeLoadBalancerHTTPListenerAttributeResponseBodyAclIds struct {
	AclId []*string `json:"AclId,omitempty" xml:"AclId,omitempty" type:"Repeated"`
}

func (s DescribeLoadBalancerHTTPListenerAttributeResponseBodyAclIds) String() string {
	return dara.Prettify(s)
}

func (s DescribeLoadBalancerHTTPListenerAttributeResponseBodyAclIds) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBodyAclIds) GetAclId() []*string {
	return s.AclId
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBodyAclIds) SetAclId(v []*string) *DescribeLoadBalancerHTTPListenerAttributeResponseBodyAclIds {
	s.AclId = v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBodyAclIds) Validate() error {
	return dara.Validate(s)
}

type DescribeLoadBalancerHTTPListenerAttributeResponseBodyRules struct {
	Rule []*DescribeLoadBalancerHTTPListenerAttributeResponseBodyRulesRule `json:"Rule,omitempty" xml:"Rule,omitempty" type:"Repeated"`
}

func (s DescribeLoadBalancerHTTPListenerAttributeResponseBodyRules) String() string {
	return dara.Prettify(s)
}

func (s DescribeLoadBalancerHTTPListenerAttributeResponseBodyRules) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBodyRules) GetRule() []*DescribeLoadBalancerHTTPListenerAttributeResponseBodyRulesRule {
	return s.Rule
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBodyRules) SetRule(v []*DescribeLoadBalancerHTTPListenerAttributeResponseBodyRulesRule) *DescribeLoadBalancerHTTPListenerAttributeResponseBodyRules {
	s.Rule = v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBodyRules) Validate() error {
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

type DescribeLoadBalancerHTTPListenerAttributeResponseBodyRulesRule struct {
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
	// 1234
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The name of the forwarding rule.
	//
	// example:
	//
	// test
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
	// rsp-uf6w******
	VServerGroupId *string `json:"VServerGroupId,omitempty" xml:"VServerGroupId,omitempty"`
}

func (s DescribeLoadBalancerHTTPListenerAttributeResponseBodyRulesRule) String() string {
	return dara.Prettify(s)
}

func (s DescribeLoadBalancerHTTPListenerAttributeResponseBodyRulesRule) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBodyRulesRule) GetDomain() *string {
	return s.Domain
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBodyRulesRule) GetRuleId() *string {
	return s.RuleId
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBodyRulesRule) GetRuleName() *string {
	return s.RuleName
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBodyRulesRule) GetUrl() *string {
	return s.Url
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBodyRulesRule) GetVServerGroupId() *string {
	return s.VServerGroupId
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBodyRulesRule) SetDomain(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBodyRulesRule {
	s.Domain = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBodyRulesRule) SetRuleId(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBodyRulesRule {
	s.RuleId = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBodyRulesRule) SetRuleName(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBodyRulesRule {
	s.RuleName = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBodyRulesRule) SetUrl(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBodyRulesRule {
	s.Url = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBodyRulesRule) SetVServerGroupId(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBodyRulesRule {
	s.VServerGroupId = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBodyRulesRule) Validate() error {
	return dara.Validate(s)
}

type DescribeLoadBalancerHTTPListenerAttributeResponseBodyTags struct {
	Tag []*DescribeLoadBalancerHTTPListenerAttributeResponseBodyTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeLoadBalancerHTTPListenerAttributeResponseBodyTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeLoadBalancerHTTPListenerAttributeResponseBodyTags) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBodyTags) GetTag() []*DescribeLoadBalancerHTTPListenerAttributeResponseBodyTagsTag {
	return s.Tag
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBodyTags) SetTag(v []*DescribeLoadBalancerHTTPListenerAttributeResponseBodyTagsTag) *DescribeLoadBalancerHTTPListenerAttributeResponseBodyTags {
	s.Tag = v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBodyTags) Validate() error {
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

type DescribeLoadBalancerHTTPListenerAttributeResponseBodyTagsTag struct {
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

func (s DescribeLoadBalancerHTTPListenerAttributeResponseBodyTagsTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeLoadBalancerHTTPListenerAttributeResponseBodyTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBodyTagsTag) GetTagKey() *string {
	return s.TagKey
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBodyTagsTag) GetTagValue() *string {
	return s.TagValue
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBodyTagsTag) SetTagKey(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBodyTagsTag {
	s.TagKey = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBodyTagsTag) SetTagValue(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBodyTagsTag {
	s.TagValue = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBodyTagsTag) Validate() error {
	return dara.Validate(s)
}
