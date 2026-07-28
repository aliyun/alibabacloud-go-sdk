// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetEmbodiedAIPlatformPasswordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *ResetEmbodiedAIPlatformPasswordRequest
	GetDBClusterId() *string
	SetPassword(v string) *ResetEmbodiedAIPlatformPasswordRequest
	GetPassword() *string
	SetPlatformName(v string) *ResetEmbodiedAIPlatformPasswordRequest
	GetPlatformName() *string
	SetRegionId(v string) *ResetEmbodiedAIPlatformPasswordRequest
	GetRegionId() *string
}

type ResetEmbodiedAIPlatformPasswordRequest struct {
	// The cluster ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-bp11q28kvl688****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The new password for the admin user of the embodied intelligence platform.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123*******
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The name of the embodied intelligence multimodal data platform.
	//
	// This parameter is required.
	//
	// example:
	//
	// platform1
	PlatformName *string `json:"PlatformName,omitempty" xml:"PlatformName,omitempty"`
	// The region ID of the instance.
	//
	// > You can call the DescribeRegions operation to query the region ID of a specified Data Lakehouse Edition cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ResetEmbodiedAIPlatformPasswordRequest) String() string {
	return dara.Prettify(s)
}

func (s ResetEmbodiedAIPlatformPasswordRequest) GoString() string {
	return s.String()
}

func (s *ResetEmbodiedAIPlatformPasswordRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ResetEmbodiedAIPlatformPasswordRequest) GetPassword() *string {
	return s.Password
}

func (s *ResetEmbodiedAIPlatformPasswordRequest) GetPlatformName() *string {
	return s.PlatformName
}

func (s *ResetEmbodiedAIPlatformPasswordRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ResetEmbodiedAIPlatformPasswordRequest) SetDBClusterId(v string) *ResetEmbodiedAIPlatformPasswordRequest {
	s.DBClusterId = &v
	return s
}

func (s *ResetEmbodiedAIPlatformPasswordRequest) SetPassword(v string) *ResetEmbodiedAIPlatformPasswordRequest {
	s.Password = &v
	return s
}

func (s *ResetEmbodiedAIPlatformPasswordRequest) SetPlatformName(v string) *ResetEmbodiedAIPlatformPasswordRequest {
	s.PlatformName = &v
	return s
}

func (s *ResetEmbodiedAIPlatformPasswordRequest) SetRegionId(v string) *ResetEmbodiedAIPlatformPasswordRequest {
	s.RegionId = &v
	return s
}

func (s *ResetEmbodiedAIPlatformPasswordRequest) Validate() error {
	return dara.Validate(s)
}
