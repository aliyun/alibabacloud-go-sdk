// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLoadBalancerUDPListenerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackendServerPort(v int32) *CreateLoadBalancerUDPListenerRequest
	GetBackendServerPort() *int32
	SetDescription(v string) *CreateLoadBalancerUDPListenerRequest
	GetDescription() *string
	SetEipTransmit(v string) *CreateLoadBalancerUDPListenerRequest
	GetEipTransmit() *string
	SetEstablishedTimeout(v int32) *CreateLoadBalancerUDPListenerRequest
	GetEstablishedTimeout() *int32
	SetHealthCheckConnectPort(v int32) *CreateLoadBalancerUDPListenerRequest
	GetHealthCheckConnectPort() *int32
	SetHealthCheckConnectTimeout(v int32) *CreateLoadBalancerUDPListenerRequest
	GetHealthCheckConnectTimeout() *int32
	SetHealthCheckExp(v string) *CreateLoadBalancerUDPListenerRequest
	GetHealthCheckExp() *string
	SetHealthCheckInterval(v int32) *CreateLoadBalancerUDPListenerRequest
	GetHealthCheckInterval() *int32
	SetHealthCheckReq(v string) *CreateLoadBalancerUDPListenerRequest
	GetHealthCheckReq() *string
	SetHealthyThreshold(v int32) *CreateLoadBalancerUDPListenerRequest
	GetHealthyThreshold() *int32
	SetListenerPort(v int32) *CreateLoadBalancerUDPListenerRequest
	GetListenerPort() *int32
	SetLoadBalancerId(v string) *CreateLoadBalancerUDPListenerRequest
	GetLoadBalancerId() *string
	SetScheduler(v string) *CreateLoadBalancerUDPListenerRequest
	GetScheduler() *string
	SetUnhealthyThreshold(v int32) *CreateLoadBalancerUDPListenerRequest
	GetUnhealthyThreshold() *int32
}

type CreateLoadBalancerUDPListenerRequest struct {
	// The port used by the backend servers. Valid values: **1*	- to **65535**.
	//
	// example:
	//
	// 8080
	BackendServerPort *int32 `json:"BackendServerPort,omitempty" xml:"BackendServerPort,omitempty"`
	// The name of the listener. The name must be **1*	- to **80*	- characters long.
	//
	// > The name cannot start with `http://` or `https://`.
	//
	// example:
	//
	// example
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether EIP pass-through is enabled. Valid values:
	//
	// - **on**: Enables EIP pass-through.
	//
	// - **off*	- (default): Disables EIP pass-through.
	//
	// example:
	//
	// on
	EipTransmit *string `json:"EipTransmit,omitempty" xml:"EipTransmit,omitempty"`
	// The timeout for established connections. Valid values: **10*	- to **900**. Default value: **900**. Unit: seconds.
	//
	// example:
	//
	// 500
	EstablishedTimeout *int32 `json:"EstablishedTimeout,omitempty" xml:"EstablishedTimeout,omitempty"`
	// The port used for health checks. Valid values: **1*	- to **65535**. If you do not specify this parameter, the port specified by `BackendServerPort` is used.
	//
	// example:
	//
	// 8080
	HealthCheckConnectPort *int32 `json:"HealthCheckConnectPort,omitempty" xml:"HealthCheckConnectPort,omitempty"`
	// The maximum time to wait for a response to a health check. A health check fails if a backend ENS instance does not respond within this time.
	//
	// - Default value: 5.
	//
	// - Valid values: **1*	- to **300**.
	//
	// - Unit: seconds.
	//
	// > If this timeout is shorter than the health check interval (`HealthCheckInterval`), the interval value is used for the timeout instead.
	//
	// example:
	//
	// 100
	HealthCheckConnectTimeout *int32 `json:"HealthCheckConnectTimeout,omitempty" xml:"HealthCheckConnectTimeout,omitempty"`
	// The expected response string for UDP health checks. The string must be up to 64 characters long and contain only letters and digits.
	//
	// example:
	//
	// ok
	HealthCheckExp *string `json:"HealthCheckExp,omitempty" xml:"HealthCheckExp,omitempty"`
	// The interval between health checks, in seconds. Valid values: **1*	- to **50**. Default value: **2**.
	//
	// example:
	//
	// 3
	HealthCheckInterval *int32 `json:"HealthCheckInterval,omitempty" xml:"HealthCheckInterval,omitempty"`
	// The request string for UDP health checks. The string must be up to 64 characters long and contain only letters and digits.
	//
	// example:
	//
	// hello
	HealthCheckReq *string `json:"HealthCheckReq,omitempty" xml:"HealthCheckReq,omitempty"`
	// The number of consecutive successful health checks required for a backend server to be considered healthy. Valid values: **2*	- to **10**. Default value: **3**.
	//
	// example:
	//
	// 4
	HealthyThreshold *int32 `json:"HealthyThreshold,omitempty" xml:"HealthyThreshold,omitempty"`
	// The listener port used to receive requests and forward them to backend servers. Valid values: **1*	- to **65535**.
	//
	// > For UDP listeners, ports 250, 4789, and 4790 are system-reserved and not available for use.
	//
	// This parameter is required.
	//
	// example:
	//
	// 80
	ListenerPort *int32 `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	// The ID of the load balancer instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// lb-5q73cv04zeyh43lh74lp4****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// The scheduling algorithm. Valid values:
	//
	// - **wrr*	- (default): Weighted Round Robin. Backend servers with a higher weight receive more requests.
	//
	// - **wlc**: Weighted Least Connections. Requests are distributed based on both the weight and the number of current connections of each backend server. If multiple backend servers have the same weight, the one with the fewest active connections receives the next request.
	//
	// - **rr**: Round Robin. Requests are distributed to backend servers sequentially.
	//
	// - **sch**: Source IP Hash. Requests from the same source IP address are routed to the same backend server.
	//
	// - **tch**: Four-tuple Hash. Requests are routed to backend servers based on a hash of four values: source IP address, destination IP address, source port, and destination port. Flows with the same four values are routed to the same backend server.
	//
	// - **qch**: QUIC Connection ID Hash. Requests with the same QUIC Connection ID are routed to the same backend server.
	//
	// - **iqch**: A consistent hash algorithm for iQUIC Connection IDs. Flows with the same second, third, and fourth bytes in the iQUIC CID are routed to the same backend server.
	//
	// example:
	//
	// wrr
	Scheduler *string `json:"Scheduler,omitempty" xml:"Scheduler,omitempty"`
	// The number of consecutive failed health checks required for a backend server to be considered unhealthy. Valid values: **2*	- to **10**. Default value: **3**.
	//
	// example:
	//
	// 4
	UnhealthyThreshold *int32 `json:"UnhealthyThreshold,omitempty" xml:"UnhealthyThreshold,omitempty"`
}

func (s CreateLoadBalancerUDPListenerRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateLoadBalancerUDPListenerRequest) GoString() string {
	return s.String()
}

func (s *CreateLoadBalancerUDPListenerRequest) GetBackendServerPort() *int32 {
	return s.BackendServerPort
}

func (s *CreateLoadBalancerUDPListenerRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateLoadBalancerUDPListenerRequest) GetEipTransmit() *string {
	return s.EipTransmit
}

func (s *CreateLoadBalancerUDPListenerRequest) GetEstablishedTimeout() *int32 {
	return s.EstablishedTimeout
}

func (s *CreateLoadBalancerUDPListenerRequest) GetHealthCheckConnectPort() *int32 {
	return s.HealthCheckConnectPort
}

func (s *CreateLoadBalancerUDPListenerRequest) GetHealthCheckConnectTimeout() *int32 {
	return s.HealthCheckConnectTimeout
}

func (s *CreateLoadBalancerUDPListenerRequest) GetHealthCheckExp() *string {
	return s.HealthCheckExp
}

func (s *CreateLoadBalancerUDPListenerRequest) GetHealthCheckInterval() *int32 {
	return s.HealthCheckInterval
}

func (s *CreateLoadBalancerUDPListenerRequest) GetHealthCheckReq() *string {
	return s.HealthCheckReq
}

func (s *CreateLoadBalancerUDPListenerRequest) GetHealthyThreshold() *int32 {
	return s.HealthyThreshold
}

func (s *CreateLoadBalancerUDPListenerRequest) GetListenerPort() *int32 {
	return s.ListenerPort
}

func (s *CreateLoadBalancerUDPListenerRequest) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *CreateLoadBalancerUDPListenerRequest) GetScheduler() *string {
	return s.Scheduler
}

func (s *CreateLoadBalancerUDPListenerRequest) GetUnhealthyThreshold() *int32 {
	return s.UnhealthyThreshold
}

func (s *CreateLoadBalancerUDPListenerRequest) SetBackendServerPort(v int32) *CreateLoadBalancerUDPListenerRequest {
	s.BackendServerPort = &v
	return s
}

func (s *CreateLoadBalancerUDPListenerRequest) SetDescription(v string) *CreateLoadBalancerUDPListenerRequest {
	s.Description = &v
	return s
}

func (s *CreateLoadBalancerUDPListenerRequest) SetEipTransmit(v string) *CreateLoadBalancerUDPListenerRequest {
	s.EipTransmit = &v
	return s
}

func (s *CreateLoadBalancerUDPListenerRequest) SetEstablishedTimeout(v int32) *CreateLoadBalancerUDPListenerRequest {
	s.EstablishedTimeout = &v
	return s
}

func (s *CreateLoadBalancerUDPListenerRequest) SetHealthCheckConnectPort(v int32) *CreateLoadBalancerUDPListenerRequest {
	s.HealthCheckConnectPort = &v
	return s
}

func (s *CreateLoadBalancerUDPListenerRequest) SetHealthCheckConnectTimeout(v int32) *CreateLoadBalancerUDPListenerRequest {
	s.HealthCheckConnectTimeout = &v
	return s
}

func (s *CreateLoadBalancerUDPListenerRequest) SetHealthCheckExp(v string) *CreateLoadBalancerUDPListenerRequest {
	s.HealthCheckExp = &v
	return s
}

func (s *CreateLoadBalancerUDPListenerRequest) SetHealthCheckInterval(v int32) *CreateLoadBalancerUDPListenerRequest {
	s.HealthCheckInterval = &v
	return s
}

func (s *CreateLoadBalancerUDPListenerRequest) SetHealthCheckReq(v string) *CreateLoadBalancerUDPListenerRequest {
	s.HealthCheckReq = &v
	return s
}

func (s *CreateLoadBalancerUDPListenerRequest) SetHealthyThreshold(v int32) *CreateLoadBalancerUDPListenerRequest {
	s.HealthyThreshold = &v
	return s
}

func (s *CreateLoadBalancerUDPListenerRequest) SetListenerPort(v int32) *CreateLoadBalancerUDPListenerRequest {
	s.ListenerPort = &v
	return s
}

func (s *CreateLoadBalancerUDPListenerRequest) SetLoadBalancerId(v string) *CreateLoadBalancerUDPListenerRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *CreateLoadBalancerUDPListenerRequest) SetScheduler(v string) *CreateLoadBalancerUDPListenerRequest {
	s.Scheduler = &v
	return s
}

func (s *CreateLoadBalancerUDPListenerRequest) SetUnhealthyThreshold(v int32) *CreateLoadBalancerUDPListenerRequest {
	s.UnhealthyThreshold = &v
	return s
}

func (s *CreateLoadBalancerUDPListenerRequest) Validate() error {
	return dara.Validate(s)
}
