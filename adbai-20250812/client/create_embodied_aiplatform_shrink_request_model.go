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
	SetPlatformName(v string) *CreateEmbodiedAIPlatformShrinkRequest
	GetPlatformName() *string
	SetRayConfigShrink(v string) *CreateEmbodiedAIPlatformShrinkRequest
	GetRayConfigShrink() *string
	SetRegionId(v string) *CreateEmbodiedAIPlatformShrinkRequest
	GetRegionId() *string
	SetWebserverSpecName(v string) *CreateEmbodiedAIPlatformShrinkRequest
	GetWebserverSpecName() *string
}

type CreateEmbodiedAIPlatformShrinkRequest struct {
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
	// platform1
	PlatformName    *string `json:"PlatformName,omitempty" xml:"PlatformName,omitempty"`
	RayConfigShrink *string `json:"RayConfig,omitempty" xml:"RayConfig,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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

func (s *CreateEmbodiedAIPlatformShrinkRequest) GetPlatformName() *string {
	return s.PlatformName
}

func (s *CreateEmbodiedAIPlatformShrinkRequest) GetRayConfigShrink() *string {
	return s.RayConfigShrink
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

func (s *CreateEmbodiedAIPlatformShrinkRequest) SetPlatformName(v string) *CreateEmbodiedAIPlatformShrinkRequest {
	s.PlatformName = &v
	return s
}

func (s *CreateEmbodiedAIPlatformShrinkRequest) SetRayConfigShrink(v string) *CreateEmbodiedAIPlatformShrinkRequest {
	s.RayConfigShrink = &v
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
