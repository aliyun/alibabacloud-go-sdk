// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAppDomainCertificateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DeleteAppDomainCertificateResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *DeleteAppDomainCertificateResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *DeleteAppDomainCertificateResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *DeleteAppDomainCertificateResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *DeleteAppDomainCertificateResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *DeleteAppDomainCertificateResponseBody
	GetErrorArgs() []interface{}
	SetModule(v *DeleteAppDomainCertificateResponseBodyModule) *DeleteAppDomainCertificateResponseBody
	GetModule() *DeleteAppDomainCertificateResponseBodyModule
	SetRequestId(v string) *DeleteAppDomainCertificateResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *DeleteAppDomainCertificateResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *DeleteAppDomainCertificateResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *DeleteAppDomainCertificateResponseBody
	GetSynchro() *bool
}

type DeleteAppDomainCertificateResponseBody struct {
	// Details of permission check failure
	//
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// Whether retry is allowed. Values:
	//
	// - false: Not allowed to retry.
	//
	// - true: Allowed to retry.
	//
	// example:
	//
	// False
	AllowRetry *bool `json:"AllowRetry,omitempty" xml:"AllowRetry,omitempty"`
	// Application name. Allows a combination of numbers, letters, and hyphens (-). Must start with a letter, cannot end with a hyphen (-), and must not exceed 36 characters.
	//
	// example:
	//
	// spring-cloud-b
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// Dynamic code, not in use, please ignore
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
	Module *DeleteAppDomainCertificateResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
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

func (s DeleteAppDomainCertificateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAppDomainCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAppDomainCertificateResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DeleteAppDomainCertificateResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *DeleteAppDomainCertificateResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *DeleteAppDomainCertificateResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *DeleteAppDomainCertificateResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *DeleteAppDomainCertificateResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *DeleteAppDomainCertificateResponseBody) GetModule() *DeleteAppDomainCertificateResponseBodyModule {
	return s.Module
}

func (s *DeleteAppDomainCertificateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAppDomainCertificateResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *DeleteAppDomainCertificateResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *DeleteAppDomainCertificateResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *DeleteAppDomainCertificateResponseBody) SetAccessDeniedDetail(v string) *DeleteAppDomainCertificateResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DeleteAppDomainCertificateResponseBody) SetAllowRetry(v bool) *DeleteAppDomainCertificateResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *DeleteAppDomainCertificateResponseBody) SetAppName(v string) *DeleteAppDomainCertificateResponseBody {
	s.AppName = &v
	return s
}

func (s *DeleteAppDomainCertificateResponseBody) SetDynamicCode(v string) *DeleteAppDomainCertificateResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DeleteAppDomainCertificateResponseBody) SetDynamicMessage(v string) *DeleteAppDomainCertificateResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DeleteAppDomainCertificateResponseBody) SetErrorArgs(v []interface{}) *DeleteAppDomainCertificateResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *DeleteAppDomainCertificateResponseBody) SetModule(v *DeleteAppDomainCertificateResponseBodyModule) *DeleteAppDomainCertificateResponseBody {
	s.Module = v
	return s
}

func (s *DeleteAppDomainCertificateResponseBody) SetRequestId(v string) *DeleteAppDomainCertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAppDomainCertificateResponseBody) SetRootErrorCode(v string) *DeleteAppDomainCertificateResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *DeleteAppDomainCertificateResponseBody) SetRootErrorMsg(v string) *DeleteAppDomainCertificateResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *DeleteAppDomainCertificateResponseBody) SetSynchro(v bool) *DeleteAppDomainCertificateResponseBody {
	s.Synchro = &v
	return s
}

func (s *DeleteAppDomainCertificateResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteAppDomainCertificateResponseBodyModule struct {
	// Whether it was successful
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteAppDomainCertificateResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s DeleteAppDomainCertificateResponseBodyModule) GoString() string {
	return s.String()
}

func (s *DeleteAppDomainCertificateResponseBodyModule) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteAppDomainCertificateResponseBodyModule) SetSuccess(v bool) *DeleteAppDomainCertificateResponseBodyModule {
	s.Success = &v
	return s
}

func (s *DeleteAppDomainCertificateResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
