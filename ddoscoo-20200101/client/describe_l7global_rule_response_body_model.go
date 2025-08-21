// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeL7GlobalRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGlobalRules(v []*DescribeL7GlobalRuleResponseBodyGlobalRules) *DescribeL7GlobalRuleResponseBody
	GetGlobalRules() []*DescribeL7GlobalRuleResponseBodyGlobalRules
	SetRequestId(v string) *DescribeL7GlobalRuleResponseBody
	GetRequestId() *string
}

type DescribeL7GlobalRuleResponseBody struct {
	GlobalRules []*DescribeL7GlobalRuleResponseBodyGlobalRules `json:"GlobalRules,omitempty" xml:"GlobalRules,omitempty" type:"Repeated"`
	// example:
	//
	// CF33B4C3-196E-4015-AADD-5CAD00057B80
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeL7GlobalRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeL7GlobalRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeL7GlobalRuleResponseBody) GetGlobalRules() []*DescribeL7GlobalRuleResponseBodyGlobalRules {
	return s.GlobalRules
}

func (s *DescribeL7GlobalRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeL7GlobalRuleResponseBody) SetGlobalRules(v []*DescribeL7GlobalRuleResponseBodyGlobalRules) *DescribeL7GlobalRuleResponseBody {
	s.GlobalRules = v
	return s
}

func (s *DescribeL7GlobalRuleResponseBody) SetRequestId(v string) *DescribeL7GlobalRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeL7GlobalRuleResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeL7GlobalRuleResponseBodyGlobalRules struct {
	// example:
	//
	// watch
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// example:
	//
	// watch
	ActionDefault *string `json:"ActionDefault,omitempty" xml:"ActionDefault,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 1
	Enabled *int64 `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// example:
	//
	// global_1
	RuleId   *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s DescribeL7GlobalRuleResponseBodyGlobalRules) String() string {
	return dara.Prettify(s)
}

func (s DescribeL7GlobalRuleResponseBodyGlobalRules) GoString() string {
	return s.String()
}

func (s *DescribeL7GlobalRuleResponseBodyGlobalRules) GetAction() *string {
	return s.Action
}

func (s *DescribeL7GlobalRuleResponseBodyGlobalRules) GetActionDefault() *string {
	return s.ActionDefault
}

func (s *DescribeL7GlobalRuleResponseBodyGlobalRules) GetDescription() *string {
	return s.Description
}

func (s *DescribeL7GlobalRuleResponseBodyGlobalRules) GetEnabled() *int64 {
	return s.Enabled
}

func (s *DescribeL7GlobalRuleResponseBodyGlobalRules) GetRuleId() *string {
	return s.RuleId
}

func (s *DescribeL7GlobalRuleResponseBodyGlobalRules) GetRuleName() *string {
	return s.RuleName
}

func (s *DescribeL7GlobalRuleResponseBodyGlobalRules) SetAction(v string) *DescribeL7GlobalRuleResponseBodyGlobalRules {
	s.Action = &v
	return s
}

func (s *DescribeL7GlobalRuleResponseBodyGlobalRules) SetActionDefault(v string) *DescribeL7GlobalRuleResponseBodyGlobalRules {
	s.ActionDefault = &v
	return s
}

func (s *DescribeL7GlobalRuleResponseBodyGlobalRules) SetDescription(v string) *DescribeL7GlobalRuleResponseBodyGlobalRules {
	s.Description = &v
	return s
}

func (s *DescribeL7GlobalRuleResponseBodyGlobalRules) SetEnabled(v int64) *DescribeL7GlobalRuleResponseBodyGlobalRules {
	s.Enabled = &v
	return s
}

func (s *DescribeL7GlobalRuleResponseBodyGlobalRules) SetRuleId(v string) *DescribeL7GlobalRuleResponseBodyGlobalRules {
	s.RuleId = &v
	return s
}

func (s *DescribeL7GlobalRuleResponseBodyGlobalRules) SetRuleName(v string) *DescribeL7GlobalRuleResponseBodyGlobalRules {
	s.RuleName = &v
	return s
}

func (s *DescribeL7GlobalRuleResponseBodyGlobalRules) Validate() error {
	return dara.Validate(s)
}
