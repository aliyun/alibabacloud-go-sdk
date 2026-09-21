// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWafRuleMatch interface {
	dara.Model
	String() string
	GoString() string
	SetConvertToLower(v bool) *WafRuleMatch
	GetConvertToLower() *bool
	SetCriteria(v []*WafRuleMatch) *WafRuleMatch
	GetCriteria() []*WafRuleMatch
	SetLogic(v string) *WafRuleMatch
	GetLogic() *string
	SetMatchOperator(v string) *WafRuleMatch
	GetMatchOperator() *string
	SetMatchType(v string) *WafRuleMatch
	GetMatchType() *string
	SetMatchValue(v interface{}) *WafRuleMatch
	GetMatchValue() interface{}
	SetNegate(v bool) *WafRuleMatch
	GetNegate() *bool
	SetParent(v string) *WafRuleMatch
	GetParent() *string
}

type WafRuleMatch struct {
	// The case-insensitive value setting.
	ConvertToLower *bool `json:"ConvertToLower,omitempty" xml:"ConvertToLower,omitempty"`
	// The logic list.
	Criteria []*WafRuleMatch `json:"Criteria,omitempty" xml:"Criteria,omitempty" type:"Repeated"`
	// The logical relationship.
	//
	// example:
	//
	// and
	Logic *string `json:"Logic,omitempty" xml:"Logic,omitempty"`
	// The match operator.
	//
	// example:
	//
	// eq
	MatchOperator *string `json:"MatchOperator,omitempty" xml:"MatchOperator,omitempty"`
	// The match field.
	//
	// example:
	//
	// ip.src
	MatchType *string `json:"MatchType,omitempty" xml:"MatchType,omitempty"`
	// The match value.
	//
	// example:
	//
	// 1.1.1.1
	MatchValue interface{} `json:"MatchValue,omitempty" xml:"MatchValue,omitempty"`
	// The negation of the match result.
	Negate *bool `json:"Negate,omitempty" xml:"Negate,omitempty"`
	// The primary row key of the parent group. This is used for two-level drop-down positioning of enumeration subkey fields. For example, the Parent of ali.websdk.umid is ali.websdk.
	//
	// example:
	//
	// ali.websdk
	Parent *string `json:"Parent,omitempty" xml:"Parent,omitempty"`
}

func (s WafRuleMatch) String() string {
	return dara.Prettify(s)
}

func (s WafRuleMatch) GoString() string {
	return s.String()
}

func (s *WafRuleMatch) GetConvertToLower() *bool {
	return s.ConvertToLower
}

func (s *WafRuleMatch) GetCriteria() []*WafRuleMatch {
	return s.Criteria
}

func (s *WafRuleMatch) GetLogic() *string {
	return s.Logic
}

func (s *WafRuleMatch) GetMatchOperator() *string {
	return s.MatchOperator
}

func (s *WafRuleMatch) GetMatchType() *string {
	return s.MatchType
}

func (s *WafRuleMatch) GetMatchValue() interface{} {
	return s.MatchValue
}

func (s *WafRuleMatch) GetNegate() *bool {
	return s.Negate
}

func (s *WafRuleMatch) GetParent() *string {
	return s.Parent
}

func (s *WafRuleMatch) SetConvertToLower(v bool) *WafRuleMatch {
	s.ConvertToLower = &v
	return s
}

func (s *WafRuleMatch) SetCriteria(v []*WafRuleMatch) *WafRuleMatch {
	s.Criteria = v
	return s
}

func (s *WafRuleMatch) SetLogic(v string) *WafRuleMatch {
	s.Logic = &v
	return s
}

func (s *WafRuleMatch) SetMatchOperator(v string) *WafRuleMatch {
	s.MatchOperator = &v
	return s
}

func (s *WafRuleMatch) SetMatchType(v string) *WafRuleMatch {
	s.MatchType = &v
	return s
}

func (s *WafRuleMatch) SetMatchValue(v interface{}) *WafRuleMatch {
	s.MatchValue = v
	return s
}

func (s *WafRuleMatch) SetNegate(v bool) *WafRuleMatch {
	s.Negate = &v
	return s
}

func (s *WafRuleMatch) SetParent(v string) *WafRuleMatch {
	s.Parent = &v
	return s
}

func (s *WafRuleMatch) Validate() error {
	if s.Criteria != nil {
		for _, item := range s.Criteria {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
