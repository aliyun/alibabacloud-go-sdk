// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWafRuleMatch2 interface {
	dara.Model
	String() string
	GoString() string
	SetConvertToLower(v bool) *WafRuleMatch2
	GetConvertToLower() *bool
	SetCriteria(v []*WafRuleMatch2Criteria) *WafRuleMatch2
	GetCriteria() []*WafRuleMatch2Criteria
	SetLogic(v string) *WafRuleMatch2
	GetLogic() *string
	SetMatchOperator(v string) *WafRuleMatch2
	GetMatchOperator() *string
	SetMatchType(v string) *WafRuleMatch2
	GetMatchType() *string
	SetMatchValue(v interface{}) *WafRuleMatch2
	GetMatchValue() interface{}
	SetNegate(v bool) *WafRuleMatch2
	GetNegate() *bool
}

type WafRuleMatch2 struct {
	ConvertToLower *bool                    `json:"ConvertToLower,omitempty" xml:"ConvertToLower,omitempty"`
	Criteria       []*WafRuleMatch2Criteria `json:"Criteria,omitempty" xml:"Criteria,omitempty" type:"Repeated"`
	Logic          *string                  `json:"Logic,omitempty" xml:"Logic,omitempty"`
	MatchOperator  *string                  `json:"MatchOperator,omitempty" xml:"MatchOperator,omitempty"`
	MatchType      *string                  `json:"MatchType,omitempty" xml:"MatchType,omitempty"`
	MatchValue     interface{}              `json:"MatchValue,omitempty" xml:"MatchValue,omitempty"`
	Negate         *bool                    `json:"Negate,omitempty" xml:"Negate,omitempty"`
}

func (s WafRuleMatch2) String() string {
	return dara.Prettify(s)
}

func (s WafRuleMatch2) GoString() string {
	return s.String()
}

func (s *WafRuleMatch2) GetConvertToLower() *bool {
	return s.ConvertToLower
}

func (s *WafRuleMatch2) GetCriteria() []*WafRuleMatch2Criteria {
	return s.Criteria
}

func (s *WafRuleMatch2) GetLogic() *string {
	return s.Logic
}

func (s *WafRuleMatch2) GetMatchOperator() *string {
	return s.MatchOperator
}

func (s *WafRuleMatch2) GetMatchType() *string {
	return s.MatchType
}

func (s *WafRuleMatch2) GetMatchValue() interface{} {
	return s.MatchValue
}

func (s *WafRuleMatch2) GetNegate() *bool {
	return s.Negate
}

func (s *WafRuleMatch2) SetConvertToLower(v bool) *WafRuleMatch2 {
	s.ConvertToLower = &v
	return s
}

func (s *WafRuleMatch2) SetCriteria(v []*WafRuleMatch2Criteria) *WafRuleMatch2 {
	s.Criteria = v
	return s
}

func (s *WafRuleMatch2) SetLogic(v string) *WafRuleMatch2 {
	s.Logic = &v
	return s
}

func (s *WafRuleMatch2) SetMatchOperator(v string) *WafRuleMatch2 {
	s.MatchOperator = &v
	return s
}

func (s *WafRuleMatch2) SetMatchType(v string) *WafRuleMatch2 {
	s.MatchType = &v
	return s
}

func (s *WafRuleMatch2) SetMatchValue(v interface{}) *WafRuleMatch2 {
	s.MatchValue = v
	return s
}

func (s *WafRuleMatch2) SetNegate(v bool) *WafRuleMatch2 {
	s.Negate = &v
	return s
}

func (s *WafRuleMatch2) Validate() error {
	return dara.Validate(s)
}

type WafRuleMatch2Criteria struct {
	ConvertToLower *bool                            `json:"ConvertToLower,omitempty" xml:"ConvertToLower,omitempty"`
	Criteria       []*WafRuleMatch2CriteriaCriteria `json:"Criteria,omitempty" xml:"Criteria,omitempty" type:"Repeated"`
	Logic          *string                          `json:"Logic,omitempty" xml:"Logic,omitempty"`
	MatchOperator  *string                          `json:"MatchOperator,omitempty" xml:"MatchOperator,omitempty"`
	MatchType      *string                          `json:"MatchType,omitempty" xml:"MatchType,omitempty"`
	MatchValue     interface{}                      `json:"MatchValue,omitempty" xml:"MatchValue,omitempty"`
	Negate         *bool                            `json:"Negate,omitempty" xml:"Negate,omitempty"`
}

func (s WafRuleMatch2Criteria) String() string {
	return dara.Prettify(s)
}

func (s WafRuleMatch2Criteria) GoString() string {
	return s.String()
}

func (s *WafRuleMatch2Criteria) GetConvertToLower() *bool {
	return s.ConvertToLower
}

func (s *WafRuleMatch2Criteria) GetCriteria() []*WafRuleMatch2CriteriaCriteria {
	return s.Criteria
}

func (s *WafRuleMatch2Criteria) GetLogic() *string {
	return s.Logic
}

func (s *WafRuleMatch2Criteria) GetMatchOperator() *string {
	return s.MatchOperator
}

func (s *WafRuleMatch2Criteria) GetMatchType() *string {
	return s.MatchType
}

func (s *WafRuleMatch2Criteria) GetMatchValue() interface{} {
	return s.MatchValue
}

func (s *WafRuleMatch2Criteria) GetNegate() *bool {
	return s.Negate
}

func (s *WafRuleMatch2Criteria) SetConvertToLower(v bool) *WafRuleMatch2Criteria {
	s.ConvertToLower = &v
	return s
}

func (s *WafRuleMatch2Criteria) SetCriteria(v []*WafRuleMatch2CriteriaCriteria) *WafRuleMatch2Criteria {
	s.Criteria = v
	return s
}

func (s *WafRuleMatch2Criteria) SetLogic(v string) *WafRuleMatch2Criteria {
	s.Logic = &v
	return s
}

func (s *WafRuleMatch2Criteria) SetMatchOperator(v string) *WafRuleMatch2Criteria {
	s.MatchOperator = &v
	return s
}

func (s *WafRuleMatch2Criteria) SetMatchType(v string) *WafRuleMatch2Criteria {
	s.MatchType = &v
	return s
}

func (s *WafRuleMatch2Criteria) SetMatchValue(v interface{}) *WafRuleMatch2Criteria {
	s.MatchValue = v
	return s
}

func (s *WafRuleMatch2Criteria) SetNegate(v bool) *WafRuleMatch2Criteria {
	s.Negate = &v
	return s
}

func (s *WafRuleMatch2Criteria) Validate() error {
	return dara.Validate(s)
}

type WafRuleMatch2CriteriaCriteria struct {
	ConvertToLower *bool                                    `json:"ConvertToLower,omitempty" xml:"ConvertToLower,omitempty"`
	Criteria       []*WafRuleMatch2CriteriaCriteriaCriteria `json:"Criteria,omitempty" xml:"Criteria,omitempty" type:"Repeated"`
	Logic          *string                                  `json:"Logic,omitempty" xml:"Logic,omitempty"`
	MatchOperator  *string                                  `json:"MatchOperator,omitempty" xml:"MatchOperator,omitempty"`
	MatchType      *string                                  `json:"MatchType,omitempty" xml:"MatchType,omitempty"`
	MatchValue     interface{}                              `json:"MatchValue,omitempty" xml:"MatchValue,omitempty"`
	Negate         *bool                                    `json:"Negate,omitempty" xml:"Negate,omitempty"`
}

func (s WafRuleMatch2CriteriaCriteria) String() string {
	return dara.Prettify(s)
}

func (s WafRuleMatch2CriteriaCriteria) GoString() string {
	return s.String()
}

func (s *WafRuleMatch2CriteriaCriteria) GetConvertToLower() *bool {
	return s.ConvertToLower
}

func (s *WafRuleMatch2CriteriaCriteria) GetCriteria() []*WafRuleMatch2CriteriaCriteriaCriteria {
	return s.Criteria
}

func (s *WafRuleMatch2CriteriaCriteria) GetLogic() *string {
	return s.Logic
}

func (s *WafRuleMatch2CriteriaCriteria) GetMatchOperator() *string {
	return s.MatchOperator
}

func (s *WafRuleMatch2CriteriaCriteria) GetMatchType() *string {
	return s.MatchType
}

func (s *WafRuleMatch2CriteriaCriteria) GetMatchValue() interface{} {
	return s.MatchValue
}

func (s *WafRuleMatch2CriteriaCriteria) GetNegate() *bool {
	return s.Negate
}

func (s *WafRuleMatch2CriteriaCriteria) SetConvertToLower(v bool) *WafRuleMatch2CriteriaCriteria {
	s.ConvertToLower = &v
	return s
}

func (s *WafRuleMatch2CriteriaCriteria) SetCriteria(v []*WafRuleMatch2CriteriaCriteriaCriteria) *WafRuleMatch2CriteriaCriteria {
	s.Criteria = v
	return s
}

func (s *WafRuleMatch2CriteriaCriteria) SetLogic(v string) *WafRuleMatch2CriteriaCriteria {
	s.Logic = &v
	return s
}

func (s *WafRuleMatch2CriteriaCriteria) SetMatchOperator(v string) *WafRuleMatch2CriteriaCriteria {
	s.MatchOperator = &v
	return s
}

func (s *WafRuleMatch2CriteriaCriteria) SetMatchType(v string) *WafRuleMatch2CriteriaCriteria {
	s.MatchType = &v
	return s
}

func (s *WafRuleMatch2CriteriaCriteria) SetMatchValue(v interface{}) *WafRuleMatch2CriteriaCriteria {
	s.MatchValue = v
	return s
}

func (s *WafRuleMatch2CriteriaCriteria) SetNegate(v bool) *WafRuleMatch2CriteriaCriteria {
	s.Negate = &v
	return s
}

func (s *WafRuleMatch2CriteriaCriteria) Validate() error {
	return dara.Validate(s)
}

type WafRuleMatch2CriteriaCriteriaCriteria struct {
	ConvertToLower *bool       `json:"ConvertToLower,omitempty" xml:"ConvertToLower,omitempty"`
	MatchOperator  *string     `json:"MatchOperator,omitempty" xml:"MatchOperator,omitempty"`
	MatchType      *string     `json:"MatchType,omitempty" xml:"MatchType,omitempty"`
	MatchValue     interface{} `json:"MatchValue,omitempty" xml:"MatchValue,omitempty"`
	Negate         *bool       `json:"Negate,omitempty" xml:"Negate,omitempty"`
}

func (s WafRuleMatch2CriteriaCriteriaCriteria) String() string {
	return dara.Prettify(s)
}

func (s WafRuleMatch2CriteriaCriteriaCriteria) GoString() string {
	return s.String()
}

func (s *WafRuleMatch2CriteriaCriteriaCriteria) GetConvertToLower() *bool {
	return s.ConvertToLower
}

func (s *WafRuleMatch2CriteriaCriteriaCriteria) GetMatchOperator() *string {
	return s.MatchOperator
}

func (s *WafRuleMatch2CriteriaCriteriaCriteria) GetMatchType() *string {
	return s.MatchType
}

func (s *WafRuleMatch2CriteriaCriteriaCriteria) GetMatchValue() interface{} {
	return s.MatchValue
}

func (s *WafRuleMatch2CriteriaCriteriaCriteria) GetNegate() *bool {
	return s.Negate
}

func (s *WafRuleMatch2CriteriaCriteriaCriteria) SetConvertToLower(v bool) *WafRuleMatch2CriteriaCriteriaCriteria {
	s.ConvertToLower = &v
	return s
}

func (s *WafRuleMatch2CriteriaCriteriaCriteria) SetMatchOperator(v string) *WafRuleMatch2CriteriaCriteriaCriteria {
	s.MatchOperator = &v
	return s
}

func (s *WafRuleMatch2CriteriaCriteriaCriteria) SetMatchType(v string) *WafRuleMatch2CriteriaCriteriaCriteria {
	s.MatchType = &v
	return s
}

func (s *WafRuleMatch2CriteriaCriteriaCriteria) SetMatchValue(v interface{}) *WafRuleMatch2CriteriaCriteriaCriteria {
	s.MatchValue = v
	return s
}

func (s *WafRuleMatch2CriteriaCriteriaCriteria) SetNegate(v bool) *WafRuleMatch2CriteriaCriteriaCriteria {
	s.Negate = &v
	return s
}

func (s *WafRuleMatch2CriteriaCriteriaCriteria) Validate() error {
	return dara.Validate(s)
}
