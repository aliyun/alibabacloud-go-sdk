// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckResourceMeasureResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CheckResourceMeasureResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *CheckResourceMeasureResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *CheckResourceMeasureResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *CheckResourceMeasureResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *CheckResourceMeasureResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *CheckResourceMeasureResponseBody
	GetErrorArgs() []interface{}
	SetModule(v *CheckResourceMeasureResponseBodyModule) *CheckResourceMeasureResponseBody
	GetModule() *CheckResourceMeasureResponseBodyModule
	SetRequestId(v string) *CheckResourceMeasureResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *CheckResourceMeasureResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *CheckResourceMeasureResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *CheckResourceMeasureResponseBody
	GetSynchro() *bool
}

type CheckResourceMeasureResponseBody struct {
	// The detailed reason why access was denied.
	//
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// Indicates whether retry is allowed.
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
	// The dynamic error message, which is used to replace the `%s` placeholder in the **ErrMessage*	- response element.
	//
	// > If **ErrMessage*	- returns **The Value of Input Parameter %s is not valid*	- and **DynamicMessage*	- returns **DtsJobId**, the request parameter **DtsJobId*	- is invalid.
	//
	// example:
	//
	// SYSTEM_ERROR
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The error parameters.
	ErrorArgs []interface{} `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	// The application module.
	Module *CheckResourceMeasureResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
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
	// A reserved parameter.
	//
	// example:
	//
	// True
	Synchro *bool `json:"Synchro,omitempty" xml:"Synchro,omitempty"`
}

func (s CheckResourceMeasureResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckResourceMeasureResponseBody) GoString() string {
	return s.String()
}

func (s *CheckResourceMeasureResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CheckResourceMeasureResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *CheckResourceMeasureResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *CheckResourceMeasureResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *CheckResourceMeasureResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *CheckResourceMeasureResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *CheckResourceMeasureResponseBody) GetModule() *CheckResourceMeasureResponseBodyModule {
	return s.Module
}

func (s *CheckResourceMeasureResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckResourceMeasureResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *CheckResourceMeasureResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *CheckResourceMeasureResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *CheckResourceMeasureResponseBody) SetAccessDeniedDetail(v string) *CheckResourceMeasureResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CheckResourceMeasureResponseBody) SetAllowRetry(v bool) *CheckResourceMeasureResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *CheckResourceMeasureResponseBody) SetAppName(v string) *CheckResourceMeasureResponseBody {
	s.AppName = &v
	return s
}

func (s *CheckResourceMeasureResponseBody) SetDynamicCode(v string) *CheckResourceMeasureResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *CheckResourceMeasureResponseBody) SetDynamicMessage(v string) *CheckResourceMeasureResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *CheckResourceMeasureResponseBody) SetErrorArgs(v []interface{}) *CheckResourceMeasureResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *CheckResourceMeasureResponseBody) SetModule(v *CheckResourceMeasureResponseBodyModule) *CheckResourceMeasureResponseBody {
	s.Module = v
	return s
}

func (s *CheckResourceMeasureResponseBody) SetRequestId(v string) *CheckResourceMeasureResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckResourceMeasureResponseBody) SetRootErrorCode(v string) *CheckResourceMeasureResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *CheckResourceMeasureResponseBody) SetRootErrorMsg(v string) *CheckResourceMeasureResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *CheckResourceMeasureResponseBody) SetSynchro(v bool) *CheckResourceMeasureResponseBody {
	s.Synchro = &v
	return s
}

func (s *CheckResourceMeasureResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CheckResourceMeasureResponseBodyModule struct {
	// The error code when the check fails. This value is null when the check passes.
	//
	// example:
	//
	// Success
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message when the check fails. This value is null when the check passes.
	//
	// example:
	//
	// stream push failed
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// Indicates whether the check is passed.
	//
	// example:
	//
	// False
	Passed *bool `json:"Passed,omitempty" xml:"Passed,omitempty"`
	// The resource identifier.
	//
	// example:
	//
	// TransitRouterVpcAttachment
	ResourceCode *string `json:"ResourceCode,omitempty" xml:"ResourceCode,omitempty"`
}

func (s CheckResourceMeasureResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s CheckResourceMeasureResponseBodyModule) GoString() string {
	return s.String()
}

func (s *CheckResourceMeasureResponseBodyModule) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CheckResourceMeasureResponseBodyModule) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CheckResourceMeasureResponseBodyModule) GetPassed() *bool {
	return s.Passed
}

func (s *CheckResourceMeasureResponseBodyModule) GetResourceCode() *string {
	return s.ResourceCode
}

func (s *CheckResourceMeasureResponseBodyModule) SetErrorCode(v string) *CheckResourceMeasureResponseBodyModule {
	s.ErrorCode = &v
	return s
}

func (s *CheckResourceMeasureResponseBodyModule) SetErrorMessage(v string) *CheckResourceMeasureResponseBodyModule {
	s.ErrorMessage = &v
	return s
}

func (s *CheckResourceMeasureResponseBodyModule) SetPassed(v bool) *CheckResourceMeasureResponseBodyModule {
	s.Passed = &v
	return s
}

func (s *CheckResourceMeasureResponseBodyModule) SetResourceCode(v string) *CheckResourceMeasureResponseBodyModule {
	s.ResourceCode = &v
	return s
}

func (s *CheckResourceMeasureResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
