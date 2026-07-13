// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWorkerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentType(v string) *CreateWorkerRequest
	GetAgentType() *string
	SetAgents(v string) *CreateWorkerRequest
	GetAgents() *string
	SetChannels(v []*CreateWorkerRequestChannels) *CreateWorkerRequest
	GetChannels() []*CreateWorkerRequestChannels
	SetClientToken(v string) *CreateWorkerRequest
	GetClientToken() *string
	SetCredentials(v []*CreateWorkerRequestCredentials) *CreateWorkerRequest
	GetCredentials() []*CreateWorkerRequestCredentials
	SetDeployType(v string) *CreateWorkerRequest
	GetDeployType() *string
	SetGroups(v []*CreateWorkerRequestGroups) *CreateWorkerRequest
	GetGroups() []*CreateWorkerRequestGroups
	SetInstanceId(v string) *CreateWorkerRequest
	GetInstanceId() *string
	SetLimitConfig(v *CreateWorkerRequestLimitConfig) *CreateWorkerRequest
	GetLimitConfig() *CreateWorkerRequestLimitConfig
	SetMcpServers(v []*CreateWorkerRequestMcpServers) *CreateWorkerRequest
	GetMcpServers() []*CreateWorkerRequestMcpServers
	SetModel(v *CreateWorkerRequestModel) *CreateWorkerRequest
	GetModel() *CreateWorkerRequestModel
	SetName(v string) *CreateWorkerRequest
	GetName() *string
	SetSkills(v []*CreateWorkerRequestSkills) *CreateWorkerRequest
	GetSkills() []*CreateWorkerRequestSkills
	SetSoul(v string) *CreateWorkerRequest
	GetSoul() *string
	SetSubagents(v []*CreateWorkerRequestSubagents) *CreateWorkerRequest
	GetSubagents() []*CreateWorkerRequestSubagents
	SetTemplate(v *CreateWorkerRequestTemplate) *CreateWorkerRequest
	GetTemplate() *CreateWorkerRequestTemplate
	SetVersionCode(v string) *CreateWorkerRequest
	GetVersionCode() *string
}

type CreateWorkerRequest struct {
	AgentType   *string                           `json:"AgentType,omitempty" xml:"AgentType,omitempty"`
	Agents      *string                           `json:"Agents,omitempty" xml:"Agents,omitempty"`
	Channels    []*CreateWorkerRequestChannels    `json:"Channels,omitempty" xml:"Channels,omitempty" type:"Repeated"`
	ClientToken *string                           `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Credentials []*CreateWorkerRequestCredentials `json:"Credentials,omitempty" xml:"Credentials,omitempty" type:"Repeated"`
	DeployType  *string                           `json:"DeployType,omitempty" xml:"DeployType,omitempty"`
	Groups      []*CreateWorkerRequestGroups      `json:"Groups,omitempty" xml:"Groups,omitempty" type:"Repeated"`
	InstanceId  *string                           `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	LimitConfig *CreateWorkerRequestLimitConfig   `json:"LimitConfig,omitempty" xml:"LimitConfig,omitempty" type:"Struct"`
	McpServers  []*CreateWorkerRequestMcpServers  `json:"McpServers,omitempty" xml:"McpServers,omitempty" type:"Repeated"`
	Model       *CreateWorkerRequestModel         `json:"Model,omitempty" xml:"Model,omitempty" type:"Struct"`
	Name        *string                           `json:"Name,omitempty" xml:"Name,omitempty"`
	Skills      []*CreateWorkerRequestSkills      `json:"Skills,omitempty" xml:"Skills,omitempty" type:"Repeated"`
	Soul        *string                           `json:"Soul,omitempty" xml:"Soul,omitempty"`
	Subagents   []*CreateWorkerRequestSubagents   `json:"Subagents,omitempty" xml:"Subagents,omitempty" type:"Repeated"`
	Template    *CreateWorkerRequestTemplate      `json:"Template,omitempty" xml:"Template,omitempty" type:"Struct"`
	VersionCode *string                           `json:"VersionCode,omitempty" xml:"VersionCode,omitempty"`
}

func (s CreateWorkerRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkerRequest) GoString() string {
	return s.String()
}

func (s *CreateWorkerRequest) GetAgentType() *string {
	return s.AgentType
}

func (s *CreateWorkerRequest) GetAgents() *string {
	return s.Agents
}

func (s *CreateWorkerRequest) GetChannels() []*CreateWorkerRequestChannels {
	return s.Channels
}

func (s *CreateWorkerRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateWorkerRequest) GetCredentials() []*CreateWorkerRequestCredentials {
	return s.Credentials
}

func (s *CreateWorkerRequest) GetDeployType() *string {
	return s.DeployType
}

func (s *CreateWorkerRequest) GetGroups() []*CreateWorkerRequestGroups {
	return s.Groups
}

func (s *CreateWorkerRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateWorkerRequest) GetLimitConfig() *CreateWorkerRequestLimitConfig {
	return s.LimitConfig
}

func (s *CreateWorkerRequest) GetMcpServers() []*CreateWorkerRequestMcpServers {
	return s.McpServers
}

func (s *CreateWorkerRequest) GetModel() *CreateWorkerRequestModel {
	return s.Model
}

func (s *CreateWorkerRequest) GetName() *string {
	return s.Name
}

func (s *CreateWorkerRequest) GetSkills() []*CreateWorkerRequestSkills {
	return s.Skills
}

func (s *CreateWorkerRequest) GetSoul() *string {
	return s.Soul
}

func (s *CreateWorkerRequest) GetSubagents() []*CreateWorkerRequestSubagents {
	return s.Subagents
}

func (s *CreateWorkerRequest) GetTemplate() *CreateWorkerRequestTemplate {
	return s.Template
}

func (s *CreateWorkerRequest) GetVersionCode() *string {
	return s.VersionCode
}

func (s *CreateWorkerRequest) SetAgentType(v string) *CreateWorkerRequest {
	s.AgentType = &v
	return s
}

func (s *CreateWorkerRequest) SetAgents(v string) *CreateWorkerRequest {
	s.Agents = &v
	return s
}

func (s *CreateWorkerRequest) SetChannels(v []*CreateWorkerRequestChannels) *CreateWorkerRequest {
	s.Channels = v
	return s
}

func (s *CreateWorkerRequest) SetClientToken(v string) *CreateWorkerRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateWorkerRequest) SetCredentials(v []*CreateWorkerRequestCredentials) *CreateWorkerRequest {
	s.Credentials = v
	return s
}

func (s *CreateWorkerRequest) SetDeployType(v string) *CreateWorkerRequest {
	s.DeployType = &v
	return s
}

func (s *CreateWorkerRequest) SetGroups(v []*CreateWorkerRequestGroups) *CreateWorkerRequest {
	s.Groups = v
	return s
}

func (s *CreateWorkerRequest) SetInstanceId(v string) *CreateWorkerRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateWorkerRequest) SetLimitConfig(v *CreateWorkerRequestLimitConfig) *CreateWorkerRequest {
	s.LimitConfig = v
	return s
}

func (s *CreateWorkerRequest) SetMcpServers(v []*CreateWorkerRequestMcpServers) *CreateWorkerRequest {
	s.McpServers = v
	return s
}

func (s *CreateWorkerRequest) SetModel(v *CreateWorkerRequestModel) *CreateWorkerRequest {
	s.Model = v
	return s
}

func (s *CreateWorkerRequest) SetName(v string) *CreateWorkerRequest {
	s.Name = &v
	return s
}

func (s *CreateWorkerRequest) SetSkills(v []*CreateWorkerRequestSkills) *CreateWorkerRequest {
	s.Skills = v
	return s
}

func (s *CreateWorkerRequest) SetSoul(v string) *CreateWorkerRequest {
	s.Soul = &v
	return s
}

func (s *CreateWorkerRequest) SetSubagents(v []*CreateWorkerRequestSubagents) *CreateWorkerRequest {
	s.Subagents = v
	return s
}

func (s *CreateWorkerRequest) SetTemplate(v *CreateWorkerRequestTemplate) *CreateWorkerRequest {
	s.Template = v
	return s
}

func (s *CreateWorkerRequest) SetVersionCode(v string) *CreateWorkerRequest {
	s.VersionCode = &v
	return s
}

func (s *CreateWorkerRequest) Validate() error {
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
	if s.Groups != nil {
		for _, item := range s.Groups {
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
	if s.Subagents != nil {
		for _, item := range s.Subagents {
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

type CreateWorkerRequestChannels struct {
	Config  *CreateWorkerRequestChannelsConfig  `json:"Config,omitempty" xml:"Config,omitempty" type:"Struct"`
	Enabled *bool                               `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	Secrets *CreateWorkerRequestChannelsSecrets `json:"Secrets,omitempty" xml:"Secrets,omitempty" type:"Struct"`
	Type    *string                             `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateWorkerRequestChannels) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkerRequestChannels) GoString() string {
	return s.String()
}

func (s *CreateWorkerRequestChannels) GetConfig() *CreateWorkerRequestChannelsConfig {
	return s.Config
}

func (s *CreateWorkerRequestChannels) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreateWorkerRequestChannels) GetSecrets() *CreateWorkerRequestChannelsSecrets {
	return s.Secrets
}

func (s *CreateWorkerRequestChannels) GetType() *string {
	return s.Type
}

func (s *CreateWorkerRequestChannels) SetConfig(v *CreateWorkerRequestChannelsConfig) *CreateWorkerRequestChannels {
	s.Config = v
	return s
}

func (s *CreateWorkerRequestChannels) SetEnabled(v bool) *CreateWorkerRequestChannels {
	s.Enabled = &v
	return s
}

func (s *CreateWorkerRequestChannels) SetSecrets(v *CreateWorkerRequestChannelsSecrets) *CreateWorkerRequestChannels {
	s.Secrets = v
	return s
}

func (s *CreateWorkerRequestChannels) SetType(v string) *CreateWorkerRequestChannels {
	s.Type = &v
	return s
}

func (s *CreateWorkerRequestChannels) Validate() error {
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

type CreateWorkerRequestChannelsConfig struct {
	CardTemplateId   *string `json:"CardTemplateId,omitempty" xml:"CardTemplateId,omitempty"`
	ClientId         *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	Extension        *string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	MessageType      *string `json:"MessageType,omitempty" xml:"MessageType,omitempty"`
	RobotCode        *string `json:"RobotCode,omitempty" xml:"RobotCode,omitempty"`
	ShowThinking     *bool   `json:"ShowThinking,omitempty" xml:"ShowThinking,omitempty"`
	ShowToolCalls    *bool   `json:"ShowToolCalls,omitempty" xml:"ShowToolCalls,omitempty"`
	StreamingEnabled *bool   `json:"StreamingEnabled,omitempty" xml:"StreamingEnabled,omitempty"`
}

func (s CreateWorkerRequestChannelsConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkerRequestChannelsConfig) GoString() string {
	return s.String()
}

func (s *CreateWorkerRequestChannelsConfig) GetCardTemplateId() *string {
	return s.CardTemplateId
}

func (s *CreateWorkerRequestChannelsConfig) GetClientId() *string {
	return s.ClientId
}

func (s *CreateWorkerRequestChannelsConfig) GetExtension() *string {
	return s.Extension
}

func (s *CreateWorkerRequestChannelsConfig) GetMessageType() *string {
	return s.MessageType
}

func (s *CreateWorkerRequestChannelsConfig) GetRobotCode() *string {
	return s.RobotCode
}

func (s *CreateWorkerRequestChannelsConfig) GetShowThinking() *bool {
	return s.ShowThinking
}

func (s *CreateWorkerRequestChannelsConfig) GetShowToolCalls() *bool {
	return s.ShowToolCalls
}

func (s *CreateWorkerRequestChannelsConfig) GetStreamingEnabled() *bool {
	return s.StreamingEnabled
}

func (s *CreateWorkerRequestChannelsConfig) SetCardTemplateId(v string) *CreateWorkerRequestChannelsConfig {
	s.CardTemplateId = &v
	return s
}

func (s *CreateWorkerRequestChannelsConfig) SetClientId(v string) *CreateWorkerRequestChannelsConfig {
	s.ClientId = &v
	return s
}

func (s *CreateWorkerRequestChannelsConfig) SetExtension(v string) *CreateWorkerRequestChannelsConfig {
	s.Extension = &v
	return s
}

func (s *CreateWorkerRequestChannelsConfig) SetMessageType(v string) *CreateWorkerRequestChannelsConfig {
	s.MessageType = &v
	return s
}

func (s *CreateWorkerRequestChannelsConfig) SetRobotCode(v string) *CreateWorkerRequestChannelsConfig {
	s.RobotCode = &v
	return s
}

func (s *CreateWorkerRequestChannelsConfig) SetShowThinking(v bool) *CreateWorkerRequestChannelsConfig {
	s.ShowThinking = &v
	return s
}

func (s *CreateWorkerRequestChannelsConfig) SetShowToolCalls(v bool) *CreateWorkerRequestChannelsConfig {
	s.ShowToolCalls = &v
	return s
}

func (s *CreateWorkerRequestChannelsConfig) SetStreamingEnabled(v bool) *CreateWorkerRequestChannelsConfig {
	s.StreamingEnabled = &v
	return s
}

func (s *CreateWorkerRequestChannelsConfig) Validate() error {
	return dara.Validate(s)
}

type CreateWorkerRequestChannelsSecrets struct {
	ClientSecret *string `json:"ClientSecret,omitempty" xml:"ClientSecret,omitempty"`
}

func (s CreateWorkerRequestChannelsSecrets) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkerRequestChannelsSecrets) GoString() string {
	return s.String()
}

func (s *CreateWorkerRequestChannelsSecrets) GetClientSecret() *string {
	return s.ClientSecret
}

func (s *CreateWorkerRequestChannelsSecrets) SetClientSecret(v string) *CreateWorkerRequestChannelsSecrets {
	s.ClientSecret = &v
	return s
}

func (s *CreateWorkerRequestChannelsSecrets) Validate() error {
	return dara.Validate(s)
}

type CreateWorkerRequestCredentials struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateWorkerRequestCredentials) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkerRequestCredentials) GoString() string {
	return s.String()
}

func (s *CreateWorkerRequestCredentials) GetName() *string {
	return s.Name
}

func (s *CreateWorkerRequestCredentials) SetName(v string) *CreateWorkerRequestCredentials {
	s.Name = &v
	return s
}

func (s *CreateWorkerRequestCredentials) Validate() error {
	return dara.Validate(s)
}

type CreateWorkerRequestGroups struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateWorkerRequestGroups) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkerRequestGroups) GoString() string {
	return s.String()
}

func (s *CreateWorkerRequestGroups) GetName() *string {
	return s.Name
}

func (s *CreateWorkerRequestGroups) GetRole() *string {
	return s.Role
}

func (s *CreateWorkerRequestGroups) GetType() *string {
	return s.Type
}

func (s *CreateWorkerRequestGroups) SetName(v string) *CreateWorkerRequestGroups {
	s.Name = &v
	return s
}

func (s *CreateWorkerRequestGroups) SetRole(v string) *CreateWorkerRequestGroups {
	s.Role = &v
	return s
}

func (s *CreateWorkerRequestGroups) SetType(v string) *CreateWorkerRequestGroups {
	s.Type = &v
	return s
}

func (s *CreateWorkerRequestGroups) Validate() error {
	return dara.Validate(s)
}

type CreateWorkerRequestLimitConfig struct {
	// example:
	//
	// token
	LimitType *string `json:"LimitType,omitempty" xml:"LimitType,omitempty"`
	// example:
	//
	// day
	PeriodType *string `json:"PeriodType,omitempty" xml:"PeriodType,omitempty"`
	UsageLimit *int64  `json:"UsageLimit,omitempty" xml:"UsageLimit,omitempty"`
}

func (s CreateWorkerRequestLimitConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkerRequestLimitConfig) GoString() string {
	return s.String()
}

func (s *CreateWorkerRequestLimitConfig) GetLimitType() *string {
	return s.LimitType
}

func (s *CreateWorkerRequestLimitConfig) GetPeriodType() *string {
	return s.PeriodType
}

func (s *CreateWorkerRequestLimitConfig) GetUsageLimit() *int64 {
	return s.UsageLimit
}

func (s *CreateWorkerRequestLimitConfig) SetLimitType(v string) *CreateWorkerRequestLimitConfig {
	s.LimitType = &v
	return s
}

func (s *CreateWorkerRequestLimitConfig) SetPeriodType(v string) *CreateWorkerRequestLimitConfig {
	s.PeriodType = &v
	return s
}

func (s *CreateWorkerRequestLimitConfig) SetUsageLimit(v int64) *CreateWorkerRequestLimitConfig {
	s.UsageLimit = &v
	return s
}

func (s *CreateWorkerRequestLimitConfig) Validate() error {
	return dara.Validate(s)
}

type CreateWorkerRequestMcpServers struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateWorkerRequestMcpServers) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkerRequestMcpServers) GoString() string {
	return s.String()
}

func (s *CreateWorkerRequestMcpServers) GetName() *string {
	return s.Name
}

func (s *CreateWorkerRequestMcpServers) SetName(v string) *CreateWorkerRequestMcpServers {
	s.Name = &v
	return s
}

func (s *CreateWorkerRequestMcpServers) Validate() error {
	return dara.Validate(s)
}

type CreateWorkerRequestModel struct {
	ModelName     *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	ModelProvider *string `json:"ModelProvider,omitempty" xml:"ModelProvider,omitempty"`
}

func (s CreateWorkerRequestModel) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkerRequestModel) GoString() string {
	return s.String()
}

func (s *CreateWorkerRequestModel) GetModelName() *string {
	return s.ModelName
}

func (s *CreateWorkerRequestModel) GetModelProvider() *string {
	return s.ModelProvider
}

func (s *CreateWorkerRequestModel) SetModelName(v string) *CreateWorkerRequestModel {
	s.ModelName = &v
	return s
}

func (s *CreateWorkerRequestModel) SetModelProvider(v string) *CreateWorkerRequestModel {
	s.ModelProvider = &v
	return s
}

func (s *CreateWorkerRequestModel) Validate() error {
	return dara.Validate(s)
}

type CreateWorkerRequestSkills struct {
	Label   *string `json:"Label,omitempty" xml:"Label,omitempty"`
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s CreateWorkerRequestSkills) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkerRequestSkills) GoString() string {
	return s.String()
}

func (s *CreateWorkerRequestSkills) GetLabel() *string {
	return s.Label
}

func (s *CreateWorkerRequestSkills) GetName() *string {
	return s.Name
}

func (s *CreateWorkerRequestSkills) GetVersion() *string {
	return s.Version
}

func (s *CreateWorkerRequestSkills) SetLabel(v string) *CreateWorkerRequestSkills {
	s.Label = &v
	return s
}

func (s *CreateWorkerRequestSkills) SetName(v string) *CreateWorkerRequestSkills {
	s.Name = &v
	return s
}

func (s *CreateWorkerRequestSkills) SetVersion(v string) *CreateWorkerRequestSkills {
	s.Version = &v
	return s
}

func (s *CreateWorkerRequestSkills) Validate() error {
	return dara.Validate(s)
}

type CreateWorkerRequestSubagents struct {
	Agents     *string                               `json:"Agents,omitempty" xml:"Agents,omitempty"`
	Skills     []*CreateWorkerRequestSubagentsSkills `json:"Skills,omitempty" xml:"Skills,omitempty" type:"Repeated"`
	SubagentId *string                               `json:"SubagentId,omitempty" xml:"SubagentId,omitempty"`
}

func (s CreateWorkerRequestSubagents) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkerRequestSubagents) GoString() string {
	return s.String()
}

func (s *CreateWorkerRequestSubagents) GetAgents() *string {
	return s.Agents
}

func (s *CreateWorkerRequestSubagents) GetSkills() []*CreateWorkerRequestSubagentsSkills {
	return s.Skills
}

func (s *CreateWorkerRequestSubagents) GetSubagentId() *string {
	return s.SubagentId
}

func (s *CreateWorkerRequestSubagents) SetAgents(v string) *CreateWorkerRequestSubagents {
	s.Agents = &v
	return s
}

func (s *CreateWorkerRequestSubagents) SetSkills(v []*CreateWorkerRequestSubagentsSkills) *CreateWorkerRequestSubagents {
	s.Skills = v
	return s
}

func (s *CreateWorkerRequestSubagents) SetSubagentId(v string) *CreateWorkerRequestSubagents {
	s.SubagentId = &v
	return s
}

func (s *CreateWorkerRequestSubagents) Validate() error {
	if s.Skills != nil {
		for _, item := range s.Skills {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateWorkerRequestSubagentsSkills struct {
	Label   *string `json:"Label,omitempty" xml:"Label,omitempty"`
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s CreateWorkerRequestSubagentsSkills) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkerRequestSubagentsSkills) GoString() string {
	return s.String()
}

func (s *CreateWorkerRequestSubagentsSkills) GetLabel() *string {
	return s.Label
}

func (s *CreateWorkerRequestSubagentsSkills) GetName() *string {
	return s.Name
}

func (s *CreateWorkerRequestSubagentsSkills) GetVersion() *string {
	return s.Version
}

func (s *CreateWorkerRequestSubagentsSkills) SetLabel(v string) *CreateWorkerRequestSubagentsSkills {
	s.Label = &v
	return s
}

func (s *CreateWorkerRequestSubagentsSkills) SetName(v string) *CreateWorkerRequestSubagentsSkills {
	s.Name = &v
	return s
}

func (s *CreateWorkerRequestSubagentsSkills) SetVersion(v string) *CreateWorkerRequestSubagentsSkills {
	s.Version = &v
	return s
}

func (s *CreateWorkerRequestSubagentsSkills) Validate() error {
	return dara.Validate(s)
}

type CreateWorkerRequestTemplate struct {
	Label   *string `json:"Label,omitempty" xml:"Label,omitempty"`
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s CreateWorkerRequestTemplate) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkerRequestTemplate) GoString() string {
	return s.String()
}

func (s *CreateWorkerRequestTemplate) GetLabel() *string {
	return s.Label
}

func (s *CreateWorkerRequestTemplate) GetName() *string {
	return s.Name
}

func (s *CreateWorkerRequestTemplate) GetVersion() *string {
	return s.Version
}

func (s *CreateWorkerRequestTemplate) SetLabel(v string) *CreateWorkerRequestTemplate {
	s.Label = &v
	return s
}

func (s *CreateWorkerRequestTemplate) SetName(v string) *CreateWorkerRequestTemplate {
	s.Name = &v
	return s
}

func (s *CreateWorkerRequestTemplate) SetVersion(v string) *CreateWorkerRequestTemplate {
	s.Version = &v
	return s
}

func (s *CreateWorkerRequestTemplate) Validate() error {
	return dara.Validate(s)
}
