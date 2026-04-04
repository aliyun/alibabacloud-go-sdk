// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAppAssistantAgentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CreateAppAssistantAgentResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *CreateAppAssistantAgentResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *CreateAppAssistantAgentResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *CreateAppAssistantAgentResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *CreateAppAssistantAgentResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *CreateAppAssistantAgentResponseBody
	GetErrorArgs() []interface{}
	SetModule(v *CreateAppAssistantAgentResponseBodyModule) *CreateAppAssistantAgentResponseBody
	GetModule() *CreateAppAssistantAgentResponseBodyModule
	SetRequestId(v string) *CreateAppAssistantAgentResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *CreateAppAssistantAgentResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *CreateAppAssistantAgentResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *CreateAppAssistantAgentResponseBody
	GetSynchro() *bool
}

type CreateAppAssistantAgentResponseBody struct {
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
	DynamicMessage *string                                    `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorArgs      []interface{}                              `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	Module         *CreateAppAssistantAgentResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
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

func (s CreateAppAssistantAgentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAppAssistantAgentResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAppAssistantAgentResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CreateAppAssistantAgentResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *CreateAppAssistantAgentResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *CreateAppAssistantAgentResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *CreateAppAssistantAgentResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *CreateAppAssistantAgentResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *CreateAppAssistantAgentResponseBody) GetModule() *CreateAppAssistantAgentResponseBodyModule {
	return s.Module
}

func (s *CreateAppAssistantAgentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAppAssistantAgentResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *CreateAppAssistantAgentResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *CreateAppAssistantAgentResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *CreateAppAssistantAgentResponseBody) SetAccessDeniedDetail(v string) *CreateAppAssistantAgentResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CreateAppAssistantAgentResponseBody) SetAllowRetry(v bool) *CreateAppAssistantAgentResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *CreateAppAssistantAgentResponseBody) SetAppName(v string) *CreateAppAssistantAgentResponseBody {
	s.AppName = &v
	return s
}

func (s *CreateAppAssistantAgentResponseBody) SetDynamicCode(v string) *CreateAppAssistantAgentResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *CreateAppAssistantAgentResponseBody) SetDynamicMessage(v string) *CreateAppAssistantAgentResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *CreateAppAssistantAgentResponseBody) SetErrorArgs(v []interface{}) *CreateAppAssistantAgentResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *CreateAppAssistantAgentResponseBody) SetModule(v *CreateAppAssistantAgentResponseBodyModule) *CreateAppAssistantAgentResponseBody {
	s.Module = v
	return s
}

func (s *CreateAppAssistantAgentResponseBody) SetRequestId(v string) *CreateAppAssistantAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAppAssistantAgentResponseBody) SetRootErrorCode(v string) *CreateAppAssistantAgentResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *CreateAppAssistantAgentResponseBody) SetRootErrorMsg(v string) *CreateAppAssistantAgentResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *CreateAppAssistantAgentResponseBody) SetSynchro(v bool) *CreateAppAssistantAgentResponseBody {
	s.Synchro = &v
	return s
}

func (s *CreateAppAssistantAgentResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateAppAssistantAgentResponseBodyModule struct {
	// example:
	//
	// duanwei@qianrutest
	AgentId   *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	AgentName *string `json:"AgentName,omitempty" xml:"AgentName,omitempty"`
	// example:
	//
	// WD20250703155602000001
	BizId       *string                                               `json:"BizId,omitempty" xml:"BizId,omitempty"`
	Credential  *CreateAppAssistantAgentResponseBodyModuleCredential  `json:"Credential,omitempty" xml:"Credential,omitempty" type:"Struct"`
	EmbedConfig *CreateAppAssistantAgentResponseBodyModuleEmbedConfig `json:"EmbedConfig,omitempty" xml:"EmbedConfig,omitempty" type:"Struct"`
	ExtraParams map[string]*string                                    `json:"ExtraParams,omitempty" xml:"ExtraParams,omitempty"`
	// example:
	//
	// 1740479834
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 2025-08-28T02:25:41Z
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// WA12313131313
	PlatformAppId *string `json:"PlatformAppId,omitempty" xml:"PlatformAppId,omitempty"`
	// example:
	//
	// VMWARE
	PlatformType *string `json:"PlatformType,omitempty" xml:"PlatformType,omitempty"`
	// example:
	//
	// NORMAL
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 1231313131
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s CreateAppAssistantAgentResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s CreateAppAssistantAgentResponseBodyModule) GoString() string {
	return s.String()
}

func (s *CreateAppAssistantAgentResponseBodyModule) GetAgentId() *string {
	return s.AgentId
}

func (s *CreateAppAssistantAgentResponseBodyModule) GetAgentName() *string {
	return s.AgentName
}

func (s *CreateAppAssistantAgentResponseBodyModule) GetBizId() *string {
	return s.BizId
}

func (s *CreateAppAssistantAgentResponseBodyModule) GetCredential() *CreateAppAssistantAgentResponseBodyModuleCredential {
	return s.Credential
}

func (s *CreateAppAssistantAgentResponseBodyModule) GetEmbedConfig() *CreateAppAssistantAgentResponseBodyModuleEmbedConfig {
	return s.EmbedConfig
}

func (s *CreateAppAssistantAgentResponseBodyModule) GetExtraParams() map[string]*string {
	return s.ExtraParams
}

func (s *CreateAppAssistantAgentResponseBodyModule) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *CreateAppAssistantAgentResponseBodyModule) GetGmtModified() *string {
	return s.GmtModified
}

func (s *CreateAppAssistantAgentResponseBodyModule) GetPlatformAppId() *string {
	return s.PlatformAppId
}

func (s *CreateAppAssistantAgentResponseBodyModule) GetPlatformType() *string {
	return s.PlatformType
}

func (s *CreateAppAssistantAgentResponseBodyModule) GetStatus() *string {
	return s.Status
}

func (s *CreateAppAssistantAgentResponseBodyModule) GetUserId() *string {
	return s.UserId
}

func (s *CreateAppAssistantAgentResponseBodyModule) SetAgentId(v string) *CreateAppAssistantAgentResponseBodyModule {
	s.AgentId = &v
	return s
}

func (s *CreateAppAssistantAgentResponseBodyModule) SetAgentName(v string) *CreateAppAssistantAgentResponseBodyModule {
	s.AgentName = &v
	return s
}

func (s *CreateAppAssistantAgentResponseBodyModule) SetBizId(v string) *CreateAppAssistantAgentResponseBodyModule {
	s.BizId = &v
	return s
}

func (s *CreateAppAssistantAgentResponseBodyModule) SetCredential(v *CreateAppAssistantAgentResponseBodyModuleCredential) *CreateAppAssistantAgentResponseBodyModule {
	s.Credential = v
	return s
}

func (s *CreateAppAssistantAgentResponseBodyModule) SetEmbedConfig(v *CreateAppAssistantAgentResponseBodyModuleEmbedConfig) *CreateAppAssistantAgentResponseBodyModule {
	s.EmbedConfig = v
	return s
}

func (s *CreateAppAssistantAgentResponseBodyModule) SetExtraParams(v map[string]*string) *CreateAppAssistantAgentResponseBodyModule {
	s.ExtraParams = v
	return s
}

func (s *CreateAppAssistantAgentResponseBodyModule) SetGmtCreate(v string) *CreateAppAssistantAgentResponseBodyModule {
	s.GmtCreate = &v
	return s
}

func (s *CreateAppAssistantAgentResponseBodyModule) SetGmtModified(v string) *CreateAppAssistantAgentResponseBodyModule {
	s.GmtModified = &v
	return s
}

func (s *CreateAppAssistantAgentResponseBodyModule) SetPlatformAppId(v string) *CreateAppAssistantAgentResponseBodyModule {
	s.PlatformAppId = &v
	return s
}

func (s *CreateAppAssistantAgentResponseBodyModule) SetPlatformType(v string) *CreateAppAssistantAgentResponseBodyModule {
	s.PlatformType = &v
	return s
}

func (s *CreateAppAssistantAgentResponseBodyModule) SetStatus(v string) *CreateAppAssistantAgentResponseBodyModule {
	s.Status = &v
	return s
}

func (s *CreateAppAssistantAgentResponseBodyModule) SetUserId(v string) *CreateAppAssistantAgentResponseBodyModule {
	s.UserId = &v
	return s
}

func (s *CreateAppAssistantAgentResponseBodyModule) Validate() error {
	if s.Credential != nil {
		if err := s.Credential.Validate(); err != nil {
			return err
		}
	}
	if s.EmbedConfig != nil {
		if err := s.EmbedConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateAppAssistantAgentResponseBodyModuleCredential struct {
	Extra map[string]*string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	// example:
	//
	// Test
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s CreateAppAssistantAgentResponseBodyModuleCredential) String() string {
	return dara.Prettify(s)
}

func (s CreateAppAssistantAgentResponseBodyModuleCredential) GoString() string {
	return s.String()
}

func (s *CreateAppAssistantAgentResponseBodyModuleCredential) GetExtra() map[string]*string {
	return s.Extra
}

func (s *CreateAppAssistantAgentResponseBodyModuleCredential) GetUsername() *string {
	return s.Username
}

func (s *CreateAppAssistantAgentResponseBodyModuleCredential) SetExtra(v map[string]*string) *CreateAppAssistantAgentResponseBodyModuleCredential {
	s.Extra = v
	return s
}

func (s *CreateAppAssistantAgentResponseBodyModuleCredential) SetUsername(v string) *CreateAppAssistantAgentResponseBodyModuleCredential {
	s.Username = &v
	return s
}

func (s *CreateAppAssistantAgentResponseBodyModuleCredential) Validate() error {
	return dara.Validate(s)
}

type CreateAppAssistantAgentResponseBodyModuleEmbedConfig struct {
	Extra     map[string]*string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	RawScript *string            `json:"RawScript,omitempty" xml:"RawScript,omitempty"`
}

func (s CreateAppAssistantAgentResponseBodyModuleEmbedConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateAppAssistantAgentResponseBodyModuleEmbedConfig) GoString() string {
	return s.String()
}

func (s *CreateAppAssistantAgentResponseBodyModuleEmbedConfig) GetExtra() map[string]*string {
	return s.Extra
}

func (s *CreateAppAssistantAgentResponseBodyModuleEmbedConfig) GetRawScript() *string {
	return s.RawScript
}

func (s *CreateAppAssistantAgentResponseBodyModuleEmbedConfig) SetExtra(v map[string]*string) *CreateAppAssistantAgentResponseBodyModuleEmbedConfig {
	s.Extra = v
	return s
}

func (s *CreateAppAssistantAgentResponseBodyModuleEmbedConfig) SetRawScript(v string) *CreateAppAssistantAgentResponseBodyModuleEmbedConfig {
	s.RawScript = &v
	return s
}

func (s *CreateAppAssistantAgentResponseBodyModuleEmbedConfig) Validate() error {
	return dara.Validate(s)
}
