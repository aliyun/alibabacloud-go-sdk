// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAppCommoditySpecificationsV2ForPartnerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ListAppCommoditySpecificationsV2ForPartnerResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *ListAppCommoditySpecificationsV2ForPartnerResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *ListAppCommoditySpecificationsV2ForPartnerResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *ListAppCommoditySpecificationsV2ForPartnerResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *ListAppCommoditySpecificationsV2ForPartnerResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *ListAppCommoditySpecificationsV2ForPartnerResponseBody
	GetErrorArgs() []interface{}
	SetMaxResults(v int32) *ListAppCommoditySpecificationsV2ForPartnerResponseBody
	GetMaxResults() *int32
	SetModule(v *ListAppCommoditySpecificationsV2ForPartnerResponseBodyModule) *ListAppCommoditySpecificationsV2ForPartnerResponseBody
	GetModule() *ListAppCommoditySpecificationsV2ForPartnerResponseBodyModule
	SetNextToken(v string) *ListAppCommoditySpecificationsV2ForPartnerResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListAppCommoditySpecificationsV2ForPartnerResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *ListAppCommoditySpecificationsV2ForPartnerResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *ListAppCommoditySpecificationsV2ForPartnerResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *ListAppCommoditySpecificationsV2ForPartnerResponseBody
	GetSynchro() *bool
}

type ListAppCommoditySpecificationsV2ForPartnerResponseBody struct {
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
	MaxResults *int32                                                        `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	Module     *ListAppCommoditySpecificationsV2ForPartnerResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
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

func (s ListAppCommoditySpecificationsV2ForPartnerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAppCommoditySpecificationsV2ForPartnerResponseBody) GoString() string {
	return s.String()
}

func (s *ListAppCommoditySpecificationsV2ForPartnerResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ListAppCommoditySpecificationsV2ForPartnerResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *ListAppCommoditySpecificationsV2ForPartnerResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *ListAppCommoditySpecificationsV2ForPartnerResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *ListAppCommoditySpecificationsV2ForPartnerResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *ListAppCommoditySpecificationsV2ForPartnerResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *ListAppCommoditySpecificationsV2ForPartnerResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAppCommoditySpecificationsV2ForPartnerResponseBody) GetModule() *ListAppCommoditySpecificationsV2ForPartnerResponseBodyModule {
	return s.Module
}

func (s *ListAppCommoditySpecificationsV2ForPartnerResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAppCommoditySpecificationsV2ForPartnerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAppCommoditySpecificationsV2ForPartnerResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *ListAppCommoditySpecificationsV2ForPartnerResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *ListAppCommoditySpecificationsV2ForPartnerResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *ListAppCommoditySpecificationsV2ForPartnerResponseBody) SetAccessDeniedDetail(v string) *ListAppCommoditySpecificationsV2ForPartnerResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListAppCommoditySpecificationsV2ForPartnerResponseBody) SetAllowRetry(v bool) *ListAppCommoditySpecificationsV2ForPartnerResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *ListAppCommoditySpecificationsV2ForPartnerResponseBody) SetAppName(v string) *ListAppCommoditySpecificationsV2ForPartnerResponseBody {
	s.AppName = &v
	return s
}

func (s *ListAppCommoditySpecificationsV2ForPartnerResponseBody) SetDynamicCode(v string) *ListAppCommoditySpecificationsV2ForPartnerResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *ListAppCommoditySpecificationsV2ForPartnerResponseBody) SetDynamicMessage(v string) *ListAppCommoditySpecificationsV2ForPartnerResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ListAppCommoditySpecificationsV2ForPartnerResponseBody) SetErrorArgs(v []interface{}) *ListAppCommoditySpecificationsV2ForPartnerResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *ListAppCommoditySpecificationsV2ForPartnerResponseBody) SetMaxResults(v int32) *ListAppCommoditySpecificationsV2ForPartnerResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListAppCommoditySpecificationsV2ForPartnerResponseBody) SetModule(v *ListAppCommoditySpecificationsV2ForPartnerResponseBodyModule) *ListAppCommoditySpecificationsV2ForPartnerResponseBody {
	s.Module = v
	return s
}

func (s *ListAppCommoditySpecificationsV2ForPartnerResponseBody) SetNextToken(v string) *ListAppCommoditySpecificationsV2ForPartnerResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListAppCommoditySpecificationsV2ForPartnerResponseBody) SetRequestId(v string) *ListAppCommoditySpecificationsV2ForPartnerResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAppCommoditySpecificationsV2ForPartnerResponseBody) SetRootErrorCode(v string) *ListAppCommoditySpecificationsV2ForPartnerResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *ListAppCommoditySpecificationsV2ForPartnerResponseBody) SetRootErrorMsg(v string) *ListAppCommoditySpecificationsV2ForPartnerResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *ListAppCommoditySpecificationsV2ForPartnerResponseBody) SetSynchro(v bool) *ListAppCommoditySpecificationsV2ForPartnerResponseBody {
	s.Synchro = &v
	return s
}

func (s *ListAppCommoditySpecificationsV2ForPartnerResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListAppCommoditySpecificationsV2ForPartnerResponseBodyModule struct {
	// example:
	//
	// 12
	CurrentPageNum *int32                                                              `json:"CurrentPageNum,omitempty" xml:"CurrentPageNum,omitempty"`
	Data           []*ListAppCommoditySpecificationsV2ForPartnerResponseBodyModuleData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Next           *ListAppCommoditySpecificationsV2ForPartnerResponseBodyModuleNext   `json:"Next,omitempty" xml:"Next,omitempty" type:"Struct"`
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
	PrePage     *bool `json:"PrePage,omitempty" xml:"PrePage,omitempty"`
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

func (s ListAppCommoditySpecificationsV2ForPartnerResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s ListAppCommoditySpecificationsV2ForPartnerResponseBodyModule) GoString() string {
	return s.String()
}

func (s *ListAppCommoditySpecificationsV2ForPartnerResponseBodyModule) GetCurrentPageNum() *int32 {
	return s.CurrentPageNum
}

func (s *ListAppCommoditySpecificationsV2ForPartnerResponseBodyModule) GetData() []*ListAppCommoditySpecificationsV2ForPartnerResponseBodyModuleData {
	return s.Data
}

func (s *ListAppCommoditySpecificationsV2ForPartnerResponseBodyModule) GetNext() *ListAppCommoditySpecificationsV2ForPartnerResponseBodyModuleNext {
	return s.Next
}

func (s *ListAppCommoditySpecificationsV2ForPartnerResponseBodyModule) GetNextPage() *bool {
	return s.NextPage
}

func (s *ListAppCommoditySpecificationsV2ForPartnerResponseBodyModule) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAppCommoditySpecificationsV2ForPartnerResponseBodyModule) GetPrePage() *bool {
	return s.PrePage
}

func (s *ListAppCommoditySpecificationsV2ForPartnerResponseBodyModule) GetResultLimit() *bool {
	return s.ResultLimit
}

func (s *ListAppCommoditySpecificationsV2ForPartnerResponseBodyModule) GetTotalItemNum() *int32 {
	return s.TotalItemNum
}

func (s *ListAppCommoditySpecificationsV2ForPartnerResponseBodyModule) GetTotalPageNum() *int32 {
	return s.TotalPageNum
}

func (s *ListAppCommoditySpecificationsV2ForPartnerResponseBodyModule) SetCurrentPageNum(v int32) *ListAppCommoditySpecificationsV2ForPartnerResponseBodyModule {
	s.CurrentPageNum = &v
	return s
}

func (s *ListAppCommoditySpecificationsV2ForPartnerResponseBodyModule) SetData(v []*ListAppCommoditySpecificationsV2ForPartnerResponseBodyModuleData) *ListAppCommoditySpecificationsV2ForPartnerResponseBodyModule {
	s.Data = v
	return s
}

func (s *ListAppCommoditySpecificationsV2ForPartnerResponseBodyModule) SetNext(v *ListAppCommoditySpecificationsV2ForPartnerResponseBodyModuleNext) *ListAppCommoditySpecificationsV2ForPartnerResponseBodyModule {
	s.Next = v
	return s
}

func (s *ListAppCommoditySpecificationsV2ForPartnerResponseBodyModule) SetNextPage(v bool) *ListAppCommoditySpecificationsV2ForPartnerResponseBodyModule {
	s.NextPage = &v
	return s
}

func (s *ListAppCommoditySpecificationsV2ForPartnerResponseBodyModule) SetPageSize(v int32) *ListAppCommoditySpecificationsV2ForPartnerResponseBodyModule {
	s.PageSize = &v
	return s
}

func (s *ListAppCommoditySpecificationsV2ForPartnerResponseBodyModule) SetPrePage(v bool) *ListAppCommoditySpecificationsV2ForPartnerResponseBodyModule {
	s.PrePage = &v
	return s
}

func (s *ListAppCommoditySpecificationsV2ForPartnerResponseBodyModule) SetResultLimit(v bool) *ListAppCommoditySpecificationsV2ForPartnerResponseBodyModule {
	s.ResultLimit = &v
	return s
}

func (s *ListAppCommoditySpecificationsV2ForPartnerResponseBodyModule) SetTotalItemNum(v int32) *ListAppCommoditySpecificationsV2ForPartnerResponseBodyModule {
	s.TotalItemNum = &v
	return s
}

func (s *ListAppCommoditySpecificationsV2ForPartnerResponseBodyModule) SetTotalPageNum(v int32) *ListAppCommoditySpecificationsV2ForPartnerResponseBodyModule {
	s.TotalPageNum = &v
	return s
}

func (s *ListAppCommoditySpecificationsV2ForPartnerResponseBodyModule) Validate() error {
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

type ListAppCommoditySpecificationsV2ForPartnerResponseBodyModuleData struct {
	// example:
	//
	// 200
	Code       *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Components map[string]*ModuleDataComponentsValue `json:"Components,omitempty" xml:"Components,omitempty"`
	Name       *string                               `json:"Name,omitempty" xml:"Name,omitempty"`
	Properties map[string]*ModuleDataPropertiesValue `json:"Properties,omitempty" xml:"Properties,omitempty"`
}

func (s ListAppCommoditySpecificationsV2ForPartnerResponseBodyModuleData) String() string {
	return dara.Prettify(s)
}

func (s ListAppCommoditySpecificationsV2ForPartnerResponseBodyModuleData) GoString() string {
	return s.String()
}

func (s *ListAppCommoditySpecificationsV2ForPartnerResponseBodyModuleData) GetCode() *string {
	return s.Code
}

func (s *ListAppCommoditySpecificationsV2ForPartnerResponseBodyModuleData) GetComponents() map[string]*ModuleDataComponentsValue {
	return s.Components
}

func (s *ListAppCommoditySpecificationsV2ForPartnerResponseBodyModuleData) GetName() *string {
	return s.Name
}

func (s *ListAppCommoditySpecificationsV2ForPartnerResponseBodyModuleData) GetProperties() map[string]*ModuleDataPropertiesValue {
	return s.Properties
}

func (s *ListAppCommoditySpecificationsV2ForPartnerResponseBodyModuleData) SetCode(v string) *ListAppCommoditySpecificationsV2ForPartnerResponseBodyModuleData {
	s.Code = &v
	return s
}

func (s *ListAppCommoditySpecificationsV2ForPartnerResponseBodyModuleData) SetComponents(v map[string]*ModuleDataComponentsValue) *ListAppCommoditySpecificationsV2ForPartnerResponseBodyModuleData {
	s.Components = v
	return s
}

func (s *ListAppCommoditySpecificationsV2ForPartnerResponseBodyModuleData) SetName(v string) *ListAppCommoditySpecificationsV2ForPartnerResponseBodyModuleData {
	s.Name = &v
	return s
}

func (s *ListAppCommoditySpecificationsV2ForPartnerResponseBodyModuleData) SetProperties(v map[string]*ModuleDataPropertiesValue) *ListAppCommoditySpecificationsV2ForPartnerResponseBodyModuleData {
	s.Properties = v
	return s
}

func (s *ListAppCommoditySpecificationsV2ForPartnerResponseBodyModuleData) Validate() error {
	return dara.Validate(s)
}

type ListAppCommoditySpecificationsV2ForPartnerResponseBodyModuleNext struct {
	// example:
	//
	// OK
	Code       *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Components map[string]*ModuleNextComponentsValue `json:"Components,omitempty" xml:"Components,omitempty"`
	// example:
	//
	// docs
	Name       *string                               `json:"Name,omitempty" xml:"Name,omitempty"`
	Properties map[string]*ModuleNextPropertiesValue `json:"Properties,omitempty" xml:"Properties,omitempty"`
}

func (s ListAppCommoditySpecificationsV2ForPartnerResponseBodyModuleNext) String() string {
	return dara.Prettify(s)
}

func (s ListAppCommoditySpecificationsV2ForPartnerResponseBodyModuleNext) GoString() string {
	return s.String()
}

func (s *ListAppCommoditySpecificationsV2ForPartnerResponseBodyModuleNext) GetCode() *string {
	return s.Code
}

func (s *ListAppCommoditySpecificationsV2ForPartnerResponseBodyModuleNext) GetComponents() map[string]*ModuleNextComponentsValue {
	return s.Components
}

func (s *ListAppCommoditySpecificationsV2ForPartnerResponseBodyModuleNext) GetName() *string {
	return s.Name
}

func (s *ListAppCommoditySpecificationsV2ForPartnerResponseBodyModuleNext) GetProperties() map[string]*ModuleNextPropertiesValue {
	return s.Properties
}

func (s *ListAppCommoditySpecificationsV2ForPartnerResponseBodyModuleNext) SetCode(v string) *ListAppCommoditySpecificationsV2ForPartnerResponseBodyModuleNext {
	s.Code = &v
	return s
}

func (s *ListAppCommoditySpecificationsV2ForPartnerResponseBodyModuleNext) SetComponents(v map[string]*ModuleNextComponentsValue) *ListAppCommoditySpecificationsV2ForPartnerResponseBodyModuleNext {
	s.Components = v
	return s
}

func (s *ListAppCommoditySpecificationsV2ForPartnerResponseBodyModuleNext) SetName(v string) *ListAppCommoditySpecificationsV2ForPartnerResponseBodyModuleNext {
	s.Name = &v
	return s
}

func (s *ListAppCommoditySpecificationsV2ForPartnerResponseBodyModuleNext) SetProperties(v map[string]*ModuleNextPropertiesValue) *ListAppCommoditySpecificationsV2ForPartnerResponseBodyModuleNext {
	s.Properties = v
	return s
}

func (s *ListAppCommoditySpecificationsV2ForPartnerResponseBodyModuleNext) Validate() error {
	return dara.Validate(s)
}
