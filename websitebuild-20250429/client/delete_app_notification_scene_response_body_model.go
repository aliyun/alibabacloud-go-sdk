// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAppNotificationSceneResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DeleteAppNotificationSceneResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *DeleteAppNotificationSceneResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *DeleteAppNotificationSceneResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *DeleteAppNotificationSceneResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *DeleteAppNotificationSceneResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *DeleteAppNotificationSceneResponseBody
	GetErrorArgs() []interface{}
	SetModule(v bool) *DeleteAppNotificationSceneResponseBody
	GetModule() *bool
	SetRequestId(v string) *DeleteAppNotificationSceneResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *DeleteAppNotificationSceneResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *DeleteAppNotificationSceneResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *DeleteAppNotificationSceneResponseBody
	GetSynchro() *bool
}

type DeleteAppNotificationSceneResponseBody struct {
	// The details of the permission verification failure.
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
	// ish-intelligence-store-platform-admin-web
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The dynamic error code.
	//
	// example:
	//
	// ERROR-oo1
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// The dynamic error message, which is used to replace the `%s` variable in the **ErrMessage*	- response parameter.
	//
	// > If **ErrMessage*	- returns **The Value of Input Parameter %s is not valid*	- and **DynamicMessage*	- returns **DtsJobId**, the request parameter **DtsJobId*	- is invalid.
	//
	// example:
	//
	// SYSTEM_ERROR
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The returned error parameters.
	ErrorArgs []interface{} `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	// Indicates whether the move operation is successful.
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
	// System exception.
	RootErrorMsg *string `json:"RootErrorMsg,omitempty" xml:"RootErrorMsg,omitempty"`
	// Indicates whether the request is synchronously processed.
	//
	// example:
	//
	// True
	Synchro *bool `json:"Synchro,omitempty" xml:"Synchro,omitempty"`
}

func (s DeleteAppNotificationSceneResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAppNotificationSceneResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAppNotificationSceneResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DeleteAppNotificationSceneResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *DeleteAppNotificationSceneResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *DeleteAppNotificationSceneResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *DeleteAppNotificationSceneResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *DeleteAppNotificationSceneResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *DeleteAppNotificationSceneResponseBody) GetModule() *bool {
	return s.Module
}

func (s *DeleteAppNotificationSceneResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAppNotificationSceneResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *DeleteAppNotificationSceneResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *DeleteAppNotificationSceneResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *DeleteAppNotificationSceneResponseBody) SetAccessDeniedDetail(v string) *DeleteAppNotificationSceneResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DeleteAppNotificationSceneResponseBody) SetAllowRetry(v bool) *DeleteAppNotificationSceneResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *DeleteAppNotificationSceneResponseBody) SetAppName(v string) *DeleteAppNotificationSceneResponseBody {
	s.AppName = &v
	return s
}

func (s *DeleteAppNotificationSceneResponseBody) SetDynamicCode(v string) *DeleteAppNotificationSceneResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DeleteAppNotificationSceneResponseBody) SetDynamicMessage(v string) *DeleteAppNotificationSceneResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DeleteAppNotificationSceneResponseBody) SetErrorArgs(v []interface{}) *DeleteAppNotificationSceneResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *DeleteAppNotificationSceneResponseBody) SetModule(v bool) *DeleteAppNotificationSceneResponseBody {
	s.Module = &v
	return s
}

func (s *DeleteAppNotificationSceneResponseBody) SetRequestId(v string) *DeleteAppNotificationSceneResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAppNotificationSceneResponseBody) SetRootErrorCode(v string) *DeleteAppNotificationSceneResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *DeleteAppNotificationSceneResponseBody) SetRootErrorMsg(v string) *DeleteAppNotificationSceneResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *DeleteAppNotificationSceneResponseBody) SetSynchro(v bool) *DeleteAppNotificationSceneResponseBody {
	s.Synchro = &v
	return s
}

func (s *DeleteAppNotificationSceneResponseBody) Validate() error {
	return dara.Validate(s)
}
