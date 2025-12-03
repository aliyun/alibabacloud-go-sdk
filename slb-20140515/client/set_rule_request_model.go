// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCookie(v string) *SetRuleRequest
	GetCookie() *string
	SetCookieTimeout(v int32) *SetRuleRequest
	GetCookieTimeout() *int32
	SetHealthCheck(v string) *SetRuleRequest
	GetHealthCheck() *string
	SetHealthCheckConnectPort(v int32) *SetRuleRequest
	GetHealthCheckConnectPort() *int32
	SetHealthCheckDomain(v string) *SetRuleRequest
	GetHealthCheckDomain() *string
	SetHealthCheckHttpCode(v string) *SetRuleRequest
	GetHealthCheckHttpCode() *string
	SetHealthCheckInterval(v int32) *SetRuleRequest
	GetHealthCheckInterval() *int32
	SetHealthCheckTimeout(v int32) *SetRuleRequest
	GetHealthCheckTimeout() *int32
	SetHealthCheckURI(v string) *SetRuleRequest
	GetHealthCheckURI() *string
	SetHealthyThreshold(v int32) *SetRuleRequest
	GetHealthyThreshold() *int32
	SetListenerSync(v string) *SetRuleRequest
	GetListenerSync() *string
	SetOwnerAccount(v string) *SetRuleRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *SetRuleRequest
	GetOwnerId() *int64
	SetRegionId(v string) *SetRuleRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *SetRuleRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *SetRuleRequest
	GetResourceOwnerId() *int64
	SetRuleId(v string) *SetRuleRequest
	GetRuleId() *string
	SetRuleName(v string) *SetRuleRequest
	GetRuleName() *string
	SetScheduler(v string) *SetRuleRequest
	GetScheduler() *string
	SetStickySession(v string) *SetRuleRequest
	GetStickySession() *string
	SetStickySessionType(v string) *SetRuleRequest
	GetStickySessionType() *string
	SetUnhealthyThreshold(v int32) *SetRuleRequest
	GetUnhealthyThreshold() *int32
	SetVServerGroupId(v string) *SetRuleRequest
	GetVServerGroupId() *string
}

type SetRuleRequest struct {
	// The cookie that is configured on the server.
	//
	// The cookie must be 1 to 200 characters in length and can contain only ASCII characters and digits. It cannot contain commas (,), semicolons (;), or space characters. It cannot start with a dollar sign ($).
	//
	// >  This parameter is required and takes effect if **StickySession*	- is set to **on*	- and **StickySessionType*	- is set to **server**.
	//
	// example:
	//
	// 23ffsa
	Cookie *string `json:"Cookie,omitempty" xml:"Cookie,omitempty"`
	// The timeout period of a cookie. Unit: seconds. Valid values: **1*	- to **86400**.
	//
	// >  This parameter is required and takes effect if **StickySession*	- is set to **on*	- and **StickySessionType*	- is set to **insert**.
	//
	// example:
	//
	// 123
	CookieTimeout *int32 `json:"CookieTimeout,omitempty" xml:"CookieTimeout,omitempty"`
	// Specifies whether to enable the health check feature. Valid values:
	//
	// 	- **on**: yes
	//
	// 	- **off**: no
	//
	// >  This parameter is required and takes effect if the **ListenerSync*	- parameter is set to **off**.
	//
	// example:
	//
	// off
	HealthCheck *string `json:"HealthCheck,omitempty" xml:"HealthCheck,omitempty"`
	// The port that is used for health checks. Valid values: **1*	- to **65535**.
	//
	// >  This parameter takes effect when the **HealthCheck*	- parameter is set to **on**.
	//
	// example:
	//
	// 80
	HealthCheckConnectPort *int32 `json:"HealthCheckConnectPort,omitempty" xml:"HealthCheckConnectPort,omitempty"`
	// The domain name that is used for health checks. Valid values:
	//
	// 	- **$_ip**: the private IP address of a backend server. If you do not set this parameter or set the parameter to $_ip, the SLB instance uses the private IP address of each backend server for health checks.
	//
	// 	- **domain**: The domain name must be 1 to 80 characters in length, and can contain letters, digits, periods (.), and hyphens (-).
	//
	// >  This parameter takes effect if the **HealthCheck*	- parameter is set to **on**.
	//
	// example:
	//
	// $_ip
	HealthCheckDomain *string `json:"HealthCheckDomain,omitempty" xml:"HealthCheckDomain,omitempty"`
	// The HTTP status code for a successful health check. Multiple HTTP status codes are separated by commas (,).
	//
	// Valid values: **http_2xx**, **http_3xx**, **http_4xx**, and **http_5xx**.
	//
	// >  This parameter is required and takes effect if the **HealthCheck*	- parameter is set to **on**.
	//
	// example:
	//
	// http_2xx
	HealthCheckHttpCode *string `json:"HealthCheckHttpCode,omitempty" xml:"HealthCheckHttpCode,omitempty"`
	// The interval between two consecutive health checks. Unit: seconds. Valid values: **1*	- to **50**.
	//
	// >  This parameter is required and takes effect if the **HealthCheck*	- parameter is set to **on**.
	//
	// example:
	//
	// 20
	HealthCheckInterval *int32 `json:"HealthCheckInterval,omitempty" xml:"HealthCheckInterval,omitempty"`
	// The timeout period of a health check response. If a backend server, such as an Elastic Compute Service (ECS) instance, does not return a health check response within the specified timeout period, the server fails the health check. Unit: seconds. Valid values: **1*	- to **300**.
	//
	// >  This parameter is required and takes effect if the **HealthCheck*	- parameter is set to **on**.
	//
	// example:
	//
	// 20
	HealthCheckTimeout *int32 `json:"HealthCheckTimeout,omitempty" xml:"HealthCheckTimeout,omitempty"`
	// The URI that is used for health checks.
	//
	// >  This parameter is required and takes effect if the **HealthCheck*	- parameter is set to **on**.
	//
	// example:
	//
	// /example
	HealthCheckURI *string `json:"HealthCheckURI,omitempty" xml:"HealthCheckURI,omitempty"`
	// The number of times that an unhealthy backend server must consecutively pass health checks before it is declared healthy. In this case, the health status is changed from **fail*	- to **success**.
	//
	// Valid values: **2*	- to **10**.
	//
	// >  This parameter is required and takes effect if the **HealthCheck*	- parameter is set to **on**.
	//
	// example:
	//
	// 4
	HealthyThreshold *int32 `json:"HealthyThreshold,omitempty" xml:"HealthyThreshold,omitempty"`
	// Specifies whether to use the scheduling algorithm, session persistence, and health check configurations of the listener. Valid values:
	//
	// 	- **on**: uses the configurations of the listener.
	//
	// 	- **off**: does not use the configurations of the listener. You can customize the health check and session persistence configurations for the forwarding rule.
	//
	// example:
	//
	// off
	ListenerSync *string `json:"ListenerSync,omitempty" xml:"ListenerSync,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the Classic Load Balancer (CLB) instance is deployed.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/27584.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the forwarding rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// rule-3ejhkt****
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The name of the forwarding rule. The name must be 1 to 40 characters in length, and can contain letters, digits, hyphens (-), forward slashes (/), periods (.), and underscores (_).
	//
	// > On the same listener, the forwarding rule names must be unique.
	//
	// example:
	//
	// doctest
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The scheduling algorithm. Valid values:
	//
	// 	- **wrr**: Backend servers with higher weights receive more requests than those with lower weights.
	//
	// 	- **rr**: Requests are distributed to backend servers in sequence.
	//
	// >  This parameter is required and takes effect if the **ListenerSync*	- parameter is set to **off**.
	//
	// example:
	//
	// wrr
	Scheduler *string `json:"Scheduler,omitempty" xml:"Scheduler,omitempty"`
	// Specifies whether to enable session persistence. Valid values:
	//
	// 	- **on**: yes
	//
	// 	- **off**: no
	//
	// This parameter is required and takes effect if the **ListenerSync*	- parameter is set to **off**.
	//
	// example:
	//
	// off
	StickySession *string `json:"StickySession,omitempty" xml:"StickySession,omitempty"`
	// The method that is used to handle a cookie. Valid values:
	//
	// 	- **insert**: inserts a cookie.
	//
	//     CLB inserts the backend server ID as a cookie into the first HTTP or HTTPS response that is sent to a client. The next request from the client will contain this cookie, and the listener will distribute this request to the recorded backend server.
	//
	// 	- **server**: rewrites a cookie.
	//
	//     When CLB detects a user-defined cookie, it overwrites the original cookie with the user-defined cookie. The next request from the client will contain the user-defined cookie, and the listener will distribute this request to the recorded backend server.
	//
	// >  This parameter is required and takes effect if the **StickySession*	- parameter is set to **on**.
	//
	// example:
	//
	// insert
	StickySessionType *string `json:"StickySessionType,omitempty" xml:"StickySessionType,omitempty"`
	// The number of times that a healthy backend server must consecutively fail health checks before it is declared unhealthy. In this case, the health status is changed from **success*	- to **fail**.
	//
	// Valid values: **2*	- to **10**.
	//
	// >  This parameter is required and takes effect if the **HealthCheck*	- parameter is set to **on**.
	//
	// example:
	//
	// 4
	UnhealthyThreshold *int32 `json:"UnhealthyThreshold,omitempty" xml:"UnhealthyThreshold,omitempty"`
	// The ID of the vServer group that is associated with the forwarding rule.
	//
	// example:
	//
	// rsp-cige6****
	VServerGroupId *string `json:"VServerGroupId,omitempty" xml:"VServerGroupId,omitempty"`
}

func (s SetRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s SetRuleRequest) GoString() string {
	return s.String()
}

func (s *SetRuleRequest) GetCookie() *string {
	return s.Cookie
}

func (s *SetRuleRequest) GetCookieTimeout() *int32 {
	return s.CookieTimeout
}

func (s *SetRuleRequest) GetHealthCheck() *string {
	return s.HealthCheck
}

func (s *SetRuleRequest) GetHealthCheckConnectPort() *int32 {
	return s.HealthCheckConnectPort
}

func (s *SetRuleRequest) GetHealthCheckDomain() *string {
	return s.HealthCheckDomain
}

func (s *SetRuleRequest) GetHealthCheckHttpCode() *string {
	return s.HealthCheckHttpCode
}

func (s *SetRuleRequest) GetHealthCheckInterval() *int32 {
	return s.HealthCheckInterval
}

func (s *SetRuleRequest) GetHealthCheckTimeout() *int32 {
	return s.HealthCheckTimeout
}

func (s *SetRuleRequest) GetHealthCheckURI() *string {
	return s.HealthCheckURI
}

func (s *SetRuleRequest) GetHealthyThreshold() *int32 {
	return s.HealthyThreshold
}

func (s *SetRuleRequest) GetListenerSync() *string {
	return s.ListenerSync
}

func (s *SetRuleRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *SetRuleRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SetRuleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SetRuleRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SetRuleRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *SetRuleRequest) GetRuleId() *string {
	return s.RuleId
}

func (s *SetRuleRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *SetRuleRequest) GetScheduler() *string {
	return s.Scheduler
}

func (s *SetRuleRequest) GetStickySession() *string {
	return s.StickySession
}

func (s *SetRuleRequest) GetStickySessionType() *string {
	return s.StickySessionType
}

func (s *SetRuleRequest) GetUnhealthyThreshold() *int32 {
	return s.UnhealthyThreshold
}

func (s *SetRuleRequest) GetVServerGroupId() *string {
	return s.VServerGroupId
}

func (s *SetRuleRequest) SetCookie(v string) *SetRuleRequest {
	s.Cookie = &v
	return s
}

func (s *SetRuleRequest) SetCookieTimeout(v int32) *SetRuleRequest {
	s.CookieTimeout = &v
	return s
}

func (s *SetRuleRequest) SetHealthCheck(v string) *SetRuleRequest {
	s.HealthCheck = &v
	return s
}

func (s *SetRuleRequest) SetHealthCheckConnectPort(v int32) *SetRuleRequest {
	s.HealthCheckConnectPort = &v
	return s
}

func (s *SetRuleRequest) SetHealthCheckDomain(v string) *SetRuleRequest {
	s.HealthCheckDomain = &v
	return s
}

func (s *SetRuleRequest) SetHealthCheckHttpCode(v string) *SetRuleRequest {
	s.HealthCheckHttpCode = &v
	return s
}

func (s *SetRuleRequest) SetHealthCheckInterval(v int32) *SetRuleRequest {
	s.HealthCheckInterval = &v
	return s
}

func (s *SetRuleRequest) SetHealthCheckTimeout(v int32) *SetRuleRequest {
	s.HealthCheckTimeout = &v
	return s
}

func (s *SetRuleRequest) SetHealthCheckURI(v string) *SetRuleRequest {
	s.HealthCheckURI = &v
	return s
}

func (s *SetRuleRequest) SetHealthyThreshold(v int32) *SetRuleRequest {
	s.HealthyThreshold = &v
	return s
}

func (s *SetRuleRequest) SetListenerSync(v string) *SetRuleRequest {
	s.ListenerSync = &v
	return s
}

func (s *SetRuleRequest) SetOwnerAccount(v string) *SetRuleRequest {
	s.OwnerAccount = &v
	return s
}

func (s *SetRuleRequest) SetOwnerId(v int64) *SetRuleRequest {
	s.OwnerId = &v
	return s
}

func (s *SetRuleRequest) SetRegionId(v string) *SetRuleRequest {
	s.RegionId = &v
	return s
}

func (s *SetRuleRequest) SetResourceOwnerAccount(v string) *SetRuleRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SetRuleRequest) SetResourceOwnerId(v int64) *SetRuleRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SetRuleRequest) SetRuleId(v string) *SetRuleRequest {
	s.RuleId = &v
	return s
}

func (s *SetRuleRequest) SetRuleName(v string) *SetRuleRequest {
	s.RuleName = &v
	return s
}

func (s *SetRuleRequest) SetScheduler(v string) *SetRuleRequest {
	s.Scheduler = &v
	return s
}

func (s *SetRuleRequest) SetStickySession(v string) *SetRuleRequest {
	s.StickySession = &v
	return s
}

func (s *SetRuleRequest) SetStickySessionType(v string) *SetRuleRequest {
	s.StickySessionType = &v
	return s
}

func (s *SetRuleRequest) SetUnhealthyThreshold(v int32) *SetRuleRequest {
	s.UnhealthyThreshold = &v
	return s
}

func (s *SetRuleRequest) SetVServerGroupId(v string) *SetRuleRequest {
	s.VServerGroupId = &v
	return s
}

func (s *SetRuleRequest) Validate() error {
	return dara.Validate(s)
}
