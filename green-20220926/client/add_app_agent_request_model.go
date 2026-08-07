// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddAppAgentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentId(v string) *AddAppAgentRequest
	GetAgentId() *string
	SetAgentName(v string) *AddAppAgentRequest
	GetAgentName() *string
	SetAppId(v string) *AddAppAgentRequest
	GetAppId() *string
	SetEnable(v bool) *AddAppAgentRequest
	GetEnable() *bool
	SetRegionId(v string) *AddAppAgentRequest
	GetRegionId() *string
	SetResourceType(v string) *AddAppAgentRequest
	GetResourceType() *string
}

type AddAppAgentRequest struct {
	// Agent ID
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

func (s AddAppAgentRequest) String() string {
	return dara.Prettify(s)
}

func (s AddAppAgentRequest) GoString() string {
	return s.String()
}

func (s *AddAppAgentRequest) GetAgentId() *string {
	return s.AgentId
}

func (s *AddAppAgentRequest) GetAgentName() *string {
	return s.AgentName
}

func (s *AddAppAgentRequest) GetAppId() *string {
	return s.AppId
}

func (s *AddAppAgentRequest) GetEnable() *bool {
	return s.Enable
}

func (s *AddAppAgentRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AddAppAgentRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *AddAppAgentRequest) SetAgentId(v string) *AddAppAgentRequest {
	s.AgentId = &v
	return s
}

func (s *AddAppAgentRequest) SetAgentName(v string) *AddAppAgentRequest {
	s.AgentName = &v
	return s
}

func (s *AddAppAgentRequest) SetAppId(v string) *AddAppAgentRequest {
	s.AppId = &v
	return s
}

func (s *AddAppAgentRequest) SetEnable(v bool) *AddAppAgentRequest {
	s.Enable = &v
	return s
}

func (s *AddAppAgentRequest) SetRegionId(v string) *AddAppAgentRequest {
	s.RegionId = &v
	return s
}

func (s *AddAppAgentRequest) SetResourceType(v string) *AddAppAgentRequest {
	s.ResourceType = &v
	return s
}

func (s *AddAppAgentRequest) Validate() error {
	return dara.Validate(s)
}
