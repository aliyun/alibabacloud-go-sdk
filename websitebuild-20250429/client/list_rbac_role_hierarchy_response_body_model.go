// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRbacRoleHierarchyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ListRbacRoleHierarchyResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *ListRbacRoleHierarchyResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *ListRbacRoleHierarchyResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *ListRbacRoleHierarchyResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *ListRbacRoleHierarchyResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *ListRbacRoleHierarchyResponseBody
	GetErrorArgs() []interface{}
	SetMaxResults(v int32) *ListRbacRoleHierarchyResponseBody
	GetMaxResults() *int32
	SetModule(v *ListRbacRoleHierarchyResponseBodyModule) *ListRbacRoleHierarchyResponseBody
	GetModule() *ListRbacRoleHierarchyResponseBodyModule
	SetNextToken(v string) *ListRbacRoleHierarchyResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListRbacRoleHierarchyResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *ListRbacRoleHierarchyResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *ListRbacRoleHierarchyResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *ListRbacRoleHierarchyResponseBody
	GetSynchro() *bool
}

type ListRbacRoleHierarchyResponseBody struct {
	// The access denied details.
	//
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// Indicates whether a retry is allowed. Valid values:
	//
	// - false: No retry is allowed.
	//
	// - true: A retry is allowed.
	//
	// example:
	//
	// False
	AllowRetry *bool `json:"AllowRetry,omitempty" xml:"AllowRetry,omitempty"`
	// The application name.
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
	// The dynamic error message, which is used to replace the `%s` variable in the **ErrMessage*	- response parameter.
	//
	// > If **ErrMessage*	- returns **The Value of Input Parameter %s is not valid*	- and **DynamicMessage*	- returns **DtsJobId**, the value of the request parameter **DtsJobId*	- is invalid.
	//
	// example:
	//
	// SYSTEM_ERROR
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The error parameters.
	ErrorArgs []interface{} `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	// The number of entries per query.
	//
	// Valid values: 10 to 100. Default value: 20.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The response object.
	Module *ListRbacRoleHierarchyResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
	// The token for the next query. This parameter is empty if no more results exist.
	//
	// example:
	//
	// 0l45bkwM022Dt+rOvPi/oQ==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
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
	// 系统异常
	RootErrorMsg *string `json:"RootErrorMsg,omitempty" xml:"RootErrorMsg,omitempty"`
	// Indicates whether the request is synchronously processed.
	//
	// example:
	//
	// True
	Synchro *bool `json:"Synchro,omitempty" xml:"Synchro,omitempty"`
}

func (s ListRbacRoleHierarchyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRbacRoleHierarchyResponseBody) GoString() string {
	return s.String()
}

func (s *ListRbacRoleHierarchyResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ListRbacRoleHierarchyResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *ListRbacRoleHierarchyResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *ListRbacRoleHierarchyResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *ListRbacRoleHierarchyResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *ListRbacRoleHierarchyResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *ListRbacRoleHierarchyResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListRbacRoleHierarchyResponseBody) GetModule() *ListRbacRoleHierarchyResponseBodyModule {
	return s.Module
}

func (s *ListRbacRoleHierarchyResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListRbacRoleHierarchyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRbacRoleHierarchyResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *ListRbacRoleHierarchyResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *ListRbacRoleHierarchyResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *ListRbacRoleHierarchyResponseBody) SetAccessDeniedDetail(v string) *ListRbacRoleHierarchyResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListRbacRoleHierarchyResponseBody) SetAllowRetry(v bool) *ListRbacRoleHierarchyResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *ListRbacRoleHierarchyResponseBody) SetAppName(v string) *ListRbacRoleHierarchyResponseBody {
	s.AppName = &v
	return s
}

func (s *ListRbacRoleHierarchyResponseBody) SetDynamicCode(v string) *ListRbacRoleHierarchyResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *ListRbacRoleHierarchyResponseBody) SetDynamicMessage(v string) *ListRbacRoleHierarchyResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ListRbacRoleHierarchyResponseBody) SetErrorArgs(v []interface{}) *ListRbacRoleHierarchyResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *ListRbacRoleHierarchyResponseBody) SetMaxResults(v int32) *ListRbacRoleHierarchyResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListRbacRoleHierarchyResponseBody) SetModule(v *ListRbacRoleHierarchyResponseBodyModule) *ListRbacRoleHierarchyResponseBody {
	s.Module = v
	return s
}

func (s *ListRbacRoleHierarchyResponseBody) SetNextToken(v string) *ListRbacRoleHierarchyResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListRbacRoleHierarchyResponseBody) SetRequestId(v string) *ListRbacRoleHierarchyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRbacRoleHierarchyResponseBody) SetRootErrorCode(v string) *ListRbacRoleHierarchyResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *ListRbacRoleHierarchyResponseBody) SetRootErrorMsg(v string) *ListRbacRoleHierarchyResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *ListRbacRoleHierarchyResponseBody) SetSynchro(v bool) *ListRbacRoleHierarchyResponseBody {
	s.Synchro = &v
	return s
}

func (s *ListRbacRoleHierarchyResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListRbacRoleHierarchyResponseBodyModule struct {
	// The current page number.
	//
	// example:
	//
	// 1
	CurrentPageNum *int32 `json:"CurrentPageNum,omitempty" xml:"CurrentPageNum,omitempty"`
	// The request results.
	Data []*ListRbacRoleHierarchyResponseBodyModuleData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The decision weight.
	Next *ListRbacRoleHierarchyResponseBodyModuleNext `json:"Next,omitempty" xml:"Next,omitempty" type:"Struct"`
	// Indicates whether a next page exists.
	//
	// example:
	//
	// False
	NextPage *bool `json:"NextPage,omitempty" xml:"NextPage,omitempty"`
	// The page size.
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
	// The server processes up to 1,000 most recent records beyond the pagination limit. If the results exceed 1,000 records, **ResultLimit*	- is **true**. In this case, narrow the time range and search again. Otherwise, **ResultLimit*	- is **false**.
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

func (s ListRbacRoleHierarchyResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s ListRbacRoleHierarchyResponseBodyModule) GoString() string {
	return s.String()
}

func (s *ListRbacRoleHierarchyResponseBodyModule) GetCurrentPageNum() *int32 {
	return s.CurrentPageNum
}

func (s *ListRbacRoleHierarchyResponseBodyModule) GetData() []*ListRbacRoleHierarchyResponseBodyModuleData {
	return s.Data
}

func (s *ListRbacRoleHierarchyResponseBodyModule) GetNext() *ListRbacRoleHierarchyResponseBodyModuleNext {
	return s.Next
}

func (s *ListRbacRoleHierarchyResponseBodyModule) GetNextPage() *bool {
	return s.NextPage
}

func (s *ListRbacRoleHierarchyResponseBodyModule) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListRbacRoleHierarchyResponseBodyModule) GetPrePage() *bool {
	return s.PrePage
}

func (s *ListRbacRoleHierarchyResponseBodyModule) GetResultLimit() *bool {
	return s.ResultLimit
}

func (s *ListRbacRoleHierarchyResponseBodyModule) GetTotalItemNum() *int32 {
	return s.TotalItemNum
}

func (s *ListRbacRoleHierarchyResponseBodyModule) GetTotalPageNum() *int32 {
	return s.TotalPageNum
}

func (s *ListRbacRoleHierarchyResponseBodyModule) SetCurrentPageNum(v int32) *ListRbacRoleHierarchyResponseBodyModule {
	s.CurrentPageNum = &v
	return s
}

func (s *ListRbacRoleHierarchyResponseBodyModule) SetData(v []*ListRbacRoleHierarchyResponseBodyModuleData) *ListRbacRoleHierarchyResponseBodyModule {
	s.Data = v
	return s
}

func (s *ListRbacRoleHierarchyResponseBodyModule) SetNext(v *ListRbacRoleHierarchyResponseBodyModuleNext) *ListRbacRoleHierarchyResponseBodyModule {
	s.Next = v
	return s
}

func (s *ListRbacRoleHierarchyResponseBodyModule) SetNextPage(v bool) *ListRbacRoleHierarchyResponseBodyModule {
	s.NextPage = &v
	return s
}

func (s *ListRbacRoleHierarchyResponseBodyModule) SetPageSize(v int32) *ListRbacRoleHierarchyResponseBodyModule {
	s.PageSize = &v
	return s
}

func (s *ListRbacRoleHierarchyResponseBodyModule) SetPrePage(v bool) *ListRbacRoleHierarchyResponseBodyModule {
	s.PrePage = &v
	return s
}

func (s *ListRbacRoleHierarchyResponseBodyModule) SetResultLimit(v bool) *ListRbacRoleHierarchyResponseBodyModule {
	s.ResultLimit = &v
	return s
}

func (s *ListRbacRoleHierarchyResponseBodyModule) SetTotalItemNum(v int32) *ListRbacRoleHierarchyResponseBodyModule {
	s.TotalItemNum = &v
	return s
}

func (s *ListRbacRoleHierarchyResponseBodyModule) SetTotalPageNum(v int32) *ListRbacRoleHierarchyResponseBodyModule {
	s.TotalPageNum = &v
	return s
}

func (s *ListRbacRoleHierarchyResponseBodyModule) Validate() error {
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

type ListRbacRoleHierarchyResponseBodyModuleData struct {
	// The child role ID.
	//
	// example:
	//
	// fc94cc51-310f-4485-adb2-ed8c706aff3b
	ChildRoleId *string `json:"ChildRoleId,omitempty" xml:"ChildRoleId,omitempty"`
	// The parent role ID.
	//
	// example:
	//
	// 71e07711-9a17-49f4-9f83-387a60ee5b64
	ParentRoleId *string `json:"ParentRoleId,omitempty" xml:"ParentRoleId,omitempty"`
}

func (s ListRbacRoleHierarchyResponseBodyModuleData) String() string {
	return dara.Prettify(s)
}

func (s ListRbacRoleHierarchyResponseBodyModuleData) GoString() string {
	return s.String()
}

func (s *ListRbacRoleHierarchyResponseBodyModuleData) GetChildRoleId() *string {
	return s.ChildRoleId
}

func (s *ListRbacRoleHierarchyResponseBodyModuleData) GetParentRoleId() *string {
	return s.ParentRoleId
}

func (s *ListRbacRoleHierarchyResponseBodyModuleData) SetChildRoleId(v string) *ListRbacRoleHierarchyResponseBodyModuleData {
	s.ChildRoleId = &v
	return s
}

func (s *ListRbacRoleHierarchyResponseBodyModuleData) SetParentRoleId(v string) *ListRbacRoleHierarchyResponseBodyModuleData {
	s.ParentRoleId = &v
	return s
}

func (s *ListRbacRoleHierarchyResponseBodyModuleData) Validate() error {
	return dara.Validate(s)
}

type ListRbacRoleHierarchyResponseBodyModuleNext struct {
	// The child role ID.
	//
	// example:
	//
	// fc94cc51-310f-4485-adb2-ed8c706aff3b
	ChildRoleId *string `json:"ChildRoleId,omitempty" xml:"ChildRoleId,omitempty"`
	// The parent role ID.
	//
	// example:
	//
	// 71e07711-9a17-49f4-9f83-387a60ee5b64
	ParentRoleId *string `json:"ParentRoleId,omitempty" xml:"ParentRoleId,omitempty"`
}

func (s ListRbacRoleHierarchyResponseBodyModuleNext) String() string {
	return dara.Prettify(s)
}

func (s ListRbacRoleHierarchyResponseBodyModuleNext) GoString() string {
	return s.String()
}

func (s *ListRbacRoleHierarchyResponseBodyModuleNext) GetChildRoleId() *string {
	return s.ChildRoleId
}

func (s *ListRbacRoleHierarchyResponseBodyModuleNext) GetParentRoleId() *string {
	return s.ParentRoleId
}

func (s *ListRbacRoleHierarchyResponseBodyModuleNext) SetChildRoleId(v string) *ListRbacRoleHierarchyResponseBodyModuleNext {
	s.ChildRoleId = &v
	return s
}

func (s *ListRbacRoleHierarchyResponseBodyModuleNext) SetParentRoleId(v string) *ListRbacRoleHierarchyResponseBodyModuleNext {
	s.ParentRoleId = &v
	return s
}

func (s *ListRbacRoleHierarchyResponseBodyModuleNext) Validate() error {
	return dara.Validate(s)
}
