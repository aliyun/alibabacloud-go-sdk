// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFileProtectRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlertLevel(v int32) *UpdateFileProtectRuleRequest
	GetAlertLevel() *int32
	SetFileOps(v []*string) *UpdateFileProtectRuleRequest
	GetFileOps() []*string
	SetFilePaths(v []*string) *UpdateFileProtectRuleRequest
	GetFilePaths() []*string
	SetId(v int64) *UpdateFileProtectRuleRequest
	GetId() *int64
	SetProcPaths(v []*string) *UpdateFileProtectRuleRequest
	GetProcPaths() []*string
	SetRuleAction(v string) *UpdateFileProtectRuleRequest
	GetRuleAction() *string
	SetRuleName(v string) *UpdateFileProtectRuleRequest
	GetRuleName() *string
	SetStatus(v int32) *UpdateFileProtectRuleRequest
	GetStatus() *int32
}

type UpdateFileProtectRuleRequest struct {
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
	// The operations that you want to perform on the files.
	//
	// This parameter is required.
	FileOps []*string `json:"FileOps,omitempty" xml:"FileOps,omitempty" type:"Repeated"`
	// The paths to the monitored files. Wildcard characters are supported.
	//
	// This parameter is required.
	FilePaths []*string `json:"FilePaths,omitempty" xml:"FilePaths,omitempty" type:"Repeated"`
	// The ID of the rule.
	//
	// example:
	//
	// 1062
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The paths to the monitored processes.
	//
	// This parameter is required.
	ProcPaths []*string `json:"ProcPaths,omitempty" xml:"ProcPaths,omitempty" type:"Repeated"`
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
	// This parameter is required.
	//
	// example:
	//
	// test-rule-1
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The status of the rule. Valid values:
	//
	// 	- **0**: disabled
	//
	// 	- **1**: enabled
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateFileProtectRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateFileProtectRuleRequest) GoString() string {
	return s.String()
}

func (s *UpdateFileProtectRuleRequest) GetAlertLevel() *int32 {
	return s.AlertLevel
}

func (s *UpdateFileProtectRuleRequest) GetFileOps() []*string {
	return s.FileOps
}

func (s *UpdateFileProtectRuleRequest) GetFilePaths() []*string {
	return s.FilePaths
}

func (s *UpdateFileProtectRuleRequest) GetId() *int64 {
	return s.Id
}

func (s *UpdateFileProtectRuleRequest) GetProcPaths() []*string {
	return s.ProcPaths
}

func (s *UpdateFileProtectRuleRequest) GetRuleAction() *string {
	return s.RuleAction
}

func (s *UpdateFileProtectRuleRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *UpdateFileProtectRuleRequest) GetStatus() *int32 {
	return s.Status
}

func (s *UpdateFileProtectRuleRequest) SetAlertLevel(v int32) *UpdateFileProtectRuleRequest {
	s.AlertLevel = &v
	return s
}

func (s *UpdateFileProtectRuleRequest) SetFileOps(v []*string) *UpdateFileProtectRuleRequest {
	s.FileOps = v
	return s
}

func (s *UpdateFileProtectRuleRequest) SetFilePaths(v []*string) *UpdateFileProtectRuleRequest {
	s.FilePaths = v
	return s
}

func (s *UpdateFileProtectRuleRequest) SetId(v int64) *UpdateFileProtectRuleRequest {
	s.Id = &v
	return s
}

func (s *UpdateFileProtectRuleRequest) SetProcPaths(v []*string) *UpdateFileProtectRuleRequest {
	s.ProcPaths = v
	return s
}

func (s *UpdateFileProtectRuleRequest) SetRuleAction(v string) *UpdateFileProtectRuleRequest {
	s.RuleAction = &v
	return s
}

func (s *UpdateFileProtectRuleRequest) SetRuleName(v string) *UpdateFileProtectRuleRequest {
	s.RuleName = &v
	return s
}

func (s *UpdateFileProtectRuleRequest) SetStatus(v int32) *UpdateFileProtectRuleRequest {
	s.Status = &v
	return s
}

func (s *UpdateFileProtectRuleRequest) Validate() error {
	return dara.Validate(s)
}
