// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAppAgentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentId(v string) *ModifyAppAgentRequest
	GetAgentId() *string
	SetAgentName(v string) *ModifyAppAgentRequest
	GetAgentName() *string
	SetAppId(v string) *ModifyAppAgentRequest
	GetAppId() *string
	SetConfig(v string) *ModifyAppAgentRequest
	GetConfig() *string
	SetEnable(v bool) *ModifyAppAgentRequest
	GetEnable() *bool
	SetRegionId(v string) *ModifyAppAgentRequest
	GetRegionId() *string
	SetResourceType(v string) *ModifyAppAgentRequest
	GetResourceType() *string
}

type ModifyAppAgentRequest struct {
	// Agent ID。
	//
	// example:
	//
	// ag.abcxxx
	AgentId *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// The agent name.
	//
	// example:
	//
	// Agent1
	AgentName *string `json:"AgentName,omitempty" xml:"AgentName,omitempty"`
	// App ID。
	//
	// example:
	//
	// txt_check_agent_01
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The configuration details.
	//
	// example:
	//
	// {"model":"default","scene":"0swLgojx","labelConfig":[{"label":"Abuse","labelDefinition":"Text content in the reviewed text that contains abusive language"}]}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// Specifies whether to enable the agent. Valid values:
	//
	// - **true**: Enabled.
	//
	// - **false**: Disabled.
	//
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource type.
	//
	// example:
	//
	// agent_text
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s ModifyAppAgentRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyAppAgentRequest) GoString() string {
	return s.String()
}

func (s *ModifyAppAgentRequest) GetAgentId() *string {
	return s.AgentId
}

func (s *ModifyAppAgentRequest) GetAgentName() *string {
	return s.AgentName
}

func (s *ModifyAppAgentRequest) GetAppId() *string {
	return s.AppId
}

func (s *ModifyAppAgentRequest) GetConfig() *string {
	return s.Config
}

func (s *ModifyAppAgentRequest) GetEnable() *bool {
	return s.Enable
}

func (s *ModifyAppAgentRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyAppAgentRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *ModifyAppAgentRequest) SetAgentId(v string) *ModifyAppAgentRequest {
	s.AgentId = &v
	return s
}

func (s *ModifyAppAgentRequest) SetAgentName(v string) *ModifyAppAgentRequest {
	s.AgentName = &v
	return s
}

func (s *ModifyAppAgentRequest) SetAppId(v string) *ModifyAppAgentRequest {
	s.AppId = &v
	return s
}

func (s *ModifyAppAgentRequest) SetConfig(v string) *ModifyAppAgentRequest {
	s.Config = &v
	return s
}

func (s *ModifyAppAgentRequest) SetEnable(v bool) *ModifyAppAgentRequest {
	s.Enable = &v
	return s
}

func (s *ModifyAppAgentRequest) SetRegionId(v string) *ModifyAppAgentRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyAppAgentRequest) SetResourceType(v string) *ModifyAppAgentRequest {
	s.ResourceType = &v
	return s
}

func (s *ModifyAppAgentRequest) Validate() error {
	return dara.Validate(s)
}
