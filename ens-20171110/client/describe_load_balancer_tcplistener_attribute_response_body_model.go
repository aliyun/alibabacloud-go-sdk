// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLoadBalancerTCPListenerAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBackendServerPort(v int32) *DescribeLoadBalancerTCPListenerAttributeResponseBody
	GetBackendServerPort() *int32
	SetBandwidth(v int32) *DescribeLoadBalancerTCPListenerAttributeResponseBody
	GetBandwidth() *int32
	SetDescription(v string) *DescribeLoadBalancerTCPListenerAttributeResponseBody
	GetDescription() *string
	SetEipTransmit(v string) *DescribeLoadBalancerTCPListenerAttributeResponseBody
	GetEipTransmit() *string
	SetEstablishedTimeout(v int32) *DescribeLoadBalancerTCPListenerAttributeResponseBody
	GetEstablishedTimeout() *int32
	SetHealthCheck(v string) *DescribeLoadBalancerTCPListenerAttributeResponseBody
	GetHealthCheck() *string
	SetHealthCheckConnectPort(v int32) *DescribeLoadBalancerTCPListenerAttributeResponseBody
	GetHealthCheckConnectPort() *int32
	SetHealthCheckConnectTimeout(v int32) *DescribeLoadBalancerTCPListenerAttributeResponseBody
	GetHealthCheckConnectTimeout() *int32
	SetHealthCheckDomain(v string) *DescribeLoadBalancerTCPListenerAttributeResponseBody
	GetHealthCheckDomain() *string
	SetHealthCheckHttpCode(v string) *DescribeLoadBalancerTCPListenerAttributeResponseBody
	GetHealthCheckHttpCode() *string
	SetHealthCheckInterval(v int32) *DescribeLoadBalancerTCPListenerAttributeResponseBody
	GetHealthCheckInterval() *int32
	SetHealthCheckType(v string) *DescribeLoadBalancerTCPListenerAttributeResponseBody
	GetHealthCheckType() *string
	SetHealthCheckURI(v string) *DescribeLoadBalancerTCPListenerAttributeResponseBody
	GetHealthCheckURI() *string
	SetHealthyThreshold(v int32) *DescribeLoadBalancerTCPListenerAttributeResponseBody
	GetHealthyThreshold() *int32
	SetListenerPort(v int32) *DescribeLoadBalancerTCPListenerAttributeResponseBody
	GetListenerPort() *int32
	SetPersistenceTimeout(v int32) *DescribeLoadBalancerTCPListenerAttributeResponseBody
	GetPersistenceTimeout() *int32
	SetRequestId(v string) *DescribeLoadBalancerTCPListenerAttributeResponseBody
	GetRequestId() *string
	SetScheduler(v string) *DescribeLoadBalancerTCPListenerAttributeResponseBody
	GetScheduler() *string
	SetStatus(v string) *DescribeLoadBalancerTCPListenerAttributeResponseBody
	GetStatus() *string
	SetUnhealthyThreshold(v int32) *DescribeLoadBalancerTCPListenerAttributeResponseBody
	GetUnhealthyThreshold() *int32
}

type DescribeLoadBalancerTCPListenerAttributeResponseBody struct {
	// The port used by the backend server of the load balancing instance. Valid values: **1*	- to **65535**.
	//
	// example:
	//
	// 8080
	BackendServerPort *int32 `json:"BackendServerPort,omitempty" xml:"BackendServerPort,omitempty"`
	// The peak bandwidth of the load balancing instance. Default value: -1, which indicates that no bandwidth throttling is applied.
	//
	// example:
	//
	// -1
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The name of the listener.
	//
	// example:
	//
	// example
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to enable EIP pass-through. Valid values:
	//
	// - **on**: enabled
	//
	// - **off*	- (default): disabled
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
	// Specifies whether to enable the health check feature. Valid values:
	//
	// - **on**: enabled
	//
	// - **off**: disabled
	//
	// example:
	//
	// on
	HealthCheck *string `json:"HealthCheck,omitempty" xml:"HealthCheck,omitempty"`
	// The port used for health checks. Valid values: **1*	- to **65535**. If the Settings for this parameter are not specified, the backend service port (BackendServerPort) is used.
	//
	// > This parameter takes effect only when HealthCheck is set to on.
	//
	// example:
	//
	// 8000
	HealthCheckConnectPort *int32 `json:"HealthCheckConnectPort,omitempty" xml:"HealthCheckConnectPort,omitempty"`
	// The timeout period for receiving a health check response. If a backend ENS instance does not respond within the specified timeout period, the health check fails.
	//
	// - Default value: 5.
	//
	// - Valid values: **1*	- to **300**.
	//
	// - Unit: seconds.
	//
	// > - This parameter takes effect only when HealthCheck is set to on.
	//
	// > - If the value of HealthCheckConnectTimeout is less than the value of HealthCheckInterval, HealthCheckConnectTimeout becomes invalid and the timeout period is the value of HealthCheckInterval.
	//
	// example:
	//
	// 100
	HealthCheckConnectTimeout *int32 `json:"HealthCheckConnectTimeout,omitempty" xml:"HealthCheckConnectTimeout,omitempty"`
	// The domain name used for health checks.
	//
	// > This parameter takes effect only when HealthCheck is set to on.
	//
	// example:
	//
	// example.com
	HealthCheckDomain *string `json:"HealthCheckDomain,omitempty" xml:"HealthCheckDomain,omitempty"`
	// The HTTP status code that indicates a successful health check. Valid values:
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
	// example:
	//
	// 5
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
	// The URI used for health checks. The URI must be **1*	- to **80*	- characters in length.
	//
	// > - The URI cannot contain only a forward slash (`/`), but it must start with a forward slash.
	//
	// > - This parameter takes effect only when HealthCheck is set to on.
	//
	// example:
	//
	// /example/index.html
	HealthCheckURI *string `json:"HealthCheckURI,omitempty" xml:"HealthCheckURI,omitempty"`
	// The number of consecutive successful health checks required to change the health check status of a backend server from fail (unreachable) to success (reachable). Valid values: **2*	- to **10**.
	//
	// > This parameter takes effect only when HealthCheck is set to on.
	//
	// example:
	//
	// 4
	HealthyThreshold *int32 `json:"HealthyThreshold,omitempty" xml:"HealthyThreshold,omitempty"`
	// The frontend port used by the load balancing instance.
	//
	// example:
	//
	// 80
	ListenerPort *int32 `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	// The timeout period for session persistence.
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
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The scheduling algorithm. Valid values:
	//
	// - **wrr**: Backend servers with higher weights are polled more frequently.
	//
	// - **wlc**: Backend servers are polled based on their weights and actual loads (number of connections). If weights are the same, backend servers with fewer current connections are polled more frequently.
	//
	// - **rr**: External requests are distributed to backend servers in sequence.
	//
	// - **sch**: Consistent hashing is based on source IP addresses. Requests from the same source IP address are scheduled to the same backend server.
	//
	// - **tch**: Consistent hashing is based on quadruples (source IP address, destination IP address, source port, and destination port). Requests from the same stream are scheduled to the same backend server.
	//
	// example:
	//
	// wrr
	Scheduler *string `json:"Scheduler,omitempty" xml:"Scheduler,omitempty"`
	// The status of the listener. Valid values:
	//
	// - **Running**: running normally
	//
	// - **Stopped**: stopped
	//
	// - **Starting**: starting
	//
	// - **Configuring**: configuring
	//
	// - **Stopping**: stopping
	//
	// example:
	//
	// Stopped
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The number of consecutive failed health checks required to change the health check status of a backend server from success (reachable) to fail (unreachable). Valid values: **2*	- to **10**.
	//
	// > This parameter takes effect only when HealthCheck is set to on.
	//
	// example:
	//
	// 4
	UnhealthyThreshold *int32 `json:"UnhealthyThreshold,omitempty" xml:"UnhealthyThreshold,omitempty"`
}

func (s DescribeLoadBalancerTCPListenerAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLoadBalancerTCPListenerAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) GetBackendServerPort() *int32 {
	return s.BackendServerPort
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) GetDescription() *string {
	return s.Description
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) GetEipTransmit() *string {
	return s.EipTransmit
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) GetEstablishedTimeout() *int32 {
	return s.EstablishedTimeout
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) GetHealthCheck() *string {
	return s.HealthCheck
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) GetHealthCheckConnectPort() *int32 {
	return s.HealthCheckConnectPort
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) GetHealthCheckConnectTimeout() *int32 {
	return s.HealthCheckConnectTimeout
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) GetHealthCheckDomain() *string {
	return s.HealthCheckDomain
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) GetHealthCheckHttpCode() *string {
	return s.HealthCheckHttpCode
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) GetHealthCheckInterval() *int32 {
	return s.HealthCheckInterval
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) GetHealthCheckType() *string {
	return s.HealthCheckType
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) GetHealthCheckURI() *string {
	return s.HealthCheckURI
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) GetHealthyThreshold() *int32 {
	return s.HealthyThreshold
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) GetListenerPort() *int32 {
	return s.ListenerPort
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) GetPersistenceTimeout() *int32 {
	return s.PersistenceTimeout
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) GetScheduler() *string {
	return s.Scheduler
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) GetUnhealthyThreshold() *int32 {
	return s.UnhealthyThreshold
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) SetBackendServerPort(v int32) *DescribeLoadBalancerTCPListenerAttributeResponseBody {
	s.BackendServerPort = &v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) SetBandwidth(v int32) *DescribeLoadBalancerTCPListenerAttributeResponseBody {
	s.Bandwidth = &v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) SetDescription(v string) *DescribeLoadBalancerTCPListenerAttributeResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) SetEipTransmit(v string) *DescribeLoadBalancerTCPListenerAttributeResponseBody {
	s.EipTransmit = &v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) SetEstablishedTimeout(v int32) *DescribeLoadBalancerTCPListenerAttributeResponseBody {
	s.EstablishedTimeout = &v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) SetHealthCheck(v string) *DescribeLoadBalancerTCPListenerAttributeResponseBody {
	s.HealthCheck = &v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) SetHealthCheckConnectPort(v int32) *DescribeLoadBalancerTCPListenerAttributeResponseBody {
	s.HealthCheckConnectPort = &v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) SetHealthCheckConnectTimeout(v int32) *DescribeLoadBalancerTCPListenerAttributeResponseBody {
	s.HealthCheckConnectTimeout = &v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) SetHealthCheckDomain(v string) *DescribeLoadBalancerTCPListenerAttributeResponseBody {
	s.HealthCheckDomain = &v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) SetHealthCheckHttpCode(v string) *DescribeLoadBalancerTCPListenerAttributeResponseBody {
	s.HealthCheckHttpCode = &v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) SetHealthCheckInterval(v int32) *DescribeLoadBalancerTCPListenerAttributeResponseBody {
	s.HealthCheckInterval = &v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) SetHealthCheckType(v string) *DescribeLoadBalancerTCPListenerAttributeResponseBody {
	s.HealthCheckType = &v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) SetHealthCheckURI(v string) *DescribeLoadBalancerTCPListenerAttributeResponseBody {
	s.HealthCheckURI = &v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) SetHealthyThreshold(v int32) *DescribeLoadBalancerTCPListenerAttributeResponseBody {
	s.HealthyThreshold = &v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) SetListenerPort(v int32) *DescribeLoadBalancerTCPListenerAttributeResponseBody {
	s.ListenerPort = &v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) SetPersistenceTimeout(v int32) *DescribeLoadBalancerTCPListenerAttributeResponseBody {
	s.PersistenceTimeout = &v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) SetRequestId(v string) *DescribeLoadBalancerTCPListenerAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) SetScheduler(v string) *DescribeLoadBalancerTCPListenerAttributeResponseBody {
	s.Scheduler = &v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) SetStatus(v string) *DescribeLoadBalancerTCPListenerAttributeResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) SetUnhealthyThreshold(v int32) *DescribeLoadBalancerTCPListenerAttributeResponseBody {
	s.UnhealthyThreshold = &v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
