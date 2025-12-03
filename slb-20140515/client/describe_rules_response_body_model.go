// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeRulesResponseBody
	GetRequestId() *string
	SetRules(v *DescribeRulesResponseBodyRules) *DescribeRulesResponseBody
	GetRules() *DescribeRulesResponseBodyRules
}

type DescribeRulesResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 9DEC9C28-AB05-4DDF-9A78-6B08EC9CE18C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The forwarding rules.
	Rules *DescribeRulesResponseBodyRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Struct"`
}

func (s DescribeRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRulesResponseBody) GetRules() *DescribeRulesResponseBodyRules {
	return s.Rules
}

func (s *DescribeRulesResponseBody) SetRequestId(v string) *DescribeRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRulesResponseBody) SetRules(v *DescribeRulesResponseBodyRules) *DescribeRulesResponseBody {
	s.Rules = v
	return s
}

func (s *DescribeRulesResponseBody) Validate() error {
	if s.Rules != nil {
		if err := s.Rules.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeRulesResponseBodyRules struct {
	Rule []*DescribeRulesResponseBodyRulesRule `json:"Rule,omitempty" xml:"Rule,omitempty" type:"Repeated"`
}

func (s DescribeRulesResponseBodyRules) String() string {
	return dara.Prettify(s)
}

func (s DescribeRulesResponseBodyRules) GoString() string {
	return s.String()
}

func (s *DescribeRulesResponseBodyRules) GetRule() []*DescribeRulesResponseBodyRulesRule {
	return s.Rule
}

func (s *DescribeRulesResponseBodyRules) SetRule(v []*DescribeRulesResponseBodyRulesRule) *DescribeRulesResponseBodyRules {
	s.Rule = v
	return s
}

func (s *DescribeRulesResponseBodyRules) Validate() error {
	if s.Rule != nil {
		for _, item := range s.Rule {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRulesResponseBodyRulesRule struct {
	// The cookie that is configured on the backend server.
	//
	// The value must be 1 to 200 characters in length, and can contain only ASCII letters and digits. It cannot contain commas (,), semicolons (;), or spaces. It cannot start with a dollar sign ($).
	//
	// >  If you set the **StickySession*	- parameter to **on*	- and the **StickySessionType*	- parameter to **server**, this parameter is required.
	//
	// example:
	//
	// 23
	Cookie *string `json:"Cookie,omitempty" xml:"Cookie,omitempty"`
	// The timeout period of a cookie. Valid values: **1 to 86400**. Unit: seconds.
	//
	// >  If you set the **StickySession*	- parameter to **on*	- and the **StickySessionType*	- parameter to **insert**, this parameter is required.
	//
	// example:
	//
	// 56
	CookieTimeout *int32 `json:"CookieTimeout,omitempty" xml:"CookieTimeout,omitempty"`
	// The requested domain name specified in the forwarding rule.
	//
	// example:
	//
	// test.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// Indicates whether health checks are enabled.
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
	// Valid values: **1 to 65535**.
	//
	// >  If you set the **HealthCheck*	- parameter to **on**, this parameter is required. If this parameter is empty but **HealthCheck*	- is set to **on**, the listener port is used for health checks.
	//
	// example:
	//
	// 45
	HealthCheckConnectPort *int32 `json:"HealthCheckConnectPort,omitempty" xml:"HealthCheckConnectPort,omitempty"`
	// The domain name that is used for health checks. Valid values:
	//
	// 	- **$_ip**: The private IP address of the backend server.
	//
	//     If you do not set this parameter or set the parameter to $_ip, the SLB instance uses the private IP address of each backend server as the domain name for health checks.
	//
	// 	- **domain**: The domain name must be 1 to 80 characters in length. The domain name can contain only letters, digits, periods (.),and hyphens (-).
	//
	// >  If you set the **HealthCheck*	- parameter to **on**, this parameter is required.
	//
	// example:
	//
	// www.domain.com
	HealthCheckDomain *string `json:"HealthCheckDomain,omitempty" xml:"HealthCheckDomain,omitempty"`
	// The HTTP status code that indicates a successful health check. Multiple HTTP status codes are separated by commas (,). Default value: **http_2xx**.
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
	// Valid values: **1 to 50**. Unit: seconds.
	//
	// >  If you set the **HealthCheck*	- parameter to **on**, this parameter is required.
	//
	// example:
	//
	// 5
	HealthCheckInterval *int32 `json:"HealthCheckInterval,omitempty" xml:"HealthCheckInterval,omitempty"`
	// The timeout period of a health check response. If a backend ECS instance does not respond within the specified timeout period, the ECS instance fails the health check. Unit: seconds
	//
	// Valid values: **1 to 300**.
	//
	// >  When you set the **HealthCheck*	- parameter to **on**, this parameter takes effect.
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
	// /example
	HealthCheckURI *string `json:"HealthCheckURI,omitempty" xml:"HealthCheckURI,omitempty"`
	// Specifies the number of successful health checks that must be consecutively performed before a backend server can be declared healthy (from **fail*	- to **success**).
	//
	// Valid values: **2 to 10**.
	//
	// >  If you set the **HealthCheck*	- parameter to **on**, this parameter is required.
	//
	// example:
	//
	// 5
	HealthyThreshold *int32 `json:"HealthyThreshold,omitempty" xml:"HealthyThreshold,omitempty"`
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
	// The ID of the forwarding rule.
	//
	// example:
	//
	// rule-tybqi6****
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The name of the forwarding rule. The name must be 1 to 80 characters in length, and can contain only letters, digits, hyphens (-), forward slashes (/), periods (.),and underscores (_).
	//
	// >  The name of each forwarding rule must be unique within a listener.
	//
	// example:
	//
	// Rule2
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
	// Specifies whether to enable session persistence.
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
	// Specifies the number of times that a healthy backend server must consecutively fail health checks before it is declared unhealthy (from **success*	- to **fail**).
	//
	// Valid values: **2 to 10**
	//
	// >  If you set the **HealthCheck*	- parameter to **on**, this parameter is required.
	//
	// example:
	//
	// 2
	UnhealthyThreshold *int32 `json:"UnhealthyThreshold,omitempty" xml:"UnhealthyThreshold,omitempty"`
	// The requested path specified in the forwarding rule.
	//
	// example:
	//
	// /cache
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// The ID of the destination vServer group specified in the forwarding rule.
	//
	// example:
	//
	// rsp-6cejjzl****
	VServerGroupId *string `json:"VServerGroupId,omitempty" xml:"VServerGroupId,omitempty"`
}

func (s DescribeRulesResponseBodyRulesRule) String() string {
	return dara.Prettify(s)
}

func (s DescribeRulesResponseBodyRulesRule) GoString() string {
	return s.String()
}

func (s *DescribeRulesResponseBodyRulesRule) GetCookie() *string {
	return s.Cookie
}

func (s *DescribeRulesResponseBodyRulesRule) GetCookieTimeout() *int32 {
	return s.CookieTimeout
}

func (s *DescribeRulesResponseBodyRulesRule) GetDomain() *string {
	return s.Domain
}

func (s *DescribeRulesResponseBodyRulesRule) GetHealthCheck() *string {
	return s.HealthCheck
}

func (s *DescribeRulesResponseBodyRulesRule) GetHealthCheckConnectPort() *int32 {
	return s.HealthCheckConnectPort
}

func (s *DescribeRulesResponseBodyRulesRule) GetHealthCheckDomain() *string {
	return s.HealthCheckDomain
}

func (s *DescribeRulesResponseBodyRulesRule) GetHealthCheckHttpCode() *string {
	return s.HealthCheckHttpCode
}

func (s *DescribeRulesResponseBodyRulesRule) GetHealthCheckInterval() *int32 {
	return s.HealthCheckInterval
}

func (s *DescribeRulesResponseBodyRulesRule) GetHealthCheckTimeout() *int32 {
	return s.HealthCheckTimeout
}

func (s *DescribeRulesResponseBodyRulesRule) GetHealthCheckURI() *string {
	return s.HealthCheckURI
}

func (s *DescribeRulesResponseBodyRulesRule) GetHealthyThreshold() *int32 {
	return s.HealthyThreshold
}

func (s *DescribeRulesResponseBodyRulesRule) GetListenerSync() *string {
	return s.ListenerSync
}

func (s *DescribeRulesResponseBodyRulesRule) GetRuleId() *string {
	return s.RuleId
}

func (s *DescribeRulesResponseBodyRulesRule) GetRuleName() *string {
	return s.RuleName
}

func (s *DescribeRulesResponseBodyRulesRule) GetScheduler() *string {
	return s.Scheduler
}

func (s *DescribeRulesResponseBodyRulesRule) GetStickySession() *string {
	return s.StickySession
}

func (s *DescribeRulesResponseBodyRulesRule) GetStickySessionType() *string {
	return s.StickySessionType
}

func (s *DescribeRulesResponseBodyRulesRule) GetUnhealthyThreshold() *int32 {
	return s.UnhealthyThreshold
}

func (s *DescribeRulesResponseBodyRulesRule) GetUrl() *string {
	return s.Url
}

func (s *DescribeRulesResponseBodyRulesRule) GetVServerGroupId() *string {
	return s.VServerGroupId
}

func (s *DescribeRulesResponseBodyRulesRule) SetCookie(v string) *DescribeRulesResponseBodyRulesRule {
	s.Cookie = &v
	return s
}

func (s *DescribeRulesResponseBodyRulesRule) SetCookieTimeout(v int32) *DescribeRulesResponseBodyRulesRule {
	s.CookieTimeout = &v
	return s
}

func (s *DescribeRulesResponseBodyRulesRule) SetDomain(v string) *DescribeRulesResponseBodyRulesRule {
	s.Domain = &v
	return s
}

func (s *DescribeRulesResponseBodyRulesRule) SetHealthCheck(v string) *DescribeRulesResponseBodyRulesRule {
	s.HealthCheck = &v
	return s
}

func (s *DescribeRulesResponseBodyRulesRule) SetHealthCheckConnectPort(v int32) *DescribeRulesResponseBodyRulesRule {
	s.HealthCheckConnectPort = &v
	return s
}

func (s *DescribeRulesResponseBodyRulesRule) SetHealthCheckDomain(v string) *DescribeRulesResponseBodyRulesRule {
	s.HealthCheckDomain = &v
	return s
}

func (s *DescribeRulesResponseBodyRulesRule) SetHealthCheckHttpCode(v string) *DescribeRulesResponseBodyRulesRule {
	s.HealthCheckHttpCode = &v
	return s
}

func (s *DescribeRulesResponseBodyRulesRule) SetHealthCheckInterval(v int32) *DescribeRulesResponseBodyRulesRule {
	s.HealthCheckInterval = &v
	return s
}

func (s *DescribeRulesResponseBodyRulesRule) SetHealthCheckTimeout(v int32) *DescribeRulesResponseBodyRulesRule {
	s.HealthCheckTimeout = &v
	return s
}

func (s *DescribeRulesResponseBodyRulesRule) SetHealthCheckURI(v string) *DescribeRulesResponseBodyRulesRule {
	s.HealthCheckURI = &v
	return s
}

func (s *DescribeRulesResponseBodyRulesRule) SetHealthyThreshold(v int32) *DescribeRulesResponseBodyRulesRule {
	s.HealthyThreshold = &v
	return s
}

func (s *DescribeRulesResponseBodyRulesRule) SetListenerSync(v string) *DescribeRulesResponseBodyRulesRule {
	s.ListenerSync = &v
	return s
}

func (s *DescribeRulesResponseBodyRulesRule) SetRuleId(v string) *DescribeRulesResponseBodyRulesRule {
	s.RuleId = &v
	return s
}

func (s *DescribeRulesResponseBodyRulesRule) SetRuleName(v string) *DescribeRulesResponseBodyRulesRule {
	s.RuleName = &v
	return s
}

func (s *DescribeRulesResponseBodyRulesRule) SetScheduler(v string) *DescribeRulesResponseBodyRulesRule {
	s.Scheduler = &v
	return s
}

func (s *DescribeRulesResponseBodyRulesRule) SetStickySession(v string) *DescribeRulesResponseBodyRulesRule {
	s.StickySession = &v
	return s
}

func (s *DescribeRulesResponseBodyRulesRule) SetStickySessionType(v string) *DescribeRulesResponseBodyRulesRule {
	s.StickySessionType = &v
	return s
}

func (s *DescribeRulesResponseBodyRulesRule) SetUnhealthyThreshold(v int32) *DescribeRulesResponseBodyRulesRule {
	s.UnhealthyThreshold = &v
	return s
}

func (s *DescribeRulesResponseBodyRulesRule) SetUrl(v string) *DescribeRulesResponseBodyRulesRule {
	s.Url = &v
	return s
}

func (s *DescribeRulesResponseBodyRulesRule) SetVServerGroupId(v string) *DescribeRulesResponseBodyRulesRule {
	s.VServerGroupId = &v
	return s
}

func (s *DescribeRulesResponseBodyRulesRule) Validate() error {
	return dara.Validate(s)
}
