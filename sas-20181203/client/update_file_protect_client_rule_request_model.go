// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFileProtectClientRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlertLevel(v int32) *UpdateFileProtectClientRuleRequest
	GetAlertLevel() *int32
	SetExcludeUsers(v []*string) *UpdateFileProtectClientRuleRequest
	GetExcludeUsers() []*string
	SetFileOps(v []*string) *UpdateFileProtectClientRuleRequest
	GetFileOps() []*string
	SetFilePaths(v []*string) *UpdateFileProtectClientRuleRequest
	GetFilePaths() []*string
	SetFileTypes(v []*string) *UpdateFileProtectClientRuleRequest
	GetFileTypes() []*string
	SetId(v int64) *UpdateFileProtectClientRuleRequest
	GetId() *int64
	SetProcPaths(v []*string) *UpdateFileProtectClientRuleRequest
	GetProcPaths() []*string
	SetRuleAction(v string) *UpdateFileProtectClientRuleRequest
	GetRuleAction() *string
	SetRuleName(v string) *UpdateFileProtectClientRuleRequest
	GetRuleName() *string
	SetStatus(v int32) *UpdateFileProtectClientRuleRequest
	GetStatus() *int32
}

type UpdateFileProtectClientRuleRequest struct {
	// example:
	//
	// 0
	AlertLevel   *int32    `json:"AlertLevel,omitempty" xml:"AlertLevel,omitempty"`
	ExcludeUsers []*string `json:"ExcludeUsers,omitempty" xml:"ExcludeUsers,omitempty" type:"Repeated"`
	// This parameter is required.
	FileOps []*string `json:"FileOps,omitempty" xml:"FileOps,omitempty" type:"Repeated"`
	// This parameter is required.
	FilePaths []*string `json:"FilePaths,omitempty" xml:"FilePaths,omitempty" type:"Repeated"`
	FileTypes []*string `json:"FileTypes,omitempty" xml:"FileTypes,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 245
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	ProcPaths []*string `json:"ProcPaths,omitempty" xml:"ProcPaths,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// pass
	RuleAction *string `json:"RuleAction,omitempty" xml:"RuleAction,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// text-001
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateFileProtectClientRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateFileProtectClientRuleRequest) GoString() string {
	return s.String()
}

func (s *UpdateFileProtectClientRuleRequest) GetAlertLevel() *int32 {
	return s.AlertLevel
}

func (s *UpdateFileProtectClientRuleRequest) GetExcludeUsers() []*string {
	return s.ExcludeUsers
}

func (s *UpdateFileProtectClientRuleRequest) GetFileOps() []*string {
	return s.FileOps
}

func (s *UpdateFileProtectClientRuleRequest) GetFilePaths() []*string {
	return s.FilePaths
}

func (s *UpdateFileProtectClientRuleRequest) GetFileTypes() []*string {
	return s.FileTypes
}

func (s *UpdateFileProtectClientRuleRequest) GetId() *int64 {
	return s.Id
}

func (s *UpdateFileProtectClientRuleRequest) GetProcPaths() []*string {
	return s.ProcPaths
}

func (s *UpdateFileProtectClientRuleRequest) GetRuleAction() *string {
	return s.RuleAction
}

func (s *UpdateFileProtectClientRuleRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *UpdateFileProtectClientRuleRequest) GetStatus() *int32 {
	return s.Status
}

func (s *UpdateFileProtectClientRuleRequest) SetAlertLevel(v int32) *UpdateFileProtectClientRuleRequest {
	s.AlertLevel = &v
	return s
}

func (s *UpdateFileProtectClientRuleRequest) SetExcludeUsers(v []*string) *UpdateFileProtectClientRuleRequest {
	s.ExcludeUsers = v
	return s
}

func (s *UpdateFileProtectClientRuleRequest) SetFileOps(v []*string) *UpdateFileProtectClientRuleRequest {
	s.FileOps = v
	return s
}

func (s *UpdateFileProtectClientRuleRequest) SetFilePaths(v []*string) *UpdateFileProtectClientRuleRequest {
	s.FilePaths = v
	return s
}

func (s *UpdateFileProtectClientRuleRequest) SetFileTypes(v []*string) *UpdateFileProtectClientRuleRequest {
	s.FileTypes = v
	return s
}

func (s *UpdateFileProtectClientRuleRequest) SetId(v int64) *UpdateFileProtectClientRuleRequest {
	s.Id = &v
	return s
}

func (s *UpdateFileProtectClientRuleRequest) SetProcPaths(v []*string) *UpdateFileProtectClientRuleRequest {
	s.ProcPaths = v
	return s
}

func (s *UpdateFileProtectClientRuleRequest) SetRuleAction(v string) *UpdateFileProtectClientRuleRequest {
	s.RuleAction = &v
	return s
}

func (s *UpdateFileProtectClientRuleRequest) SetRuleName(v string) *UpdateFileProtectClientRuleRequest {
	s.RuleName = &v
	return s
}

func (s *UpdateFileProtectClientRuleRequest) SetStatus(v int32) *UpdateFileProtectClientRuleRequest {
	s.Status = &v
	return s
}

func (s *UpdateFileProtectClientRuleRequest) Validate() error {
	return dara.Validate(s)
}
