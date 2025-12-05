// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListRulesResponseBody
	GetRequestId() *string
	SetRules(v []*ListRulesResponseBodyRules) *ListRulesResponseBody
	GetRules() []*ListRulesResponseBodyRules
	SetTotalCount(v int64) *ListRulesResponseBody
	GetTotalCount() *int64
}

type ListRulesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// EC9BF0F4-8983-491A-BC8C-1B4DD94976DE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The authorization rules that are returned.
	Rules []*ListRulesResponseBodyRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	// The total number of authorization rules that are returned.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRulesResponseBody) GetRules() []*ListRulesResponseBodyRules {
	return s.Rules
}

func (s *ListRulesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListRulesResponseBody) SetRequestId(v string) *ListRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRulesResponseBody) SetRules(v []*ListRulesResponseBodyRules) *ListRulesResponseBody {
	s.Rules = v
	return s
}

func (s *ListRulesResponseBody) SetTotalCount(v int64) *ListRulesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListRulesResponseBody) Validate() error {
	if s.Rules != nil {
		for _, item := range s.Rules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListRulesResponseBodyRules struct {
	// The remarks of the authorization rule.
	//
	// example:
	//
	// comment
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The end time of the validity period of the authorization rule. The value is a timestamp. Unit: seconds.
	//
	// example:
	//
	// 1709258400
	EffectiveEndTime *int64 `json:"EffectiveEndTime,omitempty" xml:"EffectiveEndTime,omitempty"`
	// The start time of the validity period of the authorization rule. The value is a timestamp. Unit: seconds.
	//
	// example:
	//
	// 1685499134
	EffectiveStartTime *int64 `json:"EffectiveStartTime,omitempty" xml:"EffectiveStartTime,omitempty"`
	// The authorization rule ID.
	//
	// example:
	//
	// 13
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The name of the authorization rule.
	//
	// example:
	//
	// rule
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The state of the authorization rule.
	//
	// 	- **Enabled**
	//
	// 	- **Disabled**
	//
	// example:
	//
	// Enabled
	RuleState *string `json:"RuleState,omitempty" xml:"RuleState,omitempty"`
}

func (s ListRulesResponseBodyRules) String() string {
	return dara.Prettify(s)
}

func (s ListRulesResponseBodyRules) GoString() string {
	return s.String()
}

func (s *ListRulesResponseBodyRules) GetComment() *string {
	return s.Comment
}

func (s *ListRulesResponseBodyRules) GetEffectiveEndTime() *int64 {
	return s.EffectiveEndTime
}

func (s *ListRulesResponseBodyRules) GetEffectiveStartTime() *int64 {
	return s.EffectiveStartTime
}

func (s *ListRulesResponseBodyRules) GetRuleId() *string {
	return s.RuleId
}

func (s *ListRulesResponseBodyRules) GetRuleName() *string {
	return s.RuleName
}

func (s *ListRulesResponseBodyRules) GetRuleState() *string {
	return s.RuleState
}

func (s *ListRulesResponseBodyRules) SetComment(v string) *ListRulesResponseBodyRules {
	s.Comment = &v
	return s
}

func (s *ListRulesResponseBodyRules) SetEffectiveEndTime(v int64) *ListRulesResponseBodyRules {
	s.EffectiveEndTime = &v
	return s
}

func (s *ListRulesResponseBodyRules) SetEffectiveStartTime(v int64) *ListRulesResponseBodyRules {
	s.EffectiveStartTime = &v
	return s
}

func (s *ListRulesResponseBodyRules) SetRuleId(v string) *ListRulesResponseBodyRules {
	s.RuleId = &v
	return s
}

func (s *ListRulesResponseBodyRules) SetRuleName(v string) *ListRulesResponseBodyRules {
	s.RuleName = &v
	return s
}

func (s *ListRulesResponseBodyRules) SetRuleState(v string) *ListRulesResponseBodyRules {
	s.RuleState = &v
	return s
}

func (s *ListRulesResponseBodyRules) Validate() error {
	return dara.Validate(s)
}
