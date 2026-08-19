// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTraceSiteResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *TraceSiteResponseBody
	GetRequestId() *string
	SetStatusCode(v int32) *TraceSiteResponseBody
	GetStatusCode() *int32
	SetSuccess(v bool) *TraceSiteResponseBody
	GetSuccess() *bool
	SetTrace(v []*TraceSiteResponseBodyTrace) *TraceSiteResponseBody
	GetTrace() []*TraceSiteResponseBodyTrace
}

type TraceSiteResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 9574AFDC-ABF1-5068-AAE3-6958CEBD8740
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The HTTP status code of the request.
	//
	// example:
	//
	// 200
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The trace information of the call chain.
	Trace []*TraceSiteResponseBodyTrace `json:"Trace,omitempty" xml:"Trace,omitempty" type:"Repeated"`
}

func (s TraceSiteResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TraceSiteResponseBody) GoString() string {
	return s.String()
}

func (s *TraceSiteResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TraceSiteResponseBody) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TraceSiteResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *TraceSiteResponseBody) GetTrace() []*TraceSiteResponseBodyTrace {
	return s.Trace
}

func (s *TraceSiteResponseBody) SetRequestId(v string) *TraceSiteResponseBody {
	s.RequestId = &v
	return s
}

func (s *TraceSiteResponseBody) SetStatusCode(v int32) *TraceSiteResponseBody {
	s.StatusCode = &v
	return s
}

func (s *TraceSiteResponseBody) SetSuccess(v bool) *TraceSiteResponseBody {
	s.Success = &v
	return s
}

func (s *TraceSiteResponseBody) SetTrace(v []*TraceSiteResponseBodyTrace) *TraceSiteResponseBody {
	s.Trace = v
	return s
}

func (s *TraceSiteResponseBody) Validate() error {
	if s.Trace != nil {
		for _, item := range s.Trace {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type TraceSiteResponseBodyTrace struct {
	// Indicates whether the module is matched. Valid values: true and false.
	//
	// example:
	//
	// true
	Matched *bool `json:"Matched,omitempty" xml:"Matched,omitempty"`
	// The feature module.
	//
	// example:
	//
	// SecRules
	StepModuleName *string `json:"StepModuleName,omitempty" xml:"StepModuleName,omitempty"`
	// The matching results of rules in the feature module.
	Trace []*TraceSiteResponseBodyTraceTrace `json:"Trace,omitempty" xml:"Trace,omitempty" type:"Repeated"`
}

func (s TraceSiteResponseBodyTrace) String() string {
	return dara.Prettify(s)
}

func (s TraceSiteResponseBodyTrace) GoString() string {
	return s.String()
}

func (s *TraceSiteResponseBodyTrace) GetMatched() *bool {
	return s.Matched
}

func (s *TraceSiteResponseBodyTrace) GetStepModuleName() *string {
	return s.StepModuleName
}

func (s *TraceSiteResponseBodyTrace) GetTrace() []*TraceSiteResponseBodyTraceTrace {
	return s.Trace
}

func (s *TraceSiteResponseBodyTrace) SetMatched(v bool) *TraceSiteResponseBodyTrace {
	s.Matched = &v
	return s
}

func (s *TraceSiteResponseBodyTrace) SetStepModuleName(v string) *TraceSiteResponseBodyTrace {
	s.StepModuleName = &v
	return s
}

func (s *TraceSiteResponseBodyTrace) SetTrace(v []*TraceSiteResponseBodyTraceTrace) *TraceSiteResponseBodyTrace {
	s.Trace = v
	return s
}

func (s *TraceSiteResponseBodyTrace) Validate() error {
	if s.Trace != nil {
		for _, item := range s.Trace {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type TraceSiteResponseBodyTraceTrace struct {
	// The action to perform.
	//
	// example:
	//
	// deny
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// The configuration type.
	//
	// example:
	//
	// rule
	ConfigType *string `json:"ConfigType,omitempty" xml:"ConfigType,omitempty"`
	// The mitigation capability (China).
	//
	// example:
	//
	// cn300
	DdosLevelDomestic *string `json:"DdosLevelDomestic,omitempty" xml:"DdosLevelDomestic,omitempty"`
	// The mitigation capability (global, excluding China).
	//
	// example:
	//
	// unlimit
	DdosLevelOversea *string `json:"DdosLevelOversea,omitempty" xml:"DdosLevelOversea,omitempty"`
	// The environment.
	//
	// example:
	//
	// Production
	EnvName *string `json:"EnvName,omitempty" xml:"EnvName,omitempty"`
	// The rule expression.
	//
	// example:
	//
	// (ip.geoip.country eq \\"CN\\")
	Expression *string `json:"Expression,omitempty" xml:"Expression,omitempty"`
	// The mitigation capability.
	//
	// example:
	//
	// week
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// The load balancer domain name.
	//
	// example:
	//
	// Ib.test.example.com
	LoadBalancerName *string `json:"LoadBalancerName,omitempty" xml:"LoadBalancerName,omitempty"`
	// The name of the origin pool.
	//
	// example:
	//
	// 21212.origin-pool.example.com
	OriginPoolName *string `json:"OriginPoolName,omitempty" xml:"OriginPoolName,omitempty"`
	// The routine ID.
	//
	// example:
	//
	// test.1097011697834102
	RoutineId *string `json:"RoutineId,omitempty" xml:"RoutineId,omitempty"`
	// The security rule ID.
	//
	// example:
	//
	// 1297141
	RuleId *int32 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The name of the matched rule.
	//
	// example:
	//
	// cache_test
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The version.
	//
	// example:
	//
	// 0
	SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
	// The security-related rule type.
	//
	// example:
	//
	// l4_ddos
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The value specified in the IP access rule.
	//
	// example:
	//
	// 1.1.1.1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s TraceSiteResponseBodyTraceTrace) String() string {
	return dara.Prettify(s)
}

func (s TraceSiteResponseBodyTraceTrace) GoString() string {
	return s.String()
}

func (s *TraceSiteResponseBodyTraceTrace) GetAction() *string {
	return s.Action
}

func (s *TraceSiteResponseBodyTraceTrace) GetConfigType() *string {
	return s.ConfigType
}

func (s *TraceSiteResponseBodyTraceTrace) GetDdosLevelDomestic() *string {
	return s.DdosLevelDomestic
}

func (s *TraceSiteResponseBodyTraceTrace) GetDdosLevelOversea() *string {
	return s.DdosLevelOversea
}

func (s *TraceSiteResponseBodyTraceTrace) GetEnvName() *string {
	return s.EnvName
}

func (s *TraceSiteResponseBodyTraceTrace) GetExpression() *string {
	return s.Expression
}

func (s *TraceSiteResponseBodyTraceTrace) GetLevel() *string {
	return s.Level
}

func (s *TraceSiteResponseBodyTraceTrace) GetLoadBalancerName() *string {
	return s.LoadBalancerName
}

func (s *TraceSiteResponseBodyTraceTrace) GetOriginPoolName() *string {
	return s.OriginPoolName
}

func (s *TraceSiteResponseBodyTraceTrace) GetRoutineId() *string {
	return s.RoutineId
}

func (s *TraceSiteResponseBodyTraceTrace) GetRuleId() *int32 {
	return s.RuleId
}

func (s *TraceSiteResponseBodyTraceTrace) GetRuleName() *string {
	return s.RuleName
}

func (s *TraceSiteResponseBodyTraceTrace) GetSiteVersion() *int32 {
	return s.SiteVersion
}

func (s *TraceSiteResponseBodyTraceTrace) GetType() *string {
	return s.Type
}

func (s *TraceSiteResponseBodyTraceTrace) GetValue() *string {
	return s.Value
}

func (s *TraceSiteResponseBodyTraceTrace) SetAction(v string) *TraceSiteResponseBodyTraceTrace {
	s.Action = &v
	return s
}

func (s *TraceSiteResponseBodyTraceTrace) SetConfigType(v string) *TraceSiteResponseBodyTraceTrace {
	s.ConfigType = &v
	return s
}

func (s *TraceSiteResponseBodyTraceTrace) SetDdosLevelDomestic(v string) *TraceSiteResponseBodyTraceTrace {
	s.DdosLevelDomestic = &v
	return s
}

func (s *TraceSiteResponseBodyTraceTrace) SetDdosLevelOversea(v string) *TraceSiteResponseBodyTraceTrace {
	s.DdosLevelOversea = &v
	return s
}

func (s *TraceSiteResponseBodyTraceTrace) SetEnvName(v string) *TraceSiteResponseBodyTraceTrace {
	s.EnvName = &v
	return s
}

func (s *TraceSiteResponseBodyTraceTrace) SetExpression(v string) *TraceSiteResponseBodyTraceTrace {
	s.Expression = &v
	return s
}

func (s *TraceSiteResponseBodyTraceTrace) SetLevel(v string) *TraceSiteResponseBodyTraceTrace {
	s.Level = &v
	return s
}

func (s *TraceSiteResponseBodyTraceTrace) SetLoadBalancerName(v string) *TraceSiteResponseBodyTraceTrace {
	s.LoadBalancerName = &v
	return s
}

func (s *TraceSiteResponseBodyTraceTrace) SetOriginPoolName(v string) *TraceSiteResponseBodyTraceTrace {
	s.OriginPoolName = &v
	return s
}

func (s *TraceSiteResponseBodyTraceTrace) SetRoutineId(v string) *TraceSiteResponseBodyTraceTrace {
	s.RoutineId = &v
	return s
}

func (s *TraceSiteResponseBodyTraceTrace) SetRuleId(v int32) *TraceSiteResponseBodyTraceTrace {
	s.RuleId = &v
	return s
}

func (s *TraceSiteResponseBodyTraceTrace) SetRuleName(v string) *TraceSiteResponseBodyTraceTrace {
	s.RuleName = &v
	return s
}

func (s *TraceSiteResponseBodyTraceTrace) SetSiteVersion(v int32) *TraceSiteResponseBodyTraceTrace {
	s.SiteVersion = &v
	return s
}

func (s *TraceSiteResponseBodyTraceTrace) SetType(v string) *TraceSiteResponseBodyTraceTrace {
	s.Type = &v
	return s
}

func (s *TraceSiteResponseBodyTraceTrace) SetValue(v string) *TraceSiteResponseBodyTraceTrace {
	s.Value = &v
	return s
}

func (s *TraceSiteResponseBodyTraceTrace) Validate() error {
	return dara.Validate(s)
}
