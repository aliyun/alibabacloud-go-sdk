// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSystemConfigRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAggregationIds(v []*int32) *ListSystemConfigRulesRequest
	GetAggregationIds() []*int32
	SetCurrentPage(v int32) *ListSystemConfigRulesRequest
	GetCurrentPage() *int32
	SetLang(v string) *ListSystemConfigRulesRequest
	GetLang() *string
	SetPageSize(v int32) *ListSystemConfigRulesRequest
	GetPageSize() *int32
	SetRuleName(v string) *ListSystemConfigRulesRequest
	GetRuleName() *string
	SetRuleTypes(v []*int32) *ListSystemConfigRulesRequest
	GetRuleTypes() []*int32
	SetSelectionKey(v string) *ListSystemConfigRulesRequest
	GetSelectionKey() *string
	SetSystemType(v int32) *ListSystemConfigRulesRequest
	GetSystemType() *int32
}

type ListSystemConfigRulesRequest struct {
	// The list of system rule category IDs.
	AggregationIds []*int32 `json:"AggregationIds,omitempty" xml:"AggregationIds,omitempty" type:"Repeated"`
	// The current page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The language used for the rule name and description. Valid values:
	//
	// - **zh**: Chinese.
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The maximum number of entries to return on each page during paging in a paged query.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The rule name.
	//
	// example:
	//
	// Process startup
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The list of rule types for filtering.
	RuleTypes []*int32 `json:"RuleTypes,omitempty" xml:"RuleTypes,omitempty" type:"Repeated"`
	// The key for the switch configuration of rules associated with a policy. Use the SelectionKey saved for the same policy.
	//
	// example:
	//
	// 7efed2db-cc8c-4704-a1c3-d448c5d41e74
	SelectionKey *string `json:"SelectionKey,omitempty" xml:"SelectionKey,omitempty"`
	// The system platform identifier used for filtering. Valid values:
	//
	// - **1**: Linux.
	//
	// - **2**: Windows.
	//
	// - **3**: Mac.
	//
	// example:
	//
	// 1
	SystemType *int32 `json:"SystemType,omitempty" xml:"SystemType,omitempty"`
}

func (s ListSystemConfigRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSystemConfigRulesRequest) GoString() string {
	return s.String()
}

func (s *ListSystemConfigRulesRequest) GetAggregationIds() []*int32 {
	return s.AggregationIds
}

func (s *ListSystemConfigRulesRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListSystemConfigRulesRequest) GetLang() *string {
	return s.Lang
}

func (s *ListSystemConfigRulesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSystemConfigRulesRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *ListSystemConfigRulesRequest) GetRuleTypes() []*int32 {
	return s.RuleTypes
}

func (s *ListSystemConfigRulesRequest) GetSelectionKey() *string {
	return s.SelectionKey
}

func (s *ListSystemConfigRulesRequest) GetSystemType() *int32 {
	return s.SystemType
}

func (s *ListSystemConfigRulesRequest) SetAggregationIds(v []*int32) *ListSystemConfigRulesRequest {
	s.AggregationIds = v
	return s
}

func (s *ListSystemConfigRulesRequest) SetCurrentPage(v int32) *ListSystemConfigRulesRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListSystemConfigRulesRequest) SetLang(v string) *ListSystemConfigRulesRequest {
	s.Lang = &v
	return s
}

func (s *ListSystemConfigRulesRequest) SetPageSize(v int32) *ListSystemConfigRulesRequest {
	s.PageSize = &v
	return s
}

func (s *ListSystemConfigRulesRequest) SetRuleName(v string) *ListSystemConfigRulesRequest {
	s.RuleName = &v
	return s
}

func (s *ListSystemConfigRulesRequest) SetRuleTypes(v []*int32) *ListSystemConfigRulesRequest {
	s.RuleTypes = v
	return s
}

func (s *ListSystemConfigRulesRequest) SetSelectionKey(v string) *ListSystemConfigRulesRequest {
	s.SelectionKey = &v
	return s
}

func (s *ListSystemConfigRulesRequest) SetSystemType(v int32) *ListSystemConfigRulesRequest {
	s.SystemType = &v
	return s
}

func (s *ListSystemConfigRulesRequest) Validate() error {
	return dara.Validate(s)
}
