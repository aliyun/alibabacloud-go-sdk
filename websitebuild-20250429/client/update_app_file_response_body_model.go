// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAppFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *UpdateAppFileResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *UpdateAppFileResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *UpdateAppFileResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *UpdateAppFileResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *UpdateAppFileResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *UpdateAppFileResponseBody
	GetErrorArgs() []interface{}
	SetModule(v bool) *UpdateAppFileResponseBody
	GetModule() *bool
	SetRequestId(v string) *UpdateAppFileResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *UpdateAppFileResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *UpdateAppFileResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *UpdateAppFileResponseBody
	GetSynchro() *bool
}

type UpdateAppFileResponseBody struct {
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
	// watermark
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// ERROR-oo1
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// example:
	//
	// abc
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

func (s UpdateAppFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAppFileResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAppFileResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *UpdateAppFileResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *UpdateAppFileResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *UpdateAppFileResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *UpdateAppFileResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *UpdateAppFileResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *UpdateAppFileResponseBody) GetModule() *bool {
	return s.Module
}

func (s *UpdateAppFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAppFileResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *UpdateAppFileResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *UpdateAppFileResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *UpdateAppFileResponseBody) SetAccessDeniedDetail(v string) *UpdateAppFileResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *UpdateAppFileResponseBody) SetAllowRetry(v bool) *UpdateAppFileResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *UpdateAppFileResponseBody) SetAppName(v string) *UpdateAppFileResponseBody {
	s.AppName = &v
	return s
}

func (s *UpdateAppFileResponseBody) SetDynamicCode(v string) *UpdateAppFileResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *UpdateAppFileResponseBody) SetDynamicMessage(v string) *UpdateAppFileResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *UpdateAppFileResponseBody) SetErrorArgs(v []interface{}) *UpdateAppFileResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *UpdateAppFileResponseBody) SetModule(v bool) *UpdateAppFileResponseBody {
	s.Module = &v
	return s
}

func (s *UpdateAppFileResponseBody) SetRequestId(v string) *UpdateAppFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAppFileResponseBody) SetRootErrorCode(v string) *UpdateAppFileResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *UpdateAppFileResponseBody) SetRootErrorMsg(v string) *UpdateAppFileResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *UpdateAppFileResponseBody) SetSynchro(v bool) *UpdateAppFileResponseBody {
	s.Synchro = &v
	return s
}

func (s *UpdateAppFileResponseBody) Validate() error {
	return dara.Validate(s)
}
