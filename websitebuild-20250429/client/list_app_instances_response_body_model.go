// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAppInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ListAppInstancesResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *ListAppInstancesResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *ListAppInstancesResponseBody
	GetAppName() *string
	SetCurrentPageNum(v int32) *ListAppInstancesResponseBody
	GetCurrentPageNum() *int32
	SetData(v []*AppInstanceAggregate) *ListAppInstancesResponseBody
	GetData() []*AppInstanceAggregate
	SetDynamicCode(v string) *ListAppInstancesResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *ListAppInstancesResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *ListAppInstancesResponseBody
	GetErrorArgs() []interface{}
	SetMaxResults(v int32) *ListAppInstancesResponseBody
	GetMaxResults() *int32
	SetNextPage(v bool) *ListAppInstancesResponseBody
	GetNextPage() *bool
	SetNextToken(v string) *ListAppInstancesResponseBody
	GetNextToken() *string
	SetPageSize(v int32) *ListAppInstancesResponseBody
	GetPageSize() *int32
	SetPrePage(v bool) *ListAppInstancesResponseBody
	GetPrePage() *bool
	SetRequestId(v string) *ListAppInstancesResponseBody
	GetRequestId() *string
	SetResultLimit(v bool) *ListAppInstancesResponseBody
	GetResultLimit() *bool
	SetRootErrorCode(v string) *ListAppInstancesResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *ListAppInstancesResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *ListAppInstancesResponseBody
	GetSynchro() *bool
	SetTotalPageNum(v int32) *ListAppInstancesResponseBody
	GetTotalPageNum() *int32
}

type ListAppInstancesResponseBody struct {
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
	// dewuApp
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// 1
	CurrentPageNum *int32                  `json:"CurrentPageNum,omitempty" xml:"CurrentPageNum,omitempty"`
	Data           []*AppInstanceAggregate `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// ERROR-oo1
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// example:
	//
	// https://check-result-file-sh.oss-cn-shanghai.aliyuncs.com/u6qw3gxzu3b7sbj/u6qw3gxzu3b7sbj.diff.zip?Expires=1740975709&OSSAccessKeyId=LTAI5tKUErVCETM4ev9SELNb&Signature=FP7dDnkrLlOZHmRRORVqbLOtv9c%3D
	DynamicMessage *string       `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorArgs      []interface{} `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// False
	NextPage *bool `json:"NextPage,omitempty" xml:"NextPage,omitempty"`
	// example:
	//
	// dw+qdTi1EjVSWX/INJdYNw==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// False
	PrePage *bool `json:"PrePage,omitempty" xml:"PrePage,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 6C6B99AC-39EC-5350-874C-204128C905E6
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultLimit *bool   `json:"ResultLimit,omitempty" xml:"ResultLimit,omitempty"`
	// example:
	//
	// SYSTEM.ERROR
	RootErrorCode *string `json:"RootErrorCode,omitempty" xml:"RootErrorCode,omitempty"`
	RootErrorMsg  *string `json:"RootErrorMsg,omitempty" xml:"RootErrorMsg,omitempty"`
	// example:
	//
	// True
	Synchro *bool `json:"Synchro,omitempty" xml:"Synchro,omitempty"`
	// example:
	//
	// 4
	TotalPageNum *int32 `json:"TotalPageNum,omitempty" xml:"TotalPageNum,omitempty"`
}

func (s ListAppInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAppInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAppInstancesResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ListAppInstancesResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *ListAppInstancesResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *ListAppInstancesResponseBody) GetCurrentPageNum() *int32 {
	return s.CurrentPageNum
}

func (s *ListAppInstancesResponseBody) GetData() []*AppInstanceAggregate {
	return s.Data
}

func (s *ListAppInstancesResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *ListAppInstancesResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *ListAppInstancesResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *ListAppInstancesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAppInstancesResponseBody) GetNextPage() *bool {
	return s.NextPage
}

func (s *ListAppInstancesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAppInstancesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAppInstancesResponseBody) GetPrePage() *bool {
	return s.PrePage
}

func (s *ListAppInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAppInstancesResponseBody) GetResultLimit() *bool {
	return s.ResultLimit
}

func (s *ListAppInstancesResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *ListAppInstancesResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *ListAppInstancesResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *ListAppInstancesResponseBody) GetTotalPageNum() *int32 {
	return s.TotalPageNum
}

func (s *ListAppInstancesResponseBody) SetAccessDeniedDetail(v string) *ListAppInstancesResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListAppInstancesResponseBody) SetAllowRetry(v bool) *ListAppInstancesResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *ListAppInstancesResponseBody) SetAppName(v string) *ListAppInstancesResponseBody {
	s.AppName = &v
	return s
}

func (s *ListAppInstancesResponseBody) SetCurrentPageNum(v int32) *ListAppInstancesResponseBody {
	s.CurrentPageNum = &v
	return s
}

func (s *ListAppInstancesResponseBody) SetData(v []*AppInstanceAggregate) *ListAppInstancesResponseBody {
	s.Data = v
	return s
}

func (s *ListAppInstancesResponseBody) SetDynamicCode(v string) *ListAppInstancesResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *ListAppInstancesResponseBody) SetDynamicMessage(v string) *ListAppInstancesResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ListAppInstancesResponseBody) SetErrorArgs(v []interface{}) *ListAppInstancesResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *ListAppInstancesResponseBody) SetMaxResults(v int32) *ListAppInstancesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListAppInstancesResponseBody) SetNextPage(v bool) *ListAppInstancesResponseBody {
	s.NextPage = &v
	return s
}

func (s *ListAppInstancesResponseBody) SetNextToken(v string) *ListAppInstancesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListAppInstancesResponseBody) SetPageSize(v int32) *ListAppInstancesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListAppInstancesResponseBody) SetPrePage(v bool) *ListAppInstancesResponseBody {
	s.PrePage = &v
	return s
}

func (s *ListAppInstancesResponseBody) SetRequestId(v string) *ListAppInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAppInstancesResponseBody) SetResultLimit(v bool) *ListAppInstancesResponseBody {
	s.ResultLimit = &v
	return s
}

func (s *ListAppInstancesResponseBody) SetRootErrorCode(v string) *ListAppInstancesResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *ListAppInstancesResponseBody) SetRootErrorMsg(v string) *ListAppInstancesResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *ListAppInstancesResponseBody) SetSynchro(v bool) *ListAppInstancesResponseBody {
	s.Synchro = &v
	return s
}

func (s *ListAppInstancesResponseBody) SetTotalPageNum(v int32) *ListAppInstancesResponseBody {
	s.TotalPageNum = &v
	return s
}

func (s *ListAppInstancesResponseBody) Validate() error {
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
