// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssignRbacUserRoleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *AssignRbacUserRoleResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *AssignRbacUserRoleResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *AssignRbacUserRoleResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *AssignRbacUserRoleResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *AssignRbacUserRoleResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *AssignRbacUserRoleResponseBody
	GetErrorArgs() []interface{}
	SetModule(v string) *AssignRbacUserRoleResponseBody
	GetModule() *string
	SetRequestId(v string) *AssignRbacUserRoleResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *AssignRbacUserRoleResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *AssignRbacUserRoleResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *AssignRbacUserRoleResponseBody
	GetSynchro() *bool
}

type AssignRbacUserRoleResponseBody struct {
	// The detailed reason why access was denied.
	//
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// Indicates whether retry is allowed. Valid values:
	//
	// - false: Retry is not allowed.
	//
	// - true: Retry is allowed.
	//
	// example:
	//
	// False
	AllowRetry *bool `json:"AllowRetry,omitempty" xml:"AllowRetry,omitempty"`
	// The application name. The application with this name is queried.
	//
	// example:
	//
	// or
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The dynamic error code.
	//
	// example:
	//
	// ERROR-oo1
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// The dynamic error message, used to replace the %s variable in the ErrMessage response parameter.
	//
	// example:
	//
	// SYSTEM.ERROR
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The error parameters.
	ErrorArgs []interface{} `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	// Indicates whether the deletion was successful.
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
	// SYSTEM.EROR
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

func (s AssignRbacUserRoleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AssignRbacUserRoleResponseBody) GoString() string {
	return s.String()
}

func (s *AssignRbacUserRoleResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *AssignRbacUserRoleResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *AssignRbacUserRoleResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *AssignRbacUserRoleResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *AssignRbacUserRoleResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *AssignRbacUserRoleResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *AssignRbacUserRoleResponseBody) GetModule() *string {
	return s.Module
}

func (s *AssignRbacUserRoleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AssignRbacUserRoleResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *AssignRbacUserRoleResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *AssignRbacUserRoleResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *AssignRbacUserRoleResponseBody) SetAccessDeniedDetail(v string) *AssignRbacUserRoleResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *AssignRbacUserRoleResponseBody) SetAllowRetry(v bool) *AssignRbacUserRoleResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *AssignRbacUserRoleResponseBody) SetAppName(v string) *AssignRbacUserRoleResponseBody {
	s.AppName = &v
	return s
}

func (s *AssignRbacUserRoleResponseBody) SetDynamicCode(v string) *AssignRbacUserRoleResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *AssignRbacUserRoleResponseBody) SetDynamicMessage(v string) *AssignRbacUserRoleResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *AssignRbacUserRoleResponseBody) SetErrorArgs(v []interface{}) *AssignRbacUserRoleResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *AssignRbacUserRoleResponseBody) SetModule(v string) *AssignRbacUserRoleResponseBody {
	s.Module = &v
	return s
}

func (s *AssignRbacUserRoleResponseBody) SetRequestId(v string) *AssignRbacUserRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *AssignRbacUserRoleResponseBody) SetRootErrorCode(v string) *AssignRbacUserRoleResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *AssignRbacUserRoleResponseBody) SetRootErrorMsg(v string) *AssignRbacUserRoleResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *AssignRbacUserRoleResponseBody) SetSynchro(v bool) *AssignRbacUserRoleResponseBody {
	s.Synchro = &v
	return s
}

func (s *AssignRbacUserRoleResponseBody) Validate() error {
	return dara.Validate(s)
}
