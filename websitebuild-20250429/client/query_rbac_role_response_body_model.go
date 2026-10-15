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
	// spring-cloud-b
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
	// The error parameters returned.
	ErrorArgs []interface{} `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	// The response data.
	Module *QueryRbacRoleResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
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
	// The exception message.
	//
	// example:
	//
	// 系统异常
	RootErrorMsg *string `json:"RootErrorMsg,omitempty" xml:"RootErrorMsg,omitempty"`
	// The reserved parameter.
	//
	// example:
	//
	// True
	Synchro *bool `json:"Synchro,omitempty" xml:"Synchro,omitempty"`
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
	// The time when the task was created, in ISO 8601 format (such as 2024-01-01T00:00:00+Z).
	//
	// example:
	//
	// 2026-02-06T10:20:09Z
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	// The scenario ID.
	//
	// example:
	//
	// 16257
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The default status. Valid values:
	//
	// - 0: not default.
	//
	// - 1: default.
	//
	// example:
	//
	// true
	IsDefault *bool `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// Indicates whether the approval template is system-defined. Valid values:
	//
	// - 1: Yes.
	//
	// - 0: No.
	IsSystem *bool `json:"IsSystem,omitempty" xml:"IsSystem,omitempty"`
	// The name.
	//
	// example:
	//
	// oss://anchashiscan.oss-cn-hangzhou.aliyuncs.com/prodpopscansDZBFBGr
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The file name.
	//
	// example:
	//
	// 文件名
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
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
