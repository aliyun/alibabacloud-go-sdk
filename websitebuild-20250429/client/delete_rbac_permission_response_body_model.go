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
