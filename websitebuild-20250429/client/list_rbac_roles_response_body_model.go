// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRbacRolesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ListRbacRolesResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *ListRbacRolesResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *ListRbacRolesResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *ListRbacRolesResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *ListRbacRolesResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *ListRbacRolesResponseBody
	GetErrorArgs() []interface{}
	SetMaxResults(v int32) *ListRbacRolesResponseBody
	GetMaxResults() *int32
	SetModule(v *ListRbacRolesResponseBodyModule) *ListRbacRolesResponseBody
	GetModule() *ListRbacRolesResponseBodyModule
	SetNextToken(v string) *ListRbacRolesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListRbacRolesResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *ListRbacRolesResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *ListRbacRolesResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *ListRbacRolesResponseBody
	GetSynchro() *bool
}

type ListRbacRolesResponseBody struct {
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// False
	AllowRetry *bool `json:"AllowRetry,omitempty" xml:"AllowRetry,omitempty"`
	// example:
	//
	// or
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// ERROR-oo1
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// example:
	//
	// SYSTEM_ERROR
	DynamicMessage *string       `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorArgs      []interface{} `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	MaxResults *int32                           `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	Module     *ListRbacRolesResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
	// example:
	//
	// 0l45bkwM022Dt+rOvPi/oQ==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 6C6B99AC-39EC-5350-874C-204128C905E6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// SYSTEM.ERROR
	RootErrorCode *string `json:"RootErrorCode,omitempty" xml:"RootErrorCode,omitempty"`
	// example:
	//
	// 系统异常
	RootErrorMsg *string `json:"RootErrorMsg,omitempty" xml:"RootErrorMsg,omitempty"`
	// example:
	//
	// True
	Synchro *bool `json:"Synchro,omitempty" xml:"Synchro,omitempty"`
}

func (s ListRbacRolesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRbacRolesResponseBody) GoString() string {
	return s.String()
}

func (s *ListRbacRolesResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ListRbacRolesResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *ListRbacRolesResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *ListRbacRolesResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *ListRbacRolesResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *ListRbacRolesResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *ListRbacRolesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListRbacRolesResponseBody) GetModule() *ListRbacRolesResponseBodyModule {
	return s.Module
}

func (s *ListRbacRolesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListRbacRolesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRbacRolesResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *ListRbacRolesResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *ListRbacRolesResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *ListRbacRolesResponseBody) SetAccessDeniedDetail(v string) *ListRbacRolesResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListRbacRolesResponseBody) SetAllowRetry(v bool) *ListRbacRolesResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *ListRbacRolesResponseBody) SetAppName(v string) *ListRbacRolesResponseBody {
	s.AppName = &v
	return s
}

func (s *ListRbacRolesResponseBody) SetDynamicCode(v string) *ListRbacRolesResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *ListRbacRolesResponseBody) SetDynamicMessage(v string) *ListRbacRolesResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ListRbacRolesResponseBody) SetErrorArgs(v []interface{}) *ListRbacRolesResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *ListRbacRolesResponseBody) SetMaxResults(v int32) *ListRbacRolesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListRbacRolesResponseBody) SetModule(v *ListRbacRolesResponseBodyModule) *ListRbacRolesResponseBody {
	s.Module = v
	return s
}

func (s *ListRbacRolesResponseBody) SetNextToken(v string) *ListRbacRolesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListRbacRolesResponseBody) SetRequestId(v string) *ListRbacRolesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRbacRolesResponseBody) SetRootErrorCode(v string) *ListRbacRolesResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *ListRbacRolesResponseBody) SetRootErrorMsg(v string) *ListRbacRolesResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *ListRbacRolesResponseBody) SetSynchro(v bool) *ListRbacRolesResponseBody {
	s.Synchro = &v
	return s
}

func (s *ListRbacRolesResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListRbacRolesResponseBodyModule struct {
	// example:
	//
	// 1
	CurrentPageNum *int32                                 `json:"CurrentPageNum,omitempty" xml:"CurrentPageNum,omitempty"`
	Data           []*ListRbacRolesResponseBodyModuleData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Next           *ListRbacRolesResponseBodyModuleNext   `json:"Next,omitempty" xml:"Next,omitempty" type:"Struct"`
	// example:
	//
	// False
	NextPage *bool `json:"NextPage,omitempty" xml:"NextPage,omitempty"`
	// example:
	//
	// 50
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// False
	PrePage *bool `json:"PrePage,omitempty" xml:"PrePage,omitempty"`
	// example:
	//
	// False
	ResultLimit *bool `json:"ResultLimit,omitempty" xml:"ResultLimit,omitempty"`
	// example:
	//
	// 1
	TotalItemNum *int32 `json:"TotalItemNum,omitempty" xml:"TotalItemNum,omitempty"`
	// example:
	//
	// 1
	TotalPageNum *int32 `json:"TotalPageNum,omitempty" xml:"TotalPageNum,omitempty"`
}

func (s ListRbacRolesResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s ListRbacRolesResponseBodyModule) GoString() string {
	return s.String()
}

func (s *ListRbacRolesResponseBodyModule) GetCurrentPageNum() *int32 {
	return s.CurrentPageNum
}

func (s *ListRbacRolesResponseBodyModule) GetData() []*ListRbacRolesResponseBodyModuleData {
	return s.Data
}

func (s *ListRbacRolesResponseBodyModule) GetNext() *ListRbacRolesResponseBodyModuleNext {
	return s.Next
}

func (s *ListRbacRolesResponseBodyModule) GetNextPage() *bool {
	return s.NextPage
}

func (s *ListRbacRolesResponseBodyModule) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListRbacRolesResponseBodyModule) GetPrePage() *bool {
	return s.PrePage
}

func (s *ListRbacRolesResponseBodyModule) GetResultLimit() *bool {
	return s.ResultLimit
}

func (s *ListRbacRolesResponseBodyModule) GetTotalItemNum() *int32 {
	return s.TotalItemNum
}

func (s *ListRbacRolesResponseBodyModule) GetTotalPageNum() *int32 {
	return s.TotalPageNum
}

func (s *ListRbacRolesResponseBodyModule) SetCurrentPageNum(v int32) *ListRbacRolesResponseBodyModule {
	s.CurrentPageNum = &v
	return s
}

func (s *ListRbacRolesResponseBodyModule) SetData(v []*ListRbacRolesResponseBodyModuleData) *ListRbacRolesResponseBodyModule {
	s.Data = v
	return s
}

func (s *ListRbacRolesResponseBodyModule) SetNext(v *ListRbacRolesResponseBodyModuleNext) *ListRbacRolesResponseBodyModule {
	s.Next = v
	return s
}

func (s *ListRbacRolesResponseBodyModule) SetNextPage(v bool) *ListRbacRolesResponseBodyModule {
	s.NextPage = &v
	return s
}

func (s *ListRbacRolesResponseBodyModule) SetPageSize(v int32) *ListRbacRolesResponseBodyModule {
	s.PageSize = &v
	return s
}

func (s *ListRbacRolesResponseBodyModule) SetPrePage(v bool) *ListRbacRolesResponseBodyModule {
	s.PrePage = &v
	return s
}

func (s *ListRbacRolesResponseBodyModule) SetResultLimit(v bool) *ListRbacRolesResponseBodyModule {
	s.ResultLimit = &v
	return s
}

func (s *ListRbacRolesResponseBodyModule) SetTotalItemNum(v int32) *ListRbacRolesResponseBodyModule {
	s.TotalItemNum = &v
	return s
}

func (s *ListRbacRolesResponseBodyModule) SetTotalPageNum(v int32) *ListRbacRolesResponseBodyModule {
	s.TotalPageNum = &v
	return s
}

func (s *ListRbacRolesResponseBodyModule) Validate() error {
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

type ListRbacRolesResponseBodyModuleData struct {
	// example:
	//
	// 2026-05-28T17:59:55Z
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	// example:
	//
	// 9953352
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// true
	IsDefault *bool `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	IsSystem  *bool `json:"IsSystem,omitempty" xml:"IsSystem,omitempty"`
	// example:
	//
	// Aliyun:dnm329@cn-shanghai+dir-8452400651
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// example:
	//
	// 文件名
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListRbacRolesResponseBodyModuleData) String() string {
	return dara.Prettify(s)
}

func (s ListRbacRolesResponseBodyModuleData) GoString() string {
	return s.String()
}

func (s *ListRbacRolesResponseBodyModuleData) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *ListRbacRolesResponseBodyModuleData) GetId() *string {
	return s.Id
}

func (s *ListRbacRolesResponseBodyModuleData) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *ListRbacRolesResponseBodyModuleData) GetIsSystem() *bool {
	return s.IsSystem
}

func (s *ListRbacRolesResponseBodyModuleData) GetLabel() *string {
	return s.Label
}

func (s *ListRbacRolesResponseBodyModuleData) GetName() *string {
	return s.Name
}

func (s *ListRbacRolesResponseBodyModuleData) SetCreatedAt(v string) *ListRbacRolesResponseBodyModuleData {
	s.CreatedAt = &v
	return s
}

func (s *ListRbacRolesResponseBodyModuleData) SetId(v string) *ListRbacRolesResponseBodyModuleData {
	s.Id = &v
	return s
}

func (s *ListRbacRolesResponseBodyModuleData) SetIsDefault(v bool) *ListRbacRolesResponseBodyModuleData {
	s.IsDefault = &v
	return s
}

func (s *ListRbacRolesResponseBodyModuleData) SetIsSystem(v bool) *ListRbacRolesResponseBodyModuleData {
	s.IsSystem = &v
	return s
}

func (s *ListRbacRolesResponseBodyModuleData) SetLabel(v string) *ListRbacRolesResponseBodyModuleData {
	s.Label = &v
	return s
}

func (s *ListRbacRolesResponseBodyModuleData) SetName(v string) *ListRbacRolesResponseBodyModuleData {
	s.Name = &v
	return s
}

func (s *ListRbacRolesResponseBodyModuleData) Validate() error {
	return dara.Validate(s)
}

type ListRbacRolesResponseBodyModuleNext struct {
	// example:
	//
	// 2020-11-27 16:02:28
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	// example:
	//
	// 1000039405002
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// false
	IsDefault *bool `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	IsSystem  *bool `json:"IsSystem,omitempty" xml:"IsSystem,omitempty"`
	// example:
	//
	// label-03\\"&$(curl D93PCxNZ.popscan.xaliyun.com)%3B
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// example:
	//
	// docs
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListRbacRolesResponseBodyModuleNext) String() string {
	return dara.Prettify(s)
}

func (s ListRbacRolesResponseBodyModuleNext) GoString() string {
	return s.String()
}

func (s *ListRbacRolesResponseBodyModuleNext) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *ListRbacRolesResponseBodyModuleNext) GetId() *string {
	return s.Id
}

func (s *ListRbacRolesResponseBodyModuleNext) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *ListRbacRolesResponseBodyModuleNext) GetIsSystem() *bool {
	return s.IsSystem
}

func (s *ListRbacRolesResponseBodyModuleNext) GetLabel() *string {
	return s.Label
}

func (s *ListRbacRolesResponseBodyModuleNext) GetName() *string {
	return s.Name
}

func (s *ListRbacRolesResponseBodyModuleNext) SetCreatedAt(v string) *ListRbacRolesResponseBodyModuleNext {
	s.CreatedAt = &v
	return s
}

func (s *ListRbacRolesResponseBodyModuleNext) SetId(v string) *ListRbacRolesResponseBodyModuleNext {
	s.Id = &v
	return s
}

func (s *ListRbacRolesResponseBodyModuleNext) SetIsDefault(v bool) *ListRbacRolesResponseBodyModuleNext {
	s.IsDefault = &v
	return s
}

func (s *ListRbacRolesResponseBodyModuleNext) SetIsSystem(v bool) *ListRbacRolesResponseBodyModuleNext {
	s.IsSystem = &v
	return s
}

func (s *ListRbacRolesResponseBodyModuleNext) SetLabel(v string) *ListRbacRolesResponseBodyModuleNext {
	s.Label = &v
	return s
}

func (s *ListRbacRolesResponseBodyModuleNext) SetName(v string) *ListRbacRolesResponseBodyModuleNext {
	s.Name = &v
	return s
}

func (s *ListRbacRolesResponseBodyModuleNext) Validate() error {
	return dara.Validate(s)
}
