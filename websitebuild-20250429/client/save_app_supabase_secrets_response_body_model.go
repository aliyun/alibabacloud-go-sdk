// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveAppSupabaseSecretsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *SaveAppSupabaseSecretsResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *SaveAppSupabaseSecretsResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *SaveAppSupabaseSecretsResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *SaveAppSupabaseSecretsResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *SaveAppSupabaseSecretsResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *SaveAppSupabaseSecretsResponseBody
	GetErrorArgs() []interface{}
	SetModule(v bool) *SaveAppSupabaseSecretsResponseBody
	GetModule() *bool
	SetRequestId(v string) *SaveAppSupabaseSecretsResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *SaveAppSupabaseSecretsResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *SaveAppSupabaseSecretsResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *SaveAppSupabaseSecretsResponseBody
	GetSynchro() *bool
}

type SaveAppSupabaseSecretsResponseBody struct {
	// Detailed reason for access denial.
	//
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// Is retry allowed
	//
	// example:
	//
	// False
	AllowRetry *bool `json:"AllowRetry,omitempty" xml:"AllowRetry,omitempty"`
	// App Name.
	//
	// example:
	//
	// or
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// Dynamic error code.
	//
	// example:
	//
	// ERROR-oo1
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// Dynamic error message, used to replace `%s` in the **ErrMessage*	- error message.
	//
	// > If **ErrMessage*	- returns **The Value of Input Parameter %s is not valid*	- and **DynamicMessage*	- returns **DtsJobId**, it indicates that the provided request parameter **DtsJobId*	- is invalid.
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
	// SYSTEM.EROR
	RootErrorCode *string `json:"RootErrorCode,omitempty" xml:"RootErrorCode,omitempty"`
	// Abnormal message
	//
	// example:
	//
	// 系统异常
	RootErrorMsg *string `json:"RootErrorMsg,omitempty" xml:"RootErrorMsg,omitempty"`
	// Is processed synchronously
	//
	// example:
	//
	// True
	Synchro *bool `json:"Synchro,omitempty" xml:"Synchro,omitempty"`
}

func (s SaveAppSupabaseSecretsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveAppSupabaseSecretsResponseBody) GoString() string {
	return s.String()
}

func (s *SaveAppSupabaseSecretsResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *SaveAppSupabaseSecretsResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *SaveAppSupabaseSecretsResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *SaveAppSupabaseSecretsResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *SaveAppSupabaseSecretsResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *SaveAppSupabaseSecretsResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *SaveAppSupabaseSecretsResponseBody) GetModule() *bool {
	return s.Module
}

func (s *SaveAppSupabaseSecretsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveAppSupabaseSecretsResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *SaveAppSupabaseSecretsResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *SaveAppSupabaseSecretsResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *SaveAppSupabaseSecretsResponseBody) SetAccessDeniedDetail(v string) *SaveAppSupabaseSecretsResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *SaveAppSupabaseSecretsResponseBody) SetAllowRetry(v bool) *SaveAppSupabaseSecretsResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *SaveAppSupabaseSecretsResponseBody) SetAppName(v string) *SaveAppSupabaseSecretsResponseBody {
	s.AppName = &v
	return s
}

func (s *SaveAppSupabaseSecretsResponseBody) SetDynamicCode(v string) *SaveAppSupabaseSecretsResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *SaveAppSupabaseSecretsResponseBody) SetDynamicMessage(v string) *SaveAppSupabaseSecretsResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *SaveAppSupabaseSecretsResponseBody) SetErrorArgs(v []interface{}) *SaveAppSupabaseSecretsResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *SaveAppSupabaseSecretsResponseBody) SetModule(v bool) *SaveAppSupabaseSecretsResponseBody {
	s.Module = &v
	return s
}

func (s *SaveAppSupabaseSecretsResponseBody) SetRequestId(v string) *SaveAppSupabaseSecretsResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveAppSupabaseSecretsResponseBody) SetRootErrorCode(v string) *SaveAppSupabaseSecretsResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *SaveAppSupabaseSecretsResponseBody) SetRootErrorMsg(v string) *SaveAppSupabaseSecretsResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *SaveAppSupabaseSecretsResponseBody) SetSynchro(v bool) *SaveAppSupabaseSecretsResponseBody {
	s.Synchro = &v
	return s
}

func (s *SaveAppSupabaseSecretsResponseBody) Validate() error {
	return dara.Validate(s)
}
