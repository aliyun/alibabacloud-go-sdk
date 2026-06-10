// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAppSupabaseAuthConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *UpdateAppSupabaseAuthConfigResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *UpdateAppSupabaseAuthConfigResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *UpdateAppSupabaseAuthConfigResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *UpdateAppSupabaseAuthConfigResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *UpdateAppSupabaseAuthConfigResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *UpdateAppSupabaseAuthConfigResponseBody
	GetErrorArgs() []interface{}
	SetModule(v bool) *UpdateAppSupabaseAuthConfigResponseBody
	GetModule() *bool
	SetRequestId(v string) *UpdateAppSupabaseAuthConfigResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *UpdateAppSupabaseAuthConfigResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *UpdateAppSupabaseAuthConfigResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *UpdateAppSupabaseAuthConfigResponseBody
	GetSynchro() *bool
}

type UpdateAppSupabaseAuthConfigResponseBody struct {
	// Permission information unavailable
	//
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// Indicates whether retry is allowed
	//
	// example:
	//
	// False
	AllowRetry *bool `json:"AllowRetry,omitempty" xml:"AllowRetry,omitempty"`
	// App Name.
	//
	// example:
	//
	// spring-cloud-b
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// Dynamic code. Not currently used. Ignore it.
	//
	// example:
	//
	// ERROR-oo1
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// Dynamic message.
	//
	// example:
	//
	// SYSTEM_ERROR
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// Returned error parameters
	ErrorArgs []interface{} `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	// Task object
	//
	// example:
	//
	// true
	Module *bool `json:"Module,omitempty" xml:"Module,omitempty"`
	// Id of the request
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
	// Abnormal message
	//
	// example:
	//
	// 系统异常
	RootErrorMsg *string `json:"RootErrorMsg,omitempty" xml:"RootErrorMsg,omitempty"`
	// Indicates whether processing is synchronous
	//
	// example:
	//
	// True
	Synchro *bool `json:"Synchro,omitempty" xml:"Synchro,omitempty"`
}

func (s UpdateAppSupabaseAuthConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAppSupabaseAuthConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAppSupabaseAuthConfigResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *UpdateAppSupabaseAuthConfigResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *UpdateAppSupabaseAuthConfigResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *UpdateAppSupabaseAuthConfigResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *UpdateAppSupabaseAuthConfigResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *UpdateAppSupabaseAuthConfigResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *UpdateAppSupabaseAuthConfigResponseBody) GetModule() *bool {
	return s.Module
}

func (s *UpdateAppSupabaseAuthConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAppSupabaseAuthConfigResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *UpdateAppSupabaseAuthConfigResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *UpdateAppSupabaseAuthConfigResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *UpdateAppSupabaseAuthConfigResponseBody) SetAccessDeniedDetail(v string) *UpdateAppSupabaseAuthConfigResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *UpdateAppSupabaseAuthConfigResponseBody) SetAllowRetry(v bool) *UpdateAppSupabaseAuthConfigResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *UpdateAppSupabaseAuthConfigResponseBody) SetAppName(v string) *UpdateAppSupabaseAuthConfigResponseBody {
	s.AppName = &v
	return s
}

func (s *UpdateAppSupabaseAuthConfigResponseBody) SetDynamicCode(v string) *UpdateAppSupabaseAuthConfigResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *UpdateAppSupabaseAuthConfigResponseBody) SetDynamicMessage(v string) *UpdateAppSupabaseAuthConfigResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *UpdateAppSupabaseAuthConfigResponseBody) SetErrorArgs(v []interface{}) *UpdateAppSupabaseAuthConfigResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *UpdateAppSupabaseAuthConfigResponseBody) SetModule(v bool) *UpdateAppSupabaseAuthConfigResponseBody {
	s.Module = &v
	return s
}

func (s *UpdateAppSupabaseAuthConfigResponseBody) SetRequestId(v string) *UpdateAppSupabaseAuthConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAppSupabaseAuthConfigResponseBody) SetRootErrorCode(v string) *UpdateAppSupabaseAuthConfigResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *UpdateAppSupabaseAuthConfigResponseBody) SetRootErrorMsg(v string) *UpdateAppSupabaseAuthConfigResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *UpdateAppSupabaseAuthConfigResponseBody) SetSynchro(v bool) *UpdateAppSupabaseAuthConfigResponseBody {
	s.Synchro = &v
	return s
}

func (s *UpdateAppSupabaseAuthConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
