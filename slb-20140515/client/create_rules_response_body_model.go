// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateRulesResponseBody
	GetRequestId() *string
	SetRules(v *CreateRulesResponseBodyRules) *CreateRulesResponseBody
	GetRules() *CreateRulesResponseBodyRules
}

type CreateRulesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 9DEC9C28-AB05-4DDF-9A78-6B08EC9CE18C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The forwarding rules.
	Rules *CreateRulesResponseBodyRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Struct"`
}

func (s CreateRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateRulesResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateRulesResponseBody) GetRules() *CreateRulesResponseBodyRules {
	return s.Rules
}

func (s *CreateRulesResponseBody) SetRequestId(v string) *CreateRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRulesResponseBody) SetRules(v *CreateRulesResponseBodyRules) *CreateRulesResponseBody {
	s.Rules = v
	return s
}

func (s *CreateRulesResponseBody) Validate() error {
	if s.Rules != nil {
		if err := s.Rules.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateRulesResponseBodyRules struct {
	Rule []*CreateRulesResponseBodyRulesRule `json:"Rule,omitempty" xml:"Rule,omitempty" type:"Repeated"`
}

func (s CreateRulesResponseBodyRules) String() string {
	return dara.Prettify(s)
}

func (s CreateRulesResponseBodyRules) GoString() string {
	return s.String()
}

func (s *CreateRulesResponseBodyRules) GetRule() []*CreateRulesResponseBodyRulesRule {
	return s.Rule
}

func (s *CreateRulesResponseBodyRules) SetRule(v []*CreateRulesResponseBodyRulesRule) *CreateRulesResponseBodyRules {
	s.Rule = v
	return s
}

func (s *CreateRulesResponseBodyRules) Validate() error {
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

type CreateRulesResponseBodyRulesRule struct {
	// The forwarding rule ID.
	//
	// example:
	//
	// rule-bp12jzy0*****
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The name of the forwarding rule.
	//
	// example:
	//
	// Rule2
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s CreateRulesResponseBodyRulesRule) String() string {
	return dara.Prettify(s)
}

func (s CreateRulesResponseBodyRulesRule) GoString() string {
	return s.String()
}

func (s *CreateRulesResponseBodyRulesRule) GetRuleId() *string {
	return s.RuleId
}

func (s *CreateRulesResponseBodyRulesRule) GetRuleName() *string {
	return s.RuleName
}

func (s *CreateRulesResponseBodyRulesRule) SetRuleId(v string) *CreateRulesResponseBodyRulesRule {
	s.RuleId = &v
	return s
}

func (s *CreateRulesResponseBodyRulesRule) SetRuleName(v string) *CreateRulesResponseBodyRulesRule {
	s.RuleName = &v
	return s
}

func (s *CreateRulesResponseBodyRulesRule) Validate() error {
	return dara.Validate(s)
}
