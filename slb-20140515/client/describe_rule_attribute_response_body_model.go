// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRuleAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCookie(v string) *DescribeRuleAttributeResponseBody
	GetCookie() *string
	SetCookieTimeout(v int32) *DescribeRuleAttributeResponseBody
	GetCookieTimeout() *int32
	SetDomain(v string) *DescribeRuleAttributeResponseBody
	GetDomain() *string
	SetHealthCheck(v string) *DescribeRuleAttributeResponseBody
	GetHealthCheck() *string
	SetHealthCheckConnectPort(v int32) *DescribeRuleAttributeResponseBody
	GetHealthCheckConnectPort() *int32
	SetHealthCheckDomain(v string) *DescribeRuleAttributeResponseBody
	GetHealthCheckDomain() *string
	SetHealthCheckHttpCode(v string) *DescribeRuleAttributeResponseBody
	GetHealthCheckHttpCode() *string
	SetHealthCheckInterval(v int32) *DescribeRuleAttributeResponseBody
	GetHealthCheckInterval() *int32
	SetHealthCheckTimeout(v int32) *DescribeRuleAttributeResponseBody
	GetHealthCheckTimeout() *int32
	SetHealthCheckURI(v string) *DescribeRuleAttributeResponseBody
	GetHealthCheckURI() *string
	SetHealthyThreshold(v int32) *DescribeRuleAttributeResponseBody
	GetHealthyThreshold() *int32
	SetListenerPort(v string) *DescribeRuleAttributeResponseBody
	GetListenerPort() *string
	SetListenerSync(v string) *DescribeRuleAttributeResponseBody
	GetListenerSync() *string
	SetLoadBalancerId(v string) *DescribeRuleAttributeResponseBody
	GetLoadBalancerId() *string
	SetRequestId(v string) *DescribeRuleAttributeResponseBody
	GetRequestId() *string
	SetRuleId(v string) *DescribeRuleAttributeResponseBody
	GetRuleId() *string
	SetRuleName(v string) *DescribeRuleAttributeResponseBody
	GetRuleName() *string
	SetScheduler(v string) *DescribeRuleAttributeResponseBody
	GetScheduler() *string
	SetStickySession(v string) *DescribeRuleAttributeResponseBody
	GetStickySession() *string
	SetStickySessionType(v string) *DescribeRuleAttributeResponseBody
	GetStickySessionType() *string
	SetUnhealthyThreshold(v int32) *DescribeRuleAttributeResponseBody
	GetUnhealthyThreshold() *int32
	SetUrl(v string) *DescribeRuleAttributeResponseBody
	GetUrl() *string
	SetVServerGroupId(v string) *DescribeRuleAttributeResponseBody
	GetVServerGroupId() *string
}

type DescribeRuleAttributeResponseBody struct {
	// The cookie to be configured on the backend server.
	//
	// The cookie must be 1 to 200 characters in length and can contain ASCII letters and digits. It cannot contain commas (,), semicolons (;), or whitespace characters. It cannot start with a dollar sign ($).
	//
	// If you set the **StickySession*	- parameter to **on*	- and the **StickySessionType*	- parameter to **server**, this parameter is required.
	//
	// example:
	//
	// wwe
	Cookie *string `json:"Cookie,omitempty" xml:"Cookie,omitempty"`
	// The timeout period of a cookie.
	//
	// Valid values: **1 to 86400**. Unit: seconds.
	//
	// >  If you set the **StickySession*	- parameter to **on*	- and the **StickySessionType*	- parameter to **insert**, this parameter is required.
	//
	// example:
	//
	// 12
	CookieTimeout *int32 `json:"CookieTimeout,omitempty" xml:"CookieTimeout,omitempty"`
	// The domain name that is configured in the forwarding rule.
	//
	// example:
	//
	// www.example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// Specifies whether to enable health checks.
	//
	// Valid values: **on*	- and **off**.
	//
	// >  If you set the **ListenerSync*	- parameter to **off**, this parameter is required. If you set the parameter to **on**, the configuration of the listener is used.
	//
	// example:
	//
	// off
	HealthCheck *string `json:"HealthCheck,omitempty" xml:"HealthCheck,omitempty"`
	// The backend port that is used for health checks.
	//
	// Valid values: **1*	- to **65535**.
	//
	// >  If you set the **HealthCheck*	- parameter to **on**, this parameter is required. If this parameter is empty but **HealthCheck*	- is set to **on**, the listener port is used for health checks.
	//
	// example:
	//
	// 23
	HealthCheckConnectPort *int32 `json:"HealthCheckConnectPort,omitempty" xml:"HealthCheckConnectPort,omitempty"`
	// The domain name that is used for health checks. Valid values:
	//
	// 	- **$_ip**: The private IP address of the backend server. If the $_ip parameter is set or the HealthCheckDomain parameter is not set, SLB uses the private IP addresses of backend servers as the domain names for health checks.
	//
	// 	- **domain**: The domain name must be 1 to 80 characters in length. It can contain only letters, digits, periods (.),and hyphens (-).
	//
	// >  If you set the **HealthCheck*	- parameter to **on**, this parameter is required.
	//
	// example:
	//
	// www.example.com
	HealthCheckDomain *string `json:"HealthCheckDomain,omitempty" xml:"HealthCheckDomain,omitempty"`
	// The HTTP status code that indicates a successful health check. Separate multiple HTTP status codes with commas (,). Default value: **http_2xx**.
	//
	// Valid values: **http_2xx**, **http_3xx**, **http_4xx**, and **http_5xx**.
	//
	// >  If you set the **HealthCheck*	- parameter to **on**, this parameter is required.
	//
	// example:
	//
	// http_3xx
	HealthCheckHttpCode *string `json:"HealthCheckHttpCode,omitempty" xml:"HealthCheckHttpCode,omitempty"`
	// The time interval between two consecutive health checks.
	//
	// Valid values: **1*	- to **50**. Unit: seconds.
	//
	// >  If you set the **HealthCheck*	- parameter to **on**, this parameter is required.
	//
	// example:
	//
	// 34
	HealthCheckInterval *int32 `json:"HealthCheckInterval,omitempty" xml:"HealthCheckInterval,omitempty"`
	// The timeout period of a health check response. If a backend ECS instance does not respond within the specified timeout period, the ECS instance fails the health check.
	//
	// Valid values: **1*	- to **300**. Unit: seconds.
	//
	// >  If you set the **HealthCheck*	- parameter to **on**, this parameter is required.
	//
	// example:
	//
	// 34
	HealthCheckTimeout *int32 `json:"HealthCheckTimeout,omitempty" xml:"HealthCheckTimeout,omitempty"`
	// The URI that is used for health checks.
	//
	// >  If you set the **HealthCheck*	- parameter to **on**, this parameter is required.
	//
	// example:
	//
	// /rest
	HealthCheckURI *string `json:"HealthCheckURI,omitempty" xml:"HealthCheckURI,omitempty"`
	// The number of consecutive successful health checks that must occur before an unhealthy backend server is declared healthy. In this case, the health check state is changed from **fail*	- to **success**.
	//
	// Valid values: **2*	- to **10**.
	//
	// >  If you set the **HealthCheck*	- parameter to **on**, this parameter is required.
	//
	// example:
	//
	// 2
	HealthyThreshold *int32 `json:"HealthyThreshold,omitempty" xml:"HealthyThreshold,omitempty"`
	// The listener port that is used by the SLB instance.
	//
	// example:
	//
	// 90
	ListenerPort *string `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	// Indicates whether the forwarding rule uses the scheduling algorithm, session persistence, and health check configurations of the listener.
	//
	// Valid values: **on*	- and **off**.
	//
	// 	- **off**: does not use the configurations of the listener. You can customize health check and session persistence configurations for the forwarding rule.
	//
	// 	- **on**: uses the configurations of the listener.
	//
	// example:
	//
	// off
	ListenerSync *string `json:"ListenerSync,omitempty" xml:"ListenerSync,omitempty"`
	// The ID of the SLB instance.
	//
	// example:
	//
	// lb-bp1ca0zt07t934wxe****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 9DEC9C28-AB05-4DDF-9A78-6B08EC9CE18C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the forwarding rule.
	//
	// example:
	//
	// rule-hfgnd*****
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The name of the forwarding rule.
	//
	// example:
	//
	// Rule1
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The scheduling algorithm. Valid values:
	//
	// 	- **wrr*	- (default): Backend servers that have higher weights receive more requests than backend servers that have lower weights.
	//
	// 	- **rr**: Requests are distributed to backend servers in sequence.
	//
	// >  If you set the **ListenerSync*	- parameter to **off**, this parameter is required. If you set the parameter to **on**, the configuration of the listener is used.
	//
	// example:
	//
	// wrr
	Scheduler *string `json:"Scheduler,omitempty" xml:"Scheduler,omitempty"`
	// Indicates whether session persistence is enabled.
	//
	// Valid values: **on*	- and **off**.
	//
	// >  If you set the **ListenerSync*	- parameter to **off**, this parameter is required. If you set the parameter to **on**, the configuration of the listener is used.
	//
	// example:
	//
	// off
	StickySession *string `json:"StickySession,omitempty" xml:"StickySession,omitempty"`
	// The method that is used to handle a cookie. Valid values:
	//
	// 	- **insert**: inserts a cookie into the response. SLB inserts a cookie (SERVERID) into the first HTTP or HTTPS response packet that is sent to a client. The next request from the client will contain this cookie, and the listener will distribute this request to the recorded backend server.
	//
	// 	- **server**: rewrites a cookie. When SLB detects a user-defined cookie, SLB overwrites the original cookie with the user-defined cookie. The next request from the client contains the user-defined cookie, and the listener distributes the request to the recorded backend server.
	//
	// >  If you set the **StickySession*	- parameter to **on**, this parameter is required.
	//
	// example:
	//
	// insert
	StickySessionType *string `json:"StickySessionType,omitempty" xml:"StickySessionType,omitempty"`
	// The number of consecutive failed health checks that must occur before a healthy backend server is declared unhealthy. In this case, the health check state is changed from **success*	- to **fail**.
	//
	// Valid values: **2*	- to **10**.
	//
	// >  If you set the **HealthCheck*	- parameter to **on**, this parameter is required.
	//
	// example:
	//
	// 3
	UnhealthyThreshold *int32 `json:"UnhealthyThreshold,omitempty" xml:"UnhealthyThreshold,omitempty"`
	// The URL that is configured in the forwarding rule.
	//
	// example:
	//
	// /cache
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// The ID of the vServer group that is associated with the forwarding rule.
	//
	// example:
	//
	// rsp-cige6j****
	VServerGroupId *string `json:"VServerGroupId,omitempty" xml:"VServerGroupId,omitempty"`
}

func (s DescribeRuleAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRuleAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRuleAttributeResponseBody) GetCookie() *string {
	return s.Cookie
}

func (s *DescribeRuleAttributeResponseBody) GetCookieTimeout() *int32 {
	return s.CookieTimeout
}

func (s *DescribeRuleAttributeResponseBody) GetDomain() *string {
	return s.Domain
}

func (s *DescribeRuleAttributeResponseBody) GetHealthCheck() *string {
	return s.HealthCheck
}

func (s *DescribeRuleAttributeResponseBody) GetHealthCheckConnectPort() *int32 {
	return s.HealthCheckConnectPort
}

func (s *DescribeRuleAttributeResponseBody) GetHealthCheckDomain() *string {
	return s.HealthCheckDomain
}

func (s *DescribeRuleAttributeResponseBody) GetHealthCheckHttpCode() *string {
	return s.HealthCheckHttpCode
}

func (s *DescribeRuleAttributeResponseBody) GetHealthCheckInterval() *int32 {
	return s.HealthCheckInterval
}

func (s *DescribeRuleAttributeResponseBody) GetHealthCheckTimeout() *int32 {
	return s.HealthCheckTimeout
}

func (s *DescribeRuleAttributeResponseBody) GetHealthCheckURI() *string {
	return s.HealthCheckURI
}

func (s *DescribeRuleAttributeResponseBody) GetHealthyThreshold() *int32 {
	return s.HealthyThreshold
}

func (s *DescribeRuleAttributeResponseBody) GetListenerPort() *string {
	return s.ListenerPort
}

func (s *DescribeRuleAttributeResponseBody) GetListenerSync() *string {
	return s.ListenerSync
}

func (s *DescribeRuleAttributeResponseBody) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *DescribeRuleAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRuleAttributeResponseBody) GetRuleId() *string {
	return s.RuleId
}

func (s *DescribeRuleAttributeResponseBody) GetRuleName() *string {
	return s.RuleName
}

func (s *DescribeRuleAttributeResponseBody) GetScheduler() *string {
	return s.Scheduler
}

func (s *DescribeRuleAttributeResponseBody) GetStickySession() *string {
	return s.StickySession
}

func (s *DescribeRuleAttributeResponseBody) GetStickySessionType() *string {
	return s.StickySessionType
}

func (s *DescribeRuleAttributeResponseBody) GetUnhealthyThreshold() *int32 {
	return s.UnhealthyThreshold
}

func (s *DescribeRuleAttributeResponseBody) GetUrl() *string {
	return s.Url
}

func (s *DescribeRuleAttributeResponseBody) GetVServerGroupId() *string {
	return s.VServerGroupId
}

func (s *DescribeRuleAttributeResponseBody) SetCookie(v string) *DescribeRuleAttributeResponseBody {
	s.Cookie = &v
	return s
}

func (s *DescribeRuleAttributeResponseBody) SetCookieTimeout(v int32) *DescribeRuleAttributeResponseBody {
	s.CookieTimeout = &v
	return s
}

func (s *DescribeRuleAttributeResponseBody) SetDomain(v string) *DescribeRuleAttributeResponseBody {
	s.Domain = &v
	return s
}

func (s *DescribeRuleAttributeResponseBody) SetHealthCheck(v string) *DescribeRuleAttributeResponseBody {
	s.HealthCheck = &v
	return s
}

func (s *DescribeRuleAttributeResponseBody) SetHealthCheckConnectPort(v int32) *DescribeRuleAttributeResponseBody {
	s.HealthCheckConnectPort = &v
	return s
}

func (s *DescribeRuleAttributeResponseBody) SetHealthCheckDomain(v string) *DescribeRuleAttributeResponseBody {
	s.HealthCheckDomain = &v
	return s
}

func (s *DescribeRuleAttributeResponseBody) SetHealthCheckHttpCode(v string) *DescribeRuleAttributeResponseBody {
	s.HealthCheckHttpCode = &v
	return s
}

func (s *DescribeRuleAttributeResponseBody) SetHealthCheckInterval(v int32) *DescribeRuleAttributeResponseBody {
	s.HealthCheckInterval = &v
	return s
}

func (s *DescribeRuleAttributeResponseBody) SetHealthCheckTimeout(v int32) *DescribeRuleAttributeResponseBody {
	s.HealthCheckTimeout = &v
	return s
}

func (s *DescribeRuleAttributeResponseBody) SetHealthCheckURI(v string) *DescribeRuleAttributeResponseBody {
	s.HealthCheckURI = &v
	return s
}

func (s *DescribeRuleAttributeResponseBody) SetHealthyThreshold(v int32) *DescribeRuleAttributeResponseBody {
	s.HealthyThreshold = &v
	return s
}

func (s *DescribeRuleAttributeResponseBody) SetListenerPort(v string) *DescribeRuleAttributeResponseBody {
	s.ListenerPort = &v
	return s
}

func (s *DescribeRuleAttributeResponseBody) SetListenerSync(v string) *DescribeRuleAttributeResponseBody {
	s.ListenerSync = &v
	return s
}

func (s *DescribeRuleAttributeResponseBody) SetLoadBalancerId(v string) *DescribeRuleAttributeResponseBody {
	s.LoadBalancerId = &v
	return s
}

func (s *DescribeRuleAttributeResponseBody) SetRequestId(v string) *DescribeRuleAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRuleAttributeResponseBody) SetRuleId(v string) *DescribeRuleAttributeResponseBody {
	s.RuleId = &v
	return s
}

func (s *DescribeRuleAttributeResponseBody) SetRuleName(v string) *DescribeRuleAttributeResponseBody {
	s.RuleName = &v
	return s
}

func (s *DescribeRuleAttributeResponseBody) SetScheduler(v string) *DescribeRuleAttributeResponseBody {
	s.Scheduler = &v
	return s
}

func (s *DescribeRuleAttributeResponseBody) SetStickySession(v string) *DescribeRuleAttributeResponseBody {
	s.StickySession = &v
	return s
}

func (s *DescribeRuleAttributeResponseBody) SetStickySessionType(v string) *DescribeRuleAttributeResponseBody {
	s.StickySessionType = &v
	return s
}

func (s *DescribeRuleAttributeResponseBody) SetUnhealthyThreshold(v int32) *DescribeRuleAttributeResponseBody {
	s.UnhealthyThreshold = &v
	return s
}

func (s *DescribeRuleAttributeResponseBody) SetUrl(v string) *DescribeRuleAttributeResponseBody {
	s.Url = &v
	return s
}

func (s *DescribeRuleAttributeResponseBody) SetVServerGroupId(v string) *DescribeRuleAttributeResponseBody {
	s.VServerGroupId = &v
	return s
}

func (s *DescribeRuleAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
