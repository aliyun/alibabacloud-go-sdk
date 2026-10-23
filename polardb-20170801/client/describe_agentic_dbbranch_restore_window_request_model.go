// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAgenticDBBranchRestoreWindowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBranchId(v string) *DescribeAgenticDBBranchRestoreWindowRequest
	GetBranchId() *string
	SetDBClusterId(v string) *DescribeAgenticDBBranchRestoreWindowRequest
	GetDBClusterId() *string
	SetProjectId(v string) *DescribeAgenticDBBranchRestoreWindowRequest
	GetProjectId() *string
	SetRegionId(v string) *DescribeAgenticDBBranchRestoreWindowRequest
	GetRegionId() *string
	SetTenantId(v string) *DescribeAgenticDBBranchRestoreWindowRequest
	GetTenantId() *string
}

type DescribeAgenticDBBranchRestoreWindowRequest struct {
	// The ID of the source branch.
	//
	// This parameter is required.
	//
	// example:
	//
	// br-1a2b3c4d5e6f
	BranchId *string `json:"BranchId,omitempty" xml:"BranchId,omitempty"`
	// The ID of the AgenticDB cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// pagc-bp1abcdef1234567
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The ID of the target project.
	//
	// This parameter is required.
	//
	// example:
	//
	// proj-a1b2c3d4e5f6
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
	// This parameter is required.
	//
	// example:
	//
	// t-4b83e0da66674951
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s DescribeAgenticDBBranchRestoreWindowRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAgenticDBBranchRestoreWindowRequest) GoString() string {
	return s.String()
}

func (s *DescribeAgenticDBBranchRestoreWindowRequest) GetBranchId() *string {
	return s.BranchId
}

func (s *DescribeAgenticDBBranchRestoreWindowRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeAgenticDBBranchRestoreWindowRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *DescribeAgenticDBBranchRestoreWindowRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAgenticDBBranchRestoreWindowRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *DescribeAgenticDBBranchRestoreWindowRequest) SetBranchId(v string) *DescribeAgenticDBBranchRestoreWindowRequest {
	s.BranchId = &v
	return s
}

func (s *DescribeAgenticDBBranchRestoreWindowRequest) SetDBClusterId(v string) *DescribeAgenticDBBranchRestoreWindowRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeAgenticDBBranchRestoreWindowRequest) SetProjectId(v string) *DescribeAgenticDBBranchRestoreWindowRequest {
	s.ProjectId = &v
	return s
}

func (s *DescribeAgenticDBBranchRestoreWindowRequest) SetRegionId(v string) *DescribeAgenticDBBranchRestoreWindowRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAgenticDBBranchRestoreWindowRequest) SetTenantId(v string) *DescribeAgenticDBBranchRestoreWindowRequest {
	s.TenantId = &v
	return s
}

func (s *DescribeAgenticDBBranchRestoreWindowRequest) Validate() error {
	return dara.Validate(s)
}
