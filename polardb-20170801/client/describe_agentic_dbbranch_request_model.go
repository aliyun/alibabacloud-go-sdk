// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAgenticDBBranchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBranchId(v string) *DescribeAgenticDBBranchRequest
	GetBranchId() *string
	SetDBClusterId(v string) *DescribeAgenticDBBranchRequest
	GetDBClusterId() *string
	SetProjectId(v string) *DescribeAgenticDBBranchRequest
	GetProjectId() *string
	SetRegionId(v string) *DescribeAgenticDBBranchRequest
	GetRegionId() *string
	SetTenantId(v string) *DescribeAgenticDBBranchRequest
	GetTenantId() *string
}

type DescribeAgenticDBBranchRequest struct {
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

func (s DescribeAgenticDBBranchRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAgenticDBBranchRequest) GoString() string {
	return s.String()
}

func (s *DescribeAgenticDBBranchRequest) GetBranchId() *string {
	return s.BranchId
}

func (s *DescribeAgenticDBBranchRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeAgenticDBBranchRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *DescribeAgenticDBBranchRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAgenticDBBranchRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *DescribeAgenticDBBranchRequest) SetBranchId(v string) *DescribeAgenticDBBranchRequest {
	s.BranchId = &v
	return s
}

func (s *DescribeAgenticDBBranchRequest) SetDBClusterId(v string) *DescribeAgenticDBBranchRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeAgenticDBBranchRequest) SetProjectId(v string) *DescribeAgenticDBBranchRequest {
	s.ProjectId = &v
	return s
}

func (s *DescribeAgenticDBBranchRequest) SetRegionId(v string) *DescribeAgenticDBBranchRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAgenticDBBranchRequest) SetTenantId(v string) *DescribeAgenticDBBranchRequest {
	s.TenantId = &v
	return s
}

func (s *DescribeAgenticDBBranchRequest) Validate() error {
	return dara.Validate(s)
}
