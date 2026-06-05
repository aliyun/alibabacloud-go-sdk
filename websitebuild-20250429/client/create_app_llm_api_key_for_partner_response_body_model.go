// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAppLlmApiKeyForPartnerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CreateAppLlmApiKeyForPartnerResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *CreateAppLlmApiKeyForPartnerResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *CreateAppLlmApiKeyForPartnerResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *CreateAppLlmApiKeyForPartnerResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *CreateAppLlmApiKeyForPartnerResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *CreateAppLlmApiKeyForPartnerResponseBody
	GetErrorArgs() []interface{}
	SetModule(v *CreateAppLlmApiKeyForPartnerResponseBodyModule) *CreateAppLlmApiKeyForPartnerResponseBody
	GetModule() *CreateAppLlmApiKeyForPartnerResponseBodyModule
	SetRequestId(v string) *CreateAppLlmApiKeyForPartnerResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *CreateAppLlmApiKeyForPartnerResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *CreateAppLlmApiKeyForPartnerResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *CreateAppLlmApiKeyForPartnerResponseBody
	GetSynchro() *bool
}

type CreateAppLlmApiKeyForPartnerResponseBody struct {
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
	// SYSTEM.ERROR
	DynamicMessage *string                                         `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorArgs      []interface{}                                   `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	Module         *CreateAppLlmApiKeyForPartnerResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 6C6B99AC-39EC-5350-874C-204128C905E6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// SYSTEM.EROR
	RootErrorCode *string `json:"RootErrorCode,omitempty" xml:"RootErrorCode,omitempty"`
	RootErrorMsg  *string `json:"RootErrorMsg,omitempty" xml:"RootErrorMsg,omitempty"`
	// example:
	//
	// True
	Synchro *bool `json:"Synchro,omitempty" xml:"Synchro,omitempty"`
}

func (s CreateAppLlmApiKeyForPartnerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAppLlmApiKeyForPartnerResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAppLlmApiKeyForPartnerResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CreateAppLlmApiKeyForPartnerResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *CreateAppLlmApiKeyForPartnerResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *CreateAppLlmApiKeyForPartnerResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *CreateAppLlmApiKeyForPartnerResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *CreateAppLlmApiKeyForPartnerResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *CreateAppLlmApiKeyForPartnerResponseBody) GetModule() *CreateAppLlmApiKeyForPartnerResponseBodyModule {
	return s.Module
}

func (s *CreateAppLlmApiKeyForPartnerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAppLlmApiKeyForPartnerResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *CreateAppLlmApiKeyForPartnerResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *CreateAppLlmApiKeyForPartnerResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *CreateAppLlmApiKeyForPartnerResponseBody) SetAccessDeniedDetail(v string) *CreateAppLlmApiKeyForPartnerResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CreateAppLlmApiKeyForPartnerResponseBody) SetAllowRetry(v bool) *CreateAppLlmApiKeyForPartnerResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *CreateAppLlmApiKeyForPartnerResponseBody) SetAppName(v string) *CreateAppLlmApiKeyForPartnerResponseBody {
	s.AppName = &v
	return s
}

func (s *CreateAppLlmApiKeyForPartnerResponseBody) SetDynamicCode(v string) *CreateAppLlmApiKeyForPartnerResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *CreateAppLlmApiKeyForPartnerResponseBody) SetDynamicMessage(v string) *CreateAppLlmApiKeyForPartnerResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *CreateAppLlmApiKeyForPartnerResponseBody) SetErrorArgs(v []interface{}) *CreateAppLlmApiKeyForPartnerResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *CreateAppLlmApiKeyForPartnerResponseBody) SetModule(v *CreateAppLlmApiKeyForPartnerResponseBodyModule) *CreateAppLlmApiKeyForPartnerResponseBody {
	s.Module = v
	return s
}

func (s *CreateAppLlmApiKeyForPartnerResponseBody) SetRequestId(v string) *CreateAppLlmApiKeyForPartnerResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAppLlmApiKeyForPartnerResponseBody) SetRootErrorCode(v string) *CreateAppLlmApiKeyForPartnerResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *CreateAppLlmApiKeyForPartnerResponseBody) SetRootErrorMsg(v string) *CreateAppLlmApiKeyForPartnerResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *CreateAppLlmApiKeyForPartnerResponseBody) SetSynchro(v bool) *CreateAppLlmApiKeyForPartnerResponseBody {
	s.Synchro = &v
	return s
}

func (s *CreateAppLlmApiKeyForPartnerResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateAppLlmApiKeyForPartnerResponseBodyModule struct {
	// example:
	//
	// xxxx-xxxxxx
	EncryptedApiKey *string `json:"EncryptedApiKey,omitempty" xml:"EncryptedApiKey,omitempty"`
}

func (s CreateAppLlmApiKeyForPartnerResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s CreateAppLlmApiKeyForPartnerResponseBodyModule) GoString() string {
	return s.String()
}

func (s *CreateAppLlmApiKeyForPartnerResponseBodyModule) GetEncryptedApiKey() *string {
	return s.EncryptedApiKey
}

func (s *CreateAppLlmApiKeyForPartnerResponseBodyModule) SetEncryptedApiKey(v string) *CreateAppLlmApiKeyForPartnerResponseBodyModule {
	s.EncryptedApiKey = &v
	return s
}

func (s *CreateAppLlmApiKeyForPartnerResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
