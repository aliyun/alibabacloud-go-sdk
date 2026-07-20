// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRbacOrgUnitResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CreateRbacOrgUnitResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *CreateRbacOrgUnitResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *CreateRbacOrgUnitResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *CreateRbacOrgUnitResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *CreateRbacOrgUnitResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *CreateRbacOrgUnitResponseBody
	GetErrorArgs() []interface{}
	SetModule(v string) *CreateRbacOrgUnitResponseBody
	GetModule() *string
	SetRequestId(v string) *CreateRbacOrgUnitResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *CreateRbacOrgUnitResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *CreateRbacOrgUnitResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *CreateRbacOrgUnitResponseBody
	GetSynchro() *bool
}

type CreateRbacOrgUnitResponseBody struct {
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

func (s CreateRbacOrgUnitResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateRbacOrgUnitResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRbacOrgUnitResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CreateRbacOrgUnitResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *CreateRbacOrgUnitResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *CreateRbacOrgUnitResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *CreateRbacOrgUnitResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *CreateRbacOrgUnitResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *CreateRbacOrgUnitResponseBody) GetModule() *string {
	return s.Module
}

func (s *CreateRbacOrgUnitResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateRbacOrgUnitResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *CreateRbacOrgUnitResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *CreateRbacOrgUnitResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *CreateRbacOrgUnitResponseBody) SetAccessDeniedDetail(v string) *CreateRbacOrgUnitResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CreateRbacOrgUnitResponseBody) SetAllowRetry(v bool) *CreateRbacOrgUnitResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *CreateRbacOrgUnitResponseBody) SetAppName(v string) *CreateRbacOrgUnitResponseBody {
	s.AppName = &v
	return s
}

func (s *CreateRbacOrgUnitResponseBody) SetDynamicCode(v string) *CreateRbacOrgUnitResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *CreateRbacOrgUnitResponseBody) SetDynamicMessage(v string) *CreateRbacOrgUnitResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *CreateRbacOrgUnitResponseBody) SetErrorArgs(v []interface{}) *CreateRbacOrgUnitResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *CreateRbacOrgUnitResponseBody) SetModule(v string) *CreateRbacOrgUnitResponseBody {
	s.Module = &v
	return s
}

func (s *CreateRbacOrgUnitResponseBody) SetRequestId(v string) *CreateRbacOrgUnitResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRbacOrgUnitResponseBody) SetRootErrorCode(v string) *CreateRbacOrgUnitResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *CreateRbacOrgUnitResponseBody) SetRootErrorMsg(v string) *CreateRbacOrgUnitResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *CreateRbacOrgUnitResponseBody) SetSynchro(v bool) *CreateRbacOrgUnitResponseBody {
	s.Synchro = &v
	return s
}

func (s *CreateRbacOrgUnitResponseBody) Validate() error {
	return dara.Validate(s)
}
