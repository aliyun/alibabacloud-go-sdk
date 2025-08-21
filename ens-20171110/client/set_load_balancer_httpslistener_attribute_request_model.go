// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetLoadBalancerHTTPSListenerAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *SetLoadBalancerHTTPSListenerAttributeRequest
	GetDescription() *string
	SetHealthCheck(v string) *SetLoadBalancerHTTPSListenerAttributeRequest
	GetHealthCheck() *string
	SetHealthCheckConnectPort(v int32) *SetLoadBalancerHTTPSListenerAttributeRequest
	GetHealthCheckConnectPort() *int32
	SetHealthCheckDomain(v string) *SetLoadBalancerHTTPSListenerAttributeRequest
	GetHealthCheckDomain() *string
	SetHealthCheckHttpCode(v string) *SetLoadBalancerHTTPSListenerAttributeRequest
	GetHealthCheckHttpCode() *string
	SetHealthCheckInterval(v int32) *SetLoadBalancerHTTPSListenerAttributeRequest
	GetHealthCheckInterval() *int32
	SetHealthCheckMethod(v string) *SetLoadBalancerHTTPSListenerAttributeRequest
	GetHealthCheckMethod() *string
	SetHealthCheckTimeout(v int32) *SetLoadBalancerHTTPSListenerAttributeRequest
	GetHealthCheckTimeout() *int32
	SetHealthCheckURI(v string) *SetLoadBalancerHTTPSListenerAttributeRequest
	GetHealthCheckURI() *string
	SetHealthyThreshold(v int32) *SetLoadBalancerHTTPSListenerAttributeRequest
	GetHealthyThreshold() *int32
	SetIdleTimeout(v int32) *SetLoadBalancerHTTPSListenerAttributeRequest
	GetIdleTimeout() *int32
	SetListenerPort(v int32) *SetLoadBalancerHTTPSListenerAttributeRequest
	GetListenerPort() *int32
	SetLoadBalancerId(v string) *SetLoadBalancerHTTPSListenerAttributeRequest
	GetLoadBalancerId() *string
	SetRequestTimeout(v int32) *SetLoadBalancerHTTPSListenerAttributeRequest
	GetRequestTimeout() *int32
	SetScheduler(v string) *SetLoadBalancerHTTPSListenerAttributeRequest
	GetScheduler() *string
	SetServerCertificateId(v string) *SetLoadBalancerHTTPSListenerAttributeRequest
	GetServerCertificateId() *string
	SetUnhealthyThreshold(v int32) *SetLoadBalancerHTTPSListenerAttributeRequest
	GetUnhealthyThreshold() *int32
}

type SetLoadBalancerHTTPSListenerAttributeRequest struct {
	// The name of the listener. The value must be **1*	- to **80*	- characters in length.
	//
	// >  The value cannot start with `http://` or `https://`.
	//
	// example:
	//
	// Monitoring instructions
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
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
	// >  This parameter takes effect only if you set HealthCheck to on.
	//
	// example:
	//
	// 7001
	HealthCheckConnectPort *int32 `json:"HealthCheckConnectPort,omitempty" xml:"HealthCheckConnectPort,omitempty"`
	// The domain name that you want to use for health checks.
	//
	// >  This parameter takes effect only if you set HealthCheck to on.
	//
	// example:
	//
	// www.example.com
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
	// >  This parameter takes effect only if you set HealthCheck to on.
	//
	// example:
	//
	// http_2xx
	HealthCheckHttpCode *string `json:"HealthCheckHttpCode,omitempty" xml:"HealthCheckHttpCode,omitempty"`
	// The interval at which health checks are performed. Valid values: **1*	- to **50**. Unit: seconds.
	//
	// >  This parameter takes effect only if you set HealthCheck to on.
	//
	// example:
	//
	// 2
	HealthCheckInterval *int32 `json:"HealthCheckInterval,omitempty" xml:"HealthCheckInterval,omitempty"`
	// The HTTP request method for health checks. Valid values:
	//
	// 	- **head*	- (default): requests the head of the page.
	//
	// 	- **get**: requests the specified part of the page and returns the entity body.
	//
	// >  This parameter takes effect only if the HealthCheck parameter is set to on.
	//
	// example:
	//
	// get
	HealthCheckMethod *string `json:"HealthCheckMethod,omitempty" xml:"HealthCheckMethod,omitempty"`
	// The timeout period of a health check response. If the backend ENS does not respond within the specified time, the health check fails.
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
	// 9
	HealthCheckTimeout *int32 `json:"HealthCheckTimeout,omitempty" xml:"HealthCheckTimeout,omitempty"`
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
	// /checkpreload.htm
	HealthCheckURI *string `json:"HealthCheckURI,omitempty" xml:"HealthCheckURI,omitempty"`
	// The number of consecutive successful health checks that must occur before an unhealthy and inaccessible backend server is declared healthy and accessible. Valid values: **2*	- to **10**.
	//
	// >  This parameter takes effect only if you set HealthCheck to on.
	//
	// example:
	//
	// 3
	HealthyThreshold *int32 `json:"HealthyThreshold,omitempty" xml:"HealthyThreshold,omitempty"`
	// The timeout period for idle connections. Default value: 15. Valid values: **1*	- to **60**. Unit: seconds.
	//
	// >  If no request is received within the specified timeout period, ELB closes the connection. When another request is received, ELB establishes a new connection.
	//
	// example:
	//
	// 15
	IdleTimeout *int32 `json:"IdleTimeout,omitempty" xml:"IdleTimeout,omitempty"`
	// The listener port whose attributes are to be modified. Valid values: **1*	- to **65535**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10002
	ListenerPort *int32 `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	// The ID of the Edge Load Balancer (ELB) instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// lb-5snthcyu1x10g7tywj7iu****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// The timeout period of requests. Default value: 60. Valid values: **1*	- to **180**. Unit: seconds.
	//
	// >  If no response is received from the backend server within the specified timeout period, ELB returns an HTTP 504 error code to the client.
	//
	// example:
	//
	// 60
	RequestTimeout *int32 `json:"RequestTimeout,omitempty" xml:"RequestTimeout,omitempty"`
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
	// The ID of the server certificate.
	//
	// example:
	//
	// 6027667
	ServerCertificateId *string `json:"ServerCertificateId,omitempty" xml:"ServerCertificateId,omitempty"`
	// The number of consecutive failed health checks that must occur before a healthy and accessible backend server is declared unhealthy and inaccessible. Valid values: **2*	- to **10**.
	//
	// >  This parameter takes effect only if you set HealthCheck to on.
	//
	// example:
	//
	// 3
	UnhealthyThreshold *int32 `json:"UnhealthyThreshold,omitempty" xml:"UnhealthyThreshold,omitempty"`
}

func (s SetLoadBalancerHTTPSListenerAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s SetLoadBalancerHTTPSListenerAttributeRequest) GoString() string {
	return s.String()
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) GetDescription() *string {
	return s.Description
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) GetHealthCheck() *string {
	return s.HealthCheck
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) GetHealthCheckConnectPort() *int32 {
	return s.HealthCheckConnectPort
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) GetHealthCheckDomain() *string {
	return s.HealthCheckDomain
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) GetHealthCheckHttpCode() *string {
	return s.HealthCheckHttpCode
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) GetHealthCheckInterval() *int32 {
	return s.HealthCheckInterval
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) GetHealthCheckMethod() *string {
	return s.HealthCheckMethod
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) GetHealthCheckTimeout() *int32 {
	return s.HealthCheckTimeout
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) GetHealthCheckURI() *string {
	return s.HealthCheckURI
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) GetHealthyThreshold() *int32 {
	return s.HealthyThreshold
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) GetIdleTimeout() *int32 {
	return s.IdleTimeout
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) GetListenerPort() *int32 {
	return s.ListenerPort
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) GetRequestTimeout() *int32 {
	return s.RequestTimeout
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) GetScheduler() *string {
	return s.Scheduler
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) GetServerCertificateId() *string {
	return s.ServerCertificateId
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) GetUnhealthyThreshold() *int32 {
	return s.UnhealthyThreshold
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) SetDescription(v string) *SetLoadBalancerHTTPSListenerAttributeRequest {
	s.Description = &v
	return s
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) SetHealthCheck(v string) *SetLoadBalancerHTTPSListenerAttributeRequest {
	s.HealthCheck = &v
	return s
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) SetHealthCheckConnectPort(v int32) *SetLoadBalancerHTTPSListenerAttributeRequest {
	s.HealthCheckConnectPort = &v
	return s
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) SetHealthCheckDomain(v string) *SetLoadBalancerHTTPSListenerAttributeRequest {
	s.HealthCheckDomain = &v
	return s
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) SetHealthCheckHttpCode(v string) *SetLoadBalancerHTTPSListenerAttributeRequest {
	s.HealthCheckHttpCode = &v
	return s
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) SetHealthCheckInterval(v int32) *SetLoadBalancerHTTPSListenerAttributeRequest {
	s.HealthCheckInterval = &v
	return s
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) SetHealthCheckMethod(v string) *SetLoadBalancerHTTPSListenerAttributeRequest {
	s.HealthCheckMethod = &v
	return s
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) SetHealthCheckTimeout(v int32) *SetLoadBalancerHTTPSListenerAttributeRequest {
	s.HealthCheckTimeout = &v
	return s
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) SetHealthCheckURI(v string) *SetLoadBalancerHTTPSListenerAttributeRequest {
	s.HealthCheckURI = &v
	return s
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) SetHealthyThreshold(v int32) *SetLoadBalancerHTTPSListenerAttributeRequest {
	s.HealthyThreshold = &v
	return s
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) SetIdleTimeout(v int32) *SetLoadBalancerHTTPSListenerAttributeRequest {
	s.IdleTimeout = &v
	return s
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) SetListenerPort(v int32) *SetLoadBalancerHTTPSListenerAttributeRequest {
	s.ListenerPort = &v
	return s
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) SetLoadBalancerId(v string) *SetLoadBalancerHTTPSListenerAttributeRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) SetRequestTimeout(v int32) *SetLoadBalancerHTTPSListenerAttributeRequest {
	s.RequestTimeout = &v
	return s
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) SetScheduler(v string) *SetLoadBalancerHTTPSListenerAttributeRequest {
	s.Scheduler = &v
	return s
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) SetServerCertificateId(v string) *SetLoadBalancerHTTPSListenerAttributeRequest {
	s.ServerCertificateId = &v
	return s
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) SetUnhealthyThreshold(v int32) *SetLoadBalancerHTTPSListenerAttributeRequest {
	s.UnhealthyThreshold = &v
	return s
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) Validate() error {
	return dara.Validate(s)
}
