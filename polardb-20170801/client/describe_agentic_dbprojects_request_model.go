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
	// example:
	//
	// analytics
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
