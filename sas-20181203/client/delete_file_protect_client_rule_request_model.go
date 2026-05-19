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
	// example:
	//
	// 0
	AlertLevel    *int32   `json:"AlertLevel,omitempty" xml:"AlertLevel,omitempty"`
	ExcludeIdList []*int64 `json:"ExcludeIdList,omitempty" xml:"ExcludeIdList,omitempty" type:"Repeated"`
	IdList        []*int64 `json:"IdList,omitempty" xml:"IdList,omitempty" type:"Repeated"`
	// example:
	//
	// linux
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	// example:
	//
	// pass
	RuleAction *string `json:"RuleAction,omitempty" xml:"RuleAction,omitempty"`
	// example:
	//
	// text-001
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
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
