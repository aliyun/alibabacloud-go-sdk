// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRetryChannelPublishResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *RetryChannelPublishResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *RetryChannelPublishResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *RetryChannelPublishResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *RetryChannelPublishResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *RetryChannelPublishResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *RetryChannelPublishResponseBody
	GetErrorArgs() []interface{}
	SetModule(v string) *RetryChannelPublishResponseBody
	GetModule() *string
	SetRequestId(v string) *RetryChannelPublishResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *RetryChannelPublishResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *RetryChannelPublishResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *RetryChannelPublishResponseBody
	GetSynchro() *bool
}

type RetryChannelPublishResponseBody struct {
	// The access denied details.
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
	// The application name. The name can contain digits, letters, and hyphens (-). It must start with a letter, cannot end with a hyphen (-), and cannot exceed 36 characters.
	//
	// example:
	//
	// ish-intelligence-store-platform-admin-web
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The dynamic error code.
	//
	// example:
	//
	// ERROR-oo1
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// The dynamic error message, which is used to replace the `%s` in the **ErrMessage*	- return parameter.
	//
	// > If **ErrMessage*	- returns **The Value of Input Parameter %s is not valid*	- and **DynamicMessage*	- returns **DtsJobId**, the request parameter **DtsJobId*	- is invalid.
	//
	// example:
	//
	// SYSTEM_ERROR
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The error parameters.
	ErrorArgs []interface{} `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	// The returned object.
	//
	// example:
	//
	// true
	Module *string `json:"Module,omitempty" xml:"Module,omitempty"`
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
	// A reserved parameter.
	//
	// example:
	//
	// True
	Synchro *bool `json:"Synchro,omitempty" xml:"Synchro,omitempty"`
}

func (s RetryChannelPublishResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RetryChannelPublishResponseBody) GoString() string {
	return s.String()
}

func (s *RetryChannelPublishResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *RetryChannelPublishResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *RetryChannelPublishResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *RetryChannelPublishResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *RetryChannelPublishResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *RetryChannelPublishResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *RetryChannelPublishResponseBody) GetModule() *string {
	return s.Module
}

func (s *RetryChannelPublishResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RetryChannelPublishResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *RetryChannelPublishResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *RetryChannelPublishResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *RetryChannelPublishResponseBody) SetAccessDeniedDetail(v string) *RetryChannelPublishResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *RetryChannelPublishResponseBody) SetAllowRetry(v bool) *RetryChannelPublishResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *RetryChannelPublishResponseBody) SetAppName(v string) *RetryChannelPublishResponseBody {
	s.AppName = &v
	return s
}

func (s *RetryChannelPublishResponseBody) SetDynamicCode(v string) *RetryChannelPublishResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *RetryChannelPublishResponseBody) SetDynamicMessage(v string) *RetryChannelPublishResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *RetryChannelPublishResponseBody) SetErrorArgs(v []interface{}) *RetryChannelPublishResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *RetryChannelPublishResponseBody) SetModule(v string) *RetryChannelPublishResponseBody {
	s.Module = &v
	return s
}

func (s *RetryChannelPublishResponseBody) SetRequestId(v string) *RetryChannelPublishResponseBody {
	s.RequestId = &v
	return s
}

func (s *RetryChannelPublishResponseBody) SetRootErrorCode(v string) *RetryChannelPublishResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *RetryChannelPublishResponseBody) SetRootErrorMsg(v string) *RetryChannelPublishResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *RetryChannelPublishResponseBody) SetSynchro(v bool) *RetryChannelPublishResponseBody {
	s.Synchro = &v
	return s
}

func (s *RetryChannelPublishResponseBody) Validate() error {
	return dara.Validate(s)
}
