// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAgenticDBComputeClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComputeClusterId(v string) *DeleteAgenticDBComputeClusterRequest
	GetComputeClusterId() *string
	SetDBClusterId(v string) *DeleteAgenticDBComputeClusterRequest
	GetDBClusterId() *string
	SetRegionId(v string) *DeleteAgenticDBComputeClusterRequest
	GetRegionId() *string
}

type DeleteAgenticDBComputeClusterRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pc-g0lsayq8c5qe
	ComputeClusterId *string `json:"ComputeClusterId,omitempty" xml:"ComputeClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pagc-bp1abcdef1234567
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteAgenticDBComputeClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAgenticDBComputeClusterRequest) GoString() string {
	return s.String()
}

func (s *DeleteAgenticDBComputeClusterRequest) GetComputeClusterId() *string {
	return s.ComputeClusterId
}

func (s *DeleteAgenticDBComputeClusterRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DeleteAgenticDBComputeClusterRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteAgenticDBComputeClusterRequest) SetComputeClusterId(v string) *DeleteAgenticDBComputeClusterRequest {
	s.ComputeClusterId = &v
	return s
}

func (s *DeleteAgenticDBComputeClusterRequest) SetDBClusterId(v string) *DeleteAgenticDBComputeClusterRequest {
	s.DBClusterId = &v
	return s
}

func (s *DeleteAgenticDBComputeClusterRequest) SetRegionId(v string) *DeleteAgenticDBComputeClusterRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteAgenticDBComputeClusterRequest) Validate() error {
	return dara.Validate(s)
}
