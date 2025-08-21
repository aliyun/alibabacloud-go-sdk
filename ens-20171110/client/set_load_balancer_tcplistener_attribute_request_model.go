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
	// >  The value cannot start with `http://` or `https://`.
	//
	// example:
	//
	// example
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to enable elastic IP address (EIP) pass-through. Valid values:
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
	// The port that is used for health checks. Valid values: **1*	- to **65535**. If you leave this parameter empty, the port specified by BackendServerPort is used for health checks.
	//
	// example:
	//
	// 8000
	HealthCheckConnectPort *int32 `json:"HealthCheckConnectPort,omitempty" xml:"HealthCheckConnectPort,omitempty"`
	// The timeout period for a health check response. If the value of HealthCheckTimeout is smaller than the value of HealthCheckInterval, the timeout period specified by HealthCheckTimeout becomes invalid, and the value of HealthCheckInterval is used as the timeout period.
	//
	// 	- Default value: 5.
	//
	// 	- Valid values: **1*	- to **300**.
	//
	// 	- Unit: seconds.
	//
	// >  If the value of the HealthCheckConnectTimeout parameter is smaller than that of the HealthCheckInterval parameter, the timeout period specified by the HealthCheckConnectTimeout parameter is ignored and the period of time specified by the HealthCheckInterval parameter is used as the timeout period.
	//
	// example:
	//
	// 100
	HealthCheckConnectTimeout *int32 `json:"HealthCheckConnectTimeout,omitempty" xml:"HealthCheckConnectTimeout,omitempty"`
	// The domain name that you want to use for health checks.
	//
	// example:
	//
	// www.aliyundoc.com
	HealthCheckDomain *string `json:"HealthCheckDomain,omitempty" xml:"HealthCheckDomain,omitempty"`
	// The HTTP status code for a successful health check. Valid values:
	//
	// 	- **http_2xx*	- (default)
	//
	// 	- **http_3xx**.
	//
	// 	- **http_4xx**
	//
	// 	- **http_5xx**
	//
	// example:
	//
	// http_2xx
	HealthCheckHttpCode *string `json:"HealthCheckHttpCode,omitempty" xml:"HealthCheckHttpCode,omitempty"`
	// The interval at which health checks are performed. Valid values: **1*	- to **50**. Unit: seconds.
	//
	// example:
	//
	// 5
	HealthCheckInterval *int32 `json:"HealthCheckInterval,omitempty" xml:"HealthCheckInterval,omitempty"`
	// The type of health checks. Valid values:
	//
	// 	- **tcp*	- (default)
	//
	// 	- **http**
	//
	// example:
	//
	// tcp
	HealthCheckType *string `json:"HealthCheckType,omitempty" xml:"HealthCheckType,omitempty"`
	// The Uniform Resource Identifier (URI) that you want to use for health checks. The URI must be **1*	- to **80*	- characters in length.
	//
	// >  The URL must start with `/` and contain characters other than `/`.
	//
	// example:
	//
	// /aliyundoc/index.html
	HealthCheckURI *string `json:"HealthCheckURI,omitempty" xml:"HealthCheckURI,omitempty"`
	// The number of consecutive successful health checks that must occur before an unhealthy and inaccessible backend server is declared healthy and accessible. Valid values: **2*	- to **10**.
	//
	// example:
	//
	// 4
	HealthyThreshold *int32 `json:"HealthyThreshold,omitempty" xml:"HealthyThreshold,omitempty"`
	// The listener port whose attributes are to be modified. Valid values: **1*	- to **65535**.
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
	// lb-5snthcyu1x10g7tywj7iu****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// The timeout period of session persistence.
	//
	// 	- Default value: 0. If the default value is used, the system disables session persistence.
	//
	// 	- Valid values: **0*	- to **3600**.
	//
	// 	- Unit: seconds.
	//
	// example:
	//
	// 0
	PersistenceTimeout *int32 `json:"PersistenceTimeout,omitempty" xml:"PersistenceTimeout,omitempty"`
	// The scheduling algorithm. Valid values:
	//
	// 	- **wrr**: Backend servers with higher weights receive more requests than those with lower weights.
	//
	// 	- **wlc**: Requests are distributed based on the weights and number of connections to backend servers. If two backend servers have the same weight, the backend server that has fewer connections receives more requests.
	//
	// 	- **rr**: Requests are distributed to backend servers in sequence.
	//
	// 	- **sch**: consistent hashing based on source IP addresses. Requests from the same source IP address are distributed to the same backend server.
	//
	// 	- **qch**: consistent hashing based on QUIC connection IDs (CIDs). Requests that contain the same QUIC CID are distributed to the same backend server.
	//
	// 	- **iqch**: consistent hashing based on three specific bytes of iQUIC CIDs. Requests with the same second, third, and fourth bytes are distributed to the same backend server.
	//
	// example:
	//
	// wrr
	Scheduler *string `json:"Scheduler,omitempty" xml:"Scheduler,omitempty"`
	// The number of consecutive failed health checks that must occur before a healthy and accessible backend server is declared unhealthy and inaccessible. Valid values: **2*	- to **10**.
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
