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
}

type WafRuleMatch struct {
	ConvertToLower *bool           `json:"ConvertToLower,omitempty" xml:"ConvertToLower,omitempty"`
	Criteria       []*WafRuleMatch `json:"Criteria,omitempty" xml:"Criteria,omitempty" type:"Repeated"`
	Logic          *string         `json:"Logic,omitempty" xml:"Logic,omitempty"`
	MatchOperator  *string         `json:"MatchOperator,omitempty" xml:"MatchOperator,omitempty"`
	MatchType      *string         `json:"MatchType,omitempty" xml:"MatchType,omitempty"`
	MatchValue     interface{}     `json:"MatchValue,omitempty" xml:"MatchValue,omitempty"`
	Negate         *bool           `json:"Negate,omitempty" xml:"Negate,omitempty"`
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
