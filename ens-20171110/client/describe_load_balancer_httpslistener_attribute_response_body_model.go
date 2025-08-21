// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLoadBalancerHTTPSListenerAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBackendServerPort(v int32) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody
	GetBackendServerPort() *int32
	SetBandwidth(v int32) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody
	GetBandwidth() *int32
	SetDescription(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody
	GetDescription() *string
	SetForwardPort(v int32) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody
	GetForwardPort() *int32
	SetHealthCheck(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody
	GetHealthCheck() *string
	SetHealthCheckConnectPort(v int32) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody
	GetHealthCheckConnectPort() *int32
	SetHealthCheckDomain(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody
	GetHealthCheckDomain() *string
	SetHealthCheckHttpCode(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody
	GetHealthCheckHttpCode() *string
	SetHealthCheckInterval(v int32) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody
	GetHealthCheckInterval() *int32
	SetHealthCheckMethod(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody
	GetHealthCheckMethod() *string
	SetHealthCheckTimeout(v int32) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody
	GetHealthCheckTimeout() *int32
	SetHealthCheckURI(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody
	GetHealthCheckURI() *string
	SetHealthyThreshold(v int32) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody
	GetHealthyThreshold() *int32
	SetIdleTimeout(v int32) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody
	GetIdleTimeout() *int32
	SetListenerForward(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody
	GetListenerForward() *string
	SetListenerPort(v int32) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody
	GetListenerPort() *int32
	SetRequestId(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody
	GetRequestId() *string
	SetRequestTimeout(v int32) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody
	GetRequestTimeout() *int32
	SetScheduler(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody
	GetScheduler() *string
	SetServerCertificateId(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody
	GetServerCertificateId() *string
	SetStatus(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody
	GetStatus() *string
	SetUnhealthyThreshold(v int32) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody
	GetUnhealthyThreshold() *int32
}

type DescribeLoadBalancerHTTPSListenerAttributeResponseBody struct {
	// The port used by the backend server of the ELB instance. Valid values: **1*	- to **65535**.
	//
	// example:
	//
	// 8080
	BackendServerPort *int32 `json:"BackendServerPort,omitempty" xml:"BackendServerPort,omitempty"`
	// The peak bandwidth of the Edge Load Balancer (ELB). The default value is -1, which indicates that the bandwidth is not limited.
	//
	// example:
	//
	// 50
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The description of the listener. The description must be **1*	- to **80*	- characters in length.
	//
	// example:
	//
	// abc
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The listener port that is used to redirect HTTP requests to HTTPS.
	//
	// example:
	//
	// 0
	ForwardPort *int32 `json:"ForwardPort,omitempty" xml:"ForwardPort,omitempty"`
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
	// >  This parameter is returned only if the HealthCheck parameter is set to on.
	//
	// example:
	//
	// 9902
	HealthCheckConnectPort *int32 `json:"HealthCheckConnectPort,omitempty" xml:"HealthCheckConnectPort,omitempty"`
	// The domain name that is used for health checks.
	//
	// >  This parameter is returned only if the HealthCheck parameter is set to on.
	//
	// example:
	//
	// www.test.com
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
	// >  This parameter is returned only if the HealthCheck parameter is set to on.
	//
	// example:
	//
	// http_2xx
	HealthCheckHttpCode *string `json:"HealthCheckHttpCode,omitempty" xml:"HealthCheckHttpCode,omitempty"`
	// The interval at which health checks are performed. Valid values: **1*	- to **50**. Unit: seconds.
	//
	// >  This parameter is returned only if the HealthCheck parameter is set to on.
	//
	// example:
	//
	// 5
	HealthCheckInterval *int32 `json:"HealthCheckInterval,omitempty" xml:"HealthCheckInterval,omitempty"`
	// The health check method used in HTTP health checks. Valid values:
	//
	// 	- **head**: requests the head of the page.
	//
	// 	- **get**: requests the specified part of the page and returns the entity body.
	//
	// >  This parameter is returned only if the HealthCheck parameter is set to on.
	//
	// example:
	//
	// head
	HealthCheckMethod *string `json:"HealthCheckMethod,omitempty" xml:"HealthCheckMethod,omitempty"`
	// The timeout period of a health check response. If a backend server does not respond within the specified timeout period, the server fails the health check.
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
	// 	- If the value of HealthCheckTimeout is smaller than the value of HealthCheckInterval, the timeout period specified by HealthCheckTimeout becomes invalid, and the value of HealthCheckInterval is used as the timeout period.
	//
	// example:
	//
	// 5
	HealthCheckTimeout *int32 `json:"HealthCheckTimeout,omitempty" xml:"HealthCheckTimeout,omitempty"`
	// The URI that is used for health checks. The URI must be **1*	- to **80*	- characters in length.
	//
	// 	- The URL must start with `/` and contain characters other than `/`.
	//
	// 	- This parameter is returned only if the HealthCheck parameter is set to on.
	//
	// example:
	//
	// /checkpreload.htm
	HealthCheckURI *string `json:"HealthCheckURI,omitempty" xml:"HealthCheckURI,omitempty"`
	// The number of consecutive successful health checks that must occur before an unhealthy and inaccessible backend server is declared healthy and accessible. Valid values: **2*	- to **10**.
	//
	// >  This parameter is returned only if the HealthCheck parameter is set to on.
	//
	// example:
	//
	// 3
	HealthyThreshold *int32 `json:"HealthyThreshold,omitempty" xml:"HealthyThreshold,omitempty"`
	// The timeout period for idle connections. Default value: 15. Valid values: **1*	- to **60**. Unit: seconds.
	//
	// >  If no request is received within the specified timeout period, ELB closes the connection. When another request is received, CLB establishes a new connection.
	//
	// example:
	//
	// 15
	IdleTimeout *int32 `json:"IdleTimeout,omitempty" xml:"IdleTimeout,omitempty"`
	// Indicates whether HTTP-to-HTTPS redirection is enabled. Valid values:
	//
	// 	- on
	//
	// 	- off
	//
	// example:
	//
	// off
	ListenerForward *string `json:"ListenerForward,omitempty" xml:"ListenerForward,omitempty"`
	// The frontend port that is used by the ELB instance. Valid values: **1*	- to **65535**.
	//
	// example:
	//
	// 8080
	ListenerPort *int32 `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The timeout period of requests. Default value: 60. Valid values: **1*	- to **180**. Unit: seconds.
	//
	// >  If no response is received from the backend server within the specified timeout period, ELB returns an HTTP 504 error code to the client.
	//
	// example:
	//
	// 60
	RequestTimeout *int32 `json:"RequestTimeout,omitempty" xml:"RequestTimeout,omitempty"`
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
	// The ID of the server certificate.
	//
	// example:
	//
	// 60276**
	ServerCertificateId *string `json:"ServerCertificateId,omitempty" xml:"ServerCertificateId,omitempty"`
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
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The number of consecutive failed health checks that must occur before a healthy and accessible backend server is declared unhealthy and inaccessible. Valid values: **2*	- to **10**.
	//
	// >  This parameter is returned only if the HealthCheck parameter is set to on.
	//
	// example:
	//
	// 3
	UnhealthyThreshold *int32 `json:"UnhealthyThreshold,omitempty" xml:"UnhealthyThreshold,omitempty"`
}

func (s DescribeLoadBalancerHTTPSListenerAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLoadBalancerHTTPSListenerAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) GetBackendServerPort() *int32 {
	return s.BackendServerPort
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) GetDescription() *string {
	return s.Description
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) GetForwardPort() *int32 {
	return s.ForwardPort
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) GetHealthCheck() *string {
	return s.HealthCheck
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) GetHealthCheckConnectPort() *int32 {
	return s.HealthCheckConnectPort
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) GetHealthCheckDomain() *string {
	return s.HealthCheckDomain
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) GetHealthCheckHttpCode() *string {
	return s.HealthCheckHttpCode
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) GetHealthCheckInterval() *int32 {
	return s.HealthCheckInterval
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) GetHealthCheckMethod() *string {
	return s.HealthCheckMethod
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) GetHealthCheckTimeout() *int32 {
	return s.HealthCheckTimeout
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) GetHealthCheckURI() *string {
	return s.HealthCheckURI
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) GetHealthyThreshold() *int32 {
	return s.HealthyThreshold
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) GetIdleTimeout() *int32 {
	return s.IdleTimeout
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) GetListenerForward() *string {
	return s.ListenerForward
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) GetListenerPort() *int32 {
	return s.ListenerPort
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) GetRequestTimeout() *int32 {
	return s.RequestTimeout
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) GetScheduler() *string {
	return s.Scheduler
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) GetServerCertificateId() *string {
	return s.ServerCertificateId
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) GetUnhealthyThreshold() *int32 {
	return s.UnhealthyThreshold
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) SetBackendServerPort(v int32) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody {
	s.BackendServerPort = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) SetBandwidth(v int32) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody {
	s.Bandwidth = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) SetDescription(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) SetForwardPort(v int32) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody {
	s.ForwardPort = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) SetHealthCheck(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody {
	s.HealthCheck = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) SetHealthCheckConnectPort(v int32) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody {
	s.HealthCheckConnectPort = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) SetHealthCheckDomain(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody {
	s.HealthCheckDomain = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) SetHealthCheckHttpCode(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody {
	s.HealthCheckHttpCode = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) SetHealthCheckInterval(v int32) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody {
	s.HealthCheckInterval = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) SetHealthCheckMethod(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody {
	s.HealthCheckMethod = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) SetHealthCheckTimeout(v int32) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody {
	s.HealthCheckTimeout = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) SetHealthCheckURI(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody {
	s.HealthCheckURI = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) SetHealthyThreshold(v int32) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody {
	s.HealthyThreshold = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) SetIdleTimeout(v int32) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody {
	s.IdleTimeout = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) SetListenerForward(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody {
	s.ListenerForward = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) SetListenerPort(v int32) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody {
	s.ListenerPort = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) SetRequestId(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) SetRequestTimeout(v int32) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody {
	s.RequestTimeout = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) SetScheduler(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody {
	s.Scheduler = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) SetServerCertificateId(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody {
	s.ServerCertificateId = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) SetStatus(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) SetUnhealthyThreshold(v int32) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody {
	s.UnhealthyThreshold = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
