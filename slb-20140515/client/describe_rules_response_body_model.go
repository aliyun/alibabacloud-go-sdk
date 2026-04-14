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
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Rules     *DescribeRulesResponseBodyRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Struct"`
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
	Cookie                 *string `json:"Cookie,omitempty" xml:"Cookie,omitempty"`
	CookieTimeout          *int32  `json:"CookieTimeout,omitempty" xml:"CookieTimeout,omitempty"`
	Domain                 *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	HealthCheck            *string `json:"HealthCheck,omitempty" xml:"HealthCheck,omitempty"`
	HealthCheckConnectPort *int32  `json:"HealthCheckConnectPort,omitempty" xml:"HealthCheckConnectPort,omitempty"`
	HealthCheckDomain      *string `json:"HealthCheckDomain,omitempty" xml:"HealthCheckDomain,omitempty"`
	HealthCheckHttpCode    *string `json:"HealthCheckHttpCode,omitempty" xml:"HealthCheckHttpCode,omitempty"`
	HealthCheckInterval    *int32  `json:"HealthCheckInterval,omitempty" xml:"HealthCheckInterval,omitempty"`
	HealthCheckTimeout     *int32  `json:"HealthCheckTimeout,omitempty" xml:"HealthCheckTimeout,omitempty"`
	HealthCheckURI         *string `json:"HealthCheckURI,omitempty" xml:"HealthCheckURI,omitempty"`
	HealthyThreshold       *int32  `json:"HealthyThreshold,omitempty" xml:"HealthyThreshold,omitempty"`
	ListenerSync           *string `json:"ListenerSync,omitempty" xml:"ListenerSync,omitempty"`
	RuleId                 *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	RuleName               *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	Scheduler              *string `json:"Scheduler,omitempty" xml:"Scheduler,omitempty"`
	StickySession          *string `json:"StickySession,omitempty" xml:"StickySession,omitempty"`
	StickySessionType      *string `json:"StickySessionType,omitempty" xml:"StickySessionType,omitempty"`
	UnhealthyThreshold     *int32  `json:"UnhealthyThreshold,omitempty" xml:"UnhealthyThreshold,omitempty"`
	Url                    *string `json:"Url,omitempty" xml:"Url,omitempty"`
	VServerGroupId         *string `json:"VServerGroupId,omitempty" xml:"VServerGroupId,omitempty"`
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
