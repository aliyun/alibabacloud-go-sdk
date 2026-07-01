// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAgenticDBBranchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBranchId(v string) *DeleteAgenticDBBranchRequest
	GetBranchId() *string
	SetDBClusterId(v string) *DeleteAgenticDBBranchRequest
	GetDBClusterId() *string
	SetProjectId(v string) *DeleteAgenticDBBranchRequest
	GetProjectId() *string
	SetRegionId(v string) *DeleteAgenticDBBranchRequest
	GetRegionId() *string
	SetTenantId(v string) *DeleteAgenticDBBranchRequest
	GetTenantId() *string
}

type DeleteAgenticDBBranchRequest struct {
	// The branch ID.
	//
	// This parameter is required.
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
	// The ID of the project to which the branch belongs.
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
	// The ID of the tenant to which the branch belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// t-4b83e0da66674951
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s DeleteAgenticDBBranchRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAgenticDBBranchRequest) GoString() string {
	return s.String()
}

func (s *DeleteAgenticDBBranchRequest) GetBranchId() *string {
	return s.BranchId
}

func (s *DeleteAgenticDBBranchRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DeleteAgenticDBBranchRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *DeleteAgenticDBBranchRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteAgenticDBBranchRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *DeleteAgenticDBBranchRequest) SetBranchId(v string) *DeleteAgenticDBBranchRequest {
	s.BranchId = &v
	return s
}

func (s *DeleteAgenticDBBranchRequest) SetDBClusterId(v string) *DeleteAgenticDBBranchRequest {
	s.DBClusterId = &v
	return s
}

func (s *DeleteAgenticDBBranchRequest) SetProjectId(v string) *DeleteAgenticDBBranchRequest {
	s.ProjectId = &v
	return s
}

func (s *DeleteAgenticDBBranchRequest) SetRegionId(v string) *DeleteAgenticDBBranchRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteAgenticDBBranchRequest) SetTenantId(v string) *DeleteAgenticDBBranchRequest {
	s.TenantId = &v
	return s
}

func (s *DeleteAgenticDBBranchRequest) Validate() error {
	return dara.Validate(s)
}
