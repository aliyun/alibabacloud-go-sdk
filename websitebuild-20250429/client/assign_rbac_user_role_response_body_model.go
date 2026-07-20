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
	AccessDeniedDetail *string       `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	AllowRetry         *bool         `json:"AllowRetry,omitempty" xml:"AllowRetry,omitempty"`
	AppName            *string       `json:"AppName,omitempty" xml:"AppName,omitempty"`
	DynamicCode        *string       `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage     *string       `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorArgs          []interface{} `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	Module             *string       `json:"Module,omitempty" xml:"Module,omitempty"`
	RequestId          *string       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RootErrorCode      *string       `json:"RootErrorCode,omitempty" xml:"RootErrorCode,omitempty"`
	RootErrorMsg       *string       `json:"RootErrorMsg,omitempty" xml:"RootErrorMsg,omitempty"`
	Synchro            *bool         `json:"Synchro,omitempty" xml:"Synchro,omitempty"`
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
