// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAgentPlatformShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAiPlatformConfigShrink(v string) *CreateAgentPlatformShrinkRequest
	GetAiPlatformConfigShrink() *string
	SetDBClusterId(v string) *CreateAgentPlatformShrinkRequest
	GetDBClusterId() *string
	SetName(v string) *CreateAgentPlatformShrinkRequest
	GetName() *string
	SetRegionId(v string) *CreateAgentPlatformShrinkRequest
	GetRegionId() *string
}

type CreateAgentPlatformShrinkRequest struct {
	// This parameter is required.
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
	// testplatform
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateAgentPlatformShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAgentPlatformShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateAgentPlatformShrinkRequest) GetAiPlatformConfigShrink() *string {
	return s.AiPlatformConfigShrink
}

func (s *CreateAgentPlatformShrinkRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *CreateAgentPlatformShrinkRequest) GetName() *string {
	return s.Name
}

func (s *CreateAgentPlatformShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateAgentPlatformShrinkRequest) SetAiPlatformConfigShrink(v string) *CreateAgentPlatformShrinkRequest {
	s.AiPlatformConfigShrink = &v
	return s
}

func (s *CreateAgentPlatformShrinkRequest) SetDBClusterId(v string) *CreateAgentPlatformShrinkRequest {
	s.DBClusterId = &v
	return s
}

func (s *CreateAgentPlatformShrinkRequest) SetName(v string) *CreateAgentPlatformShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateAgentPlatformShrinkRequest) SetRegionId(v string) *CreateAgentPlatformShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *CreateAgentPlatformShrinkRequest) Validate() error {
	return dara.Validate(s)
}
