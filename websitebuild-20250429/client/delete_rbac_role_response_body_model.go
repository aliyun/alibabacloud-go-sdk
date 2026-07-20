// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRbacRoleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DeleteRbacRoleResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *DeleteRbacRoleResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *DeleteRbacRoleResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *DeleteRbacRoleResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *DeleteRbacRoleResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *DeleteRbacRoleResponseBody
	GetErrorArgs() []interface{}
	SetModule(v bool) *DeleteRbacRoleResponseBody
	GetModule() *bool
	SetRequestId(v string) *DeleteRbacRoleResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *DeleteRbacRoleResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *DeleteRbacRoleResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *DeleteRbacRoleResponseBody
	GetSynchro() *bool
}

type DeleteRbacRoleResponseBody struct {
	// The details about the access denial.
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
	// The application name. The name can contain digits, letters, and hyphens (-). It must start with a letter and cannot end with a hyphen (-). The name cannot exceed 36 characters in length.
	//
	// example:
	//
	// spring-cloud-b
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The dynamic code. This parameter is not in use. Ignore this parameter.
	//
	// example:
	//
	// ERROR-oo1
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// The dynamic error message used to replace the `%s` variable in the **ErrMessage*	- parameter.
	//
	// > If **ErrMessage*	- returns **The Value of Input Parameter %s is not valid*	- and **DynamicMessage*	- returns **DtsJobId**, the **DtsJobId*	- request parameter is invalid.
	//
	// example:
	//
	// SYSTEM_ERROR
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The error parameters.
	ErrorArgs []interface{} `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	// The response data.
	//
	// example:
	//
	// true
	Module *bool `json:"Module,omitempty" xml:"Module,omitempty"`
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
	// A reserved parameter.
	//
	// example:
	//
	// True
	Synchro *bool `json:"Synchro,omitempty" xml:"Synchro,omitempty"`
}

func (s DeleteRbacRoleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteRbacRoleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRbacRoleResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DeleteRbacRoleResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *DeleteRbacRoleResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *DeleteRbacRoleResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *DeleteRbacRoleResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *DeleteRbacRoleResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *DeleteRbacRoleResponseBody) GetModule() *bool {
	return s.Module
}

func (s *DeleteRbacRoleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteRbacRoleResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *DeleteRbacRoleResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *DeleteRbacRoleResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *DeleteRbacRoleResponseBody) SetAccessDeniedDetail(v string) *DeleteRbacRoleResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DeleteRbacRoleResponseBody) SetAllowRetry(v bool) *DeleteRbacRoleResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *DeleteRbacRoleResponseBody) SetAppName(v string) *DeleteRbacRoleResponseBody {
	s.AppName = &v
	return s
}

func (s *DeleteRbacRoleResponseBody) SetDynamicCode(v string) *DeleteRbacRoleResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DeleteRbacRoleResponseBody) SetDynamicMessage(v string) *DeleteRbacRoleResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DeleteRbacRoleResponseBody) SetErrorArgs(v []interface{}) *DeleteRbacRoleResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *DeleteRbacRoleResponseBody) SetModule(v bool) *DeleteRbacRoleResponseBody {
	s.Module = &v
	return s
}

func (s *DeleteRbacRoleResponseBody) SetRequestId(v string) *DeleteRbacRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRbacRoleResponseBody) SetRootErrorCode(v string) *DeleteRbacRoleResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *DeleteRbacRoleResponseBody) SetRootErrorMsg(v string) *DeleteRbacRoleResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *DeleteRbacRoleResponseBody) SetSynchro(v bool) *DeleteRbacRoleResponseBody {
	s.Synchro = &v
	return s
}

func (s *DeleteRbacRoleResponseBody) Validate() error {
	return dara.Validate(s)
}
