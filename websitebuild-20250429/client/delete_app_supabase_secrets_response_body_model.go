// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAppSupabaseSecretsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DeleteAppSupabaseSecretsResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *DeleteAppSupabaseSecretsResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *DeleteAppSupabaseSecretsResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *DeleteAppSupabaseSecretsResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *DeleteAppSupabaseSecretsResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *DeleteAppSupabaseSecretsResponseBody
	GetErrorArgs() []interface{}
	SetModule(v bool) *DeleteAppSupabaseSecretsResponseBody
	GetModule() *bool
	SetRequestId(v string) *DeleteAppSupabaseSecretsResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *DeleteAppSupabaseSecretsResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *DeleteAppSupabaseSecretsResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *DeleteAppSupabaseSecretsResponseBody
	GetSynchro() *bool
}

type DeleteAppSupabaseSecretsResponseBody struct {
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
	// App Name.
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
	// Dynamic message.
	//
	// example:
	//
	// SYSTEM_ERROR
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// Returned error parameters
	ErrorArgs []interface{} `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	// Whether the shift succeeded
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

func (s DeleteAppSupabaseSecretsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAppSupabaseSecretsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAppSupabaseSecretsResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DeleteAppSupabaseSecretsResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *DeleteAppSupabaseSecretsResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *DeleteAppSupabaseSecretsResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *DeleteAppSupabaseSecretsResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *DeleteAppSupabaseSecretsResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *DeleteAppSupabaseSecretsResponseBody) GetModule() *bool {
	return s.Module
}

func (s *DeleteAppSupabaseSecretsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAppSupabaseSecretsResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *DeleteAppSupabaseSecretsResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *DeleteAppSupabaseSecretsResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *DeleteAppSupabaseSecretsResponseBody) SetAccessDeniedDetail(v string) *DeleteAppSupabaseSecretsResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DeleteAppSupabaseSecretsResponseBody) SetAllowRetry(v bool) *DeleteAppSupabaseSecretsResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *DeleteAppSupabaseSecretsResponseBody) SetAppName(v string) *DeleteAppSupabaseSecretsResponseBody {
	s.AppName = &v
	return s
}

func (s *DeleteAppSupabaseSecretsResponseBody) SetDynamicCode(v string) *DeleteAppSupabaseSecretsResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DeleteAppSupabaseSecretsResponseBody) SetDynamicMessage(v string) *DeleteAppSupabaseSecretsResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DeleteAppSupabaseSecretsResponseBody) SetErrorArgs(v []interface{}) *DeleteAppSupabaseSecretsResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *DeleteAppSupabaseSecretsResponseBody) SetModule(v bool) *DeleteAppSupabaseSecretsResponseBody {
	s.Module = &v
	return s
}

func (s *DeleteAppSupabaseSecretsResponseBody) SetRequestId(v string) *DeleteAppSupabaseSecretsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAppSupabaseSecretsResponseBody) SetRootErrorCode(v string) *DeleteAppSupabaseSecretsResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *DeleteAppSupabaseSecretsResponseBody) SetRootErrorMsg(v string) *DeleteAppSupabaseSecretsResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *DeleteAppSupabaseSecretsResponseBody) SetSynchro(v bool) *DeleteAppSupabaseSecretsResponseBody {
	s.Synchro = &v
	return s
}

func (s *DeleteAppSupabaseSecretsResponseBody) Validate() error {
	return dara.Validate(s)
}
