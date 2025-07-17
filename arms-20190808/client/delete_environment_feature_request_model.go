// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEnvironmentFeatureRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnvironmentId(v string) *DeleteEnvironmentFeatureRequest
	GetEnvironmentId() *string
	SetFeatureName(v string) *DeleteEnvironmentFeatureRequest
	GetFeatureName() *string
	SetRegionId(v string) *DeleteEnvironmentFeatureRequest
	GetRegionId() *string
}

type DeleteEnvironmentFeatureRequest struct {
	// The ID of the environment.
	//
	// This parameter is required.
	//
	// example:
	//
	// env-xxx
	EnvironmentId *string `json:"EnvironmentId,omitempty" xml:"EnvironmentId,omitempty"`
	// The feature name. Valid values: app-agent-pilot, metric-agent, ebpf-agent, and service-check.
	//
	// This parameter is required.
	//
	// example:
	//
	// metric-agent
	FeatureName *string `json:"FeatureName,omitempty" xml:"FeatureName,omitempty"`
	// The region ID. Valid values: cn-beijing and cn-hangzhou.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteEnvironmentFeatureRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteEnvironmentFeatureRequest) GoString() string {
	return s.String()
}

func (s *DeleteEnvironmentFeatureRequest) GetEnvironmentId() *string {
	return s.EnvironmentId
}

func (s *DeleteEnvironmentFeatureRequest) GetFeatureName() *string {
	return s.FeatureName
}

func (s *DeleteEnvironmentFeatureRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteEnvironmentFeatureRequest) SetEnvironmentId(v string) *DeleteEnvironmentFeatureRequest {
	s.EnvironmentId = &v
	return s
}

func (s *DeleteEnvironmentFeatureRequest) SetFeatureName(v string) *DeleteEnvironmentFeatureRequest {
	s.FeatureName = &v
	return s
}

func (s *DeleteEnvironmentFeatureRequest) SetRegionId(v string) *DeleteEnvironmentFeatureRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteEnvironmentFeatureRequest) Validate() error {
	return dara.Validate(s)
}
