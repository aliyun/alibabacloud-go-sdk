// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAppCodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *UpdateAppCodeResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *UpdateAppCodeResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *UpdateAppCodeResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *UpdateAppCodeResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *UpdateAppCodeResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *UpdateAppCodeResponseBody
	GetErrorArgs() []interface{}
	SetModule(v bool) *UpdateAppCodeResponseBody
	GetModule() *bool
	SetRequestId(v string) *UpdateAppCodeResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *UpdateAppCodeResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *UpdateAppCodeResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *UpdateAppCodeResponseBody
	GetSynchro() *bool
}

type UpdateAppCodeResponseBody struct {
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
	// ish-intelligence-store-platform-admin-web
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
	// Id of the request
	//
	// example:
	//
	// 6C6B99AC-39EC-5350-874C-204128C905E6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// SYSTEM.ERROR
	RootErrorCode *string `json:"RootErrorCode,omitempty" xml:"RootErrorCode,omitempty"`
	RootErrorMsg  *string `json:"RootErrorMsg,omitempty" xml:"RootErrorMsg,omitempty"`
	// example:
	//
	// True
	Synchro *bool `json:"Synchro,omitempty" xml:"Synchro,omitempty"`
}

func (s UpdateAppCodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAppCodeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAppCodeResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *UpdateAppCodeResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *UpdateAppCodeResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *UpdateAppCodeResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *UpdateAppCodeResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *UpdateAppCodeResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *UpdateAppCodeResponseBody) GetModule() *bool {
	return s.Module
}

func (s *UpdateAppCodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAppCodeResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *UpdateAppCodeResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *UpdateAppCodeResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *UpdateAppCodeResponseBody) SetAccessDeniedDetail(v string) *UpdateAppCodeResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *UpdateAppCodeResponseBody) SetAllowRetry(v bool) *UpdateAppCodeResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *UpdateAppCodeResponseBody) SetAppName(v string) *UpdateAppCodeResponseBody {
	s.AppName = &v
	return s
}

func (s *UpdateAppCodeResponseBody) SetDynamicCode(v string) *UpdateAppCodeResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *UpdateAppCodeResponseBody) SetDynamicMessage(v string) *UpdateAppCodeResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *UpdateAppCodeResponseBody) SetErrorArgs(v []interface{}) *UpdateAppCodeResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *UpdateAppCodeResponseBody) SetModule(v bool) *UpdateAppCodeResponseBody {
	s.Module = &v
	return s
}

func (s *UpdateAppCodeResponseBody) SetRequestId(v string) *UpdateAppCodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAppCodeResponseBody) SetRootErrorCode(v string) *UpdateAppCodeResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *UpdateAppCodeResponseBody) SetRootErrorMsg(v string) *UpdateAppCodeResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *UpdateAppCodeResponseBody) SetSynchro(v bool) *UpdateAppCodeResponseBody {
	s.Synchro = &v
	return s
}

func (s *UpdateAppCodeResponseBody) Validate() error {
	return dara.Validate(s)
}
