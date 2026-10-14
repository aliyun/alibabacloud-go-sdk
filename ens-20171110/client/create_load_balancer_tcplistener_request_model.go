// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLoadBalancerTCPListenerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackendServerPort(v int32) *CreateLoadBalancerTCPListenerRequest
	GetBackendServerPort() *int32
	SetDescription(v string) *CreateLoadBalancerTCPListenerRequest
	GetDescription() *string
	SetEipTransmit(v string) *CreateLoadBalancerTCPListenerRequest
	GetEipTransmit() *string
	SetEstablishedTimeout(v int32) *CreateLoadBalancerTCPListenerRequest
	GetEstablishedTimeout() *int32
	SetHealthCheckConnectPort(v int32) *CreateLoadBalancerTCPListenerRequest
	GetHealthCheckConnectPort() *int32
	SetHealthCheckConnectTimeout(v int32) *CreateLoadBalancerTCPListenerRequest
	GetHealthCheckConnectTimeout() *int32
	SetHealthCheckDomain(v string) *CreateLoadBalancerTCPListenerRequest
	GetHealthCheckDomain() *string
	SetHealthCheckHttpCode(v string) *CreateLoadBalancerTCPListenerRequest
	GetHealthCheckHttpCode() *string
	SetHealthCheckInterval(v int32) *CreateLoadBalancerTCPListenerRequest
	GetHealthCheckInterval() *int32
	SetHealthCheckType(v string) *CreateLoadBalancerTCPListenerRequest
	GetHealthCheckType() *string
	SetHealthCheckURI(v string) *CreateLoadBalancerTCPListenerRequest
	GetHealthCheckURI() *string
	SetHealthyThreshold(v int32) *CreateLoadBalancerTCPListenerRequest
	GetHealthyThreshold() *int32
	SetListenerPort(v int32) *CreateLoadBalancerTCPListenerRequest
	GetListenerPort() *int32
	SetLoadBalancerId(v string) *CreateLoadBalancerTCPListenerRequest
	GetLoadBalancerId() *string
	SetPersistenceTimeout(v int32) *CreateLoadBalancerTCPListenerRequest
	GetPersistenceTimeout() *int32
	SetScheduler(v string) *CreateLoadBalancerTCPListenerRequest
	GetScheduler() *string
	SetUnhealthyThreshold(v int32) *CreateLoadBalancerTCPListenerRequest
	GetUnhealthyThreshold() *int32
}

type CreateLoadBalancerTCPListenerRequest struct {
	// The port used by backend servers. Valid values: **1*	- to **65535**.
	//
	// example:
	//
	// 8080
	BackendServerPort *int32 `json:"BackendServerPort,omitempty" xml:"BackendServerPort,omitempty"`
	// The description of the listener. The description must be **1*	- to **80*	- characters in length.
	//
	// > The description cannot start with `http://` or `https://`.
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
	// The timeout period for established connections, in seconds. Valid values: **10*	- to **900**. Default value: **900**.
	//
	// example:
	//
	// 500
	EstablishedTimeout *int32 `json:"EstablishedTimeout,omitempty" xml:"EstablishedTimeout,omitempty"`
	// The port used for health checks. Valid values: **1*	- to **65535**. If you do not specify this parameter, the port specified by `BackendServerPort` is used by default.
	//
	// example:
	//
	// 8080
	HealthCheckConnectPort *int32 `json:"HealthCheckConnectPort,omitempty" xml:"HealthCheckConnectPort,omitempty"`
	// The maximum time to wait for a health check response. If a backend ENS instance does not respond within this timeout period, the health check fails.
	//
	// - Default value: 5.
	//
	// - Valid values: **1*	- to **300**.
	//
	// - Unit: seconds.
	//
	// > If the value of this parameter is less than the value of `HealthCheckInterval`, this parameter is ignored and the value of `HealthCheckInterval` is used as the timeout period.
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
	// example:
	//
	// http_2xx
	HealthCheckHttpCode *string `json:"HealthCheckHttpCode,omitempty" xml:"HealthCheckHttpCode,omitempty"`
	// The interval between health checks, in seconds. Valid values: **1*	- to **50**. Default value: **2**.
	//
	// example:
	//
	// 3
	HealthCheckInterval *int32 `json:"HealthCheckInterval,omitempty" xml:"HealthCheckInterval,omitempty"`
	// The health check type. Valid values:
	//
	// - **tcp*	- (default)
	//
	// - **http**
	//
	// example:
	//
	// tcp
	HealthCheckType *string `json:"HealthCheckType,omitempty" xml:"HealthCheckType,omitempty"`
	// The URI for health checks. The URI must be **1*	- to **80*	- characters in length.
	//
	// > The URI must start with a forward slash (/) but cannot be only a forward slash (/).
	//
	// example:
	//
	// /checkpreload.htm
	HealthCheckURI *string `json:"HealthCheckURI,omitempty" xml:"HealthCheckURI,omitempty"`
	// The number of consecutive successful health checks required for a backend server to be considered healthy (from a `fail` to `success` state). Valid values: **2*	- to **10**. Default value: **3**.
	//
	// example:
	//
	// 4
	HealthyThreshold *int32 `json:"HealthyThreshold,omitempty" xml:"HealthyThreshold,omitempty"`
	// The listener port that receives and forwards requests to backend servers. Valid values: **1*	- to **65535**.
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
	// lb-5ovkn1piwqmoqrfjdyhq4****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// The timeout period for session persistence.
	//
	// - Default value: **0**, which indicates that session persistence is disabled.
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
	// - **wrr*	- (default): Weighted Round Robin. Requests are distributed to backend servers based on their weights. Backend servers with higher weights receive more requests.
	//
	// - **wlc**: Weighted Least Connections. Requests are distributed to backend servers based on both their weights and the number of connections. If multiple backend servers have the same weight, the one with the fewest connections receives more requests.
	//
	// - **rr**: Round Robin. Requests are distributed to backend servers in sequential order.
	//
	// - **sch**: Source IP Hash. Requests are routed based on a consistent hash of the source IP address, ensuring that requests from the same source IP address are sent to the same backend server.
	//
	// - **tch**: Four-tuple Hash. Requests are routed based on a consistent hash of a four-tuple (source IP address, destination IP address, source port, and destination port), ensuring that requests from the same flow are sent to the same backend server.
	//
	// example:
	//
	// wrr
	Scheduler *string `json:"Scheduler,omitempty" xml:"Scheduler,omitempty"`
	// The number of consecutive failed health checks required for a backend server to be considered unhealthy (from a `success` to `fail` state). Valid values: **2*	- to **10**. Default value: **3**.
	//
	// example:
	//
	// 4
	UnhealthyThreshold *int32 `json:"UnhealthyThreshold,omitempty" xml:"UnhealthyThreshold,omitempty"`
}

func (s CreateLoadBalancerTCPListenerRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateLoadBalancerTCPListenerRequest) GoString() string {
	return s.String()
}

func (s *CreateLoadBalancerTCPListenerRequest) GetBackendServerPort() *int32 {
	return s.BackendServerPort
}

func (s *CreateLoadBalancerTCPListenerRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateLoadBalancerTCPListenerRequest) GetEipTransmit() *string {
	return s.EipTransmit
}

func (s *CreateLoadBalancerTCPListenerRequest) GetEstablishedTimeout() *int32 {
	return s.EstablishedTimeout
}

func (s *CreateLoadBalancerTCPListenerRequest) GetHealthCheckConnectPort() *int32 {
	return s.HealthCheckConnectPort
}

func (s *CreateLoadBalancerTCPListenerRequest) GetHealthCheckConnectTimeout() *int32 {
	return s.HealthCheckConnectTimeout
}

func (s *CreateLoadBalancerTCPListenerRequest) GetHealthCheckDomain() *string {
	return s.HealthCheckDomain
}

func (s *CreateLoadBalancerTCPListenerRequest) GetHealthCheckHttpCode() *string {
	return s.HealthCheckHttpCode
}

func (s *CreateLoadBalancerTCPListenerRequest) GetHealthCheckInterval() *int32 {
	return s.HealthCheckInterval
}

func (s *CreateLoadBalancerTCPListenerRequest) GetHealthCheckType() *string {
	return s.HealthCheckType
}

func (s *CreateLoadBalancerTCPListenerRequest) GetHealthCheckURI() *string {
	return s.HealthCheckURI
}

func (s *CreateLoadBalancerTCPListenerRequest) GetHealthyThreshold() *int32 {
	return s.HealthyThreshold
}

func (s *CreateLoadBalancerTCPListenerRequest) GetListenerPort() *int32 {
	return s.ListenerPort
}

func (s *CreateLoadBalancerTCPListenerRequest) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *CreateLoadBalancerTCPListenerRequest) GetPersistenceTimeout() *int32 {
	return s.PersistenceTimeout
}

func (s *CreateLoadBalancerTCPListenerRequest) GetScheduler() *string {
	return s.Scheduler
}

func (s *CreateLoadBalancerTCPListenerRequest) GetUnhealthyThreshold() *int32 {
	return s.UnhealthyThreshold
}

func (s *CreateLoadBalancerTCPListenerRequest) SetBackendServerPort(v int32) *CreateLoadBalancerTCPListenerRequest {
	s.BackendServerPort = &v
	return s
}

func (s *CreateLoadBalancerTCPListenerRequest) SetDescription(v string) *CreateLoadBalancerTCPListenerRequest {
	s.Description = &v
	return s
}

func (s *CreateLoadBalancerTCPListenerRequest) SetEipTransmit(v string) *CreateLoadBalancerTCPListenerRequest {
	s.EipTransmit = &v
	return s
}

func (s *CreateLoadBalancerTCPListenerRequest) SetEstablishedTimeout(v int32) *CreateLoadBalancerTCPListenerRequest {
	s.EstablishedTimeout = &v
	return s
}

func (s *CreateLoadBalancerTCPListenerRequest) SetHealthCheckConnectPort(v int32) *CreateLoadBalancerTCPListenerRequest {
	s.HealthCheckConnectPort = &v
	return s
}

func (s *CreateLoadBalancerTCPListenerRequest) SetHealthCheckConnectTimeout(v int32) *CreateLoadBalancerTCPListenerRequest {
	s.HealthCheckConnectTimeout = &v
	return s
}

func (s *CreateLoadBalancerTCPListenerRequest) SetHealthCheckDomain(v string) *CreateLoadBalancerTCPListenerRequest {
	s.HealthCheckDomain = &v
	return s
}

func (s *CreateLoadBalancerTCPListenerRequest) SetHealthCheckHttpCode(v string) *CreateLoadBalancerTCPListenerRequest {
	s.HealthCheckHttpCode = &v
	return s
}

func (s *CreateLoadBalancerTCPListenerRequest) SetHealthCheckInterval(v int32) *CreateLoadBalancerTCPListenerRequest {
	s.HealthCheckInterval = &v
	return s
}

func (s *CreateLoadBalancerTCPListenerRequest) SetHealthCheckType(v string) *CreateLoadBalancerTCPListenerRequest {
	s.HealthCheckType = &v
	return s
}

func (s *CreateLoadBalancerTCPListenerRequest) SetHealthCheckURI(v string) *CreateLoadBalancerTCPListenerRequest {
	s.HealthCheckURI = &v
	return s
}

func (s *CreateLoadBalancerTCPListenerRequest) SetHealthyThreshold(v int32) *CreateLoadBalancerTCPListenerRequest {
	s.HealthyThreshold = &v
	return s
}

func (s *CreateLoadBalancerTCPListenerRequest) SetListenerPort(v int32) *CreateLoadBalancerTCPListenerRequest {
	s.ListenerPort = &v
	return s
}

func (s *CreateLoadBalancerTCPListenerRequest) SetLoadBalancerId(v string) *CreateLoadBalancerTCPListenerRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *CreateLoadBalancerTCPListenerRequest) SetPersistenceTimeout(v int32) *CreateLoadBalancerTCPListenerRequest {
	s.PersistenceTimeout = &v
	return s
}

func (s *CreateLoadBalancerTCPListenerRequest) SetScheduler(v string) *CreateLoadBalancerTCPListenerRequest {
	s.Scheduler = &v
	return s
}

func (s *CreateLoadBalancerTCPListenerRequest) SetUnhealthyThreshold(v int32) *CreateLoadBalancerTCPListenerRequest {
	s.UnhealthyThreshold = &v
	return s
}

func (s *CreateLoadBalancerTCPListenerRequest) Validate() error {
	return dara.Validate(s)
}
