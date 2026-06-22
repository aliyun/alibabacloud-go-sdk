// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFileProtectClientRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlertLevel(v int32) *DeleteFileProtectClientRuleRequest
	GetAlertLevel() *int32
	SetExcludeIdList(v []*int64) *DeleteFileProtectClientRuleRequest
	GetExcludeIdList() []*int64
	SetIdList(v []*int64) *DeleteFileProtectClientRuleRequest
	GetIdList() []*int64
	SetPlatform(v string) *DeleteFileProtectClientRuleRequest
	GetPlatform() *string
	SetRuleAction(v string) *DeleteFileProtectClientRuleRequest
	GetRuleAction() *string
	SetRuleName(v string) *DeleteFileProtectClientRuleRequest
	GetRuleName() *string
	SetSelectAll(v bool) *DeleteFileProtectClientRuleRequest
	GetSelectAll() *bool
}

type DeleteFileProtectClientRuleRequest struct {
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
	// The list of excluded policy IDs.
	ExcludeIdList []*int64 `json:"ExcludeIdList,omitempty" xml:"ExcludeIdList,omitempty" type:"Repeated"`
	// The list of policy IDs.
	IdList []*int64 `json:"IdList,omitempty" xml:"IdList,omitempty" type:"Repeated"`
	// The type of the operating system. Valid values:
	//
	// - **windows**: Windows
	//
	// - **linux**: Linux.
	//
	// example:
	//
	// linux
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	// The action to take when the rule is triggered. Valid values:
	//
	// - **monitor**: Alert.
	//
	// - **block**: Block.
	//
	// - **pass**: Allow.
	//
	// example:
	//
	// pass
	RuleAction *string `json:"RuleAction,omitempty" xml:"RuleAction,omitempty"`
	// The name of the rule.
	//
	// example:
	//
	// text-001
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// Specifies whether to select all rules.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	SelectAll *bool `json:"SelectAll,omitempty" xml:"SelectAll,omitempty"`
}

func (s DeleteFileProtectClientRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteFileProtectClientRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteFileProtectClientRuleRequest) GetAlertLevel() *int32 {
	return s.AlertLevel
}

func (s *DeleteFileProtectClientRuleRequest) GetExcludeIdList() []*int64 {
	return s.ExcludeIdList
}

func (s *DeleteFileProtectClientRuleRequest) GetIdList() []*int64 {
	return s.IdList
}

func (s *DeleteFileProtectClientRuleRequest) GetPlatform() *string {
	return s.Platform
}

func (s *DeleteFileProtectClientRuleRequest) GetRuleAction() *string {
	return s.RuleAction
}

func (s *DeleteFileProtectClientRuleRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *DeleteFileProtectClientRuleRequest) GetSelectAll() *bool {
	return s.SelectAll
}

func (s *DeleteFileProtectClientRuleRequest) SetAlertLevel(v int32) *DeleteFileProtectClientRuleRequest {
	s.AlertLevel = &v
	return s
}

func (s *DeleteFileProtectClientRuleRequest) SetExcludeIdList(v []*int64) *DeleteFileProtectClientRuleRequest {
	s.ExcludeIdList = v
	return s
}

func (s *DeleteFileProtectClientRuleRequest) SetIdList(v []*int64) *DeleteFileProtectClientRuleRequest {
	s.IdList = v
	return s
}

func (s *DeleteFileProtectClientRuleRequest) SetPlatform(v string) *DeleteFileProtectClientRuleRequest {
	s.Platform = &v
	return s
}

func (s *DeleteFileProtectClientRuleRequest) SetRuleAction(v string) *DeleteFileProtectClientRuleRequest {
	s.RuleAction = &v
	return s
}

func (s *DeleteFileProtectClientRuleRequest) SetRuleName(v string) *DeleteFileProtectClientRuleRequest {
	s.RuleName = &v
	return s
}

func (s *DeleteFileProtectClientRuleRequest) SetSelectAll(v bool) *DeleteFileProtectClientRuleRequest {
	s.SelectAll = &v
	return s
}

func (s *DeleteFileProtectClientRuleRequest) Validate() error {
	return dara.Validate(s)
}
