// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRbacOrgTreeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ListRbacOrgTreeResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *ListRbacOrgTreeResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *ListRbacOrgTreeResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *ListRbacOrgTreeResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *ListRbacOrgTreeResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *ListRbacOrgTreeResponseBody
	GetErrorArgs() []interface{}
	SetMaxResults(v int32) *ListRbacOrgTreeResponseBody
	GetMaxResults() *int32
	SetModule(v *ListRbacOrgTreeResponseBodyModule) *ListRbacOrgTreeResponseBody
	GetModule() *ListRbacOrgTreeResponseBodyModule
	SetNextToken(v string) *ListRbacOrgTreeResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListRbacOrgTreeResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *ListRbacOrgTreeResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *ListRbacOrgTreeResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *ListRbacOrgTreeResponseBody
	GetSynchro() *bool
}

type ListRbacOrgTreeResponseBody struct {
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
	MaxResults *int32                             `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	Module     *ListRbacOrgTreeResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
	// example:
	//
	// AAAAARbaCuN6hiD08qrLdwJ9Fh3BFw8paIJ7ylB6A7Qn9JjM
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

func (s ListRbacOrgTreeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRbacOrgTreeResponseBody) GoString() string {
	return s.String()
}

func (s *ListRbacOrgTreeResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ListRbacOrgTreeResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *ListRbacOrgTreeResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *ListRbacOrgTreeResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *ListRbacOrgTreeResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *ListRbacOrgTreeResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *ListRbacOrgTreeResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListRbacOrgTreeResponseBody) GetModule() *ListRbacOrgTreeResponseBodyModule {
	return s.Module
}

func (s *ListRbacOrgTreeResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListRbacOrgTreeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRbacOrgTreeResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *ListRbacOrgTreeResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *ListRbacOrgTreeResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *ListRbacOrgTreeResponseBody) SetAccessDeniedDetail(v string) *ListRbacOrgTreeResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListRbacOrgTreeResponseBody) SetAllowRetry(v bool) *ListRbacOrgTreeResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *ListRbacOrgTreeResponseBody) SetAppName(v string) *ListRbacOrgTreeResponseBody {
	s.AppName = &v
	return s
}

func (s *ListRbacOrgTreeResponseBody) SetDynamicCode(v string) *ListRbacOrgTreeResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *ListRbacOrgTreeResponseBody) SetDynamicMessage(v string) *ListRbacOrgTreeResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ListRbacOrgTreeResponseBody) SetErrorArgs(v []interface{}) *ListRbacOrgTreeResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *ListRbacOrgTreeResponseBody) SetMaxResults(v int32) *ListRbacOrgTreeResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListRbacOrgTreeResponseBody) SetModule(v *ListRbacOrgTreeResponseBodyModule) *ListRbacOrgTreeResponseBody {
	s.Module = v
	return s
}

func (s *ListRbacOrgTreeResponseBody) SetNextToken(v string) *ListRbacOrgTreeResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListRbacOrgTreeResponseBody) SetRequestId(v string) *ListRbacOrgTreeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRbacOrgTreeResponseBody) SetRootErrorCode(v string) *ListRbacOrgTreeResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *ListRbacOrgTreeResponseBody) SetRootErrorMsg(v string) *ListRbacOrgTreeResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *ListRbacOrgTreeResponseBody) SetSynchro(v bool) *ListRbacOrgTreeResponseBody {
	s.Synchro = &v
	return s
}

func (s *ListRbacOrgTreeResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListRbacOrgTreeResponseBodyModule struct {
	// example:
	//
	// 12
	CurrentPageNum *int32                                   `json:"CurrentPageNum,omitempty" xml:"CurrentPageNum,omitempty"`
	Data           []*ListRbacOrgTreeResponseBodyModuleData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Next           *ListRbacOrgTreeResponseBodyModuleNext   `json:"Next,omitempty" xml:"Next,omitempty" type:"Struct"`
	// example:
	//
	// False
	NextPage *bool `json:"NextPage,omitempty" xml:"NextPage,omitempty"`
	// example:
	//
	// 10
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

func (s ListRbacOrgTreeResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s ListRbacOrgTreeResponseBodyModule) GoString() string {
	return s.String()
}

func (s *ListRbacOrgTreeResponseBodyModule) GetCurrentPageNum() *int32 {
	return s.CurrentPageNum
}

func (s *ListRbacOrgTreeResponseBodyModule) GetData() []*ListRbacOrgTreeResponseBodyModuleData {
	return s.Data
}

func (s *ListRbacOrgTreeResponseBodyModule) GetNext() *ListRbacOrgTreeResponseBodyModuleNext {
	return s.Next
}

func (s *ListRbacOrgTreeResponseBodyModule) GetNextPage() *bool {
	return s.NextPage
}

func (s *ListRbacOrgTreeResponseBodyModule) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListRbacOrgTreeResponseBodyModule) GetPrePage() *bool {
	return s.PrePage
}

func (s *ListRbacOrgTreeResponseBodyModule) GetResultLimit() *bool {
	return s.ResultLimit
}

func (s *ListRbacOrgTreeResponseBodyModule) GetTotalItemNum() *int32 {
	return s.TotalItemNum
}

func (s *ListRbacOrgTreeResponseBodyModule) GetTotalPageNum() *int32 {
	return s.TotalPageNum
}

func (s *ListRbacOrgTreeResponseBodyModule) SetCurrentPageNum(v int32) *ListRbacOrgTreeResponseBodyModule {
	s.CurrentPageNum = &v
	return s
}

func (s *ListRbacOrgTreeResponseBodyModule) SetData(v []*ListRbacOrgTreeResponseBodyModuleData) *ListRbacOrgTreeResponseBodyModule {
	s.Data = v
	return s
}

func (s *ListRbacOrgTreeResponseBodyModule) SetNext(v *ListRbacOrgTreeResponseBodyModuleNext) *ListRbacOrgTreeResponseBodyModule {
	s.Next = v
	return s
}

func (s *ListRbacOrgTreeResponseBodyModule) SetNextPage(v bool) *ListRbacOrgTreeResponseBodyModule {
	s.NextPage = &v
	return s
}

func (s *ListRbacOrgTreeResponseBodyModule) SetPageSize(v int32) *ListRbacOrgTreeResponseBodyModule {
	s.PageSize = &v
	return s
}

func (s *ListRbacOrgTreeResponseBodyModule) SetPrePage(v bool) *ListRbacOrgTreeResponseBodyModule {
	s.PrePage = &v
	return s
}

func (s *ListRbacOrgTreeResponseBodyModule) SetResultLimit(v bool) *ListRbacOrgTreeResponseBodyModule {
	s.ResultLimit = &v
	return s
}

func (s *ListRbacOrgTreeResponseBodyModule) SetTotalItemNum(v int32) *ListRbacOrgTreeResponseBodyModule {
	s.TotalItemNum = &v
	return s
}

func (s *ListRbacOrgTreeResponseBodyModule) SetTotalPageNum(v int32) *ListRbacOrgTreeResponseBodyModule {
	s.TotalPageNum = &v
	return s
}

func (s *ListRbacOrgTreeResponseBodyModule) Validate() error {
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

type ListRbacOrgTreeResponseBodyModuleData struct {
	// example:
	//
	// 2026-05-08T02:28:26Z
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	// example:
	//
	// FirstChildDepth
	Depth *int32 `json:"Depth,omitempty" xml:"Depth,omitempty"`
	// example:
	//
	// 9953352
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 文件名
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
}

func (s ListRbacOrgTreeResponseBodyModuleData) String() string {
	return dara.Prettify(s)
}

func (s ListRbacOrgTreeResponseBodyModuleData) GoString() string {
	return s.String()
}

func (s *ListRbacOrgTreeResponseBodyModuleData) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *ListRbacOrgTreeResponseBodyModuleData) GetDepth() *int32 {
	return s.Depth
}

func (s *ListRbacOrgTreeResponseBodyModuleData) GetId() *string {
	return s.Id
}

func (s *ListRbacOrgTreeResponseBodyModuleData) GetName() *string {
	return s.Name
}

func (s *ListRbacOrgTreeResponseBodyModuleData) GetPath() *string {
	return s.Path
}

func (s *ListRbacOrgTreeResponseBodyModuleData) SetCreatedAt(v string) *ListRbacOrgTreeResponseBodyModuleData {
	s.CreatedAt = &v
	return s
}

func (s *ListRbacOrgTreeResponseBodyModuleData) SetDepth(v int32) *ListRbacOrgTreeResponseBodyModuleData {
	s.Depth = &v
	return s
}

func (s *ListRbacOrgTreeResponseBodyModuleData) SetId(v string) *ListRbacOrgTreeResponseBodyModuleData {
	s.Id = &v
	return s
}

func (s *ListRbacOrgTreeResponseBodyModuleData) SetName(v string) *ListRbacOrgTreeResponseBodyModuleData {
	s.Name = &v
	return s
}

func (s *ListRbacOrgTreeResponseBodyModuleData) SetPath(v string) *ListRbacOrgTreeResponseBodyModuleData {
	s.Path = &v
	return s
}

func (s *ListRbacOrgTreeResponseBodyModuleData) Validate() error {
	return dara.Validate(s)
}

type ListRbacOrgTreeResponseBodyModuleNext struct {
	// example:
	//
	// 2025-12-15T02:29:22Z
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	// example:
	//
	// 2
	Depth *int32 `json:"Depth,omitempty" xml:"Depth,omitempty"`
	// example:
	//
	// 1000039405002
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// docs
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// /h6sRMdP&LTdQz/44ImvD/JtjSeUza/wVcp1i/dsxnl0FlL0/bvL/d+0~V6Fi5+R/P8ebktwGxe/l&AqBD_/Y+TstdpTo06U_Q/4i4:EDPGo/7.fIVgd.//AOJtXP5/X
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
}

func (s ListRbacOrgTreeResponseBodyModuleNext) String() string {
	return dara.Prettify(s)
}

func (s ListRbacOrgTreeResponseBodyModuleNext) GoString() string {
	return s.String()
}

func (s *ListRbacOrgTreeResponseBodyModuleNext) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *ListRbacOrgTreeResponseBodyModuleNext) GetDepth() *int32 {
	return s.Depth
}

func (s *ListRbacOrgTreeResponseBodyModuleNext) GetId() *string {
	return s.Id
}

func (s *ListRbacOrgTreeResponseBodyModuleNext) GetName() *string {
	return s.Name
}

func (s *ListRbacOrgTreeResponseBodyModuleNext) GetPath() *string {
	return s.Path
}

func (s *ListRbacOrgTreeResponseBodyModuleNext) SetCreatedAt(v string) *ListRbacOrgTreeResponseBodyModuleNext {
	s.CreatedAt = &v
	return s
}

func (s *ListRbacOrgTreeResponseBodyModuleNext) SetDepth(v int32) *ListRbacOrgTreeResponseBodyModuleNext {
	s.Depth = &v
	return s
}

func (s *ListRbacOrgTreeResponseBodyModuleNext) SetId(v string) *ListRbacOrgTreeResponseBodyModuleNext {
	s.Id = &v
	return s
}

func (s *ListRbacOrgTreeResponseBodyModuleNext) SetName(v string) *ListRbacOrgTreeResponseBodyModuleNext {
	s.Name = &v
	return s
}

func (s *ListRbacOrgTreeResponseBodyModuleNext) SetPath(v string) *ListRbacOrgTreeResponseBodyModuleNext {
	s.Path = &v
	return s
}

func (s *ListRbacOrgTreeResponseBodyModuleNext) Validate() error {
	return dara.Validate(s)
}
