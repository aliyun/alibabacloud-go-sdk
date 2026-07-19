// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRbacRoleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *UpdateRbacRoleResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *UpdateRbacRoleResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *UpdateRbacRoleResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *UpdateRbacRoleResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *UpdateRbacRoleResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *UpdateRbacRoleResponseBody
	GetErrorArgs() []interface{}
	SetModule(v string) *UpdateRbacRoleResponseBody
	GetModule() *string
	SetRequestId(v string) *UpdateRbacRoleResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *UpdateRbacRoleResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *UpdateRbacRoleResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *UpdateRbacRoleResponseBody
	GetSynchro() *bool
}

type UpdateRbacRoleResponseBody struct {
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// False
	AllowRetry *bool `json:"AllowRetry,omitempty" xml:"AllowRetry,omitempty"`
	// example:
	//
	// spring-cloud-b
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// ERROR-oo1
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// example:
	//
	// abc
	DynamicMessage *string       `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorArgs      []interface{} `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Module *string `json:"Module,omitempty" xml:"Module,omitempty"`
	// example:
	//
	// 6C6B99AC-39EC-5350-874C-204128C905E6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// SYSTEM.ERROR
	RootErrorCode *string `json:"RootErrorCode,omitempty" xml:"RootErrorCode,omitempty"`
	// example:
	//
	// 系统异常
	RootErrorMsg *string `json:"RootErrorMsg,omitempty" xml:"RootErrorMsg,omitempty"`
	// example:
	//
	// True
	Synchro *bool `json:"Synchro,omitempty" xml:"Synchro,omitempty"`
}

func (s UpdateRbacRoleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateRbacRoleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRbacRoleResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *UpdateRbacRoleResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *UpdateRbacRoleResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *UpdateRbacRoleResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *UpdateRbacRoleResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *UpdateRbacRoleResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *UpdateRbacRoleResponseBody) GetModule() *string {
	return s.Module
}

func (s *UpdateRbacRoleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateRbacRoleResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *UpdateRbacRoleResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *UpdateRbacRoleResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *UpdateRbacRoleResponseBody) SetAccessDeniedDetail(v string) *UpdateRbacRoleResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *UpdateRbacRoleResponseBody) SetAllowRetry(v bool) *UpdateRbacRoleResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *UpdateRbacRoleResponseBody) SetAppName(v string) *UpdateRbacRoleResponseBody {
	s.AppName = &v
	return s
}

func (s *UpdateRbacRoleResponseBody) SetDynamicCode(v string) *UpdateRbacRoleResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *UpdateRbacRoleResponseBody) SetDynamicMessage(v string) *UpdateRbacRoleResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *UpdateRbacRoleResponseBody) SetErrorArgs(v []interface{}) *UpdateRbacRoleResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *UpdateRbacRoleResponseBody) SetModule(v string) *UpdateRbacRoleResponseBody {
	s.Module = &v
	return s
}

func (s *UpdateRbacRoleResponseBody) SetRequestId(v string) *UpdateRbacRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateRbacRoleResponseBody) SetRootErrorCode(v string) *UpdateRbacRoleResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *UpdateRbacRoleResponseBody) SetRootErrorMsg(v string) *UpdateRbacRoleResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *UpdateRbacRoleResponseBody) SetSynchro(v bool) *UpdateRbacRoleResponseBody {
	s.Synchro = &v
	return s
}

func (s *UpdateRbacRoleResponseBody) Validate() error {
	return dara.Validate(s)
}
