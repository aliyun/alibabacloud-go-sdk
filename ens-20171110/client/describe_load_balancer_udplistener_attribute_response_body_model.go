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
	// The port used by the backend ELB server of the ELB instance. Valid values: **1*	- to **65535**.
	//
	// example:
	//
	// 8080
	BackendServerPort *int32 `json:"BackendServerPort,omitempty" xml:"BackendServerPort,omitempty"`
	// The peak bandwidth of the Edge Load Balancer (ELB) instance. The default value is -1, which indicates that the bandwidth is not limited.
	//
	// example:
	//
	// 5
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The name of the listener.
	//
	// example:
	//
	// example
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Indicates whether EIP pass-through is enabled. Valid values:
	//
	// 	- **on**
	//
	// 	- **off*	- (default)
	//
	// example:
	//
	// on
	EipTransmit *string `json:"EipTransmit,omitempty" xml:"EipTransmit,omitempty"`
	// The timeout period of a connection. Valid values: **10*	- to **900**. Unit: seconds.
	//
	// example:
	//
	// 500
	EstablishedTimeout *int32 `json:"EstablishedTimeout,omitempty" xml:"EstablishedTimeout,omitempty"`
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
	// example:
	//
	// 8000
	HealthCheckConnectPort *int32 `json:"HealthCheckConnectPort,omitempty" xml:"HealthCheckConnectPort,omitempty"`
	// The timeout period for a health check response. If a backend server does not respond within the specified timeout period, the server fails the health check.
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
	// 	- If the value of the HealthCheckTimeout property is smaller than the value of the HealthCheckInterval property, the timeout period specified by the HealthCheckTimeout property becomes invalid and the value of the HealthCheckInterval property is used as the timeout period.
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
	// The interval at which health checks are performed. Valid values: **1*	- to **50**. Unit: seconds.
	//
	// >  This parameter takes effect only if you set HealthCheck to on.
	//
	// example:
	//
	// 5
	HealthCheckInterval *int32 `json:"HealthCheckInterval,omitempty" xml:"HealthCheckInterval,omitempty"`
	// The request string for UDP listener health checks. The string can be up to 64 characters in length and can contain only letters and digits.
	//
	// example:
	//
	// hello
	HealthCheckReq *string `json:"HealthCheckReq,omitempty" xml:"HealthCheckReq,omitempty"`
	// The number of consecutive successful health checks that must occur before an unhealthy and inaccessible backend server is declared healthy and accessible. Valid values: **2*	- to **10**.
	//
	// >  This parameter takes effect only if you set HealthCheck to on.
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
	// The ID of the request.
	//
	// example:
	//
	// 5D7597CF-1630-54EC-A945-624A33F2E7E8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The routing algorithm. Valid values:
	//
	// 	- **wrr**: Backend servers with higher weights receive more requests than those with lower weights.
	//
	// 	- **wlc**: Requests are distributed based on the weight and load of each backend server. The load refers to the number of connections to a backend server. If two backend servers have the same weight, the backend server that has fewer connections receives more requests.
	//
	// 	- **rr**: Requests are distributed to backend servers in sequence.
	//
	// 	- **sch**: consistent hashing that is based on source IP addresses. Requests from the same source IP address are distributed to the same backend server.
	//
	// 	- **qch**: consistent hashing that is based on QUIC connection IDs. Requests that contain the same QUIC connection ID are distributed to the same backend server.
	//
	// 	- **iqch**: consistent hashing that is based on specific three bytes of the iQUIC CIDs. Requests whose second to fourth bytes are the same are distributed to the same backend server.
	//
	// example:
	//
	// wrr
	Scheduler *string `json:"Scheduler,omitempty" xml:"Scheduler,omitempty"`
	// The status of the listener. Valid values:
	//
	// 	- **Running**
	//
	// 	- **Stopped**
	//
	// 	- **Starting**
	//
	// 	- **Configuring**
	//
	// 	- **Stopping**
	//
	// example:
	//
	// Stopped
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The number of consecutive failed health checks that must occur before a healthy and accessible backend server is declared unhealthy and inaccessible. Valid values: **2*	- to **10**.
	//
	// >  This parameter takes effect only if you set HealthCheck to on.
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
