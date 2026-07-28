// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEmbodiedAIPlatformRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DeleteEmbodiedAIPlatformRequest
	GetDBClusterId() *string
	SetPlatformName(v string) *DeleteEmbodiedAIPlatformRequest
	GetPlatformName() *string
	SetRegionId(v string) *DeleteEmbodiedAIPlatformRequest
	GetRegionId() *string
}

type DeleteEmbodiedAIPlatformRequest struct {
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
	// > The name can contain lowercase letters, digits, and underscores. It must start with a letter and end with a letter or digit. The name can be up to 16 characters in length.
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

func (s DeleteEmbodiedAIPlatformRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteEmbodiedAIPlatformRequest) GoString() string {
	return s.String()
}

func (s *DeleteEmbodiedAIPlatformRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DeleteEmbodiedAIPlatformRequest) GetPlatformName() *string {
	return s.PlatformName
}

func (s *DeleteEmbodiedAIPlatformRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteEmbodiedAIPlatformRequest) SetDBClusterId(v string) *DeleteEmbodiedAIPlatformRequest {
	s.DBClusterId = &v
	return s
}

func (s *DeleteEmbodiedAIPlatformRequest) SetPlatformName(v string) *DeleteEmbodiedAIPlatformRequest {
	s.PlatformName = &v
	return s
}

func (s *DeleteEmbodiedAIPlatformRequest) SetRegionId(v string) *DeleteEmbodiedAIPlatformRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteEmbodiedAIPlatformRequest) Validate() error {
	return dara.Validate(s)
}
