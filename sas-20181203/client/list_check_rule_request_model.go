// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCheckRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCheckId(v int64) *ListCheckRuleRequest
	GetCheckId() *int64
	SetCheckName(v string) *ListCheckRuleRequest
	GetCheckName() *string
	SetCurrentPage(v int32) *ListCheckRuleRequest
	GetCurrentPage() *int32
	SetLang(v string) *ListCheckRuleRequest
	GetLang() *string
	SetPageSize(v int32) *ListCheckRuleRequest
	GetPageSize() *int32
	SetRuleType(v string) *ListCheckRuleRequest
	GetRuleType() *string
	SetScopeType(v string) *ListCheckRuleRequest
	GetScopeType() *string
	SetTaskSources(v []*string) *ListCheckRuleRequest
	GetTaskSources() []*string
}

type ListCheckRuleRequest struct {
	// The ID of the check item.
	//
	// > You can call the [ListCheckResult](~~ListCheckResult~~) API to get the check item ID.
	//
	// example:
	//
	// 58
	CheckId *int64 `json:"CheckId,omitempty" xml:"CheckId,omitempty"`
	// The name of the check item.
	//
	// example:
	//
	// checkName
	CheckName *string `json:"CheckName,omitempty" xml:"CheckName,omitempty"`
	// The page number displayed in a paginated query.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// Set the language type for the request and response messages. The default is **zh**. Values:
	//
	// - zh: Chinese
	//
	// - en: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The number of check items displayed per page in a paginated query. The default value is **20**, indicating 20 check items per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The type of rule. Default is **WHITE**. Values:
	//
	// - **WHITE**: Add to whitelist
	//
	// example:
	//
	// WHITE
	RuleType *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	// The scope where the rule applies. Values:
	//
	// - **INSTNACE**: Instance
	//
	// - **ITEM**: Check item
	//
	// example:
	//
	// INSTANCE
	ScopeType   *string   `json:"ScopeType,omitempty" xml:"ScopeType,omitempty"`
	TaskSources []*string `json:"TaskSources,omitempty" xml:"TaskSources,omitempty" type:"Repeated"`
}

func (s ListCheckRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCheckRuleRequest) GoString() string {
	return s.String()
}

func (s *ListCheckRuleRequest) GetCheckId() *int64 {
	return s.CheckId
}

func (s *ListCheckRuleRequest) GetCheckName() *string {
	return s.CheckName
}

func (s *ListCheckRuleRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListCheckRuleRequest) GetLang() *string {
	return s.Lang
}

func (s *ListCheckRuleRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCheckRuleRequest) GetRuleType() *string {
	return s.RuleType
}

func (s *ListCheckRuleRequest) GetScopeType() *string {
	return s.ScopeType
}

func (s *ListCheckRuleRequest) GetTaskSources() []*string {
	return s.TaskSources
}

func (s *ListCheckRuleRequest) SetCheckId(v int64) *ListCheckRuleRequest {
	s.CheckId = &v
	return s
}

func (s *ListCheckRuleRequest) SetCheckName(v string) *ListCheckRuleRequest {
	s.CheckName = &v
	return s
}

func (s *ListCheckRuleRequest) SetCurrentPage(v int32) *ListCheckRuleRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListCheckRuleRequest) SetLang(v string) *ListCheckRuleRequest {
	s.Lang = &v
	return s
}

func (s *ListCheckRuleRequest) SetPageSize(v int32) *ListCheckRuleRequest {
	s.PageSize = &v
	return s
}

func (s *ListCheckRuleRequest) SetRuleType(v string) *ListCheckRuleRequest {
	s.RuleType = &v
	return s
}

func (s *ListCheckRuleRequest) SetScopeType(v string) *ListCheckRuleRequest {
	s.ScopeType = &v
	return s
}

func (s *ListCheckRuleRequest) SetTaskSources(v []*string) *ListCheckRuleRequest {
	s.TaskSources = v
	return s
}

func (s *ListCheckRuleRequest) Validate() error {
	return dara.Validate(s)
}
