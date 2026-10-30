// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLoadBalancerHTTPSListenerAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBackendServerPort(v int32) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody
	GetBackendServerPort() *int32
	SetBandwidth(v int32) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody
	GetBandwidth() *int32
	SetDescription(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody
	GetDescription() *string
	SetForwardPort(v int32) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody
	GetForwardPort() *int32
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
	SetListenerForward(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody
	GetListenerForward() *string
	SetListenerPort(v int32) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody
	GetListenerPort() *int32
	SetRequestId(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody
	GetRequestId() *string
	SetRequestTimeout(v int32) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody
	GetRequestTimeout() *int32
	SetScheduler(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody
	GetScheduler() *string
	SetServerCertificateId(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody
	GetServerCertificateId() *string
	SetStatus(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody
	GetStatus() *string
	SetUnhealthyThreshold(v int32) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody
	GetUnhealthyThreshold() *int32
}

type DescribeLoadBalancerHTTPSListenerAttributeResponseBody struct {
	// The port used by the backend server of the load balancer instance. Valid values: **1*	- to **65535**.
	//
	// example:
	//
	// 8080
	BackendServerPort *int32 `json:"BackendServerPort,omitempty" xml:"BackendServerPort,omitempty"`
	// The peak bandwidth of the load balancer. Default value: -1, which indicates that bandwidth is not limited.
	//
	// example:
	//
	// 50
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The description of the listener. The description must be **1*	- to **80*	- characters in length.
	//
	// example:
	//
	// abc
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The HTTP-to-HTTPS forwarding port for the listener.
	//
	// example:
	//
	// 0
	ForwardPort *int32 `json:"ForwardPort,omitempty" xml:"ForwardPort,omitempty"`
	// Specifies whether to enable the health check. Valid values:
	//
	// - **on**: Enables the health check.
	//
	// - **off**: Disables the health check.
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
	// 9902
	HealthCheckConnectPort *int32 `json:"HealthCheckConnectPort,omitempty" xml:"HealthCheckConnectPort,omitempty"`
	// The domain name used for health checks.
	//
	// > This parameter takes effect only when HealthCheck is set to on.
	//
	// example:
	//
	// www.test.com
	HealthCheckDomain *string `json:"HealthCheckDomain,omitempty" xml:"HealthCheckDomain,omitempty"`
	// The HTTP status codes that indicate a successful health check. Valid values:
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
	// The timeout period for receiving a response from a health check. If the backend ENS instance does not respond within the specified timeout period, the health check fails.
	//
	// - Default value: 5.
	//
	// - Valid values: **1*	- to **300**.
	//
	// - Unit: seconds.
	//
	// > - This parameter takes effect only when HealthCheck is set to on.
	//
	// > - If the value of HealthCheckTimeout is less than the value of HealthCheckInterval, HealthCheckTimeout is invalid and the timeout period is the value of HealthCheckInterval.
	//
	// example:
	//
	// 5
	HealthCheckTimeout *int32 `json:"HealthCheckTimeout,omitempty" xml:"HealthCheckTimeout,omitempty"`
	// The URI used for health checks. The URI must be **1*	- to **80*	- characters in length.
	//
	// - The URI cannot be only a forward slash (`/`), but must start with a forward slash.
	//
	// - This parameter takes effect only when HealthCheck is set to on.
	//
	// example:
	//
	// /checkpreload.htm
	HealthCheckURI *string `json:"HealthCheckURI,omitempty" xml:"HealthCheckURI,omitempty"`
	// The number of consecutive successful health checks required to change the health check status of a backend server from fail (unreachable) to success (reachable). Valid values: **2*	- to **10**.
	//
	// > This parameter takes effect only when HealthCheck is set to on.
	//
	// example:
	//
	// 3
	HealthyThreshold *int32 `json:"HealthyThreshold,omitempty" xml:"HealthyThreshold,omitempty"`
	// The idle connection timeout period. Default value: 15. Valid values: **1*	- to **60**. Unit: seconds.
	//
	// > If no access request is received within the timeout period, the load balancer temporarily interrupts the current connection and establishes a new connection when the next request arrives.
	//
	// example:
	//
	// 15
	IdleTimeout *int32 `json:"IdleTimeout,omitempty" xml:"IdleTimeout,omitempty"`
	// Specifies whether to enable HTTP-to-HTTPS forwarding. Valid values:
	//
	// - **on**: Enables HTTP-to-HTTPS forwarding.
	//
	// - **off**: Disables HTTP-to-HTTPS forwarding.
	//
	// example:
	//
	// off
	ListenerForward *string `json:"ListenerForward,omitempty" xml:"ListenerForward,omitempty"`
	// The frontend port used by the load balancer instance. Valid values: **1*	- to **65535**.
	//
	// example:
	//
	// 8080
	ListenerPort *int32 `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The request timeout period. Default value: 60. Valid values: **1*	- to **180**. Unit: seconds.
	//
	// > If no response is received from the backend server within the timeout period, the load balancer stops waiting and returns an HTTP 504 error code to the client.
	//
	// example:
	//
	// 60
	RequestTimeout *int32 `json:"RequestTimeout,omitempty" xml:"RequestTimeout,omitempty"`
	// The scheduling algorithm. Valid values:
	//
	// - **wrr**: Backend servers with higher weights are polled more frequently.
	//
	// - **wlc**: In addition to the weight of each backend server, the actual load (number of connections) of the backend server is also considered. If the weights are the same, backend servers with fewer current connections are polled more frequently.
	//
	// - **rr**: External requests are distributed to backend servers in sequence.
	//
	// - **sch**: Consistent hashing based on source IP addresses. Requests from the same source IP address are scheduled to the same backend server.
	//
	// - **qch**: Consistent hashing based on QUIC Connection IDs. Requests with the same QUIC Connection ID are scheduled to the same backend server.
	//
	// - **iqch**: Consistent hashing based on three specific bytes of the iQUIC CID. Requests with the same 2nd to 4th bytes are scheduled to the same backend server.
	//
	// example:
	//
	// wrr
	Scheduler *string `json:"Scheduler,omitempty" xml:"Scheduler,omitempty"`
	// The ID of the server certificate.
	//
	// example:
	//
	// 60276**
	ServerCertificateId *string `json:"ServerCertificateId,omitempty" xml:"ServerCertificateId,omitempty"`
	// The current status of the listener. Valid values:
	//
	// - **Running**: The listener is running normally.
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
	// The number of consecutive failed health checks required to change the health check status of a backend server from success (reachable) to fail (unreachable). Valid values: **2*	- to **10**.
	//
	// > This parameter takes effect only when HealthCheck is set to on.
	//
	// example:
	//
	// 3
	UnhealthyThreshold *int32 `json:"UnhealthyThreshold,omitempty" xml:"UnhealthyThreshold,omitempty"`
}

func (s DescribeLoadBalancerHTTPSListenerAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLoadBalancerHTTPSListenerAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) GetBackendServerPort() *int32 {
	return s.BackendServerPort
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) GetDescription() *string {
	return s.Description
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) GetForwardPort() *int32 {
	return s.ForwardPort
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

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) GetListenerForward() *string {
	return s.ListenerForward
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) GetListenerPort() *int32 {
	return s.ListenerPort
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) GetRequestTimeout() *int32 {
	return s.RequestTimeout
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) GetScheduler() *string {
	return s.Scheduler
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) GetServerCertificateId() *string {
	return s.ServerCertificateId
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) GetUnhealthyThreshold() *int32 {
	return s.UnhealthyThreshold
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) SetBackendServerPort(v int32) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody {
	s.BackendServerPort = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) SetBandwidth(v int32) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody {
	s.Bandwidth = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) SetDescription(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) SetForwardPort(v int32) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody {
	s.ForwardPort = &v
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

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) SetListenerForward(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody {
	s.ListenerForward = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) SetListenerPort(v int32) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody {
	s.ListenerPort = &v
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

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) SetScheduler(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody {
	s.Scheduler = &v
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

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) SetUnhealthyThreshold(v int32) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody {
	s.UnhealthyThreshold = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
