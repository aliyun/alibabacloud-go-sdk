// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAgentPlatformShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAiPlatformConfigShrink(v string) *ModifyAgentPlatformShrinkRequest
	GetAiPlatformConfigShrink() *string
	SetDBClusterId(v string) *ModifyAgentPlatformShrinkRequest
	GetDBClusterId() *string
	SetName(v string) *ModifyAgentPlatformShrinkRequest
	GetName() *string
	SetRegionId(v string) *ModifyAgentPlatformShrinkRequest
	GetRegionId() *string
}

type ModifyAgentPlatformShrinkRequest struct {
	AiPlatformConfigShrink *string `json:"AiPlatformConfig,omitempty" xml:"AiPlatformConfig,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// amv-bp11q28kvl688****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test_platform
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyAgentPlatformShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyAgentPlatformShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyAgentPlatformShrinkRequest) GetAiPlatformConfigShrink() *string {
	return s.AiPlatformConfigShrink
}

func (s *ModifyAgentPlatformShrinkRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyAgentPlatformShrinkRequest) GetName() *string {
	return s.Name
}

func (s *ModifyAgentPlatformShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyAgentPlatformShrinkRequest) SetAiPlatformConfigShrink(v string) *ModifyAgentPlatformShrinkRequest {
	s.AiPlatformConfigShrink = &v
	return s
}

func (s *ModifyAgentPlatformShrinkRequest) SetDBClusterId(v string) *ModifyAgentPlatformShrinkRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyAgentPlatformShrinkRequest) SetName(v string) *ModifyAgentPlatformShrinkRequest {
	s.Name = &v
	return s
}

func (s *ModifyAgentPlatformShrinkRequest) SetRegionId(v string) *ModifyAgentPlatformShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyAgentPlatformShrinkRequest) Validate() error {
	return dara.Validate(s)
}
