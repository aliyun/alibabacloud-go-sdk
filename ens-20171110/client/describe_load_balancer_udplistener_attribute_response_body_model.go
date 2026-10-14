// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLoadBalancerUDPListenerAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBackendServerPort(v int32) *DescribeLoadBalancerUDPListenerAttributeResponseBody
	GetBackendServerPort() *int32
	SetBandwidth(v int32) *DescribeLoadBalancerUDPListenerAttributeResponseBody
	GetBandwidth() *int32
	SetDescription(v string) *DescribeLoadBalancerUDPListenerAttributeResponseBody
	GetDescription() *string
	SetEipTransmit(v string) *DescribeLoadBalancerUDPListenerAttributeResponseBody
	GetEipTransmit() *string
	SetEstablishedTimeout(v int32) *DescribeLoadBalancerUDPListenerAttributeResponseBody
	GetEstablishedTimeout() *int32
	SetHealthCheck(v string) *DescribeLoadBalancerUDPListenerAttributeResponseBody
	GetHealthCheck() *string
	SetHealthCheckConnectPort(v int32) *DescribeLoadBalancerUDPListenerAttributeResponseBody
	GetHealthCheckConnectPort() *int32
	SetHealthCheckConnectTimeout(v int32) *DescribeLoadBalancerUDPListenerAttributeResponseBody
	GetHealthCheckConnectTimeout() *int32
	SetHealthCheckExp(v string) *DescribeLoadBalancerUDPListenerAttributeResponseBody
	GetHealthCheckExp() *string
	SetHealthCheckInterval(v int32) *DescribeLoadBalancerUDPListenerAttributeResponseBody
	GetHealthCheckInterval() *int32
	SetHealthCheckReq(v string) *DescribeLoadBalancerUDPListenerAttributeResponseBody
	GetHealthCheckReq() *string
	SetHealthyThreshold(v int32) *DescribeLoadBalancerUDPListenerAttributeResponseBody
	GetHealthyThreshold() *int32
	SetListenerPort(v int32) *DescribeLoadBalancerUDPListenerAttributeResponseBody
	GetListenerPort() *int32
	SetRequestId(v string) *DescribeLoadBalancerUDPListenerAttributeResponseBody
	GetRequestId() *string
	SetScheduler(v string) *DescribeLoadBalancerUDPListenerAttributeResponseBody
	GetScheduler() *string
	SetStatus(v string) *DescribeLoadBalancerUDPListenerAttributeResponseBody
	GetStatus() *string
	SetUnhealthyThreshold(v int32) *DescribeLoadBalancerUDPListenerAttributeResponseBody
	GetUnhealthyThreshold() *int32
}

type DescribeLoadBalancerUDPListenerAttributeResponseBody struct {
	// The port used by the backend servers. Valid values: **1*	- to **65535**.
	//
	// example:
	//
	// 8080
	BackendServerPort *int32 `json:"BackendServerPort,omitempty" xml:"BackendServerPort,omitempty"`
	// The peak bandwidth of the load balancer. A value of -1 indicates that no rate limit is applied.
	//
	// example:
	//
	// 5
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The listener name.
	//
	// example:
	//
	// example
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to enable EIP passthrough. Valid values:
	//
	// - **on**: Enabled.
	//
	// - **off*	- (default): Disabled.
	//
	// example:
	//
	// on
	EipTransmit *string `json:"EipTransmit,omitempty" xml:"EipTransmit,omitempty"`
	// The connection timeout period. Valid values: **10*	- to **900**. Unit: seconds.
	//
	// example:
	//
	// 500
	EstablishedTimeout *int32 `json:"EstablishedTimeout,omitempty" xml:"EstablishedTimeout,omitempty"`
	// Specifies whether to enable health checks. Valid values:
	//
	// - **on**: Enabled.
	//
	// - **off**: Disabled.
	//
	// example:
	//
	// on
	HealthCheck *string `json:"HealthCheck,omitempty" xml:"HealthCheck,omitempty"`
	// The port used for health checks. Valid values: **1*	- to **65535**. If you do not specify this parameter, the value of `BackendServerPort` is used.
	//
	// example:
	//
	// 8000
	HealthCheckConnectPort *int32 `json:"HealthCheckConnectPort,omitempty" xml:"HealthCheckConnectPort,omitempty"`
	// The timeout period for a health check response. If a backend ENS instance does not respond within the specified timeout period, the health check fails.
	//
	// - Default value: 5.
	//
	// - Valid values: **1*	- to **300**.
	//
	// - Unit: seconds.
	//
	// > 	- This parameter takes effect only when `HealthCheck` is set to `on`.
	//
	// >
	//
	// > 	- If the value of `HealthCheckConnectTimeout` is less than the value of `HealthCheckInterval`, the value of `HealthCheckInterval` is used as the timeout period.
	//
	// example:
	//
	// 100
	HealthCheckConnectTimeout *int32 `json:"HealthCheckConnectTimeout,omitempty" xml:"HealthCheckConnectTimeout,omitempty"`
	// The response string for UDP health checks. The string can contain only letters and digits and can be up to 64 characters in length.
	//
	// example:
	//
	// ok
	HealthCheckExp *string `json:"HealthCheckExp,omitempty" xml:"HealthCheckExp,omitempty"`
	// The interval between health checks. Valid values: **1*	- to **50**. Unit: seconds.
	//
	// > This parameter takes effect only when `HealthCheck` is set to `on`.
	//
	// example:
	//
	// 5
	HealthCheckInterval *int32 `json:"HealthCheckInterval,omitempty" xml:"HealthCheckInterval,omitempty"`
	// The request string for UDP health checks. The string can contain only letters and digits and can be up to 64 characters in length.
	//
	// example:
	//
	// hello
	HealthCheckReq *string `json:"HealthCheckReq,omitempty" xml:"HealthCheckReq,omitempty"`
	// The number of consecutive successful health checks that must occur before a backend server can be declared healthy. Valid values: **2*	- to **10**.
	//
	// > This parameter takes effect only when `HealthCheck` is set to `on`.
	//
	// example:
	//
	// 4
	HealthyThreshold *int32 `json:"HealthyThreshold,omitempty" xml:"HealthyThreshold,omitempty"`
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
	// 5D7597CF-1630-54EC-A945-624A33F2E7E8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The scheduling algorithm. Valid values:
	//
	// - **wrr*	- (Weighted Round Robin): Backend servers with higher weights receive a larger proportion of requests.
	//
	// - **wlc*	- (Weighted Least Connections): Requests are distributed based on both the weight and the current connection count of each backend server. If backend servers have the same weight, the server with the fewest connections receives more new requests.
	//
	// - **rr*	- (Round Robin): Requests are distributed to backend servers sequentially.
	//
	// - **sch*	- (Source IP Hash): Routes requests from the same source IP address to the same backend server by using a consistent hash of the source IP address.
	//
	// - **tch*	- (4-tuple Hash): Routes packets from the same flow to the same backend server by using a consistent hash of the 4-tuple (source IP address, destination IP address, source port, and destination port).
	//
	// - **qch*	- (QUIC Connection ID Hash): Uses a consistent hashing algorithm on the QUIC Connection ID to ensure that requests with the same connection ID are always routed to the same backend server.
	//
	// - **iqch*	- (iQUIC CID Hash): Uses a consistent hashing algorithm on bytes 2, 3, and 4 of the iQUIC CID. Requests where these three bytes are the same are routed to the same backend server.
	//
	// example:
	//
	// wrr
	Scheduler *string `json:"Scheduler,omitempty" xml:"Scheduler,omitempty"`
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
	// Stopped
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The number of consecutive failed health checks that must occur before a backend server can be declared unhealthy. Valid values: **2*	- to **10**.
	//
	// > This parameter takes effect only when `HealthCheck` is set to `on`.
	//
	// example:
	//
	// 4
	UnhealthyThreshold *int32 `json:"UnhealthyThreshold,omitempty" xml:"UnhealthyThreshold,omitempty"`
}

func (s DescribeLoadBalancerUDPListenerAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLoadBalancerUDPListenerAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancerUDPListenerAttributeResponseBody) GetBackendServerPort() *int32 {
	return s.BackendServerPort
}

func (s *DescribeLoadBalancerUDPListenerAttributeResponseBody) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *DescribeLoadBalancerUDPListenerAttributeResponseBody) GetDescription() *string {
	return s.Description
}

func (s *DescribeLoadBalancerUDPListenerAttributeResponseBody) GetEipTransmit() *string {
	return s.EipTransmit
}

func (s *DescribeLoadBalancerUDPListenerAttributeResponseBody) GetEstablishedTimeout() *int32 {
	return s.EstablishedTimeout
}

func (s *DescribeLoadBalancerUDPListenerAttributeResponseBody) GetHealthCheck() *string {
	return s.HealthCheck
}

func (s *DescribeLoadBalancerUDPListenerAttributeResponseBody) GetHealthCheckConnectPort() *int32 {
	return s.HealthCheckConnectPort
}

func (s *DescribeLoadBalancerUDPListenerAttributeResponseBody) GetHealthCheckConnectTimeout() *int32 {
	return s.HealthCheckConnectTimeout
}

func (s *DescribeLoadBalancerUDPListenerAttributeResponseBody) GetHealthCheckExp() *string {
	return s.HealthCheckExp
}

func (s *DescribeLoadBalancerUDPListenerAttributeResponseBody) GetHealthCheckInterval() *int32 {
	return s.HealthCheckInterval
}

func (s *DescribeLoadBalancerUDPListenerAttributeResponseBody) GetHealthCheckReq() *string {
	return s.HealthCheckReq
}

func (s *DescribeLoadBalancerUDPListenerAttributeResponseBody) GetHealthyThreshold() *int32 {
	return s.HealthyThreshold
}

func (s *DescribeLoadBalancerUDPListenerAttributeResponseBody) GetListenerPort() *int32 {
	return s.ListenerPort
}

func (s *DescribeLoadBalancerUDPListenerAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLoadBalancerUDPListenerAttributeResponseBody) GetScheduler() *string {
	return s.Scheduler
}

func (s *DescribeLoadBalancerUDPListenerAttributeResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeLoadBalancerUDPListenerAttributeResponseBody) GetUnhealthyThreshold() *int32 {
	return s.UnhealthyThreshold
}

func (s *DescribeLoadBalancerUDPListenerAttributeResponseBody) SetBackendServerPort(v int32) *DescribeLoadBalancerUDPListenerAttributeResponseBody {
	s.BackendServerPort = &v
	return s
}

func (s *DescribeLoadBalancerUDPListenerAttributeResponseBody) SetBandwidth(v int32) *DescribeLoadBalancerUDPListenerAttributeResponseBody {
	s.Bandwidth = &v
	return s
}

func (s *DescribeLoadBalancerUDPListenerAttributeResponseBody) SetDescription(v string) *DescribeLoadBalancerUDPListenerAttributeResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeLoadBalancerUDPListenerAttributeResponseBody) SetEipTransmit(v string) *DescribeLoadBalancerUDPListenerAttributeResponseBody {
	s.EipTransmit = &v
	return s
}

func (s *DescribeLoadBalancerUDPListenerAttributeResponseBody) SetEstablishedTimeout(v int32) *DescribeLoadBalancerUDPListenerAttributeResponseBody {
	s.EstablishedTimeout = &v
	return s
}

func (s *DescribeLoadBalancerUDPListenerAttributeResponseBody) SetHealthCheck(v string) *DescribeLoadBalancerUDPListenerAttributeResponseBody {
	s.HealthCheck = &v
	return s
}

func (s *DescribeLoadBalancerUDPListenerAttributeResponseBody) SetHealthCheckConnectPort(v int32) *DescribeLoadBalancerUDPListenerAttributeResponseBody {
	s.HealthCheckConnectPort = &v
	return s
}

func (s *DescribeLoadBalancerUDPListenerAttributeResponseBody) SetHealthCheckConnectTimeout(v int32) *DescribeLoadBalancerUDPListenerAttributeResponseBody {
	s.HealthCheckConnectTimeout = &v
	return s
}

func (s *DescribeLoadBalancerUDPListenerAttributeResponseBody) SetHealthCheckExp(v string) *DescribeLoadBalancerUDPListenerAttributeResponseBody {
	s.HealthCheckExp = &v
	return s
}

func (s *DescribeLoadBalancerUDPListenerAttributeResponseBody) SetHealthCheckInterval(v int32) *DescribeLoadBalancerUDPListenerAttributeResponseBody {
	s.HealthCheckInterval = &v
	return s
}

func (s *DescribeLoadBalancerUDPListenerAttributeResponseBody) SetHealthCheckReq(v string) *DescribeLoadBalancerUDPListenerAttributeResponseBody {
	s.HealthCheckReq = &v
	return s
}

func (s *DescribeLoadBalancerUDPListenerAttributeResponseBody) SetHealthyThreshold(v int32) *DescribeLoadBalancerUDPListenerAttributeResponseBody {
	s.HealthyThreshold = &v
	return s
}

func (s *DescribeLoadBalancerUDPListenerAttributeResponseBody) SetListenerPort(v int32) *DescribeLoadBalancerUDPListenerAttributeResponseBody {
	s.ListenerPort = &v
	return s
}

func (s *DescribeLoadBalancerUDPListenerAttributeResponseBody) SetRequestId(v string) *DescribeLoadBalancerUDPListenerAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLoadBalancerUDPListenerAttributeResponseBody) SetScheduler(v string) *DescribeLoadBalancerUDPListenerAttributeResponseBody {
	s.Scheduler = &v
	return s
}

func (s *DescribeLoadBalancerUDPListenerAttributeResponseBody) SetStatus(v string) *DescribeLoadBalancerUDPListenerAttributeResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeLoadBalancerUDPListenerAttributeResponseBody) SetUnhealthyThreshold(v int32) *DescribeLoadBalancerUDPListenerAttributeResponseBody {
	s.UnhealthyThreshold = &v
	return s
}

func (s *DescribeLoadBalancerUDPListenerAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
