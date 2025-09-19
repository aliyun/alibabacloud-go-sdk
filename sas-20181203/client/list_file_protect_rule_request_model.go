// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFileProtectRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlertLevel(v int32) *ListFileProtectRuleRequest
	GetAlertLevel() *int32
	SetCurrentPage(v int32) *ListFileProtectRuleRequest
	GetCurrentPage() *int32
	SetPageSize(v int32) *ListFileProtectRuleRequest
	GetPageSize() *int32
	SetPlatform(v string) *ListFileProtectRuleRequest
	GetPlatform() *string
	SetRuleAction(v string) *ListFileProtectRuleRequest
	GetRuleAction() *string
	SetRuleName(v string) *ListFileProtectRuleRequest
	GetRuleName() *string
}

type ListFileProtectRuleRequest struct {
	// The severity of alerts. Valid values:
	//
	// 	- 0: does not generate alerts
	//
	// 	- 1: sends notifications
	//
	// 	- 2: suspicious
	//
	// 	- 3: high-risk
	//
	// example:
	//
	// 0
	AlertLevel *int32 `json:"AlertLevel,omitempty" xml:"AlertLevel,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The type of the operating system. Valid values:
	//
	// 	- **windows**: Windows
	//
	// 	- **linux**: Linux
	//
	// example:
	//
	// linux
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	// The handling method of the rule. Valid values:
	//
	// 	- pass: allow
	//
	// 	- alert
	//
	// example:
	//
	// pass
	RuleAction *string `json:"RuleAction,omitempty" xml:"RuleAction,omitempty"`
	// The name of the rule.
	//
	// example:
	//
	// test-rule-1
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s ListFileProtectRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s ListFileProtectRuleRequest) GoString() string {
	return s.String()
}

func (s *ListFileProtectRuleRequest) GetAlertLevel() *int32 {
	return s.AlertLevel
}

func (s *ListFileProtectRuleRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListFileProtectRuleRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListFileProtectRuleRequest) GetPlatform() *string {
	return s.Platform
}

func (s *ListFileProtectRuleRequest) GetRuleAction() *string {
	return s.RuleAction
}

func (s *ListFileProtectRuleRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *ListFileProtectRuleRequest) SetAlertLevel(v int32) *ListFileProtectRuleRequest {
	s.AlertLevel = &v
	return s
}

func (s *ListFileProtectRuleRequest) SetCurrentPage(v int32) *ListFileProtectRuleRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListFileProtectRuleRequest) SetPageSize(v int32) *ListFileProtectRuleRequest {
	s.PageSize = &v
	return s
}

func (s *ListFileProtectRuleRequest) SetPlatform(v string) *ListFileProtectRuleRequest {
	s.Platform = &v
	return s
}

func (s *ListFileProtectRuleRequest) SetRuleAction(v string) *ListFileProtectRuleRequest {
	s.RuleAction = &v
	return s
}

func (s *ListFileProtectRuleRequest) SetRuleName(v string) *ListFileProtectRuleRequest {
	s.RuleName = &v
	return s
}

func (s *ListFileProtectRuleRequest) Validate() error {
	return dara.Validate(s)
}
