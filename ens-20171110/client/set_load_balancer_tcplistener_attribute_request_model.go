// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetLoadBalancerTCPListenerAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *SetLoadBalancerTCPListenerAttributeRequest
	GetDescription() *string
	SetEipTransmit(v string) *SetLoadBalancerTCPListenerAttributeRequest
	GetEipTransmit() *string
	SetEstablishedTimeout(v int32) *SetLoadBalancerTCPListenerAttributeRequest
	GetEstablishedTimeout() *int32
	SetHealthCheckConnectPort(v int32) *SetLoadBalancerTCPListenerAttributeRequest
	GetHealthCheckConnectPort() *int32
	SetHealthCheckConnectTimeout(v int32) *SetLoadBalancerTCPListenerAttributeRequest
	GetHealthCheckConnectTimeout() *int32
	SetHealthCheckDomain(v string) *SetLoadBalancerTCPListenerAttributeRequest
	GetHealthCheckDomain() *string
	SetHealthCheckHttpCode(v string) *SetLoadBalancerTCPListenerAttributeRequest
	GetHealthCheckHttpCode() *string
	SetHealthCheckInterval(v int32) *SetLoadBalancerTCPListenerAttributeRequest
	GetHealthCheckInterval() *int32
	SetHealthCheckType(v string) *SetLoadBalancerTCPListenerAttributeRequest
	GetHealthCheckType() *string
	SetHealthCheckURI(v string) *SetLoadBalancerTCPListenerAttributeRequest
	GetHealthCheckURI() *string
	SetHealthyThreshold(v int32) *SetLoadBalancerTCPListenerAttributeRequest
	GetHealthyThreshold() *int32
	SetListenerPort(v int32) *SetLoadBalancerTCPListenerAttributeRequest
	GetListenerPort() *int32
	SetLoadBalancerId(v string) *SetLoadBalancerTCPListenerAttributeRequest
	GetLoadBalancerId() *string
	SetPersistenceTimeout(v int32) *SetLoadBalancerTCPListenerAttributeRequest
	GetPersistenceTimeout() *int32
	SetScheduler(v string) *SetLoadBalancerTCPListenerAttributeRequest
	GetScheduler() *string
	SetUnhealthyThreshold(v int32) *SetLoadBalancerTCPListenerAttributeRequest
	GetUnhealthyThreshold() *int32
}

type SetLoadBalancerTCPListenerAttributeRequest struct {
	// The description of the listener. The description must be **1*	- to **80*	- characters in length.
	//
	// > The description cannot start with `http://` or `https://`.
	//
	// example:
	//
	// example
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to enable EIP pass-through. Valid values:
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
	// The port used for health checks. Valid values: **1*	- to **65535**. If this parameter is not set, the backend service port (BackendServerPort) is used.
	//
	// example:
	//
	// 8000
	HealthCheckConnectPort *int32 `json:"HealthCheckConnectPort,omitempty" xml:"HealthCheckConnectPort,omitempty"`
	// The maximum amount of time to wait for a response from a health check. If the backend ENS does not respond within the specified time, the health check fails.
	//
	// - Default value: 5 seconds.
	//
	// - Valid values: **1*	- to **300**.
	//
	// - Unit: seconds.
	//
	// > If the value of HealthCheckConnectTimeout is less than the value of HealthCheckInterval, HealthCheckConnectTimeout becomes invalid and the timeout period is set to the value of HealthCheckInterval.
	//
	// example:
	//
	// 100
	HealthCheckConnectTimeout *int32 `json:"HealthCheckConnectTimeout,omitempty" xml:"HealthCheckConnectTimeout,omitempty"`
	// The domain name used for health checks.
	//
	// example:
	//
	// www.aliyundoc.com
	HealthCheckDomain *string `json:"HealthCheckDomain,omitempty" xml:"HealthCheckDomain,omitempty"`
	// The HTTP status code for a Normal health check. Valid values:
	//
	// - **http_2xx*	- (default).
	//
	// - **http_3xx**.
	//
	// - **http_4xx**.
	//
	// - **http_5xx**.
	//
	// example:
	//
	// http_2xx
	HealthCheckHttpCode *string `json:"HealthCheckHttpCode,omitempty" xml:"HealthCheckHttpCode,omitempty"`
	// The interval between two consecutive health checks. Valid values: **1*	- to **50**. Unit: seconds.
	//
	// example:
	//
	// 5
	HealthCheckInterval *int32 `json:"HealthCheckInterval,omitempty" xml:"HealthCheckInterval,omitempty"`
	// The health check type. Valid values:
	//
	// - **tcp*	- (default).
	//
	// - **http**.
	//
	// example:
	//
	// tcp
	HealthCheckType *string `json:"HealthCheckType,omitempty" xml:"HealthCheckType,omitempty"`
	// The URI used for health checks. The URI must be **1*	- to **80*	- characters in length.
	//
	// > The URI cannot be only `/` but must start with `/`.
	//
	// example:
	//
	// /aliyundoc/index.html
	HealthCheckURI *string `json:"HealthCheckURI,omitempty" xml:"HealthCheckURI,omitempty"`
	// The number of consecutive successful health checks required before the health check status of a backend server changes from fail (unreachable) to success (reachable). Valid values: **2*	- to **10**.
	//
	// example:
	//
	// 4
	HealthyThreshold *int32 `json:"HealthyThreshold,omitempty" xml:"HealthyThreshold,omitempty"`
	// The listener port whose attributes you want to modify. Valid values: **1*	- to **65535**.
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
	// lb-5snthcyu1x10g7tywj7iu****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// The timeout period of session persistence.
	//
	// - Default value: 0, which indicates that session persistence is disabled.
	//
	// - Valid values: **0*	- to **3600**.
	//
	// - Unit: seconds.
	//
	// example:
	//
	// 0
	PersistenceTimeout *int32 `json:"PersistenceTimeout,omitempty" xml:"PersistenceTimeout,omitempty"`
	// The scheduling algorithm. Valid values:
	//
	// - **wrr**: Backend servers with higher weights receive more requests.
	//
	// - **wlc**: Requests are distributed based on the weight and actual load (number of connections) of each backend server. When the weights are the same, backend servers with fewer connections receive more requests.
	//
	// - **rr**: Requests are distributed to backend servers in sequence.
	//
	// - **sch**: Consistent hashing based on source IP addresses. Requests from the same source IP address are distributed to the same backend server.
	//
	// - **tch**: Consistent hashing based on four factors: source IP address, destination IP address, source port, and destination port. Requests from the same flow are distributed to the same backend server.
	//
	// example:
	//
	// wrr
	Scheduler *string `json:"Scheduler,omitempty" xml:"Scheduler,omitempty"`
	// The number of consecutive failed health checks required before the health check status of a backend server changes from success (reachable) to fail (unreachable). Valid values: **2*	- to **10**.
	//
	// example:
	//
	// 4
	UnhealthyThreshold *int32 `json:"UnhealthyThreshold,omitempty" xml:"UnhealthyThreshold,omitempty"`
}

func (s SetLoadBalancerTCPListenerAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s SetLoadBalancerTCPListenerAttributeRequest) GoString() string {
	return s.String()
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) GetDescription() *string {
	return s.Description
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) GetEipTransmit() *string {
	return s.EipTransmit
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) GetEstablishedTimeout() *int32 {
	return s.EstablishedTimeout
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) GetHealthCheckConnectPort() *int32 {
	return s.HealthCheckConnectPort
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) GetHealthCheckConnectTimeout() *int32 {
	return s.HealthCheckConnectTimeout
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) GetHealthCheckDomain() *string {
	return s.HealthCheckDomain
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) GetHealthCheckHttpCode() *string {
	return s.HealthCheckHttpCode
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) GetHealthCheckInterval() *int32 {
	return s.HealthCheckInterval
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) GetHealthCheckType() *string {
	return s.HealthCheckType
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) GetHealthCheckURI() *string {
	return s.HealthCheckURI
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) GetHealthyThreshold() *int32 {
	return s.HealthyThreshold
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) GetListenerPort() *int32 {
	return s.ListenerPort
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) GetPersistenceTimeout() *int32 {
	return s.PersistenceTimeout
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) GetScheduler() *string {
	return s.Scheduler
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) GetUnhealthyThreshold() *int32 {
	return s.UnhealthyThreshold
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) SetDescription(v string) *SetLoadBalancerTCPListenerAttributeRequest {
	s.Description = &v
	return s
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) SetEipTransmit(v string) *SetLoadBalancerTCPListenerAttributeRequest {
	s.EipTransmit = &v
	return s
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) SetEstablishedTimeout(v int32) *SetLoadBalancerTCPListenerAttributeRequest {
	s.EstablishedTimeout = &v
	return s
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) SetHealthCheckConnectPort(v int32) *SetLoadBalancerTCPListenerAttributeRequest {
	s.HealthCheckConnectPort = &v
	return s
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) SetHealthCheckConnectTimeout(v int32) *SetLoadBalancerTCPListenerAttributeRequest {
	s.HealthCheckConnectTimeout = &v
	return s
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) SetHealthCheckDomain(v string) *SetLoadBalancerTCPListenerAttributeRequest {
	s.HealthCheckDomain = &v
	return s
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) SetHealthCheckHttpCode(v string) *SetLoadBalancerTCPListenerAttributeRequest {
	s.HealthCheckHttpCode = &v
	return s
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) SetHealthCheckInterval(v int32) *SetLoadBalancerTCPListenerAttributeRequest {
	s.HealthCheckInterval = &v
	return s
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) SetHealthCheckType(v string) *SetLoadBalancerTCPListenerAttributeRequest {
	s.HealthCheckType = &v
	return s
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) SetHealthCheckURI(v string) *SetLoadBalancerTCPListenerAttributeRequest {
	s.HealthCheckURI = &v
	return s
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) SetHealthyThreshold(v int32) *SetLoadBalancerTCPListenerAttributeRequest {
	s.HealthyThreshold = &v
	return s
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) SetListenerPort(v int32) *SetLoadBalancerTCPListenerAttributeRequest {
	s.ListenerPort = &v
	return s
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) SetLoadBalancerId(v string) *SetLoadBalancerTCPListenerAttributeRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) SetPersistenceTimeout(v int32) *SetLoadBalancerTCPListenerAttributeRequest {
	s.PersistenceTimeout = &v
	return s
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) SetScheduler(v string) *SetLoadBalancerTCPListenerAttributeRequest {
	s.Scheduler = &v
	return s
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) SetUnhealthyThreshold(v int32) *SetLoadBalancerTCPListenerAttributeRequest {
	s.UnhealthyThreshold = &v
	return s
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) Validate() error {
	return dara.Validate(s)
}
