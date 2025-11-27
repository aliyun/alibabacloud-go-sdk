// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRCVClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DeleteRCVClusterRequest
	GetClusterId() *string
	SetRegionId(v string) *DeleteRCVClusterRequest
	GetRegionId() *string
}

type DeleteRCVClusterRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cd21387ea640145bab79a78276c1a****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteRCVClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteRCVClusterRequest) GoString() string {
	return s.String()
}

func (s *DeleteRCVClusterRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DeleteRCVClusterRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteRCVClusterRequest) SetClusterId(v string) *DeleteRCVClusterRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteRCVClusterRequest) SetRegionId(v string) *DeleteRCVClusterRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteRCVClusterRequest) Validate() error {
	return dara.Validate(s)
}
