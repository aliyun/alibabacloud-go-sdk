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
	PlatformName *string `json:"PlatformName,omitempty" xml:"PlatformName,omitempty"`
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
