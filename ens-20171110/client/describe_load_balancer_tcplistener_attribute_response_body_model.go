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
	// The port used by the backend server of the ELB instance. Valid values: **1*	- to **65535**.
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
	// Indicates whether the health check feature is enabled. Valid values:
	//
	// 	- **on**
	//
	// 	- **off**
	//
	// example:
	//
	// on
	HealthCheck *string `json:"HealthCheck,omitempty" xml:"HealthCheck,omitempty"`
	// The port that is used for health checks. Valid values: **1*	- to **65535**. If an empty string is returned for this parameter, the port specified by BackendServerPort is used for health checks.
	//
	// >  This parameter is returned only if you set HealthCheck to on.
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
	// The domain name that is used for health checks.
	//
	// >  This parameter is returned only if you set HealthCheck to on.
	//
	// example:
	//
	// example.com
	HealthCheckDomain *string `json:"HealthCheckDomain,omitempty" xml:"HealthCheckDomain,omitempty"`
	// The HTTP status codes for a successful health check. Valid values:
	//
	// 	- **http_2xx*	- (default)
	//
	// 	- **http_3xx**
	//
	// 	- **http_4xx**
	//
	// 	- **http_5xx**
	//
	// >  This parameter is returned only if you set HealthCheck to on.
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
	// The URI used for health checks. The URI must be **1*	- to **80*	- characters in length.
	//
	// >
	//
	// 	- A URL must start with a forward slash (`/`) but cannot contain only forward slashes (`/`).
	//
	// 	- This parameter takes effect only if the HealthCheck parameter is set to on.
	//
	// example:
	//
	// /example/index.html
	HealthCheckURI *string `json:"HealthCheckURI,omitempty" xml:"HealthCheckURI,omitempty"`
	// The number of consecutive successful health checks that must occur before an unhealthy and inaccessible backend server is declared healthy and accessible. Valid values: **2*	- to **10**.
	//
	// >  This parameter is returned only if you set HealthCheck to on.
	//
	// example:
	//
	// 4
	HealthyThreshold *int32 `json:"HealthyThreshold,omitempty" xml:"HealthyThreshold,omitempty"`
	// The frontend port that is used by the ELB instance.
	//
	// example:
	//
	// 80
	ListenerPort *int32 `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
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
	// The ID of the request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
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
	// >  This parameter is returned only if you set HealthCheck to on.
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
