// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRbacRoleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *QueryRbacRoleResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *QueryRbacRoleResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *QueryRbacRoleResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *QueryRbacRoleResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *QueryRbacRoleResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *QueryRbacRoleResponseBody
	GetErrorArgs() []interface{}
	SetModule(v *QueryRbacRoleResponseBodyModule) *QueryRbacRoleResponseBody
	GetModule() *QueryRbacRoleResponseBodyModule
	SetRequestId(v string) *QueryRbacRoleResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *QueryRbacRoleResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *QueryRbacRoleResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *QueryRbacRoleResponseBody
	GetSynchro() *bool
}

type QueryRbacRoleResponseBody struct {
	AccessDeniedDetail *string                          `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	AllowRetry         *bool                            `json:"AllowRetry,omitempty" xml:"AllowRetry,omitempty"`
	AppName            *string                          `json:"AppName,omitempty" xml:"AppName,omitempty"`
	DynamicCode        *string                          `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage     *string                          `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorArgs          []interface{}                    `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	Module             *QueryRbacRoleResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
	RequestId          *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RootErrorCode      *string                          `json:"RootErrorCode,omitempty" xml:"RootErrorCode,omitempty"`
	RootErrorMsg       *string                          `json:"RootErrorMsg,omitempty" xml:"RootErrorMsg,omitempty"`
	Synchro            *bool                            `json:"Synchro,omitempty" xml:"Synchro,omitempty"`
}

func (s QueryRbacRoleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryRbacRoleResponseBody) GoString() string {
	return s.String()
}

func (s *QueryRbacRoleResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *QueryRbacRoleResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *QueryRbacRoleResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *QueryRbacRoleResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *QueryRbacRoleResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *QueryRbacRoleResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *QueryRbacRoleResponseBody) GetModule() *QueryRbacRoleResponseBodyModule {
	return s.Module
}

func (s *QueryRbacRoleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryRbacRoleResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *QueryRbacRoleResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *QueryRbacRoleResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *QueryRbacRoleResponseBody) SetAccessDeniedDetail(v string) *QueryRbacRoleResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *QueryRbacRoleResponseBody) SetAllowRetry(v bool) *QueryRbacRoleResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *QueryRbacRoleResponseBody) SetAppName(v string) *QueryRbacRoleResponseBody {
	s.AppName = &v
	return s
}

func (s *QueryRbacRoleResponseBody) SetDynamicCode(v string) *QueryRbacRoleResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *QueryRbacRoleResponseBody) SetDynamicMessage(v string) *QueryRbacRoleResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *QueryRbacRoleResponseBody) SetErrorArgs(v []interface{}) *QueryRbacRoleResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *QueryRbacRoleResponseBody) SetModule(v *QueryRbacRoleResponseBodyModule) *QueryRbacRoleResponseBody {
	s.Module = v
	return s
}

func (s *QueryRbacRoleResponseBody) SetRequestId(v string) *QueryRbacRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryRbacRoleResponseBody) SetRootErrorCode(v string) *QueryRbacRoleResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *QueryRbacRoleResponseBody) SetRootErrorMsg(v string) *QueryRbacRoleResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *QueryRbacRoleResponseBody) SetSynchro(v bool) *QueryRbacRoleResponseBody {
	s.Synchro = &v
	return s
}

func (s *QueryRbacRoleResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryRbacRoleResponseBodyModule struct {
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IsDefault *bool   `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	IsSystem  *bool   `json:"IsSystem,omitempty" xml:"IsSystem,omitempty"`
	Label     *string `json:"Label,omitempty" xml:"Label,omitempty"`
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s QueryRbacRoleResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s QueryRbacRoleResponseBodyModule) GoString() string {
	return s.String()
}

func (s *QueryRbacRoleResponseBodyModule) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *QueryRbacRoleResponseBodyModule) GetId() *string {
	return s.Id
}

func (s *QueryRbacRoleResponseBodyModule) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *QueryRbacRoleResponseBodyModule) GetIsSystem() *bool {
	return s.IsSystem
}

func (s *QueryRbacRoleResponseBodyModule) GetLabel() *string {
	return s.Label
}

func (s *QueryRbacRoleResponseBodyModule) GetName() *string {
	return s.Name
}

func (s *QueryRbacRoleResponseBodyModule) SetCreatedAt(v string) *QueryRbacRoleResponseBodyModule {
	s.CreatedAt = &v
	return s
}

func (s *QueryRbacRoleResponseBodyModule) SetId(v string) *QueryRbacRoleResponseBodyModule {
	s.Id = &v
	return s
}

func (s *QueryRbacRoleResponseBodyModule) SetIsDefault(v bool) *QueryRbacRoleResponseBodyModule {
	s.IsDefault = &v
	return s
}

func (s *QueryRbacRoleResponseBodyModule) SetIsSystem(v bool) *QueryRbacRoleResponseBodyModule {
	s.IsSystem = &v
	return s
}

func (s *QueryRbacRoleResponseBodyModule) SetLabel(v string) *QueryRbacRoleResponseBodyModule {
	s.Label = &v
	return s
}

func (s *QueryRbacRoleResponseBodyModule) SetName(v string) *QueryRbacRoleResponseBodyModule {
	s.Name = &v
	return s
}

func (s *QueryRbacRoleResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
