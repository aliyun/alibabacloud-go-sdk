// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartEnvironmentFeatureRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnvironmentId(v string) *RestartEnvironmentFeatureRequest
	GetEnvironmentId() *string
	SetFeatureName(v string) *RestartEnvironmentFeatureRequest
	GetFeatureName() *string
	SetRegionId(v string) *RestartEnvironmentFeatureRequest
	GetRegionId() *string
}

type RestartEnvironmentFeatureRequest struct {
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
	// The region ID. Default value: cn-hangzhou.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s RestartEnvironmentFeatureRequest) String() string {
	return dara.Prettify(s)
}

func (s RestartEnvironmentFeatureRequest) GoString() string {
	return s.String()
}

func (s *RestartEnvironmentFeatureRequest) GetEnvironmentId() *string {
	return s.EnvironmentId
}

func (s *RestartEnvironmentFeatureRequest) GetFeatureName() *string {
	return s.FeatureName
}

func (s *RestartEnvironmentFeatureRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RestartEnvironmentFeatureRequest) SetEnvironmentId(v string) *RestartEnvironmentFeatureRequest {
	s.EnvironmentId = &v
	return s
}

func (s *RestartEnvironmentFeatureRequest) SetFeatureName(v string) *RestartEnvironmentFeatureRequest {
	s.FeatureName = &v
	return s
}

func (s *RestartEnvironmentFeatureRequest) SetRegionId(v string) *RestartEnvironmentFeatureRequest {
	s.RegionId = &v
	return s
}

func (s *RestartEnvironmentFeatureRequest) Validate() error {
	return dara.Validate(s)
}
