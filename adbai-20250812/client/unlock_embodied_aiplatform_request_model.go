// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnlockEmbodiedAIPlatformRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *UnlockEmbodiedAIPlatformRequest
	GetDBClusterId() *string
	SetPlatformName(v string) *UnlockEmbodiedAIPlatformRequest
	GetPlatformName() *string
	SetRegionId(v string) *UnlockEmbodiedAIPlatformRequest
	GetRegionId() *string
}

type UnlockEmbodiedAIPlatformRequest struct {
	// The instance cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-bp11q28kvl688****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The name of the embodied intelligence multimodal data platform.
	//
	// > The name must start with a lowercase letter and end with a lowercase letter or digit. It can contain lowercase letters, digits, and underscores. The name can be up to 16 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// platform1
	PlatformName *string `json:"PlatformName,omitempty" xml:"PlatformName,omitempty"`
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
}

func (s UnlockEmbodiedAIPlatformRequest) String() string {
	return dara.Prettify(s)
}

func (s UnlockEmbodiedAIPlatformRequest) GoString() string {
	return s.String()
}

func (s *UnlockEmbodiedAIPlatformRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *UnlockEmbodiedAIPlatformRequest) GetPlatformName() *string {
	return s.PlatformName
}

func (s *UnlockEmbodiedAIPlatformRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UnlockEmbodiedAIPlatformRequest) SetDBClusterId(v string) *UnlockEmbodiedAIPlatformRequest {
	s.DBClusterId = &v
	return s
}

func (s *UnlockEmbodiedAIPlatformRequest) SetPlatformName(v string) *UnlockEmbodiedAIPlatformRequest {
	s.PlatformName = &v
	return s
}

func (s *UnlockEmbodiedAIPlatformRequest) SetRegionId(v string) *UnlockEmbodiedAIPlatformRequest {
	s.RegionId = &v
	return s
}

func (s *UnlockEmbodiedAIPlatformRequest) Validate() error {
	return dara.Validate(s)
}
