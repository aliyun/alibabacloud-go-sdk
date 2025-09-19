// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSystemClientRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAggregationIds(v []*int32) *ListSystemClientRulesRequest
	GetAggregationIds() []*int32
	SetCurrentPage(v int32) *ListSystemClientRulesRequest
	GetCurrentPage() *int32
	SetIsContainer(v int32) *ListSystemClientRulesRequest
	GetIsContainer() *int32
	SetLang(v string) *ListSystemClientRulesRequest
	GetLang() *string
	SetPageSize(v int32) *ListSystemClientRulesRequest
	GetPageSize() *int32
	SetRuleName(v string) *ListSystemClientRulesRequest
	GetRuleName() *string
	SetRuleTypes(v []*int32) *ListSystemClientRulesRequest
	GetRuleTypes() []*int32
	SetSystemType(v int32) *ListSystemClientRulesRequest
	GetSystemType() *int32
}

type ListSystemClientRulesRequest struct {
	// The IDs of the aggregation types for rules.
	AggregationIds []*int32 `json:"AggregationIds,omitempty" xml:"AggregationIds,omitempty" type:"Repeated"`
	// The number of the page to return.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// Specifies whether to query only container images. Valid values:
	//
	// 	- **0**: no
	//
	// 	- **1**: yes
	//
	// example:
	//
	// 1
	IsContainer *int32 `json:"IsContainer,omitempty" xml:"IsContainer,omitempty"`
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
	// The name of the system defense rule.
	//
	// example:
	//
	// Rule\\*\\*\\*\\*
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The types of the system defense rules.
	RuleTypes []*int32 `json:"RuleTypes,omitempty" xml:"RuleTypes,omitempty" type:"Repeated"`
	// The type of the OS. Valid values:
	//
	// 	- **2**: Windows
	//
	// 	- **1**: Linux
	//
	// 	- **0**: all types
	//
	// example:
	//
	// 0
	SystemType *int32 `json:"SystemType,omitempty" xml:"SystemType,omitempty"`
}

func (s ListSystemClientRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSystemClientRulesRequest) GoString() string {
	return s.String()
}

func (s *ListSystemClientRulesRequest) GetAggregationIds() []*int32 {
	return s.AggregationIds
}

func (s *ListSystemClientRulesRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListSystemClientRulesRequest) GetIsContainer() *int32 {
	return s.IsContainer
}

func (s *ListSystemClientRulesRequest) GetLang() *string {
	return s.Lang
}

func (s *ListSystemClientRulesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSystemClientRulesRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *ListSystemClientRulesRequest) GetRuleTypes() []*int32 {
	return s.RuleTypes
}

func (s *ListSystemClientRulesRequest) GetSystemType() *int32 {
	return s.SystemType
}

func (s *ListSystemClientRulesRequest) SetAggregationIds(v []*int32) *ListSystemClientRulesRequest {
	s.AggregationIds = v
	return s
}

func (s *ListSystemClientRulesRequest) SetCurrentPage(v int32) *ListSystemClientRulesRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListSystemClientRulesRequest) SetIsContainer(v int32) *ListSystemClientRulesRequest {
	s.IsContainer = &v
	return s
}

func (s *ListSystemClientRulesRequest) SetLang(v string) *ListSystemClientRulesRequest {
	s.Lang = &v
	return s
}

func (s *ListSystemClientRulesRequest) SetPageSize(v int32) *ListSystemClientRulesRequest {
	s.PageSize = &v
	return s
}

func (s *ListSystemClientRulesRequest) SetRuleName(v string) *ListSystemClientRulesRequest {
	s.RuleName = &v
	return s
}

func (s *ListSystemClientRulesRequest) SetRuleTypes(v []*int32) *ListSystemClientRulesRequest {
	s.RuleTypes = v
	return s
}

func (s *ListSystemClientRulesRequest) SetSystemType(v int32) *ListSystemClientRulesRequest {
	s.SystemType = &v
	return s
}

func (s *ListSystemClientRulesRequest) Validate() error {
	return dara.Validate(s)
}
