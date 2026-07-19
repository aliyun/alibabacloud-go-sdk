// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRbacPermissionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ListRbacPermissionsResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *ListRbacPermissionsResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *ListRbacPermissionsResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *ListRbacPermissionsResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *ListRbacPermissionsResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *ListRbacPermissionsResponseBody
	GetErrorArgs() []interface{}
	SetMaxResults(v int32) *ListRbacPermissionsResponseBody
	GetMaxResults() *int32
	SetModule(v *ListRbacPermissionsResponseBodyModule) *ListRbacPermissionsResponseBody
	GetModule() *ListRbacPermissionsResponseBodyModule
	SetNextToken(v string) *ListRbacPermissionsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListRbacPermissionsResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *ListRbacPermissionsResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *ListRbacPermissionsResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *ListRbacPermissionsResponseBody
	GetSynchro() *bool
}

type ListRbacPermissionsResponseBody struct {
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
	// spring-cloud-b
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
	MaxResults *int32                                 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	Module     *ListRbacPermissionsResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
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
	// SYSTEM.EROR
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

func (s ListRbacPermissionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRbacPermissionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListRbacPermissionsResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ListRbacPermissionsResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *ListRbacPermissionsResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *ListRbacPermissionsResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *ListRbacPermissionsResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *ListRbacPermissionsResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *ListRbacPermissionsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListRbacPermissionsResponseBody) GetModule() *ListRbacPermissionsResponseBodyModule {
	return s.Module
}

func (s *ListRbacPermissionsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListRbacPermissionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRbacPermissionsResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *ListRbacPermissionsResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *ListRbacPermissionsResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *ListRbacPermissionsResponseBody) SetAccessDeniedDetail(v string) *ListRbacPermissionsResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListRbacPermissionsResponseBody) SetAllowRetry(v bool) *ListRbacPermissionsResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *ListRbacPermissionsResponseBody) SetAppName(v string) *ListRbacPermissionsResponseBody {
	s.AppName = &v
	return s
}

func (s *ListRbacPermissionsResponseBody) SetDynamicCode(v string) *ListRbacPermissionsResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *ListRbacPermissionsResponseBody) SetDynamicMessage(v string) *ListRbacPermissionsResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ListRbacPermissionsResponseBody) SetErrorArgs(v []interface{}) *ListRbacPermissionsResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *ListRbacPermissionsResponseBody) SetMaxResults(v int32) *ListRbacPermissionsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListRbacPermissionsResponseBody) SetModule(v *ListRbacPermissionsResponseBodyModule) *ListRbacPermissionsResponseBody {
	s.Module = v
	return s
}

func (s *ListRbacPermissionsResponseBody) SetNextToken(v string) *ListRbacPermissionsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListRbacPermissionsResponseBody) SetRequestId(v string) *ListRbacPermissionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRbacPermissionsResponseBody) SetRootErrorCode(v string) *ListRbacPermissionsResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *ListRbacPermissionsResponseBody) SetRootErrorMsg(v string) *ListRbacPermissionsResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *ListRbacPermissionsResponseBody) SetSynchro(v bool) *ListRbacPermissionsResponseBody {
	s.Synchro = &v
	return s
}

func (s *ListRbacPermissionsResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListRbacPermissionsResponseBodyModule struct {
	// example:
	//
	// 1
	CurrentPageNum *int32                                       `json:"CurrentPageNum,omitempty" xml:"CurrentPageNum,omitempty"`
	Data           []*ListRbacPermissionsResponseBodyModuleData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Next           *ListRbacPermissionsResponseBodyModuleNext   `json:"Next,omitempty" xml:"Next,omitempty" type:"Struct"`
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

func (s ListRbacPermissionsResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s ListRbacPermissionsResponseBodyModule) GoString() string {
	return s.String()
}

func (s *ListRbacPermissionsResponseBodyModule) GetCurrentPageNum() *int32 {
	return s.CurrentPageNum
}

func (s *ListRbacPermissionsResponseBodyModule) GetData() []*ListRbacPermissionsResponseBodyModuleData {
	return s.Data
}

func (s *ListRbacPermissionsResponseBodyModule) GetNext() *ListRbacPermissionsResponseBodyModuleNext {
	return s.Next
}

func (s *ListRbacPermissionsResponseBodyModule) GetNextPage() *bool {
	return s.NextPage
}

func (s *ListRbacPermissionsResponseBodyModule) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListRbacPermissionsResponseBodyModule) GetPrePage() *bool {
	return s.PrePage
}

func (s *ListRbacPermissionsResponseBodyModule) GetResultLimit() *bool {
	return s.ResultLimit
}

func (s *ListRbacPermissionsResponseBodyModule) GetTotalItemNum() *int32 {
	return s.TotalItemNum
}

func (s *ListRbacPermissionsResponseBodyModule) GetTotalPageNum() *int32 {
	return s.TotalPageNum
}

func (s *ListRbacPermissionsResponseBodyModule) SetCurrentPageNum(v int32) *ListRbacPermissionsResponseBodyModule {
	s.CurrentPageNum = &v
	return s
}

func (s *ListRbacPermissionsResponseBodyModule) SetData(v []*ListRbacPermissionsResponseBodyModuleData) *ListRbacPermissionsResponseBodyModule {
	s.Data = v
	return s
}

func (s *ListRbacPermissionsResponseBodyModule) SetNext(v *ListRbacPermissionsResponseBodyModuleNext) *ListRbacPermissionsResponseBodyModule {
	s.Next = v
	return s
}

func (s *ListRbacPermissionsResponseBodyModule) SetNextPage(v bool) *ListRbacPermissionsResponseBodyModule {
	s.NextPage = &v
	return s
}

func (s *ListRbacPermissionsResponseBodyModule) SetPageSize(v int32) *ListRbacPermissionsResponseBodyModule {
	s.PageSize = &v
	return s
}

func (s *ListRbacPermissionsResponseBodyModule) SetPrePage(v bool) *ListRbacPermissionsResponseBodyModule {
	s.PrePage = &v
	return s
}

func (s *ListRbacPermissionsResponseBodyModule) SetResultLimit(v bool) *ListRbacPermissionsResponseBodyModule {
	s.ResultLimit = &v
	return s
}

func (s *ListRbacPermissionsResponseBodyModule) SetTotalItemNum(v int32) *ListRbacPermissionsResponseBodyModule {
	s.TotalItemNum = &v
	return s
}

func (s *ListRbacPermissionsResponseBodyModule) SetTotalPageNum(v int32) *ListRbacPermissionsResponseBodyModule {
	s.TotalPageNum = &v
	return s
}

func (s *ListRbacPermissionsResponseBodyModule) Validate() error {
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

type ListRbacPermissionsResponseBodyModuleData struct {
	// example:
	//
	// monitor
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// example:
	//
	// lshm-mysql-coypt
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 9953352
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// sjapi-h5.aihuishou.com-waf
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
}

func (s ListRbacPermissionsResponseBodyModuleData) String() string {
	return dara.Prettify(s)
}

func (s ListRbacPermissionsResponseBodyModuleData) GoString() string {
	return s.String()
}

func (s *ListRbacPermissionsResponseBodyModuleData) GetAction() *string {
	return s.Action
}

func (s *ListRbacPermissionsResponseBodyModuleData) GetDescription() *string {
	return s.Description
}

func (s *ListRbacPermissionsResponseBodyModuleData) GetId() *string {
	return s.Id
}

func (s *ListRbacPermissionsResponseBodyModuleData) GetResource() *string {
	return s.Resource
}

func (s *ListRbacPermissionsResponseBodyModuleData) SetAction(v string) *ListRbacPermissionsResponseBodyModuleData {
	s.Action = &v
	return s
}

func (s *ListRbacPermissionsResponseBodyModuleData) SetDescription(v string) *ListRbacPermissionsResponseBodyModuleData {
	s.Description = &v
	return s
}

func (s *ListRbacPermissionsResponseBodyModuleData) SetId(v string) *ListRbacPermissionsResponseBodyModuleData {
	s.Id = &v
	return s
}

func (s *ListRbacPermissionsResponseBodyModuleData) SetResource(v string) *ListRbacPermissionsResponseBodyModuleData {
	s.Resource = &v
	return s
}

func (s *ListRbacPermissionsResponseBodyModuleData) Validate() error {
	return dara.Validate(s)
}

type ListRbacPermissionsResponseBodyModuleNext struct {
	// example:
	//
	// accept
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// example:
	//
	// lx_supabase_test_02
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 1000039405002
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// buy
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
}

func (s ListRbacPermissionsResponseBodyModuleNext) String() string {
	return dara.Prettify(s)
}

func (s ListRbacPermissionsResponseBodyModuleNext) GoString() string {
	return s.String()
}

func (s *ListRbacPermissionsResponseBodyModuleNext) GetAction() *string {
	return s.Action
}

func (s *ListRbacPermissionsResponseBodyModuleNext) GetDescription() *string {
	return s.Description
}

func (s *ListRbacPermissionsResponseBodyModuleNext) GetId() *string {
	return s.Id
}

func (s *ListRbacPermissionsResponseBodyModuleNext) GetResource() *string {
	return s.Resource
}

func (s *ListRbacPermissionsResponseBodyModuleNext) SetAction(v string) *ListRbacPermissionsResponseBodyModuleNext {
	s.Action = &v
	return s
}

func (s *ListRbacPermissionsResponseBodyModuleNext) SetDescription(v string) *ListRbacPermissionsResponseBodyModuleNext {
	s.Description = &v
	return s
}

func (s *ListRbacPermissionsResponseBodyModuleNext) SetId(v string) *ListRbacPermissionsResponseBodyModuleNext {
	s.Id = &v
	return s
}

func (s *ListRbacPermissionsResponseBodyModuleNext) SetResource(v string) *ListRbacPermissionsResponseBodyModuleNext {
	s.Resource = &v
	return s
}

func (s *ListRbacPermissionsResponseBodyModuleNext) Validate() error {
	return dara.Validate(s)
}
