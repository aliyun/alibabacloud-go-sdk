// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFileProtectRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFileProtectList(v []*ListFileProtectRuleResponseBodyFileProtectList) *ListFileProtectRuleResponseBody
	GetFileProtectList() []*ListFileProtectRuleResponseBodyFileProtectList
	SetPageInfo(v *ListFileProtectRuleResponseBodyPageInfo) *ListFileProtectRuleResponseBody
	GetPageInfo() *ListFileProtectRuleResponseBodyPageInfo
	SetRequestId(v string) *ListFileProtectRuleResponseBody
	GetRequestId() *string
}

type ListFileProtectRuleResponseBody struct {
	// The details of returned data.
	FileProtectList []*ListFileProtectRuleResponseBodyFileProtectList `json:"FileProtectList,omitempty" xml:"FileProtectList,omitempty" type:"Repeated"`
	// The pagination information.
	PageInfo *ListFileProtectRuleResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// FBBEB173-1F43-505F-A876-C03ECDF6CE4C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListFileProtectRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListFileProtectRuleResponseBody) GoString() string {
	return s.String()
}

func (s *ListFileProtectRuleResponseBody) GetFileProtectList() []*ListFileProtectRuleResponseBodyFileProtectList {
	return s.FileProtectList
}

func (s *ListFileProtectRuleResponseBody) GetPageInfo() *ListFileProtectRuleResponseBodyPageInfo {
	return s.PageInfo
}

func (s *ListFileProtectRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListFileProtectRuleResponseBody) SetFileProtectList(v []*ListFileProtectRuleResponseBodyFileProtectList) *ListFileProtectRuleResponseBody {
	s.FileProtectList = v
	return s
}

func (s *ListFileProtectRuleResponseBody) SetPageInfo(v *ListFileProtectRuleResponseBodyPageInfo) *ListFileProtectRuleResponseBody {
	s.PageInfo = v
	return s
}

func (s *ListFileProtectRuleResponseBody) SetRequestId(v string) *ListFileProtectRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFileProtectRuleResponseBody) Validate() error {
	if s.FileProtectList != nil {
		for _, item := range s.FileProtectList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.PageInfo != nil {
		if err := s.PageInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListFileProtectRuleResponseBodyFileProtectList struct {
	// The handling method of the rule. Valid values:
	//
	// 	- pass: allow
	//
	// 	- alert
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
	AlertLevel *string `json:"AlertLevel,omitempty" xml:"AlertLevel,omitempty"`
	// The total number of affected assets.
	//
	// example:
	//
	// 12
	EffectInstanceCount *int64 `json:"EffectInstanceCount,omitempty" xml:"EffectInstanceCount,omitempty"`
	// The operations performed on the files.
	FileOps []*string `json:"FileOps,omitempty" xml:"FileOps,omitempty" type:"Repeated"`
	// The paths to the monitored files. Wildcard characters are supported.
	FilePaths []*string `json:"FilePaths,omitempty" xml:"FilePaths,omitempty" type:"Repeated"`
	// The time when the rule was created.
	//
	// example:
	//
	// 1682304179000
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the rule was last modified.
	//
	// example:
	//
	// 1682304179000
	GmtModified *int64 `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The ID of the rule.
	//
	// example:
	//
	// 1412511
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
	// The paths to the monitored processes. Wildcard characters are supported.
	ProcPaths []*string `json:"ProcPaths,omitempty" xml:"ProcPaths,omitempty" type:"Repeated"`
	// The name of the rule.
	//
	// example:
	//
	// test11
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The status of the rule. Valid values:
	//
	// 	- 0: disabled
	//
	// 	- 1: enabled
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The switch ID of the rule.
	//
	// example:
	//
	// FILE_PROTECT_RULE_SWITCH_TYPE_1693474122927
	SwitchId *string `json:"SwitchId,omitempty" xml:"SwitchId,omitempty"`
}

func (s ListFileProtectRuleResponseBodyFileProtectList) String() string {
	return dara.Prettify(s)
}

func (s ListFileProtectRuleResponseBodyFileProtectList) GoString() string {
	return s.String()
}

func (s *ListFileProtectRuleResponseBodyFileProtectList) GetAction() *string {
	return s.Action
}

func (s *ListFileProtectRuleResponseBodyFileProtectList) GetAlertLevel() *string {
	return s.AlertLevel
}

func (s *ListFileProtectRuleResponseBodyFileProtectList) GetEffectInstanceCount() *int64 {
	return s.EffectInstanceCount
}

func (s *ListFileProtectRuleResponseBodyFileProtectList) GetFileOps() []*string {
	return s.FileOps
}

func (s *ListFileProtectRuleResponseBodyFileProtectList) GetFilePaths() []*string {
	return s.FilePaths
}

func (s *ListFileProtectRuleResponseBodyFileProtectList) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *ListFileProtectRuleResponseBodyFileProtectList) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *ListFileProtectRuleResponseBodyFileProtectList) GetId() *int64 {
	return s.Id
}

func (s *ListFileProtectRuleResponseBodyFileProtectList) GetPlatform() *string {
	return s.Platform
}

func (s *ListFileProtectRuleResponseBodyFileProtectList) GetProcPaths() []*string {
	return s.ProcPaths
}

func (s *ListFileProtectRuleResponseBodyFileProtectList) GetRuleName() *string {
	return s.RuleName
}

func (s *ListFileProtectRuleResponseBodyFileProtectList) GetStatus() *int32 {
	return s.Status
}

func (s *ListFileProtectRuleResponseBodyFileProtectList) GetSwitchId() *string {
	return s.SwitchId
}

func (s *ListFileProtectRuleResponseBodyFileProtectList) SetAction(v string) *ListFileProtectRuleResponseBodyFileProtectList {
	s.Action = &v
	return s
}

func (s *ListFileProtectRuleResponseBodyFileProtectList) SetAlertLevel(v string) *ListFileProtectRuleResponseBodyFileProtectList {
	s.AlertLevel = &v
	return s
}

func (s *ListFileProtectRuleResponseBodyFileProtectList) SetEffectInstanceCount(v int64) *ListFileProtectRuleResponseBodyFileProtectList {
	s.EffectInstanceCount = &v
	return s
}

func (s *ListFileProtectRuleResponseBodyFileProtectList) SetFileOps(v []*string) *ListFileProtectRuleResponseBodyFileProtectList {
	s.FileOps = v
	return s
}

func (s *ListFileProtectRuleResponseBodyFileProtectList) SetFilePaths(v []*string) *ListFileProtectRuleResponseBodyFileProtectList {
	s.FilePaths = v
	return s
}

func (s *ListFileProtectRuleResponseBodyFileProtectList) SetGmtCreate(v int64) *ListFileProtectRuleResponseBodyFileProtectList {
	s.GmtCreate = &v
	return s
}

func (s *ListFileProtectRuleResponseBodyFileProtectList) SetGmtModified(v int64) *ListFileProtectRuleResponseBodyFileProtectList {
	s.GmtModified = &v
	return s
}

func (s *ListFileProtectRuleResponseBodyFileProtectList) SetId(v int64) *ListFileProtectRuleResponseBodyFileProtectList {
	s.Id = &v
	return s
}

func (s *ListFileProtectRuleResponseBodyFileProtectList) SetPlatform(v string) *ListFileProtectRuleResponseBodyFileProtectList {
	s.Platform = &v
	return s
}

func (s *ListFileProtectRuleResponseBodyFileProtectList) SetProcPaths(v []*string) *ListFileProtectRuleResponseBodyFileProtectList {
	s.ProcPaths = v
	return s
}

func (s *ListFileProtectRuleResponseBodyFileProtectList) SetRuleName(v string) *ListFileProtectRuleResponseBodyFileProtectList {
	s.RuleName = &v
	return s
}

func (s *ListFileProtectRuleResponseBodyFileProtectList) SetStatus(v int32) *ListFileProtectRuleResponseBodyFileProtectList {
	s.Status = &v
	return s
}

func (s *ListFileProtectRuleResponseBodyFileProtectList) SetSwitchId(v string) *ListFileProtectRuleResponseBodyFileProtectList {
	s.SwitchId = &v
	return s
}

func (s *ListFileProtectRuleResponseBodyFileProtectList) Validate() error {
	return dara.Validate(s)
}

type ListFileProtectRuleResponseBodyPageInfo struct {
	// The page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 253
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListFileProtectRuleResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s ListFileProtectRuleResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *ListFileProtectRuleResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListFileProtectRuleResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListFileProtectRuleResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListFileProtectRuleResponseBodyPageInfo) SetCurrentPage(v int32) *ListFileProtectRuleResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *ListFileProtectRuleResponseBodyPageInfo) SetPageSize(v int32) *ListFileProtectRuleResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *ListFileProtectRuleResponseBodyPageInfo) SetTotalCount(v int32) *ListFileProtectRuleResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *ListFileProtectRuleResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}
