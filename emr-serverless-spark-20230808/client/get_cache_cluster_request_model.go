// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCacheClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *GetCacheClusterRequest
	GetRegionId() *string
}

type GetCacheClusterRequest struct {
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s GetCacheClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCacheClusterRequest) GoString() string {
	return s.String()
}

func (s *GetCacheClusterRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetCacheClusterRequest) SetRegionId(v string) *GetCacheClusterRequest {
	s.RegionId = &v
	return s
}

func (s *GetCacheClusterRequest) Validate() error {
	return dara.Validate(s)
}
