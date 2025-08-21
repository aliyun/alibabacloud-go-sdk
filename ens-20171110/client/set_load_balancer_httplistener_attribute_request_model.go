// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetLoadBalancerHTTPListenerAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *SetLoadBalancerHTTPListenerAttributeRequest
	GetDescription() *string
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
	SetRequestTimeout(v int32) *SetLoadBalancerHTTPListenerAttributeRequest
	GetRequestTimeout() *int32
	SetScheduler(v string) *SetLoadBalancerHTTPListenerAttributeRequest
	GetScheduler() *string
	SetUnhealthyThreshold(v int32) *SetLoadBalancerHTTPListenerAttributeRequest
	GetUnhealthyThreshold() *int32
	SetXForwardedFor(v string) *SetLoadBalancerHTTPListenerAttributeRequest
	GetXForwardedFor() *string
}

type SetLoadBalancerHTTPListenerAttributeRequest struct {
	// The name of the listener. The value must be **1*	- to **80*	- characters in length.
	//
	// >  The value cannot start with `http://` or `https://`.
	//
	// example:
	//
	// http_8080
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
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
	// The port that is used for health checks. Valid values: **1*	- to **65535**. If you leave this parameter empty, the port specified by BackendServerPort is used for health checks.
	//
	// >  This parameter takes effect only if you set HealthCheck to on.
	//
	// example:
	//
	// 65500
	HealthCheckConnectPort *int32 `json:"HealthCheckConnectPort,omitempty" xml:"HealthCheckConnectPort,omitempty"`
	// The domain name that is used for health checks.
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
	// 	- **http_3xx**.
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
	// The interval at which health checks are performed. Valid values: **1*	- to **50**. Unit: seconds.
	//
	// >  This parameter takes effect only if you set HealthCheck to on.
	//
	// example:
	//
	// 2
	HealthCheckInterval *int32 `json:"HealthCheckInterval,omitempty" xml:"HealthCheckInterval,omitempty"`
	// The HTTP request method for health checks. Valid values:
	//
	// 	- **head**
	//
	// 	- **get**
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
	// The number of consecutive successful health checks that must occur before an unhealthy and inaccessible backend server is declared healthy and accessible. Valid values: **2*	- to **10**.
	//
	// >  This parameter takes effect only if you set HealthCheck to on.
	//
	// example:
	//
	// 3
	HealthyThreshold *int32 `json:"HealthyThreshold,omitempty" xml:"HealthyThreshold,omitempty"`
	// The timeout period for idle connections. Default value: 15. Valid values: **1*	- to **60**. Unit: seconds.
	//
	// >  If no request is received within the specified timeout period, ELB closes the connection. When another request is received, CLB establishes a new connection.
	//
	// example:
	//
	// 15
	IdleTimeout *int32 `json:"IdleTimeout,omitempty" xml:"IdleTimeout,omitempty"`
	// The listener port whose attributes are to be modified. Valid values: **1*	- to **65535**.
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
	// lb-5snthcyu1x10g7tywj7iu****
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
	// 	- **wrr**: Backend servers with higher weights receive more requests than those with lower weights.
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
	// The number of consecutive failed health checks that must occur before a healthy and accessible backend server is declared unhealthy and inaccessible. Valid values: **2*	- to **10**.
	//
	// >  This parameter takes effect only if you set HealthCheck to on.
	//
	// example:
	//
	// 3
	UnhealthyThreshold *int32 `json:"UnhealthyThreshold,omitempty" xml:"UnhealthyThreshold,omitempty"`
	// Specifies whether to use the X-Forwarded-For header to obtain the real IP address of the client. Valid values:
	//
	// 	- **on*	- (default)
	//
	// 	- **off**
	//
	// example:
	//
	// on
	XForwardedFor *string `json:"XForwardedFor,omitempty" xml:"XForwardedFor,omitempty"`
}

func (s SetLoadBalancerHTTPListenerAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s SetLoadBalancerHTTPListenerAttributeRequest) GoString() string {
	return s.String()
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) GetDescription() *string {
	return s.Description
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

func (s *SetLoadBalancerHTTPListenerAttributeRequest) GetRequestTimeout() *int32 {
	return s.RequestTimeout
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) GetScheduler() *string {
	return s.Scheduler
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) GetUnhealthyThreshold() *int32 {
	return s.UnhealthyThreshold
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) GetXForwardedFor() *string {
	return s.XForwardedFor
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) SetDescription(v string) *SetLoadBalancerHTTPListenerAttributeRequest {
	s.Description = &v
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

func (s *SetLoadBalancerHTTPListenerAttributeRequest) SetRequestTimeout(v int32) *SetLoadBalancerHTTPListenerAttributeRequest {
	s.RequestTimeout = &v
	return s
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) SetScheduler(v string) *SetLoadBalancerHTTPListenerAttributeRequest {
	s.Scheduler = &v
	return s
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) SetUnhealthyThreshold(v int32) *SetLoadBalancerHTTPListenerAttributeRequest {
	s.UnhealthyThreshold = &v
	return s
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) SetXForwardedFor(v string) *SetLoadBalancerHTTPListenerAttributeRequest {
	s.XForwardedFor = &v
	return s
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) Validate() error {
	return dara.Validate(s)
}
