// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFileProtectClientRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFileProtectList(v []*ListFileProtectClientRuleResponseBodyFileProtectList) *ListFileProtectClientRuleResponseBody
	GetFileProtectList() []*ListFileProtectClientRuleResponseBodyFileProtectList
	SetPageInfo(v *ListFileProtectClientRuleResponseBodyPageInfo) *ListFileProtectClientRuleResponseBody
	GetPageInfo() *ListFileProtectClientRuleResponseBodyPageInfo
	SetRequestId(v string) *ListFileProtectClientRuleResponseBody
	GetRequestId() *string
}

type ListFileProtectClientRuleResponseBody struct {
	FileProtectList []*ListFileProtectClientRuleResponseBodyFileProtectList `json:"FileProtectList,omitempty" xml:"FileProtectList,omitempty" type:"Repeated"`
	PageInfo        *ListFileProtectClientRuleResponseBodyPageInfo          `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// example:
	//
	// BA674E4B-00CF-5DEA-8B92-360862FB****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListFileProtectClientRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListFileProtectClientRuleResponseBody) GoString() string {
	return s.String()
}

func (s *ListFileProtectClientRuleResponseBody) GetFileProtectList() []*ListFileProtectClientRuleResponseBodyFileProtectList {
	return s.FileProtectList
}

func (s *ListFileProtectClientRuleResponseBody) GetPageInfo() *ListFileProtectClientRuleResponseBodyPageInfo {
	return s.PageInfo
}

func (s *ListFileProtectClientRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListFileProtectClientRuleResponseBody) SetFileProtectList(v []*ListFileProtectClientRuleResponseBodyFileProtectList) *ListFileProtectClientRuleResponseBody {
	s.FileProtectList = v
	return s
}

func (s *ListFileProtectClientRuleResponseBody) SetPageInfo(v *ListFileProtectClientRuleResponseBodyPageInfo) *ListFileProtectClientRuleResponseBody {
	s.PageInfo = v
	return s
}

func (s *ListFileProtectClientRuleResponseBody) SetRequestId(v string) *ListFileProtectClientRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFileProtectClientRuleResponseBody) Validate() error {
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

type ListFileProtectClientRuleResponseBodyFileProtectList struct {
	// example:
	//
	// 0
	AlertLevel *int32 `json:"AlertLevel,omitempty" xml:"AlertLevel,omitempty"`
	// example:
	//
	// 12
	EffectInstanceCount *int64    `json:"EffectInstanceCount,omitempty" xml:"EffectInstanceCount,omitempty"`
	ExcludeUsers        []*string `json:"ExcludeUsers,omitempty" xml:"ExcludeUsers,omitempty" type:"Repeated"`
	FileOps             []*string `json:"FileOps,omitempty" xml:"FileOps,omitempty" type:"Repeated"`
	FilePaths           []*string `json:"FilePaths,omitempty" xml:"FilePaths,omitempty" type:"Repeated"`
	FileTypes           []*string `json:"FileTypes,omitempty" xml:"FileTypes,omitempty" type:"Repeated"`
	// example:
	//
	// 1682304179000
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 1682304179000
	GmtModified *int64 `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// 1412511
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
	// test11
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// FILE_PROTECT_RULE_SWITCH_TYPE_******
	SwitchId *string `json:"SwitchId,omitempty" xml:"SwitchId,omitempty"`
}

func (s ListFileProtectClientRuleResponseBodyFileProtectList) String() string {
	return dara.Prettify(s)
}

func (s ListFileProtectClientRuleResponseBodyFileProtectList) GoString() string {
	return s.String()
}

func (s *ListFileProtectClientRuleResponseBodyFileProtectList) GetAlertLevel() *int32 {
	return s.AlertLevel
}

func (s *ListFileProtectClientRuleResponseBodyFileProtectList) GetEffectInstanceCount() *int64 {
	return s.EffectInstanceCount
}

func (s *ListFileProtectClientRuleResponseBodyFileProtectList) GetExcludeUsers() []*string {
	return s.ExcludeUsers
}

func (s *ListFileProtectClientRuleResponseBodyFileProtectList) GetFileOps() []*string {
	return s.FileOps
}

func (s *ListFileProtectClientRuleResponseBodyFileProtectList) GetFilePaths() []*string {
	return s.FilePaths
}

func (s *ListFileProtectClientRuleResponseBodyFileProtectList) GetFileTypes() []*string {
	return s.FileTypes
}

func (s *ListFileProtectClientRuleResponseBodyFileProtectList) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *ListFileProtectClientRuleResponseBodyFileProtectList) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *ListFileProtectClientRuleResponseBodyFileProtectList) GetId() *int64 {
	return s.Id
}

func (s *ListFileProtectClientRuleResponseBodyFileProtectList) GetPlatform() *string {
	return s.Platform
}

func (s *ListFileProtectClientRuleResponseBodyFileProtectList) GetProcPaths() []*string {
	return s.ProcPaths
}

func (s *ListFileProtectClientRuleResponseBodyFileProtectList) GetRuleAction() *string {
	return s.RuleAction
}

func (s *ListFileProtectClientRuleResponseBodyFileProtectList) GetRuleName() *string {
	return s.RuleName
}

func (s *ListFileProtectClientRuleResponseBodyFileProtectList) GetStatus() *int32 {
	return s.Status
}

func (s *ListFileProtectClientRuleResponseBodyFileProtectList) GetSwitchId() *string {
	return s.SwitchId
}

func (s *ListFileProtectClientRuleResponseBodyFileProtectList) SetAlertLevel(v int32) *ListFileProtectClientRuleResponseBodyFileProtectList {
	s.AlertLevel = &v
	return s
}

func (s *ListFileProtectClientRuleResponseBodyFileProtectList) SetEffectInstanceCount(v int64) *ListFileProtectClientRuleResponseBodyFileProtectList {
	s.EffectInstanceCount = &v
	return s
}

func (s *ListFileProtectClientRuleResponseBodyFileProtectList) SetExcludeUsers(v []*string) *ListFileProtectClientRuleResponseBodyFileProtectList {
	s.ExcludeUsers = v
	return s
}

func (s *ListFileProtectClientRuleResponseBodyFileProtectList) SetFileOps(v []*string) *ListFileProtectClientRuleResponseBodyFileProtectList {
	s.FileOps = v
	return s
}

func (s *ListFileProtectClientRuleResponseBodyFileProtectList) SetFilePaths(v []*string) *ListFileProtectClientRuleResponseBodyFileProtectList {
	s.FilePaths = v
	return s
}

func (s *ListFileProtectClientRuleResponseBodyFileProtectList) SetFileTypes(v []*string) *ListFileProtectClientRuleResponseBodyFileProtectList {
	s.FileTypes = v
	return s
}

func (s *ListFileProtectClientRuleResponseBodyFileProtectList) SetGmtCreate(v int64) *ListFileProtectClientRuleResponseBodyFileProtectList {
	s.GmtCreate = &v
	return s
}

func (s *ListFileProtectClientRuleResponseBodyFileProtectList) SetGmtModified(v int64) *ListFileProtectClientRuleResponseBodyFileProtectList {
	s.GmtModified = &v
	return s
}

func (s *ListFileProtectClientRuleResponseBodyFileProtectList) SetId(v int64) *ListFileProtectClientRuleResponseBodyFileProtectList {
	s.Id = &v
	return s
}

func (s *ListFileProtectClientRuleResponseBodyFileProtectList) SetPlatform(v string) *ListFileProtectClientRuleResponseBodyFileProtectList {
	s.Platform = &v
	return s
}

func (s *ListFileProtectClientRuleResponseBodyFileProtectList) SetProcPaths(v []*string) *ListFileProtectClientRuleResponseBodyFileProtectList {
	s.ProcPaths = v
	return s
}

func (s *ListFileProtectClientRuleResponseBodyFileProtectList) SetRuleAction(v string) *ListFileProtectClientRuleResponseBodyFileProtectList {
	s.RuleAction = &v
	return s
}

func (s *ListFileProtectClientRuleResponseBodyFileProtectList) SetRuleName(v string) *ListFileProtectClientRuleResponseBodyFileProtectList {
	s.RuleName = &v
	return s
}

func (s *ListFileProtectClientRuleResponseBodyFileProtectList) SetStatus(v int32) *ListFileProtectClientRuleResponseBodyFileProtectList {
	s.Status = &v
	return s
}

func (s *ListFileProtectClientRuleResponseBodyFileProtectList) SetSwitchId(v string) *ListFileProtectClientRuleResponseBodyFileProtectList {
	s.SwitchId = &v
	return s
}

func (s *ListFileProtectClientRuleResponseBodyFileProtectList) Validate() error {
	return dara.Validate(s)
}

type ListFileProtectClientRuleResponseBodyPageInfo struct {
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 45
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListFileProtectClientRuleResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s ListFileProtectClientRuleResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *ListFileProtectClientRuleResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListFileProtectClientRuleResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListFileProtectClientRuleResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListFileProtectClientRuleResponseBodyPageInfo) SetCurrentPage(v int32) *ListFileProtectClientRuleResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *ListFileProtectClientRuleResponseBodyPageInfo) SetPageSize(v int32) *ListFileProtectClientRuleResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *ListFileProtectClientRuleResponseBodyPageInfo) SetTotalCount(v int32) *ListFileProtectClientRuleResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *ListFileProtectClientRuleResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}
