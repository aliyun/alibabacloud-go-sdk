// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckAppVerifyCodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CheckAppVerifyCodeResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *CheckAppVerifyCodeResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *CheckAppVerifyCodeResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *CheckAppVerifyCodeResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *CheckAppVerifyCodeResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *CheckAppVerifyCodeResponseBody
	GetErrorArgs() []interface{}
	SetModule(v *CheckAppVerifyCodeResponseBodyModule) *CheckAppVerifyCodeResponseBody
	GetModule() *CheckAppVerifyCodeResponseBodyModule
	SetRequestId(v string) *CheckAppVerifyCodeResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *CheckAppVerifyCodeResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *CheckAppVerifyCodeResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *CheckAppVerifyCodeResponseBody
	GetSynchro() *bool
}

type CheckAppVerifyCodeResponseBody struct {
	// The detailed reason why access is denied.
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
	// The dynamic error message, which is used to replace the %s variable in the ErrMessage response element.
	//
	// example:
	//
	// SYSTEM_ERROR
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The error parameters.
	ErrorArgs []interface{} `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	// The response data.
	Module *CheckAppVerifyCodeResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
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
	// The error message.
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

func (s CheckAppVerifyCodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckAppVerifyCodeResponseBody) GoString() string {
	return s.String()
}

func (s *CheckAppVerifyCodeResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CheckAppVerifyCodeResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *CheckAppVerifyCodeResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *CheckAppVerifyCodeResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *CheckAppVerifyCodeResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *CheckAppVerifyCodeResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *CheckAppVerifyCodeResponseBody) GetModule() *CheckAppVerifyCodeResponseBodyModule {
	return s.Module
}

func (s *CheckAppVerifyCodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckAppVerifyCodeResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *CheckAppVerifyCodeResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *CheckAppVerifyCodeResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *CheckAppVerifyCodeResponseBody) SetAccessDeniedDetail(v string) *CheckAppVerifyCodeResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CheckAppVerifyCodeResponseBody) SetAllowRetry(v bool) *CheckAppVerifyCodeResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *CheckAppVerifyCodeResponseBody) SetAppName(v string) *CheckAppVerifyCodeResponseBody {
	s.AppName = &v
	return s
}

func (s *CheckAppVerifyCodeResponseBody) SetDynamicCode(v string) *CheckAppVerifyCodeResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *CheckAppVerifyCodeResponseBody) SetDynamicMessage(v string) *CheckAppVerifyCodeResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *CheckAppVerifyCodeResponseBody) SetErrorArgs(v []interface{}) *CheckAppVerifyCodeResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *CheckAppVerifyCodeResponseBody) SetModule(v *CheckAppVerifyCodeResponseBodyModule) *CheckAppVerifyCodeResponseBody {
	s.Module = v
	return s
}

func (s *CheckAppVerifyCodeResponseBody) SetRequestId(v string) *CheckAppVerifyCodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckAppVerifyCodeResponseBody) SetRootErrorCode(v string) *CheckAppVerifyCodeResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *CheckAppVerifyCodeResponseBody) SetRootErrorMsg(v string) *CheckAppVerifyCodeResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *CheckAppVerifyCodeResponseBody) SetSynchro(v bool) *CheckAppVerifyCodeResponseBody {
	s.Synchro = &v
	return s
}

func (s *CheckAppVerifyCodeResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CheckAppVerifyCodeResponseBodyModule struct {
	// The masked phone number or email address, used for frontend display.
	//
	// example:
	//
	// 1
	MaskedTarget *string `json:"MaskedTarget,omitempty" xml:"MaskedTarget,omitempty"`
	// The recipient ID.
	//
	// example:
	//
	// 1
	RecipientId *string `json:"RecipientId,omitempty" xml:"RecipientId,omitempty"`
}

func (s CheckAppVerifyCodeResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s CheckAppVerifyCodeResponseBodyModule) GoString() string {
	return s.String()
}

func (s *CheckAppVerifyCodeResponseBodyModule) GetMaskedTarget() *string {
	return s.MaskedTarget
}

func (s *CheckAppVerifyCodeResponseBodyModule) GetRecipientId() *string {
	return s.RecipientId
}

func (s *CheckAppVerifyCodeResponseBodyModule) SetMaskedTarget(v string) *CheckAppVerifyCodeResponseBodyModule {
	s.MaskedTarget = &v
	return s
}

func (s *CheckAppVerifyCodeResponseBodyModule) SetRecipientId(v string) *CheckAppVerifyCodeResponseBodyModule {
	s.RecipientId = &v
	return s
}

func (s *CheckAppVerifyCodeResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
