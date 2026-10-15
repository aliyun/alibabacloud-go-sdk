// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRbacPermissionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DeleteRbacPermissionResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *DeleteRbacPermissionResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *DeleteRbacPermissionResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *DeleteRbacPermissionResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *DeleteRbacPermissionResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *DeleteRbacPermissionResponseBody
	GetErrorArgs() []interface{}
	SetModule(v bool) *DeleteRbacPermissionResponseBody
	GetModule() *bool
	SetRequestId(v string) *DeleteRbacPermissionResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *DeleteRbacPermissionResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *DeleteRbacPermissionResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *DeleteRbacPermissionResponseBody
	GetSynchro() *bool
}

type DeleteRbacPermissionResponseBody struct {
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
	// dewuApp
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The dynamic error code.
	//
	// example:
	//
	// ERROR-oo1
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// The dynamic error message, which is used to replace the `%s` placeholder in the **ErrMessage*	- response parameter.
	//
	// > If **ErrMessage*	- returns **The Value of Input Parameter %s is not valid*	- and **DynamicMessage*	- returns **DtsJobId**, the value of the request parameter **DtsJobId*	- is invalid.
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

func (s DeleteRbacPermissionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteRbacPermissionResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRbacPermissionResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DeleteRbacPermissionResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *DeleteRbacPermissionResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *DeleteRbacPermissionResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *DeleteRbacPermissionResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *DeleteRbacPermissionResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *DeleteRbacPermissionResponseBody) GetModule() *bool {
	return s.Module
}

func (s *DeleteRbacPermissionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteRbacPermissionResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *DeleteRbacPermissionResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *DeleteRbacPermissionResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *DeleteRbacPermissionResponseBody) SetAccessDeniedDetail(v string) *DeleteRbacPermissionResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DeleteRbacPermissionResponseBody) SetAllowRetry(v bool) *DeleteRbacPermissionResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *DeleteRbacPermissionResponseBody) SetAppName(v string) *DeleteRbacPermissionResponseBody {
	s.AppName = &v
	return s
}

func (s *DeleteRbacPermissionResponseBody) SetDynamicCode(v string) *DeleteRbacPermissionResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DeleteRbacPermissionResponseBody) SetDynamicMessage(v string) *DeleteRbacPermissionResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DeleteRbacPermissionResponseBody) SetErrorArgs(v []interface{}) *DeleteRbacPermissionResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *DeleteRbacPermissionResponseBody) SetModule(v bool) *DeleteRbacPermissionResponseBody {
	s.Module = &v
	return s
}

func (s *DeleteRbacPermissionResponseBody) SetRequestId(v string) *DeleteRbacPermissionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRbacPermissionResponseBody) SetRootErrorCode(v string) *DeleteRbacPermissionResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *DeleteRbacPermissionResponseBody) SetRootErrorMsg(v string) *DeleteRbacPermissionResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *DeleteRbacPermissionResponseBody) SetSynchro(v bool) *DeleteRbacPermissionResponseBody {
	s.Synchro = &v
	return s
}

func (s *DeleteRbacPermissionResponseBody) Validate() error {
	return dara.Validate(s)
}
