// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSystemAggregationRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAggregationIds(v []*int32) *ListSystemAggregationRulesRequest
	GetAggregationIds() []*int32
	SetCurrentPage(v int32) *ListSystemAggregationRulesRequest
	GetCurrentPage() *int32
	SetLang(v string) *ListSystemAggregationRulesRequest
	GetLang() *string
	SetPageSize(v int32) *ListSystemAggregationRulesRequest
	GetPageSize() *int32
	SetRuleName(v string) *ListSystemAggregationRulesRequest
	GetRuleName() *string
	SetRuleTypes(v []*int32) *ListSystemAggregationRulesRequest
	GetRuleTypes() []*int32
	SetSystemType(v int32) *ListSystemAggregationRulesRequest
	GetSystemType() *int32
}

type ListSystemAggregationRulesRequest struct {
	// The IDs of the aggregation types.
	AggregationIds []*int32 `json:"AggregationIds,omitempty" xml:"AggregationIds,omitempty" type:"Repeated"`
	// The number of the page to return.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The language of the content within the request and response. Default value: **zh**. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The number of entries to return on each page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The name of the rule.
	//
	// example:
	//
	// Rule\\*\\*\\*\\*
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The types of the rules.
	RuleTypes []*int32 `json:"RuleTypes,omitempty" xml:"RuleTypes,omitempty" type:"Repeated"`
	// The type of the operating system. Valid values:
	//
	// 	- **2**: Windows
	//
	// 	- **1**: Linux
	//
	// 	- **0**: all
	//
	// example:
	//
	// 0
	SystemType *int32 `json:"SystemType,omitempty" xml:"SystemType,omitempty"`
}

func (s ListSystemAggregationRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSystemAggregationRulesRequest) GoString() string {
	return s.String()
}

func (s *ListSystemAggregationRulesRequest) GetAggregationIds() []*int32 {
	return s.AggregationIds
}

func (s *ListSystemAggregationRulesRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListSystemAggregationRulesRequest) GetLang() *string {
	return s.Lang
}

func (s *ListSystemAggregationRulesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSystemAggregationRulesRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *ListSystemAggregationRulesRequest) GetRuleTypes() []*int32 {
	return s.RuleTypes
}

func (s *ListSystemAggregationRulesRequest) GetSystemType() *int32 {
	return s.SystemType
}

func (s *ListSystemAggregationRulesRequest) SetAggregationIds(v []*int32) *ListSystemAggregationRulesRequest {
	s.AggregationIds = v
	return s
}

func (s *ListSystemAggregationRulesRequest) SetCurrentPage(v int32) *ListSystemAggregationRulesRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListSystemAggregationRulesRequest) SetLang(v string) *ListSystemAggregationRulesRequest {
	s.Lang = &v
	return s
}

func (s *ListSystemAggregationRulesRequest) SetPageSize(v int32) *ListSystemAggregationRulesRequest {
	s.PageSize = &v
	return s
}

func (s *ListSystemAggregationRulesRequest) SetRuleName(v string) *ListSystemAggregationRulesRequest {
	s.RuleName = &v
	return s
}

func (s *ListSystemAggregationRulesRequest) SetRuleTypes(v []*int32) *ListSystemAggregationRulesRequest {
	s.RuleTypes = v
	return s
}

func (s *ListSystemAggregationRulesRequest) SetSystemType(v int32) *ListSystemAggregationRulesRequest {
	s.SystemType = &v
	return s
}

func (s *ListSystemAggregationRulesRequest) Validate() error {
	return dara.Validate(s)
}
