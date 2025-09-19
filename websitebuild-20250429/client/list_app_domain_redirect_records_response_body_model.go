// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAppDomainRedirectRecordsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ListAppDomainRedirectRecordsResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *ListAppDomainRedirectRecordsResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *ListAppDomainRedirectRecordsResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *ListAppDomainRedirectRecordsResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *ListAppDomainRedirectRecordsResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *ListAppDomainRedirectRecordsResponseBody
	GetErrorArgs() []interface{}
	SetMaxResults(v int32) *ListAppDomainRedirectRecordsResponseBody
	GetMaxResults() *int32
	SetModule(v *ListAppDomainRedirectRecordsResponseBodyModule) *ListAppDomainRedirectRecordsResponseBody
	GetModule() *ListAppDomainRedirectRecordsResponseBodyModule
	SetNextToken(v string) *ListAppDomainRedirectRecordsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListAppDomainRedirectRecordsResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *ListAppDomainRedirectRecordsResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *ListAppDomainRedirectRecordsResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *ListAppDomainRedirectRecordsResponseBody
	GetSynchro() *bool
}

type ListAppDomainRedirectRecordsResponseBody struct {
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
	// https://check-result-file-sh.oss-cn-shanghai.aliyuncs.com/u6qw3gxzu3b7sbj/u6qw3gxzu3b7sbj.diff.zip?Expires=1740975709&OSSAccessKeyId=LTAI5tKUErVCETM4ev9SELNb&Signature=FP7dDnkrLlOZHmRRORVqbLOtv9c%3D
	DynamicMessage *string       `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorArgs      []interface{} `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	MaxResults *int32                                          `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	Module     *ListAppDomainRedirectRecordsResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
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
	// True
	Synchro *bool `json:"Synchro,omitempty" xml:"Synchro,omitempty"`
}

func (s ListAppDomainRedirectRecordsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAppDomainRedirectRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAppDomainRedirectRecordsResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ListAppDomainRedirectRecordsResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *ListAppDomainRedirectRecordsResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *ListAppDomainRedirectRecordsResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *ListAppDomainRedirectRecordsResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *ListAppDomainRedirectRecordsResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *ListAppDomainRedirectRecordsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAppDomainRedirectRecordsResponseBody) GetModule() *ListAppDomainRedirectRecordsResponseBodyModule {
	return s.Module
}

func (s *ListAppDomainRedirectRecordsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAppDomainRedirectRecordsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAppDomainRedirectRecordsResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *ListAppDomainRedirectRecordsResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *ListAppDomainRedirectRecordsResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *ListAppDomainRedirectRecordsResponseBody) SetAccessDeniedDetail(v string) *ListAppDomainRedirectRecordsResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListAppDomainRedirectRecordsResponseBody) SetAllowRetry(v bool) *ListAppDomainRedirectRecordsResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *ListAppDomainRedirectRecordsResponseBody) SetAppName(v string) *ListAppDomainRedirectRecordsResponseBody {
	s.AppName = &v
	return s
}

func (s *ListAppDomainRedirectRecordsResponseBody) SetDynamicCode(v string) *ListAppDomainRedirectRecordsResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *ListAppDomainRedirectRecordsResponseBody) SetDynamicMessage(v string) *ListAppDomainRedirectRecordsResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ListAppDomainRedirectRecordsResponseBody) SetErrorArgs(v []interface{}) *ListAppDomainRedirectRecordsResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *ListAppDomainRedirectRecordsResponseBody) SetMaxResults(v int32) *ListAppDomainRedirectRecordsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListAppDomainRedirectRecordsResponseBody) SetModule(v *ListAppDomainRedirectRecordsResponseBodyModule) *ListAppDomainRedirectRecordsResponseBody {
	s.Module = v
	return s
}

func (s *ListAppDomainRedirectRecordsResponseBody) SetNextToken(v string) *ListAppDomainRedirectRecordsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListAppDomainRedirectRecordsResponseBody) SetRequestId(v string) *ListAppDomainRedirectRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAppDomainRedirectRecordsResponseBody) SetRootErrorCode(v string) *ListAppDomainRedirectRecordsResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *ListAppDomainRedirectRecordsResponseBody) SetRootErrorMsg(v string) *ListAppDomainRedirectRecordsResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *ListAppDomainRedirectRecordsResponseBody) SetSynchro(v bool) *ListAppDomainRedirectRecordsResponseBody {
	s.Synchro = &v
	return s
}

func (s *ListAppDomainRedirectRecordsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListAppDomainRedirectRecordsResponseBodyModule struct {
	// example:
	//
	// 12
	CurrentPageNum *int32                                                `json:"CurrentPageNum,omitempty" xml:"CurrentPageNum,omitempty"`
	Data           []*ListAppDomainRedirectRecordsResponseBodyModuleData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Next           *ListAppDomainRedirectRecordsResponseBodyModuleNext   `json:"Next,omitempty" xml:"Next,omitempty" type:"Struct"`
	NextPage       *bool                                                 `json:"NextPage,omitempty" xml:"NextPage,omitempty"`
	// example:
	//
	// 50
	PageSize    *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PrePage     *bool  `json:"PrePage,omitempty" xml:"PrePage,omitempty"`
	ResultLimit *bool  `json:"ResultLimit,omitempty" xml:"ResultLimit,omitempty"`
	// example:
	//
	// 1
	TotalItemNum *int32 `json:"TotalItemNum,omitempty" xml:"TotalItemNum,omitempty"`
	// example:
	//
	// 1
	TotalPageNum *int32 `json:"TotalPageNum,omitempty" xml:"TotalPageNum,omitempty"`
}

func (s ListAppDomainRedirectRecordsResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s ListAppDomainRedirectRecordsResponseBodyModule) GoString() string {
	return s.String()
}

func (s *ListAppDomainRedirectRecordsResponseBodyModule) GetCurrentPageNum() *int32 {
	return s.CurrentPageNum
}

func (s *ListAppDomainRedirectRecordsResponseBodyModule) GetData() []*ListAppDomainRedirectRecordsResponseBodyModuleData {
	return s.Data
}

func (s *ListAppDomainRedirectRecordsResponseBodyModule) GetNext() *ListAppDomainRedirectRecordsResponseBodyModuleNext {
	return s.Next
}

func (s *ListAppDomainRedirectRecordsResponseBodyModule) GetNextPage() *bool {
	return s.NextPage
}

func (s *ListAppDomainRedirectRecordsResponseBodyModule) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAppDomainRedirectRecordsResponseBodyModule) GetPrePage() *bool {
	return s.PrePage
}

func (s *ListAppDomainRedirectRecordsResponseBodyModule) GetResultLimit() *bool {
	return s.ResultLimit
}

func (s *ListAppDomainRedirectRecordsResponseBodyModule) GetTotalItemNum() *int32 {
	return s.TotalItemNum
}

func (s *ListAppDomainRedirectRecordsResponseBodyModule) GetTotalPageNum() *int32 {
	return s.TotalPageNum
}

func (s *ListAppDomainRedirectRecordsResponseBodyModule) SetCurrentPageNum(v int32) *ListAppDomainRedirectRecordsResponseBodyModule {
	s.CurrentPageNum = &v
	return s
}

func (s *ListAppDomainRedirectRecordsResponseBodyModule) SetData(v []*ListAppDomainRedirectRecordsResponseBodyModuleData) *ListAppDomainRedirectRecordsResponseBodyModule {
	s.Data = v
	return s
}

func (s *ListAppDomainRedirectRecordsResponseBodyModule) SetNext(v *ListAppDomainRedirectRecordsResponseBodyModuleNext) *ListAppDomainRedirectRecordsResponseBodyModule {
	s.Next = v
	return s
}

func (s *ListAppDomainRedirectRecordsResponseBodyModule) SetNextPage(v bool) *ListAppDomainRedirectRecordsResponseBodyModule {
	s.NextPage = &v
	return s
}

func (s *ListAppDomainRedirectRecordsResponseBodyModule) SetPageSize(v int32) *ListAppDomainRedirectRecordsResponseBodyModule {
	s.PageSize = &v
	return s
}

func (s *ListAppDomainRedirectRecordsResponseBodyModule) SetPrePage(v bool) *ListAppDomainRedirectRecordsResponseBodyModule {
	s.PrePage = &v
	return s
}

func (s *ListAppDomainRedirectRecordsResponseBodyModule) SetResultLimit(v bool) *ListAppDomainRedirectRecordsResponseBodyModule {
	s.ResultLimit = &v
	return s
}

func (s *ListAppDomainRedirectRecordsResponseBodyModule) SetTotalItemNum(v int32) *ListAppDomainRedirectRecordsResponseBodyModule {
	s.TotalItemNum = &v
	return s
}

func (s *ListAppDomainRedirectRecordsResponseBodyModule) SetTotalPageNum(v int32) *ListAppDomainRedirectRecordsResponseBodyModule {
	s.TotalPageNum = &v
	return s
}

func (s *ListAppDomainRedirectRecordsResponseBodyModule) Validate() error {
	return dara.Validate(s)
}

type ListAppDomainRedirectRecordsResponseBodyModuleData struct {
	// example:
	//
	// f0379419-433d-410e-98d9-bf5c72f47227
	RecordId *string `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	// example:
	//
	// abc.wanwang.xin
	SourceDomain *string `json:"SourceDomain,omitempty" xml:"SourceDomain,omitempty"`
	// example:
	//
	// aliyuncs.com
	TargetDomain *string `json:"TargetDomain,omitempty" xml:"TargetDomain,omitempty"`
}

func (s ListAppDomainRedirectRecordsResponseBodyModuleData) String() string {
	return dara.Prettify(s)
}

func (s ListAppDomainRedirectRecordsResponseBodyModuleData) GoString() string {
	return s.String()
}

func (s *ListAppDomainRedirectRecordsResponseBodyModuleData) GetRecordId() *string {
	return s.RecordId
}

func (s *ListAppDomainRedirectRecordsResponseBodyModuleData) GetSourceDomain() *string {
	return s.SourceDomain
}

func (s *ListAppDomainRedirectRecordsResponseBodyModuleData) GetTargetDomain() *string {
	return s.TargetDomain
}

func (s *ListAppDomainRedirectRecordsResponseBodyModuleData) SetRecordId(v string) *ListAppDomainRedirectRecordsResponseBodyModuleData {
	s.RecordId = &v
	return s
}

func (s *ListAppDomainRedirectRecordsResponseBodyModuleData) SetSourceDomain(v string) *ListAppDomainRedirectRecordsResponseBodyModuleData {
	s.SourceDomain = &v
	return s
}

func (s *ListAppDomainRedirectRecordsResponseBodyModuleData) SetTargetDomain(v string) *ListAppDomainRedirectRecordsResponseBodyModuleData {
	s.TargetDomain = &v
	return s
}

func (s *ListAppDomainRedirectRecordsResponseBodyModuleData) Validate() error {
	return dara.Validate(s)
}

type ListAppDomainRedirectRecordsResponseBodyModuleNext struct {
	// example:
	//
	// 936956504373539840
	RecordId *string `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	// example:
	//
	// abc.wanwang.xin
	SourceDomain *string `json:"SourceDomain,omitempty" xml:"SourceDomain,omitempty"`
	// example:
	//
	// aliyuncs.com
	TargetDomain *string `json:"TargetDomain,omitempty" xml:"TargetDomain,omitempty"`
}

func (s ListAppDomainRedirectRecordsResponseBodyModuleNext) String() string {
	return dara.Prettify(s)
}

func (s ListAppDomainRedirectRecordsResponseBodyModuleNext) GoString() string {
	return s.String()
}

func (s *ListAppDomainRedirectRecordsResponseBodyModuleNext) GetRecordId() *string {
	return s.RecordId
}

func (s *ListAppDomainRedirectRecordsResponseBodyModuleNext) GetSourceDomain() *string {
	return s.SourceDomain
}

func (s *ListAppDomainRedirectRecordsResponseBodyModuleNext) GetTargetDomain() *string {
	return s.TargetDomain
}

func (s *ListAppDomainRedirectRecordsResponseBodyModuleNext) SetRecordId(v string) *ListAppDomainRedirectRecordsResponseBodyModuleNext {
	s.RecordId = &v
	return s
}

func (s *ListAppDomainRedirectRecordsResponseBodyModuleNext) SetSourceDomain(v string) *ListAppDomainRedirectRecordsResponseBodyModuleNext {
	s.SourceDomain = &v
	return s
}

func (s *ListAppDomainRedirectRecordsResponseBodyModuleNext) SetTargetDomain(v string) *ListAppDomainRedirectRecordsResponseBodyModuleNext {
	s.TargetDomain = &v
	return s
}

func (s *ListAppDomainRedirectRecordsResponseBodyModuleNext) Validate() error {
	return dara.Validate(s)
}
