// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWorkerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateWorkerResponseBody
	GetCode() *string
	SetData(v *CreateWorkerResponseBodyData) *CreateWorkerResponseBody
	GetData() *CreateWorkerResponseBodyData
	SetHttpStatusCode(v int32) *CreateWorkerResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateWorkerResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateWorkerResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateWorkerResponseBody
	GetSuccess() *bool
}

type CreateWorkerResponseBody struct {
	Code           *string                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *CreateWorkerResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                        `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                         `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateWorkerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkerResponseBody) GoString() string {
	return s.String()
}

func (s *CreateWorkerResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateWorkerResponseBody) GetData() *CreateWorkerResponseBodyData {
	return s.Data
}

func (s *CreateWorkerResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateWorkerResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateWorkerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateWorkerResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateWorkerResponseBody) SetCode(v string) *CreateWorkerResponseBody {
	s.Code = &v
	return s
}

func (s *CreateWorkerResponseBody) SetData(v *CreateWorkerResponseBodyData) *CreateWorkerResponseBody {
	s.Data = v
	return s
}

func (s *CreateWorkerResponseBody) SetHttpStatusCode(v int32) *CreateWorkerResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateWorkerResponseBody) SetMessage(v string) *CreateWorkerResponseBody {
	s.Message = &v
	return s
}

func (s *CreateWorkerResponseBody) SetRequestId(v string) *CreateWorkerResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateWorkerResponseBody) SetSuccess(v bool) *CreateWorkerResponseBody {
	s.Success = &v
	return s
}

func (s *CreateWorkerResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateWorkerResponseBodyData struct {
	AgentType   *string                                    `json:"AgentType,omitempty" xml:"AgentType,omitempty"`
	Agents      *string                                    `json:"Agents,omitempty" xml:"Agents,omitempty"`
	Credentials []*CreateWorkerResponseBodyDataCredentials `json:"Credentials,omitempty" xml:"Credentials,omitempty" type:"Repeated"`
	DeployType  *string                                    `json:"DeployType,omitempty" xml:"DeployType,omitempty"`
	Groups      []*CreateWorkerResponseBodyDataGroups      `json:"Groups,omitempty" xml:"Groups,omitempty" type:"Repeated"`
	InstanceId  *string                                    `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	LimitConfig *CreateWorkerResponseBodyDataLimitConfig   `json:"LimitConfig,omitempty" xml:"LimitConfig,omitempty" type:"Struct"`
	McpServers  []*CreateWorkerResponseBodyDataMcpServers  `json:"McpServers,omitempty" xml:"McpServers,omitempty" type:"Repeated"`
	Model       *CreateWorkerResponseBodyDataModel         `json:"Model,omitempty" xml:"Model,omitempty" type:"Struct"`
	Name        *string                                    `json:"Name,omitempty" xml:"Name,omitempty"`
	Skills      []*CreateWorkerResponseBodyDataSkills      `json:"Skills,omitempty" xml:"Skills,omitempty" type:"Repeated"`
	Soul        *string                                    `json:"Soul,omitempty" xml:"Soul,omitempty"`
	StartTime   *string                                    `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status      *string                                    `json:"Status,omitempty" xml:"Status,omitempty"`
	Template    *CreateWorkerResponseBodyDataTemplate      `json:"Template,omitempty" xml:"Template,omitempty" type:"Struct"`
	VersionCode *string                                    `json:"VersionCode,omitempty" xml:"VersionCode,omitempty"`
}

func (s CreateWorkerResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkerResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateWorkerResponseBodyData) GetAgentType() *string {
	return s.AgentType
}

func (s *CreateWorkerResponseBodyData) GetAgents() *string {
	return s.Agents
}

func (s *CreateWorkerResponseBodyData) GetCredentials() []*CreateWorkerResponseBodyDataCredentials {
	return s.Credentials
}

func (s *CreateWorkerResponseBodyData) GetDeployType() *string {
	return s.DeployType
}

func (s *CreateWorkerResponseBodyData) GetGroups() []*CreateWorkerResponseBodyDataGroups {
	return s.Groups
}

func (s *CreateWorkerResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateWorkerResponseBodyData) GetLimitConfig() *CreateWorkerResponseBodyDataLimitConfig {
	return s.LimitConfig
}

func (s *CreateWorkerResponseBodyData) GetMcpServers() []*CreateWorkerResponseBodyDataMcpServers {
	return s.McpServers
}

func (s *CreateWorkerResponseBodyData) GetModel() *CreateWorkerResponseBodyDataModel {
	return s.Model
}

func (s *CreateWorkerResponseBodyData) GetName() *string {
	return s.Name
}

func (s *CreateWorkerResponseBodyData) GetSkills() []*CreateWorkerResponseBodyDataSkills {
	return s.Skills
}

func (s *CreateWorkerResponseBodyData) GetSoul() *string {
	return s.Soul
}

func (s *CreateWorkerResponseBodyData) GetStartTime() *string {
	return s.StartTime
}

func (s *CreateWorkerResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *CreateWorkerResponseBodyData) GetTemplate() *CreateWorkerResponseBodyDataTemplate {
	return s.Template
}

func (s *CreateWorkerResponseBodyData) GetVersionCode() *string {
	return s.VersionCode
}

func (s *CreateWorkerResponseBodyData) SetAgentType(v string) *CreateWorkerResponseBodyData {
	s.AgentType = &v
	return s
}

func (s *CreateWorkerResponseBodyData) SetAgents(v string) *CreateWorkerResponseBodyData {
	s.Agents = &v
	return s
}

func (s *CreateWorkerResponseBodyData) SetCredentials(v []*CreateWorkerResponseBodyDataCredentials) *CreateWorkerResponseBodyData {
	s.Credentials = v
	return s
}

func (s *CreateWorkerResponseBodyData) SetDeployType(v string) *CreateWorkerResponseBodyData {
	s.DeployType = &v
	return s
}

func (s *CreateWorkerResponseBodyData) SetGroups(v []*CreateWorkerResponseBodyDataGroups) *CreateWorkerResponseBodyData {
	s.Groups = v
	return s
}

func (s *CreateWorkerResponseBodyData) SetInstanceId(v string) *CreateWorkerResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *CreateWorkerResponseBodyData) SetLimitConfig(v *CreateWorkerResponseBodyDataLimitConfig) *CreateWorkerResponseBodyData {
	s.LimitConfig = v
	return s
}

func (s *CreateWorkerResponseBodyData) SetMcpServers(v []*CreateWorkerResponseBodyDataMcpServers) *CreateWorkerResponseBodyData {
	s.McpServers = v
	return s
}

func (s *CreateWorkerResponseBodyData) SetModel(v *CreateWorkerResponseBodyDataModel) *CreateWorkerResponseBodyData {
	s.Model = v
	return s
}

func (s *CreateWorkerResponseBodyData) SetName(v string) *CreateWorkerResponseBodyData {
	s.Name = &v
	return s
}

func (s *CreateWorkerResponseBodyData) SetSkills(v []*CreateWorkerResponseBodyDataSkills) *CreateWorkerResponseBodyData {
	s.Skills = v
	return s
}

func (s *CreateWorkerResponseBodyData) SetSoul(v string) *CreateWorkerResponseBodyData {
	s.Soul = &v
	return s
}

func (s *CreateWorkerResponseBodyData) SetStartTime(v string) *CreateWorkerResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *CreateWorkerResponseBodyData) SetStatus(v string) *CreateWorkerResponseBodyData {
	s.Status = &v
	return s
}

func (s *CreateWorkerResponseBodyData) SetTemplate(v *CreateWorkerResponseBodyDataTemplate) *CreateWorkerResponseBodyData {
	s.Template = v
	return s
}

func (s *CreateWorkerResponseBodyData) SetVersionCode(v string) *CreateWorkerResponseBodyData {
	s.VersionCode = &v
	return s
}

func (s *CreateWorkerResponseBodyData) Validate() error {
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

type CreateWorkerResponseBodyDataCredentials struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateWorkerResponseBodyDataCredentials) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkerResponseBodyDataCredentials) GoString() string {
	return s.String()
}

func (s *CreateWorkerResponseBodyDataCredentials) GetName() *string {
	return s.Name
}

func (s *CreateWorkerResponseBodyDataCredentials) SetName(v string) *CreateWorkerResponseBodyDataCredentials {
	s.Name = &v
	return s
}

func (s *CreateWorkerResponseBodyDataCredentials) Validate() error {
	return dara.Validate(s)
}

type CreateWorkerResponseBodyDataGroups struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateWorkerResponseBodyDataGroups) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkerResponseBodyDataGroups) GoString() string {
	return s.String()
}

func (s *CreateWorkerResponseBodyDataGroups) GetName() *string {
	return s.Name
}

func (s *CreateWorkerResponseBodyDataGroups) GetRole() *string {
	return s.Role
}

func (s *CreateWorkerResponseBodyDataGroups) GetType() *string {
	return s.Type
}

func (s *CreateWorkerResponseBodyDataGroups) SetName(v string) *CreateWorkerResponseBodyDataGroups {
	s.Name = &v
	return s
}

func (s *CreateWorkerResponseBodyDataGroups) SetRole(v string) *CreateWorkerResponseBodyDataGroups {
	s.Role = &v
	return s
}

func (s *CreateWorkerResponseBodyDataGroups) SetType(v string) *CreateWorkerResponseBodyDataGroups {
	s.Type = &v
	return s
}

func (s *CreateWorkerResponseBodyDataGroups) Validate() error {
	return dara.Validate(s)
}

type CreateWorkerResponseBodyDataLimitConfig struct {
	LimitType  *string `json:"LimitType,omitempty" xml:"LimitType,omitempty"`
	PeriodType *string `json:"PeriodType,omitempty" xml:"PeriodType,omitempty"`
	UsageLimit *int64  `json:"UsageLimit,omitempty" xml:"UsageLimit,omitempty"`
}

func (s CreateWorkerResponseBodyDataLimitConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkerResponseBodyDataLimitConfig) GoString() string {
	return s.String()
}

func (s *CreateWorkerResponseBodyDataLimitConfig) GetLimitType() *string {
	return s.LimitType
}

func (s *CreateWorkerResponseBodyDataLimitConfig) GetPeriodType() *string {
	return s.PeriodType
}

func (s *CreateWorkerResponseBodyDataLimitConfig) GetUsageLimit() *int64 {
	return s.UsageLimit
}

func (s *CreateWorkerResponseBodyDataLimitConfig) SetLimitType(v string) *CreateWorkerResponseBodyDataLimitConfig {
	s.LimitType = &v
	return s
}

func (s *CreateWorkerResponseBodyDataLimitConfig) SetPeriodType(v string) *CreateWorkerResponseBodyDataLimitConfig {
	s.PeriodType = &v
	return s
}

func (s *CreateWorkerResponseBodyDataLimitConfig) SetUsageLimit(v int64) *CreateWorkerResponseBodyDataLimitConfig {
	s.UsageLimit = &v
	return s
}

func (s *CreateWorkerResponseBodyDataLimitConfig) Validate() error {
	return dara.Validate(s)
}

type CreateWorkerResponseBodyDataMcpServers struct {
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Transport *string `json:"Transport,omitempty" xml:"Transport,omitempty"`
	Url       *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s CreateWorkerResponseBodyDataMcpServers) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkerResponseBodyDataMcpServers) GoString() string {
	return s.String()
}

func (s *CreateWorkerResponseBodyDataMcpServers) GetName() *string {
	return s.Name
}

func (s *CreateWorkerResponseBodyDataMcpServers) GetTransport() *string {
	return s.Transport
}

func (s *CreateWorkerResponseBodyDataMcpServers) GetUrl() *string {
	return s.Url
}

func (s *CreateWorkerResponseBodyDataMcpServers) SetName(v string) *CreateWorkerResponseBodyDataMcpServers {
	s.Name = &v
	return s
}

func (s *CreateWorkerResponseBodyDataMcpServers) SetTransport(v string) *CreateWorkerResponseBodyDataMcpServers {
	s.Transport = &v
	return s
}

func (s *CreateWorkerResponseBodyDataMcpServers) SetUrl(v string) *CreateWorkerResponseBodyDataMcpServers {
	s.Url = &v
	return s
}

func (s *CreateWorkerResponseBodyDataMcpServers) Validate() error {
	return dara.Validate(s)
}

type CreateWorkerResponseBodyDataModel struct {
	ModelName     *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	ModelProvider *string `json:"ModelProvider,omitempty" xml:"ModelProvider,omitempty"`
}

func (s CreateWorkerResponseBodyDataModel) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkerResponseBodyDataModel) GoString() string {
	return s.String()
}

func (s *CreateWorkerResponseBodyDataModel) GetModelName() *string {
	return s.ModelName
}

func (s *CreateWorkerResponseBodyDataModel) GetModelProvider() *string {
	return s.ModelProvider
}

func (s *CreateWorkerResponseBodyDataModel) SetModelName(v string) *CreateWorkerResponseBodyDataModel {
	s.ModelName = &v
	return s
}

func (s *CreateWorkerResponseBodyDataModel) SetModelProvider(v string) *CreateWorkerResponseBodyDataModel {
	s.ModelProvider = &v
	return s
}

func (s *CreateWorkerResponseBodyDataModel) Validate() error {
	return dara.Validate(s)
}

type CreateWorkerResponseBodyDataSkills struct {
	Label   *string `json:"Label,omitempty" xml:"Label,omitempty"`
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s CreateWorkerResponseBodyDataSkills) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkerResponseBodyDataSkills) GoString() string {
	return s.String()
}

func (s *CreateWorkerResponseBodyDataSkills) GetLabel() *string {
	return s.Label
}

func (s *CreateWorkerResponseBodyDataSkills) GetName() *string {
	return s.Name
}

func (s *CreateWorkerResponseBodyDataSkills) GetVersion() *string {
	return s.Version
}

func (s *CreateWorkerResponseBodyDataSkills) SetLabel(v string) *CreateWorkerResponseBodyDataSkills {
	s.Label = &v
	return s
}

func (s *CreateWorkerResponseBodyDataSkills) SetName(v string) *CreateWorkerResponseBodyDataSkills {
	s.Name = &v
	return s
}

func (s *CreateWorkerResponseBodyDataSkills) SetVersion(v string) *CreateWorkerResponseBodyDataSkills {
	s.Version = &v
	return s
}

func (s *CreateWorkerResponseBodyDataSkills) Validate() error {
	return dara.Validate(s)
}

type CreateWorkerResponseBodyDataTemplate struct {
	Label   *string `json:"Label,omitempty" xml:"Label,omitempty"`
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s CreateWorkerResponseBodyDataTemplate) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkerResponseBodyDataTemplate) GoString() string {
	return s.String()
}

func (s *CreateWorkerResponseBodyDataTemplate) GetLabel() *string {
	return s.Label
}

func (s *CreateWorkerResponseBodyDataTemplate) GetName() *string {
	return s.Name
}

func (s *CreateWorkerResponseBodyDataTemplate) GetVersion() *string {
	return s.Version
}

func (s *CreateWorkerResponseBodyDataTemplate) SetLabel(v string) *CreateWorkerResponseBodyDataTemplate {
	s.Label = &v
	return s
}

func (s *CreateWorkerResponseBodyDataTemplate) SetName(v string) *CreateWorkerResponseBodyDataTemplate {
	s.Name = &v
	return s
}

func (s *CreateWorkerResponseBodyDataTemplate) SetVersion(v string) *CreateWorkerResponseBodyDataTemplate {
	s.Version = &v
	return s
}

func (s *CreateWorkerResponseBodyDataTemplate) Validate() error {
	return dara.Validate(s)
}
