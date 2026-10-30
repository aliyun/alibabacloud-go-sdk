// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLoadBalancerHTTPListenerAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBackendServerPort(v int32) *DescribeLoadBalancerHTTPListenerAttributeResponseBody
	GetBackendServerPort() *int32
	SetBandwidth(v int32) *DescribeLoadBalancerHTTPListenerAttributeResponseBody
	GetBandwidth() *int32
	SetDescription(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBody
	GetDescription() *string
	SetForwardPort(v int32) *DescribeLoadBalancerHTTPListenerAttributeResponseBody
	GetForwardPort() *int32
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
	SetRequestId(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBody
	GetRequestId() *string
	SetRequestTimeout(v int32) *DescribeLoadBalancerHTTPListenerAttributeResponseBody
	GetRequestTimeout() *int32
	SetScheduler(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBody
	GetScheduler() *string
	SetServerCertificateId(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBody
	GetServerCertificateId() *string
	SetStatus(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBody
	GetStatus() *string
	SetUnhealthyThreshold(v int32) *DescribeLoadBalancerHTTPListenerAttributeResponseBody
	GetUnhealthyThreshold() *int32
	SetXForwardedFor(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBody
	GetXForwardedFor() *string
}

type DescribeLoadBalancerHTTPListenerAttributeResponseBody struct {
	// The port used by the backend server of the load balancing instance. Valid values: **1*	- to **65535**.
	//
	// example:
	//
	// 8080
	BackendServerPort *int32 `json:"BackendServerPort,omitempty" xml:"BackendServerPort,omitempty"`
	// The peak bandwidth of load balancing. Default value: -1, which indicates that bandwidth is not limited.
	//
	// example:
	//
	// -1
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The description of the listener. The description must be **1*	- to **80*	- characters in length.
	//
	// > It cannot start with `http://` or `https://`.
	//
	// example:
	//
	// abc
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The forwarding port for HTTP-to-HTTPS listeners.
	//
	// example:
	//
	// 10002
	ForwardPort *int32 `json:"ForwardPort,omitempty" xml:"ForwardPort,omitempty"`
	// Indicates whether health checks are enabled. Valid values:
	//
	// - **on**: enabled.
	//
	// - **off**: disabled.
	//
	// example:
	//
	// on
	HealthCheck *string `json:"HealthCheck,omitempty" xml:"HealthCheck,omitempty"`
	// The port used for health checks. Valid values: **1*	- to **65535**. If this parameter is not specified, the backend server port (BackendServerPort) is used.
	//
	// > This parameter takes effect only when HealthCheck is set to on.
	//
	// example:
	//
	// 5000
	HealthCheckConnectPort *int32 `json:"HealthCheckConnectPort,omitempty" xml:"HealthCheckConnectPort,omitempty"`
	// The domain name used for health checks.
	//
	// > This parameter takes effect only when HealthCheck is set to on.
	//
	// example:
	//
	// www.example.com
	HealthCheckDomain *string `json:"HealthCheckDomain,omitempty" xml:"HealthCheckDomain,omitempty"`
	// The HTTP status code that indicates a successful health check. Valid values:
	//
	// - **http_2xx*	- (default)
	//
	// - **http_3xx**
	//
	// - **http_4xx**
	//
	// - **http_5xx**
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
	// 5
	HealthCheckInterval *int32 `json:"HealthCheckInterval,omitempty" xml:"HealthCheckInterval,omitempty"`
	// The HTTP method used for health checks. Valid values:
	//
	// - **head**: Requests only the header of the page.
	//
	// - **get**: Requests the specified page information and returns the entity body.
	//
	// > This parameter takes effect only when HealthCheck is set to on.
	//
	// example:
	//
	// head
	HealthCheckMethod *string `json:"HealthCheckMethod,omitempty" xml:"HealthCheckMethod,omitempty"`
	// The timeout period for a health check response. If a backend ENS instance does not respond within the specified timeout period, the health check fails.
	//
	// - Default value: 5.
	//
	// - Valid values: **1*	- to **300**.
	//
	// - Unit: seconds.
	//
	// > - This parameter takes effect only when HealthCheck is set to on.
	//
	// > - If the value of HealthCheckTimeout is less than the value of HealthCheckInterval, HealthCheckTimeout becomes invalid and the timeout period equals the value of HealthCheckInterval.
	//
	// example:
	//
	// 10
	HealthCheckTimeout *int32 `json:"HealthCheckTimeout,omitempty" xml:"HealthCheckTimeout,omitempty"`
	// The URI used for health checks. The URI must be **1*	- to **80*	- characters in length.
	//
	// > - The URI cannot be only a forward slash (`/`), but it must start with a forward slash (`/`).
	//
	// > - This parameter takes effect only when HealthCheck is set to on.
	//
	// example:
	//
	// /checkpreload.htm
	HealthCheckURI *string `json:"HealthCheckURI,omitempty" xml:"HealthCheckURI,omitempty"`
	// The number of consecutive successful health checks that must occur before the health check status of a backend server is changed from fail (unreachable) to success (reachable). Valid values: **2*	- to **10**.
	//
	// > This parameter takes effect only when HealthCheck is set to on.
	//
	// example:
	//
	// 3
	HealthyThreshold *int32 `json:"HealthyThreshold,omitempty" xml:"HealthyThreshold,omitempty"`
	// The idle connection timeout period. Default value: 15. Valid values: **1*	- to **60**. Unit: seconds.
	//
	// > If no request is received within the specified timeout period, load balancing temporarily interrupts the current connection and establishes a new connection when the next request arrives.
	//
	// example:
	//
	// 15
	IdleTimeout *int32 `json:"IdleTimeout,omitempty" xml:"IdleTimeout,omitempty"`
	// Indicates whether HTTP-to-HTTPS forwarding is enabled. Valid values:
	//
	// - **on**: enabled.
	//
	// - **off*	- (default): disabled.
	//
	// example:
	//
	// off
	ListenerForward *string `json:"ListenerForward,omitempty" xml:"ListenerForward,omitempty"`
	// The listener port.
	//
	// example:
	//
	// 80
	ListenerPort *int32 `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The request timeout period. Default value: 60. Valid values: **1*	- to **180**. Unit: seconds.
	//
	// > If a backend server does not return a response within the specified timeout period, load balancing stops waiting and returns an HTTP 504 error code to the client.
	//
	// example:
	//
	// 60
	RequestTimeout *int32 `json:"RequestTimeout,omitempty" xml:"RequestTimeout,omitempty"`
	// The scheduling algorithm. Valid values:
	//
	// - **wrr**: Backend servers with higher weights receive more requests.
	//
	// - **wlc**: Requests are distributed based on the weights and actual loads (connections) of backend servers. If two backend servers have the same weight, the backend server with fewer connections receives more requests.
	//
	// - **rr**: Requests are evenly and sequentially distributed to backend servers.
	//
	// - **sch**: Consistent hashing based on source IP addresses. Requests from the same source IP address are distributed to the same backend server.
	//
	// - **qch**: Consistent hashing based on QUIC Connection IDs. Requests with the same QUIC Connection ID are distributed to the same backend server.
	//
	// - **iqch**: Consistent hashing based on specific three bytes of the iQUIC CID. Requests with the same second to fourth bytes are distributed to the same backend server.
	//
	// example:
	//
	// wrr
	Scheduler *string `json:"Scheduler,omitempty" xml:"Scheduler,omitempty"`
	// The ID of the server certificate.
	//
	// example:
	//
	// 6027667
	ServerCertificateId *string `json:"ServerCertificateId,omitempty" xml:"ServerCertificateId,omitempty"`
	// The status of the listener. Valid values:
	//
	// - **Running**: The listener is running.
	//
	// - **Stopped**: The listener is stopped.
	//
	// - **Starting**: The listener is starting.
	//
	// - **Configuring**: The listener is being configured.
	//
	// - **Stopping**: The listener is being stopped.
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The number of consecutive failed health checks that must occur before the health check status of a backend server is changed from success (reachable) to fail (unreachable). Valid values: **2*	- to **10**.
	//
	// > This parameter takes effect only when HealthCheck is set to on.
	//
	// example:
	//
	// 3
	UnhealthyThreshold *int32 `json:"UnhealthyThreshold,omitempty" xml:"UnhealthyThreshold,omitempty"`
	// Indicates whether the X-Forwarded-For header is used to retrieve the originating IP addresses of clients. Valid values:
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

func (s DescribeLoadBalancerHTTPListenerAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLoadBalancerHTTPListenerAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) GetBackendServerPort() *int32 {
	return s.BackendServerPort
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) GetDescription() *string {
	return s.Description
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) GetForwardPort() *int32 {
	return s.ForwardPort
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

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) GetRequestTimeout() *int32 {
	return s.RequestTimeout
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) GetScheduler() *string {
	return s.Scheduler
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) GetServerCertificateId() *string {
	return s.ServerCertificateId
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) GetUnhealthyThreshold() *int32 {
	return s.UnhealthyThreshold
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) GetXForwardedFor() *string {
	return s.XForwardedFor
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) SetBackendServerPort(v int32) *DescribeLoadBalancerHTTPListenerAttributeResponseBody {
	s.BackendServerPort = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) SetBandwidth(v int32) *DescribeLoadBalancerHTTPListenerAttributeResponseBody {
	s.Bandwidth = &v
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

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) SetRequestId(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) SetRequestTimeout(v int32) *DescribeLoadBalancerHTTPListenerAttributeResponseBody {
	s.RequestTimeout = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) SetScheduler(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBody {
	s.Scheduler = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) SetServerCertificateId(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBody {
	s.ServerCertificateId = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) SetStatus(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) SetUnhealthyThreshold(v int32) *DescribeLoadBalancerHTTPListenerAttributeResponseBody {
	s.UnhealthyThreshold = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) SetXForwardedFor(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBody {
	s.XForwardedFor = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
