// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAgenticDBBranchEndpointsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBranchId(v string) *DescribeAgenticDBBranchEndpointsRequest
	GetBranchId() *string
	SetDBClusterId(v string) *DescribeAgenticDBBranchEndpointsRequest
	GetDBClusterId() *string
	SetProjectId(v string) *DescribeAgenticDBBranchEndpointsRequest
	GetProjectId() *string
	SetRegionId(v string) *DescribeAgenticDBBranchEndpointsRequest
	GetRegionId() *string
	SetTenantId(v string) *DescribeAgenticDBBranchEndpointsRequest
	GetTenantId() *string
}

type DescribeAgenticDBBranchEndpointsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// br-7g8h9i0j1k2l
	BranchId *string `json:"BranchId,omitempty" xml:"BranchId,omitempty"`
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
	// proj-a1b2c3d4e5f6
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// t-4b83e0da66674951
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s DescribeAgenticDBBranchEndpointsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAgenticDBBranchEndpointsRequest) GoString() string {
	return s.String()
}

func (s *DescribeAgenticDBBranchEndpointsRequest) GetBranchId() *string {
	return s.BranchId
}

func (s *DescribeAgenticDBBranchEndpointsRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeAgenticDBBranchEndpointsRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *DescribeAgenticDBBranchEndpointsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAgenticDBBranchEndpointsRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *DescribeAgenticDBBranchEndpointsRequest) SetBranchId(v string) *DescribeAgenticDBBranchEndpointsRequest {
	s.BranchId = &v
	return s
}

func (s *DescribeAgenticDBBranchEndpointsRequest) SetDBClusterId(v string) *DescribeAgenticDBBranchEndpointsRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeAgenticDBBranchEndpointsRequest) SetProjectId(v string) *DescribeAgenticDBBranchEndpointsRequest {
	s.ProjectId = &v
	return s
}

func (s *DescribeAgenticDBBranchEndpointsRequest) SetRegionId(v string) *DescribeAgenticDBBranchEndpointsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAgenticDBBranchEndpointsRequest) SetTenantId(v string) *DescribeAgenticDBBranchEndpointsRequest {
	s.TenantId = &v
	return s
}

func (s *DescribeAgenticDBBranchEndpointsRequest) Validate() error {
	return dara.Validate(s)
}
