// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateAppTemplateLikeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *OperateAppTemplateLikeResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *OperateAppTemplateLikeResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *OperateAppTemplateLikeResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *OperateAppTemplateLikeResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *OperateAppTemplateLikeResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *OperateAppTemplateLikeResponseBody
	GetErrorArgs() []interface{}
	SetModule(v bool) *OperateAppTemplateLikeResponseBody
	GetModule() *bool
	SetRequestId(v string) *OperateAppTemplateLikeResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *OperateAppTemplateLikeResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *OperateAppTemplateLikeResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *OperateAppTemplateLikeResponseBody
	GetSynchro() *bool
}

type OperateAppTemplateLikeResponseBody struct {
	// Permission denied information
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
	// Application name. It can contain digits, letters, and hyphens (-). It must start with a letter, cannot end with a hyphen (-), and must be no more than 36 characters in length.
	//
	// example:
	//
	// ish-intelligence-store-platform-admin-web
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// Dynamic error code.
	//
	// example:
	//
	// ERROR-oo1
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// Dynamic message.
	//
	// example:
	//
	// SYSTEM_ERROR
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// Returned error parameters
	ErrorArgs []interface{} `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	// Response data
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
	// Error code
	//
	// example:
	//
	// SYSTEM.ERROR
	RootErrorCode *string `json:"RootErrorCode,omitempty" xml:"RootErrorCode,omitempty"`
	// Abnormal message
	//
	// example:
	//
	// 系统异常
	RootErrorMsg *string `json:"RootErrorMsg,omitempty" xml:"RootErrorMsg,omitempty"`
	// Reserved parameter.
	//
	// example:
	//
	// True
	Synchro *bool `json:"Synchro,omitempty" xml:"Synchro,omitempty"`
}

func (s OperateAppTemplateLikeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OperateAppTemplateLikeResponseBody) GoString() string {
	return s.String()
}

func (s *OperateAppTemplateLikeResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *OperateAppTemplateLikeResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *OperateAppTemplateLikeResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *OperateAppTemplateLikeResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *OperateAppTemplateLikeResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *OperateAppTemplateLikeResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *OperateAppTemplateLikeResponseBody) GetModule() *bool {
	return s.Module
}

func (s *OperateAppTemplateLikeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OperateAppTemplateLikeResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *OperateAppTemplateLikeResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *OperateAppTemplateLikeResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *OperateAppTemplateLikeResponseBody) SetAccessDeniedDetail(v string) *OperateAppTemplateLikeResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *OperateAppTemplateLikeResponseBody) SetAllowRetry(v bool) *OperateAppTemplateLikeResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *OperateAppTemplateLikeResponseBody) SetAppName(v string) *OperateAppTemplateLikeResponseBody {
	s.AppName = &v
	return s
}

func (s *OperateAppTemplateLikeResponseBody) SetDynamicCode(v string) *OperateAppTemplateLikeResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *OperateAppTemplateLikeResponseBody) SetDynamicMessage(v string) *OperateAppTemplateLikeResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *OperateAppTemplateLikeResponseBody) SetErrorArgs(v []interface{}) *OperateAppTemplateLikeResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *OperateAppTemplateLikeResponseBody) SetModule(v bool) *OperateAppTemplateLikeResponseBody {
	s.Module = &v
	return s
}

func (s *OperateAppTemplateLikeResponseBody) SetRequestId(v string) *OperateAppTemplateLikeResponseBody {
	s.RequestId = &v
	return s
}

func (s *OperateAppTemplateLikeResponseBody) SetRootErrorCode(v string) *OperateAppTemplateLikeResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *OperateAppTemplateLikeResponseBody) SetRootErrorMsg(v string) *OperateAppTemplateLikeResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *OperateAppTemplateLikeResponseBody) SetSynchro(v bool) *OperateAppTemplateLikeResponseBody {
	s.Synchro = &v
	return s
}

func (s *OperateAppTemplateLikeResponseBody) Validate() error {
	return dara.Validate(s)
}
