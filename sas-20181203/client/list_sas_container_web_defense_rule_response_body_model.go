// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSasContainerWebDefenseRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContainerWebDefenseRuleList(v []*ListSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRuleList) *ListSasContainerWebDefenseRuleResponseBody
	GetContainerWebDefenseRuleList() []*ListSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRuleList
	SetPageInfo(v *ListSasContainerWebDefenseRuleResponseBodyPageInfo) *ListSasContainerWebDefenseRuleResponseBody
	GetPageInfo() *ListSasContainerWebDefenseRuleResponseBodyPageInfo
	SetRequestId(v string) *ListSasContainerWebDefenseRuleResponseBody
	GetRequestId() *string
}

type ListSasContainerWebDefenseRuleResponseBody struct {
	// The rules for container tamper-proofing.
	ContainerWebDefenseRuleList []*ListSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRuleList `json:"ContainerWebDefenseRuleList,omitempty" xml:"ContainerWebDefenseRuleList,omitempty" type:"Repeated"`
	// The pagination information.
	PageInfo *ListSasContainerWebDefenseRuleResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The request ID, which is used to query logs and troubleshoot issues.
	//
	// example:
	//
	// 8C376***AE74FB4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListSasContainerWebDefenseRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSasContainerWebDefenseRuleResponseBody) GoString() string {
	return s.String()
}

func (s *ListSasContainerWebDefenseRuleResponseBody) GetContainerWebDefenseRuleList() []*ListSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRuleList {
	return s.ContainerWebDefenseRuleList
}

func (s *ListSasContainerWebDefenseRuleResponseBody) GetPageInfo() *ListSasContainerWebDefenseRuleResponseBodyPageInfo {
	return s.PageInfo
}

func (s *ListSasContainerWebDefenseRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSasContainerWebDefenseRuleResponseBody) SetContainerWebDefenseRuleList(v []*ListSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRuleList) *ListSasContainerWebDefenseRuleResponseBody {
	s.ContainerWebDefenseRuleList = v
	return s
}

func (s *ListSasContainerWebDefenseRuleResponseBody) SetPageInfo(v *ListSasContainerWebDefenseRuleResponseBodyPageInfo) *ListSasContainerWebDefenseRuleResponseBody {
	s.PageInfo = v
	return s
}

func (s *ListSasContainerWebDefenseRuleResponseBody) SetRequestId(v string) *ListSasContainerWebDefenseRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSasContainerWebDefenseRuleResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRuleList struct {
	// The user ID.
	//
	// example:
	//
	// 1766185894104675
	AliUid *int64 `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// The number of the applications.
	//
	// example:
	//
	// 10
	ApptotalCount *int32 `json:"ApptotalCount,omitempty" xml:"ApptotalCount,omitempty"`
	// The creation time. Unit: milliseconds.
	//
	// example:
	//
	// 1698978109000
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The timestamp when the alert event was last modified. Unit: milliseconds.
	//
	// example:
	//
	// 1678852686000
	GmtModified *int64 `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The ID of the rule.
	//
	// example:
	//
	// 420336648
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The paths that are protected.
	PathConfDTOList []*ListSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRuleListPathConfDTOList `json:"PathConfDTOList,omitempty" xml:"PathConfDTOList,omitempty" type:"Repeated"`
	// The name of the rule.
	//
	// example:
	//
	// test2
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The status of the rule. Valid values:
	//
	// 	- **1**: enabled
	//
	// 	- **0**: disabled
	//
	// example:
	//
	// 1
	RuleStatus *int32 `json:"RuleStatus,omitempty" xml:"RuleStatus,omitempty"`
}

func (s ListSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRuleList) String() string {
	return dara.Prettify(s)
}

func (s ListSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRuleList) GoString() string {
	return s.String()
}

func (s *ListSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRuleList) GetAliUid() *int64 {
	return s.AliUid
}

func (s *ListSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRuleList) GetApptotalCount() *int32 {
	return s.ApptotalCount
}

func (s *ListSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRuleList) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *ListSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRuleList) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *ListSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRuleList) GetId() *int64 {
	return s.Id
}

func (s *ListSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRuleList) GetPathConfDTOList() []*ListSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRuleListPathConfDTOList {
	return s.PathConfDTOList
}

func (s *ListSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRuleList) GetRuleName() *string {
	return s.RuleName
}

func (s *ListSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRuleList) GetRuleStatus() *int32 {
	return s.RuleStatus
}

func (s *ListSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRuleList) SetAliUid(v int64) *ListSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRuleList {
	s.AliUid = &v
	return s
}

func (s *ListSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRuleList) SetApptotalCount(v int32) *ListSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRuleList {
	s.ApptotalCount = &v
	return s
}

func (s *ListSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRuleList) SetGmtCreate(v int64) *ListSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRuleList {
	s.GmtCreate = &v
	return s
}

func (s *ListSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRuleList) SetGmtModified(v int64) *ListSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRuleList {
	s.GmtModified = &v
	return s
}

func (s *ListSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRuleList) SetId(v int64) *ListSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRuleList {
	s.Id = &v
	return s
}

func (s *ListSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRuleList) SetPathConfDTOList(v []*ListSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRuleListPathConfDTOList) *ListSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRuleList {
	s.PathConfDTOList = v
	return s
}

func (s *ListSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRuleList) SetRuleName(v string) *ListSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRuleList {
	s.RuleName = &v
	return s
}

func (s *ListSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRuleList) SetRuleStatus(v int32) *ListSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRuleList {
	s.RuleStatus = &v
	return s
}

func (s *ListSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRuleList) Validate() error {
	return dara.Validate(s)
}

type ListSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRuleListPathConfDTOList struct {
	// The backup paths.
	//
	// example:
	//
	// /tmp/test
	BackupPath *string `json:"BackupPath,omitempty" xml:"BackupPath,omitempty"`
	// The prevention mode. Valid values:
	//
	// 	- **block**
	//
	// 	- **audit**
	//
	// example:
	//
	// audit
	DefenseMode *string `json:"DefenseMode,omitempty" xml:"DefenseMode,omitempty"`
	// The protected path.
	//
	// example:
	//
	// /usr/test/
	DefensePath *string `json:"DefensePath,omitempty" xml:"DefensePath,omitempty"`
	// The file that is excluded.
	//
	// example:
	//
	// php
	ExcludeFile *string `json:"ExcludeFile,omitempty" xml:"ExcludeFile,omitempty"`
	// The path to the file that is excluded.
	//
	// example:
	//
	// /usr/tt
	ExcludeFilePath *string `json:"ExcludeFilePath,omitempty" xml:"ExcludeFilePath,omitempty"`
	// The type of the file that is excluded.
	//
	// example:
	//
	// jsp
	ExcludeFileType *string `json:"ExcludeFileType,omitempty" xml:"ExcludeFileType,omitempty"`
	// The protection mode. Valid values:
	//
	// 	- **0**: basic mode (whitelist)
	//
	// 	- **1**: complex mode (blacklist)
	//
	// example:
	//
	// 0
	GuardType *int32 `json:"GuardType,omitempty" xml:"GuardType,omitempty"`
	// The file that is included.
	//
	// example:
	//
	// /usr/test
	IncludeFile *string `json:"IncludeFile,omitempty" xml:"IncludeFile,omitempty"`
	// The type of the file that is included.
	//
	// example:
	//
	// php
	IncludeFileType *string `json:"IncludeFileType,omitempty" xml:"IncludeFileType,omitempty"`
	// The processes that are added to the whitelist.
	ProcessPathList []*string `json:"ProcessPathList,omitempty" xml:"ProcessPathList,omitempty" type:"Repeated"`
}

func (s ListSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRuleListPathConfDTOList) String() string {
	return dara.Prettify(s)
}

func (s ListSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRuleListPathConfDTOList) GoString() string {
	return s.String()
}

func (s *ListSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRuleListPathConfDTOList) GetBackupPath() *string {
	return s.BackupPath
}

func (s *ListSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRuleListPathConfDTOList) GetDefenseMode() *string {
	return s.DefenseMode
}

func (s *ListSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRuleListPathConfDTOList) GetDefensePath() *string {
	return s.DefensePath
}

func (s *ListSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRuleListPathConfDTOList) GetExcludeFile() *string {
	return s.ExcludeFile
}

func (s *ListSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRuleListPathConfDTOList) GetExcludeFilePath() *string {
	return s.ExcludeFilePath
}

func (s *ListSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRuleListPathConfDTOList) GetExcludeFileType() *string {
	return s.ExcludeFileType
}

func (s *ListSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRuleListPathConfDTOList) GetGuardType() *int32 {
	return s.GuardType
}

func (s *ListSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRuleListPathConfDTOList) GetIncludeFile() *string {
	return s.IncludeFile
}

func (s *ListSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRuleListPathConfDTOList) GetIncludeFileType() *string {
	return s.IncludeFileType
}

func (s *ListSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRuleListPathConfDTOList) GetProcessPathList() []*string {
	return s.ProcessPathList
}

func (s *ListSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRuleListPathConfDTOList) SetBackupPath(v string) *ListSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRuleListPathConfDTOList {
	s.BackupPath = &v
	return s
}

func (s *ListSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRuleListPathConfDTOList) SetDefenseMode(v string) *ListSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRuleListPathConfDTOList {
	s.DefenseMode = &v
	return s
}

func (s *ListSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRuleListPathConfDTOList) SetDefensePath(v string) *ListSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRuleListPathConfDTOList {
	s.DefensePath = &v
	return s
}

func (s *ListSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRuleListPathConfDTOList) SetExcludeFile(v string) *ListSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRuleListPathConfDTOList {
	s.ExcludeFile = &v
	return s
}

func (s *ListSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRuleListPathConfDTOList) SetExcludeFilePath(v string) *ListSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRuleListPathConfDTOList {
	s.ExcludeFilePath = &v
	return s
}

func (s *ListSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRuleListPathConfDTOList) SetExcludeFileType(v string) *ListSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRuleListPathConfDTOList {
	s.ExcludeFileType = &v
	return s
}

func (s *ListSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRuleListPathConfDTOList) SetGuardType(v int32) *ListSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRuleListPathConfDTOList {
	s.GuardType = &v
	return s
}

func (s *ListSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRuleListPathConfDTOList) SetIncludeFile(v string) *ListSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRuleListPathConfDTOList {
	s.IncludeFile = &v
	return s
}

func (s *ListSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRuleListPathConfDTOList) SetIncludeFileType(v string) *ListSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRuleListPathConfDTOList {
	s.IncludeFileType = &v
	return s
}

func (s *ListSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRuleListPathConfDTOList) SetProcessPathList(v []*string) *ListSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRuleListPathConfDTOList {
	s.ProcessPathList = v
	return s
}

func (s *ListSasContainerWebDefenseRuleResponseBodyContainerWebDefenseRuleListPathConfDTOList) Validate() error {
	return dara.Validate(s)
}

type ListSasContainerWebDefenseRuleResponseBodyPageInfo struct {
	// The number of entries returned on the current page.
	//
	// example:
	//
	// 4
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
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
	// 42
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListSasContainerWebDefenseRuleResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s ListSasContainerWebDefenseRuleResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *ListSasContainerWebDefenseRuleResponseBodyPageInfo) GetCount() *int32 {
	return s.Count
}

func (s *ListSasContainerWebDefenseRuleResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListSasContainerWebDefenseRuleResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSasContainerWebDefenseRuleResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListSasContainerWebDefenseRuleResponseBodyPageInfo) SetCount(v int32) *ListSasContainerWebDefenseRuleResponseBodyPageInfo {
	s.Count = &v
	return s
}

func (s *ListSasContainerWebDefenseRuleResponseBodyPageInfo) SetCurrentPage(v int32) *ListSasContainerWebDefenseRuleResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *ListSasContainerWebDefenseRuleResponseBodyPageInfo) SetPageSize(v int32) *ListSasContainerWebDefenseRuleResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *ListSasContainerWebDefenseRuleResponseBodyPageInfo) SetTotalCount(v int32) *ListSasContainerWebDefenseRuleResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *ListSasContainerWebDefenseRuleResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}
