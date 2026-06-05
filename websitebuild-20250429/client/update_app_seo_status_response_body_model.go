// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAppSeoStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *UpdateAppSeoStatusResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *UpdateAppSeoStatusResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *UpdateAppSeoStatusResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *UpdateAppSeoStatusResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *UpdateAppSeoStatusResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *UpdateAppSeoStatusResponseBody
	GetErrorArgs() []interface{}
	SetModule(v bool) *UpdateAppSeoStatusResponseBody
	GetModule() *bool
	SetRequestId(v string) *UpdateAppSeoStatusResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *UpdateAppSeoStatusResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *UpdateAppSeoStatusResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *UpdateAppSeoStatusResponseBody
	GetSynchro() *bool
}

type UpdateAppSeoStatusResponseBody struct {
	// Detailed reason for access denial.
	//
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// Whether retry is allowed
	//
	// example:
	//
	// False
	AllowRetry *bool `json:"AllowRetry,omitempty" xml:"AllowRetry,omitempty"`
	// Application name
	//
	// example:
	//
	// spring-cloud-b
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// Dynamic error code.
	//
	// example:
	//
	// ERROR-oo1
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// Dynamic error message, used to replace the `%s` in the **ErrMessage*	- error message.
	//
	// > If **ErrMessage*	- returns **The Value of Input Parameter %s is not valid**, and **DynamicMessage*	- returns **DtsJobId**, it indicates that the input request parameter **DtsJobId*	- is invalid.
	//
	// example:
	//
	// abc
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// Returned error parameters
	ErrorArgs []interface{} `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	// Response data
	//
	// example:
	//
	// true
	Module *bool `json:"Module,omitempty" xml:"Module,omitempty"`
	// ID of the request
	//
	// example:
	//
	// 6C6B99AC-39EC-5350-874C-204128C905E6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Error code
	//
	// example:
	//
	// SYSTEM.ERROR
	RootErrorCode *string `json:"RootErrorCode,omitempty" xml:"RootErrorCode,omitempty"`
	// Exception message
	//
	// example:
	//
	// 系统异常
	RootErrorMsg *string `json:"RootErrorMsg,omitempty" xml:"RootErrorMsg,omitempty"`
	// Reserved parameter.
	//
	// example:
	//
	// True
	Synchro *bool `json:"Synchro,omitempty" xml:"Synchro,omitempty"`
}

func (s UpdateAppSeoStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAppSeoStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAppSeoStatusResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *UpdateAppSeoStatusResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *UpdateAppSeoStatusResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *UpdateAppSeoStatusResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *UpdateAppSeoStatusResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *UpdateAppSeoStatusResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *UpdateAppSeoStatusResponseBody) GetModule() *bool {
	return s.Module
}

func (s *UpdateAppSeoStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAppSeoStatusResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *UpdateAppSeoStatusResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *UpdateAppSeoStatusResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *UpdateAppSeoStatusResponseBody) SetAccessDeniedDetail(v string) *UpdateAppSeoStatusResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *UpdateAppSeoStatusResponseBody) SetAllowRetry(v bool) *UpdateAppSeoStatusResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *UpdateAppSeoStatusResponseBody) SetAppName(v string) *UpdateAppSeoStatusResponseBody {
	s.AppName = &v
	return s
}

func (s *UpdateAppSeoStatusResponseBody) SetDynamicCode(v string) *UpdateAppSeoStatusResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *UpdateAppSeoStatusResponseBody) SetDynamicMessage(v string) *UpdateAppSeoStatusResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *UpdateAppSeoStatusResponseBody) SetErrorArgs(v []interface{}) *UpdateAppSeoStatusResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *UpdateAppSeoStatusResponseBody) SetModule(v bool) *UpdateAppSeoStatusResponseBody {
	s.Module = &v
	return s
}

func (s *UpdateAppSeoStatusResponseBody) SetRequestId(v string) *UpdateAppSeoStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAppSeoStatusResponseBody) SetRootErrorCode(v string) *UpdateAppSeoStatusResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *UpdateAppSeoStatusResponseBody) SetRootErrorMsg(v string) *UpdateAppSeoStatusResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *UpdateAppSeoStatusResponseBody) SetSynchro(v bool) *UpdateAppSeoStatusResponseBody {
	s.Synchro = &v
	return s
}

func (s *UpdateAppSeoStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
