// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAgenticDBProjectsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeAgenticDBProjectsRequest
	GetDBClusterId() *string
	SetPageNumber(v int32) *DescribeAgenticDBProjectsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeAgenticDBProjectsRequest
	GetPageSize() *int32
	SetProjectId(v string) *DescribeAgenticDBProjectsRequest
	GetProjectId() *string
	SetProjectName(v string) *DescribeAgenticDBProjectsRequest
	GetProjectName() *string
	SetRegionId(v string) *DescribeAgenticDBProjectsRequest
	GetRegionId() *string
	SetTenantId(v string) *DescribeAgenticDBProjectsRequest
	GetTenantId() *string
}

type DescribeAgenticDBProjectsRequest struct {
	// The AgenticDB cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pagc-bp1abcdef1234567
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 30. Maximum value: 100.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The project ID for exact match.
	//
	// example:
	//
	// proj-a1b2c3d4e5f6
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The project name for fuzzy match.
	//
	// example:
	//
	// analytics
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The tenant ID to which the project belongs.
	//
	// example:
	//
	// t-4b83e0da66674951
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s DescribeAgenticDBProjectsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAgenticDBProjectsRequest) GoString() string {
	return s.String()
}

func (s *DescribeAgenticDBProjectsRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeAgenticDBProjectsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeAgenticDBProjectsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeAgenticDBProjectsRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *DescribeAgenticDBProjectsRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *DescribeAgenticDBProjectsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAgenticDBProjectsRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *DescribeAgenticDBProjectsRequest) SetDBClusterId(v string) *DescribeAgenticDBProjectsRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeAgenticDBProjectsRequest) SetPageNumber(v int32) *DescribeAgenticDBProjectsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeAgenticDBProjectsRequest) SetPageSize(v int32) *DescribeAgenticDBProjectsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAgenticDBProjectsRequest) SetProjectId(v string) *DescribeAgenticDBProjectsRequest {
	s.ProjectId = &v
	return s
}

func (s *DescribeAgenticDBProjectsRequest) SetProjectName(v string) *DescribeAgenticDBProjectsRequest {
	s.ProjectName = &v
	return s
}

func (s *DescribeAgenticDBProjectsRequest) SetRegionId(v string) *DescribeAgenticDBProjectsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAgenticDBProjectsRequest) SetTenantId(v string) *DescribeAgenticDBProjectsRequest {
	s.TenantId = &v
	return s
}

func (s *DescribeAgenticDBProjectsRequest) Validate() error {
	return dara.Validate(s)
}
