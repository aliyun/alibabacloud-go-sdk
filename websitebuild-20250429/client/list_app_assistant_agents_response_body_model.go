// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAppAssistantAgentsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ListAppAssistantAgentsResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *ListAppAssistantAgentsResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *ListAppAssistantAgentsResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *ListAppAssistantAgentsResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *ListAppAssistantAgentsResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *ListAppAssistantAgentsResponseBody
	GetErrorArgs() []interface{}
	SetModule(v []*ListAppAssistantAgentsResponseBodyModule) *ListAppAssistantAgentsResponseBody
	GetModule() []*ListAppAssistantAgentsResponseBodyModule
	SetRequestId(v string) *ListAppAssistantAgentsResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *ListAppAssistantAgentsResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *ListAppAssistantAgentsResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *ListAppAssistantAgentsResponseBody
	GetSynchro() *bool
}

type ListAppAssistantAgentsResponseBody struct {
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
	// spring-cloud-b
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// ERROR-oo1
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// example:
	//
	// SYSTEM_ERROR
	DynamicMessage *string                                     `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorArgs      []interface{}                               `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	Module         []*ListAppAssistantAgentsResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Repeated"`
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

func (s ListAppAssistantAgentsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAppAssistantAgentsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAppAssistantAgentsResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ListAppAssistantAgentsResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *ListAppAssistantAgentsResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *ListAppAssistantAgentsResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *ListAppAssistantAgentsResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *ListAppAssistantAgentsResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *ListAppAssistantAgentsResponseBody) GetModule() []*ListAppAssistantAgentsResponseBodyModule {
	return s.Module
}

func (s *ListAppAssistantAgentsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAppAssistantAgentsResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *ListAppAssistantAgentsResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *ListAppAssistantAgentsResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *ListAppAssistantAgentsResponseBody) SetAccessDeniedDetail(v string) *ListAppAssistantAgentsResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListAppAssistantAgentsResponseBody) SetAllowRetry(v bool) *ListAppAssistantAgentsResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *ListAppAssistantAgentsResponseBody) SetAppName(v string) *ListAppAssistantAgentsResponseBody {
	s.AppName = &v
	return s
}

func (s *ListAppAssistantAgentsResponseBody) SetDynamicCode(v string) *ListAppAssistantAgentsResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *ListAppAssistantAgentsResponseBody) SetDynamicMessage(v string) *ListAppAssistantAgentsResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ListAppAssistantAgentsResponseBody) SetErrorArgs(v []interface{}) *ListAppAssistantAgentsResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *ListAppAssistantAgentsResponseBody) SetModule(v []*ListAppAssistantAgentsResponseBodyModule) *ListAppAssistantAgentsResponseBody {
	s.Module = v
	return s
}

func (s *ListAppAssistantAgentsResponseBody) SetRequestId(v string) *ListAppAssistantAgentsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAppAssistantAgentsResponseBody) SetRootErrorCode(v string) *ListAppAssistantAgentsResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *ListAppAssistantAgentsResponseBody) SetRootErrorMsg(v string) *ListAppAssistantAgentsResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *ListAppAssistantAgentsResponseBody) SetSynchro(v bool) *ListAppAssistantAgentsResponseBody {
	s.Synchro = &v
	return s
}

func (s *ListAppAssistantAgentsResponseBody) Validate() error {
	if s.Module != nil {
		for _, item := range s.Module {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAppAssistantAgentsResponseBodyModule struct {
	// example:
	//
	// liyang1_v@soulapp
	AgentId   *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	AgentName *string `json:"AgentName,omitempty" xml:"AgentName,omitempty"`
	// example:
	//
	// WD20250703155602000001
	BizId       *string                                              `json:"BizId,omitempty" xml:"BizId,omitempty"`
	Credential  *ListAppAssistantAgentsResponseBodyModuleCredential  `json:"Credential,omitempty" xml:"Credential,omitempty" type:"Struct"`
	EmbedConfig *ListAppAssistantAgentsResponseBodyModuleEmbedConfig `json:"EmbedConfig,omitempty" xml:"EmbedConfig,omitempty" type:"Struct"`
	ExtraParams map[string]*string                                   `json:"ExtraParams,omitempty" xml:"ExtraParams,omitempty"`
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
	// WA12313123131
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
	// 1231311312
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListAppAssistantAgentsResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s ListAppAssistantAgentsResponseBodyModule) GoString() string {
	return s.String()
}

func (s *ListAppAssistantAgentsResponseBodyModule) GetAgentId() *string {
	return s.AgentId
}

func (s *ListAppAssistantAgentsResponseBodyModule) GetAgentName() *string {
	return s.AgentName
}

func (s *ListAppAssistantAgentsResponseBodyModule) GetBizId() *string {
	return s.BizId
}

func (s *ListAppAssistantAgentsResponseBodyModule) GetCredential() *ListAppAssistantAgentsResponseBodyModuleCredential {
	return s.Credential
}

func (s *ListAppAssistantAgentsResponseBodyModule) GetEmbedConfig() *ListAppAssistantAgentsResponseBodyModuleEmbedConfig {
	return s.EmbedConfig
}

func (s *ListAppAssistantAgentsResponseBodyModule) GetExtraParams() map[string]*string {
	return s.ExtraParams
}

func (s *ListAppAssistantAgentsResponseBodyModule) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListAppAssistantAgentsResponseBodyModule) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListAppAssistantAgentsResponseBodyModule) GetPlatformAppId() *string {
	return s.PlatformAppId
}

func (s *ListAppAssistantAgentsResponseBodyModule) GetPlatformType() *string {
	return s.PlatformType
}

func (s *ListAppAssistantAgentsResponseBodyModule) GetStatus() *string {
	return s.Status
}

func (s *ListAppAssistantAgentsResponseBodyModule) GetUserId() *string {
	return s.UserId
}

func (s *ListAppAssistantAgentsResponseBodyModule) SetAgentId(v string) *ListAppAssistantAgentsResponseBodyModule {
	s.AgentId = &v
	return s
}

func (s *ListAppAssistantAgentsResponseBodyModule) SetAgentName(v string) *ListAppAssistantAgentsResponseBodyModule {
	s.AgentName = &v
	return s
}

func (s *ListAppAssistantAgentsResponseBodyModule) SetBizId(v string) *ListAppAssistantAgentsResponseBodyModule {
	s.BizId = &v
	return s
}

func (s *ListAppAssistantAgentsResponseBodyModule) SetCredential(v *ListAppAssistantAgentsResponseBodyModuleCredential) *ListAppAssistantAgentsResponseBodyModule {
	s.Credential = v
	return s
}

func (s *ListAppAssistantAgentsResponseBodyModule) SetEmbedConfig(v *ListAppAssistantAgentsResponseBodyModuleEmbedConfig) *ListAppAssistantAgentsResponseBodyModule {
	s.EmbedConfig = v
	return s
}

func (s *ListAppAssistantAgentsResponseBodyModule) SetExtraParams(v map[string]*string) *ListAppAssistantAgentsResponseBodyModule {
	s.ExtraParams = v
	return s
}

func (s *ListAppAssistantAgentsResponseBodyModule) SetGmtCreate(v string) *ListAppAssistantAgentsResponseBodyModule {
	s.GmtCreate = &v
	return s
}

func (s *ListAppAssistantAgentsResponseBodyModule) SetGmtModified(v string) *ListAppAssistantAgentsResponseBodyModule {
	s.GmtModified = &v
	return s
}

func (s *ListAppAssistantAgentsResponseBodyModule) SetPlatformAppId(v string) *ListAppAssistantAgentsResponseBodyModule {
	s.PlatformAppId = &v
	return s
}

func (s *ListAppAssistantAgentsResponseBodyModule) SetPlatformType(v string) *ListAppAssistantAgentsResponseBodyModule {
	s.PlatformType = &v
	return s
}

func (s *ListAppAssistantAgentsResponseBodyModule) SetStatus(v string) *ListAppAssistantAgentsResponseBodyModule {
	s.Status = &v
	return s
}

func (s *ListAppAssistantAgentsResponseBodyModule) SetUserId(v string) *ListAppAssistantAgentsResponseBodyModule {
	s.UserId = &v
	return s
}

func (s *ListAppAssistantAgentsResponseBodyModule) Validate() error {
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

type ListAppAssistantAgentsResponseBodyModuleCredential struct {
	// API Key
	//
	// example:
	//
	// akm-xxxxxxx
	ApiKey *string `json:"ApiKey,omitempty" xml:"ApiKey,omitempty"`
	// API Secret
	//
	// example:
	//
	// sk-xxxxxx
	ApiSecret *string            `json:"ApiSecret,omitempty" xml:"ApiSecret,omitempty"`
	Extra     map[string]*string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	// example:
	//
	// ***
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// example:
	//
	// Test
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s ListAppAssistantAgentsResponseBodyModuleCredential) String() string {
	return dara.Prettify(s)
}

func (s ListAppAssistantAgentsResponseBodyModuleCredential) GoString() string {
	return s.String()
}

func (s *ListAppAssistantAgentsResponseBodyModuleCredential) GetApiKey() *string {
	return s.ApiKey
}

func (s *ListAppAssistantAgentsResponseBodyModuleCredential) GetApiSecret() *string {
	return s.ApiSecret
}

func (s *ListAppAssistantAgentsResponseBodyModuleCredential) GetExtra() map[string]*string {
	return s.Extra
}

func (s *ListAppAssistantAgentsResponseBodyModuleCredential) GetPassword() *string {
	return s.Password
}

func (s *ListAppAssistantAgentsResponseBodyModuleCredential) GetUsername() *string {
	return s.Username
}

func (s *ListAppAssistantAgentsResponseBodyModuleCredential) SetApiKey(v string) *ListAppAssistantAgentsResponseBodyModuleCredential {
	s.ApiKey = &v
	return s
}

func (s *ListAppAssistantAgentsResponseBodyModuleCredential) SetApiSecret(v string) *ListAppAssistantAgentsResponseBodyModuleCredential {
	s.ApiSecret = &v
	return s
}

func (s *ListAppAssistantAgentsResponseBodyModuleCredential) SetExtra(v map[string]*string) *ListAppAssistantAgentsResponseBodyModuleCredential {
	s.Extra = v
	return s
}

func (s *ListAppAssistantAgentsResponseBodyModuleCredential) SetPassword(v string) *ListAppAssistantAgentsResponseBodyModuleCredential {
	s.Password = &v
	return s
}

func (s *ListAppAssistantAgentsResponseBodyModuleCredential) SetUsername(v string) *ListAppAssistantAgentsResponseBodyModuleCredential {
	s.Username = &v
	return s
}

func (s *ListAppAssistantAgentsResponseBodyModuleCredential) Validate() error {
	return dara.Validate(s)
}

type ListAppAssistantAgentsResponseBodyModuleEmbedConfig struct {
	Extra     map[string]*string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	RawScript *string            `json:"RawScript,omitempty" xml:"RawScript,omitempty"`
}

func (s ListAppAssistantAgentsResponseBodyModuleEmbedConfig) String() string {
	return dara.Prettify(s)
}

func (s ListAppAssistantAgentsResponseBodyModuleEmbedConfig) GoString() string {
	return s.String()
}

func (s *ListAppAssistantAgentsResponseBodyModuleEmbedConfig) GetExtra() map[string]*string {
	return s.Extra
}

func (s *ListAppAssistantAgentsResponseBodyModuleEmbedConfig) GetRawScript() *string {
	return s.RawScript
}

func (s *ListAppAssistantAgentsResponseBodyModuleEmbedConfig) SetExtra(v map[string]*string) *ListAppAssistantAgentsResponseBodyModuleEmbedConfig {
	s.Extra = v
	return s
}

func (s *ListAppAssistantAgentsResponseBodyModuleEmbedConfig) SetRawScript(v string) *ListAppAssistantAgentsResponseBodyModuleEmbedConfig {
	s.RawScript = &v
	return s
}

func (s *ListAppAssistantAgentsResponseBodyModuleEmbedConfig) Validate() error {
	return dara.Validate(s)
}
