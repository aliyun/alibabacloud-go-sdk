// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFileProtectClientRuleStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlertLevel(v int32) *UpdateFileProtectClientRuleStatusRequest
	GetAlertLevel() *int32
	SetExcludeIdList(v []*int64) *UpdateFileProtectClientRuleStatusRequest
	GetExcludeIdList() []*int64
	SetIdList(v []*int64) *UpdateFileProtectClientRuleStatusRequest
	GetIdList() []*int64
	SetPlatform(v string) *UpdateFileProtectClientRuleStatusRequest
	GetPlatform() *string
	SetRuleAction(v string) *UpdateFileProtectClientRuleStatusRequest
	GetRuleAction() *string
	SetRuleName(v string) *UpdateFileProtectClientRuleStatusRequest
	GetRuleName() *string
	SetSelectAll(v bool) *UpdateFileProtectClientRuleStatusRequest
	GetSelectAll() *bool
	SetStatus(v int32) *UpdateFileProtectClientRuleStatusRequest
	GetStatus() *int32
}

type UpdateFileProtectClientRuleStatusRequest struct {
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
	// test
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true
	SelectAll *bool `json:"SelectAll,omitempty" xml:"SelectAll,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateFileProtectClientRuleStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateFileProtectClientRuleStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateFileProtectClientRuleStatusRequest) GetAlertLevel() *int32 {
	return s.AlertLevel
}

func (s *UpdateFileProtectClientRuleStatusRequest) GetExcludeIdList() []*int64 {
	return s.ExcludeIdList
}

func (s *UpdateFileProtectClientRuleStatusRequest) GetIdList() []*int64 {
	return s.IdList
}

func (s *UpdateFileProtectClientRuleStatusRequest) GetPlatform() *string {
	return s.Platform
}

func (s *UpdateFileProtectClientRuleStatusRequest) GetRuleAction() *string {
	return s.RuleAction
}

func (s *UpdateFileProtectClientRuleStatusRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *UpdateFileProtectClientRuleStatusRequest) GetSelectAll() *bool {
	return s.SelectAll
}

func (s *UpdateFileProtectClientRuleStatusRequest) GetStatus() *int32 {
	return s.Status
}

func (s *UpdateFileProtectClientRuleStatusRequest) SetAlertLevel(v int32) *UpdateFileProtectClientRuleStatusRequest {
	s.AlertLevel = &v
	return s
}

func (s *UpdateFileProtectClientRuleStatusRequest) SetExcludeIdList(v []*int64) *UpdateFileProtectClientRuleStatusRequest {
	s.ExcludeIdList = v
	return s
}

func (s *UpdateFileProtectClientRuleStatusRequest) SetIdList(v []*int64) *UpdateFileProtectClientRuleStatusRequest {
	s.IdList = v
	return s
}

func (s *UpdateFileProtectClientRuleStatusRequest) SetPlatform(v string) *UpdateFileProtectClientRuleStatusRequest {
	s.Platform = &v
	return s
}

func (s *UpdateFileProtectClientRuleStatusRequest) SetRuleAction(v string) *UpdateFileProtectClientRuleStatusRequest {
	s.RuleAction = &v
	return s
}

func (s *UpdateFileProtectClientRuleStatusRequest) SetRuleName(v string) *UpdateFileProtectClientRuleStatusRequest {
	s.RuleName = &v
	return s
}

func (s *UpdateFileProtectClientRuleStatusRequest) SetSelectAll(v bool) *UpdateFileProtectClientRuleStatusRequest {
	s.SelectAll = &v
	return s
}

func (s *UpdateFileProtectClientRuleStatusRequest) SetStatus(v int32) *UpdateFileProtectClientRuleStatusRequest {
	s.Status = &v
	return s
}

func (s *UpdateFileProtectClientRuleStatusRequest) Validate() error {
	return dara.Validate(s)
}
