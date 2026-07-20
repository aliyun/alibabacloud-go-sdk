// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRbacOrgUnitResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *UpdateRbacOrgUnitResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *UpdateRbacOrgUnitResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *UpdateRbacOrgUnitResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *UpdateRbacOrgUnitResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *UpdateRbacOrgUnitResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *UpdateRbacOrgUnitResponseBody
	GetErrorArgs() []interface{}
	SetModule(v string) *UpdateRbacOrgUnitResponseBody
	GetModule() *string
	SetRequestId(v string) *UpdateRbacOrgUnitResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *UpdateRbacOrgUnitResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *UpdateRbacOrgUnitResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *UpdateRbacOrgUnitResponseBody
	GetSynchro() *bool
}

type UpdateRbacOrgUnitResponseBody struct {
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

func (s UpdateRbacOrgUnitResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateRbacOrgUnitResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRbacOrgUnitResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *UpdateRbacOrgUnitResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *UpdateRbacOrgUnitResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *UpdateRbacOrgUnitResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *UpdateRbacOrgUnitResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *UpdateRbacOrgUnitResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *UpdateRbacOrgUnitResponseBody) GetModule() *string {
	return s.Module
}

func (s *UpdateRbacOrgUnitResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateRbacOrgUnitResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *UpdateRbacOrgUnitResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *UpdateRbacOrgUnitResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *UpdateRbacOrgUnitResponseBody) SetAccessDeniedDetail(v string) *UpdateRbacOrgUnitResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *UpdateRbacOrgUnitResponseBody) SetAllowRetry(v bool) *UpdateRbacOrgUnitResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *UpdateRbacOrgUnitResponseBody) SetAppName(v string) *UpdateRbacOrgUnitResponseBody {
	s.AppName = &v
	return s
}

func (s *UpdateRbacOrgUnitResponseBody) SetDynamicCode(v string) *UpdateRbacOrgUnitResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *UpdateRbacOrgUnitResponseBody) SetDynamicMessage(v string) *UpdateRbacOrgUnitResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *UpdateRbacOrgUnitResponseBody) SetErrorArgs(v []interface{}) *UpdateRbacOrgUnitResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *UpdateRbacOrgUnitResponseBody) SetModule(v string) *UpdateRbacOrgUnitResponseBody {
	s.Module = &v
	return s
}

func (s *UpdateRbacOrgUnitResponseBody) SetRequestId(v string) *UpdateRbacOrgUnitResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateRbacOrgUnitResponseBody) SetRootErrorCode(v string) *UpdateRbacOrgUnitResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *UpdateRbacOrgUnitResponseBody) SetRootErrorMsg(v string) *UpdateRbacOrgUnitResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *UpdateRbacOrgUnitResponseBody) SetSynchro(v bool) *UpdateRbacOrgUnitResponseBody {
	s.Synchro = &v
	return s
}

func (s *UpdateRbacOrgUnitResponseBody) Validate() error {
	return dara.Validate(s)
}
