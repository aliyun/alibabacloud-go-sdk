// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetLoadBalancerUDPListenerAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *SetLoadBalancerUDPListenerAttributeRequest
	GetDescription() *string
	SetEipTransmit(v string) *SetLoadBalancerUDPListenerAttributeRequest
	GetEipTransmit() *string
	SetEstablishedTimeout(v int32) *SetLoadBalancerUDPListenerAttributeRequest
	GetEstablishedTimeout() *int32
	SetHealthCheckConnectPort(v int32) *SetLoadBalancerUDPListenerAttributeRequest
	GetHealthCheckConnectPort() *int32
	SetHealthCheckConnectTimeout(v int32) *SetLoadBalancerUDPListenerAttributeRequest
	GetHealthCheckConnectTimeout() *int32
	SetHealthCheckExp(v string) *SetLoadBalancerUDPListenerAttributeRequest
	GetHealthCheckExp() *string
	SetHealthCheckInterval(v int32) *SetLoadBalancerUDPListenerAttributeRequest
	GetHealthCheckInterval() *int32
	SetHealthCheckReq(v string) *SetLoadBalancerUDPListenerAttributeRequest
	GetHealthCheckReq() *string
	SetHealthyThreshold(v int32) *SetLoadBalancerUDPListenerAttributeRequest
	GetHealthyThreshold() *int32
	SetListenerPort(v int32) *SetLoadBalancerUDPListenerAttributeRequest
	GetListenerPort() *int32
	SetLoadBalancerId(v string) *SetLoadBalancerUDPListenerAttributeRequest
	GetLoadBalancerId() *string
	SetScheduler(v string) *SetLoadBalancerUDPListenerAttributeRequest
	GetScheduler() *string
	SetUnhealthyThreshold(v int32) *SetLoadBalancerUDPListenerAttributeRequest
	GetUnhealthyThreshold() *int32
}

type SetLoadBalancerUDPListenerAttributeRequest struct {
	// The name of the listener. The name must be **1*	- to **80*	- characters long.
	//
	// > The name cannot start with `http://` or `https://`.
	//
	// example:
	//
	// example
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to enable EIP passthrough. Valid values:
	//
	// - **on**: Enables EIP passthrough.
	//
	// - **off*	- (default): Disables EIP passthrough.
	//
	// example:
	//
	// on
	EipTransmit *string `json:"EipTransmit,omitempty" xml:"EipTransmit,omitempty"`
	// The timeout for established connections, in seconds. The value must be an integer from **10*	- to **900**. The default value is **900**.
	//
	// example:
	//
	// 500
	EstablishedTimeout *int32 `json:"EstablishedTimeout,omitempty" xml:"EstablishedTimeout,omitempty"`
	// The port used for health checks. Valid values: **1*	- to **65535**. Defaults to the backend server port (`BackendServerPort`) if not specified.
	//
	// example:
	//
	// 8080
	HealthCheckConnectPort *int32 `json:"HealthCheckConnectPort,omitempty" xml:"HealthCheckConnectPort,omitempty"`
	// The timeout for a health check response. If a backend ENS instance does not respond within this timeout period, the health check fails.
	//
	// - Default value: 5
	//
	// - Valid values: **1*	- to **300**
	//
	// - Unit: seconds
	//
	// > If the value of `HealthCheckConnectTimeout` is less than the value of `HealthCheckInterval`, the system ignores `HealthCheckConnectTimeout` and uses the value of `HealthCheckInterval` as the timeout period.
	//
	// example:
	//
	// 100
	HealthCheckConnectTimeout *int32 `json:"HealthCheckConnectTimeout,omitempty" xml:"HealthCheckConnectTimeout,omitempty"`
	// The expected response string for a UDP listener health check. The string can contain only letters and digits and must be at most 64 characters long.
	//
	// example:
	//
	// ok
	HealthCheckExp *string `json:"HealthCheckExp,omitempty" xml:"HealthCheckExp,omitempty"`
	// The interval, in seconds, between consecutive health checks. Valid values: **1*	- to **50**.
	//
	// example:
	//
	// 5
	HealthCheckInterval *int32 `json:"HealthCheckInterval,omitempty" xml:"HealthCheckInterval,omitempty"`
	// The request string for a UDP listener health check. The string can contain only letters and digits and must be at most 64 characters long.
	//
	// example:
	//
	// hello
	HealthCheckReq *string `json:"HealthCheckReq,omitempty" xml:"HealthCheckReq,omitempty"`
	// The number of consecutive successful health checks required to mark an unhealthy backend server as healthy. Valid values: **2*	- to **10**.
	//
	// example:
	//
	// 4
	HealthyThreshold *int32 `json:"HealthyThreshold,omitempty" xml:"HealthyThreshold,omitempty"`
	// The listener port to modify. Valid values: **1*	- to **65535**.
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
	// lb-5pzipr2fszqtl2xf64uy5****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// The scheduling algorithm. Valid values:
	//
	// - **wrr**: Weighted Round Robin. Backend servers with higher weights receive a larger proportion of requests.
	//
	// - **wlc**: Weighted Least Connections. Requests are distributed based on both the weights and the number of active connections of backend servers. If servers have the same weight, the server with the fewest active connections receives more requests.
	//
	// - **rr**: Round Robin. Requests are distributed to backend servers in sequence.
	//
	// - **sch**: Source IP Hash. A consistent hash based on the source IP address. Requests from the same source IP address are routed to the same backend server.
	//
	// - **tch**: 4-Tuple Hash. A consistent hash based on the 4-tuple (source IP, destination IP, source port, and destination port). Flows with the same 4-tuple are routed to the same backend server.
	//
	// - **qch**: QUIC Connection ID Hash. A consistent hash based on the QUIC Connection ID. Requests with the same QUIC Connection ID are routed to the same backend server.
	//
	// - **iqch**: iQUIC CID Hash. A consistent hash based on three specific bytes (from the second to the fourth byte) of the iQUIC CID. Requests with identical values in these three bytes are routed to the same backend server.
	//
	// example:
	//
	// wrr
	Scheduler *string `json:"Scheduler,omitempty" xml:"Scheduler,omitempty"`
	// The number of consecutive failed health checks required to mark a healthy backend server as unhealthy. Valid values: **2*	- to **10**.
	//
	// example:
	//
	// 4
	UnhealthyThreshold *int32 `json:"UnhealthyThreshold,omitempty" xml:"UnhealthyThreshold,omitempty"`
}

func (s SetLoadBalancerUDPListenerAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s SetLoadBalancerUDPListenerAttributeRequest) GoString() string {
	return s.String()
}

func (s *SetLoadBalancerUDPListenerAttributeRequest) GetDescription() *string {
	return s.Description
}

func (s *SetLoadBalancerUDPListenerAttributeRequest) GetEipTransmit() *string {
	return s.EipTransmit
}

func (s *SetLoadBalancerUDPListenerAttributeRequest) GetEstablishedTimeout() *int32 {
	return s.EstablishedTimeout
}

func (s *SetLoadBalancerUDPListenerAttributeRequest) GetHealthCheckConnectPort() *int32 {
	return s.HealthCheckConnectPort
}

func (s *SetLoadBalancerUDPListenerAttributeRequest) GetHealthCheckConnectTimeout() *int32 {
	return s.HealthCheckConnectTimeout
}

func (s *SetLoadBalancerUDPListenerAttributeRequest) GetHealthCheckExp() *string {
	return s.HealthCheckExp
}

func (s *SetLoadBalancerUDPListenerAttributeRequest) GetHealthCheckInterval() *int32 {
	return s.HealthCheckInterval
}

func (s *SetLoadBalancerUDPListenerAttributeRequest) GetHealthCheckReq() *string {
	return s.HealthCheckReq
}

func (s *SetLoadBalancerUDPListenerAttributeRequest) GetHealthyThreshold() *int32 {
	return s.HealthyThreshold
}

func (s *SetLoadBalancerUDPListenerAttributeRequest) GetListenerPort() *int32 {
	return s.ListenerPort
}

func (s *SetLoadBalancerUDPListenerAttributeRequest) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *SetLoadBalancerUDPListenerAttributeRequest) GetScheduler() *string {
	return s.Scheduler
}

func (s *SetLoadBalancerUDPListenerAttributeRequest) GetUnhealthyThreshold() *int32 {
	return s.UnhealthyThreshold
}

func (s *SetLoadBalancerUDPListenerAttributeRequest) SetDescription(v string) *SetLoadBalancerUDPListenerAttributeRequest {
	s.Description = &v
	return s
}

func (s *SetLoadBalancerUDPListenerAttributeRequest) SetEipTransmit(v string) *SetLoadBalancerUDPListenerAttributeRequest {
	s.EipTransmit = &v
	return s
}

func (s *SetLoadBalancerUDPListenerAttributeRequest) SetEstablishedTimeout(v int32) *SetLoadBalancerUDPListenerAttributeRequest {
	s.EstablishedTimeout = &v
	return s
}

func (s *SetLoadBalancerUDPListenerAttributeRequest) SetHealthCheckConnectPort(v int32) *SetLoadBalancerUDPListenerAttributeRequest {
	s.HealthCheckConnectPort = &v
	return s
}

func (s *SetLoadBalancerUDPListenerAttributeRequest) SetHealthCheckConnectTimeout(v int32) *SetLoadBalancerUDPListenerAttributeRequest {
	s.HealthCheckConnectTimeout = &v
	return s
}

func (s *SetLoadBalancerUDPListenerAttributeRequest) SetHealthCheckExp(v string) *SetLoadBalancerUDPListenerAttributeRequest {
	s.HealthCheckExp = &v
	return s
}

func (s *SetLoadBalancerUDPListenerAttributeRequest) SetHealthCheckInterval(v int32) *SetLoadBalancerUDPListenerAttributeRequest {
	s.HealthCheckInterval = &v
	return s
}

func (s *SetLoadBalancerUDPListenerAttributeRequest) SetHealthCheckReq(v string) *SetLoadBalancerUDPListenerAttributeRequest {
	s.HealthCheckReq = &v
	return s
}

func (s *SetLoadBalancerUDPListenerAttributeRequest) SetHealthyThreshold(v int32) *SetLoadBalancerUDPListenerAttributeRequest {
	s.HealthyThreshold = &v
	return s
}

func (s *SetLoadBalancerUDPListenerAttributeRequest) SetListenerPort(v int32) *SetLoadBalancerUDPListenerAttributeRequest {
	s.ListenerPort = &v
	return s
}

func (s *SetLoadBalancerUDPListenerAttributeRequest) SetLoadBalancerId(v string) *SetLoadBalancerUDPListenerAttributeRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *SetLoadBalancerUDPListenerAttributeRequest) SetScheduler(v string) *SetLoadBalancerUDPListenerAttributeRequest {
	s.Scheduler = &v
	return s
}

func (s *SetLoadBalancerUDPListenerAttributeRequest) SetUnhealthyThreshold(v int32) *SetLoadBalancerUDPListenerAttributeRequest {
	s.UnhealthyThreshold = &v
	return s
}

func (s *SetLoadBalancerUDPListenerAttributeRequest) Validate() error {
	return dara.Validate(s)
}
