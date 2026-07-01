// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAgenticDBBranchLineageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBranchId(v string) *DescribeAgenticDBBranchLineageRequest
	GetBranchId() *string
	SetDBClusterId(v string) *DescribeAgenticDBBranchLineageRequest
	GetDBClusterId() *string
	SetIncludeDestroying(v bool) *DescribeAgenticDBBranchLineageRequest
	GetIncludeDestroying() *bool
	SetMaxViewDepth(v int32) *DescribeAgenticDBBranchLineageRequest
	GetMaxViewDepth() *int32
	SetProjectId(v string) *DescribeAgenticDBBranchLineageRequest
	GetProjectId() *string
	SetRegionId(v string) *DescribeAgenticDBBranchLineageRequest
	GetRegionId() *string
	SetTenantId(v string) *DescribeAgenticDBBranchLineageRequest
	GetTenantId() *string
}

type DescribeAgenticDBBranchLineageRequest struct {
	// The anchor branch ID. If this parameter is not specified, the primary branch of the project is used by default.
	//
	// example:
	//
	// br-7g8h9i0j1k2l
	BranchId *string `json:"BranchId,omitempty" xml:"BranchId,omitempty"`
	// The AgenticDB cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pagc-bp1abcdef1234567
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// Specifies whether to include branches in the Destroying state. Default value: false.
	//
	// example:
	//
	// false
	IncludeDestroying *bool `json:"IncludeDestroying,omitempty" xml:"IncludeDestroying,omitempty"`
	// example:
	//
	// 3
	MaxViewDepth *int32 `json:"MaxViewDepth,omitempty" xml:"MaxViewDepth,omitempty"`
	// The project ID.
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

func (s DescribeAgenticDBBranchLineageRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAgenticDBBranchLineageRequest) GoString() string {
	return s.String()
}

func (s *DescribeAgenticDBBranchLineageRequest) GetBranchId() *string {
	return s.BranchId
}

func (s *DescribeAgenticDBBranchLineageRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeAgenticDBBranchLineageRequest) GetIncludeDestroying() *bool {
	return s.IncludeDestroying
}

func (s *DescribeAgenticDBBranchLineageRequest) GetMaxViewDepth() *int32 {
	return s.MaxViewDepth
}

func (s *DescribeAgenticDBBranchLineageRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *DescribeAgenticDBBranchLineageRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAgenticDBBranchLineageRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *DescribeAgenticDBBranchLineageRequest) SetBranchId(v string) *DescribeAgenticDBBranchLineageRequest {
	s.BranchId = &v
	return s
}

func (s *DescribeAgenticDBBranchLineageRequest) SetDBClusterId(v string) *DescribeAgenticDBBranchLineageRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeAgenticDBBranchLineageRequest) SetIncludeDestroying(v bool) *DescribeAgenticDBBranchLineageRequest {
	s.IncludeDestroying = &v
	return s
}

func (s *DescribeAgenticDBBranchLineageRequest) SetMaxViewDepth(v int32) *DescribeAgenticDBBranchLineageRequest {
	s.MaxViewDepth = &v
	return s
}

func (s *DescribeAgenticDBBranchLineageRequest) SetProjectId(v string) *DescribeAgenticDBBranchLineageRequest {
	s.ProjectId = &v
	return s
}

func (s *DescribeAgenticDBBranchLineageRequest) SetRegionId(v string) *DescribeAgenticDBBranchLineageRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAgenticDBBranchLineageRequest) SetTenantId(v string) *DescribeAgenticDBBranchLineageRequest {
	s.TenantId = &v
	return s
}

func (s *DescribeAgenticDBBranchLineageRequest) Validate() error {
	return dara.Validate(s)
}
