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
	// The listener name. The name must be **1*	- to **80*	- characters in length.
	//
	// > The name cannot start with `http://` or `https://`.
	//
	// example:
	//
	// Listener description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to enable health checks. Valid values:
	//
	// - **on**: enabled.
	//
	// - **off**: disabled.
	//
	// example:
	//
	// on
	HealthCheck *string `json:"HealthCheck,omitempty" xml:"HealthCheck,omitempty"`
	// The port used for health checks. Valid values: **1*	- to **65535**. If this parameter is not configured in Settings, the backend service port (BackendServerPort) is used.
	//
	// > This parameter takes effect only when HealthCheck is set to on.
	//
	// example:
	//
	// 65500
	HealthCheckConnectPort *int32 `json:"HealthCheckConnectPort,omitempty" xml:"HealthCheckConnectPort,omitempty"`
	// The domain name used for health checks.
	//
	// > This parameter takes effect only when HealthCheck is set to on.
	//
	// example:
	//
	// www.aliyundoc.com
	HealthCheckDomain *string `json:"HealthCheckDomain,omitempty" xml:"HealthCheckDomain,omitempty"`
	// The HTTP status code that indicates a successful health check. Valid values:
	//
	// - **http_2xx*	- (default).
	//
	// - **http_3xx**.
	//
	// - **http_4xx**.
	//
	// - **http_5xx**.
	//
	// > This parameter takes effect only when HealthCheck is set to on.
	//
	// example:
	//
	// http_2xx
	HealthCheckHttpCode *string `json:"HealthCheckHttpCode,omitempty" xml:"HealthCheckHttpCode,omitempty"`
	// The interval between two consecutive health checks. Valid values: **1*	- to **50**. Unit: seconds.
	//
	// > This parameter takes effect only when HealthCheck is set to on.
	//
	// example:
	//
	// 2
	HealthCheckInterval *int32 `json:"HealthCheckInterval,omitempty" xml:"HealthCheckInterval,omitempty"`
	// The health check method for the HTTP listener. Valid values:
	//
	// - **head**.
	//
	// - **get**.
	//
	// > This parameter takes effect only when HealthCheck is set to on.
	//
	// example:
	//
	// head
	HealthCheckMethod *string `json:"HealthCheckMethod,omitempty" xml:"HealthCheckMethod,omitempty"`
	// The amount of time to wait for a response from a health check. If the backend ECS instance does not respond within the specified time, the health check fails.
	//
	// - Default value: 5 seconds.
	//
	// - Valid values: **1*	- to **300**.
	//
	// - Unit: seconds.
	//
	// > - This parameter takes effect only when HealthCheck is set to on.
	//
	// >- If the value of HealthCheckTimeout is less than the value of HealthCheckInterval, HealthCheckTimeout becomes ineffective and the timeout period equals the value of HealthCheckInterval.
	//
	// example:
	//
	// 5
	HealthCheckTimeout *int32 `json:"HealthCheckTimeout,omitempty" xml:"HealthCheckTimeout,omitempty"`
	// The URI used for health checks. The URI must be **1*	- to **80*	- characters in length.
	//
	// > - The URI cannot be only `/` but must start with `/`.
	//
	// >- This parameter takes effect only when HealthCheck is set to on.
	//
	// example:
	//
	// /checkpreload.htm
	HealthCheckURI *string `json:"HealthCheckURI,omitempty" xml:"HealthCheckURI,omitempty"`
	// The number of consecutive successful health checks required before a backend server is declared healthy (changing the status from fail to success). Valid values: **2*	- to **10**.
	//
	// > This parameter takes effect only when HealthCheck is set to on.
	//
	// example:
	//
	// 3
	HealthyThreshold *int32 `json:"HealthyThreshold,omitempty" xml:"HealthyThreshold,omitempty"`
	// The idle connection timeout period. Default value: 15 seconds. Valid values: **1*	- to **60**. Unit: seconds.
	//
	// > If no access request is received within the timeout period, the load balancing instance temporarily breaks the current connection and re-establishes a new connection when the next request arrives.
	//
	// example:
	//
	// 15
	IdleTimeout *int32 `json:"IdleTimeout,omitempty" xml:"IdleTimeout,omitempty"`
	// The listener port whose attributes you want to modify. Valid values: **1*	- to **65535**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8080
	ListenerPort *int32 `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	// The ID of the load balancing instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// lb-5snthcyu1x10g7tywj7iu****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// The request timeout period. Default value: 60 seconds. Valid values: **1*	- to **180**. Unit: seconds.
	//
	// > If no response is received from the backend server within the timeout period, the load balancing instance stops waiting and returns an HTTP 504 error code to the client.
	//
	// example:
	//
	// 60
	RequestTimeout *int32 `json:"RequestTimeout,omitempty" xml:"RequestTimeout,omitempty"`
	// The scheduling algorithm. Valid values:
	//
	// - **wrr**: Backend servers with higher weights are polled more frequently.
	//
	// - **wlc**: Requests are distributed based on the weight and actual load (number of connections) of each backend server. When the weights are the same, backend servers with fewer connections are polled more frequently.
	//
	// - **rr**: Requests are distributed to backend servers in sequence.
	//
	// - **sch**: Consistent hashing based on source IP addresses. Requests from the same source IP address are distributed to the same backend server.
	//
	// - **qch**: Consistent hashing based on QUIC Connection IDs. Requests with the same QUIC Connection ID are distributed to the same backend server.
	//
	// - **iqch**: Consistent hashing based on specific three bytes of the iQUIC CID. Requests with the same 2nd to 4th bytes are distributed to the same backend server.
	//
	// example:
	//
	// wrr
	Scheduler *string `json:"Scheduler,omitempty" xml:"Scheduler,omitempty"`
	// The number of consecutive failed health checks required before a backend server is declared unhealthy (changing the status from success to fail). Valid values: **2*	- to **10**.
	//
	// > This parameter takes effect only when HealthCheck is set to on.
	//
	// example:
	//
	// 3
	UnhealthyThreshold *int32 `json:"UnhealthyThreshold,omitempty" xml:"UnhealthyThreshold,omitempty"`
	// Specifies whether to use the X-Forwarded-For header to retrieve the originating IP address of the client. Valid values:
	//
	// - **on*	- (default): enabled.
	//
	// - **off**: disabled.
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
