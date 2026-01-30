// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAgentServiceConfig interface {
	dara.Model
	String() string
	GoString() string
	SetAddress(v string) *AgentServiceConfig
	GetAddress() *string
	SetCustomConfig(v *AgentServiceConfigCustomConfig) *AgentServiceConfig
	GetCustomConfig() *AgentServiceConfigCustomConfig
	SetDashScopeConfig(v *AgentServiceConfigDashScopeConfig) *AgentServiceConfig
	GetDashScopeConfig() *AgentServiceConfigDashScopeConfig
	SetDifyConfig(v *AgentServiceConfigDifyConfig) *AgentServiceConfig
	GetDifyConfig() *AgentServiceConfigDifyConfig
	SetEnableHealthCheck(v bool) *AgentServiceConfig
	GetEnableHealthCheck() *bool
	SetEnableOutlierDetection(v bool) *AgentServiceConfig
	GetEnableOutlierDetection() *bool
	SetProtocols(v []*string) *AgentServiceConfig
	GetProtocols() []*string
	SetProvider(v string) *AgentServiceConfig
	GetProvider() *string
}

type AgentServiceConfig struct {
	// This parameter is required.
	//
	// example:
	//
	// https://dashscope.aliyuncs.com/api/v1
	Address                *string                            `json:"address,omitempty" xml:"address,omitempty"`
	CustomConfig           *AgentServiceConfigCustomConfig    `json:"customConfig,omitempty" xml:"customConfig,omitempty" type:"Struct"`
	DashScopeConfig        *AgentServiceConfigDashScopeConfig `json:"dashScopeConfig,omitempty" xml:"dashScopeConfig,omitempty" type:"Struct"`
	DifyConfig             *AgentServiceConfigDifyConfig      `json:"difyConfig,omitempty" xml:"difyConfig,omitempty" type:"Struct"`
	EnableHealthCheck      *bool                              `json:"enableHealthCheck,omitempty" xml:"enableHealthCheck,omitempty"`
	EnableOutlierDetection *bool                              `json:"enableOutlierDetection,omitempty" xml:"enableOutlierDetection,omitempty"`
	Protocols              []*string                          `json:"protocols,omitempty" xml:"protocols,omitempty" type:"Repeated"`
	// This parameter is required.
	Provider *string `json:"provider,omitempty" xml:"provider,omitempty"`
}

func (s AgentServiceConfig) String() string {
	return dara.Prettify(s)
}

func (s AgentServiceConfig) GoString() string {
	return s.String()
}

func (s *AgentServiceConfig) GetAddress() *string {
	return s.Address
}

func (s *AgentServiceConfig) GetCustomConfig() *AgentServiceConfigCustomConfig {
	return s.CustomConfig
}

func (s *AgentServiceConfig) GetDashScopeConfig() *AgentServiceConfigDashScopeConfig {
	return s.DashScopeConfig
}

func (s *AgentServiceConfig) GetDifyConfig() *AgentServiceConfigDifyConfig {
	return s.DifyConfig
}

func (s *AgentServiceConfig) GetEnableHealthCheck() *bool {
	return s.EnableHealthCheck
}

func (s *AgentServiceConfig) GetEnableOutlierDetection() *bool {
	return s.EnableOutlierDetection
}

func (s *AgentServiceConfig) GetProtocols() []*string {
	return s.Protocols
}

func (s *AgentServiceConfig) GetProvider() *string {
	return s.Provider
}

func (s *AgentServiceConfig) SetAddress(v string) *AgentServiceConfig {
	s.Address = &v
	return s
}

func (s *AgentServiceConfig) SetCustomConfig(v *AgentServiceConfigCustomConfig) *AgentServiceConfig {
	s.CustomConfig = v
	return s
}

func (s *AgentServiceConfig) SetDashScopeConfig(v *AgentServiceConfigDashScopeConfig) *AgentServiceConfig {
	s.DashScopeConfig = v
	return s
}

func (s *AgentServiceConfig) SetDifyConfig(v *AgentServiceConfigDifyConfig) *AgentServiceConfig {
	s.DifyConfig = v
	return s
}

func (s *AgentServiceConfig) SetEnableHealthCheck(v bool) *AgentServiceConfig {
	s.EnableHealthCheck = &v
	return s
}

func (s *AgentServiceConfig) SetEnableOutlierDetection(v bool) *AgentServiceConfig {
	s.EnableOutlierDetection = &v
	return s
}

func (s *AgentServiceConfig) SetProtocols(v []*string) *AgentServiceConfig {
	s.Protocols = v
	return s
}

func (s *AgentServiceConfig) SetProvider(v string) *AgentServiceConfig {
	s.Provider = &v
	return s
}

func (s *AgentServiceConfig) Validate() error {
	if s.CustomConfig != nil {
		if err := s.CustomConfig.Validate(); err != nil {
			return err
		}
	}
	if s.DashScopeConfig != nil {
		if err := s.DashScopeConfig.Validate(); err != nil {
			return err
		}
	}
	if s.DifyConfig != nil {
		if err := s.DifyConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AgentServiceConfigCustomConfig struct {
	ApiKey             *string `json:"apiKey,omitempty" xml:"apiKey,omitempty"`
	ApiKeyGenerateMode *string `json:"apiKeyGenerateMode,omitempty" xml:"apiKeyGenerateMode,omitempty"`
}

func (s AgentServiceConfigCustomConfig) String() string {
	return dara.Prettify(s)
}

func (s AgentServiceConfigCustomConfig) GoString() string {
	return s.String()
}

func (s *AgentServiceConfigCustomConfig) GetApiKey() *string {
	return s.ApiKey
}

func (s *AgentServiceConfigCustomConfig) GetApiKeyGenerateMode() *string {
	return s.ApiKeyGenerateMode
}

func (s *AgentServiceConfigCustomConfig) SetApiKey(v string) *AgentServiceConfigCustomConfig {
	s.ApiKey = &v
	return s
}

func (s *AgentServiceConfigCustomConfig) SetApiKeyGenerateMode(v string) *AgentServiceConfigCustomConfig {
	s.ApiKeyGenerateMode = &v
	return s
}

func (s *AgentServiceConfigCustomConfig) Validate() error {
	return dara.Validate(s)
}

type AgentServiceConfigDashScopeConfig struct {
	AppCredentials []*AgentServiceConfigDashScopeConfigAppCredentials `json:"appCredentials,omitempty" xml:"appCredentials,omitempty" type:"Repeated"`
}

func (s AgentServiceConfigDashScopeConfig) String() string {
	return dara.Prettify(s)
}

func (s AgentServiceConfigDashScopeConfig) GoString() string {
	return s.String()
}

func (s *AgentServiceConfigDashScopeConfig) GetAppCredentials() []*AgentServiceConfigDashScopeConfigAppCredentials {
	return s.AppCredentials
}

func (s *AgentServiceConfigDashScopeConfig) SetAppCredentials(v []*AgentServiceConfigDashScopeConfigAppCredentials) *AgentServiceConfigDashScopeConfig {
	s.AppCredentials = v
	return s
}

func (s *AgentServiceConfigDashScopeConfig) Validate() error {
	if s.AppCredentials != nil {
		for _, item := range s.AppCredentials {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AgentServiceConfigDashScopeConfigAppCredentials struct {
	ApiKey *string `json:"apiKey,omitempty" xml:"apiKey,omitempty"`
	AppId  *string `json:"appId,omitempty" xml:"appId,omitempty"`
}

func (s AgentServiceConfigDashScopeConfigAppCredentials) String() string {
	return dara.Prettify(s)
}

func (s AgentServiceConfigDashScopeConfigAppCredentials) GoString() string {
	return s.String()
}

func (s *AgentServiceConfigDashScopeConfigAppCredentials) GetApiKey() *string {
	return s.ApiKey
}

func (s *AgentServiceConfigDashScopeConfigAppCredentials) GetAppId() *string {
	return s.AppId
}

func (s *AgentServiceConfigDashScopeConfigAppCredentials) SetApiKey(v string) *AgentServiceConfigDashScopeConfigAppCredentials {
	s.ApiKey = &v
	return s
}

func (s *AgentServiceConfigDashScopeConfigAppCredentials) SetAppId(v string) *AgentServiceConfigDashScopeConfigAppCredentials {
	s.AppId = &v
	return s
}

func (s *AgentServiceConfigDashScopeConfigAppCredentials) Validate() error {
	return dara.Validate(s)
}

type AgentServiceConfigDifyConfig struct {
	ApiKey  *string `json:"apiKey,omitempty" xml:"apiKey,omitempty"`
	BotType *string `json:"botType,omitempty" xml:"botType,omitempty"`
}

func (s AgentServiceConfigDifyConfig) String() string {
	return dara.Prettify(s)
}

func (s AgentServiceConfigDifyConfig) GoString() string {
	return s.String()
}

func (s *AgentServiceConfigDifyConfig) GetApiKey() *string {
	return s.ApiKey
}

func (s *AgentServiceConfigDifyConfig) GetBotType() *string {
	return s.BotType
}

func (s *AgentServiceConfigDifyConfig) SetApiKey(v string) *AgentServiceConfigDifyConfig {
	s.ApiKey = &v
	return s
}

func (s *AgentServiceConfigDifyConfig) SetBotType(v string) *AgentServiceConfigDifyConfig {
	s.BotType = &v
	return s
}

func (s *AgentServiceConfigDifyConfig) Validate() error {
	return dara.Validate(s)
}
