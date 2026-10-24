// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSystemConfigRulesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAggregationIdsShrink(v string) *ListSystemConfigRulesShrinkRequest
	GetAggregationIdsShrink() *string
	SetCurrentPage(v int32) *ListSystemConfigRulesShrinkRequest
	GetCurrentPage() *int32
	SetLang(v string) *ListSystemConfigRulesShrinkRequest
	GetLang() *string
	SetPageSize(v int32) *ListSystemConfigRulesShrinkRequest
	GetPageSize() *int32
	SetRuleName(v string) *ListSystemConfigRulesShrinkRequest
	GetRuleName() *string
	SetRuleTypesShrink(v string) *ListSystemConfigRulesShrinkRequest
	GetRuleTypesShrink() *string
	SetSelectionKey(v string) *ListSystemConfigRulesShrinkRequest
	GetSelectionKey() *string
	SetSystemType(v int32) *ListSystemConfigRulesShrinkRequest
	GetSystemType() *int32
}

type ListSystemConfigRulesShrinkRequest struct {
	// The list of system rule category IDs.
	AggregationIdsShrink *string `json:"AggregationIds,omitempty" xml:"AggregationIds,omitempty"`
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
	RuleTypesShrink *string `json:"RuleTypes,omitempty" xml:"RuleTypes,omitempty"`
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

func (s ListSystemConfigRulesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSystemConfigRulesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListSystemConfigRulesShrinkRequest) GetAggregationIdsShrink() *string {
	return s.AggregationIdsShrink
}

func (s *ListSystemConfigRulesShrinkRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListSystemConfigRulesShrinkRequest) GetLang() *string {
	return s.Lang
}

func (s *ListSystemConfigRulesShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSystemConfigRulesShrinkRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *ListSystemConfigRulesShrinkRequest) GetRuleTypesShrink() *string {
	return s.RuleTypesShrink
}

func (s *ListSystemConfigRulesShrinkRequest) GetSelectionKey() *string {
	return s.SelectionKey
}

func (s *ListSystemConfigRulesShrinkRequest) GetSystemType() *int32 {
	return s.SystemType
}

func (s *ListSystemConfigRulesShrinkRequest) SetAggregationIdsShrink(v string) *ListSystemConfigRulesShrinkRequest {
	s.AggregationIdsShrink = &v
	return s
}

func (s *ListSystemConfigRulesShrinkRequest) SetCurrentPage(v int32) *ListSystemConfigRulesShrinkRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListSystemConfigRulesShrinkRequest) SetLang(v string) *ListSystemConfigRulesShrinkRequest {
	s.Lang = &v
	return s
}

func (s *ListSystemConfigRulesShrinkRequest) SetPageSize(v int32) *ListSystemConfigRulesShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListSystemConfigRulesShrinkRequest) SetRuleName(v string) *ListSystemConfigRulesShrinkRequest {
	s.RuleName = &v
	return s
}

func (s *ListSystemConfigRulesShrinkRequest) SetRuleTypesShrink(v string) *ListSystemConfigRulesShrinkRequest {
	s.RuleTypesShrink = &v
	return s
}

func (s *ListSystemConfigRulesShrinkRequest) SetSelectionKey(v string) *ListSystemConfigRulesShrinkRequest {
	s.SelectionKey = &v
	return s
}

func (s *ListSystemConfigRulesShrinkRequest) SetSystemType(v int32) *ListSystemConfigRulesShrinkRequest {
	s.SystemType = &v
	return s
}

func (s *ListSystemConfigRulesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
