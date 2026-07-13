// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetWorkerResponseBody
	GetCode() *string
	SetData(v *GetWorkerResponseBodyData) *GetWorkerResponseBody
	GetData() *GetWorkerResponseBodyData
	SetHttpStatusCode(v int32) *GetWorkerResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetWorkerResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetWorkerResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetWorkerResponseBody
	GetSuccess() *bool
}

type GetWorkerResponseBody struct {
	Code           *string                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *GetWorkerResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                     `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                    `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                      `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetWorkerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetWorkerResponseBody) GoString() string {
	return s.String()
}

func (s *GetWorkerResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetWorkerResponseBody) GetData() *GetWorkerResponseBodyData {
	return s.Data
}

func (s *GetWorkerResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetWorkerResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetWorkerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetWorkerResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetWorkerResponseBody) SetCode(v string) *GetWorkerResponseBody {
	s.Code = &v
	return s
}

func (s *GetWorkerResponseBody) SetData(v *GetWorkerResponseBodyData) *GetWorkerResponseBody {
	s.Data = v
	return s
}

func (s *GetWorkerResponseBody) SetHttpStatusCode(v int32) *GetWorkerResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetWorkerResponseBody) SetMessage(v string) *GetWorkerResponseBody {
	s.Message = &v
	return s
}

func (s *GetWorkerResponseBody) SetRequestId(v string) *GetWorkerResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWorkerResponseBody) SetSuccess(v bool) *GetWorkerResponseBody {
	s.Success = &v
	return s
}

func (s *GetWorkerResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetWorkerResponseBodyData struct {
	AgentType   *string                                 `json:"AgentType,omitempty" xml:"AgentType,omitempty"`
	Agents      *string                                 `json:"Agents,omitempty" xml:"Agents,omitempty"`
	Channels    []*GetWorkerResponseBodyDataChannels    `json:"Channels,omitempty" xml:"Channels,omitempty" type:"Repeated"`
	Credentials []*GetWorkerResponseBodyDataCredentials `json:"Credentials,omitempty" xml:"Credentials,omitempty" type:"Repeated"`
	DeployType  *string                                 `json:"DeployType,omitempty" xml:"DeployType,omitempty"`
	Groups      []*GetWorkerResponseBodyDataGroups      `json:"Groups,omitempty" xml:"Groups,omitempty" type:"Repeated"`
	InstanceId  *string                                 `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	LimitConfig *GetWorkerResponseBodyDataLimitConfig   `json:"LimitConfig,omitempty" xml:"LimitConfig,omitempty" type:"Struct"`
	McpServers  []*GetWorkerResponseBodyDataMcpServers  `json:"McpServers,omitempty" xml:"McpServers,omitempty" type:"Repeated"`
	Model       *GetWorkerResponseBodyDataModel         `json:"Model,omitempty" xml:"Model,omitempty" type:"Struct"`
	Name        *string                                 `json:"Name,omitempty" xml:"Name,omitempty"`
	RegionId    *string                                 `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Skills      []*GetWorkerResponseBodyDataSkills      `json:"Skills,omitempty" xml:"Skills,omitempty" type:"Repeated"`
	Soul        *string                                 `json:"Soul,omitempty" xml:"Soul,omitempty"`
	StartTime   *string                                 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status      *string                                 `json:"Status,omitempty" xml:"Status,omitempty"`
	Subagents   []*GetWorkerResponseBodyDataSubagents   `json:"Subagents,omitempty" xml:"Subagents,omitempty" type:"Repeated"`
	Template    *GetWorkerResponseBodyDataTemplate      `json:"Template,omitempty" xml:"Template,omitempty" type:"Struct"`
	VersionCode *string                                 `json:"VersionCode,omitempty" xml:"VersionCode,omitempty"`
}

func (s GetWorkerResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetWorkerResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetWorkerResponseBodyData) GetAgentType() *string {
	return s.AgentType
}

func (s *GetWorkerResponseBodyData) GetAgents() *string {
	return s.Agents
}

func (s *GetWorkerResponseBodyData) GetChannels() []*GetWorkerResponseBodyDataChannels {
	return s.Channels
}

func (s *GetWorkerResponseBodyData) GetCredentials() []*GetWorkerResponseBodyDataCredentials {
	return s.Credentials
}

func (s *GetWorkerResponseBodyData) GetDeployType() *string {
	return s.DeployType
}

func (s *GetWorkerResponseBodyData) GetGroups() []*GetWorkerResponseBodyDataGroups {
	return s.Groups
}

func (s *GetWorkerResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetWorkerResponseBodyData) GetLimitConfig() *GetWorkerResponseBodyDataLimitConfig {
	return s.LimitConfig
}

func (s *GetWorkerResponseBodyData) GetMcpServers() []*GetWorkerResponseBodyDataMcpServers {
	return s.McpServers
}

func (s *GetWorkerResponseBodyData) GetModel() *GetWorkerResponseBodyDataModel {
	return s.Model
}

func (s *GetWorkerResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetWorkerResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *GetWorkerResponseBodyData) GetSkills() []*GetWorkerResponseBodyDataSkills {
	return s.Skills
}

func (s *GetWorkerResponseBodyData) GetSoul() *string {
	return s.Soul
}

func (s *GetWorkerResponseBodyData) GetStartTime() *string {
	return s.StartTime
}

func (s *GetWorkerResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetWorkerResponseBodyData) GetSubagents() []*GetWorkerResponseBodyDataSubagents {
	return s.Subagents
}

func (s *GetWorkerResponseBodyData) GetTemplate() *GetWorkerResponseBodyDataTemplate {
	return s.Template
}

func (s *GetWorkerResponseBodyData) GetVersionCode() *string {
	return s.VersionCode
}

func (s *GetWorkerResponseBodyData) SetAgentType(v string) *GetWorkerResponseBodyData {
	s.AgentType = &v
	return s
}

func (s *GetWorkerResponseBodyData) SetAgents(v string) *GetWorkerResponseBodyData {
	s.Agents = &v
	return s
}

func (s *GetWorkerResponseBodyData) SetChannels(v []*GetWorkerResponseBodyDataChannels) *GetWorkerResponseBodyData {
	s.Channels = v
	return s
}

func (s *GetWorkerResponseBodyData) SetCredentials(v []*GetWorkerResponseBodyDataCredentials) *GetWorkerResponseBodyData {
	s.Credentials = v
	return s
}

func (s *GetWorkerResponseBodyData) SetDeployType(v string) *GetWorkerResponseBodyData {
	s.DeployType = &v
	return s
}

func (s *GetWorkerResponseBodyData) SetGroups(v []*GetWorkerResponseBodyDataGroups) *GetWorkerResponseBodyData {
	s.Groups = v
	return s
}

func (s *GetWorkerResponseBodyData) SetInstanceId(v string) *GetWorkerResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *GetWorkerResponseBodyData) SetLimitConfig(v *GetWorkerResponseBodyDataLimitConfig) *GetWorkerResponseBodyData {
	s.LimitConfig = v
	return s
}

func (s *GetWorkerResponseBodyData) SetMcpServers(v []*GetWorkerResponseBodyDataMcpServers) *GetWorkerResponseBodyData {
	s.McpServers = v
	return s
}

func (s *GetWorkerResponseBodyData) SetModel(v *GetWorkerResponseBodyDataModel) *GetWorkerResponseBodyData {
	s.Model = v
	return s
}

func (s *GetWorkerResponseBodyData) SetName(v string) *GetWorkerResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetWorkerResponseBodyData) SetRegionId(v string) *GetWorkerResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *GetWorkerResponseBodyData) SetSkills(v []*GetWorkerResponseBodyDataSkills) *GetWorkerResponseBodyData {
	s.Skills = v
	return s
}

func (s *GetWorkerResponseBodyData) SetSoul(v string) *GetWorkerResponseBodyData {
	s.Soul = &v
	return s
}

func (s *GetWorkerResponseBodyData) SetStartTime(v string) *GetWorkerResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *GetWorkerResponseBodyData) SetStatus(v string) *GetWorkerResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetWorkerResponseBodyData) SetSubagents(v []*GetWorkerResponseBodyDataSubagents) *GetWorkerResponseBodyData {
	s.Subagents = v
	return s
}

func (s *GetWorkerResponseBodyData) SetTemplate(v *GetWorkerResponseBodyDataTemplate) *GetWorkerResponseBodyData {
	s.Template = v
	return s
}

func (s *GetWorkerResponseBodyData) SetVersionCode(v string) *GetWorkerResponseBodyData {
	s.VersionCode = &v
	return s
}

func (s *GetWorkerResponseBodyData) Validate() error {
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

type GetWorkerResponseBodyDataChannels struct {
	Config       *GetWorkerResponseBodyDataChannelsConfig       `json:"Config,omitempty" xml:"Config,omitempty" type:"Struct"`
	Enabled      *bool                                          `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	SecretStatus *GetWorkerResponseBodyDataChannelsSecretStatus `json:"SecretStatus,omitempty" xml:"SecretStatus,omitempty" type:"Struct"`
	Type         *string                                        `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetWorkerResponseBodyDataChannels) String() string {
	return dara.Prettify(s)
}

func (s GetWorkerResponseBodyDataChannels) GoString() string {
	return s.String()
}

func (s *GetWorkerResponseBodyDataChannels) GetConfig() *GetWorkerResponseBodyDataChannelsConfig {
	return s.Config
}

func (s *GetWorkerResponseBodyDataChannels) GetEnabled() *bool {
	return s.Enabled
}

func (s *GetWorkerResponseBodyDataChannels) GetSecretStatus() *GetWorkerResponseBodyDataChannelsSecretStatus {
	return s.SecretStatus
}

func (s *GetWorkerResponseBodyDataChannels) GetType() *string {
	return s.Type
}

func (s *GetWorkerResponseBodyDataChannels) SetConfig(v *GetWorkerResponseBodyDataChannelsConfig) *GetWorkerResponseBodyDataChannels {
	s.Config = v
	return s
}

func (s *GetWorkerResponseBodyDataChannels) SetEnabled(v bool) *GetWorkerResponseBodyDataChannels {
	s.Enabled = &v
	return s
}

func (s *GetWorkerResponseBodyDataChannels) SetSecretStatus(v *GetWorkerResponseBodyDataChannelsSecretStatus) *GetWorkerResponseBodyDataChannels {
	s.SecretStatus = v
	return s
}

func (s *GetWorkerResponseBodyDataChannels) SetType(v string) *GetWorkerResponseBodyDataChannels {
	s.Type = &v
	return s
}

func (s *GetWorkerResponseBodyDataChannels) Validate() error {
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

type GetWorkerResponseBodyDataChannelsConfig struct {
	CardTemplateId   *string `json:"CardTemplateId,omitempty" xml:"CardTemplateId,omitempty"`
	ClientId         *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	Extension        *string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	MessageType      *string `json:"MessageType,omitempty" xml:"MessageType,omitempty"`
	RobotCode        *string `json:"RobotCode,omitempty" xml:"RobotCode,omitempty"`
	ShowThinking     *bool   `json:"ShowThinking,omitempty" xml:"ShowThinking,omitempty"`
	ShowToolCalls    *bool   `json:"ShowToolCalls,omitempty" xml:"ShowToolCalls,omitempty"`
	StreamingEnabled *bool   `json:"StreamingEnabled,omitempty" xml:"StreamingEnabled,omitempty"`
}

func (s GetWorkerResponseBodyDataChannelsConfig) String() string {
	return dara.Prettify(s)
}

func (s GetWorkerResponseBodyDataChannelsConfig) GoString() string {
	return s.String()
}

func (s *GetWorkerResponseBodyDataChannelsConfig) GetCardTemplateId() *string {
	return s.CardTemplateId
}

func (s *GetWorkerResponseBodyDataChannelsConfig) GetClientId() *string {
	return s.ClientId
}

func (s *GetWorkerResponseBodyDataChannelsConfig) GetExtension() *string {
	return s.Extension
}

func (s *GetWorkerResponseBodyDataChannelsConfig) GetMessageType() *string {
	return s.MessageType
}

func (s *GetWorkerResponseBodyDataChannelsConfig) GetRobotCode() *string {
	return s.RobotCode
}

func (s *GetWorkerResponseBodyDataChannelsConfig) GetShowThinking() *bool {
	return s.ShowThinking
}

func (s *GetWorkerResponseBodyDataChannelsConfig) GetShowToolCalls() *bool {
	return s.ShowToolCalls
}

func (s *GetWorkerResponseBodyDataChannelsConfig) GetStreamingEnabled() *bool {
	return s.StreamingEnabled
}

func (s *GetWorkerResponseBodyDataChannelsConfig) SetCardTemplateId(v string) *GetWorkerResponseBodyDataChannelsConfig {
	s.CardTemplateId = &v
	return s
}

func (s *GetWorkerResponseBodyDataChannelsConfig) SetClientId(v string) *GetWorkerResponseBodyDataChannelsConfig {
	s.ClientId = &v
	return s
}

func (s *GetWorkerResponseBodyDataChannelsConfig) SetExtension(v string) *GetWorkerResponseBodyDataChannelsConfig {
	s.Extension = &v
	return s
}

func (s *GetWorkerResponseBodyDataChannelsConfig) SetMessageType(v string) *GetWorkerResponseBodyDataChannelsConfig {
	s.MessageType = &v
	return s
}

func (s *GetWorkerResponseBodyDataChannelsConfig) SetRobotCode(v string) *GetWorkerResponseBodyDataChannelsConfig {
	s.RobotCode = &v
	return s
}

func (s *GetWorkerResponseBodyDataChannelsConfig) SetShowThinking(v bool) *GetWorkerResponseBodyDataChannelsConfig {
	s.ShowThinking = &v
	return s
}

func (s *GetWorkerResponseBodyDataChannelsConfig) SetShowToolCalls(v bool) *GetWorkerResponseBodyDataChannelsConfig {
	s.ShowToolCalls = &v
	return s
}

func (s *GetWorkerResponseBodyDataChannelsConfig) SetStreamingEnabled(v bool) *GetWorkerResponseBodyDataChannelsConfig {
	s.StreamingEnabled = &v
	return s
}

func (s *GetWorkerResponseBodyDataChannelsConfig) Validate() error {
	return dara.Validate(s)
}

type GetWorkerResponseBodyDataChannelsSecretStatus struct {
	ClientSecret *bool `json:"ClientSecret,omitempty" xml:"ClientSecret,omitempty"`
}

func (s GetWorkerResponseBodyDataChannelsSecretStatus) String() string {
	return dara.Prettify(s)
}

func (s GetWorkerResponseBodyDataChannelsSecretStatus) GoString() string {
	return s.String()
}

func (s *GetWorkerResponseBodyDataChannelsSecretStatus) GetClientSecret() *bool {
	return s.ClientSecret
}

func (s *GetWorkerResponseBodyDataChannelsSecretStatus) SetClientSecret(v bool) *GetWorkerResponseBodyDataChannelsSecretStatus {
	s.ClientSecret = &v
	return s
}

func (s *GetWorkerResponseBodyDataChannelsSecretStatus) Validate() error {
	return dara.Validate(s)
}

type GetWorkerResponseBodyDataCredentials struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetWorkerResponseBodyDataCredentials) String() string {
	return dara.Prettify(s)
}

func (s GetWorkerResponseBodyDataCredentials) GoString() string {
	return s.String()
}

func (s *GetWorkerResponseBodyDataCredentials) GetName() *string {
	return s.Name
}

func (s *GetWorkerResponseBodyDataCredentials) SetName(v string) *GetWorkerResponseBodyDataCredentials {
	s.Name = &v
	return s
}

func (s *GetWorkerResponseBodyDataCredentials) Validate() error {
	return dara.Validate(s)
}

type GetWorkerResponseBodyDataGroups struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetWorkerResponseBodyDataGroups) String() string {
	return dara.Prettify(s)
}

func (s GetWorkerResponseBodyDataGroups) GoString() string {
	return s.String()
}

func (s *GetWorkerResponseBodyDataGroups) GetName() *string {
	return s.Name
}

func (s *GetWorkerResponseBodyDataGroups) GetRole() *string {
	return s.Role
}

func (s *GetWorkerResponseBodyDataGroups) GetType() *string {
	return s.Type
}

func (s *GetWorkerResponseBodyDataGroups) SetName(v string) *GetWorkerResponseBodyDataGroups {
	s.Name = &v
	return s
}

func (s *GetWorkerResponseBodyDataGroups) SetRole(v string) *GetWorkerResponseBodyDataGroups {
	s.Role = &v
	return s
}

func (s *GetWorkerResponseBodyDataGroups) SetType(v string) *GetWorkerResponseBodyDataGroups {
	s.Type = &v
	return s
}

func (s *GetWorkerResponseBodyDataGroups) Validate() error {
	return dara.Validate(s)
}

type GetWorkerResponseBodyDataLimitConfig struct {
	LimitType  *string `json:"LimitType,omitempty" xml:"LimitType,omitempty"`
	OverLimit  *bool   `json:"OverLimit,omitempty" xml:"OverLimit,omitempty"`
	PeriodType *string `json:"PeriodType,omitempty" xml:"PeriodType,omitempty"`
	RuleStatus *string `json:"RuleStatus,omitempty" xml:"RuleStatus,omitempty"`
	UsageLimit *int64  `json:"UsageLimit,omitempty" xml:"UsageLimit,omitempty"`
	UsedAmount *int64  `json:"UsedAmount,omitempty" xml:"UsedAmount,omitempty"`
}

func (s GetWorkerResponseBodyDataLimitConfig) String() string {
	return dara.Prettify(s)
}

func (s GetWorkerResponseBodyDataLimitConfig) GoString() string {
	return s.String()
}

func (s *GetWorkerResponseBodyDataLimitConfig) GetLimitType() *string {
	return s.LimitType
}

func (s *GetWorkerResponseBodyDataLimitConfig) GetOverLimit() *bool {
	return s.OverLimit
}

func (s *GetWorkerResponseBodyDataLimitConfig) GetPeriodType() *string {
	return s.PeriodType
}

func (s *GetWorkerResponseBodyDataLimitConfig) GetRuleStatus() *string {
	return s.RuleStatus
}

func (s *GetWorkerResponseBodyDataLimitConfig) GetUsageLimit() *int64 {
	return s.UsageLimit
}

func (s *GetWorkerResponseBodyDataLimitConfig) GetUsedAmount() *int64 {
	return s.UsedAmount
}

func (s *GetWorkerResponseBodyDataLimitConfig) SetLimitType(v string) *GetWorkerResponseBodyDataLimitConfig {
	s.LimitType = &v
	return s
}

func (s *GetWorkerResponseBodyDataLimitConfig) SetOverLimit(v bool) *GetWorkerResponseBodyDataLimitConfig {
	s.OverLimit = &v
	return s
}

func (s *GetWorkerResponseBodyDataLimitConfig) SetPeriodType(v string) *GetWorkerResponseBodyDataLimitConfig {
	s.PeriodType = &v
	return s
}

func (s *GetWorkerResponseBodyDataLimitConfig) SetRuleStatus(v string) *GetWorkerResponseBodyDataLimitConfig {
	s.RuleStatus = &v
	return s
}

func (s *GetWorkerResponseBodyDataLimitConfig) SetUsageLimit(v int64) *GetWorkerResponseBodyDataLimitConfig {
	s.UsageLimit = &v
	return s
}

func (s *GetWorkerResponseBodyDataLimitConfig) SetUsedAmount(v int64) *GetWorkerResponseBodyDataLimitConfig {
	s.UsedAmount = &v
	return s
}

func (s *GetWorkerResponseBodyDataLimitConfig) Validate() error {
	return dara.Validate(s)
}

type GetWorkerResponseBodyDataMcpServers struct {
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Transport *string `json:"Transport,omitempty" xml:"Transport,omitempty"`
	Url       *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetWorkerResponseBodyDataMcpServers) String() string {
	return dara.Prettify(s)
}

func (s GetWorkerResponseBodyDataMcpServers) GoString() string {
	return s.String()
}

func (s *GetWorkerResponseBodyDataMcpServers) GetName() *string {
	return s.Name
}

func (s *GetWorkerResponseBodyDataMcpServers) GetTransport() *string {
	return s.Transport
}

func (s *GetWorkerResponseBodyDataMcpServers) GetUrl() *string {
	return s.Url
}

func (s *GetWorkerResponseBodyDataMcpServers) SetName(v string) *GetWorkerResponseBodyDataMcpServers {
	s.Name = &v
	return s
}

func (s *GetWorkerResponseBodyDataMcpServers) SetTransport(v string) *GetWorkerResponseBodyDataMcpServers {
	s.Transport = &v
	return s
}

func (s *GetWorkerResponseBodyDataMcpServers) SetUrl(v string) *GetWorkerResponseBodyDataMcpServers {
	s.Url = &v
	return s
}

func (s *GetWorkerResponseBodyDataMcpServers) Validate() error {
	return dara.Validate(s)
}

type GetWorkerResponseBodyDataModel struct {
	ModelName     *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	ModelProvider *string `json:"ModelProvider,omitempty" xml:"ModelProvider,omitempty"`
}

func (s GetWorkerResponseBodyDataModel) String() string {
	return dara.Prettify(s)
}

func (s GetWorkerResponseBodyDataModel) GoString() string {
	return s.String()
}

func (s *GetWorkerResponseBodyDataModel) GetModelName() *string {
	return s.ModelName
}

func (s *GetWorkerResponseBodyDataModel) GetModelProvider() *string {
	return s.ModelProvider
}

func (s *GetWorkerResponseBodyDataModel) SetModelName(v string) *GetWorkerResponseBodyDataModel {
	s.ModelName = &v
	return s
}

func (s *GetWorkerResponseBodyDataModel) SetModelProvider(v string) *GetWorkerResponseBodyDataModel {
	s.ModelProvider = &v
	return s
}

func (s *GetWorkerResponseBodyDataModel) Validate() error {
	return dara.Validate(s)
}

type GetWorkerResponseBodyDataSkills struct {
	Label   *string `json:"Label,omitempty" xml:"Label,omitempty"`
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetWorkerResponseBodyDataSkills) String() string {
	return dara.Prettify(s)
}

func (s GetWorkerResponseBodyDataSkills) GoString() string {
	return s.String()
}

func (s *GetWorkerResponseBodyDataSkills) GetLabel() *string {
	return s.Label
}

func (s *GetWorkerResponseBodyDataSkills) GetName() *string {
	return s.Name
}

func (s *GetWorkerResponseBodyDataSkills) GetVersion() *string {
	return s.Version
}

func (s *GetWorkerResponseBodyDataSkills) SetLabel(v string) *GetWorkerResponseBodyDataSkills {
	s.Label = &v
	return s
}

func (s *GetWorkerResponseBodyDataSkills) SetName(v string) *GetWorkerResponseBodyDataSkills {
	s.Name = &v
	return s
}

func (s *GetWorkerResponseBodyDataSkills) SetVersion(v string) *GetWorkerResponseBodyDataSkills {
	s.Version = &v
	return s
}

func (s *GetWorkerResponseBodyDataSkills) Validate() error {
	return dara.Validate(s)
}

type GetWorkerResponseBodyDataSubagents struct {
	Description *string                                     `json:"Description,omitempty" xml:"Description,omitempty"`
	HasAgentsMd *bool                                       `json:"HasAgentsMd,omitempty" xml:"HasAgentsMd,omitempty"`
	Name        *string                                     `json:"Name,omitempty" xml:"Name,omitempty"`
	Skills      []*GetWorkerResponseBodyDataSubagentsSkills `json:"Skills,omitempty" xml:"Skills,omitempty" type:"Repeated"`
	SourcePath  *string                                     `json:"SourcePath,omitempty" xml:"SourcePath,omitempty"`
	SubagentId  *string                                     `json:"SubagentId,omitempty" xml:"SubagentId,omitempty"`
}

func (s GetWorkerResponseBodyDataSubagents) String() string {
	return dara.Prettify(s)
}

func (s GetWorkerResponseBodyDataSubagents) GoString() string {
	return s.String()
}

func (s *GetWorkerResponseBodyDataSubagents) GetDescription() *string {
	return s.Description
}

func (s *GetWorkerResponseBodyDataSubagents) GetHasAgentsMd() *bool {
	return s.HasAgentsMd
}

func (s *GetWorkerResponseBodyDataSubagents) GetName() *string {
	return s.Name
}

func (s *GetWorkerResponseBodyDataSubagents) GetSkills() []*GetWorkerResponseBodyDataSubagentsSkills {
	return s.Skills
}

func (s *GetWorkerResponseBodyDataSubagents) GetSourcePath() *string {
	return s.SourcePath
}

func (s *GetWorkerResponseBodyDataSubagents) GetSubagentId() *string {
	return s.SubagentId
}

func (s *GetWorkerResponseBodyDataSubagents) SetDescription(v string) *GetWorkerResponseBodyDataSubagents {
	s.Description = &v
	return s
}

func (s *GetWorkerResponseBodyDataSubagents) SetHasAgentsMd(v bool) *GetWorkerResponseBodyDataSubagents {
	s.HasAgentsMd = &v
	return s
}

func (s *GetWorkerResponseBodyDataSubagents) SetName(v string) *GetWorkerResponseBodyDataSubagents {
	s.Name = &v
	return s
}

func (s *GetWorkerResponseBodyDataSubagents) SetSkills(v []*GetWorkerResponseBodyDataSubagentsSkills) *GetWorkerResponseBodyDataSubagents {
	s.Skills = v
	return s
}

func (s *GetWorkerResponseBodyDataSubagents) SetSourcePath(v string) *GetWorkerResponseBodyDataSubagents {
	s.SourcePath = &v
	return s
}

func (s *GetWorkerResponseBodyDataSubagents) SetSubagentId(v string) *GetWorkerResponseBodyDataSubagents {
	s.SubagentId = &v
	return s
}

func (s *GetWorkerResponseBodyDataSubagents) Validate() error {
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

type GetWorkerResponseBodyDataSubagentsSkills struct {
	HasSkillMd *bool   `json:"HasSkillMd,omitempty" xml:"HasSkillMd,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	SkillId    *string `json:"SkillId,omitempty" xml:"SkillId,omitempty"`
}

func (s GetWorkerResponseBodyDataSubagentsSkills) String() string {
	return dara.Prettify(s)
}

func (s GetWorkerResponseBodyDataSubagentsSkills) GoString() string {
	return s.String()
}

func (s *GetWorkerResponseBodyDataSubagentsSkills) GetHasSkillMd() *bool {
	return s.HasSkillMd
}

func (s *GetWorkerResponseBodyDataSubagentsSkills) GetName() *string {
	return s.Name
}

func (s *GetWorkerResponseBodyDataSubagentsSkills) GetSkillId() *string {
	return s.SkillId
}

func (s *GetWorkerResponseBodyDataSubagentsSkills) SetHasSkillMd(v bool) *GetWorkerResponseBodyDataSubagentsSkills {
	s.HasSkillMd = &v
	return s
}

func (s *GetWorkerResponseBodyDataSubagentsSkills) SetName(v string) *GetWorkerResponseBodyDataSubagentsSkills {
	s.Name = &v
	return s
}

func (s *GetWorkerResponseBodyDataSubagentsSkills) SetSkillId(v string) *GetWorkerResponseBodyDataSubagentsSkills {
	s.SkillId = &v
	return s
}

func (s *GetWorkerResponseBodyDataSubagentsSkills) Validate() error {
	return dara.Validate(s)
}

type GetWorkerResponseBodyDataTemplate struct {
	Label   *string `json:"Label,omitempty" xml:"Label,omitempty"`
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetWorkerResponseBodyDataTemplate) String() string {
	return dara.Prettify(s)
}

func (s GetWorkerResponseBodyDataTemplate) GoString() string {
	return s.String()
}

func (s *GetWorkerResponseBodyDataTemplate) GetLabel() *string {
	return s.Label
}

func (s *GetWorkerResponseBodyDataTemplate) GetName() *string {
	return s.Name
}

func (s *GetWorkerResponseBodyDataTemplate) GetVersion() *string {
	return s.Version
}

func (s *GetWorkerResponseBodyDataTemplate) SetLabel(v string) *GetWorkerResponseBodyDataTemplate {
	s.Label = &v
	return s
}

func (s *GetWorkerResponseBodyDataTemplate) SetName(v string) *GetWorkerResponseBodyDataTemplate {
	s.Name = &v
	return s
}

func (s *GetWorkerResponseBodyDataTemplate) SetVersion(v string) *GetWorkerResponseBodyDataTemplate {
	s.Version = &v
	return s
}

func (s *GetWorkerResponseBodyDataTemplate) Validate() error {
	return dara.Validate(s)
}
