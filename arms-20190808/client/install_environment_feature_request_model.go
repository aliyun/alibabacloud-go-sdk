// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallEnvironmentFeatureRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliyunLang(v string) *InstallEnvironmentFeatureRequest
	GetAliyunLang() *string
	SetConfig(v string) *InstallEnvironmentFeatureRequest
	GetConfig() *string
	SetEnvironmentId(v string) *InstallEnvironmentFeatureRequest
	GetEnvironmentId() *string
	SetFeatureName(v string) *InstallEnvironmentFeatureRequest
	GetFeatureName() *string
	SetFeatureVersion(v string) *InstallEnvironmentFeatureRequest
	GetFeatureVersion() *string
	SetRegion(v string) *InstallEnvironmentFeatureRequest
	GetRegion() *string
	SetRegionId(v string) *InstallEnvironmentFeatureRequest
	GetRegionId() *string
}

type InstallEnvironmentFeatureRequest struct {
	// The language. Valid values: zh and en. Default value: zh.
	//
	// example:
	//
	// zh
	AliyunLang *string `json:"AliyunLang,omitempty" xml:"AliyunLang,omitempty"`
	// The metadata of the feature.
	//
	// example:
	//
	// {\\"continuous\\":true,\\"dataRevision\\":2}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
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
	// 	- app-agent-pilot
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- metric-agent
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// This parameter is required.
	//
	// example:
	//
	// metric-agent
	FeatureName *string `json:"FeatureName,omitempty" xml:"FeatureName,omitempty"`
	// The version of the feature.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1.1.17
	FeatureVersion *string `json:"FeatureVersion,omitempty" xml:"FeatureVersion,omitempty"`
	// The region ID of the feature.
	//
	// example:
	//
	// cn-shenzhen
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s InstallEnvironmentFeatureRequest) String() string {
	return dara.Prettify(s)
}

func (s InstallEnvironmentFeatureRequest) GoString() string {
	return s.String()
}

func (s *InstallEnvironmentFeatureRequest) GetAliyunLang() *string {
	return s.AliyunLang
}

func (s *InstallEnvironmentFeatureRequest) GetConfig() *string {
	return s.Config
}

func (s *InstallEnvironmentFeatureRequest) GetEnvironmentId() *string {
	return s.EnvironmentId
}

func (s *InstallEnvironmentFeatureRequest) GetFeatureName() *string {
	return s.FeatureName
}

func (s *InstallEnvironmentFeatureRequest) GetFeatureVersion() *string {
	return s.FeatureVersion
}

func (s *InstallEnvironmentFeatureRequest) GetRegion() *string {
	return s.Region
}

func (s *InstallEnvironmentFeatureRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *InstallEnvironmentFeatureRequest) SetAliyunLang(v string) *InstallEnvironmentFeatureRequest {
	s.AliyunLang = &v
	return s
}

func (s *InstallEnvironmentFeatureRequest) SetConfig(v string) *InstallEnvironmentFeatureRequest {
	s.Config = &v
	return s
}

func (s *InstallEnvironmentFeatureRequest) SetEnvironmentId(v string) *InstallEnvironmentFeatureRequest {
	s.EnvironmentId = &v
	return s
}

func (s *InstallEnvironmentFeatureRequest) SetFeatureName(v string) *InstallEnvironmentFeatureRequest {
	s.FeatureName = &v
	return s
}

func (s *InstallEnvironmentFeatureRequest) SetFeatureVersion(v string) *InstallEnvironmentFeatureRequest {
	s.FeatureVersion = &v
	return s
}

func (s *InstallEnvironmentFeatureRequest) SetRegion(v string) *InstallEnvironmentFeatureRequest {
	s.Region = &v
	return s
}

func (s *InstallEnvironmentFeatureRequest) SetRegionId(v string) *InstallEnvironmentFeatureRequest {
	s.RegionId = &v
	return s
}

func (s *InstallEnvironmentFeatureRequest) Validate() error {
	return dara.Validate(s)
}
