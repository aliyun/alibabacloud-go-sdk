// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMaterialFileListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *QueryMaterialFileListResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *QueryMaterialFileListResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *QueryMaterialFileListResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *QueryMaterialFileListResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *QueryMaterialFileListResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *QueryMaterialFileListResponseBody
	GetErrorArgs() []interface{}
	SetErrorCode(v string) *QueryMaterialFileListResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *QueryMaterialFileListResponseBody
	GetErrorMsg() *string
	SetMaxResults(v int32) *QueryMaterialFileListResponseBody
	GetMaxResults() *int32
	SetModule(v *QueryMaterialFileListResponseBodyModule) *QueryMaterialFileListResponseBody
	GetModule() *QueryMaterialFileListResponseBodyModule
	SetNextToken(v string) *QueryMaterialFileListResponseBody
	GetNextToken() *string
	SetRequestId(v string) *QueryMaterialFileListResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *QueryMaterialFileListResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *QueryMaterialFileListResponseBody
	GetRootErrorMsg() *string
	SetSuccess(v bool) *QueryMaterialFileListResponseBody
	GetSuccess() *bool
	SetSynchro(v bool) *QueryMaterialFileListResponseBody
	GetSynchro() *bool
}

type QueryMaterialFileListResponseBody struct {
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
	// https://check-result-file-sh.oss-cn-shanghai.aliyuncs.com/u6qw3gxzu3b7sbj/u6qw3gxzu3b7sbj.diff.zip?Expires=1740975709&OSSAccessKeyId=LTAI5tKUErVCETM4ev9SELNb&Signature=FP7dDnkrLlOZHmRRORVqbLOtv9c%3D
	DynamicMessage *string       `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorArgs      []interface{} `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// aliuid:1998006665794443 assumeRole not exist,serviceName:realtimelogpush.dcdnservices.aliyuncs.com
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// example:
	//
	// 10
	MaxResults *int32                                   `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	Module     *QueryMaterialFileListResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
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
	// example:
	//
	// SYSTEM.ERROR
	RootErrorCode *string `json:"RootErrorCode,omitempty" xml:"RootErrorCode,omitempty"`
	RootErrorMsg  *string `json:"RootErrorMsg,omitempty" xml:"RootErrorMsg,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// True
	Synchro *bool `json:"Synchro,omitempty" xml:"Synchro,omitempty"`
}

func (s QueryMaterialFileListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryMaterialFileListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMaterialFileListResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *QueryMaterialFileListResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *QueryMaterialFileListResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *QueryMaterialFileListResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *QueryMaterialFileListResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *QueryMaterialFileListResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *QueryMaterialFileListResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *QueryMaterialFileListResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *QueryMaterialFileListResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *QueryMaterialFileListResponseBody) GetModule() *QueryMaterialFileListResponseBodyModule {
	return s.Module
}

func (s *QueryMaterialFileListResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *QueryMaterialFileListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryMaterialFileListResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *QueryMaterialFileListResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *QueryMaterialFileListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryMaterialFileListResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *QueryMaterialFileListResponseBody) SetAccessDeniedDetail(v string) *QueryMaterialFileListResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *QueryMaterialFileListResponseBody) SetAllowRetry(v bool) *QueryMaterialFileListResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *QueryMaterialFileListResponseBody) SetAppName(v string) *QueryMaterialFileListResponseBody {
	s.AppName = &v
	return s
}

func (s *QueryMaterialFileListResponseBody) SetDynamicCode(v string) *QueryMaterialFileListResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *QueryMaterialFileListResponseBody) SetDynamicMessage(v string) *QueryMaterialFileListResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *QueryMaterialFileListResponseBody) SetErrorArgs(v []interface{}) *QueryMaterialFileListResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *QueryMaterialFileListResponseBody) SetErrorCode(v string) *QueryMaterialFileListResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *QueryMaterialFileListResponseBody) SetErrorMsg(v string) *QueryMaterialFileListResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *QueryMaterialFileListResponseBody) SetMaxResults(v int32) *QueryMaterialFileListResponseBody {
	s.MaxResults = &v
	return s
}

func (s *QueryMaterialFileListResponseBody) SetModule(v *QueryMaterialFileListResponseBodyModule) *QueryMaterialFileListResponseBody {
	s.Module = v
	return s
}

func (s *QueryMaterialFileListResponseBody) SetNextToken(v string) *QueryMaterialFileListResponseBody {
	s.NextToken = &v
	return s
}

func (s *QueryMaterialFileListResponseBody) SetRequestId(v string) *QueryMaterialFileListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryMaterialFileListResponseBody) SetRootErrorCode(v string) *QueryMaterialFileListResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *QueryMaterialFileListResponseBody) SetRootErrorMsg(v string) *QueryMaterialFileListResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *QueryMaterialFileListResponseBody) SetSuccess(v bool) *QueryMaterialFileListResponseBody {
	s.Success = &v
	return s
}

func (s *QueryMaterialFileListResponseBody) SetSynchro(v bool) *QueryMaterialFileListResponseBody {
	s.Synchro = &v
	return s
}

func (s *QueryMaterialFileListResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryMaterialFileListResponseBodyModule struct {
	// example:
	//
	// 1
	CurrentPageNum *int32             `json:"CurrentPageNum,omitempty" xml:"CurrentPageNum,omitempty"`
	Data           []*AppMaterialFile `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	NextPage       *bool              `json:"NextPage,omitempty" xml:"NextPage,omitempty"`
	// example:
	//
	// 50
	PageSize    *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PrePage     *bool  `json:"PrePage,omitempty" xml:"PrePage,omitempty"`
	ResultLimit *bool  `json:"ResultLimit,omitempty" xml:"ResultLimit,omitempty"`
	// example:
	//
	// 0
	TotalItemNum *int32 `json:"TotalItemNum,omitempty" xml:"TotalItemNum,omitempty"`
	// example:
	//
	// 1
	TotalPageNum *int32 `json:"TotalPageNum,omitempty" xml:"TotalPageNum,omitempty"`
}

func (s QueryMaterialFileListResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s QueryMaterialFileListResponseBodyModule) GoString() string {
	return s.String()
}

func (s *QueryMaterialFileListResponseBodyModule) GetCurrentPageNum() *int32 {
	return s.CurrentPageNum
}

func (s *QueryMaterialFileListResponseBodyModule) GetData() []*AppMaterialFile {
	return s.Data
}

func (s *QueryMaterialFileListResponseBodyModule) GetNextPage() *bool {
	return s.NextPage
}

func (s *QueryMaterialFileListResponseBodyModule) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryMaterialFileListResponseBodyModule) GetPrePage() *bool {
	return s.PrePage
}

func (s *QueryMaterialFileListResponseBodyModule) GetResultLimit() *bool {
	return s.ResultLimit
}

func (s *QueryMaterialFileListResponseBodyModule) GetTotalItemNum() *int32 {
	return s.TotalItemNum
}

func (s *QueryMaterialFileListResponseBodyModule) GetTotalPageNum() *int32 {
	return s.TotalPageNum
}

func (s *QueryMaterialFileListResponseBodyModule) SetCurrentPageNum(v int32) *QueryMaterialFileListResponseBodyModule {
	s.CurrentPageNum = &v
	return s
}

func (s *QueryMaterialFileListResponseBodyModule) SetData(v []*AppMaterialFile) *QueryMaterialFileListResponseBodyModule {
	s.Data = v
	return s
}

func (s *QueryMaterialFileListResponseBodyModule) SetNextPage(v bool) *QueryMaterialFileListResponseBodyModule {
	s.NextPage = &v
	return s
}

func (s *QueryMaterialFileListResponseBodyModule) SetPageSize(v int32) *QueryMaterialFileListResponseBodyModule {
	s.PageSize = &v
	return s
}

func (s *QueryMaterialFileListResponseBodyModule) SetPrePage(v bool) *QueryMaterialFileListResponseBodyModule {
	s.PrePage = &v
	return s
}

func (s *QueryMaterialFileListResponseBodyModule) SetResultLimit(v bool) *QueryMaterialFileListResponseBodyModule {
	s.ResultLimit = &v
	return s
}

func (s *QueryMaterialFileListResponseBodyModule) SetTotalItemNum(v int32) *QueryMaterialFileListResponseBodyModule {
	s.TotalItemNum = &v
	return s
}

func (s *QueryMaterialFileListResponseBodyModule) SetTotalPageNum(v int32) *QueryMaterialFileListResponseBodyModule {
	s.TotalPageNum = &v
	return s
}

func (s *QueryMaterialFileListResponseBodyModule) Validate() error {
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
