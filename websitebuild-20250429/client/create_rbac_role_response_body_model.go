// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRbacRoleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CreateRbacRoleResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *CreateRbacRoleResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *CreateRbacRoleResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *CreateRbacRoleResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *CreateRbacRoleResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *CreateRbacRoleResponseBody
	GetErrorArgs() []interface{}
	SetModule(v string) *CreateRbacRoleResponseBody
	GetModule() *string
	SetRequestId(v string) *CreateRbacRoleResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *CreateRbacRoleResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *CreateRbacRoleResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *CreateRbacRoleResponseBody
	GetSynchro() *bool
}

type CreateRbacRoleResponseBody struct {
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
	// ish-intelligence-store-platform-admin-web
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
	// Indicates whether the request is synchronously processed.
	//
	// example:
	//
	// True
	Synchro *bool `json:"Synchro,omitempty" xml:"Synchro,omitempty"`
}

func (s CreateRbacRoleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateRbacRoleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRbacRoleResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CreateRbacRoleResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *CreateRbacRoleResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *CreateRbacRoleResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *CreateRbacRoleResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *CreateRbacRoleResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *CreateRbacRoleResponseBody) GetModule() *string {
	return s.Module
}

func (s *CreateRbacRoleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateRbacRoleResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *CreateRbacRoleResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *CreateRbacRoleResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *CreateRbacRoleResponseBody) SetAccessDeniedDetail(v string) *CreateRbacRoleResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CreateRbacRoleResponseBody) SetAllowRetry(v bool) *CreateRbacRoleResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *CreateRbacRoleResponseBody) SetAppName(v string) *CreateRbacRoleResponseBody {
	s.AppName = &v
	return s
}

func (s *CreateRbacRoleResponseBody) SetDynamicCode(v string) *CreateRbacRoleResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *CreateRbacRoleResponseBody) SetDynamicMessage(v string) *CreateRbacRoleResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *CreateRbacRoleResponseBody) SetErrorArgs(v []interface{}) *CreateRbacRoleResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *CreateRbacRoleResponseBody) SetModule(v string) *CreateRbacRoleResponseBody {
	s.Module = &v
	return s
}

func (s *CreateRbacRoleResponseBody) SetRequestId(v string) *CreateRbacRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRbacRoleResponseBody) SetRootErrorCode(v string) *CreateRbacRoleResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *CreateRbacRoleResponseBody) SetRootErrorMsg(v string) *CreateRbacRoleResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *CreateRbacRoleResponseBody) SetSynchro(v bool) *CreateRbacRoleResponseBody {
	s.Synchro = &v
	return s
}

func (s *CreateRbacRoleResponseBody) Validate() error {
	return dara.Validate(s)
}
