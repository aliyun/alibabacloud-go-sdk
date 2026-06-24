// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnlineAppInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *OnlineAppInstanceResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *OnlineAppInstanceResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *OnlineAppInstanceResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *OnlineAppInstanceResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *OnlineAppInstanceResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *OnlineAppInstanceResponseBody
	GetErrorArgs() []interface{}
	SetModule(v map[string]interface{}) *OnlineAppInstanceResponseBody
	GetModule() map[string]interface{}
	SetRequestId(v string) *OnlineAppInstanceResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *OnlineAppInstanceResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *OnlineAppInstanceResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *OnlineAppInstanceResponseBody
	GetSynchro() *bool
}

type OnlineAppInstanceResponseBody struct {
	// The authentication failure details.
	//
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// Indicates whether retries are allowed.
	//
	// example:
	//
	// False
	AllowRetry *bool `json:"AllowRetry,omitempty" xml:"AllowRetry,omitempty"`
	// The application name. The name can contain digits, letters, and hyphens (-). It must start with a letter and cannot end with a hyphen (-). The name can be up to 36 characters in length.
	//
	// example:
	//
	// stg.mlp_finance_cap_refund_hi
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The dynamic error code. This parameter will be deprecated.
	//
	// example:
	//
	// ERROR-oo1
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// The placeholder in the dynamic error message.
	//
	// example:
	//
	// xxxxx
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The error parameters.
	ErrorArgs []interface{} `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	// The application module.
	//
	// example:
	//
	// {\\"StorageSize\\": \\"3.29 MB\\", \\"FileNum\\": 30}
	Module map[string]interface{} `json:"Module,omitempty" xml:"Module,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CAB8FBB7-F93D-596D-8BA9-FB278ADF9C22
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
	// Indicates whether the request is processed synchronously.
	//
	// example:
	//
	// True
	Synchro *bool `json:"Synchro,omitempty" xml:"Synchro,omitempty"`
}

func (s OnlineAppInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OnlineAppInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *OnlineAppInstanceResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *OnlineAppInstanceResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *OnlineAppInstanceResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *OnlineAppInstanceResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *OnlineAppInstanceResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *OnlineAppInstanceResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *OnlineAppInstanceResponseBody) GetModule() map[string]interface{} {
	return s.Module
}

func (s *OnlineAppInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OnlineAppInstanceResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *OnlineAppInstanceResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *OnlineAppInstanceResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *OnlineAppInstanceResponseBody) SetAccessDeniedDetail(v string) *OnlineAppInstanceResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *OnlineAppInstanceResponseBody) SetAllowRetry(v bool) *OnlineAppInstanceResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *OnlineAppInstanceResponseBody) SetAppName(v string) *OnlineAppInstanceResponseBody {
	s.AppName = &v
	return s
}

func (s *OnlineAppInstanceResponseBody) SetDynamicCode(v string) *OnlineAppInstanceResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *OnlineAppInstanceResponseBody) SetDynamicMessage(v string) *OnlineAppInstanceResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *OnlineAppInstanceResponseBody) SetErrorArgs(v []interface{}) *OnlineAppInstanceResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *OnlineAppInstanceResponseBody) SetModule(v map[string]interface{}) *OnlineAppInstanceResponseBody {
	s.Module = v
	return s
}

func (s *OnlineAppInstanceResponseBody) SetRequestId(v string) *OnlineAppInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *OnlineAppInstanceResponseBody) SetRootErrorCode(v string) *OnlineAppInstanceResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *OnlineAppInstanceResponseBody) SetRootErrorMsg(v string) *OnlineAppInstanceResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *OnlineAppInstanceResponseBody) SetSynchro(v bool) *OnlineAppInstanceResponseBody {
	s.Synchro = &v
	return s
}

func (s *OnlineAppInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
