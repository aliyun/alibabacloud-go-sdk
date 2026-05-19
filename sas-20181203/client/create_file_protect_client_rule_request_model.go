// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFileProtectClientRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlertLevel(v int32) *CreateFileProtectClientRuleRequest
	GetAlertLevel() *int32
	SetClientToken(v string) *CreateFileProtectClientRuleRequest
	GetClientToken() *string
	SetExcludeUsers(v []*string) *CreateFileProtectClientRuleRequest
	GetExcludeUsers() []*string
	SetFileOps(v []*string) *CreateFileProtectClientRuleRequest
	GetFileOps() []*string
	SetFilePaths(v []*string) *CreateFileProtectClientRuleRequest
	GetFilePaths() []*string
	SetFileTypes(v []*string) *CreateFileProtectClientRuleRequest
	GetFileTypes() []*string
	SetPlatform(v string) *CreateFileProtectClientRuleRequest
	GetPlatform() *string
	SetProcPaths(v []*string) *CreateFileProtectClientRuleRequest
	GetProcPaths() []*string
	SetRuleAction(v string) *CreateFileProtectClientRuleRequest
	GetRuleAction() *string
	SetRuleName(v string) *CreateFileProtectClientRuleRequest
	GetRuleName() *string
	SetStatus(v int32) *CreateFileProtectClientRuleRequest
	GetStatus() *int32
	SetSwitchId(v string) *CreateFileProtectClientRuleRequest
	GetSwitchId() *string
}

type CreateFileProtectClientRuleRequest struct {
	// example:
	//
	// 0
	AlertLevel *int32 `json:"AlertLevel,omitempty" xml:"AlertLevel,omitempty"`
	// example:
	//
	// xxx
	ClientToken  *string   `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ExcludeUsers []*string `json:"ExcludeUsers,omitempty" xml:"ExcludeUsers,omitempty" type:"Repeated"`
	// This parameter is required.
	FileOps []*string `json:"FileOps,omitempty" xml:"FileOps,omitempty" type:"Repeated"`
	// This parameter is required.
	FilePaths []*string `json:"FilePaths,omitempty" xml:"FilePaths,omitempty" type:"Repeated"`
	FileTypes []*string `json:"FileTypes,omitempty" xml:"FileTypes,omitempty" type:"Repeated"`
	// example:
	//
	// linux
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	// This parameter is required.
	ProcPaths []*string `json:"ProcPaths,omitempty" xml:"ProcPaths,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// pass
	RuleAction *string `json:"RuleAction,omitempty" xml:"RuleAction,omitempty"`
	// This parameter is required.
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// FILE_PROTECT_RULE_SWITCH_TYPE_0000
	SwitchId *string `json:"SwitchId,omitempty" xml:"SwitchId,omitempty"`
}

func (s CreateFileProtectClientRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateFileProtectClientRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateFileProtectClientRuleRequest) GetAlertLevel() *int32 {
	return s.AlertLevel
}

func (s *CreateFileProtectClientRuleRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateFileProtectClientRuleRequest) GetExcludeUsers() []*string {
	return s.ExcludeUsers
}

func (s *CreateFileProtectClientRuleRequest) GetFileOps() []*string {
	return s.FileOps
}

func (s *CreateFileProtectClientRuleRequest) GetFilePaths() []*string {
	return s.FilePaths
}

func (s *CreateFileProtectClientRuleRequest) GetFileTypes() []*string {
	return s.FileTypes
}

func (s *CreateFileProtectClientRuleRequest) GetPlatform() *string {
	return s.Platform
}

func (s *CreateFileProtectClientRuleRequest) GetProcPaths() []*string {
	return s.ProcPaths
}

func (s *CreateFileProtectClientRuleRequest) GetRuleAction() *string {
	return s.RuleAction
}

func (s *CreateFileProtectClientRuleRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *CreateFileProtectClientRuleRequest) GetStatus() *int32 {
	return s.Status
}

func (s *CreateFileProtectClientRuleRequest) GetSwitchId() *string {
	return s.SwitchId
}

func (s *CreateFileProtectClientRuleRequest) SetAlertLevel(v int32) *CreateFileProtectClientRuleRequest {
	s.AlertLevel = &v
	return s
}

func (s *CreateFileProtectClientRuleRequest) SetClientToken(v string) *CreateFileProtectClientRuleRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateFileProtectClientRuleRequest) SetExcludeUsers(v []*string) *CreateFileProtectClientRuleRequest {
	s.ExcludeUsers = v
	return s
}

func (s *CreateFileProtectClientRuleRequest) SetFileOps(v []*string) *CreateFileProtectClientRuleRequest {
	s.FileOps = v
	return s
}

func (s *CreateFileProtectClientRuleRequest) SetFilePaths(v []*string) *CreateFileProtectClientRuleRequest {
	s.FilePaths = v
	return s
}

func (s *CreateFileProtectClientRuleRequest) SetFileTypes(v []*string) *CreateFileProtectClientRuleRequest {
	s.FileTypes = v
	return s
}

func (s *CreateFileProtectClientRuleRequest) SetPlatform(v string) *CreateFileProtectClientRuleRequest {
	s.Platform = &v
	return s
}

func (s *CreateFileProtectClientRuleRequest) SetProcPaths(v []*string) *CreateFileProtectClientRuleRequest {
	s.ProcPaths = v
	return s
}

func (s *CreateFileProtectClientRuleRequest) SetRuleAction(v string) *CreateFileProtectClientRuleRequest {
	s.RuleAction = &v
	return s
}

func (s *CreateFileProtectClientRuleRequest) SetRuleName(v string) *CreateFileProtectClientRuleRequest {
	s.RuleName = &v
	return s
}

func (s *CreateFileProtectClientRuleRequest) SetStatus(v int32) *CreateFileProtectClientRuleRequest {
	s.Status = &v
	return s
}

func (s *CreateFileProtectClientRuleRequest) SetSwitchId(v string) *CreateFileProtectClientRuleRequest {
	s.SwitchId = &v
	return s
}

func (s *CreateFileProtectClientRuleRequest) Validate() error {
	return dara.Validate(s)
}
