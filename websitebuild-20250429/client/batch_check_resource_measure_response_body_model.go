// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchCheckResourceMeasureResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *BatchCheckResourceMeasureResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *BatchCheckResourceMeasureResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *BatchCheckResourceMeasureResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *BatchCheckResourceMeasureResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *BatchCheckResourceMeasureResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *BatchCheckResourceMeasureResponseBody
	GetErrorArgs() []interface{}
	SetModule(v *BatchCheckResourceMeasureResponseBodyModule) *BatchCheckResourceMeasureResponseBody
	GetModule() *BatchCheckResourceMeasureResponseBodyModule
	SetRequestId(v string) *BatchCheckResourceMeasureResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *BatchCheckResourceMeasureResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *BatchCheckResourceMeasureResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *BatchCheckResourceMeasureResponseBody
	GetSynchro() *bool
}

type BatchCheckResourceMeasureResponseBody struct {
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
	// watermark
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// ERROR-oo1
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// example:
	//
	// SYSTEM_ERROR
	DynamicMessage *string                                      `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorArgs      []interface{}                                `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	Module         *BatchCheckResourceMeasureResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
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

func (s BatchCheckResourceMeasureResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchCheckResourceMeasureResponseBody) GoString() string {
	return s.String()
}

func (s *BatchCheckResourceMeasureResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *BatchCheckResourceMeasureResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *BatchCheckResourceMeasureResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *BatchCheckResourceMeasureResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *BatchCheckResourceMeasureResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *BatchCheckResourceMeasureResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *BatchCheckResourceMeasureResponseBody) GetModule() *BatchCheckResourceMeasureResponseBodyModule {
	return s.Module
}

func (s *BatchCheckResourceMeasureResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchCheckResourceMeasureResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *BatchCheckResourceMeasureResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *BatchCheckResourceMeasureResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *BatchCheckResourceMeasureResponseBody) SetAccessDeniedDetail(v string) *BatchCheckResourceMeasureResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *BatchCheckResourceMeasureResponseBody) SetAllowRetry(v bool) *BatchCheckResourceMeasureResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *BatchCheckResourceMeasureResponseBody) SetAppName(v string) *BatchCheckResourceMeasureResponseBody {
	s.AppName = &v
	return s
}

func (s *BatchCheckResourceMeasureResponseBody) SetDynamicCode(v string) *BatchCheckResourceMeasureResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *BatchCheckResourceMeasureResponseBody) SetDynamicMessage(v string) *BatchCheckResourceMeasureResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *BatchCheckResourceMeasureResponseBody) SetErrorArgs(v []interface{}) *BatchCheckResourceMeasureResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *BatchCheckResourceMeasureResponseBody) SetModule(v *BatchCheckResourceMeasureResponseBodyModule) *BatchCheckResourceMeasureResponseBody {
	s.Module = v
	return s
}

func (s *BatchCheckResourceMeasureResponseBody) SetRequestId(v string) *BatchCheckResourceMeasureResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchCheckResourceMeasureResponseBody) SetRootErrorCode(v string) *BatchCheckResourceMeasureResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *BatchCheckResourceMeasureResponseBody) SetRootErrorMsg(v string) *BatchCheckResourceMeasureResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *BatchCheckResourceMeasureResponseBody) SetSynchro(v bool) *BatchCheckResourceMeasureResponseBody {
	s.Synchro = &v
	return s
}

func (s *BatchCheckResourceMeasureResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BatchCheckResourceMeasureResponseBodyModule struct {
	AllPassed *bool                          `json:"AllPassed,omitempty" xml:"AllPassed,omitempty"`
	Results   map[string]*ModuleResultsValue `json:"Results,omitempty" xml:"Results,omitempty"`
}

func (s BatchCheckResourceMeasureResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s BatchCheckResourceMeasureResponseBodyModule) GoString() string {
	return s.String()
}

func (s *BatchCheckResourceMeasureResponseBodyModule) GetAllPassed() *bool {
	return s.AllPassed
}

func (s *BatchCheckResourceMeasureResponseBodyModule) GetResults() map[string]*ModuleResultsValue {
	return s.Results
}

func (s *BatchCheckResourceMeasureResponseBodyModule) SetAllPassed(v bool) *BatchCheckResourceMeasureResponseBodyModule {
	s.AllPassed = &v
	return s
}

func (s *BatchCheckResourceMeasureResponseBodyModule) SetResults(v map[string]*ModuleResultsValue) *BatchCheckResourceMeasureResponseBodyModule {
	s.Results = v
	return s
}

func (s *BatchCheckResourceMeasureResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
