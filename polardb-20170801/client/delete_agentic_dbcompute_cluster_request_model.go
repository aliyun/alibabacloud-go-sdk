// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAgenticDBComputeClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBranchId(v string) *DeleteAgenticDBComputeClusterRequest
	GetBranchId() *string
	SetComputeClusterId(v string) *DeleteAgenticDBComputeClusterRequest
	GetComputeClusterId() *string
	SetDBClusterId(v string) *DeleteAgenticDBComputeClusterRequest
	GetDBClusterId() *string
	SetProjectId(v string) *DeleteAgenticDBComputeClusterRequest
	GetProjectId() *string
	SetRegionId(v string) *DeleteAgenticDBComputeClusterRequest
	GetRegionId() *string
	SetTenantId(v string) *DeleteAgenticDBComputeClusterRequest
	GetTenantId() *string
}

type DeleteAgenticDBComputeClusterRequest struct {
	// The branch ID.
	//
	// example:
	//
	// br-69f762b1a44f49c487d64b9e71
	BranchId *string `json:"BranchId,omitempty" xml:"BranchId,omitempty"`
	// The branch compute cluster ID.
	//
	// example:
	//
	// pc-g0lsayq8c5qe
	ComputeClusterId *string `json:"ComputeClusterId,omitempty" xml:"ComputeClusterId,omitempty"`
	// The AgenticDB cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pagc-bp1abcdef1234567
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The project ID to which the resource belongs.
	//
	// example:
	//
	// proj-7140b4c74b3a44978c825bac77
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// t-51121616fa9e43e98cc90e4afa
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s DeleteAgenticDBComputeClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAgenticDBComputeClusterRequest) GoString() string {
	return s.String()
}

func (s *DeleteAgenticDBComputeClusterRequest) GetBranchId() *string {
	return s.BranchId
}

func (s *DeleteAgenticDBComputeClusterRequest) GetComputeClusterId() *string {
	return s.ComputeClusterId
}

func (s *DeleteAgenticDBComputeClusterRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DeleteAgenticDBComputeClusterRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *DeleteAgenticDBComputeClusterRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteAgenticDBComputeClusterRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *DeleteAgenticDBComputeClusterRequest) SetBranchId(v string) *DeleteAgenticDBComputeClusterRequest {
	s.BranchId = &v
	return s
}

func (s *DeleteAgenticDBComputeClusterRequest) SetComputeClusterId(v string) *DeleteAgenticDBComputeClusterRequest {
	s.ComputeClusterId = &v
	return s
}

func (s *DeleteAgenticDBComputeClusterRequest) SetDBClusterId(v string) *DeleteAgenticDBComputeClusterRequest {
	s.DBClusterId = &v
	return s
}

func (s *DeleteAgenticDBComputeClusterRequest) SetProjectId(v string) *DeleteAgenticDBComputeClusterRequest {
	s.ProjectId = &v
	return s
}

func (s *DeleteAgenticDBComputeClusterRequest) SetRegionId(v string) *DeleteAgenticDBComputeClusterRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteAgenticDBComputeClusterRequest) SetTenantId(v string) *DeleteAgenticDBComputeClusterRequest {
	s.TenantId = &v
	return s
}

func (s *DeleteAgenticDBComputeClusterRequest) Validate() error {
	return dara.Validate(s)
}
