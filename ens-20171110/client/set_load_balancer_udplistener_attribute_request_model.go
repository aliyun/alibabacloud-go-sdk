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
	// The name of the listener. The name must be **1*	- to **80*	- characters in length.
	//
	// > The name cannot start with `http://` or `https://`.
	//
	// example:
	//
	// example
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to enable EIP pass-through. Valid values:
	//
	// - **on**: Enables EIP pass-through.
	//
	// - **off**: Disables EIP pass-through. This is the default value.
	//
	// example:
	//
	// on
	EipTransmit *string `json:"EipTransmit,omitempty" xml:"EipTransmit,omitempty"`
	// The connection timeout period. Valid values: **10*	- to **900**. Default value: **900**. Unit: seconds.
	//
	// example:
	//
	// 500
	EstablishedTimeout *int32 `json:"EstablishedTimeout,omitempty" xml:"EstablishedTimeout,omitempty"`
	// The port used for health checks. Valid values: **1*	- to **65535**. If no Settings are specified for this parameter, the backend service port (BackendServerPort) is used.
	//
	// example:
	//
	// 8080
	HealthCheckConnectPort *int32 `json:"HealthCheckConnectPort,omitempty" xml:"HealthCheckConnectPort,omitempty"`
	// The timeout period for a health check response. If a backend ENS instance does not respond within the specified timeout period, the health check fails.
	//
	// - Default value: 5.
	//
	// - Valid values: **1*	- to **300**.
	//
	// - Unit: seconds.
	//
	// > If the value of HealthCheckConnectTimeout is less than the value of HealthCheckInterval, HealthCheckConnectTimeout is invalid and the timeout period is the value of HealthCheckInterval.
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
	// The interval between two consecutive health checks. Valid values: **1*	- to **50**. Unit: seconds.
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
	// The number of consecutive successful health checks that must occur before a backend server is declared healthy (from fail to success). Valid values: **2*	- to **10**.
	//
	// example:
	//
	// 4
	HealthyThreshold *int32 `json:"HealthyThreshold,omitempty" xml:"HealthyThreshold,omitempty"`
	// The frontend port of the listener whose attributes you want to modify. Valid values: **1*	- to **65535**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 80
	ListenerPort *int32 `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	// The ID of the load balancing instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// lb-5pzipr2fszqtl2xf64uy5****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
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
	// - **tch**: Consistent hashing based on quadruples (source IP address, destination IP address, source port, and destination port). Requests from the same stream are distributed to the same backend server.
	//
	// - **qch**: Consistent hashing based on QUIC connection IDs. Requests with the same QUIC connection ID are distributed to the same backend server.
	//
	// - **iqch**: Consistent hashing based on specific three bytes of the iQUIC CID. Requests with the same second to fourth bytes are distributed to the same backend server.
	//
	// example:
	//
	// wrr
	Scheduler *string `json:"Scheduler,omitempty" xml:"Scheduler,omitempty"`
	// The number of consecutive failed health checks that must occur before a backend server is declared unhealthy (from success to fail). Valid values: **2*	- to **10**.
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
