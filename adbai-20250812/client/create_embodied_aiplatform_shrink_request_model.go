// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEmbodiedAIPlatformShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *CreateEmbodiedAIPlatformShrinkRequest
	GetDBClusterId() *string
	SetDeviceCount(v int32) *CreateEmbodiedAIPlatformShrinkRequest
	GetDeviceCount() *int32
	SetPlatformName(v string) *CreateEmbodiedAIPlatformShrinkRequest
	GetPlatformName() *string
	SetRayConfigShrink(v string) *CreateEmbodiedAIPlatformShrinkRequest
	GetRayConfigShrink() *string
	SetRayTrainConfigShrink(v string) *CreateEmbodiedAIPlatformShrinkRequest
	GetRayTrainConfigShrink() *string
	SetRegionId(v string) *CreateEmbodiedAIPlatformShrinkRequest
	GetRegionId() *string
	SetWebserverSpecName(v string) *CreateEmbodiedAIPlatformShrinkRequest
	GetWebserverSpecName() *string
}

type CreateEmbodiedAIPlatformShrinkRequest struct {
	// The instance cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-bp11q28kvl688****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The ontology count.
	//
	// example:
	//
	// 3
	DeviceCount *int32 `json:"DeviceCount,omitempty" xml:"DeviceCount,omitempty"`
	// The name of the embodied intelligence multimodal data platform.
	//
	// > The name can contain lowercase letters, digits, and underscores. It must start with a letter and end with a letter or digit. The name can be up to 16 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// platform1
	PlatformName *string `json:"PlatformName,omitempty" xml:"PlatformName,omitempty"`
	// The Ray specification information of the platform.
	RayConfigShrink *string `json:"RayConfig,omitempty" xml:"RayConfig,omitempty"`
	// The development and training resource configuration.
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
	// The webserver specification of the platform.
	//
	// example:
	//
	// large
	WebserverSpecName *string `json:"WebserverSpecName,omitempty" xml:"WebserverSpecName,omitempty"`
}

func (s CreateEmbodiedAIPlatformShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateEmbodiedAIPlatformShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateEmbodiedAIPlatformShrinkRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *CreateEmbodiedAIPlatformShrinkRequest) GetDeviceCount() *int32 {
	return s.DeviceCount
}

func (s *CreateEmbodiedAIPlatformShrinkRequest) GetPlatformName() *string {
	return s.PlatformName
}

func (s *CreateEmbodiedAIPlatformShrinkRequest) GetRayConfigShrink() *string {
	return s.RayConfigShrink
}

func (s *CreateEmbodiedAIPlatformShrinkRequest) GetRayTrainConfigShrink() *string {
	return s.RayTrainConfigShrink
}

func (s *CreateEmbodiedAIPlatformShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateEmbodiedAIPlatformShrinkRequest) GetWebserverSpecName() *string {
	return s.WebserverSpecName
}

func (s *CreateEmbodiedAIPlatformShrinkRequest) SetDBClusterId(v string) *CreateEmbodiedAIPlatformShrinkRequest {
	s.DBClusterId = &v
	return s
}

func (s *CreateEmbodiedAIPlatformShrinkRequest) SetDeviceCount(v int32) *CreateEmbodiedAIPlatformShrinkRequest {
	s.DeviceCount = &v
	return s
}

func (s *CreateEmbodiedAIPlatformShrinkRequest) SetPlatformName(v string) *CreateEmbodiedAIPlatformShrinkRequest {
	s.PlatformName = &v
	return s
}

func (s *CreateEmbodiedAIPlatformShrinkRequest) SetRayConfigShrink(v string) *CreateEmbodiedAIPlatformShrinkRequest {
	s.RayConfigShrink = &v
	return s
}

func (s *CreateEmbodiedAIPlatformShrinkRequest) SetRayTrainConfigShrink(v string) *CreateEmbodiedAIPlatformShrinkRequest {
	s.RayTrainConfigShrink = &v
	return s
}

func (s *CreateEmbodiedAIPlatformShrinkRequest) SetRegionId(v string) *CreateEmbodiedAIPlatformShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *CreateEmbodiedAIPlatformShrinkRequest) SetWebserverSpecName(v string) *CreateEmbodiedAIPlatformShrinkRequest {
	s.WebserverSpecName = &v
	return s
}

func (s *CreateEmbodiedAIPlatformShrinkRequest) Validate() error {
	return dara.Validate(s)
}
