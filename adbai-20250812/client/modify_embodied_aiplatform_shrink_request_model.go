// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyEmbodiedAIPlatformShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *ModifyEmbodiedAIPlatformShrinkRequest
	GetDBClusterId() *string
	SetDeviceCount(v string) *ModifyEmbodiedAIPlatformShrinkRequest
	GetDeviceCount() *string
	SetPlatformName(v string) *ModifyEmbodiedAIPlatformShrinkRequest
	GetPlatformName() *string
	SetRayConfigShrink(v string) *ModifyEmbodiedAIPlatformShrinkRequest
	GetRayConfigShrink() *string
	SetRayTrainConfigShrink(v string) *ModifyEmbodiedAIPlatformShrinkRequest
	GetRayTrainConfigShrink() *string
	SetRegionId(v string) *ModifyEmbodiedAIPlatformShrinkRequest
	GetRegionId() *string
	SetWebserverSpecName(v string) *ModifyEmbodiedAIPlatformShrinkRequest
	GetWebserverSpecName() *string
}

type ModifyEmbodiedAIPlatformShrinkRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-bp11q28kvl688****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	DeviceCount *string `json:"DeviceCount,omitempty" xml:"DeviceCount,omitempty"`
	// The name of the embodied intelligence multimodal data platform.
	//
	// > The name can contain lowercase letters, digits, and underscores (_). It must start with a letter and end with a letter or digit. The name can be up to 16 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// eap_platform
	PlatformName *string `json:"PlatformName,omitempty" xml:"PlatformName,omitempty"`
	// The Ray specification information of the platform.
	RayConfigShrink      *string `json:"RayConfig,omitempty" xml:"RayConfig,omitempty"`
	RayTrainConfigShrink *string `json:"RayTrainConfig,omitempty" xml:"RayTrainConfig,omitempty"`
	// The region ID.
	//
	// > You can call the DescribeRegions operation to query the region ID of a specified Data Lakehouse Edition cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The Webserver specification of the platform.
	//
	// example:
	//
	// large
	WebserverSpecName *string `json:"WebserverSpecName,omitempty" xml:"WebserverSpecName,omitempty"`
}

func (s ModifyEmbodiedAIPlatformShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyEmbodiedAIPlatformShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyEmbodiedAIPlatformShrinkRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyEmbodiedAIPlatformShrinkRequest) GetDeviceCount() *string {
	return s.DeviceCount
}

func (s *ModifyEmbodiedAIPlatformShrinkRequest) GetPlatformName() *string {
	return s.PlatformName
}

func (s *ModifyEmbodiedAIPlatformShrinkRequest) GetRayConfigShrink() *string {
	return s.RayConfigShrink
}

func (s *ModifyEmbodiedAIPlatformShrinkRequest) GetRayTrainConfigShrink() *string {
	return s.RayTrainConfigShrink
}

func (s *ModifyEmbodiedAIPlatformShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyEmbodiedAIPlatformShrinkRequest) GetWebserverSpecName() *string {
	return s.WebserverSpecName
}

func (s *ModifyEmbodiedAIPlatformShrinkRequest) SetDBClusterId(v string) *ModifyEmbodiedAIPlatformShrinkRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyEmbodiedAIPlatformShrinkRequest) SetDeviceCount(v string) *ModifyEmbodiedAIPlatformShrinkRequest {
	s.DeviceCount = &v
	return s
}

func (s *ModifyEmbodiedAIPlatformShrinkRequest) SetPlatformName(v string) *ModifyEmbodiedAIPlatformShrinkRequest {
	s.PlatformName = &v
	return s
}

func (s *ModifyEmbodiedAIPlatformShrinkRequest) SetRayConfigShrink(v string) *ModifyEmbodiedAIPlatformShrinkRequest {
	s.RayConfigShrink = &v
	return s
}

func (s *ModifyEmbodiedAIPlatformShrinkRequest) SetRayTrainConfigShrink(v string) *ModifyEmbodiedAIPlatformShrinkRequest {
	s.RayTrainConfigShrink = &v
	return s
}

func (s *ModifyEmbodiedAIPlatformShrinkRequest) SetRegionId(v string) *ModifyEmbodiedAIPlatformShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyEmbodiedAIPlatformShrinkRequest) SetWebserverSpecName(v string) *ModifyEmbodiedAIPlatformShrinkRequest {
	s.WebserverSpecName = &v
	return s
}

func (s *ModifyEmbodiedAIPlatformShrinkRequest) Validate() error {
	return dara.Validate(s)
}
