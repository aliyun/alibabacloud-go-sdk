// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetListGpuClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *GetListGpuClusterRequest
	GetRegionId() *string
	SetTimeRegion(v interface{}) *GetListGpuClusterRequest
	GetTimeRegion() interface{}
}

type GetListGpuClusterRequest struct {
	// The ID of the region to query.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The query time range, in milliseconds, in the format of [start, end]. Example: [13254125000, 52451241000].
	//
	// example:
	//
	// [1757820000, 1757823600]
	TimeRegion interface{} `json:"timeRegion,omitempty" xml:"timeRegion,omitempty"`
}

func (s GetListGpuClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s GetListGpuClusterRequest) GoString() string {
	return s.String()
}

func (s *GetListGpuClusterRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetListGpuClusterRequest) GetTimeRegion() interface{} {
	return s.TimeRegion
}

func (s *GetListGpuClusterRequest) SetRegionId(v string) *GetListGpuClusterRequest {
	s.RegionId = &v
	return s
}

func (s *GetListGpuClusterRequest) SetTimeRegion(v interface{}) *GetListGpuClusterRequest {
	s.TimeRegion = v
	return s
}

func (s *GetListGpuClusterRequest) Validate() error {
	return dara.Validate(s)
}
