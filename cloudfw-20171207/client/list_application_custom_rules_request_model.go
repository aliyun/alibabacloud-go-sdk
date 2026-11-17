// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApplicationCustomRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationCustomRuleId(v string) *ListApplicationCustomRulesRequest
	GetApplicationCustomRuleId() *string
	SetApplicationCustomRuleName(v string) *ListApplicationCustomRulesRequest
	GetApplicationCustomRuleName() *string
	SetApplicationId(v int64) *ListApplicationCustomRulesRequest
	GetApplicationId() *int64
	SetApplicationTemplateId(v string) *ListApplicationCustomRulesRequest
	GetApplicationTemplateId() *string
	SetLang(v string) *ListApplicationCustomRulesRequest
	GetLang() *string
	SetPageNo(v int64) *ListApplicationCustomRulesRequest
	GetPageNo() *int64
	SetPageSize(v int64) *ListApplicationCustomRulesRequest
	GetPageSize() *int64
	SetPriority(v int64) *ListApplicationCustomRulesRequest
	GetPriority() *int64
	SetRuleAction(v string) *ListApplicationCustomRulesRequest
	GetRuleAction() *string
}

type ListApplicationCustomRulesRequest struct {
	// The ID of the application custom rule. Exact match is supported.
	//
	// example:
	//
	// acr-7d2e4f6a8b0c1d3e
	ApplicationCustomRuleId *string `json:"ApplicationCustomRuleId,omitempty" xml:"ApplicationCustomRuleId,omitempty"`
	// The name of the custom rule. Fuzzy match is supported.
	//
	// example:
	//
	// Custom rule
	ApplicationCustomRuleName *string `json:"ApplicationCustomRuleName,omitempty" xml:"ApplicationCustomRuleName,omitempty"`
	// The ID of the application. This parameter is used to filter custom rules associated with the application. The application must exist in the built-in application signature library.
	//
	// example:
	//
	// 10012150000
	ApplicationId *int64 `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The ID of the application template to query. This parameter is required. The template must exist within the current Alibaba Cloud account.
	//
	// This parameter is required.
	//
	// example:
	//
	// ati-4f0691c56b8d4f06
	ApplicationTemplateId *string `json:"ApplicationTemplateId,omitempty" xml:"ApplicationTemplateId,omitempty"`
	// The language of the returned information. Valid values:
	//
	// - **zh**: Chinese. This is the default value.
	//
	// - **en**: English.
	//
	// This parameter determines whether the returned application names are in Chinese or English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The page number. Valid values: integers greater than or equal to 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. Valid values: integers greater than or equal to 1. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The priority of the rule. Exact match is supported. Valid values: integers from 1 to 200.
	//
	// example:
	//
	// 1
	Priority *int64 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The action of the rule. Exact match is supported. Valid values:
	//
	// - **Allow**: allows the traffic.
	//
	// - **Alert**: monitors the traffic.
	//
	// - **Deny**: blocks the traffic.
	//
	// example:
	//
	// Deny
	RuleAction *string `json:"RuleAction,omitempty" xml:"RuleAction,omitempty"`
}

func (s ListApplicationCustomRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationCustomRulesRequest) GoString() string {
	return s.String()
}

func (s *ListApplicationCustomRulesRequest) GetApplicationCustomRuleId() *string {
	return s.ApplicationCustomRuleId
}

func (s *ListApplicationCustomRulesRequest) GetApplicationCustomRuleName() *string {
	return s.ApplicationCustomRuleName
}

func (s *ListApplicationCustomRulesRequest) GetApplicationId() *int64 {
	return s.ApplicationId
}

func (s *ListApplicationCustomRulesRequest) GetApplicationTemplateId() *string {
	return s.ApplicationTemplateId
}

func (s *ListApplicationCustomRulesRequest) GetLang() *string {
	return s.Lang
}

func (s *ListApplicationCustomRulesRequest) GetPageNo() *int64 {
	return s.PageNo
}

func (s *ListApplicationCustomRulesRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListApplicationCustomRulesRequest) GetPriority() *int64 {
	return s.Priority
}

func (s *ListApplicationCustomRulesRequest) GetRuleAction() *string {
	return s.RuleAction
}

func (s *ListApplicationCustomRulesRequest) SetApplicationCustomRuleId(v string) *ListApplicationCustomRulesRequest {
	s.ApplicationCustomRuleId = &v
	return s
}

func (s *ListApplicationCustomRulesRequest) SetApplicationCustomRuleName(v string) *ListApplicationCustomRulesRequest {
	s.ApplicationCustomRuleName = &v
	return s
}

func (s *ListApplicationCustomRulesRequest) SetApplicationId(v int64) *ListApplicationCustomRulesRequest {
	s.ApplicationId = &v
	return s
}

func (s *ListApplicationCustomRulesRequest) SetApplicationTemplateId(v string) *ListApplicationCustomRulesRequest {
	s.ApplicationTemplateId = &v
	return s
}

func (s *ListApplicationCustomRulesRequest) SetLang(v string) *ListApplicationCustomRulesRequest {
	s.Lang = &v
	return s
}

func (s *ListApplicationCustomRulesRequest) SetPageNo(v int64) *ListApplicationCustomRulesRequest {
	s.PageNo = &v
	return s
}

func (s *ListApplicationCustomRulesRequest) SetPageSize(v int64) *ListApplicationCustomRulesRequest {
	s.PageSize = &v
	return s
}

func (s *ListApplicationCustomRulesRequest) SetPriority(v int64) *ListApplicationCustomRulesRequest {
	s.Priority = &v
	return s
}

func (s *ListApplicationCustomRulesRequest) SetRuleAction(v string) *ListApplicationCustomRulesRequest {
	s.RuleAction = &v
	return s
}

func (s *ListApplicationCustomRulesRequest) Validate() error {
	return dara.Validate(s)
}
