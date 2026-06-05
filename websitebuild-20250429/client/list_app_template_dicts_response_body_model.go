// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAppTemplateDictsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ListAppTemplateDictsResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *ListAppTemplateDictsResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *ListAppTemplateDictsResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *ListAppTemplateDictsResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *ListAppTemplateDictsResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *ListAppTemplateDictsResponseBody
	GetErrorArgs() []interface{}
	SetMaxResults(v int32) *ListAppTemplateDictsResponseBody
	GetMaxResults() *int32
	SetModule(v []*ListAppTemplateDictsResponseBodyModule) *ListAppTemplateDictsResponseBody
	GetModule() []*ListAppTemplateDictsResponseBodyModule
	SetNextToken(v string) *ListAppTemplateDictsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListAppTemplateDictsResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *ListAppTemplateDictsResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *ListAppTemplateDictsResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *ListAppTemplateDictsResponseBody
	GetSynchro() *bool
}

type ListAppTemplateDictsResponseBody struct {
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
	MaxResults *int32                                    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	Module     []*ListAppTemplateDictsResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Repeated"`
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

func (s ListAppTemplateDictsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAppTemplateDictsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAppTemplateDictsResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ListAppTemplateDictsResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *ListAppTemplateDictsResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *ListAppTemplateDictsResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *ListAppTemplateDictsResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *ListAppTemplateDictsResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *ListAppTemplateDictsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAppTemplateDictsResponseBody) GetModule() []*ListAppTemplateDictsResponseBodyModule {
	return s.Module
}

func (s *ListAppTemplateDictsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAppTemplateDictsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAppTemplateDictsResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *ListAppTemplateDictsResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *ListAppTemplateDictsResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *ListAppTemplateDictsResponseBody) SetAccessDeniedDetail(v string) *ListAppTemplateDictsResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListAppTemplateDictsResponseBody) SetAllowRetry(v bool) *ListAppTemplateDictsResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *ListAppTemplateDictsResponseBody) SetAppName(v string) *ListAppTemplateDictsResponseBody {
	s.AppName = &v
	return s
}

func (s *ListAppTemplateDictsResponseBody) SetDynamicCode(v string) *ListAppTemplateDictsResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *ListAppTemplateDictsResponseBody) SetDynamicMessage(v string) *ListAppTemplateDictsResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ListAppTemplateDictsResponseBody) SetErrorArgs(v []interface{}) *ListAppTemplateDictsResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *ListAppTemplateDictsResponseBody) SetMaxResults(v int32) *ListAppTemplateDictsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListAppTemplateDictsResponseBody) SetModule(v []*ListAppTemplateDictsResponseBodyModule) *ListAppTemplateDictsResponseBody {
	s.Module = v
	return s
}

func (s *ListAppTemplateDictsResponseBody) SetNextToken(v string) *ListAppTemplateDictsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListAppTemplateDictsResponseBody) SetRequestId(v string) *ListAppTemplateDictsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAppTemplateDictsResponseBody) SetRootErrorCode(v string) *ListAppTemplateDictsResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *ListAppTemplateDictsResponseBody) SetRootErrorMsg(v string) *ListAppTemplateDictsResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *ListAppTemplateDictsResponseBody) SetSynchro(v bool) *ListAppTemplateDictsResponseBody {
	s.Synchro = &v
	return s
}

func (s *ListAppTemplateDictsResponseBody) Validate() error {
	if s.Module != nil {
		for _, item := range s.Module {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAppTemplateDictsResponseBodyModule struct {
	// example:
	//
	// abc
	DictCode *string `json:"DictCode,omitempty" xml:"DictCode,omitempty"`
	// example:
	//
	// label
	DictLabel *string `json:"DictLabel,omitempty" xml:"DictLabel,omitempty"`
	// example:
	//
	// product_version
	DictType *string `json:"DictType,omitempty" xml:"DictType,omitempty"`
	// example:
	//
	// abc
	DictValue    *string `json:"DictValue,omitempty" xml:"DictValue,omitempty"`
	HasTemplates *bool   `json:"HasTemplates,omitempty" xml:"HasTemplates,omitempty"`
	// example:
	//
	// desc
	SortOrder *int32 `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
}

func (s ListAppTemplateDictsResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s ListAppTemplateDictsResponseBodyModule) GoString() string {
	return s.String()
}

func (s *ListAppTemplateDictsResponseBodyModule) GetDictCode() *string {
	return s.DictCode
}

func (s *ListAppTemplateDictsResponseBodyModule) GetDictLabel() *string {
	return s.DictLabel
}

func (s *ListAppTemplateDictsResponseBodyModule) GetDictType() *string {
	return s.DictType
}

func (s *ListAppTemplateDictsResponseBodyModule) GetDictValue() *string {
	return s.DictValue
}

func (s *ListAppTemplateDictsResponseBodyModule) GetHasTemplates() *bool {
	return s.HasTemplates
}

func (s *ListAppTemplateDictsResponseBodyModule) GetSortOrder() *int32 {
	return s.SortOrder
}

func (s *ListAppTemplateDictsResponseBodyModule) SetDictCode(v string) *ListAppTemplateDictsResponseBodyModule {
	s.DictCode = &v
	return s
}

func (s *ListAppTemplateDictsResponseBodyModule) SetDictLabel(v string) *ListAppTemplateDictsResponseBodyModule {
	s.DictLabel = &v
	return s
}

func (s *ListAppTemplateDictsResponseBodyModule) SetDictType(v string) *ListAppTemplateDictsResponseBodyModule {
	s.DictType = &v
	return s
}

func (s *ListAppTemplateDictsResponseBodyModule) SetDictValue(v string) *ListAppTemplateDictsResponseBodyModule {
	s.DictValue = &v
	return s
}

func (s *ListAppTemplateDictsResponseBodyModule) SetHasTemplates(v bool) *ListAppTemplateDictsResponseBodyModule {
	s.HasTemplates = &v
	return s
}

func (s *ListAppTemplateDictsResponseBodyModule) SetSortOrder(v int32) *ListAppTemplateDictsResponseBodyModule {
	s.SortOrder = &v
	return s
}

func (s *ListAppTemplateDictsResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
