// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteServerlessClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DeleteServerlessClusterRequest
	GetClusterId() *string
	SetRegionId(v string) *DeleteServerlessClusterRequest
	GetRegionId() *string
	SetZoneId(v string) *DeleteServerlessClusterRequest
	GetZoneId() *string
}

type DeleteServerlessClusterRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// sh-bp1pj13wh9****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-shenzhen
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-shenzhen-e
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DeleteServerlessClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteServerlessClusterRequest) GoString() string {
	return s.String()
}

func (s *DeleteServerlessClusterRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DeleteServerlessClusterRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteServerlessClusterRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *DeleteServerlessClusterRequest) SetClusterId(v string) *DeleteServerlessClusterRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteServerlessClusterRequest) SetRegionId(v string) *DeleteServerlessClusterRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteServerlessClusterRequest) SetZoneId(v string) *DeleteServerlessClusterRequest {
	s.ZoneId = &v
	return s
}

func (s *DeleteServerlessClusterRequest) Validate() error {
	return dara.Validate(s)
}
