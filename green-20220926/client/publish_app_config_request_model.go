// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishAppConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentId(v string) *PublishAppConfigRequest
	GetAgentId() *string
	SetAgentName(v string) *PublishAppConfigRequest
	GetAgentName() *string
	SetAppId(v string) *PublishAppConfigRequest
	GetAppId() *string
	SetConfig(v string) *PublishAppConfigRequest
	GetConfig() *string
	SetEnable(v bool) *PublishAppConfigRequest
	GetEnable() *bool
	SetRegionId(v string) *PublishAppConfigRequest
	GetRegionId() *string
	SetResourceType(v string) *PublishAppConfigRequest
	GetResourceType() *string
}

type PublishAppConfigRequest struct {
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
	// {"agentItemConfigs": "[{\\"agentId\\":\\"ag.abcxxx\\",\\"enable\\":true,\\"name\\":\\"Agent1\\"}]"}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// Specifies whether to enable the feature. Valid values:
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

func (s PublishAppConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s PublishAppConfigRequest) GoString() string {
	return s.String()
}

func (s *PublishAppConfigRequest) GetAgentId() *string {
	return s.AgentId
}

func (s *PublishAppConfigRequest) GetAgentName() *string {
	return s.AgentName
}

func (s *PublishAppConfigRequest) GetAppId() *string {
	return s.AppId
}

func (s *PublishAppConfigRequest) GetConfig() *string {
	return s.Config
}

func (s *PublishAppConfigRequest) GetEnable() *bool {
	return s.Enable
}

func (s *PublishAppConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *PublishAppConfigRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *PublishAppConfigRequest) SetAgentId(v string) *PublishAppConfigRequest {
	s.AgentId = &v
	return s
}

func (s *PublishAppConfigRequest) SetAgentName(v string) *PublishAppConfigRequest {
	s.AgentName = &v
	return s
}

func (s *PublishAppConfigRequest) SetAppId(v string) *PublishAppConfigRequest {
	s.AppId = &v
	return s
}

func (s *PublishAppConfigRequest) SetConfig(v string) *PublishAppConfigRequest {
	s.Config = &v
	return s
}

func (s *PublishAppConfigRequest) SetEnable(v bool) *PublishAppConfigRequest {
	s.Enable = &v
	return s
}

func (s *PublishAppConfigRequest) SetRegionId(v string) *PublishAppConfigRequest {
	s.RegionId = &v
	return s
}

func (s *PublishAppConfigRequest) SetResourceType(v string) *PublishAppConfigRequest {
	s.ResourceType = &v
	return s
}

func (s *PublishAppConfigRequest) Validate() error {
	return dara.Validate(s)
}
