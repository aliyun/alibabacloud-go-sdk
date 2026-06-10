// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMaterialTaskListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *QueryMaterialTaskListResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *QueryMaterialTaskListResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *QueryMaterialTaskListResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *QueryMaterialTaskListResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *QueryMaterialTaskListResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *QueryMaterialTaskListResponseBody
	GetErrorArgs() []interface{}
	SetErrorCode(v string) *QueryMaterialTaskListResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *QueryMaterialTaskListResponseBody
	GetErrorMsg() *string
	SetMaxResults(v int32) *QueryMaterialTaskListResponseBody
	GetMaxResults() *int32
	SetModule(v *QueryMaterialTaskListResponseBodyModule) *QueryMaterialTaskListResponseBody
	GetModule() *QueryMaterialTaskListResponseBodyModule
	SetNextToken(v string) *QueryMaterialTaskListResponseBody
	GetNextToken() *string
	SetRequestId(v string) *QueryMaterialTaskListResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *QueryMaterialTaskListResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *QueryMaterialTaskListResponseBody
	GetRootErrorMsg() *string
	SetSuccess(v bool) *QueryMaterialTaskListResponseBody
	GetSuccess() *bool
	SetSynchro(v bool) *QueryMaterialTaskListResponseBody
	GetSynchro() *bool
}

type QueryMaterialTaskListResponseBody struct {
	// Detailed reason for access denial.
	//
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// Indicates whether retry is allowed. Valid values:
	//
	// - false: Retry is not allowed.
	//
	// - true: Retry is allowed.
	//
	// example:
	//
	// False
	AllowRetry *bool `json:"AllowRetry,omitempty" xml:"AllowRetry,omitempty"`
	// App Name.
	//
	// example:
	//
	// or
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// dynamic error code.
	//
	// example:
	//
	// ERROR-oo1
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// dynamic message. Not used currently. Ignore it.
	//
	// example:
	//
	// SYSTEM_ERROR
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// Returned error parameters
	ErrorArgs []interface{} `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	// Error code. The value is interpreted as follows: If the request succeeded, the ErrorCode field is not returned. If the request failed, the ErrorCode field is returned. For specific details, see the error code list in this topic.
	//
	// example:
	//
	// 0
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// error message.
	//
	// example:
	//
	// aliuid:1998006665794443 assumeRole not exist,serviceName:realtimelogpush.dcdnservices.aliyuncs.com
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// Number of results returned per query.
	//
	// Valid values: 10 to 100. Default Value: 20.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// Response data
	Module *QueryMaterialTaskListResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
	// Token for starting the next query. This value is empty if there is no next query.
	//
	// example:
	//
	// AAAAARbaCuN6hiD08qrLdwJ9Fh3BFw8paIJ7ylB6A7Qn9JjM
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 6C6B99AC-39EC-5350-874C-204128C905E6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// error code
	//
	// example:
	//
	// SYSTEM.ERROR
	RootErrorCode *string `json:"RootErrorCode,omitempty" xml:"RootErrorCode,omitempty"`
	// abnormal message
	//
	// example:
	//
	// 系统异常
	RootErrorMsg *string `json:"RootErrorMsg,omitempty" xml:"RootErrorMsg,omitempty"`
	// Indicates whether the Request succeeded.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// Indicates whether processing is synchronous.
	//
	// example:
	//
	// True
	Synchro *bool `json:"Synchro,omitempty" xml:"Synchro,omitempty"`
}

func (s QueryMaterialTaskListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryMaterialTaskListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMaterialTaskListResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *QueryMaterialTaskListResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *QueryMaterialTaskListResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *QueryMaterialTaskListResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *QueryMaterialTaskListResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *QueryMaterialTaskListResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *QueryMaterialTaskListResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *QueryMaterialTaskListResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *QueryMaterialTaskListResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *QueryMaterialTaskListResponseBody) GetModule() *QueryMaterialTaskListResponseBodyModule {
	return s.Module
}

func (s *QueryMaterialTaskListResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *QueryMaterialTaskListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryMaterialTaskListResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *QueryMaterialTaskListResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *QueryMaterialTaskListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryMaterialTaskListResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *QueryMaterialTaskListResponseBody) SetAccessDeniedDetail(v string) *QueryMaterialTaskListResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *QueryMaterialTaskListResponseBody) SetAllowRetry(v bool) *QueryMaterialTaskListResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *QueryMaterialTaskListResponseBody) SetAppName(v string) *QueryMaterialTaskListResponseBody {
	s.AppName = &v
	return s
}

func (s *QueryMaterialTaskListResponseBody) SetDynamicCode(v string) *QueryMaterialTaskListResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *QueryMaterialTaskListResponseBody) SetDynamicMessage(v string) *QueryMaterialTaskListResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *QueryMaterialTaskListResponseBody) SetErrorArgs(v []interface{}) *QueryMaterialTaskListResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *QueryMaterialTaskListResponseBody) SetErrorCode(v string) *QueryMaterialTaskListResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *QueryMaterialTaskListResponseBody) SetErrorMsg(v string) *QueryMaterialTaskListResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *QueryMaterialTaskListResponseBody) SetMaxResults(v int32) *QueryMaterialTaskListResponseBody {
	s.MaxResults = &v
	return s
}

func (s *QueryMaterialTaskListResponseBody) SetModule(v *QueryMaterialTaskListResponseBodyModule) *QueryMaterialTaskListResponseBody {
	s.Module = v
	return s
}

func (s *QueryMaterialTaskListResponseBody) SetNextToken(v string) *QueryMaterialTaskListResponseBody {
	s.NextToken = &v
	return s
}

func (s *QueryMaterialTaskListResponseBody) SetRequestId(v string) *QueryMaterialTaskListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryMaterialTaskListResponseBody) SetRootErrorCode(v string) *QueryMaterialTaskListResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *QueryMaterialTaskListResponseBody) SetRootErrorMsg(v string) *QueryMaterialTaskListResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *QueryMaterialTaskListResponseBody) SetSuccess(v bool) *QueryMaterialTaskListResponseBody {
	s.Success = &v
	return s
}

func (s *QueryMaterialTaskListResponseBody) SetSynchro(v bool) *QueryMaterialTaskListResponseBody {
	s.Synchro = &v
	return s
}

func (s *QueryMaterialTaskListResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryMaterialTaskListResponseBodyModule struct {
	// Current page number.
	//
	// example:
	//
	// 12
	CurrentPageNum *int32 `json:"CurrentPageNum,omitempty" xml:"CurrentPageNum,omitempty"`
	// Request result.
	Data []*AppMaterialTask `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// Indicates whether a next page exists.
	//
	// example:
	//
	// False
	NextPage *bool `json:"NextPage,omitempty" xml:"NextPage,omitempty"`
	// Page size.
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
	// In addition to pagination limits, the server-side processes at most the latest 1 000 records for the current query. If the result exceeds 1 000 records, **ResultLimit*	- is **true**. In this case, narrow the time range and search again. Otherwise, **ResultLimit*	- is **false**.
	ResultLimit *bool `json:"ResultLimit,omitempty" xml:"ResultLimit,omitempty"`
	// Total number of records.
	//
	// example:
	//
	// 1
	TotalItemNum *int32 `json:"TotalItemNum,omitempty" xml:"TotalItemNum,omitempty"`
	// Total number of pages.
	//
	// example:
	//
	// 1
	TotalPageNum *int32 `json:"TotalPageNum,omitempty" xml:"TotalPageNum,omitempty"`
}

func (s QueryMaterialTaskListResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s QueryMaterialTaskListResponseBodyModule) GoString() string {
	return s.String()
}

func (s *QueryMaterialTaskListResponseBodyModule) GetCurrentPageNum() *int32 {
	return s.CurrentPageNum
}

func (s *QueryMaterialTaskListResponseBodyModule) GetData() []*AppMaterialTask {
	return s.Data
}

func (s *QueryMaterialTaskListResponseBodyModule) GetNextPage() *bool {
	return s.NextPage
}

func (s *QueryMaterialTaskListResponseBodyModule) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryMaterialTaskListResponseBodyModule) GetPrePage() *bool {
	return s.PrePage
}

func (s *QueryMaterialTaskListResponseBodyModule) GetResultLimit() *bool {
	return s.ResultLimit
}

func (s *QueryMaterialTaskListResponseBodyModule) GetTotalItemNum() *int32 {
	return s.TotalItemNum
}

func (s *QueryMaterialTaskListResponseBodyModule) GetTotalPageNum() *int32 {
	return s.TotalPageNum
}

func (s *QueryMaterialTaskListResponseBodyModule) SetCurrentPageNum(v int32) *QueryMaterialTaskListResponseBodyModule {
	s.CurrentPageNum = &v
	return s
}

func (s *QueryMaterialTaskListResponseBodyModule) SetData(v []*AppMaterialTask) *QueryMaterialTaskListResponseBodyModule {
	s.Data = v
	return s
}

func (s *QueryMaterialTaskListResponseBodyModule) SetNextPage(v bool) *QueryMaterialTaskListResponseBodyModule {
	s.NextPage = &v
	return s
}

func (s *QueryMaterialTaskListResponseBodyModule) SetPageSize(v int32) *QueryMaterialTaskListResponseBodyModule {
	s.PageSize = &v
	return s
}

func (s *QueryMaterialTaskListResponseBodyModule) SetPrePage(v bool) *QueryMaterialTaskListResponseBodyModule {
	s.PrePage = &v
	return s
}

func (s *QueryMaterialTaskListResponseBodyModule) SetResultLimit(v bool) *QueryMaterialTaskListResponseBodyModule {
	s.ResultLimit = &v
	return s
}

func (s *QueryMaterialTaskListResponseBodyModule) SetTotalItemNum(v int32) *QueryMaterialTaskListResponseBodyModule {
	s.TotalItemNum = &v
	return s
}

func (s *QueryMaterialTaskListResponseBodyModule) SetTotalPageNum(v int32) *QueryMaterialTaskListResponseBodyModule {
	s.TotalPageNum = &v
	return s
}

func (s *QueryMaterialTaskListResponseBodyModule) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
