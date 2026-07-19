// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveRbacRoleHierarchyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *RemoveRbacRoleHierarchyResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *RemoveRbacRoleHierarchyResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *RemoveRbacRoleHierarchyResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *RemoveRbacRoleHierarchyResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *RemoveRbacRoleHierarchyResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *RemoveRbacRoleHierarchyResponseBody
	GetErrorArgs() []interface{}
	SetModule(v bool) *RemoveRbacRoleHierarchyResponseBody
	GetModule() *bool
	SetRequestId(v string) *RemoveRbacRoleHierarchyResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *RemoveRbacRoleHierarchyResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *RemoveRbacRoleHierarchyResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *RemoveRbacRoleHierarchyResponseBody
	GetSynchro() *bool
}

type RemoveRbacRoleHierarchyResponseBody struct {
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
	// SYSTEM_ERROR
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

func (s RemoveRbacRoleHierarchyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveRbacRoleHierarchyResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveRbacRoleHierarchyResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *RemoveRbacRoleHierarchyResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *RemoveRbacRoleHierarchyResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *RemoveRbacRoleHierarchyResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *RemoveRbacRoleHierarchyResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *RemoveRbacRoleHierarchyResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *RemoveRbacRoleHierarchyResponseBody) GetModule() *bool {
	return s.Module
}

func (s *RemoveRbacRoleHierarchyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveRbacRoleHierarchyResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *RemoveRbacRoleHierarchyResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *RemoveRbacRoleHierarchyResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *RemoveRbacRoleHierarchyResponseBody) SetAccessDeniedDetail(v string) *RemoveRbacRoleHierarchyResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *RemoveRbacRoleHierarchyResponseBody) SetAllowRetry(v bool) *RemoveRbacRoleHierarchyResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *RemoveRbacRoleHierarchyResponseBody) SetAppName(v string) *RemoveRbacRoleHierarchyResponseBody {
	s.AppName = &v
	return s
}

func (s *RemoveRbacRoleHierarchyResponseBody) SetDynamicCode(v string) *RemoveRbacRoleHierarchyResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *RemoveRbacRoleHierarchyResponseBody) SetDynamicMessage(v string) *RemoveRbacRoleHierarchyResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *RemoveRbacRoleHierarchyResponseBody) SetErrorArgs(v []interface{}) *RemoveRbacRoleHierarchyResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *RemoveRbacRoleHierarchyResponseBody) SetModule(v bool) *RemoveRbacRoleHierarchyResponseBody {
	s.Module = &v
	return s
}

func (s *RemoveRbacRoleHierarchyResponseBody) SetRequestId(v string) *RemoveRbacRoleHierarchyResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveRbacRoleHierarchyResponseBody) SetRootErrorCode(v string) *RemoveRbacRoleHierarchyResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *RemoveRbacRoleHierarchyResponseBody) SetRootErrorMsg(v string) *RemoveRbacRoleHierarchyResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *RemoveRbacRoleHierarchyResponseBody) SetSynchro(v bool) *RemoveRbacRoleHierarchyResponseBody {
	s.Synchro = &v
	return s
}

func (s *RemoveRbacRoleHierarchyResponseBody) Validate() error {
	return dara.Validate(s)
}
