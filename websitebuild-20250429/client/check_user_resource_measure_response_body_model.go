// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckUserResourceMeasureResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CheckUserResourceMeasureResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *CheckUserResourceMeasureResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *CheckUserResourceMeasureResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *CheckUserResourceMeasureResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *CheckUserResourceMeasureResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *CheckUserResourceMeasureResponseBody
	GetErrorArgs() []interface{}
	SetModule(v *CheckUserResourceMeasureResponseBodyModule) *CheckUserResourceMeasureResponseBody
	GetModule() *CheckUserResourceMeasureResponseBodyModule
	SetRequestId(v string) *CheckUserResourceMeasureResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *CheckUserResourceMeasureResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *CheckUserResourceMeasureResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *CheckUserResourceMeasureResponseBody
	GetSynchro() *bool
}

type CheckUserResourceMeasureResponseBody struct {
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
	DynamicMessage *string                                     `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorArgs      []interface{}                               `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	Module         *CheckUserResourceMeasureResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
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

func (s CheckUserResourceMeasureResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckUserResourceMeasureResponseBody) GoString() string {
	return s.String()
}

func (s *CheckUserResourceMeasureResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CheckUserResourceMeasureResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *CheckUserResourceMeasureResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *CheckUserResourceMeasureResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *CheckUserResourceMeasureResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *CheckUserResourceMeasureResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *CheckUserResourceMeasureResponseBody) GetModule() *CheckUserResourceMeasureResponseBodyModule {
	return s.Module
}

func (s *CheckUserResourceMeasureResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckUserResourceMeasureResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *CheckUserResourceMeasureResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *CheckUserResourceMeasureResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *CheckUserResourceMeasureResponseBody) SetAccessDeniedDetail(v string) *CheckUserResourceMeasureResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CheckUserResourceMeasureResponseBody) SetAllowRetry(v bool) *CheckUserResourceMeasureResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *CheckUserResourceMeasureResponseBody) SetAppName(v string) *CheckUserResourceMeasureResponseBody {
	s.AppName = &v
	return s
}

func (s *CheckUserResourceMeasureResponseBody) SetDynamicCode(v string) *CheckUserResourceMeasureResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *CheckUserResourceMeasureResponseBody) SetDynamicMessage(v string) *CheckUserResourceMeasureResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *CheckUserResourceMeasureResponseBody) SetErrorArgs(v []interface{}) *CheckUserResourceMeasureResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *CheckUserResourceMeasureResponseBody) SetModule(v *CheckUserResourceMeasureResponseBodyModule) *CheckUserResourceMeasureResponseBody {
	s.Module = v
	return s
}

func (s *CheckUserResourceMeasureResponseBody) SetRequestId(v string) *CheckUserResourceMeasureResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckUserResourceMeasureResponseBody) SetRootErrorCode(v string) *CheckUserResourceMeasureResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *CheckUserResourceMeasureResponseBody) SetRootErrorMsg(v string) *CheckUserResourceMeasureResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *CheckUserResourceMeasureResponseBody) SetSynchro(v bool) *CheckUserResourceMeasureResponseBody {
	s.Synchro = &v
	return s
}

func (s *CheckUserResourceMeasureResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CheckUserResourceMeasureResponseBodyModule struct {
	// example:
	//
	// Success
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// stream push failed
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// False
	Passed *bool `json:"Passed,omitempty" xml:"Passed,omitempty"`
	// example:
	//
	// TransitRouterVpcAttachment
	ResourceCode *string `json:"ResourceCode,omitempty" xml:"ResourceCode,omitempty"`
}

func (s CheckUserResourceMeasureResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s CheckUserResourceMeasureResponseBodyModule) GoString() string {
	return s.String()
}

func (s *CheckUserResourceMeasureResponseBodyModule) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CheckUserResourceMeasureResponseBodyModule) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CheckUserResourceMeasureResponseBodyModule) GetPassed() *bool {
	return s.Passed
}

func (s *CheckUserResourceMeasureResponseBodyModule) GetResourceCode() *string {
	return s.ResourceCode
}

func (s *CheckUserResourceMeasureResponseBodyModule) SetErrorCode(v string) *CheckUserResourceMeasureResponseBodyModule {
	s.ErrorCode = &v
	return s
}

func (s *CheckUserResourceMeasureResponseBodyModule) SetErrorMessage(v string) *CheckUserResourceMeasureResponseBodyModule {
	s.ErrorMessage = &v
	return s
}

func (s *CheckUserResourceMeasureResponseBodyModule) SetPassed(v bool) *CheckUserResourceMeasureResponseBodyModule {
	s.Passed = &v
	return s
}

func (s *CheckUserResourceMeasureResponseBodyModule) SetResourceCode(v string) *CheckUserResourceMeasureResponseBodyModule {
	s.ResourceCode = &v
	return s
}

func (s *CheckUserResourceMeasureResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
