// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFileProtectClientRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetFileProtectClientRuleResponseBodyData) *GetFileProtectClientRuleResponseBody
	GetData() *GetFileProtectClientRuleResponseBodyData
	SetRequestId(v string) *GetFileProtectClientRuleResponseBody
	GetRequestId() *string
}

type GetFileProtectClientRuleResponseBody struct {
	Data *GetFileProtectClientRuleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 79CFF74D-E967-5407-8A78-EE03B925****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetFileProtectClientRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetFileProtectClientRuleResponseBody) GoString() string {
	return s.String()
}

func (s *GetFileProtectClientRuleResponseBody) GetData() *GetFileProtectClientRuleResponseBodyData {
	return s.Data
}

func (s *GetFileProtectClientRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetFileProtectClientRuleResponseBody) SetData(v *GetFileProtectClientRuleResponseBodyData) *GetFileProtectClientRuleResponseBody {
	s.Data = v
	return s
}

func (s *GetFileProtectClientRuleResponseBody) SetRequestId(v string) *GetFileProtectClientRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFileProtectClientRuleResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetFileProtectClientRuleResponseBodyData struct {
	// example:
	//
	// 1
	AlertLevel   *int32    `json:"AlertLevel,omitempty" xml:"AlertLevel,omitempty"`
	ExcludeUsers []*string `json:"ExcludeUsers,omitempty" xml:"ExcludeUsers,omitempty" type:"Repeated"`
	FileOps      []*string `json:"FileOps,omitempty" xml:"FileOps,omitempty" type:"Repeated"`
	FilePaths    []*string `json:"FilePaths,omitempty" xml:"FilePaths,omitempty" type:"Repeated"`
	FileTypes    []*string `json:"FileTypes,omitempty" xml:"FileTypes,omitempty" type:"Repeated"`
	// example:
	//
	// 3119
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// linux
	Platform  *string   `json:"Platform,omitempty" xml:"Platform,omitempty"`
	ProcPaths []*string `json:"ProcPaths,omitempty" xml:"ProcPaths,omitempty" type:"Repeated"`
	// example:
	//
	// pass
	RuleAction *string `json:"RuleAction,omitempty" xml:"RuleAction,omitempty"`
	// example:
	//
	// test
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// USER-CONTAINER-RULE-SWITCH-TYPE_***
	SwitchId *string `json:"SwitchId,omitempty" xml:"SwitchId,omitempty"`
}

func (s GetFileProtectClientRuleResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetFileProtectClientRuleResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetFileProtectClientRuleResponseBodyData) GetAlertLevel() *int32 {
	return s.AlertLevel
}

func (s *GetFileProtectClientRuleResponseBodyData) GetExcludeUsers() []*string {
	return s.ExcludeUsers
}

func (s *GetFileProtectClientRuleResponseBodyData) GetFileOps() []*string {
	return s.FileOps
}

func (s *GetFileProtectClientRuleResponseBodyData) GetFilePaths() []*string {
	return s.FilePaths
}

func (s *GetFileProtectClientRuleResponseBodyData) GetFileTypes() []*string {
	return s.FileTypes
}

func (s *GetFileProtectClientRuleResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *GetFileProtectClientRuleResponseBodyData) GetPlatform() *string {
	return s.Platform
}

func (s *GetFileProtectClientRuleResponseBodyData) GetProcPaths() []*string {
	return s.ProcPaths
}

func (s *GetFileProtectClientRuleResponseBodyData) GetRuleAction() *string {
	return s.RuleAction
}

func (s *GetFileProtectClientRuleResponseBodyData) GetRuleName() *string {
	return s.RuleName
}

func (s *GetFileProtectClientRuleResponseBodyData) GetStatus() *int32 {
	return s.Status
}

func (s *GetFileProtectClientRuleResponseBodyData) GetSwitchId() *string {
	return s.SwitchId
}

func (s *GetFileProtectClientRuleResponseBodyData) SetAlertLevel(v int32) *GetFileProtectClientRuleResponseBodyData {
	s.AlertLevel = &v
	return s
}

func (s *GetFileProtectClientRuleResponseBodyData) SetExcludeUsers(v []*string) *GetFileProtectClientRuleResponseBodyData {
	s.ExcludeUsers = v
	return s
}

func (s *GetFileProtectClientRuleResponseBodyData) SetFileOps(v []*string) *GetFileProtectClientRuleResponseBodyData {
	s.FileOps = v
	return s
}

func (s *GetFileProtectClientRuleResponseBodyData) SetFilePaths(v []*string) *GetFileProtectClientRuleResponseBodyData {
	s.FilePaths = v
	return s
}

func (s *GetFileProtectClientRuleResponseBodyData) SetFileTypes(v []*string) *GetFileProtectClientRuleResponseBodyData {
	s.FileTypes = v
	return s
}

func (s *GetFileProtectClientRuleResponseBodyData) SetId(v int64) *GetFileProtectClientRuleResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetFileProtectClientRuleResponseBodyData) SetPlatform(v string) *GetFileProtectClientRuleResponseBodyData {
	s.Platform = &v
	return s
}

func (s *GetFileProtectClientRuleResponseBodyData) SetProcPaths(v []*string) *GetFileProtectClientRuleResponseBodyData {
	s.ProcPaths = v
	return s
}

func (s *GetFileProtectClientRuleResponseBodyData) SetRuleAction(v string) *GetFileProtectClientRuleResponseBodyData {
	s.RuleAction = &v
	return s
}

func (s *GetFileProtectClientRuleResponseBodyData) SetRuleName(v string) *GetFileProtectClientRuleResponseBodyData {
	s.RuleName = &v
	return s
}

func (s *GetFileProtectClientRuleResponseBodyData) SetStatus(v int32) *GetFileProtectClientRuleResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetFileProtectClientRuleResponseBodyData) SetSwitchId(v string) *GetFileProtectClientRuleResponseBodyData {
	s.SwitchId = &v
	return s
}

func (s *GetFileProtectClientRuleResponseBodyData) Validate() error {
	return dara.Validate(s)
}
