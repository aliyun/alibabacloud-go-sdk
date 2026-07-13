// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWorkerShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentType(v string) *CreateWorkerShrinkRequest
	GetAgentType() *string
	SetAgents(v string) *CreateWorkerShrinkRequest
	GetAgents() *string
	SetChannelsShrink(v string) *CreateWorkerShrinkRequest
	GetChannelsShrink() *string
	SetClientToken(v string) *CreateWorkerShrinkRequest
	GetClientToken() *string
	SetCredentialsShrink(v string) *CreateWorkerShrinkRequest
	GetCredentialsShrink() *string
	SetDeployType(v string) *CreateWorkerShrinkRequest
	GetDeployType() *string
	SetGroupsShrink(v string) *CreateWorkerShrinkRequest
	GetGroupsShrink() *string
	SetInstanceId(v string) *CreateWorkerShrinkRequest
	GetInstanceId() *string
	SetLimitConfigShrink(v string) *CreateWorkerShrinkRequest
	GetLimitConfigShrink() *string
	SetMcpServersShrink(v string) *CreateWorkerShrinkRequest
	GetMcpServersShrink() *string
	SetModelShrink(v string) *CreateWorkerShrinkRequest
	GetModelShrink() *string
	SetName(v string) *CreateWorkerShrinkRequest
	GetName() *string
	SetSkillsShrink(v string) *CreateWorkerShrinkRequest
	GetSkillsShrink() *string
	SetSoul(v string) *CreateWorkerShrinkRequest
	GetSoul() *string
	SetSubagentsShrink(v string) *CreateWorkerShrinkRequest
	GetSubagentsShrink() *string
	SetTemplateShrink(v string) *CreateWorkerShrinkRequest
	GetTemplateShrink() *string
	SetVersionCode(v string) *CreateWorkerShrinkRequest
	GetVersionCode() *string
}

type CreateWorkerShrinkRequest struct {
	AgentType         *string `json:"AgentType,omitempty" xml:"AgentType,omitempty"`
	Agents            *string `json:"Agents,omitempty" xml:"Agents,omitempty"`
	ChannelsShrink    *string `json:"Channels,omitempty" xml:"Channels,omitempty"`
	ClientToken       *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	CredentialsShrink *string `json:"Credentials,omitempty" xml:"Credentials,omitempty"`
	DeployType        *string `json:"DeployType,omitempty" xml:"DeployType,omitempty"`
	GroupsShrink      *string `json:"Groups,omitempty" xml:"Groups,omitempty"`
	InstanceId        *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	LimitConfigShrink *string `json:"LimitConfig,omitempty" xml:"LimitConfig,omitempty"`
	McpServersShrink  *string `json:"McpServers,omitempty" xml:"McpServers,omitempty"`
	ModelShrink       *string `json:"Model,omitempty" xml:"Model,omitempty"`
	Name              *string `json:"Name,omitempty" xml:"Name,omitempty"`
	SkillsShrink      *string `json:"Skills,omitempty" xml:"Skills,omitempty"`
	Soul              *string `json:"Soul,omitempty" xml:"Soul,omitempty"`
	SubagentsShrink   *string `json:"Subagents,omitempty" xml:"Subagents,omitempty"`
	TemplateShrink    *string `json:"Template,omitempty" xml:"Template,omitempty"`
	VersionCode       *string `json:"VersionCode,omitempty" xml:"VersionCode,omitempty"`
}

func (s CreateWorkerShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkerShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateWorkerShrinkRequest) GetAgentType() *string {
	return s.AgentType
}

func (s *CreateWorkerShrinkRequest) GetAgents() *string {
	return s.Agents
}

func (s *CreateWorkerShrinkRequest) GetChannelsShrink() *string {
	return s.ChannelsShrink
}

func (s *CreateWorkerShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateWorkerShrinkRequest) GetCredentialsShrink() *string {
	return s.CredentialsShrink
}

func (s *CreateWorkerShrinkRequest) GetDeployType() *string {
	return s.DeployType
}

func (s *CreateWorkerShrinkRequest) GetGroupsShrink() *string {
	return s.GroupsShrink
}

func (s *CreateWorkerShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateWorkerShrinkRequest) GetLimitConfigShrink() *string {
	return s.LimitConfigShrink
}

func (s *CreateWorkerShrinkRequest) GetMcpServersShrink() *string {
	return s.McpServersShrink
}

func (s *CreateWorkerShrinkRequest) GetModelShrink() *string {
	return s.ModelShrink
}

func (s *CreateWorkerShrinkRequest) GetName() *string {
	return s.Name
}

func (s *CreateWorkerShrinkRequest) GetSkillsShrink() *string {
	return s.SkillsShrink
}

func (s *CreateWorkerShrinkRequest) GetSoul() *string {
	return s.Soul
}

func (s *CreateWorkerShrinkRequest) GetSubagentsShrink() *string {
	return s.SubagentsShrink
}

func (s *CreateWorkerShrinkRequest) GetTemplateShrink() *string {
	return s.TemplateShrink
}

func (s *CreateWorkerShrinkRequest) GetVersionCode() *string {
	return s.VersionCode
}

func (s *CreateWorkerShrinkRequest) SetAgentType(v string) *CreateWorkerShrinkRequest {
	s.AgentType = &v
	return s
}

func (s *CreateWorkerShrinkRequest) SetAgents(v string) *CreateWorkerShrinkRequest {
	s.Agents = &v
	return s
}

func (s *CreateWorkerShrinkRequest) SetChannelsShrink(v string) *CreateWorkerShrinkRequest {
	s.ChannelsShrink = &v
	return s
}

func (s *CreateWorkerShrinkRequest) SetClientToken(v string) *CreateWorkerShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateWorkerShrinkRequest) SetCredentialsShrink(v string) *CreateWorkerShrinkRequest {
	s.CredentialsShrink = &v
	return s
}

func (s *CreateWorkerShrinkRequest) SetDeployType(v string) *CreateWorkerShrinkRequest {
	s.DeployType = &v
	return s
}

func (s *CreateWorkerShrinkRequest) SetGroupsShrink(v string) *CreateWorkerShrinkRequest {
	s.GroupsShrink = &v
	return s
}

func (s *CreateWorkerShrinkRequest) SetInstanceId(v string) *CreateWorkerShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateWorkerShrinkRequest) SetLimitConfigShrink(v string) *CreateWorkerShrinkRequest {
	s.LimitConfigShrink = &v
	return s
}

func (s *CreateWorkerShrinkRequest) SetMcpServersShrink(v string) *CreateWorkerShrinkRequest {
	s.McpServersShrink = &v
	return s
}

func (s *CreateWorkerShrinkRequest) SetModelShrink(v string) *CreateWorkerShrinkRequest {
	s.ModelShrink = &v
	return s
}

func (s *CreateWorkerShrinkRequest) SetName(v string) *CreateWorkerShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateWorkerShrinkRequest) SetSkillsShrink(v string) *CreateWorkerShrinkRequest {
	s.SkillsShrink = &v
	return s
}

func (s *CreateWorkerShrinkRequest) SetSoul(v string) *CreateWorkerShrinkRequest {
	s.Soul = &v
	return s
}

func (s *CreateWorkerShrinkRequest) SetSubagentsShrink(v string) *CreateWorkerShrinkRequest {
	s.SubagentsShrink = &v
	return s
}

func (s *CreateWorkerShrinkRequest) SetTemplateShrink(v string) *CreateWorkerShrinkRequest {
	s.TemplateShrink = &v
	return s
}

func (s *CreateWorkerShrinkRequest) SetVersionCode(v string) *CreateWorkerShrinkRequest {
	s.VersionCode = &v
	return s
}

func (s *CreateWorkerShrinkRequest) Validate() error {
	return dara.Validate(s)
}
