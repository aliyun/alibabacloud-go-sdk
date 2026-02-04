// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetAppDomainCertificateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *SetAppDomainCertificateResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *SetAppDomainCertificateResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *SetAppDomainCertificateResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *SetAppDomainCertificateResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *SetAppDomainCertificateResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *SetAppDomainCertificateResponseBody
	GetErrorArgs() []interface{}
	SetModule(v *SetAppDomainCertificateResponseBodyModule) *SetAppDomainCertificateResponseBody
	GetModule() *SetAppDomainCertificateResponseBodyModule
	SetRequestId(v string) *SetAppDomainCertificateResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *SetAppDomainCertificateResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *SetAppDomainCertificateResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *SetAppDomainCertificateResponseBody
	GetSynchro() *bool
}

type SetAppDomainCertificateResponseBody struct {
	// No permission information
	//
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// Whether to allow retry. Values:
	//
	// - false: Do not allow retry.
	//
	// - true: Allow retry.
	//
	// example:
	//
	// False
	AllowRetry *bool `json:"AllowRetry,omitempty" xml:"AllowRetry,omitempty"`
	// 应用名称
	//
	// example:
	//
	// mar
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// Dynamic error code
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
	// Return error parameters
	ErrorArgs []interface{} `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	// Response data
	Module *SetAppDomainCertificateResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
	// ID of the request
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
	// Exception message
	//
	// example:
	//
	// 系统异常
	RootErrorMsg *string `json:"RootErrorMsg,omitempty" xml:"RootErrorMsg,omitempty"`
	// Whether it is processed synchronously
	//
	// example:
	//
	// True
	Synchro *bool `json:"Synchro,omitempty" xml:"Synchro,omitempty"`
}

func (s SetAppDomainCertificateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetAppDomainCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *SetAppDomainCertificateResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *SetAppDomainCertificateResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *SetAppDomainCertificateResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *SetAppDomainCertificateResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *SetAppDomainCertificateResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *SetAppDomainCertificateResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *SetAppDomainCertificateResponseBody) GetModule() *SetAppDomainCertificateResponseBodyModule {
	return s.Module
}

func (s *SetAppDomainCertificateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetAppDomainCertificateResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *SetAppDomainCertificateResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *SetAppDomainCertificateResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *SetAppDomainCertificateResponseBody) SetAccessDeniedDetail(v string) *SetAppDomainCertificateResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *SetAppDomainCertificateResponseBody) SetAllowRetry(v bool) *SetAppDomainCertificateResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *SetAppDomainCertificateResponseBody) SetAppName(v string) *SetAppDomainCertificateResponseBody {
	s.AppName = &v
	return s
}

func (s *SetAppDomainCertificateResponseBody) SetDynamicCode(v string) *SetAppDomainCertificateResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *SetAppDomainCertificateResponseBody) SetDynamicMessage(v string) *SetAppDomainCertificateResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *SetAppDomainCertificateResponseBody) SetErrorArgs(v []interface{}) *SetAppDomainCertificateResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *SetAppDomainCertificateResponseBody) SetModule(v *SetAppDomainCertificateResponseBodyModule) *SetAppDomainCertificateResponseBody {
	s.Module = v
	return s
}

func (s *SetAppDomainCertificateResponseBody) SetRequestId(v string) *SetAppDomainCertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetAppDomainCertificateResponseBody) SetRootErrorCode(v string) *SetAppDomainCertificateResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *SetAppDomainCertificateResponseBody) SetRootErrorMsg(v string) *SetAppDomainCertificateResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *SetAppDomainCertificateResponseBody) SetSynchro(v bool) *SetAppDomainCertificateResponseBody {
	s.Synchro = &v
	return s
}

func (s *SetAppDomainCertificateResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SetAppDomainCertificateResponseBodyModule struct {
	// Whether it is successful
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SetAppDomainCertificateResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s SetAppDomainCertificateResponseBodyModule) GoString() string {
	return s.String()
}

func (s *SetAppDomainCertificateResponseBodyModule) GetSuccess() *bool {
	return s.Success
}

func (s *SetAppDomainCertificateResponseBodyModule) SetSuccess(v bool) *SetAppDomainCertificateResponseBodyModule {
	s.Success = &v
	return s
}

func (s *SetAppDomainCertificateResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
