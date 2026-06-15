// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopCacheClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *StopCacheClusterRequest
	GetRegionId() *string
}

type StopCacheClusterRequest struct {
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s StopCacheClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s StopCacheClusterRequest) GoString() string {
	return s.String()
}

func (s *StopCacheClusterRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *StopCacheClusterRequest) SetRegionId(v string) *StopCacheClusterRequest {
	s.RegionId = &v
	return s
}

func (s *StopCacheClusterRequest) Validate() error {
	return dara.Validate(s)
}
