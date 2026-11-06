// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourceGroupAgentModelSettingsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentPlatform(v string) *GetResourceGroupAgentModelSettingsRequest
	GetAgentPlatform() *string
	SetAgentProvider(v string) *GetResourceGroupAgentModelSettingsRequest
	GetAgentProvider() *string
	SetBizType(v int32) *GetResourceGroupAgentModelSettingsRequest
	GetBizType() *int32
	SetResourceGroupIds(v []*string) *GetResourceGroupAgentModelSettingsRequest
	GetResourceGroupIds() []*string
}

type GetResourceGroupAgentModelSettingsRequest struct {
	// The agent platform type.
	//
	// example:
	//
	// ENTERPRISE
	AgentPlatform *string `json:"AgentPlatform,omitempty" xml:"AgentPlatform,omitempty"`
	// The agent provider name.
	//
	// This parameter is required.
	//
	// example:
	//
	// OpenClaw
	AgentProvider *string `json:"AgentProvider,omitempty" xml:"AgentProvider,omitempty"`
	// The business type.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	BizType *int32 `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// The list of resource group IDs.
	//
	// This parameter is required.
	ResourceGroupIds []*string `json:"ResourceGroupIds,omitempty" xml:"ResourceGroupIds,omitempty" type:"Repeated"`
}

func (s GetResourceGroupAgentModelSettingsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetResourceGroupAgentModelSettingsRequest) GoString() string {
	return s.String()
}

func (s *GetResourceGroupAgentModelSettingsRequest) GetAgentPlatform() *string {
	return s.AgentPlatform
}

func (s *GetResourceGroupAgentModelSettingsRequest) GetAgentProvider() *string {
	return s.AgentProvider
}

func (s *GetResourceGroupAgentModelSettingsRequest) GetBizType() *int32 {
	return s.BizType
}

func (s *GetResourceGroupAgentModelSettingsRequest) GetResourceGroupIds() []*string {
	return s.ResourceGroupIds
}

func (s *GetResourceGroupAgentModelSettingsRequest) SetAgentPlatform(v string) *GetResourceGroupAgentModelSettingsRequest {
	s.AgentPlatform = &v
	return s
}

func (s *GetResourceGroupAgentModelSettingsRequest) SetAgentProvider(v string) *GetResourceGroupAgentModelSettingsRequest {
	s.AgentProvider = &v
	return s
}

func (s *GetResourceGroupAgentModelSettingsRequest) SetBizType(v int32) *GetResourceGroupAgentModelSettingsRequest {
	s.BizType = &v
	return s
}

func (s *GetResourceGroupAgentModelSettingsRequest) SetResourceGroupIds(v []*string) *GetResourceGroupAgentModelSettingsRequest {
	s.ResourceGroupIds = v
	return s
}

func (s *GetResourceGroupAgentModelSettingsRequest) Validate() error {
	return dara.Validate(s)
}
