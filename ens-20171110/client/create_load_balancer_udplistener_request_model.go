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
	// The port used by the backend ELB server of the ELB instance. Valid values: **1*	- to **65535**.
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
	// example
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to enable Elastic IP address (EIP) pass-through. Valid values:
	//
	// 	- **on**
	//
	// 	- **off*	- (default)
	//
	// example:
	//
	// on
	EipTransmit *string `json:"EipTransmit,omitempty" xml:"EipTransmit,omitempty"`
	// The timeout period of a connection. Valid values: **10*	- to **900**. Default value: **900**. Unit: seconds.
	//
	// example:
	//
	// 500
	EstablishedTimeout *int32 `json:"EstablishedTimeout,omitempty" xml:"EstablishedTimeout,omitempty"`
	// The port that is used for health checks. Valid values: **1*	- to **65535**. If you leave this parameter empty, the port specified by BackendServerPort is used for health checks.
	//
	// example:
	//
	// 8080
	HealthCheckConnectPort *int32 `json:"HealthCheckConnectPort,omitempty" xml:"HealthCheckConnectPort,omitempty"`
	// The timeout period for a health check response. If a backend server does not respond within the specified timeout period, the server fails the health check.
	//
	// 	- Default value: 5.
	//
	// 	- Valid values: **1*	- to **300**.
	//
	// 	- Unit: seconds.
	//
	// >  If the value of the HealthCheckConnectTimeout parameter is smaller than that of the HealthCheckInterval parameter, the timeout period specified by the HealthCheckConnectTimeout parameter becomes invalid and the value of the HealthCheckInterval parameter is used as the timeout period.
	//
	// example:
	//
	// 100
	HealthCheckConnectTimeout *int32 `json:"HealthCheckConnectTimeout,omitempty" xml:"HealthCheckConnectTimeout,omitempty"`
	// The response string for UDP listener health checks. The string can be up to 64 characters in length and can contain only letters and digits.
	//
	// example:
	//
	// ok
	HealthCheckExp *string `json:"HealthCheckExp,omitempty" xml:"HealthCheckExp,omitempty"`
	// The interval at which health checks are performed. Valid values: **1*	- to **50**. Default value: **2**. Unit: seconds.
	//
	// example:
	//
	// 3
	HealthCheckInterval *int32 `json:"HealthCheckInterval,omitempty" xml:"HealthCheckInterval,omitempty"`
	// The request string for UDP listener health checks. The string can be up to 64 characters in length and can contain only letters and digits.
	//
	// example:
	//
	// hello
	HealthCheckReq *string `json:"HealthCheckReq,omitempty" xml:"HealthCheckReq,omitempty"`
	// The number of consecutive successful health checks that must occur before an unhealthy and inaccessible backend server is declared healthy and accessible. Valid values: **2*	- to **10**. Default value: **3**.
	//
	// example:
	//
	// 4
	HealthyThreshold *int32 `json:"HealthyThreshold,omitempty" xml:"HealthyThreshold,omitempty"`
	// The listener port that is used by Edge Load Balancer (ELB) to receive requests and forward the requests to backend servers. Valid values: **1*	- to **65535**.
	//
	// >  You cannot specify ports 250, 4789, or 4790 for UDP listeners. They are system reserved ports.
	//
	// This parameter is required.
	//
	// example:
	//
	// 80
	ListenerPort *int32 `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	// The ID of the Edge Load Balancer (ELB) instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// lb-5q73cv04zeyh43lh74lp4****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// The scheduling algorithm. Valid values:
	//
	// 	- **wrr**: Backend servers with higher weights receive more requests than backend servers with lower weights. This is the default value.
	//
	// 	- **wlc**: Requests are distributed based on the weight and load of each backend server. The load refers to the number of connections on a backend server. If two backend servers have the same weight, the backend server that has fewer connections receives more requests.
	//
	// 	- **rr**: Requests are distributed to backend servers in sequence.
	//
	// 	- **sch**: Consistent hashing that is based on source IP addresses. Requests from the same source IP address are distributed to the same backend server.
	//
	// 	- **qch**: Consistent hashing based on Quick UDP Internet Connection (QUIC) IDs. Requests that contain the same QUIC ID are scheduled to the same backend server.
	//
	// 	- **iqch**: Consistent hashing based on three specific bytes of iQUIC CID. Requests with the same second, third, and forth bytes are scheduled to the same backend server.
	//
	// example:
	//
	// wrr
	Scheduler *string `json:"Scheduler,omitempty" xml:"Scheduler,omitempty"`
	// The number of consecutive failed health checks that must occur before a healthy and accessible backend server is declared unhealthy and inaccessible. Valid values: **2*	- to **10**. Default value: **3**.
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
