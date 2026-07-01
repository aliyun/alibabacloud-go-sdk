// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAgenticDBBranchesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBranchName(v string) *DescribeAgenticDBBranchesRequest
	GetBranchName() *string
	SetDBClusterId(v string) *DescribeAgenticDBBranchesRequest
	GetDBClusterId() *string
	SetPageNumber(v int32) *DescribeAgenticDBBranchesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeAgenticDBBranchesRequest
	GetPageSize() *int32
	SetProjectId(v string) *DescribeAgenticDBBranchesRequest
	GetProjectId() *string
	SetRegionId(v string) *DescribeAgenticDBBranchesRequest
	GetRegionId() *string
	SetStatus(v string) *DescribeAgenticDBBranchesRequest
	GetStatus() *string
	SetTenantId(v string) *DescribeAgenticDBBranchesRequest
	GetTenantId() *string
}

type DescribeAgenticDBBranchesRequest struct {
	// example:
	//
	// feature
	BranchName *string `json:"BranchName,omitempty" xml:"BranchName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pagc-bp1abcdef1234567
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
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
	// example:
	//
	// Active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// t-4b83e0da66674951
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s DescribeAgenticDBBranchesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAgenticDBBranchesRequest) GoString() string {
	return s.String()
}

func (s *DescribeAgenticDBBranchesRequest) GetBranchName() *string {
	return s.BranchName
}

func (s *DescribeAgenticDBBranchesRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeAgenticDBBranchesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeAgenticDBBranchesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeAgenticDBBranchesRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *DescribeAgenticDBBranchesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAgenticDBBranchesRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeAgenticDBBranchesRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *DescribeAgenticDBBranchesRequest) SetBranchName(v string) *DescribeAgenticDBBranchesRequest {
	s.BranchName = &v
	return s
}

func (s *DescribeAgenticDBBranchesRequest) SetDBClusterId(v string) *DescribeAgenticDBBranchesRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeAgenticDBBranchesRequest) SetPageNumber(v int32) *DescribeAgenticDBBranchesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeAgenticDBBranchesRequest) SetPageSize(v int32) *DescribeAgenticDBBranchesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAgenticDBBranchesRequest) SetProjectId(v string) *DescribeAgenticDBBranchesRequest {
	s.ProjectId = &v
	return s
}

func (s *DescribeAgenticDBBranchesRequest) SetRegionId(v string) *DescribeAgenticDBBranchesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAgenticDBBranchesRequest) SetStatus(v string) *DescribeAgenticDBBranchesRequest {
	s.Status = &v
	return s
}

func (s *DescribeAgenticDBBranchesRequest) SetTenantId(v string) *DescribeAgenticDBBranchesRequest {
	s.TenantId = &v
	return s
}

func (s *DescribeAgenticDBBranchesRequest) Validate() error {
	return dara.Validate(s)
}
