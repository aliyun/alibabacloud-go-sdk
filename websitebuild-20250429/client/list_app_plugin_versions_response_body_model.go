// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAppPluginVersionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ListAppPluginVersionsResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *ListAppPluginVersionsResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *ListAppPluginVersionsResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *ListAppPluginVersionsResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *ListAppPluginVersionsResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *ListAppPluginVersionsResponseBody
	GetErrorArgs() []interface{}
	SetModule(v *ListAppPluginVersionsResponseBodyModule) *ListAppPluginVersionsResponseBody
	GetModule() *ListAppPluginVersionsResponseBodyModule
	SetRequestId(v string) *ListAppPluginVersionsResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *ListAppPluginVersionsResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *ListAppPluginVersionsResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *ListAppPluginVersionsResponseBody
	GetSynchro() *bool
}

type ListAppPluginVersionsResponseBody struct {
	// The detailed reason why access was denied.
	//
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// Indicates whether retry is allowed.
	//
	// example:
	//
	// False
	AllowRetry *bool `json:"AllowRetry,omitempty" xml:"AllowRetry,omitempty"`
	// The application name. The application with this name is queried.
	//
	// example:
	//
	// spring-cloud-b
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The dynamic error code.
	//
	// example:
	//
	// ERROR-oo1
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// The dynamic message. Currently not in use. Ignore this parameter.
	//
	// example:
	//
	// abc
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The error parameters returned.
	ErrorArgs []interface{} `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	// The response data.
	Module *ListAppPluginVersionsResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 6C6B99AC-39EC-5350-874C-204128C905E6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The error code.
	//
	// example:
	//
	// SYSTEM.ERROR
	RootErrorCode *string `json:"RootErrorCode,omitempty" xml:"RootErrorCode,omitempty"`
	// The exception message.
	//
	// example:
	//
	// System exception.
	RootErrorMsg *string `json:"RootErrorMsg,omitempty" xml:"RootErrorMsg,omitempty"`
	// The reserved parameter.
	//
	// example:
	//
	// True
	Synchro *bool `json:"Synchro,omitempty" xml:"Synchro,omitempty"`
}

func (s ListAppPluginVersionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAppPluginVersionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAppPluginVersionsResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ListAppPluginVersionsResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *ListAppPluginVersionsResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *ListAppPluginVersionsResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *ListAppPluginVersionsResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *ListAppPluginVersionsResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *ListAppPluginVersionsResponseBody) GetModule() *ListAppPluginVersionsResponseBodyModule {
	return s.Module
}

func (s *ListAppPluginVersionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAppPluginVersionsResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *ListAppPluginVersionsResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *ListAppPluginVersionsResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *ListAppPluginVersionsResponseBody) SetAccessDeniedDetail(v string) *ListAppPluginVersionsResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListAppPluginVersionsResponseBody) SetAllowRetry(v bool) *ListAppPluginVersionsResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *ListAppPluginVersionsResponseBody) SetAppName(v string) *ListAppPluginVersionsResponseBody {
	s.AppName = &v
	return s
}

func (s *ListAppPluginVersionsResponseBody) SetDynamicCode(v string) *ListAppPluginVersionsResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *ListAppPluginVersionsResponseBody) SetDynamicMessage(v string) *ListAppPluginVersionsResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ListAppPluginVersionsResponseBody) SetErrorArgs(v []interface{}) *ListAppPluginVersionsResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *ListAppPluginVersionsResponseBody) SetModule(v *ListAppPluginVersionsResponseBodyModule) *ListAppPluginVersionsResponseBody {
	s.Module = v
	return s
}

func (s *ListAppPluginVersionsResponseBody) SetRequestId(v string) *ListAppPluginVersionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAppPluginVersionsResponseBody) SetRootErrorCode(v string) *ListAppPluginVersionsResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *ListAppPluginVersionsResponseBody) SetRootErrorMsg(v string) *ListAppPluginVersionsResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *ListAppPluginVersionsResponseBody) SetSynchro(v bool) *ListAppPluginVersionsResponseBody {
	s.Synchro = &v
	return s
}

func (s *ListAppPluginVersionsResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListAppPluginVersionsResponseBodyModule struct {
	// The current page number.
	//
	// example:
	//
	// 12
	CurrentPageNum *int32 `json:"CurrentPageNum,omitempty" xml:"CurrentPageNum,omitempty"`
	// The query results.
	Data []*ListAppPluginVersionsResponseBodyModuleData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The decision weight.
	Next *ListAppPluginVersionsResponseBodyModuleNext `json:"Next,omitempty" xml:"Next,omitempty" type:"Struct"`
	// Indicates whether a next page exists.
	//
	// example:
	//
	// False
	NextPage *bool `json:"NextPage,omitempty" xml:"NextPage,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Indicates whether a previous page exists.
	//
	// example:
	//
	// False
	PrePage *bool `json:"PrePage,omitempty" xml:"PrePage,omitempty"`
	// Apart from pagination limits, the server processes a maximum of 1000 recent records for the current query. If the results exceed 1000 records, **ResultLimit*	- is **true**, and you need to narrow the time range and search again. Otherwise, **ResultLimit*	- is **false**.
	//
	// example:
	//
	// False
	ResultLimit *bool `json:"ResultLimit,omitempty" xml:"ResultLimit,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 1
	TotalItemNum *int32 `json:"TotalItemNum,omitempty" xml:"TotalItemNum,omitempty"`
	// The total number of pages.
	//
	// example:
	//
	// 1
	TotalPageNum *int32 `json:"TotalPageNum,omitempty" xml:"TotalPageNum,omitempty"`
}

func (s ListAppPluginVersionsResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s ListAppPluginVersionsResponseBodyModule) GoString() string {
	return s.String()
}

func (s *ListAppPluginVersionsResponseBodyModule) GetCurrentPageNum() *int32 {
	return s.CurrentPageNum
}

func (s *ListAppPluginVersionsResponseBodyModule) GetData() []*ListAppPluginVersionsResponseBodyModuleData {
	return s.Data
}

func (s *ListAppPluginVersionsResponseBodyModule) GetNext() *ListAppPluginVersionsResponseBodyModuleNext {
	return s.Next
}

func (s *ListAppPluginVersionsResponseBodyModule) GetNextPage() *bool {
	return s.NextPage
}

func (s *ListAppPluginVersionsResponseBodyModule) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAppPluginVersionsResponseBodyModule) GetPrePage() *bool {
	return s.PrePage
}

func (s *ListAppPluginVersionsResponseBodyModule) GetResultLimit() *bool {
	return s.ResultLimit
}

func (s *ListAppPluginVersionsResponseBodyModule) GetTotalItemNum() *int32 {
	return s.TotalItemNum
}

func (s *ListAppPluginVersionsResponseBodyModule) GetTotalPageNum() *int32 {
	return s.TotalPageNum
}

func (s *ListAppPluginVersionsResponseBodyModule) SetCurrentPageNum(v int32) *ListAppPluginVersionsResponseBodyModule {
	s.CurrentPageNum = &v
	return s
}

func (s *ListAppPluginVersionsResponseBodyModule) SetData(v []*ListAppPluginVersionsResponseBodyModuleData) *ListAppPluginVersionsResponseBodyModule {
	s.Data = v
	return s
}

func (s *ListAppPluginVersionsResponseBodyModule) SetNext(v *ListAppPluginVersionsResponseBodyModuleNext) *ListAppPluginVersionsResponseBodyModule {
	s.Next = v
	return s
}

func (s *ListAppPluginVersionsResponseBodyModule) SetNextPage(v bool) *ListAppPluginVersionsResponseBodyModule {
	s.NextPage = &v
	return s
}

func (s *ListAppPluginVersionsResponseBodyModule) SetPageSize(v int32) *ListAppPluginVersionsResponseBodyModule {
	s.PageSize = &v
	return s
}

func (s *ListAppPluginVersionsResponseBodyModule) SetPrePage(v bool) *ListAppPluginVersionsResponseBodyModule {
	s.PrePage = &v
	return s
}

func (s *ListAppPluginVersionsResponseBodyModule) SetResultLimit(v bool) *ListAppPluginVersionsResponseBodyModule {
	s.ResultLimit = &v
	return s
}

func (s *ListAppPluginVersionsResponseBodyModule) SetTotalItemNum(v int32) *ListAppPluginVersionsResponseBodyModule {
	s.TotalItemNum = &v
	return s
}

func (s *ListAppPluginVersionsResponseBodyModule) SetTotalPageNum(v int32) *ListAppPluginVersionsResponseBodyModule {
	s.TotalPageNum = &v
	return s
}

func (s *ListAppPluginVersionsResponseBodyModule) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Next != nil {
		if err := s.Next.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListAppPluginVersionsResponseBodyModuleData struct {
	// The version changelog.
	//
	// example:
	//
	// x x x
	Changelog *string `json:"Changelog,omitempty" xml:"Changelog,omitempty"`
	// The commit ID.
	//
	// example:
	//
	// xxx
	CommitSha *string `json:"CommitSha,omitempty" xml:"CommitSha,omitempty"`
	// The creator.
	//
	// example:
	//
	// yinr
	CreatedBy *string `json:"CreatedBy,omitempty" xml:"CreatedBy,omitempty"`
	// The current number of files in the project.
	//
	// example:
	//
	// 4
	FileCount *int32 `json:"FileCount,omitempty" xml:"FileCount,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2026-06-24T02:29:03.000+00:00
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The file status.
	//
	// example:
	//
	// RUNNING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The application instance version.
	//
	// example:
	//
	// V2.0
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListAppPluginVersionsResponseBodyModuleData) String() string {
	return dara.Prettify(s)
}

func (s ListAppPluginVersionsResponseBodyModuleData) GoString() string {
	return s.String()
}

func (s *ListAppPluginVersionsResponseBodyModuleData) GetChangelog() *string {
	return s.Changelog
}

func (s *ListAppPluginVersionsResponseBodyModuleData) GetCommitSha() *string {
	return s.CommitSha
}

func (s *ListAppPluginVersionsResponseBodyModuleData) GetCreatedBy() *string {
	return s.CreatedBy
}

func (s *ListAppPluginVersionsResponseBodyModuleData) GetFileCount() *int32 {
	return s.FileCount
}

func (s *ListAppPluginVersionsResponseBodyModuleData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListAppPluginVersionsResponseBodyModuleData) GetStatus() *string {
	return s.Status
}

func (s *ListAppPluginVersionsResponseBodyModuleData) GetVersion() *string {
	return s.Version
}

func (s *ListAppPluginVersionsResponseBodyModuleData) SetChangelog(v string) *ListAppPluginVersionsResponseBodyModuleData {
	s.Changelog = &v
	return s
}

func (s *ListAppPluginVersionsResponseBodyModuleData) SetCommitSha(v string) *ListAppPluginVersionsResponseBodyModuleData {
	s.CommitSha = &v
	return s
}

func (s *ListAppPluginVersionsResponseBodyModuleData) SetCreatedBy(v string) *ListAppPluginVersionsResponseBodyModuleData {
	s.CreatedBy = &v
	return s
}

func (s *ListAppPluginVersionsResponseBodyModuleData) SetFileCount(v int32) *ListAppPluginVersionsResponseBodyModuleData {
	s.FileCount = &v
	return s
}

func (s *ListAppPluginVersionsResponseBodyModuleData) SetGmtCreate(v string) *ListAppPluginVersionsResponseBodyModuleData {
	s.GmtCreate = &v
	return s
}

func (s *ListAppPluginVersionsResponseBodyModuleData) SetStatus(v string) *ListAppPluginVersionsResponseBodyModuleData {
	s.Status = &v
	return s
}

func (s *ListAppPluginVersionsResponseBodyModuleData) SetVersion(v string) *ListAppPluginVersionsResponseBodyModuleData {
	s.Version = &v
	return s
}

func (s *ListAppPluginVersionsResponseBodyModuleData) Validate() error {
	return dara.Validate(s)
}

type ListAppPluginVersionsResponseBodyModuleNext struct {
	// The version changelog.
	//
	// example:
	//
	// xxx
	Changelog *string `json:"Changelog,omitempty" xml:"Changelog,omitempty"`
	// The commit ID.
	//
	// example:
	//
	// xxx
	CommitSha *string `json:"CommitSha,omitempty" xml:"CommitSha,omitempty"`
	// The creator.
	//
	// example:
	//
	// adminSubuser(264859988988056208)
	CreatedBy *string `json:"CreatedBy,omitempty" xml:"CreatedBy,omitempty"`
	// The current number of files in the project.
	//
	// example:
	//
	// 4
	FileCount *int32 `json:"FileCount,omitempty" xml:"FileCount,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2025-09-17 20:43:21
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// Valid values: unknown, init, testing, online.
	//
	// example:
	//
	// Succeeded
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The application instance version.
	//
	// example:
	//
	// V2.0
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListAppPluginVersionsResponseBodyModuleNext) String() string {
	return dara.Prettify(s)
}

func (s ListAppPluginVersionsResponseBodyModuleNext) GoString() string {
	return s.String()
}

func (s *ListAppPluginVersionsResponseBodyModuleNext) GetChangelog() *string {
	return s.Changelog
}

func (s *ListAppPluginVersionsResponseBodyModuleNext) GetCommitSha() *string {
	return s.CommitSha
}

func (s *ListAppPluginVersionsResponseBodyModuleNext) GetCreatedBy() *string {
	return s.CreatedBy
}

func (s *ListAppPluginVersionsResponseBodyModuleNext) GetFileCount() *int32 {
	return s.FileCount
}

func (s *ListAppPluginVersionsResponseBodyModuleNext) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListAppPluginVersionsResponseBodyModuleNext) GetStatus() *string {
	return s.Status
}

func (s *ListAppPluginVersionsResponseBodyModuleNext) GetVersion() *string {
	return s.Version
}

func (s *ListAppPluginVersionsResponseBodyModuleNext) SetChangelog(v string) *ListAppPluginVersionsResponseBodyModuleNext {
	s.Changelog = &v
	return s
}

func (s *ListAppPluginVersionsResponseBodyModuleNext) SetCommitSha(v string) *ListAppPluginVersionsResponseBodyModuleNext {
	s.CommitSha = &v
	return s
}

func (s *ListAppPluginVersionsResponseBodyModuleNext) SetCreatedBy(v string) *ListAppPluginVersionsResponseBodyModuleNext {
	s.CreatedBy = &v
	return s
}

func (s *ListAppPluginVersionsResponseBodyModuleNext) SetFileCount(v int32) *ListAppPluginVersionsResponseBodyModuleNext {
	s.FileCount = &v
	return s
}

func (s *ListAppPluginVersionsResponseBodyModuleNext) SetGmtCreate(v string) *ListAppPluginVersionsResponseBodyModuleNext {
	s.GmtCreate = &v
	return s
}

func (s *ListAppPluginVersionsResponseBodyModuleNext) SetStatus(v string) *ListAppPluginVersionsResponseBodyModuleNext {
	s.Status = &v
	return s
}

func (s *ListAppPluginVersionsResponseBodyModuleNext) SetVersion(v string) *ListAppPluginVersionsResponseBodyModuleNext {
	s.Version = &v
	return s
}

func (s *ListAppPluginVersionsResponseBodyModuleNext) Validate() error {
	return dara.Validate(s)
}
