// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAppAssistantAgentSsoLoginResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CreateAppAssistantAgentSsoLoginResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *CreateAppAssistantAgentSsoLoginResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *CreateAppAssistantAgentSsoLoginResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *CreateAppAssistantAgentSsoLoginResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *CreateAppAssistantAgentSsoLoginResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *CreateAppAssistantAgentSsoLoginResponseBody
	GetErrorArgs() []interface{}
	SetModule(v *CreateAppAssistantAgentSsoLoginResponseBodyModule) *CreateAppAssistantAgentSsoLoginResponseBody
	GetModule() *CreateAppAssistantAgentSsoLoginResponseBodyModule
	SetRequestId(v string) *CreateAppAssistantAgentSsoLoginResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *CreateAppAssistantAgentSsoLoginResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *CreateAppAssistantAgentSsoLoginResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *CreateAppAssistantAgentSsoLoginResponseBody
	GetSynchro() *bool
}

type CreateAppAssistantAgentSsoLoginResponseBody struct {
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
	// dewuApp
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// ERROR-oo1
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// example:
	//
	// SYSTEM_ERROR
	DynamicMessage *string                                            `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorArgs      []interface{}                                      `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	Module         *CreateAppAssistantAgentSsoLoginResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
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

func (s CreateAppAssistantAgentSsoLoginResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAppAssistantAgentSsoLoginResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAppAssistantAgentSsoLoginResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CreateAppAssistantAgentSsoLoginResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *CreateAppAssistantAgentSsoLoginResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *CreateAppAssistantAgentSsoLoginResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *CreateAppAssistantAgentSsoLoginResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *CreateAppAssistantAgentSsoLoginResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *CreateAppAssistantAgentSsoLoginResponseBody) GetModule() *CreateAppAssistantAgentSsoLoginResponseBodyModule {
	return s.Module
}

func (s *CreateAppAssistantAgentSsoLoginResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAppAssistantAgentSsoLoginResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *CreateAppAssistantAgentSsoLoginResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *CreateAppAssistantAgentSsoLoginResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *CreateAppAssistantAgentSsoLoginResponseBody) SetAccessDeniedDetail(v string) *CreateAppAssistantAgentSsoLoginResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CreateAppAssistantAgentSsoLoginResponseBody) SetAllowRetry(v bool) *CreateAppAssistantAgentSsoLoginResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *CreateAppAssistantAgentSsoLoginResponseBody) SetAppName(v string) *CreateAppAssistantAgentSsoLoginResponseBody {
	s.AppName = &v
	return s
}

func (s *CreateAppAssistantAgentSsoLoginResponseBody) SetDynamicCode(v string) *CreateAppAssistantAgentSsoLoginResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *CreateAppAssistantAgentSsoLoginResponseBody) SetDynamicMessage(v string) *CreateAppAssistantAgentSsoLoginResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *CreateAppAssistantAgentSsoLoginResponseBody) SetErrorArgs(v []interface{}) *CreateAppAssistantAgentSsoLoginResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *CreateAppAssistantAgentSsoLoginResponseBody) SetModule(v *CreateAppAssistantAgentSsoLoginResponseBodyModule) *CreateAppAssistantAgentSsoLoginResponseBody {
	s.Module = v
	return s
}

func (s *CreateAppAssistantAgentSsoLoginResponseBody) SetRequestId(v string) *CreateAppAssistantAgentSsoLoginResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAppAssistantAgentSsoLoginResponseBody) SetRootErrorCode(v string) *CreateAppAssistantAgentSsoLoginResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *CreateAppAssistantAgentSsoLoginResponseBody) SetRootErrorMsg(v string) *CreateAppAssistantAgentSsoLoginResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *CreateAppAssistantAgentSsoLoginResponseBody) SetSynchro(v bool) *CreateAppAssistantAgentSsoLoginResponseBody {
	s.Synchro = &v
	return s
}

func (s *CreateAppAssistantAgentSsoLoginResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateAppAssistantAgentSsoLoginResponseBodyModule struct {
	// example:
	//
	// 2025-07-30T16:00Z
	ExpireTime *int64             `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	Extra      map[string]*string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	// example:
	//
	// VMWARE
	PlatformType *string `json:"PlatformType,omitempty" xml:"PlatformType,omitempty"`
	// example:
	//
	// https://sso.agent
	SsoUrl *string `json:"SsoUrl,omitempty" xml:"SsoUrl,omitempty"`
}

func (s CreateAppAssistantAgentSsoLoginResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s CreateAppAssistantAgentSsoLoginResponseBodyModule) GoString() string {
	return s.String()
}

func (s *CreateAppAssistantAgentSsoLoginResponseBodyModule) GetExpireTime() *int64 {
	return s.ExpireTime
}

func (s *CreateAppAssistantAgentSsoLoginResponseBodyModule) GetExtra() map[string]*string {
	return s.Extra
}

func (s *CreateAppAssistantAgentSsoLoginResponseBodyModule) GetPlatformType() *string {
	return s.PlatformType
}

func (s *CreateAppAssistantAgentSsoLoginResponseBodyModule) GetSsoUrl() *string {
	return s.SsoUrl
}

func (s *CreateAppAssistantAgentSsoLoginResponseBodyModule) SetExpireTime(v int64) *CreateAppAssistantAgentSsoLoginResponseBodyModule {
	s.ExpireTime = &v
	return s
}

func (s *CreateAppAssistantAgentSsoLoginResponseBodyModule) SetExtra(v map[string]*string) *CreateAppAssistantAgentSsoLoginResponseBodyModule {
	s.Extra = v
	return s
}

func (s *CreateAppAssistantAgentSsoLoginResponseBodyModule) SetPlatformType(v string) *CreateAppAssistantAgentSsoLoginResponseBodyModule {
	s.PlatformType = &v
	return s
}

func (s *CreateAppAssistantAgentSsoLoginResponseBodyModule) SetSsoUrl(v string) *CreateAppAssistantAgentSsoLoginResponseBodyModule {
	s.SsoUrl = &v
	return s
}

func (s *CreateAppAssistantAgentSsoLoginResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
