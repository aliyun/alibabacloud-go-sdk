// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFileProtectClientRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlertLevel(v int32) *ListFileProtectClientRuleRequest
	GetAlertLevel() *int32
	SetCurrentPage(v int32) *ListFileProtectClientRuleRequest
	GetCurrentPage() *int32
	SetPageSize(v int32) *ListFileProtectClientRuleRequest
	GetPageSize() *int32
	SetPlatform(v string) *ListFileProtectClientRuleRequest
	GetPlatform() *string
	SetRuleAction(v string) *ListFileProtectClientRuleRequest
	GetRuleAction() *string
	SetRuleName(v string) *ListFileProtectClientRuleRequest
	GetRuleName() *string
}

type ListFileProtectClientRuleRequest struct {
	// The alert notification level. Valid values:
	//
	// - 0: no alert
	//
	// - 1: reminder
	//
	// - 2: suspicious
	//
	// - 3: high-risk.
	//
	// example:
	//
	// 0
	AlertLevel *int32 `json:"AlertLevel,omitempty" xml:"AlertLevel,omitempty"`
	// The page number of the current page in a paged query.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The maximum number of entries per page for a paged query.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The operating system type. Valid values:
	//
	// - **windows**: Windows
	//
	// - **linux**: Linux.
	//
	// example:
	//
	// linux
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	// The rule action. Valid values:
	//
	// - **block**: blocks the request.
	//
	// - **monitor**: monitors the request.
	//
	// - **pass**: allows the request.
	//
	// example:
	//
	// pass
	RuleAction *string `json:"RuleAction,omitempty" xml:"RuleAction,omitempty"`
	// The rule name.
	//
	// example:
	//
	// tetsRule
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s ListFileProtectClientRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s ListFileProtectClientRuleRequest) GoString() string {
	return s.String()
}

func (s *ListFileProtectClientRuleRequest) GetAlertLevel() *int32 {
	return s.AlertLevel
}

func (s *ListFileProtectClientRuleRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListFileProtectClientRuleRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListFileProtectClientRuleRequest) GetPlatform() *string {
	return s.Platform
}

func (s *ListFileProtectClientRuleRequest) GetRuleAction() *string {
	return s.RuleAction
}

func (s *ListFileProtectClientRuleRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *ListFileProtectClientRuleRequest) SetAlertLevel(v int32) *ListFileProtectClientRuleRequest {
	s.AlertLevel = &v
	return s
}

func (s *ListFileProtectClientRuleRequest) SetCurrentPage(v int32) *ListFileProtectClientRuleRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListFileProtectClientRuleRequest) SetPageSize(v int32) *ListFileProtectClientRuleRequest {
	s.PageSize = &v
	return s
}

func (s *ListFileProtectClientRuleRequest) SetPlatform(v string) *ListFileProtectClientRuleRequest {
	s.Platform = &v
	return s
}

func (s *ListFileProtectClientRuleRequest) SetRuleAction(v string) *ListFileProtectClientRuleRequest {
	s.RuleAction = &v
	return s
}

func (s *ListFileProtectClientRuleRequest) SetRuleName(v string) *ListFileProtectClientRuleRequest {
	s.RuleName = &v
	return s
}

func (s *ListFileProtectClientRuleRequest) Validate() error {
	return dara.Validate(s)
}
