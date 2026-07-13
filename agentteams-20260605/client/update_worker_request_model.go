// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWorkerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgents(v string) *UpdateWorkerRequest
	GetAgents() *string
	SetChannels(v []*UpdateWorkerRequestChannels) *UpdateWorkerRequest
	GetChannels() []*UpdateWorkerRequestChannels
	SetClientToken(v string) *UpdateWorkerRequest
	GetClientToken() *string
	SetCredentials(v []*UpdateWorkerRequestCredentials) *UpdateWorkerRequest
	GetCredentials() []*UpdateWorkerRequestCredentials
	SetInstanceId(v string) *UpdateWorkerRequest
	GetInstanceId() *string
	SetLimitConfig(v *UpdateWorkerRequestLimitConfig) *UpdateWorkerRequest
	GetLimitConfig() *UpdateWorkerRequestLimitConfig
	SetMcpServers(v []*UpdateWorkerRequestMcpServers) *UpdateWorkerRequest
	GetMcpServers() []*UpdateWorkerRequestMcpServers
	SetModel(v *UpdateWorkerRequestModel) *UpdateWorkerRequest
	GetModel() *UpdateWorkerRequestModel
	SetName(v string) *UpdateWorkerRequest
	GetName() *string
	SetSkills(v []*UpdateWorkerRequestSkills) *UpdateWorkerRequest
	GetSkills() []*UpdateWorkerRequestSkills
	SetSoul(v string) *UpdateWorkerRequest
	GetSoul() *string
	SetTemplate(v *UpdateWorkerRequestTemplate) *UpdateWorkerRequest
	GetTemplate() *UpdateWorkerRequestTemplate
	SetVersionCode(v string) *UpdateWorkerRequest
	GetVersionCode() *string
}

type UpdateWorkerRequest struct {
	Agents      *string                           `json:"Agents,omitempty" xml:"Agents,omitempty"`
	Channels    []*UpdateWorkerRequestChannels    `json:"Channels,omitempty" xml:"Channels,omitempty" type:"Repeated"`
	ClientToken *string                           `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Credentials []*UpdateWorkerRequestCredentials `json:"Credentials,omitempty" xml:"Credentials,omitempty" type:"Repeated"`
	// This parameter is required.
	InstanceId  *string                          `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	LimitConfig *UpdateWorkerRequestLimitConfig  `json:"LimitConfig,omitempty" xml:"LimitConfig,omitempty" type:"Struct"`
	McpServers  []*UpdateWorkerRequestMcpServers `json:"McpServers,omitempty" xml:"McpServers,omitempty" type:"Repeated"`
	Model       *UpdateWorkerRequestModel        `json:"Model,omitempty" xml:"Model,omitempty" type:"Struct"`
	// This parameter is required.
	Name        *string                      `json:"Name,omitempty" xml:"Name,omitempty"`
	Skills      []*UpdateWorkerRequestSkills `json:"Skills,omitempty" xml:"Skills,omitempty" type:"Repeated"`
	Soul        *string                      `json:"Soul,omitempty" xml:"Soul,omitempty"`
	Template    *UpdateWorkerRequestTemplate `json:"Template,omitempty" xml:"Template,omitempty" type:"Struct"`
	VersionCode *string                      `json:"VersionCode,omitempty" xml:"VersionCode,omitempty"`
}

func (s UpdateWorkerRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkerRequest) GoString() string {
	return s.String()
}

func (s *UpdateWorkerRequest) GetAgents() *string {
	return s.Agents
}

func (s *UpdateWorkerRequest) GetChannels() []*UpdateWorkerRequestChannels {
	return s.Channels
}

func (s *UpdateWorkerRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateWorkerRequest) GetCredentials() []*UpdateWorkerRequestCredentials {
	return s.Credentials
}

func (s *UpdateWorkerRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateWorkerRequest) GetLimitConfig() *UpdateWorkerRequestLimitConfig {
	return s.LimitConfig
}

func (s *UpdateWorkerRequest) GetMcpServers() []*UpdateWorkerRequestMcpServers {
	return s.McpServers
}

func (s *UpdateWorkerRequest) GetModel() *UpdateWorkerRequestModel {
	return s.Model
}

func (s *UpdateWorkerRequest) GetName() *string {
	return s.Name
}

func (s *UpdateWorkerRequest) GetSkills() []*UpdateWorkerRequestSkills {
	return s.Skills
}

func (s *UpdateWorkerRequest) GetSoul() *string {
	return s.Soul
}

func (s *UpdateWorkerRequest) GetTemplate() *UpdateWorkerRequestTemplate {
	return s.Template
}

func (s *UpdateWorkerRequest) GetVersionCode() *string {
	return s.VersionCode
}

func (s *UpdateWorkerRequest) SetAgents(v string) *UpdateWorkerRequest {
	s.Agents = &v
	return s
}

func (s *UpdateWorkerRequest) SetChannels(v []*UpdateWorkerRequestChannels) *UpdateWorkerRequest {
	s.Channels = v
	return s
}

func (s *UpdateWorkerRequest) SetClientToken(v string) *UpdateWorkerRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateWorkerRequest) SetCredentials(v []*UpdateWorkerRequestCredentials) *UpdateWorkerRequest {
	s.Credentials = v
	return s
}

func (s *UpdateWorkerRequest) SetInstanceId(v string) *UpdateWorkerRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateWorkerRequest) SetLimitConfig(v *UpdateWorkerRequestLimitConfig) *UpdateWorkerRequest {
	s.LimitConfig = v
	return s
}

func (s *UpdateWorkerRequest) SetMcpServers(v []*UpdateWorkerRequestMcpServers) *UpdateWorkerRequest {
	s.McpServers = v
	return s
}

func (s *UpdateWorkerRequest) SetModel(v *UpdateWorkerRequestModel) *UpdateWorkerRequest {
	s.Model = v
	return s
}

func (s *UpdateWorkerRequest) SetName(v string) *UpdateWorkerRequest {
	s.Name = &v
	return s
}

func (s *UpdateWorkerRequest) SetSkills(v []*UpdateWorkerRequestSkills) *UpdateWorkerRequest {
	s.Skills = v
	return s
}

func (s *UpdateWorkerRequest) SetSoul(v string) *UpdateWorkerRequest {
	s.Soul = &v
	return s
}

func (s *UpdateWorkerRequest) SetTemplate(v *UpdateWorkerRequestTemplate) *UpdateWorkerRequest {
	s.Template = v
	return s
}

func (s *UpdateWorkerRequest) SetVersionCode(v string) *UpdateWorkerRequest {
	s.VersionCode = &v
	return s
}

func (s *UpdateWorkerRequest) Validate() error {
	if s.Channels != nil {
		for _, item := range s.Channels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Credentials != nil {
		for _, item := range s.Credentials {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.LimitConfig != nil {
		if err := s.LimitConfig.Validate(); err != nil {
			return err
		}
	}
	if s.McpServers != nil {
		for _, item := range s.McpServers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Model != nil {
		if err := s.Model.Validate(); err != nil {
			return err
		}
	}
	if s.Skills != nil {
		for _, item := range s.Skills {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Template != nil {
		if err := s.Template.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateWorkerRequestChannels struct {
	Config  *UpdateWorkerRequestChannelsConfig  `json:"Config,omitempty" xml:"Config,omitempty" type:"Struct"`
	Enabled *bool                               `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	Secrets *UpdateWorkerRequestChannelsSecrets `json:"Secrets,omitempty" xml:"Secrets,omitempty" type:"Struct"`
	Type    *string                             `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateWorkerRequestChannels) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkerRequestChannels) GoString() string {
	return s.String()
}

func (s *UpdateWorkerRequestChannels) GetConfig() *UpdateWorkerRequestChannelsConfig {
	return s.Config
}

func (s *UpdateWorkerRequestChannels) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdateWorkerRequestChannels) GetSecrets() *UpdateWorkerRequestChannelsSecrets {
	return s.Secrets
}

func (s *UpdateWorkerRequestChannels) GetType() *string {
	return s.Type
}

func (s *UpdateWorkerRequestChannels) SetConfig(v *UpdateWorkerRequestChannelsConfig) *UpdateWorkerRequestChannels {
	s.Config = v
	return s
}

func (s *UpdateWorkerRequestChannels) SetEnabled(v bool) *UpdateWorkerRequestChannels {
	s.Enabled = &v
	return s
}

func (s *UpdateWorkerRequestChannels) SetSecrets(v *UpdateWorkerRequestChannelsSecrets) *UpdateWorkerRequestChannels {
	s.Secrets = v
	return s
}

func (s *UpdateWorkerRequestChannels) SetType(v string) *UpdateWorkerRequestChannels {
	s.Type = &v
	return s
}

func (s *UpdateWorkerRequestChannels) Validate() error {
	if s.Config != nil {
		if err := s.Config.Validate(); err != nil {
			return err
		}
	}
	if s.Secrets != nil {
		if err := s.Secrets.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateWorkerRequestChannelsConfig struct {
	CardTemplateId   *string `json:"CardTemplateId,omitempty" xml:"CardTemplateId,omitempty"`
	ClientId         *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	Extension        *string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	MessageType      *string `json:"MessageType,omitempty" xml:"MessageType,omitempty"`
	RobotCode        *string `json:"RobotCode,omitempty" xml:"RobotCode,omitempty"`
	ShowThinking     *bool   `json:"ShowThinking,omitempty" xml:"ShowThinking,omitempty"`
	ShowToolCalls    *bool   `json:"ShowToolCalls,omitempty" xml:"ShowToolCalls,omitempty"`
	StreamingEnabled *bool   `json:"StreamingEnabled,omitempty" xml:"StreamingEnabled,omitempty"`
}

func (s UpdateWorkerRequestChannelsConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkerRequestChannelsConfig) GoString() string {
	return s.String()
}

func (s *UpdateWorkerRequestChannelsConfig) GetCardTemplateId() *string {
	return s.CardTemplateId
}

func (s *UpdateWorkerRequestChannelsConfig) GetClientId() *string {
	return s.ClientId
}

func (s *UpdateWorkerRequestChannelsConfig) GetExtension() *string {
	return s.Extension
}

func (s *UpdateWorkerRequestChannelsConfig) GetMessageType() *string {
	return s.MessageType
}

func (s *UpdateWorkerRequestChannelsConfig) GetRobotCode() *string {
	return s.RobotCode
}

func (s *UpdateWorkerRequestChannelsConfig) GetShowThinking() *bool {
	return s.ShowThinking
}

func (s *UpdateWorkerRequestChannelsConfig) GetShowToolCalls() *bool {
	return s.ShowToolCalls
}

func (s *UpdateWorkerRequestChannelsConfig) GetStreamingEnabled() *bool {
	return s.StreamingEnabled
}

func (s *UpdateWorkerRequestChannelsConfig) SetCardTemplateId(v string) *UpdateWorkerRequestChannelsConfig {
	s.CardTemplateId = &v
	return s
}

func (s *UpdateWorkerRequestChannelsConfig) SetClientId(v string) *UpdateWorkerRequestChannelsConfig {
	s.ClientId = &v
	return s
}

func (s *UpdateWorkerRequestChannelsConfig) SetExtension(v string) *UpdateWorkerRequestChannelsConfig {
	s.Extension = &v
	return s
}

func (s *UpdateWorkerRequestChannelsConfig) SetMessageType(v string) *UpdateWorkerRequestChannelsConfig {
	s.MessageType = &v
	return s
}

func (s *UpdateWorkerRequestChannelsConfig) SetRobotCode(v string) *UpdateWorkerRequestChannelsConfig {
	s.RobotCode = &v
	return s
}

func (s *UpdateWorkerRequestChannelsConfig) SetShowThinking(v bool) *UpdateWorkerRequestChannelsConfig {
	s.ShowThinking = &v
	return s
}

func (s *UpdateWorkerRequestChannelsConfig) SetShowToolCalls(v bool) *UpdateWorkerRequestChannelsConfig {
	s.ShowToolCalls = &v
	return s
}

func (s *UpdateWorkerRequestChannelsConfig) SetStreamingEnabled(v bool) *UpdateWorkerRequestChannelsConfig {
	s.StreamingEnabled = &v
	return s
}

func (s *UpdateWorkerRequestChannelsConfig) Validate() error {
	return dara.Validate(s)
}

type UpdateWorkerRequestChannelsSecrets struct {
	ClientSecret *string `json:"ClientSecret,omitempty" xml:"ClientSecret,omitempty"`
}

func (s UpdateWorkerRequestChannelsSecrets) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkerRequestChannelsSecrets) GoString() string {
	return s.String()
}

func (s *UpdateWorkerRequestChannelsSecrets) GetClientSecret() *string {
	return s.ClientSecret
}

func (s *UpdateWorkerRequestChannelsSecrets) SetClientSecret(v string) *UpdateWorkerRequestChannelsSecrets {
	s.ClientSecret = &v
	return s
}

func (s *UpdateWorkerRequestChannelsSecrets) Validate() error {
	return dara.Validate(s)
}

type UpdateWorkerRequestCredentials struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateWorkerRequestCredentials) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkerRequestCredentials) GoString() string {
	return s.String()
}

func (s *UpdateWorkerRequestCredentials) GetName() *string {
	return s.Name
}

func (s *UpdateWorkerRequestCredentials) SetName(v string) *UpdateWorkerRequestCredentials {
	s.Name = &v
	return s
}

func (s *UpdateWorkerRequestCredentials) Validate() error {
	return dara.Validate(s)
}

type UpdateWorkerRequestLimitConfig struct {
	LimitType  *string `json:"LimitType,omitempty" xml:"LimitType,omitempty"`
	PeriodType *string `json:"PeriodType,omitempty" xml:"PeriodType,omitempty"`
	UsageLimit *int64  `json:"UsageLimit,omitempty" xml:"UsageLimit,omitempty"`
}

func (s UpdateWorkerRequestLimitConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkerRequestLimitConfig) GoString() string {
	return s.String()
}

func (s *UpdateWorkerRequestLimitConfig) GetLimitType() *string {
	return s.LimitType
}

func (s *UpdateWorkerRequestLimitConfig) GetPeriodType() *string {
	return s.PeriodType
}

func (s *UpdateWorkerRequestLimitConfig) GetUsageLimit() *int64 {
	return s.UsageLimit
}

func (s *UpdateWorkerRequestLimitConfig) SetLimitType(v string) *UpdateWorkerRequestLimitConfig {
	s.LimitType = &v
	return s
}

func (s *UpdateWorkerRequestLimitConfig) SetPeriodType(v string) *UpdateWorkerRequestLimitConfig {
	s.PeriodType = &v
	return s
}

func (s *UpdateWorkerRequestLimitConfig) SetUsageLimit(v int64) *UpdateWorkerRequestLimitConfig {
	s.UsageLimit = &v
	return s
}

func (s *UpdateWorkerRequestLimitConfig) Validate() error {
	return dara.Validate(s)
}

type UpdateWorkerRequestMcpServers struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateWorkerRequestMcpServers) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkerRequestMcpServers) GoString() string {
	return s.String()
}

func (s *UpdateWorkerRequestMcpServers) GetName() *string {
	return s.Name
}

func (s *UpdateWorkerRequestMcpServers) SetName(v string) *UpdateWorkerRequestMcpServers {
	s.Name = &v
	return s
}

func (s *UpdateWorkerRequestMcpServers) Validate() error {
	return dara.Validate(s)
}

type UpdateWorkerRequestModel struct {
	ModelName     *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	ModelProvider *string `json:"ModelProvider,omitempty" xml:"ModelProvider,omitempty"`
}

func (s UpdateWorkerRequestModel) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkerRequestModel) GoString() string {
	return s.String()
}

func (s *UpdateWorkerRequestModel) GetModelName() *string {
	return s.ModelName
}

func (s *UpdateWorkerRequestModel) GetModelProvider() *string {
	return s.ModelProvider
}

func (s *UpdateWorkerRequestModel) SetModelName(v string) *UpdateWorkerRequestModel {
	s.ModelName = &v
	return s
}

func (s *UpdateWorkerRequestModel) SetModelProvider(v string) *UpdateWorkerRequestModel {
	s.ModelProvider = &v
	return s
}

func (s *UpdateWorkerRequestModel) Validate() error {
	return dara.Validate(s)
}

type UpdateWorkerRequestSkills struct {
	Label   *string `json:"Label,omitempty" xml:"Label,omitempty"`
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s UpdateWorkerRequestSkills) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkerRequestSkills) GoString() string {
	return s.String()
}

func (s *UpdateWorkerRequestSkills) GetLabel() *string {
	return s.Label
}

func (s *UpdateWorkerRequestSkills) GetName() *string {
	return s.Name
}

func (s *UpdateWorkerRequestSkills) GetVersion() *string {
	return s.Version
}

func (s *UpdateWorkerRequestSkills) SetLabel(v string) *UpdateWorkerRequestSkills {
	s.Label = &v
	return s
}

func (s *UpdateWorkerRequestSkills) SetName(v string) *UpdateWorkerRequestSkills {
	s.Name = &v
	return s
}

func (s *UpdateWorkerRequestSkills) SetVersion(v string) *UpdateWorkerRequestSkills {
	s.Version = &v
	return s
}

func (s *UpdateWorkerRequestSkills) Validate() error {
	return dara.Validate(s)
}

type UpdateWorkerRequestTemplate struct {
	Label   *string `json:"Label,omitempty" xml:"Label,omitempty"`
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s UpdateWorkerRequestTemplate) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkerRequestTemplate) GoString() string {
	return s.String()
}

func (s *UpdateWorkerRequestTemplate) GetLabel() *string {
	return s.Label
}

func (s *UpdateWorkerRequestTemplate) GetName() *string {
	return s.Name
}

func (s *UpdateWorkerRequestTemplate) GetVersion() *string {
	return s.Version
}

func (s *UpdateWorkerRequestTemplate) SetLabel(v string) *UpdateWorkerRequestTemplate {
	s.Label = &v
	return s
}

func (s *UpdateWorkerRequestTemplate) SetName(v string) *UpdateWorkerRequestTemplate {
	s.Name = &v
	return s
}

func (s *UpdateWorkerRequestTemplate) SetVersion(v string) *UpdateWorkerRequestTemplate {
	s.Version = &v
	return s
}

func (s *UpdateWorkerRequestTemplate) Validate() error {
	return dara.Validate(s)
}
