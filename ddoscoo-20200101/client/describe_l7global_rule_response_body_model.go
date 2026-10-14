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
	// The global template rule information.
	GlobalRules []*DescribeL7GlobalRuleResponseBodyGlobalRules `json:"GlobalRules,omitempty" xml:"GlobalRules,omitempty" type:"Repeated"`
	// The ID of the request. The ID is a unique identifier that Alibaba Cloud generates for the request. You can use the ID to troubleshoot issues.
	//
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
	if s.GlobalRules != nil {
		for _, item := range s.GlobalRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeL7GlobalRuleResponseBodyGlobalRules struct {
	// The action that the rule performs. Valid values:
	//
	// - **watch**: Monitor.
	//
	// - **block**: Block.
	//
	// - **challenge**: JavaScript Challenge.
	//
	// example:
	//
	// watch
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// The default action of the rule. Valid values:
	//
	// - **watch**: Monitor.
	//
	// - **block**: Block.
	//
	// - **challenge**: JavaScript Challenge.
	//
	// example:
	//
	// watch
	ActionDefault *string `json:"ActionDefault,omitempty" xml:"ActionDefault,omitempty"`
	// The rule description.
	//
	// example:
	//
	// HTTP请求疑似模拟浏览器，校验浏览器Header等组合特征#4
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Indicates whether the rule is enabled. Valid values:
	//
	// - **1**: Enabled.
	//
	// - **0**: Disabled.
	//
	// example:
	//
	// 1
	Enabled *int64 `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The rule ID.
	//
	// example:
	//
	// global_1
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The rule name.
	//
	// example:
	//
	// 模拟浏览器请求
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
