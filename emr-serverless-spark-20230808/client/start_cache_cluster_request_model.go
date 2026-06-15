// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartCacheClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *StartCacheClusterRequest
	GetRegionId() *string
}

type StartCacheClusterRequest struct {
	// Region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s StartCacheClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s StartCacheClusterRequest) GoString() string {
	return s.String()
}

func (s *StartCacheClusterRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *StartCacheClusterRequest) SetRegionId(v string) *StartCacheClusterRequest {
	s.RegionId = &v
	return s
}

func (s *StartCacheClusterRequest) Validate() error {
	return dara.Validate(s)
}
