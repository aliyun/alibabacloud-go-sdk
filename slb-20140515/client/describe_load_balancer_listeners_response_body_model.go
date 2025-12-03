// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLoadBalancerListenersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetListeners(v []*DescribeLoadBalancerListenersResponseBodyListeners) *DescribeLoadBalancerListenersResponseBody
	GetListeners() []*DescribeLoadBalancerListenersResponseBodyListeners
	SetMaxResults(v int32) *DescribeLoadBalancerListenersResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeLoadBalancerListenersResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeLoadBalancerListenersResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeLoadBalancerListenersResponseBody
	GetTotalCount() *int32
}

type DescribeLoadBalancerListenersResponseBody struct {
	// A list of listeners of the CLB instance.
	//
	// >  This parameter is not returned if the CLB instance does not have a listener.
	Listeners []*DescribeLoadBalancerListenersResponseBodyListeners `json:"Listeners,omitempty" xml:"Listeners,omitempty" type:"Repeated"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that determines the start point of the query. Valid values:
	//
	// 	- If **NextToken*	- is empty, it indicates that no subsequent query is to be sent.
	//
	// 	- If a value is returned for **NextToken**, the value is the token that determines the start point of the next query.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 365F4154-92F6-4AE4-92F8-7FF34B540710
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeLoadBalancerListenersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLoadBalancerListenersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancerListenersResponseBody) GetListeners() []*DescribeLoadBalancerListenersResponseBodyListeners {
	return s.Listeners
}

func (s *DescribeLoadBalancerListenersResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeLoadBalancerListenersResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeLoadBalancerListenersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLoadBalancerListenersResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeLoadBalancerListenersResponseBody) SetListeners(v []*DescribeLoadBalancerListenersResponseBodyListeners) *DescribeLoadBalancerListenersResponseBody {
	s.Listeners = v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBody) SetMaxResults(v int32) *DescribeLoadBalancerListenersResponseBody {
	s.MaxResults = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBody) SetNextToken(v string) *DescribeLoadBalancerListenersResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBody) SetRequestId(v string) *DescribeLoadBalancerListenersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBody) SetTotalCount(v int32) *DescribeLoadBalancerListenersResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBody) Validate() error {
	if s.Listeners != nil {
		for _, item := range s.Listeners {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeLoadBalancerListenersResponseBodyListeners struct {
	// The ID of the access control list (ACL).
	//
	// example:
	//
	// nacl-a2do9e413e0spzasx****
	AclId *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
	// The IDs of the ACLs.
	AclIds []*string `json:"AclIds,omitempty" xml:"AclIds,omitempty" type:"Repeated"`
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
	// The type of access control. Valid values:
	//
	// 	- **white**: The listener forwards requests only from IP addresses and CIDR blocks on the whitelist. Your service may be adversely affected if the whitelist is not properly configured. If a whitelist is configured, the listener forwards requests only from IP addresses that are added to the whitelist.
	//
	// If you configure a whitelist but no IP address is added to the whitelist, the listener forwards all requests.
	//
	// 	- **black**: The listener blocks requests from IP addresses and CIDR blocks on the blacklist.
	//
	// If you configure a blacklist but no IP address is added to the blacklist, the listener forwards all requests.
	//
	// example:
	//
	// white
	AclType *string `json:"AclType,omitempty" xml:"AclType,omitempty"`
	// The port of the backend server.
	//
	// >  This parameter takes effect only when the `VServerGroupId` and `MasterSlaveServerGroupId` parameters are both empty.
	//
	// example:
	//
	// 80
	BackendServerPort *int32 `json:"BackendServerPort,omitempty" xml:"BackendServerPort,omitempty"`
	// The maximum bandwidth of the listener. Unit: Mbit/s.
	//
	// example:
	//
	// -1
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The description of the listener.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The configurations of the HTTP listener.
	HTTPListenerConfig *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig `json:"HTTPListenerConfig,omitempty" xml:"HTTPListenerConfig,omitempty" type:"Struct"`
	// The configurations of the HTTPS listener.
	HTTPSListenerConfig *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig `json:"HTTPSListenerConfig,omitempty" xml:"HTTPSListenerConfig,omitempty" type:"Struct"`
	// The listener port.
	//
	// example:
	//
	// 80
	ListenerPort *int32 `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	// The protocol used by the listener.
	//
	// example:
	//
	// http
	ListenerProtocol *string `json:"ListenerProtocol,omitempty" xml:"ListenerProtocol,omitempty"`
	// The ID of the CLB instance.
	//
	// example:
	//
	// lb-bp1b6c719dfa****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// The scheduling algorithm. Valid values:
	//
	// 	- **wrr**: Backend servers with higher weights receive more requests than those with lower weights.
	//
	// 	- **rr**: Requests are distributed to backend servers in sequence.
	//
	// 	- **sch**: consistent hashing that is based on source IP addresses. Requests from the same source IP address are distributed to the same backend server.
	//
	// 	- **tch**: specifies consistent hashing based on the source IP address, destination IP address, source port, and destination port. Requests that have the same four factors are distributed to the same backend server.
	//
	// 	- **qch**: specifies consistent hashing based on Quick UDP Internet Connection (QUIC) IDs. Requests that contain the same QUIC ID are scheduled to the same backend server.
	//
	// >  Only high-performance CLB instances support the **sch**, **tch**, and **qch*	- consistent hashing algorithms.
	//
	// example:
	//
	// wrr
	Scheduler *string `json:"Scheduler,omitempty" xml:"Scheduler,omitempty"`
	// The status of the listener. Valid values:
	//
	// 	- **running**
	//
	// 	- **stopped**
	//
	// example:
	//
	// running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The configurations of the TCP listener.
	TCPListenerConfig *DescribeLoadBalancerListenersResponseBodyListenersTCPListenerConfig `json:"TCPListenerConfig,omitempty" xml:"TCPListenerConfig,omitempty" type:"Struct"`
	// A list of tags.
	Tags []*DescribeLoadBalancerListenersResponseBodyListenersTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The configurations of the UDP listener.
	UDPListenerConfig *DescribeLoadBalancerListenersResponseBodyListenersUDPListenerConfig `json:"UDPListenerConfig,omitempty" xml:"UDPListenerConfig,omitempty" type:"Struct"`
	// The ID of the vServer group associated with the listener.
	//
	// example:
	//
	// rsp-cige6j****
	VServerGroupId *string `json:"VServerGroupId,omitempty" xml:"VServerGroupId,omitempty"`
}

func (s DescribeLoadBalancerListenersResponseBodyListeners) String() string {
	return dara.Prettify(s)
}

func (s DescribeLoadBalancerListenersResponseBodyListeners) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancerListenersResponseBodyListeners) GetAclId() *string {
	return s.AclId
}

func (s *DescribeLoadBalancerListenersResponseBodyListeners) GetAclIds() []*string {
	return s.AclIds
}

func (s *DescribeLoadBalancerListenersResponseBodyListeners) GetAclStatus() *string {
	return s.AclStatus
}

func (s *DescribeLoadBalancerListenersResponseBodyListeners) GetAclType() *string {
	return s.AclType
}

func (s *DescribeLoadBalancerListenersResponseBodyListeners) GetBackendServerPort() *int32 {
	return s.BackendServerPort
}

func (s *DescribeLoadBalancerListenersResponseBodyListeners) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *DescribeLoadBalancerListenersResponseBodyListeners) GetDescription() *string {
	return s.Description
}

func (s *DescribeLoadBalancerListenersResponseBodyListeners) GetHTTPListenerConfig() *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig {
	return s.HTTPListenerConfig
}

func (s *DescribeLoadBalancerListenersResponseBodyListeners) GetHTTPSListenerConfig() *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig {
	return s.HTTPSListenerConfig
}

func (s *DescribeLoadBalancerListenersResponseBodyListeners) GetListenerPort() *int32 {
	return s.ListenerPort
}

func (s *DescribeLoadBalancerListenersResponseBodyListeners) GetListenerProtocol() *string {
	return s.ListenerProtocol
}

func (s *DescribeLoadBalancerListenersResponseBodyListeners) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *DescribeLoadBalancerListenersResponseBodyListeners) GetScheduler() *string {
	return s.Scheduler
}

func (s *DescribeLoadBalancerListenersResponseBodyListeners) GetStatus() *string {
	return s.Status
}

func (s *DescribeLoadBalancerListenersResponseBodyListeners) GetTCPListenerConfig() *DescribeLoadBalancerListenersResponseBodyListenersTCPListenerConfig {
	return s.TCPListenerConfig
}

func (s *DescribeLoadBalancerListenersResponseBodyListeners) GetTags() []*DescribeLoadBalancerListenersResponseBodyListenersTags {
	return s.Tags
}

func (s *DescribeLoadBalancerListenersResponseBodyListeners) GetUDPListenerConfig() *DescribeLoadBalancerListenersResponseBodyListenersUDPListenerConfig {
	return s.UDPListenerConfig
}

func (s *DescribeLoadBalancerListenersResponseBodyListeners) GetVServerGroupId() *string {
	return s.VServerGroupId
}

func (s *DescribeLoadBalancerListenersResponseBodyListeners) SetAclId(v string) *DescribeLoadBalancerListenersResponseBodyListeners {
	s.AclId = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListeners) SetAclIds(v []*string) *DescribeLoadBalancerListenersResponseBodyListeners {
	s.AclIds = v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListeners) SetAclStatus(v string) *DescribeLoadBalancerListenersResponseBodyListeners {
	s.AclStatus = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListeners) SetAclType(v string) *DescribeLoadBalancerListenersResponseBodyListeners {
	s.AclType = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListeners) SetBackendServerPort(v int32) *DescribeLoadBalancerListenersResponseBodyListeners {
	s.BackendServerPort = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListeners) SetBandwidth(v int32) *DescribeLoadBalancerListenersResponseBodyListeners {
	s.Bandwidth = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListeners) SetDescription(v string) *DescribeLoadBalancerListenersResponseBodyListeners {
	s.Description = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListeners) SetHTTPListenerConfig(v *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig) *DescribeLoadBalancerListenersResponseBodyListeners {
	s.HTTPListenerConfig = v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListeners) SetHTTPSListenerConfig(v *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig) *DescribeLoadBalancerListenersResponseBodyListeners {
	s.HTTPSListenerConfig = v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListeners) SetListenerPort(v int32) *DescribeLoadBalancerListenersResponseBodyListeners {
	s.ListenerPort = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListeners) SetListenerProtocol(v string) *DescribeLoadBalancerListenersResponseBodyListeners {
	s.ListenerProtocol = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListeners) SetLoadBalancerId(v string) *DescribeLoadBalancerListenersResponseBodyListeners {
	s.LoadBalancerId = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListeners) SetScheduler(v string) *DescribeLoadBalancerListenersResponseBodyListeners {
	s.Scheduler = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListeners) SetStatus(v string) *DescribeLoadBalancerListenersResponseBodyListeners {
	s.Status = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListeners) SetTCPListenerConfig(v *DescribeLoadBalancerListenersResponseBodyListenersTCPListenerConfig) *DescribeLoadBalancerListenersResponseBodyListeners {
	s.TCPListenerConfig = v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListeners) SetTags(v []*DescribeLoadBalancerListenersResponseBodyListenersTags) *DescribeLoadBalancerListenersResponseBodyListeners {
	s.Tags = v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListeners) SetUDPListenerConfig(v *DescribeLoadBalancerListenersResponseBodyListenersUDPListenerConfig) *DescribeLoadBalancerListenersResponseBodyListeners {
	s.UDPListenerConfig = v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListeners) SetVServerGroupId(v string) *DescribeLoadBalancerListenersResponseBodyListeners {
	s.VServerGroupId = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListeners) Validate() error {
	if s.HTTPListenerConfig != nil {
		if err := s.HTTPListenerConfig.Validate(); err != nil {
			return err
		}
	}
	if s.HTTPSListenerConfig != nil {
		if err := s.HTTPSListenerConfig.Validate(); err != nil {
			return err
		}
	}
	if s.TCPListenerConfig != nil {
		if err := s.TCPListenerConfig.Validate(); err != nil {
			return err
		}
	}
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.UDPListenerConfig != nil {
		if err := s.UDPListenerConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig struct {
	// The cookie configures for the server.
	//
	// example:
	//
	// B490B5EBF6F3CD402E515D22BCDA****
	Cookie *string `json:"Cookie,omitempty" xml:"Cookie,omitempty"`
	// The maximum amount of time to wait before the session cookie expires. Unit: seconds.
	//
	// Valid values: **1*	- to **86400**.
	//
	// example:
	//
	// 500
	CookieTimeout *int32 `json:"CookieTimeout,omitempty" xml:"CookieTimeout,omitempty"`
	// The listener port that is used for HTTP-to-HTTPS redirection.
	//
	// >  If the **ListenerForward*	- parameter is set to **off**, this parameter is not returned.
	//
	// example:
	//
	// 80
	ForwardPort *int32 `json:"ForwardPort,omitempty" xml:"ForwardPort,omitempty"`
	// Indicates whether GZIP compression is enabled. Valid values:
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
	// >  This parameter takes effect only when **HealthCheck*	- is set to **on**.
	//
	// example:
	//
	// 8080
	HealthCheckConnectPort *int32 `json:"HealthCheckConnectPort,omitempty" xml:"HealthCheckConnectPort,omitempty"`
	// The domain name that is used for health checks.
	//
	// example:
	//
	// www.example.com
	HealthCheckDomain *string `json:"HealthCheckDomain,omitempty" xml:"HealthCheckDomain,omitempty"`
	// The HTTP status code that indicates a healthy backend server.
	//
	// example:
	//
	// http_2xx,http_3xx
	HealthCheckHttpCode *string `json:"HealthCheckHttpCode,omitempty" xml:"HealthCheckHttpCode,omitempty"`
	// The HTTP version for health checks.
	//
	// example:
	//
	// HTTP 1.0
	HealthCheckHttpVersion *string `json:"HealthCheckHttpVersion,omitempty" xml:"HealthCheckHttpVersion,omitempty"`
	// The interval at which health checks are performed. Unit: seconds.
	//
	// example:
	//
	// 5
	HealthCheckInterval *int32 `json:"HealthCheckInterval,omitempty" xml:"HealthCheckInterval,omitempty"`
	// The health check method. Valid values: **head*	- and **get**.
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
	// The protocol that you want to use for health checks.
	//
	// example:
	//
	// tcp
	HealthCheckType *string `json:"HealthCheckType,omitempty" xml:"HealthCheckType,omitempty"`
	// The URI that is used for health checks.
	//
	// example:
	//
	// /test/index.html
	HealthCheckURI *string `json:"HealthCheckURI,omitempty" xml:"HealthCheckURI,omitempty"`
	// The number of times that an unhealthy backend server must consecutively pass health checks before it is declared healthy. In this case, the health check status of the backend server changes from **fail*	- to **success**.
	//
	// Valid values: **2*	- to **10**.
	//
	// example:
	//
	// 4
	HealthyThreshold *int32 `json:"HealthyThreshold,omitempty" xml:"HealthyThreshold,omitempty"`
	// The timeout period of an idle connection. Unit: seconds. Valid values: **1*	- to **60**.
	//
	// If no request is received within the specified timeout period, CLB closes the connection. When a request is received, CLB establishes a new connection.
	//
	// example:
	//
	// 15
	IdleTimeout *int32 `json:"IdleTimeout,omitempty" xml:"IdleTimeout,omitempty"`
	// Indicates whether HTTP-to-HTTPS redirection is enabled for the listener. Valid values:
	//
	// 	- **on**
	//
	// 	- **off**
	//
	// example:
	//
	// on
	ListenerForward *string `json:"ListenerForward,omitempty" xml:"ListenerForward,omitempty"`
	// The timeout period of a request. Unit: seconds. Valid values: **1*	- to **180**.
	//
	// If no response is received from a backend server during the request timeout period, CLB sends the `HTTP 504` status code to the client.
	//
	// example:
	//
	// 60
	RequestTimeout *int32 `json:"RequestTimeout,omitempty" xml:"RequestTimeout,omitempty"`
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
	// The method used to handle the cookie. Valid values:
	//
	// 	- **insert**: inserts a cookie. CLB inserts the SERVERID cookie to the HTTP or HTTPS response to the first request from a client. Subsequent requests that carry the SERVERID cookie from the client are forwarded to the same backend server as the first request.
	//
	// 	- **server**: rewrites the original cookie. CLB rewrites the custom cookies in requests from a client. Subsequent requests from the client that carry the new cookie are forwarded to the same backend server as the first request.
	//
	// example:
	//
	// insert
	StickySessionType *string `json:"StickySessionType,omitempty" xml:"StickySessionType,omitempty"`
	// The number of times that a healthy backend server must consecutively fail health checks before it is declared unhealthy. In this case, the health check status of the backend server changes from **success*	- to **fail**.
	//
	// Valid values: **2*	- to **10**.
	//
	// example:
	//
	// 4
	UnhealthyThreshold *int32 `json:"UnhealthyThreshold,omitempty" xml:"UnhealthyThreshold,omitempty"`
	// Indicates whether the `X-Forwarded-For` header is used to preserve client IP addresses. Valid values:
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
	// on
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
	// Indicates whether the `SLB-IP` header is used to retrieve the VIP of the client. Valid values:
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
	// on
	XForwardedFor_SLBPORT *string `json:"XForwardedFor_SLBPORT,omitempty" xml:"XForwardedFor_SLBPORT,omitempty"`
	// Indicates whether the `X-Forwarded-Proto` header is used to obtain the listener protocol. Valid values:
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

func (s DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig) GetCookie() *string {
	return s.Cookie
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig) GetCookieTimeout() *int32 {
	return s.CookieTimeout
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig) GetForwardPort() *int32 {
	return s.ForwardPort
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig) GetGzip() *string {
	return s.Gzip
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig) GetHealthCheck() *string {
	return s.HealthCheck
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig) GetHealthCheckConnectPort() *int32 {
	return s.HealthCheckConnectPort
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig) GetHealthCheckDomain() *string {
	return s.HealthCheckDomain
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig) GetHealthCheckHttpCode() *string {
	return s.HealthCheckHttpCode
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig) GetHealthCheckHttpVersion() *string {
	return s.HealthCheckHttpVersion
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig) GetHealthCheckInterval() *int32 {
	return s.HealthCheckInterval
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig) GetHealthCheckMethod() *string {
	return s.HealthCheckMethod
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig) GetHealthCheckTimeout() *int32 {
	return s.HealthCheckTimeout
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig) GetHealthCheckType() *string {
	return s.HealthCheckType
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig) GetHealthCheckURI() *string {
	return s.HealthCheckURI
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig) GetHealthyThreshold() *int32 {
	return s.HealthyThreshold
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig) GetIdleTimeout() *int32 {
	return s.IdleTimeout
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig) GetListenerForward() *string {
	return s.ListenerForward
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig) GetRequestTimeout() *int32 {
	return s.RequestTimeout
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig) GetStickySession() *string {
	return s.StickySession
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig) GetStickySessionType() *string {
	return s.StickySessionType
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig) GetUnhealthyThreshold() *int32 {
	return s.UnhealthyThreshold
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig) GetXForwardedFor() *string {
	return s.XForwardedFor
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig) GetXForwardedFor_ClientSrcPort() *string {
	return s.XForwardedFor_ClientSrcPort
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig) GetXForwardedFor_SLBID() *string {
	return s.XForwardedFor_SLBID
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig) GetXForwardedFor_SLBIP() *string {
	return s.XForwardedFor_SLBIP
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig) GetXForwardedFor_SLBPORT() *string {
	return s.XForwardedFor_SLBPORT
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig) GetXForwardedFor_proto() *string {
	return s.XForwardedFor_proto
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig) SetCookie(v string) *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig {
	s.Cookie = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig) SetCookieTimeout(v int32) *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig {
	s.CookieTimeout = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig) SetForwardPort(v int32) *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig {
	s.ForwardPort = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig) SetGzip(v string) *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig {
	s.Gzip = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig) SetHealthCheck(v string) *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig {
	s.HealthCheck = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig) SetHealthCheckConnectPort(v int32) *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig {
	s.HealthCheckConnectPort = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig) SetHealthCheckDomain(v string) *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig {
	s.HealthCheckDomain = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig) SetHealthCheckHttpCode(v string) *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig {
	s.HealthCheckHttpCode = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig) SetHealthCheckHttpVersion(v string) *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig {
	s.HealthCheckHttpVersion = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig) SetHealthCheckInterval(v int32) *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig {
	s.HealthCheckInterval = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig) SetHealthCheckMethod(v string) *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig {
	s.HealthCheckMethod = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig) SetHealthCheckTimeout(v int32) *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig {
	s.HealthCheckTimeout = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig) SetHealthCheckType(v string) *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig {
	s.HealthCheckType = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig) SetHealthCheckURI(v string) *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig {
	s.HealthCheckURI = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig) SetHealthyThreshold(v int32) *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig {
	s.HealthyThreshold = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig) SetIdleTimeout(v int32) *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig {
	s.IdleTimeout = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig) SetListenerForward(v string) *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig {
	s.ListenerForward = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig) SetRequestTimeout(v int32) *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig {
	s.RequestTimeout = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig) SetStickySession(v string) *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig {
	s.StickySession = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig) SetStickySessionType(v string) *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig {
	s.StickySessionType = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig) SetUnhealthyThreshold(v int32) *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig {
	s.UnhealthyThreshold = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig) SetXForwardedFor(v string) *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig {
	s.XForwardedFor = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig) SetXForwardedFor_ClientSrcPort(v string) *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig {
	s.XForwardedFor_ClientSrcPort = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig) SetXForwardedFor_SLBID(v string) *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig {
	s.XForwardedFor_SLBID = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig) SetXForwardedFor_SLBIP(v string) *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig {
	s.XForwardedFor_SLBIP = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig) SetXForwardedFor_SLBPORT(v string) *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig {
	s.XForwardedFor_SLBPORT = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig) SetXForwardedFor_proto(v string) *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig {
	s.XForwardedFor_proto = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig) Validate() error {
	return dara.Validate(s)
}

type DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig struct {
	// The ID of the CA certificate.
	//
	// example:
	//
	// idkp-234-cn-test-0****
	CACertificateId *string `json:"CACertificateId,omitempty" xml:"CACertificateId,omitempty"`
	// The cookie configures for the server.
	//
	// example:
	//
	// B490B5EBF6F3CD402E515D22BCDA****
	Cookie *string `json:"Cookie,omitempty" xml:"Cookie,omitempty"`
	// The maximum amount of time to wait before the session cookie expires. Unit: seconds.
	//
	// Valid values: **1*	- to **86400**.
	//
	// example:
	//
	// 500
	CookieTimeout *int32 `json:"CookieTimeout,omitempty" xml:"CookieTimeout,omitempty"`
	// Indicates whether `HTTP 2.0` is enabled. Valid values:
	//
	// 	- **on**
	//
	// 	- **off**
	//
	// example:
	//
	// off
	EnableHttp2 *string `json:"EnableHttp2,omitempty" xml:"EnableHttp2,omitempty"`
	// Indicates whether GZIP compression is enabled. Valid values:
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
	// example:
	//
	// 8080
	HealthCheckConnectPort *int32 `json:"HealthCheckConnectPort,omitempty" xml:"HealthCheckConnectPort,omitempty"`
	// The domain name that is used for health checks.
	//
	// example:
	//
	// www.example.com
	HealthCheckDomain *string `json:"HealthCheckDomain,omitempty" xml:"HealthCheckDomain,omitempty"`
	// The HTTP status code that indicates a healthy backend server.
	//
	// example:
	//
	// http_2xx,http_3xx
	HealthCheckHttpCode *string `json:"HealthCheckHttpCode,omitempty" xml:"HealthCheckHttpCode,omitempty"`
	// The HTTP version for health checks.
	//
	// example:
	//
	// HTTP 1.0
	HealthCheckHttpVersion *string `json:"HealthCheckHttpVersion,omitempty" xml:"HealthCheckHttpVersion,omitempty"`
	// The interval at which health checks are performed. Unit: seconds.
	//
	// example:
	//
	// 5
	HealthCheckInterval *int32 `json:"HealthCheckInterval,omitempty" xml:"HealthCheckInterval,omitempty"`
	// The health check method.
	//
	// example:
	//
	// get
	HealthCheckMethod *string `json:"HealthCheckMethod,omitempty" xml:"HealthCheckMethod,omitempty"`
	// The timeout period of a health check response. Unit: seconds.
	//
	// example:
	//
	// 3
	HealthCheckTimeout *int32 `json:"HealthCheckTimeout,omitempty" xml:"HealthCheckTimeout,omitempty"`
	// The protocol that you want to use for health checks.
	//
	// example:
	//
	// tcp
	HealthCheckType *string `json:"HealthCheckType,omitempty" xml:"HealthCheckType,omitempty"`
	// The URI that is used for health checks.
	//
	// example:
	//
	// /test/index.html
	HealthCheckURI *string `json:"HealthCheckURI,omitempty" xml:"HealthCheckURI,omitempty"`
	// The number of times that an unhealthy backend server must consecutively pass health checks before it is declared healthy. In this case, the health check status of the backend server changes from **fail*	- to **success**.
	//
	// Valid values: **2*	- to **10**.
	//
	// example:
	//
	// 4
	HealthyThreshold *int32 `json:"HealthyThreshold,omitempty" xml:"HealthyThreshold,omitempty"`
	// The timeout period of an idle connection. Unit: seconds. Valid values: **1*	- to **60**.
	//
	// If no request is received within the specified timeout period, CLB closes the connection. When a request is received, CLB establishes a new connection.
	//
	// example:
	//
	// 15
	IdleTimeout *int32 `json:"IdleTimeout,omitempty" xml:"IdleTimeout,omitempty"`
	// The request timeout period. Unit: seconds. Valid values: **1*	- to **180**.
	//
	// If no response is received from a backend server during the request timeout period, CLB sends the `HTTP 504` status code to the client.
	//
	// example:
	//
	// 60
	RequestTimeout *int32 `json:"RequestTimeout,omitempty" xml:"RequestTimeout,omitempty"`
	// The ID of the server certificate.
	//
	// example:
	//
	// idkp-123-cn-test-0****
	ServerCertificateId *string `json:"ServerCertificateId,omitempty" xml:"ServerCertificateId,omitempty"`
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
	// The method used to handle the cookie.
	//
	// 	- **insert**: inserts a cookie. CLB inserts the SERVERID cookie to the HTTP or HTTPS response to the first request from a client. Subsequent requests that carry the SERVERID cookie from the client are forwarded to the same backend server as the first request.
	//
	// 	- **server**: rewrites the original cookie. CLB rewrites the custom cookies in requests from a client. Subsequent requests from the client that carry the new cookie are forwarded to the same backend server as the first request.
	//
	// example:
	//
	// insert
	StickySessionType *string `json:"StickySessionType,omitempty" xml:"StickySessionType,omitempty"`
	// A TLS security policy contains TLS protocols and cipher suites available for HTTPS.
	//
	// 	- **tls_cipher_policy_1_0**:
	//
	//     Supported TLS versions: TLSv1.0, TLSv1.1, and TLSv1.2.
	//
	//     Supported cipher suites: ECDHE-RSA-AES128-GCM-SHA256, ECDHE-RSA-AES256-GCM-SHA384, ECDHE-RSA-AES128-SHA256, ECDHE-RSA-AES256-SHA384, AES128-GCM-SHA256, AES256-GCM-SHA384, AES128-SHA256, AES256-SHA256, ECDHE-RSA-AES128-SHA, ECDHE-RSA-AES256-SHA, AES128-SHA, AES256-SHA, and DES-CBC3-SHA.
	//
	// 	- **tls_cipher_policy_1_1**:
	//
	//     Supported TLS versions: TLSv1.1 and TLSv1.2.
	//
	//     Supported cipher suites: ECDHE-RSA-AES128-GCM-SHA256, ECDHE-RSA-AES256-GCM-SHA384, ECDHE-RSA-AES128-SHA256, ECDHE-RSA-AES256-SHA384, AES128-GCM-SHA256, AES256-GCM-SHA384, AES128-SHA256, AES256-SHA256, ECDHE-RSA-AES128-SHA, ECDHE-RSA-AES256-SHA, AES128-SHA, AES256-SHA, and DES-CBC3-SHA.
	//
	// 	- **tls_cipher_policy_1_2**
	//
	//     Supported TLS versions: TLSv1.2.
	//
	//     Supported cipher suites: ECDHE-RSA-AES128-GCM-SHA256, ECDHE-RSA-AES256-GCM-SHA384, ECDHE-RSA-AES128-SHA256, ECDHE-RSA-AES256-SHA384, AES128-GCM-SHA256, AES256-GCM-SHA384, AES128-SHA256, AES256-SHA256, ECDHE-RSA-AES128-SHA, ECDHE-RSA-AES256-SHA, AES128-SHA, AES256-SHA, and DES-CBC3-SHA.
	//
	// 	- **tls_cipher_policy_1_2_strict**
	//
	//     Supported TLS versions: TLSv1.2.
	//
	//     Supported cipher suites: ECDHE-RSA-AES128-GCM-SHA256, ECDHE-RSA-AES256-GCM-SHA384, ECDHE-RSA-AES128-SHA256, ECDHE-RSA-AES256-SHA384, ECDHE-RSA-AES128-SHA, and ECDHE-RSA-AES256-SHA.
	//
	// 	- **tls_cipher_policy_1_2_strict_with_1_3**
	//
	//     Supported TLS versions: TLSv1.2 and TLSv1.3.
	//
	//     Supported cipher suites: TLS_AES_128_GCM_SHA256, TLS_AES_256_GCM_SHA384, TLS_CHACHA20_POLY1305_SHA256, TLS_AES_128_CCM_SHA256, TLS_AES_128_CCM_8_SHA256, ECDHE-ECDSA-AES128-GCM-SHA256, ECDHE-ECDSA-AES256-GCM-SHA384, ECDHE-ECDSA-AES128-SHA256, ECDHE-ECDSA-AES256-SHA384, ECDHE-RSA-AES128-GCM-SHA256, ECDHE-RSA-AES256-GCM-SHA384, ECDHE-RSA-AES128-SHA256, ECDHE-RSA-AES256-SHA384, ECDHE-ECDSA-AES128-SHA, ECDHE-ECDSA-AES256-SHA, ECDHE-RSA-AES128-SHA, and ECDHE-RSA-AES256-SHA.
	//
	// example:
	//
	// tls_cipher_policy_1_0
	TLSCipherPolicy *string `json:"TLSCipherPolicy,omitempty" xml:"TLSCipherPolicy,omitempty"`
	// The number of times that a healthy backend server must consecutively fail health checks before it is declared unhealthy. In this case, the health check status of the backend server changes from **success*	- to **fail**.
	//
	// Valid values: **2*	- to **10**.
	//
	// example:
	//
	// 4
	UnhealthyThreshold *int32 `json:"UnhealthyThreshold,omitempty" xml:"UnhealthyThreshold,omitempty"`
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
	// Indicates whether the `XForwardedFor_ClientCertClientVerify` header is used to obtain the verification result of the client certificate. Valid values:
	//
	// 	- **on**
	//
	// 	- **off**
	//
	// example:
	//
	// off
	XForwardedFor_ClientCertClientVerify *string `json:"XForwardedFor_ClientCertClientVerify,omitempty" xml:"XForwardedFor_ClientCertClientVerify,omitempty"`
	// Indicates whether the `XForwardedFor_ClientCertFingerprint` header is used to obtain the fingerprint of the client certificate. Valid values:
	//
	// 	- **on**
	//
	// 	- **off**
	//
	// example:
	//
	// off
	XForwardedFor_ClientCertFingerprint *string `json:"XForwardedFor_ClientCertFingerprint,omitempty" xml:"XForwardedFor_ClientCertFingerprint,omitempty"`
	// Indicates whether the `XForwardedFor_ClientCertIssuerDN` header is used to obtain information about the authority that issues the client certificate. Valid values:
	//
	// 	- **on**
	//
	// 	- **off**
	//
	// example:
	//
	// off
	XForwardedFor_ClientCertIssuerDN *string `json:"XForwardedFor_ClientCertIssuerDN,omitempty" xml:"XForwardedFor_ClientCertIssuerDN,omitempty"`
	// Indicates whether the `XForwardedFor_ClientCertSubjectDN` header is used to obtain information about the owner of the client certificate. Valid values:
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
	// Indicates whether the `SLB-IP` header is used to retrieve the VIP of the client. Valid values:
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
	// Indicates whether the `X-Forwarded-Proto` header is used to obtain the listener protocol. Valid values:
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

func (s DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig) GetCACertificateId() *string {
	return s.CACertificateId
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig) GetCookie() *string {
	return s.Cookie
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig) GetCookieTimeout() *int32 {
	return s.CookieTimeout
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig) GetEnableHttp2() *string {
	return s.EnableHttp2
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig) GetGzip() *string {
	return s.Gzip
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig) GetHealthCheck() *string {
	return s.HealthCheck
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig) GetHealthCheckConnectPort() *int32 {
	return s.HealthCheckConnectPort
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig) GetHealthCheckDomain() *string {
	return s.HealthCheckDomain
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig) GetHealthCheckHttpCode() *string {
	return s.HealthCheckHttpCode
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig) GetHealthCheckHttpVersion() *string {
	return s.HealthCheckHttpVersion
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig) GetHealthCheckInterval() *int32 {
	return s.HealthCheckInterval
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig) GetHealthCheckMethod() *string {
	return s.HealthCheckMethod
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig) GetHealthCheckTimeout() *int32 {
	return s.HealthCheckTimeout
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig) GetHealthCheckType() *string {
	return s.HealthCheckType
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig) GetHealthCheckURI() *string {
	return s.HealthCheckURI
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig) GetHealthyThreshold() *int32 {
	return s.HealthyThreshold
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig) GetIdleTimeout() *int32 {
	return s.IdleTimeout
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig) GetRequestTimeout() *int32 {
	return s.RequestTimeout
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig) GetServerCertificateId() *string {
	return s.ServerCertificateId
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig) GetStickySession() *string {
	return s.StickySession
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig) GetStickySessionType() *string {
	return s.StickySessionType
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig) GetTLSCipherPolicy() *string {
	return s.TLSCipherPolicy
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig) GetUnhealthyThreshold() *int32 {
	return s.UnhealthyThreshold
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig) GetXForwardedFor() *string {
	return s.XForwardedFor
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig) GetXForwardedFor_ClientCertClientVerify() *string {
	return s.XForwardedFor_ClientCertClientVerify
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig) GetXForwardedFor_ClientCertFingerprint() *string {
	return s.XForwardedFor_ClientCertFingerprint
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig) GetXForwardedFor_ClientCertIssuerDN() *string {
	return s.XForwardedFor_ClientCertIssuerDN
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig) GetXForwardedFor_ClientCertSubjectDN() *string {
	return s.XForwardedFor_ClientCertSubjectDN
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig) GetXForwardedFor_ClientSrcPort() *string {
	return s.XForwardedFor_ClientSrcPort
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig) GetXForwardedFor_SLBID() *string {
	return s.XForwardedFor_SLBID
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig) GetXForwardedFor_SLBIP() *string {
	return s.XForwardedFor_SLBIP
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig) GetXForwardedFor_SLBPORT() *string {
	return s.XForwardedFor_SLBPORT
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig) GetXForwardedFor_proto() *string {
	return s.XForwardedFor_proto
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig) SetCACertificateId(v string) *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig {
	s.CACertificateId = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig) SetCookie(v string) *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig {
	s.Cookie = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig) SetCookieTimeout(v int32) *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig {
	s.CookieTimeout = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig) SetEnableHttp2(v string) *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig {
	s.EnableHttp2 = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig) SetGzip(v string) *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig {
	s.Gzip = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig) SetHealthCheck(v string) *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig {
	s.HealthCheck = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig) SetHealthCheckConnectPort(v int32) *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig {
	s.HealthCheckConnectPort = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig) SetHealthCheckDomain(v string) *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig {
	s.HealthCheckDomain = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig) SetHealthCheckHttpCode(v string) *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig {
	s.HealthCheckHttpCode = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig) SetHealthCheckHttpVersion(v string) *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig {
	s.HealthCheckHttpVersion = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig) SetHealthCheckInterval(v int32) *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig {
	s.HealthCheckInterval = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig) SetHealthCheckMethod(v string) *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig {
	s.HealthCheckMethod = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig) SetHealthCheckTimeout(v int32) *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig {
	s.HealthCheckTimeout = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig) SetHealthCheckType(v string) *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig {
	s.HealthCheckType = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig) SetHealthCheckURI(v string) *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig {
	s.HealthCheckURI = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig) SetHealthyThreshold(v int32) *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig {
	s.HealthyThreshold = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig) SetIdleTimeout(v int32) *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig {
	s.IdleTimeout = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig) SetRequestTimeout(v int32) *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig {
	s.RequestTimeout = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig) SetServerCertificateId(v string) *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig {
	s.ServerCertificateId = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig) SetStickySession(v string) *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig {
	s.StickySession = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig) SetStickySessionType(v string) *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig {
	s.StickySessionType = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig) SetTLSCipherPolicy(v string) *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig {
	s.TLSCipherPolicy = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig) SetUnhealthyThreshold(v int32) *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig {
	s.UnhealthyThreshold = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig) SetXForwardedFor(v string) *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig {
	s.XForwardedFor = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig) SetXForwardedFor_ClientCertClientVerify(v string) *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig {
	s.XForwardedFor_ClientCertClientVerify = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig) SetXForwardedFor_ClientCertFingerprint(v string) *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig {
	s.XForwardedFor_ClientCertFingerprint = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig) SetXForwardedFor_ClientCertIssuerDN(v string) *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig {
	s.XForwardedFor_ClientCertIssuerDN = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig) SetXForwardedFor_ClientCertSubjectDN(v string) *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig {
	s.XForwardedFor_ClientCertSubjectDN = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig) SetXForwardedFor_ClientSrcPort(v string) *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig {
	s.XForwardedFor_ClientSrcPort = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig) SetXForwardedFor_SLBID(v string) *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig {
	s.XForwardedFor_SLBID = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig) SetXForwardedFor_SLBIP(v string) *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig {
	s.XForwardedFor_SLBIP = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig) SetXForwardedFor_SLBPORT(v string) *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig {
	s.XForwardedFor_SLBPORT = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig) SetXForwardedFor_proto(v string) *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig {
	s.XForwardedFor_proto = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig) Validate() error {
	return dara.Validate(s)
}

type DescribeLoadBalancerListenersResponseBodyListenersTCPListenerConfig struct {
	// Indicates whether connection draining is enabled. Valid values:
	//
	// 	- **on**
	//
	// 	- **off**
	//
	// example:
	//
	// off
	ConnectionDrain *string `json:"ConnectionDrain,omitempty" xml:"ConnectionDrain,omitempty"`
	// The timeout period of connection draining. Unit: seconds.
	//
	// Value values: **10 to 900**.
	//
	// example:
	//
	// 300
	ConnectionDrainTimeout *int32 `json:"ConnectionDrainTimeout,omitempty" xml:"ConnectionDrainTimeout,omitempty"`
	// The timeout period of a connection. Unit: seconds.
	//
	// example:
	//
	// 500
	EstablishedTimeout *int32 `json:"EstablishedTimeout,omitempty" xml:"EstablishedTimeout,omitempty"`
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
	// example:
	//
	// 8080
	HealthCheckConnectPort *int32 `json:"HealthCheckConnectPort,omitempty" xml:"HealthCheckConnectPort,omitempty"`
	// The timeout period of health checks. Unit: seconds.
	//
	// Valid values: **1*	- to **300**.
	//
	// example:
	//
	// 100
	HealthCheckConnectTimeout *int32 `json:"HealthCheckConnectTimeout,omitempty" xml:"HealthCheckConnectTimeout,omitempty"`
	// The domain name that is used for health checks.
	//
	// example:
	//
	// www.example.com
	HealthCheckDomain *string `json:"HealthCheckDomain,omitempty" xml:"HealthCheckDomain,omitempty"`
	// The HTTP status code that indicates a healthy backend server.
	//
	// example:
	//
	// http_2xx,http_3xx
	HealthCheckHttpCode *string `json:"HealthCheckHttpCode,omitempty" xml:"HealthCheckHttpCode,omitempty"`
	// The interval between two consecutive health checks. Unit: seconds.
	//
	// example:
	//
	// 5
	HealthCheckInterval *int32 `json:"HealthCheckInterval,omitempty" xml:"HealthCheckInterval,omitempty"`
	// The health check method.
	//
	// example:
	//
	// get
	HealthCheckMethod *string `json:"HealthCheckMethod,omitempty" xml:"HealthCheckMethod,omitempty"`
	// The protocol that you want to use for health checks.
	//
	// example:
	//
	// tcp
	HealthCheckType *string `json:"HealthCheckType,omitempty" xml:"HealthCheckType,omitempty"`
	// The URI that is used for health checks.
	//
	// example:
	//
	// /test/index.html
	HealthCheckURI *string `json:"HealthCheckURI,omitempty" xml:"HealthCheckURI,omitempty"`
	// The number of times that an unhealthy backend server must consecutively pass health checks before it is declared healthy. In this case, the health check status of the backend server changes from **fail*	- to **success**.
	//
	// Valid values: **2*	- to **10**.
	//
	// example:
	//
	// 4
	HealthyThreshold *int32 `json:"HealthyThreshold,omitempty" xml:"HealthyThreshold,omitempty"`
	// The ID of the primary/secondary server group associated with the listener.
	//
	// example:
	//
	// rsp-0bfucw*****
	MasterSlaveServerGroupId *string `json:"MasterSlaveServerGroupId,omitempty" xml:"MasterSlaveServerGroupId,omitempty"`
	// Indicates whether session persistence is enabled. Unit: seconds.
	//
	// Valid values: **0*	- to **3600**.
	//
	// **0*	- indicates that session persistence is disabled.
	//
	// example:
	//
	// 0
	PersistenceTimeout *int32 `json:"PersistenceTimeout,omitempty" xml:"PersistenceTimeout,omitempty"`
	// Indicates whether the Proxy protocol is used to pass source client IP addresses to backend servers. Valid values:
	//
	// 	- **true**: enables the burst feature for the data disk.
	//
	// 	- **false**: The task is not being retried.
	//
	// example:
	//
	// false
	ProxyProtocolV2Enabled *string `json:"ProxyProtocolV2Enabled,omitempty" xml:"ProxyProtocolV2Enabled,omitempty"`
	// The number of times that a healthy backend server must consecutively fail health checks before it is declared unhealthy. In this case, the health check status of the backend server changes from **success*	- to **fail**.
	//
	// Valid values: **2*	- to **10**.
	//
	// example:
	//
	// 4
	UnhealthyThreshold *int32 `json:"UnhealthyThreshold,omitempty" xml:"UnhealthyThreshold,omitempty"`
}

func (s DescribeLoadBalancerListenersResponseBodyListenersTCPListenerConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeLoadBalancerListenersResponseBodyListenersTCPListenerConfig) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersTCPListenerConfig) GetConnectionDrain() *string {
	return s.ConnectionDrain
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersTCPListenerConfig) GetConnectionDrainTimeout() *int32 {
	return s.ConnectionDrainTimeout
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersTCPListenerConfig) GetEstablishedTimeout() *int32 {
	return s.EstablishedTimeout
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersTCPListenerConfig) GetHealthCheck() *string {
	return s.HealthCheck
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersTCPListenerConfig) GetHealthCheckConnectPort() *int32 {
	return s.HealthCheckConnectPort
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersTCPListenerConfig) GetHealthCheckConnectTimeout() *int32 {
	return s.HealthCheckConnectTimeout
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersTCPListenerConfig) GetHealthCheckDomain() *string {
	return s.HealthCheckDomain
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersTCPListenerConfig) GetHealthCheckHttpCode() *string {
	return s.HealthCheckHttpCode
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersTCPListenerConfig) GetHealthCheckInterval() *int32 {
	return s.HealthCheckInterval
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersTCPListenerConfig) GetHealthCheckMethod() *string {
	return s.HealthCheckMethod
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersTCPListenerConfig) GetHealthCheckType() *string {
	return s.HealthCheckType
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersTCPListenerConfig) GetHealthCheckURI() *string {
	return s.HealthCheckURI
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersTCPListenerConfig) GetHealthyThreshold() *int32 {
	return s.HealthyThreshold
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersTCPListenerConfig) GetMasterSlaveServerGroupId() *string {
	return s.MasterSlaveServerGroupId
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersTCPListenerConfig) GetPersistenceTimeout() *int32 {
	return s.PersistenceTimeout
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersTCPListenerConfig) GetProxyProtocolV2Enabled() *string {
	return s.ProxyProtocolV2Enabled
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersTCPListenerConfig) GetUnhealthyThreshold() *int32 {
	return s.UnhealthyThreshold
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersTCPListenerConfig) SetConnectionDrain(v string) *DescribeLoadBalancerListenersResponseBodyListenersTCPListenerConfig {
	s.ConnectionDrain = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersTCPListenerConfig) SetConnectionDrainTimeout(v int32) *DescribeLoadBalancerListenersResponseBodyListenersTCPListenerConfig {
	s.ConnectionDrainTimeout = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersTCPListenerConfig) SetEstablishedTimeout(v int32) *DescribeLoadBalancerListenersResponseBodyListenersTCPListenerConfig {
	s.EstablishedTimeout = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersTCPListenerConfig) SetHealthCheck(v string) *DescribeLoadBalancerListenersResponseBodyListenersTCPListenerConfig {
	s.HealthCheck = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersTCPListenerConfig) SetHealthCheckConnectPort(v int32) *DescribeLoadBalancerListenersResponseBodyListenersTCPListenerConfig {
	s.HealthCheckConnectPort = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersTCPListenerConfig) SetHealthCheckConnectTimeout(v int32) *DescribeLoadBalancerListenersResponseBodyListenersTCPListenerConfig {
	s.HealthCheckConnectTimeout = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersTCPListenerConfig) SetHealthCheckDomain(v string) *DescribeLoadBalancerListenersResponseBodyListenersTCPListenerConfig {
	s.HealthCheckDomain = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersTCPListenerConfig) SetHealthCheckHttpCode(v string) *DescribeLoadBalancerListenersResponseBodyListenersTCPListenerConfig {
	s.HealthCheckHttpCode = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersTCPListenerConfig) SetHealthCheckInterval(v int32) *DescribeLoadBalancerListenersResponseBodyListenersTCPListenerConfig {
	s.HealthCheckInterval = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersTCPListenerConfig) SetHealthCheckMethod(v string) *DescribeLoadBalancerListenersResponseBodyListenersTCPListenerConfig {
	s.HealthCheckMethod = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersTCPListenerConfig) SetHealthCheckType(v string) *DescribeLoadBalancerListenersResponseBodyListenersTCPListenerConfig {
	s.HealthCheckType = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersTCPListenerConfig) SetHealthCheckURI(v string) *DescribeLoadBalancerListenersResponseBodyListenersTCPListenerConfig {
	s.HealthCheckURI = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersTCPListenerConfig) SetHealthyThreshold(v int32) *DescribeLoadBalancerListenersResponseBodyListenersTCPListenerConfig {
	s.HealthyThreshold = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersTCPListenerConfig) SetMasterSlaveServerGroupId(v string) *DescribeLoadBalancerListenersResponseBodyListenersTCPListenerConfig {
	s.MasterSlaveServerGroupId = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersTCPListenerConfig) SetPersistenceTimeout(v int32) *DescribeLoadBalancerListenersResponseBodyListenersTCPListenerConfig {
	s.PersistenceTimeout = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersTCPListenerConfig) SetProxyProtocolV2Enabled(v string) *DescribeLoadBalancerListenersResponseBodyListenersTCPListenerConfig {
	s.ProxyProtocolV2Enabled = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersTCPListenerConfig) SetUnhealthyThreshold(v int32) *DescribeLoadBalancerListenersResponseBodyListenersTCPListenerConfig {
	s.UnhealthyThreshold = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersTCPListenerConfig) Validate() error {
	return dara.Validate(s)
}

type DescribeLoadBalancerListenersResponseBodyListenersTags struct {
	// The tag key.
	//
	// example:
	//
	// TestKey
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value.
	//
	// example:
	//
	// TestValue
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s DescribeLoadBalancerListenersResponseBodyListenersTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeLoadBalancerListenersResponseBodyListenersTags) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersTags) GetTagKey() *string {
	return s.TagKey
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersTags) GetTagValue() *string {
	return s.TagValue
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersTags) SetTagKey(v string) *DescribeLoadBalancerListenersResponseBodyListenersTags {
	s.TagKey = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersTags) SetTagValue(v string) *DescribeLoadBalancerListenersResponseBodyListenersTags {
	s.TagValue = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersTags) Validate() error {
	return dara.Validate(s)
}

type DescribeLoadBalancerListenersResponseBodyListenersUDPListenerConfig struct {
	// Indicates whether connection draining is enabled. Valid values:
	//
	// 	- **on**
	//
	// 	- **off**
	//
	// example:
	//
	// off
	ConnectionDrain *string `json:"ConnectionDrain,omitempty" xml:"ConnectionDrain,omitempty"`
	// The timeout period of connection draining. Unit: seconds.
	//
	// Value values: **10 to 900**.
	//
	// example:
	//
	// 300
	ConnectionDrainTimeout *int32 `json:"ConnectionDrainTimeout,omitempty" xml:"ConnectionDrainTimeout,omitempty"`
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
	// example:
	//
	// 8080
	HealthCheckConnectPort *int32 `json:"HealthCheckConnectPort,omitempty" xml:"HealthCheckConnectPort,omitempty"`
	// The timeout period for a health check response.
	//
	// example:
	//
	// 100
	HealthCheckConnectTimeout *int32 `json:"HealthCheckConnectTimeout,omitempty" xml:"HealthCheckConnectTimeout,omitempty"`
	// The response string of UDP health checks.
	//
	// example:
	//
	// ok
	HealthCheckExp *string `json:"HealthCheckExp,omitempty" xml:"HealthCheckExp,omitempty"`
	// The interval between two consecutive health checks. Unit: seconds.
	//
	// example:
	//
	// 5
	HealthCheckInterval *int32 `json:"HealthCheckInterval,omitempty" xml:"HealthCheckInterval,omitempty"`
	// The request string of UDP health checks.
	//
	// example:
	//
	// hello
	HealthCheckReq *string `json:"HealthCheckReq,omitempty" xml:"HealthCheckReq,omitempty"`
	// The number of times that a backend server must consecutively pass health checks before it is declared healthy.
	//
	// example:
	//
	// 4
	HealthyThreshold *int32 `json:"HealthyThreshold,omitempty" xml:"HealthyThreshold,omitempty"`
	// The ID of the primary/secondary server group that is associated with the listener.
	//
	// example:
	//
	// rsp-0bfucw****
	MasterSlaveServerGroupId *string `json:"MasterSlaveServerGroupId,omitempty" xml:"MasterSlaveServerGroupId,omitempty"`
	// Indicates whether the Proxy protocol is used to pass source client IP addresses to backend servers. Valid values:
	//
	// 	- **true**: enables the burst feature for the data disk.
	//
	// 	- **false**: The task is not being retried.
	//
	// example:
	//
	// false
	ProxyProtocolV2Enabled *string `json:"ProxyProtocolV2Enabled,omitempty" xml:"ProxyProtocolV2Enabled,omitempty"`
	// The number of times that a backend server must consecutively fail health checks before it is declared unhealthy.
	//
	// example:
	//
	// 4
	UnhealthyThreshold *int32 `json:"UnhealthyThreshold,omitempty" xml:"UnhealthyThreshold,omitempty"`
}

func (s DescribeLoadBalancerListenersResponseBodyListenersUDPListenerConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeLoadBalancerListenersResponseBodyListenersUDPListenerConfig) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersUDPListenerConfig) GetConnectionDrain() *string {
	return s.ConnectionDrain
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersUDPListenerConfig) GetConnectionDrainTimeout() *int32 {
	return s.ConnectionDrainTimeout
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersUDPListenerConfig) GetHealthCheck() *string {
	return s.HealthCheck
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersUDPListenerConfig) GetHealthCheckConnectPort() *int32 {
	return s.HealthCheckConnectPort
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersUDPListenerConfig) GetHealthCheckConnectTimeout() *int32 {
	return s.HealthCheckConnectTimeout
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersUDPListenerConfig) GetHealthCheckExp() *string {
	return s.HealthCheckExp
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersUDPListenerConfig) GetHealthCheckInterval() *int32 {
	return s.HealthCheckInterval
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersUDPListenerConfig) GetHealthCheckReq() *string {
	return s.HealthCheckReq
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersUDPListenerConfig) GetHealthyThreshold() *int32 {
	return s.HealthyThreshold
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersUDPListenerConfig) GetMasterSlaveServerGroupId() *string {
	return s.MasterSlaveServerGroupId
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersUDPListenerConfig) GetProxyProtocolV2Enabled() *string {
	return s.ProxyProtocolV2Enabled
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersUDPListenerConfig) GetUnhealthyThreshold() *int32 {
	return s.UnhealthyThreshold
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersUDPListenerConfig) SetConnectionDrain(v string) *DescribeLoadBalancerListenersResponseBodyListenersUDPListenerConfig {
	s.ConnectionDrain = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersUDPListenerConfig) SetConnectionDrainTimeout(v int32) *DescribeLoadBalancerListenersResponseBodyListenersUDPListenerConfig {
	s.ConnectionDrainTimeout = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersUDPListenerConfig) SetHealthCheck(v string) *DescribeLoadBalancerListenersResponseBodyListenersUDPListenerConfig {
	s.HealthCheck = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersUDPListenerConfig) SetHealthCheckConnectPort(v int32) *DescribeLoadBalancerListenersResponseBodyListenersUDPListenerConfig {
	s.HealthCheckConnectPort = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersUDPListenerConfig) SetHealthCheckConnectTimeout(v int32) *DescribeLoadBalancerListenersResponseBodyListenersUDPListenerConfig {
	s.HealthCheckConnectTimeout = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersUDPListenerConfig) SetHealthCheckExp(v string) *DescribeLoadBalancerListenersResponseBodyListenersUDPListenerConfig {
	s.HealthCheckExp = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersUDPListenerConfig) SetHealthCheckInterval(v int32) *DescribeLoadBalancerListenersResponseBodyListenersUDPListenerConfig {
	s.HealthCheckInterval = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersUDPListenerConfig) SetHealthCheckReq(v string) *DescribeLoadBalancerListenersResponseBodyListenersUDPListenerConfig {
	s.HealthCheckReq = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersUDPListenerConfig) SetHealthyThreshold(v int32) *DescribeLoadBalancerListenersResponseBodyListenersUDPListenerConfig {
	s.HealthyThreshold = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersUDPListenerConfig) SetMasterSlaveServerGroupId(v string) *DescribeLoadBalancerListenersResponseBodyListenersUDPListenerConfig {
	s.MasterSlaveServerGroupId = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersUDPListenerConfig) SetProxyProtocolV2Enabled(v string) *DescribeLoadBalancerListenersResponseBodyListenersUDPListenerConfig {
	s.ProxyProtocolV2Enabled = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersUDPListenerConfig) SetUnhealthyThreshold(v int32) *DescribeLoadBalancerListenersResponseBodyListenersUDPListenerConfig {
	s.UnhealthyThreshold = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersUDPListenerConfig) Validate() error {
	return dara.Validate(s)
}
