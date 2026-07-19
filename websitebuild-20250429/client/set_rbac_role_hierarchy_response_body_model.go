// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetRbacRoleHierarchyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *SetRbacRoleHierarchyResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *SetRbacRoleHierarchyResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *SetRbacRoleHierarchyResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *SetRbacRoleHierarchyResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *SetRbacRoleHierarchyResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *SetRbacRoleHierarchyResponseBody
	GetErrorArgs() []interface{}
	SetModule(v bool) *SetRbacRoleHierarchyResponseBody
	GetModule() *bool
	SetRequestId(v string) *SetRbacRoleHierarchyResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *SetRbacRoleHierarchyResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *SetRbacRoleHierarchyResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *SetRbacRoleHierarchyResponseBody
	GetSynchro() *bool
}

type SetRbacRoleHierarchyResponseBody struct {
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
	// or
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// ERROR-oo1
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// example:
	//
	// SYSTEM.ERROR
	DynamicMessage *string       `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorArgs      []interface{} `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Module *bool `json:"Module,omitempty" xml:"Module,omitempty"`
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

func (s SetRbacRoleHierarchyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetRbacRoleHierarchyResponseBody) GoString() string {
	return s.String()
}

func (s *SetRbacRoleHierarchyResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *SetRbacRoleHierarchyResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *SetRbacRoleHierarchyResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *SetRbacRoleHierarchyResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *SetRbacRoleHierarchyResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *SetRbacRoleHierarchyResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *SetRbacRoleHierarchyResponseBody) GetModule() *bool {
	return s.Module
}

func (s *SetRbacRoleHierarchyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetRbacRoleHierarchyResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *SetRbacRoleHierarchyResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *SetRbacRoleHierarchyResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *SetRbacRoleHierarchyResponseBody) SetAccessDeniedDetail(v string) *SetRbacRoleHierarchyResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *SetRbacRoleHierarchyResponseBody) SetAllowRetry(v bool) *SetRbacRoleHierarchyResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *SetRbacRoleHierarchyResponseBody) SetAppName(v string) *SetRbacRoleHierarchyResponseBody {
	s.AppName = &v
	return s
}

func (s *SetRbacRoleHierarchyResponseBody) SetDynamicCode(v string) *SetRbacRoleHierarchyResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *SetRbacRoleHierarchyResponseBody) SetDynamicMessage(v string) *SetRbacRoleHierarchyResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *SetRbacRoleHierarchyResponseBody) SetErrorArgs(v []interface{}) *SetRbacRoleHierarchyResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *SetRbacRoleHierarchyResponseBody) SetModule(v bool) *SetRbacRoleHierarchyResponseBody {
	s.Module = &v
	return s
}

func (s *SetRbacRoleHierarchyResponseBody) SetRequestId(v string) *SetRbacRoleHierarchyResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetRbacRoleHierarchyResponseBody) SetRootErrorCode(v string) *SetRbacRoleHierarchyResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *SetRbacRoleHierarchyResponseBody) SetRootErrorMsg(v string) *SetRbacRoleHierarchyResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *SetRbacRoleHierarchyResponseBody) SetSynchro(v bool) *SetRbacRoleHierarchyResponseBody {
	s.Synchro = &v
	return s
}

func (s *SetRbacRoleHierarchyResponseBody) Validate() error {
	return dara.Validate(s)
}
