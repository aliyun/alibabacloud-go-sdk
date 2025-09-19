// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFileProtectRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlertLevel(v int32) *CreateFileProtectRuleRequest
	GetAlertLevel() *int32
	SetFileOps(v []*string) *CreateFileProtectRuleRequest
	GetFileOps() []*string
	SetFilePaths(v []*string) *CreateFileProtectRuleRequest
	GetFilePaths() []*string
	SetPlatform(v string) *CreateFileProtectRuleRequest
	GetPlatform() *string
	SetProcPaths(v []*string) *CreateFileProtectRuleRequest
	GetProcPaths() []*string
	SetRuleAction(v string) *CreateFileProtectRuleRequest
	GetRuleAction() *string
	SetRuleName(v string) *CreateFileProtectRuleRequest
	GetRuleName() *string
	SetStatus(v int32) *CreateFileProtectRuleRequest
	GetStatus() *int32
	SetSwitchId(v string) *CreateFileProtectRuleRequest
	GetSwitchId() *string
}

type CreateFileProtectRuleRequest struct {
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
	// The paths to the files that you want to monitor. Wildcard characters are supported.
	//
	// This parameter is required.
	FilePaths []*string `json:"FilePaths,omitempty" xml:"FilePaths,omitempty" type:"Repeated"`
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
	// The paths to the processes that you want to monitor. Wildcard characters are supported.
	//
	// This parameter is required.
	ProcPaths []*string `json:"ProcPaths,omitempty" xml:"ProcPaths,omitempty" type:"Repeated"`
	// The handling method of the rule. Valid values:
	//
	// 	- pass: allow
	//
	// 	- alert
	//
	// This parameter is required.
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
	// tetsRule
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// Specifies whether to enable the rule. Valid values:
	//
	// 	- **1**: yes
	//
	// 	- **0**: no
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The switch ID of the rule.
	//
	// example:
	//
	// FILE_PROTECT_RULE_SWITCH_TYPE_0000
	SwitchId *string `json:"SwitchId,omitempty" xml:"SwitchId,omitempty"`
}

func (s CreateFileProtectRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateFileProtectRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateFileProtectRuleRequest) GetAlertLevel() *int32 {
	return s.AlertLevel
}

func (s *CreateFileProtectRuleRequest) GetFileOps() []*string {
	return s.FileOps
}

func (s *CreateFileProtectRuleRequest) GetFilePaths() []*string {
	return s.FilePaths
}

func (s *CreateFileProtectRuleRequest) GetPlatform() *string {
	return s.Platform
}

func (s *CreateFileProtectRuleRequest) GetProcPaths() []*string {
	return s.ProcPaths
}

func (s *CreateFileProtectRuleRequest) GetRuleAction() *string {
	return s.RuleAction
}

func (s *CreateFileProtectRuleRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *CreateFileProtectRuleRequest) GetStatus() *int32 {
	return s.Status
}

func (s *CreateFileProtectRuleRequest) GetSwitchId() *string {
	return s.SwitchId
}

func (s *CreateFileProtectRuleRequest) SetAlertLevel(v int32) *CreateFileProtectRuleRequest {
	s.AlertLevel = &v
	return s
}

func (s *CreateFileProtectRuleRequest) SetFileOps(v []*string) *CreateFileProtectRuleRequest {
	s.FileOps = v
	return s
}

func (s *CreateFileProtectRuleRequest) SetFilePaths(v []*string) *CreateFileProtectRuleRequest {
	s.FilePaths = v
	return s
}

func (s *CreateFileProtectRuleRequest) SetPlatform(v string) *CreateFileProtectRuleRequest {
	s.Platform = &v
	return s
}

func (s *CreateFileProtectRuleRequest) SetProcPaths(v []*string) *CreateFileProtectRuleRequest {
	s.ProcPaths = v
	return s
}

func (s *CreateFileProtectRuleRequest) SetRuleAction(v string) *CreateFileProtectRuleRequest {
	s.RuleAction = &v
	return s
}

func (s *CreateFileProtectRuleRequest) SetRuleName(v string) *CreateFileProtectRuleRequest {
	s.RuleName = &v
	return s
}

func (s *CreateFileProtectRuleRequest) SetStatus(v int32) *CreateFileProtectRuleRequest {
	s.Status = &v
	return s
}

func (s *CreateFileProtectRuleRequest) SetSwitchId(v string) *CreateFileProtectRuleRequest {
	s.SwitchId = &v
	return s
}

func (s *CreateFileProtectRuleRequest) Validate() error {
	return dara.Validate(s)
}
