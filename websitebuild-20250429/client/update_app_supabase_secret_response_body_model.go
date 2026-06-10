// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAppSupabaseSecretResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *UpdateAppSupabaseSecretResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *UpdateAppSupabaseSecretResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *UpdateAppSupabaseSecretResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *UpdateAppSupabaseSecretResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *UpdateAppSupabaseSecretResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *UpdateAppSupabaseSecretResponseBody
	GetErrorArgs() []interface{}
	SetModule(v bool) *UpdateAppSupabaseSecretResponseBody
	GetModule() *bool
	SetRequestId(v string) *UpdateAppSupabaseSecretResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *UpdateAppSupabaseSecretResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *UpdateAppSupabaseSecretResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *UpdateAppSupabaseSecretResponseBody
	GetSynchro() *bool
}

type UpdateAppSupabaseSecretResponseBody struct {
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
	// App name.
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
	// Dynamic error message used to replace the `%s` placeholder in the **ErrMessage*	- error message.
	//
	// > If **ErrMessage*	- returns **The Value of Input Parameter %s is not valid*	- and **DynamicMessage*	- returns **DtsJobId**, it indicates that the provided request parameter **DtsJobId*	- is invalid.
	//
	// example:
	//
	// SYSTEM_ERROR
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// Fault parameters.
	ErrorArgs []interface{} `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	// Response data
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
	// Whether processing is synchronous
	//
	// example:
	//
	// True
	Synchro *bool `json:"Synchro,omitempty" xml:"Synchro,omitempty"`
}

func (s UpdateAppSupabaseSecretResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAppSupabaseSecretResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAppSupabaseSecretResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *UpdateAppSupabaseSecretResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *UpdateAppSupabaseSecretResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *UpdateAppSupabaseSecretResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *UpdateAppSupabaseSecretResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *UpdateAppSupabaseSecretResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *UpdateAppSupabaseSecretResponseBody) GetModule() *bool {
	return s.Module
}

func (s *UpdateAppSupabaseSecretResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAppSupabaseSecretResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *UpdateAppSupabaseSecretResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *UpdateAppSupabaseSecretResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *UpdateAppSupabaseSecretResponseBody) SetAccessDeniedDetail(v string) *UpdateAppSupabaseSecretResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *UpdateAppSupabaseSecretResponseBody) SetAllowRetry(v bool) *UpdateAppSupabaseSecretResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *UpdateAppSupabaseSecretResponseBody) SetAppName(v string) *UpdateAppSupabaseSecretResponseBody {
	s.AppName = &v
	return s
}

func (s *UpdateAppSupabaseSecretResponseBody) SetDynamicCode(v string) *UpdateAppSupabaseSecretResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *UpdateAppSupabaseSecretResponseBody) SetDynamicMessage(v string) *UpdateAppSupabaseSecretResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *UpdateAppSupabaseSecretResponseBody) SetErrorArgs(v []interface{}) *UpdateAppSupabaseSecretResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *UpdateAppSupabaseSecretResponseBody) SetModule(v bool) *UpdateAppSupabaseSecretResponseBody {
	s.Module = &v
	return s
}

func (s *UpdateAppSupabaseSecretResponseBody) SetRequestId(v string) *UpdateAppSupabaseSecretResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAppSupabaseSecretResponseBody) SetRootErrorCode(v string) *UpdateAppSupabaseSecretResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *UpdateAppSupabaseSecretResponseBody) SetRootErrorMsg(v string) *UpdateAppSupabaseSecretResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *UpdateAppSupabaseSecretResponseBody) SetSynchro(v bool) *UpdateAppSupabaseSecretResponseBody {
	s.Synchro = &v
	return s
}

func (s *UpdateAppSupabaseSecretResponseBody) Validate() error {
	return dara.Validate(s)
}
