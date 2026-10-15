// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRbacPermissionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CreateRbacPermissionResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *CreateRbacPermissionResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *CreateRbacPermissionResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *CreateRbacPermissionResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *CreateRbacPermissionResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *CreateRbacPermissionResponseBody
	GetErrorArgs() []interface{}
	SetModule(v string) *CreateRbacPermissionResponseBody
	GetModule() *string
	SetRequestId(v string) *CreateRbacPermissionResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *CreateRbacPermissionResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *CreateRbacPermissionResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *CreateRbacPermissionResponseBody
	GetSynchro() *bool
}

type CreateRbacPermissionResponseBody struct {
	// The detailed reason why access is denied.
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
	// spring-cloud-b
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The dynamic error code.
	//
	// example:
	//
	// ERROR-oo1
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// The dynamic message.
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
	// The error message.
	//
	// example:
	//
	// 系统异常
	RootErrorMsg *string `json:"RootErrorMsg,omitempty" xml:"RootErrorMsg,omitempty"`
	// The reserved parameter.
	//
	// example:
	//
	// True
	Synchro *bool `json:"Synchro,omitempty" xml:"Synchro,omitempty"`
}

func (s CreateRbacPermissionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateRbacPermissionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRbacPermissionResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CreateRbacPermissionResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *CreateRbacPermissionResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *CreateRbacPermissionResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *CreateRbacPermissionResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *CreateRbacPermissionResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *CreateRbacPermissionResponseBody) GetModule() *string {
	return s.Module
}

func (s *CreateRbacPermissionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateRbacPermissionResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *CreateRbacPermissionResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *CreateRbacPermissionResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *CreateRbacPermissionResponseBody) SetAccessDeniedDetail(v string) *CreateRbacPermissionResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CreateRbacPermissionResponseBody) SetAllowRetry(v bool) *CreateRbacPermissionResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *CreateRbacPermissionResponseBody) SetAppName(v string) *CreateRbacPermissionResponseBody {
	s.AppName = &v
	return s
}

func (s *CreateRbacPermissionResponseBody) SetDynamicCode(v string) *CreateRbacPermissionResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *CreateRbacPermissionResponseBody) SetDynamicMessage(v string) *CreateRbacPermissionResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *CreateRbacPermissionResponseBody) SetErrorArgs(v []interface{}) *CreateRbacPermissionResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *CreateRbacPermissionResponseBody) SetModule(v string) *CreateRbacPermissionResponseBody {
	s.Module = &v
	return s
}

func (s *CreateRbacPermissionResponseBody) SetRequestId(v string) *CreateRbacPermissionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRbacPermissionResponseBody) SetRootErrorCode(v string) *CreateRbacPermissionResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *CreateRbacPermissionResponseBody) SetRootErrorMsg(v string) *CreateRbacPermissionResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *CreateRbacPermissionResponseBody) SetSynchro(v bool) *CreateRbacPermissionResponseBody {
	s.Synchro = &v
	return s
}

func (s *CreateRbacPermissionResponseBody) Validate() error {
	return dara.Validate(s)
}
