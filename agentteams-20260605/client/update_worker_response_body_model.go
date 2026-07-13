// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWorkerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateWorkerResponseBody
	GetCode() *string
	SetData(v *UpdateWorkerResponseBodyData) *UpdateWorkerResponseBody
	GetData() *UpdateWorkerResponseBodyData
	SetHttpStatusCode(v int32) *UpdateWorkerResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateWorkerResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateWorkerResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateWorkerResponseBody
	GetSuccess() *bool
}

type UpdateWorkerResponseBody struct {
	Code           *string                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *UpdateWorkerResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                        `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                         `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateWorkerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkerResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateWorkerResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateWorkerResponseBody) GetData() *UpdateWorkerResponseBodyData {
	return s.Data
}

func (s *UpdateWorkerResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateWorkerResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateWorkerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateWorkerResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateWorkerResponseBody) SetCode(v string) *UpdateWorkerResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateWorkerResponseBody) SetData(v *UpdateWorkerResponseBodyData) *UpdateWorkerResponseBody {
	s.Data = v
	return s
}

func (s *UpdateWorkerResponseBody) SetHttpStatusCode(v int32) *UpdateWorkerResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateWorkerResponseBody) SetMessage(v string) *UpdateWorkerResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateWorkerResponseBody) SetRequestId(v string) *UpdateWorkerResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateWorkerResponseBody) SetSuccess(v bool) *UpdateWorkerResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateWorkerResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateWorkerResponseBodyData struct {
	AgentType   *string                                    `json:"AgentType,omitempty" xml:"AgentType,omitempty"`
	Agents      *string                                    `json:"Agents,omitempty" xml:"Agents,omitempty"`
	Channels    []*UpdateWorkerResponseBodyDataChannels    `json:"Channels,omitempty" xml:"Channels,omitempty" type:"Repeated"`
	Credentials []*UpdateWorkerResponseBodyDataCredentials `json:"Credentials,omitempty" xml:"Credentials,omitempty" type:"Repeated"`
	DeployType  *string                                    `json:"DeployType,omitempty" xml:"DeployType,omitempty"`
	Groups      []*UpdateWorkerResponseBodyDataGroups      `json:"Groups,omitempty" xml:"Groups,omitempty" type:"Repeated"`
	InstanceId  *string                                    `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	LimitConfig *UpdateWorkerResponseBodyDataLimitConfig   `json:"LimitConfig,omitempty" xml:"LimitConfig,omitempty" type:"Struct"`
	McpServers  []*UpdateWorkerResponseBodyDataMcpServers  `json:"McpServers,omitempty" xml:"McpServers,omitempty" type:"Repeated"`
	Model       *UpdateWorkerResponseBodyDataModel         `json:"Model,omitempty" xml:"Model,omitempty" type:"Struct"`
	Name        *string                                    `json:"Name,omitempty" xml:"Name,omitempty"`
	RegionId    *string                                    `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Skills      []*UpdateWorkerResponseBodyDataSkills      `json:"Skills,omitempty" xml:"Skills,omitempty" type:"Repeated"`
	Soul        *string                                    `json:"Soul,omitempty" xml:"Soul,omitempty"`
	StartTime   *string                                    `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status      *string                                    `json:"Status,omitempty" xml:"Status,omitempty"`
	Template    *UpdateWorkerResponseBodyDataTemplate      `json:"Template,omitempty" xml:"Template,omitempty" type:"Struct"`
	VersionCode *string                                    `json:"VersionCode,omitempty" xml:"VersionCode,omitempty"`
}

func (s UpdateWorkerResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkerResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateWorkerResponseBodyData) GetAgentType() *string {
	return s.AgentType
}

func (s *UpdateWorkerResponseBodyData) GetAgents() *string {
	return s.Agents
}

func (s *UpdateWorkerResponseBodyData) GetChannels() []*UpdateWorkerResponseBodyDataChannels {
	return s.Channels
}

func (s *UpdateWorkerResponseBodyData) GetCredentials() []*UpdateWorkerResponseBodyDataCredentials {
	return s.Credentials
}

func (s *UpdateWorkerResponseBodyData) GetDeployType() *string {
	return s.DeployType
}

func (s *UpdateWorkerResponseBodyData) GetGroups() []*UpdateWorkerResponseBodyDataGroups {
	return s.Groups
}

func (s *UpdateWorkerResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateWorkerResponseBodyData) GetLimitConfig() *UpdateWorkerResponseBodyDataLimitConfig {
	return s.LimitConfig
}

func (s *UpdateWorkerResponseBodyData) GetMcpServers() []*UpdateWorkerResponseBodyDataMcpServers {
	return s.McpServers
}

func (s *UpdateWorkerResponseBodyData) GetModel() *UpdateWorkerResponseBodyDataModel {
	return s.Model
}

func (s *UpdateWorkerResponseBodyData) GetName() *string {
	return s.Name
}

func (s *UpdateWorkerResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateWorkerResponseBodyData) GetSkills() []*UpdateWorkerResponseBodyDataSkills {
	return s.Skills
}

func (s *UpdateWorkerResponseBodyData) GetSoul() *string {
	return s.Soul
}

func (s *UpdateWorkerResponseBodyData) GetStartTime() *string {
	return s.StartTime
}

func (s *UpdateWorkerResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *UpdateWorkerResponseBodyData) GetTemplate() *UpdateWorkerResponseBodyDataTemplate {
	return s.Template
}

func (s *UpdateWorkerResponseBodyData) GetVersionCode() *string {
	return s.VersionCode
}

func (s *UpdateWorkerResponseBodyData) SetAgentType(v string) *UpdateWorkerResponseBodyData {
	s.AgentType = &v
	return s
}

func (s *UpdateWorkerResponseBodyData) SetAgents(v string) *UpdateWorkerResponseBodyData {
	s.Agents = &v
	return s
}

func (s *UpdateWorkerResponseBodyData) SetChannels(v []*UpdateWorkerResponseBodyDataChannels) *UpdateWorkerResponseBodyData {
	s.Channels = v
	return s
}

func (s *UpdateWorkerResponseBodyData) SetCredentials(v []*UpdateWorkerResponseBodyDataCredentials) *UpdateWorkerResponseBodyData {
	s.Credentials = v
	return s
}

func (s *UpdateWorkerResponseBodyData) SetDeployType(v string) *UpdateWorkerResponseBodyData {
	s.DeployType = &v
	return s
}

func (s *UpdateWorkerResponseBodyData) SetGroups(v []*UpdateWorkerResponseBodyDataGroups) *UpdateWorkerResponseBodyData {
	s.Groups = v
	return s
}

func (s *UpdateWorkerResponseBodyData) SetInstanceId(v string) *UpdateWorkerResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *UpdateWorkerResponseBodyData) SetLimitConfig(v *UpdateWorkerResponseBodyDataLimitConfig) *UpdateWorkerResponseBodyData {
	s.LimitConfig = v
	return s
}

func (s *UpdateWorkerResponseBodyData) SetMcpServers(v []*UpdateWorkerResponseBodyDataMcpServers) *UpdateWorkerResponseBodyData {
	s.McpServers = v
	return s
}

func (s *UpdateWorkerResponseBodyData) SetModel(v *UpdateWorkerResponseBodyDataModel) *UpdateWorkerResponseBodyData {
	s.Model = v
	return s
}

func (s *UpdateWorkerResponseBodyData) SetName(v string) *UpdateWorkerResponseBodyData {
	s.Name = &v
	return s
}

func (s *UpdateWorkerResponseBodyData) SetRegionId(v string) *UpdateWorkerResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *UpdateWorkerResponseBodyData) SetSkills(v []*UpdateWorkerResponseBodyDataSkills) *UpdateWorkerResponseBodyData {
	s.Skills = v
	return s
}

func (s *UpdateWorkerResponseBodyData) SetSoul(v string) *UpdateWorkerResponseBodyData {
	s.Soul = &v
	return s
}

func (s *UpdateWorkerResponseBodyData) SetStartTime(v string) *UpdateWorkerResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *UpdateWorkerResponseBodyData) SetStatus(v string) *UpdateWorkerResponseBodyData {
	s.Status = &v
	return s
}

func (s *UpdateWorkerResponseBodyData) SetTemplate(v *UpdateWorkerResponseBodyDataTemplate) *UpdateWorkerResponseBodyData {
	s.Template = v
	return s
}

func (s *UpdateWorkerResponseBodyData) SetVersionCode(v string) *UpdateWorkerResponseBodyData {
	s.VersionCode = &v
	return s
}

func (s *UpdateWorkerResponseBodyData) Validate() error {
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
	if s.Template != nil {
		if err := s.Template.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateWorkerResponseBodyDataChannels struct {
	Config       *UpdateWorkerResponseBodyDataChannelsConfig       `json:"Config,omitempty" xml:"Config,omitempty" type:"Struct"`
	Enabled      *bool                                             `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	SecretStatus *UpdateWorkerResponseBodyDataChannelsSecretStatus `json:"SecretStatus,omitempty" xml:"SecretStatus,omitempty" type:"Struct"`
	Type         *string                                           `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateWorkerResponseBodyDataChannels) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkerResponseBodyDataChannels) GoString() string {
	return s.String()
}

func (s *UpdateWorkerResponseBodyDataChannels) GetConfig() *UpdateWorkerResponseBodyDataChannelsConfig {
	return s.Config
}

func (s *UpdateWorkerResponseBodyDataChannels) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdateWorkerResponseBodyDataChannels) GetSecretStatus() *UpdateWorkerResponseBodyDataChannelsSecretStatus {
	return s.SecretStatus
}

func (s *UpdateWorkerResponseBodyDataChannels) GetType() *string {
	return s.Type
}

func (s *UpdateWorkerResponseBodyDataChannels) SetConfig(v *UpdateWorkerResponseBodyDataChannelsConfig) *UpdateWorkerResponseBodyDataChannels {
	s.Config = v
	return s
}

func (s *UpdateWorkerResponseBodyDataChannels) SetEnabled(v bool) *UpdateWorkerResponseBodyDataChannels {
	s.Enabled = &v
	return s
}

func (s *UpdateWorkerResponseBodyDataChannels) SetSecretStatus(v *UpdateWorkerResponseBodyDataChannelsSecretStatus) *UpdateWorkerResponseBodyDataChannels {
	s.SecretStatus = v
	return s
}

func (s *UpdateWorkerResponseBodyDataChannels) SetType(v string) *UpdateWorkerResponseBodyDataChannels {
	s.Type = &v
	return s
}

func (s *UpdateWorkerResponseBodyDataChannels) Validate() error {
	if s.Config != nil {
		if err := s.Config.Validate(); err != nil {
			return err
		}
	}
	if s.SecretStatus != nil {
		if err := s.SecretStatus.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateWorkerResponseBodyDataChannelsConfig struct {
	CardTemplateId   *string `json:"CardTemplateId,omitempty" xml:"CardTemplateId,omitempty"`
	ClientId         *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	Extension        *string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	MessageType      *string `json:"MessageType,omitempty" xml:"MessageType,omitempty"`
	RobotCode        *string `json:"RobotCode,omitempty" xml:"RobotCode,omitempty"`
	ShowThinking     *bool   `json:"ShowThinking,omitempty" xml:"ShowThinking,omitempty"`
	ShowToolCalls    *bool   `json:"ShowToolCalls,omitempty" xml:"ShowToolCalls,omitempty"`
	StreamingEnabled *bool   `json:"StreamingEnabled,omitempty" xml:"StreamingEnabled,omitempty"`
}

func (s UpdateWorkerResponseBodyDataChannelsConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkerResponseBodyDataChannelsConfig) GoString() string {
	return s.String()
}

func (s *UpdateWorkerResponseBodyDataChannelsConfig) GetCardTemplateId() *string {
	return s.CardTemplateId
}

func (s *UpdateWorkerResponseBodyDataChannelsConfig) GetClientId() *string {
	return s.ClientId
}

func (s *UpdateWorkerResponseBodyDataChannelsConfig) GetExtension() *string {
	return s.Extension
}

func (s *UpdateWorkerResponseBodyDataChannelsConfig) GetMessageType() *string {
	return s.MessageType
}

func (s *UpdateWorkerResponseBodyDataChannelsConfig) GetRobotCode() *string {
	return s.RobotCode
}

func (s *UpdateWorkerResponseBodyDataChannelsConfig) GetShowThinking() *bool {
	return s.ShowThinking
}

func (s *UpdateWorkerResponseBodyDataChannelsConfig) GetShowToolCalls() *bool {
	return s.ShowToolCalls
}

func (s *UpdateWorkerResponseBodyDataChannelsConfig) GetStreamingEnabled() *bool {
	return s.StreamingEnabled
}

func (s *UpdateWorkerResponseBodyDataChannelsConfig) SetCardTemplateId(v string) *UpdateWorkerResponseBodyDataChannelsConfig {
	s.CardTemplateId = &v
	return s
}

func (s *UpdateWorkerResponseBodyDataChannelsConfig) SetClientId(v string) *UpdateWorkerResponseBodyDataChannelsConfig {
	s.ClientId = &v
	return s
}

func (s *UpdateWorkerResponseBodyDataChannelsConfig) SetExtension(v string) *UpdateWorkerResponseBodyDataChannelsConfig {
	s.Extension = &v
	return s
}

func (s *UpdateWorkerResponseBodyDataChannelsConfig) SetMessageType(v string) *UpdateWorkerResponseBodyDataChannelsConfig {
	s.MessageType = &v
	return s
}

func (s *UpdateWorkerResponseBodyDataChannelsConfig) SetRobotCode(v string) *UpdateWorkerResponseBodyDataChannelsConfig {
	s.RobotCode = &v
	return s
}

func (s *UpdateWorkerResponseBodyDataChannelsConfig) SetShowThinking(v bool) *UpdateWorkerResponseBodyDataChannelsConfig {
	s.ShowThinking = &v
	return s
}

func (s *UpdateWorkerResponseBodyDataChannelsConfig) SetShowToolCalls(v bool) *UpdateWorkerResponseBodyDataChannelsConfig {
	s.ShowToolCalls = &v
	return s
}

func (s *UpdateWorkerResponseBodyDataChannelsConfig) SetStreamingEnabled(v bool) *UpdateWorkerResponseBodyDataChannelsConfig {
	s.StreamingEnabled = &v
	return s
}

func (s *UpdateWorkerResponseBodyDataChannelsConfig) Validate() error {
	return dara.Validate(s)
}

type UpdateWorkerResponseBodyDataChannelsSecretStatus struct {
	ClientSecret *bool `json:"ClientSecret,omitempty" xml:"ClientSecret,omitempty"`
}

func (s UpdateWorkerResponseBodyDataChannelsSecretStatus) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkerResponseBodyDataChannelsSecretStatus) GoString() string {
	return s.String()
}

func (s *UpdateWorkerResponseBodyDataChannelsSecretStatus) GetClientSecret() *bool {
	return s.ClientSecret
}

func (s *UpdateWorkerResponseBodyDataChannelsSecretStatus) SetClientSecret(v bool) *UpdateWorkerResponseBodyDataChannelsSecretStatus {
	s.ClientSecret = &v
	return s
}

func (s *UpdateWorkerResponseBodyDataChannelsSecretStatus) Validate() error {
	return dara.Validate(s)
}

type UpdateWorkerResponseBodyDataCredentials struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateWorkerResponseBodyDataCredentials) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkerResponseBodyDataCredentials) GoString() string {
	return s.String()
}

func (s *UpdateWorkerResponseBodyDataCredentials) GetName() *string {
	return s.Name
}

func (s *UpdateWorkerResponseBodyDataCredentials) SetName(v string) *UpdateWorkerResponseBodyDataCredentials {
	s.Name = &v
	return s
}

func (s *UpdateWorkerResponseBodyDataCredentials) Validate() error {
	return dara.Validate(s)
}

type UpdateWorkerResponseBodyDataGroups struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateWorkerResponseBodyDataGroups) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkerResponseBodyDataGroups) GoString() string {
	return s.String()
}

func (s *UpdateWorkerResponseBodyDataGroups) GetName() *string {
	return s.Name
}

func (s *UpdateWorkerResponseBodyDataGroups) GetRole() *string {
	return s.Role
}

func (s *UpdateWorkerResponseBodyDataGroups) GetType() *string {
	return s.Type
}

func (s *UpdateWorkerResponseBodyDataGroups) SetName(v string) *UpdateWorkerResponseBodyDataGroups {
	s.Name = &v
	return s
}

func (s *UpdateWorkerResponseBodyDataGroups) SetRole(v string) *UpdateWorkerResponseBodyDataGroups {
	s.Role = &v
	return s
}

func (s *UpdateWorkerResponseBodyDataGroups) SetType(v string) *UpdateWorkerResponseBodyDataGroups {
	s.Type = &v
	return s
}

func (s *UpdateWorkerResponseBodyDataGroups) Validate() error {
	return dara.Validate(s)
}

type UpdateWorkerResponseBodyDataLimitConfig struct {
	LimitType  *string `json:"LimitType,omitempty" xml:"LimitType,omitempty"`
	PeriodType *string `json:"PeriodType,omitempty" xml:"PeriodType,omitempty"`
	UsageLimit *int64  `json:"UsageLimit,omitempty" xml:"UsageLimit,omitempty"`
}

func (s UpdateWorkerResponseBodyDataLimitConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkerResponseBodyDataLimitConfig) GoString() string {
	return s.String()
}

func (s *UpdateWorkerResponseBodyDataLimitConfig) GetLimitType() *string {
	return s.LimitType
}

func (s *UpdateWorkerResponseBodyDataLimitConfig) GetPeriodType() *string {
	return s.PeriodType
}

func (s *UpdateWorkerResponseBodyDataLimitConfig) GetUsageLimit() *int64 {
	return s.UsageLimit
}

func (s *UpdateWorkerResponseBodyDataLimitConfig) SetLimitType(v string) *UpdateWorkerResponseBodyDataLimitConfig {
	s.LimitType = &v
	return s
}

func (s *UpdateWorkerResponseBodyDataLimitConfig) SetPeriodType(v string) *UpdateWorkerResponseBodyDataLimitConfig {
	s.PeriodType = &v
	return s
}

func (s *UpdateWorkerResponseBodyDataLimitConfig) SetUsageLimit(v int64) *UpdateWorkerResponseBodyDataLimitConfig {
	s.UsageLimit = &v
	return s
}

func (s *UpdateWorkerResponseBodyDataLimitConfig) Validate() error {
	return dara.Validate(s)
}

type UpdateWorkerResponseBodyDataMcpServers struct {
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Transport *string `json:"Transport,omitempty" xml:"Transport,omitempty"`
	Url       *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s UpdateWorkerResponseBodyDataMcpServers) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkerResponseBodyDataMcpServers) GoString() string {
	return s.String()
}

func (s *UpdateWorkerResponseBodyDataMcpServers) GetName() *string {
	return s.Name
}

func (s *UpdateWorkerResponseBodyDataMcpServers) GetTransport() *string {
	return s.Transport
}

func (s *UpdateWorkerResponseBodyDataMcpServers) GetUrl() *string {
	return s.Url
}

func (s *UpdateWorkerResponseBodyDataMcpServers) SetName(v string) *UpdateWorkerResponseBodyDataMcpServers {
	s.Name = &v
	return s
}

func (s *UpdateWorkerResponseBodyDataMcpServers) SetTransport(v string) *UpdateWorkerResponseBodyDataMcpServers {
	s.Transport = &v
	return s
}

func (s *UpdateWorkerResponseBodyDataMcpServers) SetUrl(v string) *UpdateWorkerResponseBodyDataMcpServers {
	s.Url = &v
	return s
}

func (s *UpdateWorkerResponseBodyDataMcpServers) Validate() error {
	return dara.Validate(s)
}

type UpdateWorkerResponseBodyDataModel struct {
	ModelName     *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	ModelProvider *string `json:"ModelProvider,omitempty" xml:"ModelProvider,omitempty"`
}

func (s UpdateWorkerResponseBodyDataModel) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkerResponseBodyDataModel) GoString() string {
	return s.String()
}

func (s *UpdateWorkerResponseBodyDataModel) GetModelName() *string {
	return s.ModelName
}

func (s *UpdateWorkerResponseBodyDataModel) GetModelProvider() *string {
	return s.ModelProvider
}

func (s *UpdateWorkerResponseBodyDataModel) SetModelName(v string) *UpdateWorkerResponseBodyDataModel {
	s.ModelName = &v
	return s
}

func (s *UpdateWorkerResponseBodyDataModel) SetModelProvider(v string) *UpdateWorkerResponseBodyDataModel {
	s.ModelProvider = &v
	return s
}

func (s *UpdateWorkerResponseBodyDataModel) Validate() error {
	return dara.Validate(s)
}

type UpdateWorkerResponseBodyDataSkills struct {
	Label   *string `json:"Label,omitempty" xml:"Label,omitempty"`
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s UpdateWorkerResponseBodyDataSkills) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkerResponseBodyDataSkills) GoString() string {
	return s.String()
}

func (s *UpdateWorkerResponseBodyDataSkills) GetLabel() *string {
	return s.Label
}

func (s *UpdateWorkerResponseBodyDataSkills) GetName() *string {
	return s.Name
}

func (s *UpdateWorkerResponseBodyDataSkills) GetVersion() *string {
	return s.Version
}

func (s *UpdateWorkerResponseBodyDataSkills) SetLabel(v string) *UpdateWorkerResponseBodyDataSkills {
	s.Label = &v
	return s
}

func (s *UpdateWorkerResponseBodyDataSkills) SetName(v string) *UpdateWorkerResponseBodyDataSkills {
	s.Name = &v
	return s
}

func (s *UpdateWorkerResponseBodyDataSkills) SetVersion(v string) *UpdateWorkerResponseBodyDataSkills {
	s.Version = &v
	return s
}

func (s *UpdateWorkerResponseBodyDataSkills) Validate() error {
	return dara.Validate(s)
}

type UpdateWorkerResponseBodyDataTemplate struct {
	Label   *string `json:"Label,omitempty" xml:"Label,omitempty"`
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s UpdateWorkerResponseBodyDataTemplate) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkerResponseBodyDataTemplate) GoString() string {
	return s.String()
}

func (s *UpdateWorkerResponseBodyDataTemplate) GetLabel() *string {
	return s.Label
}

func (s *UpdateWorkerResponseBodyDataTemplate) GetName() *string {
	return s.Name
}

func (s *UpdateWorkerResponseBodyDataTemplate) GetVersion() *string {
	return s.Version
}

func (s *UpdateWorkerResponseBodyDataTemplate) SetLabel(v string) *UpdateWorkerResponseBodyDataTemplate {
	s.Label = &v
	return s
}

func (s *UpdateWorkerResponseBodyDataTemplate) SetName(v string) *UpdateWorkerResponseBodyDataTemplate {
	s.Name = &v
	return s
}

func (s *UpdateWorkerResponseBodyDataTemplate) SetVersion(v string) *UpdateWorkerResponseBodyDataTemplate {
	s.Version = &v
	return s
}

func (s *UpdateWorkerResponseBodyDataTemplate) Validate() error {
	return dara.Validate(s)
}
