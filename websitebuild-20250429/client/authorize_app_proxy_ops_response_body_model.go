// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthorizeAppProxyOpsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *AuthorizeAppProxyOpsResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *AuthorizeAppProxyOpsResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *AuthorizeAppProxyOpsResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *AuthorizeAppProxyOpsResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *AuthorizeAppProxyOpsResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *AuthorizeAppProxyOpsResponseBody
	GetErrorArgs() []interface{}
	SetModule(v *AuthorizeAppProxyOpsResponseBodyModule) *AuthorizeAppProxyOpsResponseBody
	GetModule() *AuthorizeAppProxyOpsResponseBodyModule
	SetRequestId(v string) *AuthorizeAppProxyOpsResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *AuthorizeAppProxyOpsResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *AuthorizeAppProxyOpsResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *AuthorizeAppProxyOpsResponseBody
	GetSynchro() *bool
}

type AuthorizeAppProxyOpsResponseBody struct {
	// The deprecated parameter.
	//
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// Indicates whether retry is allowed.
	//
	// example:
	//
	// False
	AllowRetry *bool `json:"AllowRetry,omitempty" xml:"AllowRetry,omitempty"`
	// The application name.
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
	// The dynamic error message.
	//
	// example:
	//
	// xxx
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The error parameters.
	ErrorArgs []interface{} `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	// The response object.
	Module *AuthorizeAppProxyOpsResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
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
	// The root error message.
	//
	// example:
	//
	// 系统异常
	RootErrorMsg *string `json:"RootErrorMsg,omitempty" xml:"RootErrorMsg,omitempty"`
	// Indicates whether the request is synchronously processed.
	//
	// example:
	//
	// True
	Synchro *bool `json:"Synchro,omitempty" xml:"Synchro,omitempty"`
}

func (s AuthorizeAppProxyOpsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AuthorizeAppProxyOpsResponseBody) GoString() string {
	return s.String()
}

func (s *AuthorizeAppProxyOpsResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *AuthorizeAppProxyOpsResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *AuthorizeAppProxyOpsResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *AuthorizeAppProxyOpsResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *AuthorizeAppProxyOpsResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *AuthorizeAppProxyOpsResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *AuthorizeAppProxyOpsResponseBody) GetModule() *AuthorizeAppProxyOpsResponseBodyModule {
	return s.Module
}

func (s *AuthorizeAppProxyOpsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AuthorizeAppProxyOpsResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *AuthorizeAppProxyOpsResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *AuthorizeAppProxyOpsResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *AuthorizeAppProxyOpsResponseBody) SetAccessDeniedDetail(v string) *AuthorizeAppProxyOpsResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *AuthorizeAppProxyOpsResponseBody) SetAllowRetry(v bool) *AuthorizeAppProxyOpsResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *AuthorizeAppProxyOpsResponseBody) SetAppName(v string) *AuthorizeAppProxyOpsResponseBody {
	s.AppName = &v
	return s
}

func (s *AuthorizeAppProxyOpsResponseBody) SetDynamicCode(v string) *AuthorizeAppProxyOpsResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *AuthorizeAppProxyOpsResponseBody) SetDynamicMessage(v string) *AuthorizeAppProxyOpsResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *AuthorizeAppProxyOpsResponseBody) SetErrorArgs(v []interface{}) *AuthorizeAppProxyOpsResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *AuthorizeAppProxyOpsResponseBody) SetModule(v *AuthorizeAppProxyOpsResponseBodyModule) *AuthorizeAppProxyOpsResponseBody {
	s.Module = v
	return s
}

func (s *AuthorizeAppProxyOpsResponseBody) SetRequestId(v string) *AuthorizeAppProxyOpsResponseBody {
	s.RequestId = &v
	return s
}

func (s *AuthorizeAppProxyOpsResponseBody) SetRootErrorCode(v string) *AuthorizeAppProxyOpsResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *AuthorizeAppProxyOpsResponseBody) SetRootErrorMsg(v string) *AuthorizeAppProxyOpsResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *AuthorizeAppProxyOpsResponseBody) SetSynchro(v bool) *AuthorizeAppProxyOpsResponseBody {
	s.Synchro = &v
	return s
}

func (s *AuthorizeAppProxyOpsResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AuthorizeAppProxyOpsResponseBodyModule struct {
	// Indicates whether the service-linked role is authorized.
	//
	// example:
	//
	// Y
	Authorized *bool `json:"Authorized,omitempty" xml:"Authorized,omitempty"`
	// The RAM service role authorization link returned when the service-linked role is not authorized.
	//
	// example:
	//
	// url
	RamAuthLink *string `json:"RamAuthLink,omitempty" xml:"RamAuthLink,omitempty"`
	// The authorization URL built with tmpTicket, returned when the service-linked role is authorized.
	//
	// example:
	//
	// 12345rert
	TmpTicket *string `json:"TmpTicket,omitempty" xml:"TmpTicket,omitempty"`
}

func (s AuthorizeAppProxyOpsResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s AuthorizeAppProxyOpsResponseBodyModule) GoString() string {
	return s.String()
}

func (s *AuthorizeAppProxyOpsResponseBodyModule) GetAuthorized() *bool {
	return s.Authorized
}

func (s *AuthorizeAppProxyOpsResponseBodyModule) GetRamAuthLink() *string {
	return s.RamAuthLink
}

func (s *AuthorizeAppProxyOpsResponseBodyModule) GetTmpTicket() *string {
	return s.TmpTicket
}

func (s *AuthorizeAppProxyOpsResponseBodyModule) SetAuthorized(v bool) *AuthorizeAppProxyOpsResponseBodyModule {
	s.Authorized = &v
	return s
}

func (s *AuthorizeAppProxyOpsResponseBodyModule) SetRamAuthLink(v string) *AuthorizeAppProxyOpsResponseBodyModule {
	s.RamAuthLink = &v
	return s
}

func (s *AuthorizeAppProxyOpsResponseBodyModule) SetTmpTicket(v string) *AuthorizeAppProxyOpsResponseBodyModule {
	s.TmpTicket = &v
	return s
}

func (s *AuthorizeAppProxyOpsResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
