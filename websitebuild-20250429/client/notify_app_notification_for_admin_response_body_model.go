// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iNotifyAppNotificationForAdminResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *NotifyAppNotificationForAdminResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *NotifyAppNotificationForAdminResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *NotifyAppNotificationForAdminResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *NotifyAppNotificationForAdminResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *NotifyAppNotificationForAdminResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *NotifyAppNotificationForAdminResponseBody
	GetErrorArgs() []interface{}
	SetModule(v bool) *NotifyAppNotificationForAdminResponseBody
	GetModule() *bool
	SetRequestId(v string) *NotifyAppNotificationForAdminResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *NotifyAppNotificationForAdminResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *NotifyAppNotificationForAdminResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *NotifyAppNotificationForAdminResponseBody
	GetSynchro() *bool
}

type NotifyAppNotificationForAdminResponseBody struct {
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
	// The application name. The name must start with a letter and can contain letters, digits, underscores (_), and hyphens (-). The name can be up to 36 characters in length.
	//
	// example:
	//
	// dewuApp
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The dynamic error code.
	//
	// example:
	//
	// ERROR-oo1
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// The dynamic error message, which is used to replace the %s variable in the ErrMessage response parameter.
	//
	// example:
	//
	// SYSTEM_ERROR
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The error parameters.
	ErrorArgs []interface{} `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	// The response data.
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

func (s NotifyAppNotificationForAdminResponseBody) String() string {
	return dara.Prettify(s)
}

func (s NotifyAppNotificationForAdminResponseBody) GoString() string {
	return s.String()
}

func (s *NotifyAppNotificationForAdminResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *NotifyAppNotificationForAdminResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *NotifyAppNotificationForAdminResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *NotifyAppNotificationForAdminResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *NotifyAppNotificationForAdminResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *NotifyAppNotificationForAdminResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *NotifyAppNotificationForAdminResponseBody) GetModule() *bool {
	return s.Module
}

func (s *NotifyAppNotificationForAdminResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *NotifyAppNotificationForAdminResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *NotifyAppNotificationForAdminResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *NotifyAppNotificationForAdminResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *NotifyAppNotificationForAdminResponseBody) SetAccessDeniedDetail(v string) *NotifyAppNotificationForAdminResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *NotifyAppNotificationForAdminResponseBody) SetAllowRetry(v bool) *NotifyAppNotificationForAdminResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *NotifyAppNotificationForAdminResponseBody) SetAppName(v string) *NotifyAppNotificationForAdminResponseBody {
	s.AppName = &v
	return s
}

func (s *NotifyAppNotificationForAdminResponseBody) SetDynamicCode(v string) *NotifyAppNotificationForAdminResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *NotifyAppNotificationForAdminResponseBody) SetDynamicMessage(v string) *NotifyAppNotificationForAdminResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *NotifyAppNotificationForAdminResponseBody) SetErrorArgs(v []interface{}) *NotifyAppNotificationForAdminResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *NotifyAppNotificationForAdminResponseBody) SetModule(v bool) *NotifyAppNotificationForAdminResponseBody {
	s.Module = &v
	return s
}

func (s *NotifyAppNotificationForAdminResponseBody) SetRequestId(v string) *NotifyAppNotificationForAdminResponseBody {
	s.RequestId = &v
	return s
}

func (s *NotifyAppNotificationForAdminResponseBody) SetRootErrorCode(v string) *NotifyAppNotificationForAdminResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *NotifyAppNotificationForAdminResponseBody) SetRootErrorMsg(v string) *NotifyAppNotificationForAdminResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *NotifyAppNotificationForAdminResponseBody) SetSynchro(v bool) *NotifyAppNotificationForAdminResponseBody {
	s.Synchro = &v
	return s
}

func (s *NotifyAppNotificationForAdminResponseBody) Validate() error {
	return dara.Validate(s)
}
