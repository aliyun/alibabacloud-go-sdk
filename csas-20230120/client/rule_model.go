// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRule interface {
	dara.Model
	String() string
	GoString() string
	SetCombinator(v string) *Rule
	GetCombinator() *string
	SetId(v string) *Rule
	GetId() *string
	SetName(v string) *Rule
	GetName() *string
	SetOperator(v string) *Rule
	GetOperator() *string
	SetRuleSubType(v string) *Rule
	GetRuleSubType() *string
	SetRuleType(v string) *Rule
	GetRuleType() *string
	SetRules(v []*Rule) *Rule
	GetRules() []*Rule
	SetValues(v []*string) *Rule
	GetValues() []*string
}

type Rule struct {
	// if can be null:
	// true
	Combinator *string `json:"Combinator,omitempty" xml:"Combinator,omitempty"`
	Id         *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// if can be null:
	// true
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// if can be null:
	// true
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// if can be null:
	// true
	RuleSubType *string `json:"RuleSubType,omitempty" xml:"RuleSubType,omitempty"`
	// if can be null:
	// true
	RuleType *string   `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	Rules    []*Rule   `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	Values   []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s Rule) String() string {
	return dara.Prettify(s)
}

func (s Rule) GoString() string {
	return s.String()
}

func (s *Rule) GetCombinator() *string {
	return s.Combinator
}

func (s *Rule) GetId() *string {
	return s.Id
}

func (s *Rule) GetName() *string {
	return s.Name
}

func (s *Rule) GetOperator() *string {
	return s.Operator
}

func (s *Rule) GetRuleSubType() *string {
	return s.RuleSubType
}

func (s *Rule) GetRuleType() *string {
	return s.RuleType
}

func (s *Rule) GetRules() []*Rule {
	return s.Rules
}

func (s *Rule) GetValues() []*string {
	return s.Values
}

func (s *Rule) SetCombinator(v string) *Rule {
	s.Combinator = &v
	return s
}

func (s *Rule) SetId(v string) *Rule {
	s.Id = &v
	return s
}

func (s *Rule) SetName(v string) *Rule {
	s.Name = &v
	return s
}

func (s *Rule) SetOperator(v string) *Rule {
	s.Operator = &v
	return s
}

func (s *Rule) SetRuleSubType(v string) *Rule {
	s.RuleSubType = &v
	return s
}

func (s *Rule) SetRuleType(v string) *Rule {
	s.RuleType = &v
	return s
}

func (s *Rule) SetRules(v []*Rule) *Rule {
	s.Rules = v
	return s
}

func (s *Rule) SetValues(v []*string) *Rule {
	s.Values = v
	return s
}

func (s *Rule) Validate() error {
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
