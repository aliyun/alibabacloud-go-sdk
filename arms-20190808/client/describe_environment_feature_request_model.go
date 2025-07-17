// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEnvironmentFeatureRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliyunLang(v string) *DescribeEnvironmentFeatureRequest
	GetAliyunLang() *string
	SetEnvironmentId(v string) *DescribeEnvironmentFeatureRequest
	GetEnvironmentId() *string
	SetFeatureName(v string) *DescribeEnvironmentFeatureRequest
	GetFeatureName() *string
	SetRegionId(v string) *DescribeEnvironmentFeatureRequest
	GetRegionId() *string
}

type DescribeEnvironmentFeatureRequest struct {
	// The language. Valid values: en and zh.
	//
	// example:
	//
	// en
	AliyunLang *string `json:"AliyunLang,omitempty" xml:"AliyunLang,omitempty"`
	// The environment ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// env-xxx
	EnvironmentId *string `json:"EnvironmentId,omitempty" xml:"EnvironmentId,omitempty"`
	// The name of the feature.
	//
	// Valid values:
	//
	// 	- app-agent-pilot: App Pilot agent
	//
	// 	- arms-cmonitor: ARMS CMonitor agent
	//
	// 	- metric-agent: Prometheus agent
	//
	// This parameter is required.
	//
	// example:
	//
	// metric-agent
	FeatureName *string `json:"FeatureName,omitempty" xml:"FeatureName,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeEnvironmentFeatureRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnvironmentFeatureRequest) GoString() string {
	return s.String()
}

func (s *DescribeEnvironmentFeatureRequest) GetAliyunLang() *string {
	return s.AliyunLang
}

func (s *DescribeEnvironmentFeatureRequest) GetEnvironmentId() *string {
	return s.EnvironmentId
}

func (s *DescribeEnvironmentFeatureRequest) GetFeatureName() *string {
	return s.FeatureName
}

func (s *DescribeEnvironmentFeatureRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeEnvironmentFeatureRequest) SetAliyunLang(v string) *DescribeEnvironmentFeatureRequest {
	s.AliyunLang = &v
	return s
}

func (s *DescribeEnvironmentFeatureRequest) SetEnvironmentId(v string) *DescribeEnvironmentFeatureRequest {
	s.EnvironmentId = &v
	return s
}

func (s *DescribeEnvironmentFeatureRequest) SetFeatureName(v string) *DescribeEnvironmentFeatureRequest {
	s.FeatureName = &v
	return s
}

func (s *DescribeEnvironmentFeatureRequest) SetRegionId(v string) *DescribeEnvironmentFeatureRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeEnvironmentFeatureRequest) Validate() error {
	return dara.Validate(s)
}
