// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListContainerDefenseRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConditions(v []*ListContainerDefenseRuleRequestConditions) *ListContainerDefenseRuleRequest
	GetConditions() []*ListContainerDefenseRuleRequestConditions
	SetCurrentPage(v int32) *ListContainerDefenseRuleRequest
	GetCurrentPage() *int32
	SetIsDefaultRule(v int32) *ListContainerDefenseRuleRequest
	GetIsDefaultRule() *int32
	SetLang(v string) *ListContainerDefenseRuleRequest
	GetLang() *string
	SetPageSize(v int32) *ListContainerDefenseRuleRequest
	GetPageSize() *int32
	SetRuleType(v int32) *ListContainerDefenseRuleRequest
	GetRuleType() *int32
}

type ListContainerDefenseRuleRequest struct {
	// The details of the condition.
	Conditions []*ListContainerDefenseRuleRequestConditions `json:"Conditions,omitempty" xml:"Conditions,omitempty" type:"Repeated"`
	// The number of the page to return. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// Specifies whether to query system rules.
	//
	// >  This parameter is deprecated.
	//
	// example:
	//
	// 1
	IsDefaultRule *int32 `json:"IsDefaultRule,omitempty" xml:"IsDefaultRule,omitempty"`
	// The language of the content within the request and response. Default value: **zh**. Valid values:
	//
	// 	- **zh**: Chinese.
	//
	// 	- **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The number of entries per page. Default value: 20. If you leave this parameter empty, 20 entries are returned on each page.
	//
	// >  We recommend that you do not leave this parameter empty.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The rule type. Valid values:
	//
	// 	- 1: system rule
	//
	// 	- 2: user-defined rule
	//
	// example:
	//
	// 1
	RuleType *int32 `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
}

func (s ListContainerDefenseRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s ListContainerDefenseRuleRequest) GoString() string {
	return s.String()
}

func (s *ListContainerDefenseRuleRequest) GetConditions() []*ListContainerDefenseRuleRequestConditions {
	return s.Conditions
}

func (s *ListContainerDefenseRuleRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListContainerDefenseRuleRequest) GetIsDefaultRule() *int32 {
	return s.IsDefaultRule
}

func (s *ListContainerDefenseRuleRequest) GetLang() *string {
	return s.Lang
}

func (s *ListContainerDefenseRuleRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListContainerDefenseRuleRequest) GetRuleType() *int32 {
	return s.RuleType
}

func (s *ListContainerDefenseRuleRequest) SetConditions(v []*ListContainerDefenseRuleRequestConditions) *ListContainerDefenseRuleRequest {
	s.Conditions = v
	return s
}

func (s *ListContainerDefenseRuleRequest) SetCurrentPage(v int32) *ListContainerDefenseRuleRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListContainerDefenseRuleRequest) SetIsDefaultRule(v int32) *ListContainerDefenseRuleRequest {
	s.IsDefaultRule = &v
	return s
}

func (s *ListContainerDefenseRuleRequest) SetLang(v string) *ListContainerDefenseRuleRequest {
	s.Lang = &v
	return s
}

func (s *ListContainerDefenseRuleRequest) SetPageSize(v int32) *ListContainerDefenseRuleRequest {
	s.PageSize = &v
	return s
}

func (s *ListContainerDefenseRuleRequest) SetRuleType(v int32) *ListContainerDefenseRuleRequest {
	s.RuleType = &v
	return s
}

func (s *ListContainerDefenseRuleRequest) Validate() error {
	if s.Conditions != nil {
		for _, item := range s.Conditions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListContainerDefenseRuleRequestConditions struct {
	// The condition type. Valid values:
	//
	// 	- **ruleName**: the rule name
	//
	// example:
	//
	// ruleName
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The rule content.
	//
	// example:
	//
	// auto-test-rule-**
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListContainerDefenseRuleRequestConditions) String() string {
	return dara.Prettify(s)
}

func (s ListContainerDefenseRuleRequestConditions) GoString() string {
	return s.String()
}

func (s *ListContainerDefenseRuleRequestConditions) GetType() *string {
	return s.Type
}

func (s *ListContainerDefenseRuleRequestConditions) GetValue() *string {
	return s.Value
}

func (s *ListContainerDefenseRuleRequestConditions) SetType(v string) *ListContainerDefenseRuleRequestConditions {
	s.Type = &v
	return s
}

func (s *ListContainerDefenseRuleRequestConditions) SetValue(v string) *ListContainerDefenseRuleRequestConditions {
	s.Value = &v
	return s
}

func (s *ListContainerDefenseRuleRequestConditions) Validate() error {
	return dara.Validate(s)
}
