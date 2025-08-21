// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLoadBalancerHTTPListenerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackendServerPort(v int32) *CreateLoadBalancerHTTPListenerRequest
	GetBackendServerPort() *int32
	SetDescription(v string) *CreateLoadBalancerHTTPListenerRequest
	GetDescription() *string
	SetForwardPort(v int32) *CreateLoadBalancerHTTPListenerRequest
	GetForwardPort() *int32
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
	SetRequestTimeout(v int32) *CreateLoadBalancerHTTPListenerRequest
	GetRequestTimeout() *int32
	SetScheduler(v string) *CreateLoadBalancerHTTPListenerRequest
	GetScheduler() *string
	SetUnhealthyThreshold(v int32) *CreateLoadBalancerHTTPListenerRequest
	GetUnhealthyThreshold() *int32
	SetXForwardedFor(v string) *CreateLoadBalancerHTTPListenerRequest
	GetXForwardedFor() *string
}

type CreateLoadBalancerHTTPListenerRequest struct {
	// The port used by the backend server of the ELB instance. Valid values: **1*	- to **65535**.
	//
	// example:
	//
	// 8080
	BackendServerPort *int32 `json:"BackendServerPort,omitempty" xml:"BackendServerPort,omitempty"`
	// The name of the listener. The value must be **1*	- to **80*	- characters in length.
	//
	// >  The value cannot start with `http://` or `https://`.
	//
	// example:
	//
	// Monitoring instructions
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The listener port that is used to redirect HTTP requests to HTTPS.
	//
	// example:
	//
	// 0
	ForwardPort *int32 `json:"ForwardPort,omitempty" xml:"ForwardPort,omitempty"`
	// Specifies whether to enable the health check feature. Valid values:
	//
	// 	- **on**
	//
	// 	- **off*	- (default)
	//
	// This parameter is required.
	//
	// example:
	//
	// on
	HealthCheck *string `json:"HealthCheck,omitempty" xml:"HealthCheck,omitempty"`
	// The backend port that is used for health checks. Valid values: **1*	- to **65535**.
	//
	// >  This parameter takes effect only if you set HealthCheck to on.
	//
	// example:
	//
	// 30040
	HealthCheckConnectPort *int32 `json:"HealthCheckConnectPort,omitempty" xml:"HealthCheckConnectPort,omitempty"`
	// The domain name that you want to use for health checks.
	//
	// >  This parameter takes effect only if you set HealthCheck to on.
	//
	// example:
	//
	// www.aliyundoc.com
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
	// 	- **head*	- (default)
	//
	// 	- **get**
	//
	// >  This parameter takes effect only if the HealthCheck parameter is set to on.
	//
	// example:
	//
	// head
	HealthCheckMethod *string `json:"HealthCheckMethod,omitempty" xml:"HealthCheckMethod,omitempty"`
	// The timeout period of a health check response. If a backend server does not respond within the specified timeout period, the server fails the health check.
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
	// >
	//
	// 	- A URL must start with a forward slash (`/`) but cannot contain only forward slashes (`/`).
	//
	// 	- This parameter takes effect only if the HealthCheck parameter is set to on.
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
	// >  If no request is received within the specified timeout period, ELB closes the connection. When a request is received, ELB creates a new connection.
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
	// The listener port that is used by Edge Load Balancer (ELB) to receive requests and forward the requests to backend servers. Valid values: **1*	- to **65535**.
	//
	// >  We recommend that you use port 80 for HTTP.
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
	// lb-5s7crik3yo3bp03gqrbp5****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// The timeout period of a request. Default value: 60. Valid values: **1*	- to **180**. Unit: seconds.
	//
	// >  If no response is received from the backend server within the specified timeout period, ALB returns an HTTP 504 error code to the client.
	//
	// example:
	//
	// 60
	RequestTimeout *int32 `json:"RequestTimeout,omitempty" xml:"RequestTimeout,omitempty"`
	// The scheduling algorithm. Valid values:
	//
	// 	- **wrr*	- (default): Backend servers with higher weights receive more requests than backend servers with lower weights.
	//
	// 	- **wlc**: Requests are distributed based on the weights and number of connections to backend servers. If two backend servers have the same weight, the backend server that has fewer connections receives more requests.
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
	// The number of consecutive failed health checks that must occur before a healthy and accessible backend server is declared unhealthy and inaccessible. Valid values: **2*	- to **10**. Default value: **3**.
	//
	// >  This parameter takes effect only if you set HealthCheck to on.
	//
	// example:
	//
	// 3
	UnhealthyThreshold *int32 `json:"UnhealthyThreshold,omitempty" xml:"UnhealthyThreshold,omitempty"`
	// Specifies whether to use the X-Forwarded-For header to obtain the real IP address of the client. Valid values:
	//
	// 	- **on**
	//
	// 	- **off*	- (default)
	//
	// example:
	//
	// off
	XForwardedFor *string `json:"XForwardedFor,omitempty" xml:"XForwardedFor,omitempty"`
}

func (s CreateLoadBalancerHTTPListenerRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateLoadBalancerHTTPListenerRequest) GoString() string {
	return s.String()
}

func (s *CreateLoadBalancerHTTPListenerRequest) GetBackendServerPort() *int32 {
	return s.BackendServerPort
}

func (s *CreateLoadBalancerHTTPListenerRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateLoadBalancerHTTPListenerRequest) GetForwardPort() *int32 {
	return s.ForwardPort
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

func (s *CreateLoadBalancerHTTPListenerRequest) GetRequestTimeout() *int32 {
	return s.RequestTimeout
}

func (s *CreateLoadBalancerHTTPListenerRequest) GetScheduler() *string {
	return s.Scheduler
}

func (s *CreateLoadBalancerHTTPListenerRequest) GetUnhealthyThreshold() *int32 {
	return s.UnhealthyThreshold
}

func (s *CreateLoadBalancerHTTPListenerRequest) GetXForwardedFor() *string {
	return s.XForwardedFor
}

func (s *CreateLoadBalancerHTTPListenerRequest) SetBackendServerPort(v int32) *CreateLoadBalancerHTTPListenerRequest {
	s.BackendServerPort = &v
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

func (s *CreateLoadBalancerHTTPListenerRequest) SetRequestTimeout(v int32) *CreateLoadBalancerHTTPListenerRequest {
	s.RequestTimeout = &v
	return s
}

func (s *CreateLoadBalancerHTTPListenerRequest) SetScheduler(v string) *CreateLoadBalancerHTTPListenerRequest {
	s.Scheduler = &v
	return s
}

func (s *CreateLoadBalancerHTTPListenerRequest) SetUnhealthyThreshold(v int32) *CreateLoadBalancerHTTPListenerRequest {
	s.UnhealthyThreshold = &v
	return s
}

func (s *CreateLoadBalancerHTTPListenerRequest) SetXForwardedFor(v string) *CreateLoadBalancerHTTPListenerRequest {
	s.XForwardedFor = &v
	return s
}

func (s *CreateLoadBalancerHTTPListenerRequest) Validate() error {
	return dara.Validate(s)
}
