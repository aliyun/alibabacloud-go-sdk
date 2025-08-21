// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLoadBalancerHTTPSListenerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackendServerPort(v int32) *CreateLoadBalancerHTTPSListenerRequest
	GetBackendServerPort() *int32
	SetCookie(v string) *CreateLoadBalancerHTTPSListenerRequest
	GetCookie() *string
	SetCookieTimeout(v int32) *CreateLoadBalancerHTTPSListenerRequest
	GetCookieTimeout() *int32
	SetDescription(v string) *CreateLoadBalancerHTTPSListenerRequest
	GetDescription() *string
	SetForwardPort(v int32) *CreateLoadBalancerHTTPSListenerRequest
	GetForwardPort() *int32
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
	SetListenerForward(v string) *CreateLoadBalancerHTTPSListenerRequest
	GetListenerForward() *string
	SetListenerPort(v int32) *CreateLoadBalancerHTTPSListenerRequest
	GetListenerPort() *int32
	SetLoadBalancerId(v string) *CreateLoadBalancerHTTPSListenerRequest
	GetLoadBalancerId() *string
	SetRequestTimeout(v int32) *CreateLoadBalancerHTTPSListenerRequest
	GetRequestTimeout() *int32
	SetScheduler(v string) *CreateLoadBalancerHTTPSListenerRequest
	GetScheduler() *string
	SetServerCertificateId(v string) *CreateLoadBalancerHTTPSListenerRequest
	GetServerCertificateId() *string
	SetStickySessionType(v string) *CreateLoadBalancerHTTPSListenerRequest
	GetStickySessionType() *string
	SetUnhealthyThreshold(v int32) *CreateLoadBalancerHTTPSListenerRequest
	GetUnhealthyThreshold() *int32
}

type CreateLoadBalancerHTTPSListenerRequest struct {
	// The backend port that is used by the ELB instance. Valid values: **1*	- to **65535**.
	//
	// example:
	//
	// 8080
	BackendServerPort *int32 `json:"BackendServerPort,omitempty" xml:"BackendServerPort,omitempty"`
	// The cookie that is configured on the server. The cookie must be **1*	- to **200*	- characters in length and contain only ASCII characters and digits.
	//
	// >  This parameter is required if you set StickySession to on and StickySessionType to server.
	//
	// example:
	//
	// example
	Cookie *string `json:"Cookie,omitempty" xml:"Cookie,omitempty"`
	// The timeout period of a cookie. Valid values: **1*	- to **86400**. Unit: seconds.
	//
	// >  This parameter is required if you set StickySession to on and StickySessionType to insert.
	//
	// example:
	//
	// 100
	CookieTimeout *int32 `json:"CookieTimeout,omitempty" xml:"CookieTimeout,omitempty"`
	// The description of the listener. The description must be **1*	- to **80*	- characters in length.
	//
	// >  The value cannot start with `http://` or `https://`.
	//
	// example:
	//
	// https_80
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The listener port that is used to redirect HTTP requests to HTTPS.
	//
	// example:
	//
	// 0
	ForwardPort *int32 `json:"ForwardPort,omitempty" xml:"ForwardPort,omitempty"`
	// Indicates whether the health check feature is enabled. Valid values:
	//
	// 	- **on**
	//
	// 	- **off**
	//
	// This parameter is required.
	//
	// example:
	//
	// on
	HealthCheck *string `json:"HealthCheck,omitempty" xml:"HealthCheck,omitempty"`
	// The port that is used for health checks. Valid values: **1*	- to **65535**. If you leave this parameter empty, the port specified by BackendServerPort is used for health checks.
	//
	// >  This parameter takes effect only if you set HealthCheck to on.
	//
	// example:
	//
	// 11
	HealthCheckConnectPort *int32 `json:"HealthCheckConnectPort,omitempty" xml:"HealthCheckConnectPort,omitempty"`
	// The domain name that you want to use for health checks.
	//
	// >  This parameter takes effect only if you set HealthCheck to on.
	//
	// example:
	//
	// example.com
	HealthCheckDomain *string `json:"HealthCheckDomain,omitempty" xml:"HealthCheckDomain,omitempty"`
	// The HTTP status code for a successful health check. Valid values:
	//
	// 	- **http_2xx*	- (default)
	//
	// 	- **http_3xx**
	//
	// 	- **http_4xx**
	//
	// 	- **http_5xx**
	//
	// >  This parameter takes effect only if you set HealthCheck to on.
	//
	// example:
	//
	// http_2xx
	HealthCheckHttpCode *string `json:"HealthCheckHttpCode,omitempty" xml:"HealthCheckHttpCode,omitempty"`
	// The interval at which health checks are performed. Valid values: **1*	- to **50**. Default value: **2**. Unit: seconds.
	//
	// >  This parameter takes effect only if you set HealthCheck to on.
	//
	// example:
	//
	// 2
	HealthCheckInterval *int32 `json:"HealthCheckInterval,omitempty" xml:"HealthCheckInterval,omitempty"`
	// The HTTP request method for health checks. Valid values:
	//
	// 	- **head*	- (default): requests the head of the page.
	//
	// 	- **get**: requests the specified part of the page and returns the entity body.
	//
	// >  This parameter takes effect only if the HealthCheck parameter is set to on.
	//
	// example:
	//
	// head
	HealthCheckMethod *string `json:"HealthCheckMethod,omitempty" xml:"HealthCheckMethod,omitempty"`
	// The timeout period of a health check response. If a backend server does not respond within the specified timeout period, the server fails to pass the health check.
	//
	// 	- Default value: 5.
	//
	// 	- Valid values: **1*	- to **300**.
	//
	// 	- Unit: seconds.
	//
	// >
	//
	// 	- This parameter takes effect only if the HealthCheck parameter is set to on.
	//
	// 	- If the value of HealthCheckTimeout is smaller than the value of HealthCheckInterval, the timeout period specified by HealthCheckTimeout becomes invalid, and the value of HealthCheckInterval is used as the timeout period.
	//
	// example:
	//
	// 5
	HealthCheckTimeout *int32 `json:"HealthCheckTimeout,omitempty" xml:"HealthCheckTimeout,omitempty"`
	// The URI used for health checks. The URI must be **1*	- to **80*	- characters in length.
	//
	// >  A URL must start with a forward slash (`/`) but cannot contain only forward slashes (`/`).
	//
	// example:
	//
	// /checkpreload.htm
	HealthCheckURI *string `json:"HealthCheckURI,omitempty" xml:"HealthCheckURI,omitempty"`
	// The number of consecutive successful health checks that must occur before an unhealthy and inaccessible backend server is declared healthy and accessible. Valid values: **2*	- to **10**. Default value: **3**.
	//
	// >  This parameter takes effect only if you set HealthCheck to on.
	//
	// example:
	//
	// 3
	HealthyThreshold *int32 `json:"HealthyThreshold,omitempty" xml:"HealthyThreshold,omitempty"`
	// The timeout period for idle connections. Default value: 15. Valid values: **1*	- to **60**. Unit: seconds.
	//
	// >  If no request is received within the specified timeout period, ELB closes the connection. When another request is received, ELB establishes a new connection.
	//
	// example:
	//
	// 15
	IdleTimeout *int32 `json:"IdleTimeout,omitempty" xml:"IdleTimeout,omitempty"`
	// Specifies whether to enable redirection from HTTP to HTTPS. Valid values:
	//
	// 	- **on**
	//
	// 	- **off*	- (default)
	//
	// example:
	//
	// off
	ListenerForward *string `json:"ListenerForward,omitempty" xml:"ListenerForward,omitempty"`
	// The listening port that is used by Edge Load Balancer (ELB) to receive requests and forward the requests to backend servers. Valid values: **1*	- to **65535**.
	//
	// >  We recommend that you use port 443 for HTTPS.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8080
	ListenerPort *int32 `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	// The ID of the Edge Load Balancer (ELB) instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// lb-5s8w63yydi59w7klaikam****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// The timeout period of requests. Default value: 60. Valid values: **1*	- to **180**. Unit: seconds.
	//
	// >  If no response is received from the backend server within the specified timeout period, ELB returns an HTTP 504 error code to the client.
	//
	// example:
	//
	// 60
	RequestTimeout *int32 `json:"RequestTimeout,omitempty" xml:"RequestTimeout,omitempty"`
	// The scheduling algorithm. Valid values:
	//
	// 	- **wrr*	- (default): Backend servers with higher weights receive more requests than backend servers with lower weights.
	//
	// 	- **wlc**: Requests are distributed based on the weight and load of each backend server. The load refers to the number of connections on a backend server. If two backend servers have the same weight, the backend server that has fewer connections receives more requests.
	//
	// 	- **rr**: Requests are distributed to backend servers in sequence.
	//
	// 	- **sch**: consistent hashing based on source IP addresses. Requests from the same source IP address are distributed to the same backend server.
	//
	// 	- **qch**: consistent hashing based on QUIC connection IDs (CIDs). Requests that contain the same QUIC CID are distributed to the same backend server.
	//
	// 	- **iqch**: consistent hashing based on three specific bytes of iQUIC CIDs. Requests with the same second, third, and fourth bytes are distributed to the same backend server.
	//
	// example:
	//
	// wrr
	Scheduler *string `json:"Scheduler,omitempty" xml:"Scheduler,omitempty"`
	// The ID of the server certificate.
	//
	// This parameter is required.
	//
	// example:
	//
	// 60276**
	ServerCertificateId *string `json:"ServerCertificateId,omitempty" xml:"ServerCertificateId,omitempty"`
	// The method that is used to handle cookies. Valid values:
	//
	// 	- **insert**: inserts a cookie. ELB inserts a session cookie (SERVERID) into the first HTTP or HTTPS response that is sent to a client. Subsequent requests to ELB carry this cookie, and ELB determines the destination servers of the requests based on the cookies.
	//
	// 	- **server**: rewrites the original cookie. SLB rewrites the custom cookies in requests from a client. Subsequent requests from the client that carry the new cookie are forwarded to the same backend server as the first request.
	//
	// >  This parameter is required when the StickySession parameter is set to on.
	//
	// example:
	//
	// insert
	StickySessionType *string `json:"StickySessionType,omitempty" xml:"StickySessionType,omitempty"`
	// The number of consecutive failed health checks that must occur before a healthy and accessible backend server is declared unhealthy and inaccessible. Valid values: **2*	- to **10**. Default value: **3**.
	//
	// >  This parameter takes effect only if you set HealthCheck to on.
	//
	// example:
	//
	// 3
	UnhealthyThreshold *int32 `json:"UnhealthyThreshold,omitempty" xml:"UnhealthyThreshold,omitempty"`
}

func (s CreateLoadBalancerHTTPSListenerRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateLoadBalancerHTTPSListenerRequest) GoString() string {
	return s.String()
}

func (s *CreateLoadBalancerHTTPSListenerRequest) GetBackendServerPort() *int32 {
	return s.BackendServerPort
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

func (s *CreateLoadBalancerHTTPSListenerRequest) GetForwardPort() *int32 {
	return s.ForwardPort
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

func (s *CreateLoadBalancerHTTPSListenerRequest) GetListenerForward() *string {
	return s.ListenerForward
}

func (s *CreateLoadBalancerHTTPSListenerRequest) GetListenerPort() *int32 {
	return s.ListenerPort
}

func (s *CreateLoadBalancerHTTPSListenerRequest) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *CreateLoadBalancerHTTPSListenerRequest) GetRequestTimeout() *int32 {
	return s.RequestTimeout
}

func (s *CreateLoadBalancerHTTPSListenerRequest) GetScheduler() *string {
	return s.Scheduler
}

func (s *CreateLoadBalancerHTTPSListenerRequest) GetServerCertificateId() *string {
	return s.ServerCertificateId
}

func (s *CreateLoadBalancerHTTPSListenerRequest) GetStickySessionType() *string {
	return s.StickySessionType
}

func (s *CreateLoadBalancerHTTPSListenerRequest) GetUnhealthyThreshold() *int32 {
	return s.UnhealthyThreshold
}

func (s *CreateLoadBalancerHTTPSListenerRequest) SetBackendServerPort(v int32) *CreateLoadBalancerHTTPSListenerRequest {
	s.BackendServerPort = &v
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

func (s *CreateLoadBalancerHTTPSListenerRequest) SetForwardPort(v int32) *CreateLoadBalancerHTTPSListenerRequest {
	s.ForwardPort = &v
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

func (s *CreateLoadBalancerHTTPSListenerRequest) SetListenerForward(v string) *CreateLoadBalancerHTTPSListenerRequest {
	s.ListenerForward = &v
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

func (s *CreateLoadBalancerHTTPSListenerRequest) SetRequestTimeout(v int32) *CreateLoadBalancerHTTPSListenerRequest {
	s.RequestTimeout = &v
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

func (s *CreateLoadBalancerHTTPSListenerRequest) SetStickySessionType(v string) *CreateLoadBalancerHTTPSListenerRequest {
	s.StickySessionType = &v
	return s
}

func (s *CreateLoadBalancerHTTPSListenerRequest) SetUnhealthyThreshold(v int32) *CreateLoadBalancerHTTPSListenerRequest {
	s.UnhealthyThreshold = &v
	return s
}

func (s *CreateLoadBalancerHTTPSListenerRequest) Validate() error {
	return dara.Validate(s)
}
