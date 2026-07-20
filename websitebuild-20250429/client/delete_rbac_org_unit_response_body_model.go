// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRbacOrgUnitResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DeleteRbacOrgUnitResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *DeleteRbacOrgUnitResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *DeleteRbacOrgUnitResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *DeleteRbacOrgUnitResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *DeleteRbacOrgUnitResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *DeleteRbacOrgUnitResponseBody
	GetErrorArgs() []interface{}
	SetModule(v bool) *DeleteRbacOrgUnitResponseBody
	GetModule() *bool
	SetRequestId(v string) *DeleteRbacOrgUnitResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *DeleteRbacOrgUnitResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *DeleteRbacOrgUnitResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *DeleteRbacOrgUnitResponseBody
	GetSynchro() *bool
}

type DeleteRbacOrgUnitResponseBody struct {
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

func (s DeleteRbacOrgUnitResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteRbacOrgUnitResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRbacOrgUnitResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DeleteRbacOrgUnitResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *DeleteRbacOrgUnitResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *DeleteRbacOrgUnitResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *DeleteRbacOrgUnitResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *DeleteRbacOrgUnitResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *DeleteRbacOrgUnitResponseBody) GetModule() *bool {
	return s.Module
}

func (s *DeleteRbacOrgUnitResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteRbacOrgUnitResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *DeleteRbacOrgUnitResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *DeleteRbacOrgUnitResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *DeleteRbacOrgUnitResponseBody) SetAccessDeniedDetail(v string) *DeleteRbacOrgUnitResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DeleteRbacOrgUnitResponseBody) SetAllowRetry(v bool) *DeleteRbacOrgUnitResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *DeleteRbacOrgUnitResponseBody) SetAppName(v string) *DeleteRbacOrgUnitResponseBody {
	s.AppName = &v
	return s
}

func (s *DeleteRbacOrgUnitResponseBody) SetDynamicCode(v string) *DeleteRbacOrgUnitResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DeleteRbacOrgUnitResponseBody) SetDynamicMessage(v string) *DeleteRbacOrgUnitResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DeleteRbacOrgUnitResponseBody) SetErrorArgs(v []interface{}) *DeleteRbacOrgUnitResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *DeleteRbacOrgUnitResponseBody) SetModule(v bool) *DeleteRbacOrgUnitResponseBody {
	s.Module = &v
	return s
}

func (s *DeleteRbacOrgUnitResponseBody) SetRequestId(v string) *DeleteRbacOrgUnitResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRbacOrgUnitResponseBody) SetRootErrorCode(v string) *DeleteRbacOrgUnitResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *DeleteRbacOrgUnitResponseBody) SetRootErrorMsg(v string) *DeleteRbacOrgUnitResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *DeleteRbacOrgUnitResponseBody) SetSynchro(v bool) *DeleteRbacOrgUnitResponseBody {
	s.Synchro = &v
	return s
}

func (s *DeleteRbacOrgUnitResponseBody) Validate() error {
	return dara.Validate(s)
}
