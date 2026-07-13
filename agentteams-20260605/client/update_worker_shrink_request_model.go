// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWorkerShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgents(v string) *UpdateWorkerShrinkRequest
	GetAgents() *string
	SetChannelsShrink(v string) *UpdateWorkerShrinkRequest
	GetChannelsShrink() *string
	SetClientToken(v string) *UpdateWorkerShrinkRequest
	GetClientToken() *string
	SetCredentialsShrink(v string) *UpdateWorkerShrinkRequest
	GetCredentialsShrink() *string
	SetInstanceId(v string) *UpdateWorkerShrinkRequest
	GetInstanceId() *string
	SetLimitConfigShrink(v string) *UpdateWorkerShrinkRequest
	GetLimitConfigShrink() *string
	SetMcpServersShrink(v string) *UpdateWorkerShrinkRequest
	GetMcpServersShrink() *string
	SetModelShrink(v string) *UpdateWorkerShrinkRequest
	GetModelShrink() *string
	SetName(v string) *UpdateWorkerShrinkRequest
	GetName() *string
	SetSkillsShrink(v string) *UpdateWorkerShrinkRequest
	GetSkillsShrink() *string
	SetSoul(v string) *UpdateWorkerShrinkRequest
	GetSoul() *string
	SetTemplateShrink(v string) *UpdateWorkerShrinkRequest
	GetTemplateShrink() *string
	SetVersionCode(v string) *UpdateWorkerShrinkRequest
	GetVersionCode() *string
}

type UpdateWorkerShrinkRequest struct {
	Agents            *string `json:"Agents,omitempty" xml:"Agents,omitempty"`
	ChannelsShrink    *string `json:"Channels,omitempty" xml:"Channels,omitempty"`
	ClientToken       *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	CredentialsShrink *string `json:"Credentials,omitempty" xml:"Credentials,omitempty"`
	// This parameter is required.
	InstanceId        *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	LimitConfigShrink *string `json:"LimitConfig,omitempty" xml:"LimitConfig,omitempty"`
	McpServersShrink  *string `json:"McpServers,omitempty" xml:"McpServers,omitempty"`
	ModelShrink       *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// This parameter is required.
	Name           *string `json:"Name,omitempty" xml:"Name,omitempty"`
	SkillsShrink   *string `json:"Skills,omitempty" xml:"Skills,omitempty"`
	Soul           *string `json:"Soul,omitempty" xml:"Soul,omitempty"`
	TemplateShrink *string `json:"Template,omitempty" xml:"Template,omitempty"`
	VersionCode    *string `json:"VersionCode,omitempty" xml:"VersionCode,omitempty"`
}

func (s UpdateWorkerShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkerShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateWorkerShrinkRequest) GetAgents() *string {
	return s.Agents
}

func (s *UpdateWorkerShrinkRequest) GetChannelsShrink() *string {
	return s.ChannelsShrink
}

func (s *UpdateWorkerShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateWorkerShrinkRequest) GetCredentialsShrink() *string {
	return s.CredentialsShrink
}

func (s *UpdateWorkerShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateWorkerShrinkRequest) GetLimitConfigShrink() *string {
	return s.LimitConfigShrink
}

func (s *UpdateWorkerShrinkRequest) GetMcpServersShrink() *string {
	return s.McpServersShrink
}

func (s *UpdateWorkerShrinkRequest) GetModelShrink() *string {
	return s.ModelShrink
}

func (s *UpdateWorkerShrinkRequest) GetName() *string {
	return s.Name
}

func (s *UpdateWorkerShrinkRequest) GetSkillsShrink() *string {
	return s.SkillsShrink
}

func (s *UpdateWorkerShrinkRequest) GetSoul() *string {
	return s.Soul
}

func (s *UpdateWorkerShrinkRequest) GetTemplateShrink() *string {
	return s.TemplateShrink
}

func (s *UpdateWorkerShrinkRequest) GetVersionCode() *string {
	return s.VersionCode
}

func (s *UpdateWorkerShrinkRequest) SetAgents(v string) *UpdateWorkerShrinkRequest {
	s.Agents = &v
	return s
}

func (s *UpdateWorkerShrinkRequest) SetChannelsShrink(v string) *UpdateWorkerShrinkRequest {
	s.ChannelsShrink = &v
	return s
}

func (s *UpdateWorkerShrinkRequest) SetClientToken(v string) *UpdateWorkerShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateWorkerShrinkRequest) SetCredentialsShrink(v string) *UpdateWorkerShrinkRequest {
	s.CredentialsShrink = &v
	return s
}

func (s *UpdateWorkerShrinkRequest) SetInstanceId(v string) *UpdateWorkerShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateWorkerShrinkRequest) SetLimitConfigShrink(v string) *UpdateWorkerShrinkRequest {
	s.LimitConfigShrink = &v
	return s
}

func (s *UpdateWorkerShrinkRequest) SetMcpServersShrink(v string) *UpdateWorkerShrinkRequest {
	s.McpServersShrink = &v
	return s
}

func (s *UpdateWorkerShrinkRequest) SetModelShrink(v string) *UpdateWorkerShrinkRequest {
	s.ModelShrink = &v
	return s
}

func (s *UpdateWorkerShrinkRequest) SetName(v string) *UpdateWorkerShrinkRequest {
	s.Name = &v
	return s
}

func (s *UpdateWorkerShrinkRequest) SetSkillsShrink(v string) *UpdateWorkerShrinkRequest {
	s.SkillsShrink = &v
	return s
}

func (s *UpdateWorkerShrinkRequest) SetSoul(v string) *UpdateWorkerShrinkRequest {
	s.Soul = &v
	return s
}

func (s *UpdateWorkerShrinkRequest) SetTemplateShrink(v string) *UpdateWorkerShrinkRequest {
	s.TemplateShrink = &v
	return s
}

func (s *UpdateWorkerShrinkRequest) SetVersionCode(v string) *UpdateWorkerShrinkRequest {
	s.VersionCode = &v
	return s
}

func (s *UpdateWorkerShrinkRequest) Validate() error {
	return dara.Validate(s)
}
