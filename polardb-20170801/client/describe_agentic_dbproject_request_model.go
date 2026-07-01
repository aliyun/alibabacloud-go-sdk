// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAgenticDBProjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeAgenticDBProjectRequest
	GetDBClusterId() *string
	SetProjectId(v string) *DescribeAgenticDBProjectRequest
	GetProjectId() *string
	SetRegionId(v string) *DescribeAgenticDBProjectRequest
	GetRegionId() *string
	SetTenantId(v string) *DescribeAgenticDBProjectRequest
	GetTenantId() *string
}

type DescribeAgenticDBProjectRequest struct {
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

func (s DescribeAgenticDBProjectRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAgenticDBProjectRequest) GoString() string {
	return s.String()
}

func (s *DescribeAgenticDBProjectRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeAgenticDBProjectRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *DescribeAgenticDBProjectRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAgenticDBProjectRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *DescribeAgenticDBProjectRequest) SetDBClusterId(v string) *DescribeAgenticDBProjectRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeAgenticDBProjectRequest) SetProjectId(v string) *DescribeAgenticDBProjectRequest {
	s.ProjectId = &v
	return s
}

func (s *DescribeAgenticDBProjectRequest) SetRegionId(v string) *DescribeAgenticDBProjectRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAgenticDBProjectRequest) SetTenantId(v string) *DescribeAgenticDBProjectRequest {
	s.TenantId = &v
	return s
}

func (s *DescribeAgenticDBProjectRequest) Validate() error {
	return dara.Validate(s)
}
