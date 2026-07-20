// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeRbacUserRoleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *RevokeRbacUserRoleResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *RevokeRbacUserRoleResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *RevokeRbacUserRoleResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *RevokeRbacUserRoleResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *RevokeRbacUserRoleResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *RevokeRbacUserRoleResponseBody
	GetErrorArgs() []interface{}
	SetModule(v bool) *RevokeRbacUserRoleResponseBody
	GetModule() *bool
	SetRequestId(v string) *RevokeRbacUserRoleResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *RevokeRbacUserRoleResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *RevokeRbacUserRoleResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *RevokeRbacUserRoleResponseBody
	GetSynchro() *bool
}

type RevokeRbacUserRoleResponseBody struct {
	AccessDeniedDetail *string       `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	AllowRetry         *bool         `json:"AllowRetry,omitempty" xml:"AllowRetry,omitempty"`
	AppName            *string       `json:"AppName,omitempty" xml:"AppName,omitempty"`
	DynamicCode        *string       `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage     *string       `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorArgs          []interface{} `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	Module             *bool         `json:"Module,omitempty" xml:"Module,omitempty"`
	RequestId          *string       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RootErrorCode      *string       `json:"RootErrorCode,omitempty" xml:"RootErrorCode,omitempty"`
	RootErrorMsg       *string       `json:"RootErrorMsg,omitempty" xml:"RootErrorMsg,omitempty"`
	Synchro            *bool         `json:"Synchro,omitempty" xml:"Synchro,omitempty"`
}

func (s RevokeRbacUserRoleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RevokeRbacUserRoleResponseBody) GoString() string {
	return s.String()
}

func (s *RevokeRbacUserRoleResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *RevokeRbacUserRoleResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *RevokeRbacUserRoleResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *RevokeRbacUserRoleResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *RevokeRbacUserRoleResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *RevokeRbacUserRoleResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *RevokeRbacUserRoleResponseBody) GetModule() *bool {
	return s.Module
}

func (s *RevokeRbacUserRoleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RevokeRbacUserRoleResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *RevokeRbacUserRoleResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *RevokeRbacUserRoleResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *RevokeRbacUserRoleResponseBody) SetAccessDeniedDetail(v string) *RevokeRbacUserRoleResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *RevokeRbacUserRoleResponseBody) SetAllowRetry(v bool) *RevokeRbacUserRoleResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *RevokeRbacUserRoleResponseBody) SetAppName(v string) *RevokeRbacUserRoleResponseBody {
	s.AppName = &v
	return s
}

func (s *RevokeRbacUserRoleResponseBody) SetDynamicCode(v string) *RevokeRbacUserRoleResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *RevokeRbacUserRoleResponseBody) SetDynamicMessage(v string) *RevokeRbacUserRoleResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *RevokeRbacUserRoleResponseBody) SetErrorArgs(v []interface{}) *RevokeRbacUserRoleResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *RevokeRbacUserRoleResponseBody) SetModule(v bool) *RevokeRbacUserRoleResponseBody {
	s.Module = &v
	return s
}

func (s *RevokeRbacUserRoleResponseBody) SetRequestId(v string) *RevokeRbacUserRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *RevokeRbacUserRoleResponseBody) SetRootErrorCode(v string) *RevokeRbacUserRoleResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *RevokeRbacUserRoleResponseBody) SetRootErrorMsg(v string) *RevokeRbacUserRoleResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *RevokeRbacUserRoleResponseBody) SetSynchro(v bool) *RevokeRbacUserRoleResponseBody {
	s.Synchro = &v
	return s
}

func (s *RevokeRbacUserRoleResponseBody) Validate() error {
	return dara.Validate(s)
}
