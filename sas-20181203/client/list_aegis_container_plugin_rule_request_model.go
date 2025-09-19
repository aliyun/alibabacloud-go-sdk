// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAegisContainerPluginRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCriteria(v string) *ListAegisContainerPluginRuleRequest
	GetCriteria() *string
	SetCurrentPage(v int32) *ListAegisContainerPluginRuleRequest
	GetCurrentPage() *int32
	SetLang(v string) *ListAegisContainerPluginRuleRequest
	GetLang() *string
	SetPageSize(v int32) *ListAegisContainerPluginRuleRequest
	GetPageSize() *int32
	SetRuleType(v int32) *ListAegisContainerPluginRuleRequest
	GetRuleType() *int32
}

type ListAegisContainerPluginRuleRequest struct {
	// The query condition.
	//
	// example:
	//
	// [{\\"name\\": \\"name\\", \\"value\\": \\"test-1818\\"}]
	Criteria *string `json:"Criteria,omitempty" xml:"Criteria,omitempty"`
	// The page number.
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
	// The number of entries per page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The type of the rule. Valid values:
	//
	// 	- **0**: custom
	//
	// 	- **1**: system
	//
	// example:
	//
	// 0
	RuleType *int32 `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
}

func (s ListAegisContainerPluginRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAegisContainerPluginRuleRequest) GoString() string {
	return s.String()
}

func (s *ListAegisContainerPluginRuleRequest) GetCriteria() *string {
	return s.Criteria
}

func (s *ListAegisContainerPluginRuleRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListAegisContainerPluginRuleRequest) GetLang() *string {
	return s.Lang
}

func (s *ListAegisContainerPluginRuleRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAegisContainerPluginRuleRequest) GetRuleType() *int32 {
	return s.RuleType
}

func (s *ListAegisContainerPluginRuleRequest) SetCriteria(v string) *ListAegisContainerPluginRuleRequest {
	s.Criteria = &v
	return s
}

func (s *ListAegisContainerPluginRuleRequest) SetCurrentPage(v int32) *ListAegisContainerPluginRuleRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListAegisContainerPluginRuleRequest) SetLang(v string) *ListAegisContainerPluginRuleRequest {
	s.Lang = &v
	return s
}

func (s *ListAegisContainerPluginRuleRequest) SetPageSize(v int32) *ListAegisContainerPluginRuleRequest {
	s.PageSize = &v
	return s
}

func (s *ListAegisContainerPluginRuleRequest) SetRuleType(v int32) *ListAegisContainerPluginRuleRequest {
	s.RuleType = &v
	return s
}

func (s *ListAegisContainerPluginRuleRequest) Validate() error {
	return dara.Validate(s)
}
