// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppSandboxPreviewUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetAppSandboxPreviewUrlResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *GetAppSandboxPreviewUrlResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *GetAppSandboxPreviewUrlResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *GetAppSandboxPreviewUrlResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *GetAppSandboxPreviewUrlResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *GetAppSandboxPreviewUrlResponseBody
	GetErrorArgs() []interface{}
	SetModule(v interface{}) *GetAppSandboxPreviewUrlResponseBody
	GetModule() interface{}
	SetRequestId(v string) *GetAppSandboxPreviewUrlResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *GetAppSandboxPreviewUrlResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *GetAppSandboxPreviewUrlResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *GetAppSandboxPreviewUrlResponseBody
	GetSynchro() *bool
}

type GetAppSandboxPreviewUrlResponseBody struct {
	// The detailed cause of access denial.
	//
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// Indicates whether retries are allowed.
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
	// The dynamic code. This field is not in use and can be ignored.
	//
	// example:
	//
	// ERROR-oo1
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// The dynamic error message, which is used to replace `%s` in the error message returned by the **ErrMessage*	- parameter. If **ErrMessage*	- returns **The Value of Input Parameter %s is not valid*	- and **DynamicMessage*	- returns **DtsJobId**, the request parameter **DtsJobId*	- is invalid.
	//
	// example:
	//
	// SYSTEM_ERROR
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The returned error parameters.
	ErrorArgs []interface{} `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	// The returned object.
	//
	// example:
	//
	// true
	Module interface{} `json:"Module,omitempty" xml:"Module,omitempty"`
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
	// System exception
	RootErrorMsg *string `json:"RootErrorMsg,omitempty" xml:"RootErrorMsg,omitempty"`
	// The reserved parameter.
	//
	// example:
	//
	// True
	Synchro *bool `json:"Synchro,omitempty" xml:"Synchro,omitempty"`
}

func (s GetAppSandboxPreviewUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAppSandboxPreviewUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetAppSandboxPreviewUrlResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetAppSandboxPreviewUrlResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *GetAppSandboxPreviewUrlResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *GetAppSandboxPreviewUrlResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *GetAppSandboxPreviewUrlResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *GetAppSandboxPreviewUrlResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *GetAppSandboxPreviewUrlResponseBody) GetModule() interface{} {
	return s.Module
}

func (s *GetAppSandboxPreviewUrlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAppSandboxPreviewUrlResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *GetAppSandboxPreviewUrlResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *GetAppSandboxPreviewUrlResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *GetAppSandboxPreviewUrlResponseBody) SetAccessDeniedDetail(v string) *GetAppSandboxPreviewUrlResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetAppSandboxPreviewUrlResponseBody) SetAllowRetry(v bool) *GetAppSandboxPreviewUrlResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *GetAppSandboxPreviewUrlResponseBody) SetAppName(v string) *GetAppSandboxPreviewUrlResponseBody {
	s.AppName = &v
	return s
}

func (s *GetAppSandboxPreviewUrlResponseBody) SetDynamicCode(v string) *GetAppSandboxPreviewUrlResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *GetAppSandboxPreviewUrlResponseBody) SetDynamicMessage(v string) *GetAppSandboxPreviewUrlResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *GetAppSandboxPreviewUrlResponseBody) SetErrorArgs(v []interface{}) *GetAppSandboxPreviewUrlResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *GetAppSandboxPreviewUrlResponseBody) SetModule(v interface{}) *GetAppSandboxPreviewUrlResponseBody {
	s.Module = v
	return s
}

func (s *GetAppSandboxPreviewUrlResponseBody) SetRequestId(v string) *GetAppSandboxPreviewUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAppSandboxPreviewUrlResponseBody) SetRootErrorCode(v string) *GetAppSandboxPreviewUrlResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *GetAppSandboxPreviewUrlResponseBody) SetRootErrorMsg(v string) *GetAppSandboxPreviewUrlResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *GetAppSandboxPreviewUrlResponseBody) SetSynchro(v bool) *GetAppSandboxPreviewUrlResponseBody {
	s.Synchro = &v
	return s
}

func (s *GetAppSandboxPreviewUrlResponseBody) Validate() error {
	return dara.Validate(s)
}
