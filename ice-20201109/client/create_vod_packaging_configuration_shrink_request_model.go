// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVodPackagingConfigurationShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigurationName(v string) *CreateVodPackagingConfigurationShrinkRequest
	GetConfigurationName() *string
	SetDescription(v string) *CreateVodPackagingConfigurationShrinkRequest
	GetDescription() *string
	SetGroupName(v string) *CreateVodPackagingConfigurationShrinkRequest
	GetGroupName() *string
	SetPackageConfigShrink(v string) *CreateVodPackagingConfigurationShrinkRequest
	GetPackageConfigShrink() *string
	SetProtocol(v string) *CreateVodPackagingConfigurationShrinkRequest
	GetProtocol() *string
}

type CreateVodPackagingConfigurationShrinkRequest struct {
	// The name of the packaging configuration. The name must be unique in an account and can be up to 128 characters in length. Letters, digits, underscores (_), and hyphens (-) are supported.
	//
	// example:
	//
	// hls_3s
	ConfigurationName *string `json:"ConfigurationName,omitempty" xml:"ConfigurationName,omitempty"`
	// The description of the packaging configuration.
	//
	// example:
	//
	// HLS 3s vod packaging
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the packaging group. The name can be up to 128 characters in length. Letters, digits, underscores (_), and hyphens (-) are supported.
	//
	// example:
	//
	// vod_hls
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The packaging configuration.
	PackageConfigShrink *string `json:"PackageConfig,omitempty" xml:"PackageConfig,omitempty"`
	// The package type.
	//
	// 	- HLS: packages content into TS segments for delivery over the HLS protocol.
	//
	// 	- HLS_CMAF: packages content into CMAF segments for delivery over the HLS protocol.
	//
	// 	- DASH: packages content for delivery over the DASH protocol.
	//
	// example:
	//
	// HLS
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
}

func (s CreateVodPackagingConfigurationShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateVodPackagingConfigurationShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateVodPackagingConfigurationShrinkRequest) GetConfigurationName() *string {
	return s.ConfigurationName
}

func (s *CreateVodPackagingConfigurationShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateVodPackagingConfigurationShrinkRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *CreateVodPackagingConfigurationShrinkRequest) GetPackageConfigShrink() *string {
	return s.PackageConfigShrink
}

func (s *CreateVodPackagingConfigurationShrinkRequest) GetProtocol() *string {
	return s.Protocol
}

func (s *CreateVodPackagingConfigurationShrinkRequest) SetConfigurationName(v string) *CreateVodPackagingConfigurationShrinkRequest {
	s.ConfigurationName = &v
	return s
}

func (s *CreateVodPackagingConfigurationShrinkRequest) SetDescription(v string) *CreateVodPackagingConfigurationShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateVodPackagingConfigurationShrinkRequest) SetGroupName(v string) *CreateVodPackagingConfigurationShrinkRequest {
	s.GroupName = &v
	return s
}

func (s *CreateVodPackagingConfigurationShrinkRequest) SetPackageConfigShrink(v string) *CreateVodPackagingConfigurationShrinkRequest {
	s.PackageConfigShrink = &v
	return s
}

func (s *CreateVodPackagingConfigurationShrinkRequest) SetProtocol(v string) *CreateVodPackagingConfigurationShrinkRequest {
	s.Protocol = &v
	return s
}

func (s *CreateVodPackagingConfigurationShrinkRequest) Validate() error {
	return dara.Validate(s)
}
