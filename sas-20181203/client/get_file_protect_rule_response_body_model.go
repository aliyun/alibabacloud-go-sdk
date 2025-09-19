// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFileProtectRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetFileProtectRuleResponseBodyData) *GetFileProtectRuleResponseBody
	GetData() *GetFileProtectRuleResponseBodyData
	SetRequestId(v string) *GetFileProtectRuleResponseBody
	GetRequestId() *string
}

type GetFileProtectRuleResponseBody struct {
	// The response parameters.
	Data *GetFileProtectRuleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// C0DF9057-67C5-574D-A2D2-0CA9AC74C4D3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetFileProtectRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetFileProtectRuleResponseBody) GoString() string {
	return s.String()
}

func (s *GetFileProtectRuleResponseBody) GetData() *GetFileProtectRuleResponseBodyData {
	return s.Data
}

func (s *GetFileProtectRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetFileProtectRuleResponseBody) SetData(v *GetFileProtectRuleResponseBodyData) *GetFileProtectRuleResponseBody {
	s.Data = v
	return s
}

func (s *GetFileProtectRuleResponseBody) SetRequestId(v string) *GetFileProtectRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFileProtectRuleResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetFileProtectRuleResponseBodyData struct {
	// The handling method of the rule. Valid values:
	//
	// 1.  pass: allow
	//
	// 2.  alert
	//
	// example:
	//
	// pass
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
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
	// The operations performed on the files.
	FileOps []*string `json:"FileOps,omitempty" xml:"FileOps,omitempty" type:"Repeated"`
	// The paths to the monitored files.
	FilePaths []*string `json:"FilePaths,omitempty" xml:"FilePaths,omitempty" type:"Repeated"`
	// The ID of the rule.
	//
	// example:
	//
	// 44616
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
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
	// The paths to the monitored processes.
	ProcPaths []*string `json:"ProcPaths,omitempty" xml:"ProcPaths,omitempty" type:"Repeated"`
	// The name of the rule.
	//
	// example:
	//
	// test-000
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The status of the rule. Valid values:
	//
	// 1.  0: disabled
	//
	// 2.  1: enabled
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

func (s GetFileProtectRuleResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetFileProtectRuleResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetFileProtectRuleResponseBodyData) GetAction() *string {
	return s.Action
}

func (s *GetFileProtectRuleResponseBodyData) GetAlertLevel() *int32 {
	return s.AlertLevel
}

func (s *GetFileProtectRuleResponseBodyData) GetFileOps() []*string {
	return s.FileOps
}

func (s *GetFileProtectRuleResponseBodyData) GetFilePaths() []*string {
	return s.FilePaths
}

func (s *GetFileProtectRuleResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *GetFileProtectRuleResponseBodyData) GetPlatform() *string {
	return s.Platform
}

func (s *GetFileProtectRuleResponseBodyData) GetProcPaths() []*string {
	return s.ProcPaths
}

func (s *GetFileProtectRuleResponseBodyData) GetRuleName() *string {
	return s.RuleName
}

func (s *GetFileProtectRuleResponseBodyData) GetStatus() *int32 {
	return s.Status
}

func (s *GetFileProtectRuleResponseBodyData) GetSwitchId() *string {
	return s.SwitchId
}

func (s *GetFileProtectRuleResponseBodyData) SetAction(v string) *GetFileProtectRuleResponseBodyData {
	s.Action = &v
	return s
}

func (s *GetFileProtectRuleResponseBodyData) SetAlertLevel(v int32) *GetFileProtectRuleResponseBodyData {
	s.AlertLevel = &v
	return s
}

func (s *GetFileProtectRuleResponseBodyData) SetFileOps(v []*string) *GetFileProtectRuleResponseBodyData {
	s.FileOps = v
	return s
}

func (s *GetFileProtectRuleResponseBodyData) SetFilePaths(v []*string) *GetFileProtectRuleResponseBodyData {
	s.FilePaths = v
	return s
}

func (s *GetFileProtectRuleResponseBodyData) SetId(v int64) *GetFileProtectRuleResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetFileProtectRuleResponseBodyData) SetPlatform(v string) *GetFileProtectRuleResponseBodyData {
	s.Platform = &v
	return s
}

func (s *GetFileProtectRuleResponseBodyData) SetProcPaths(v []*string) *GetFileProtectRuleResponseBodyData {
	s.ProcPaths = v
	return s
}

func (s *GetFileProtectRuleResponseBodyData) SetRuleName(v string) *GetFileProtectRuleResponseBodyData {
	s.RuleName = &v
	return s
}

func (s *GetFileProtectRuleResponseBodyData) SetStatus(v int32) *GetFileProtectRuleResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetFileProtectRuleResponseBodyData) SetSwitchId(v string) *GetFileProtectRuleResponseBodyData {
	s.SwitchId = &v
	return s
}

func (s *GetFileProtectRuleResponseBodyData) Validate() error {
	return dara.Validate(s)
}
