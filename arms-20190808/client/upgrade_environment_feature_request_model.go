// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeEnvironmentFeatureRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliyunLang(v string) *UpgradeEnvironmentFeatureRequest
	GetAliyunLang() *string
	SetEnvironmentId(v string) *UpgradeEnvironmentFeatureRequest
	GetEnvironmentId() *string
	SetFeatureName(v string) *UpgradeEnvironmentFeatureRequest
	GetFeatureName() *string
	SetFeatureVersion(v string) *UpgradeEnvironmentFeatureRequest
	GetFeatureVersion() *string
	SetRegionId(v string) *UpgradeEnvironmentFeatureRequest
	GetRegionId() *string
	SetValues(v string) *UpgradeEnvironmentFeatureRequest
	GetValues() *string
}

type UpgradeEnvironmentFeatureRequest struct {
	// The language. Valid values: zh and en. Default value: zh.
	//
	// example:
	//
	// zh
	AliyunLang *string `json:"AliyunLang,omitempty" xml:"AliyunLang,omitempty"`
	// The environment ID.
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
	// The version of the feature.
	//
	// example:
	//
	// 1.1.17
	FeatureVersion *string `json:"FeatureVersion,omitempty" xml:"FeatureVersion,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Specifies whether to enable service discovery. For PodAnnotation, set the value to run or mini. For PodMonitor and ServiceMonitor, set the value to true or false.
	//
	// example:
	//
	// {"PodAnnotation":"run"}
	Values *string `json:"Values,omitempty" xml:"Values,omitempty"`
}

func (s UpgradeEnvironmentFeatureRequest) String() string {
	return dara.Prettify(s)
}

func (s UpgradeEnvironmentFeatureRequest) GoString() string {
	return s.String()
}

func (s *UpgradeEnvironmentFeatureRequest) GetAliyunLang() *string {
	return s.AliyunLang
}

func (s *UpgradeEnvironmentFeatureRequest) GetEnvironmentId() *string {
	return s.EnvironmentId
}

func (s *UpgradeEnvironmentFeatureRequest) GetFeatureName() *string {
	return s.FeatureName
}

func (s *UpgradeEnvironmentFeatureRequest) GetFeatureVersion() *string {
	return s.FeatureVersion
}

func (s *UpgradeEnvironmentFeatureRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpgradeEnvironmentFeatureRequest) GetValues() *string {
	return s.Values
}

func (s *UpgradeEnvironmentFeatureRequest) SetAliyunLang(v string) *UpgradeEnvironmentFeatureRequest {
	s.AliyunLang = &v
	return s
}

func (s *UpgradeEnvironmentFeatureRequest) SetEnvironmentId(v string) *UpgradeEnvironmentFeatureRequest {
	s.EnvironmentId = &v
	return s
}

func (s *UpgradeEnvironmentFeatureRequest) SetFeatureName(v string) *UpgradeEnvironmentFeatureRequest {
	s.FeatureName = &v
	return s
}

func (s *UpgradeEnvironmentFeatureRequest) SetFeatureVersion(v string) *UpgradeEnvironmentFeatureRequest {
	s.FeatureVersion = &v
	return s
}

func (s *UpgradeEnvironmentFeatureRequest) SetRegionId(v string) *UpgradeEnvironmentFeatureRequest {
	s.RegionId = &v
	return s
}

func (s *UpgradeEnvironmentFeatureRequest) SetValues(v string) *UpgradeEnvironmentFeatureRequest {
	s.Values = &v
	return s
}

func (s *UpgradeEnvironmentFeatureRequest) Validate() error {
	return dara.Validate(s)
}
