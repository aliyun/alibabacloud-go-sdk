// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCopyAppPluginConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CopyAppPluginConfigResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *CopyAppPluginConfigResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *CopyAppPluginConfigResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *CopyAppPluginConfigResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *CopyAppPluginConfigResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *CopyAppPluginConfigResponseBody
	GetErrorArgs() []interface{}
	SetModule(v bool) *CopyAppPluginConfigResponseBody
	GetModule() *bool
	SetRequestId(v string) *CopyAppPluginConfigResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *CopyAppPluginConfigResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *CopyAppPluginConfigResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *CopyAppPluginConfigResponseBody
	GetSynchro() *bool
}

type CopyAppPluginConfigResponseBody struct {
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

func (s CopyAppPluginConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CopyAppPluginConfigResponseBody) GoString() string {
	return s.String()
}

func (s *CopyAppPluginConfigResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CopyAppPluginConfigResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *CopyAppPluginConfigResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *CopyAppPluginConfigResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *CopyAppPluginConfigResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *CopyAppPluginConfigResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *CopyAppPluginConfigResponseBody) GetModule() *bool {
	return s.Module
}

func (s *CopyAppPluginConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CopyAppPluginConfigResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *CopyAppPluginConfigResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *CopyAppPluginConfigResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *CopyAppPluginConfigResponseBody) SetAccessDeniedDetail(v string) *CopyAppPluginConfigResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CopyAppPluginConfigResponseBody) SetAllowRetry(v bool) *CopyAppPluginConfigResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *CopyAppPluginConfigResponseBody) SetAppName(v string) *CopyAppPluginConfigResponseBody {
	s.AppName = &v
	return s
}

func (s *CopyAppPluginConfigResponseBody) SetDynamicCode(v string) *CopyAppPluginConfigResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *CopyAppPluginConfigResponseBody) SetDynamicMessage(v string) *CopyAppPluginConfigResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *CopyAppPluginConfigResponseBody) SetErrorArgs(v []interface{}) *CopyAppPluginConfigResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *CopyAppPluginConfigResponseBody) SetModule(v bool) *CopyAppPluginConfigResponseBody {
	s.Module = &v
	return s
}

func (s *CopyAppPluginConfigResponseBody) SetRequestId(v string) *CopyAppPluginConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *CopyAppPluginConfigResponseBody) SetRootErrorCode(v string) *CopyAppPluginConfigResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *CopyAppPluginConfigResponseBody) SetRootErrorMsg(v string) *CopyAppPluginConfigResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *CopyAppPluginConfigResponseBody) SetSynchro(v bool) *CopyAppPluginConfigResponseBody {
	s.Synchro = &v
	return s
}

func (s *CopyAppPluginConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
