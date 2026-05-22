// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWaitingRoomRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListWaitingRoomRulesResponseBody
	GetRequestId() *string
	SetWaitingRoomRules(v []*ListWaitingRoomRulesResponseBodyWaitingRoomRules) *ListWaitingRoomRulesResponseBody
	GetWaitingRoomRules() []*ListWaitingRoomRulesResponseBodyWaitingRoomRules
}

type ListWaitingRoomRulesResponseBody struct {
	RequestId        *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	WaitingRoomRules []*ListWaitingRoomRulesResponseBodyWaitingRoomRules `json:"WaitingRoomRules,omitempty" xml:"WaitingRoomRules,omitempty" type:"Repeated"`
}

func (s ListWaitingRoomRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListWaitingRoomRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListWaitingRoomRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListWaitingRoomRulesResponseBody) GetWaitingRoomRules() []*ListWaitingRoomRulesResponseBodyWaitingRoomRules {
	return s.WaitingRoomRules
}

func (s *ListWaitingRoomRulesResponseBody) SetRequestId(v string) *ListWaitingRoomRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListWaitingRoomRulesResponseBody) SetWaitingRoomRules(v []*ListWaitingRoomRulesResponseBodyWaitingRoomRules) *ListWaitingRoomRulesResponseBody {
	s.WaitingRoomRules = v
	return s
}

func (s *ListWaitingRoomRulesResponseBody) Validate() error {
	if s.WaitingRoomRules != nil {
		for _, item := range s.WaitingRoomRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListWaitingRoomRulesResponseBodyWaitingRoomRules struct {
	Rule              *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	RuleEnable        *string `json:"RuleEnable,omitempty" xml:"RuleEnable,omitempty"`
	RuleName          *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	WaitingRoomRuleId *int64  `json:"WaitingRoomRuleId,omitempty" xml:"WaitingRoomRuleId,omitempty"`
}

func (s ListWaitingRoomRulesResponseBodyWaitingRoomRules) String() string {
	return dara.Prettify(s)
}

func (s ListWaitingRoomRulesResponseBodyWaitingRoomRules) GoString() string {
	return s.String()
}

func (s *ListWaitingRoomRulesResponseBodyWaitingRoomRules) GetRule() *string {
	return s.Rule
}

func (s *ListWaitingRoomRulesResponseBodyWaitingRoomRules) GetRuleEnable() *string {
	return s.RuleEnable
}

func (s *ListWaitingRoomRulesResponseBodyWaitingRoomRules) GetRuleName() *string {
	return s.RuleName
}

func (s *ListWaitingRoomRulesResponseBodyWaitingRoomRules) GetWaitingRoomRuleId() *int64 {
	return s.WaitingRoomRuleId
}

func (s *ListWaitingRoomRulesResponseBodyWaitingRoomRules) SetRule(v string) *ListWaitingRoomRulesResponseBodyWaitingRoomRules {
	s.Rule = &v
	return s
}

func (s *ListWaitingRoomRulesResponseBodyWaitingRoomRules) SetRuleEnable(v string) *ListWaitingRoomRulesResponseBodyWaitingRoomRules {
	s.RuleEnable = &v
	return s
}

func (s *ListWaitingRoomRulesResponseBodyWaitingRoomRules) SetRuleName(v string) *ListWaitingRoomRulesResponseBodyWaitingRoomRules {
	s.RuleName = &v
	return s
}

func (s *ListWaitingRoomRulesResponseBodyWaitingRoomRules) SetWaitingRoomRuleId(v int64) *ListWaitingRoomRulesResponseBodyWaitingRoomRules {
	s.WaitingRoomRuleId = &v
	return s
}

func (s *ListWaitingRoomRulesResponseBodyWaitingRoomRules) Validate() error {
	return dara.Validate(s)
}
