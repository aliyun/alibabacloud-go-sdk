// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAgenticDBBranchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBranchName(v string) *CreateAgenticDBBranchRequest
	GetBranchName() *string
	SetDBClusterId(v string) *CreateAgenticDBBranchRequest
	GetDBClusterId() *string
	SetDescription(v string) *CreateAgenticDBBranchRequest
	GetDescription() *string
	SetParentBranchId(v string) *CreateAgenticDBBranchRequest
	GetParentBranchId() *string
	SetProjectId(v string) *CreateAgenticDBBranchRequest
	GetProjectId() *string
	SetRegionId(v string) *CreateAgenticDBBranchRequest
	GetRegionId() *string
	SetTenantId(v string) *CreateAgenticDBBranchRequest
	GetTenantId() *string
}

type CreateAgenticDBBranchRequest struct {
	// The name of the branch.
	//
	// This parameter is required.
	//
	// example:
	//
	// feature-analytics
	BranchName *string `json:"BranchName,omitempty" xml:"BranchName,omitempty"`
	// The AgenticDB cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pagc-bp1abcdef1234567
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The description of the branch.
	//
	// example:
	//
	// Feature branch for analytics
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the parent branch. If this parameter is not specified, the branch is derived from the main branch by default.
	//
	// example:
	//
	// br-1a2b3c4d5e6f
	ParentBranchId *string `json:"ParentBranchId,omitempty" xml:"ParentBranchId,omitempty"`
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

func (s CreateAgenticDBBranchRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAgenticDBBranchRequest) GoString() string {
	return s.String()
}

func (s *CreateAgenticDBBranchRequest) GetBranchName() *string {
	return s.BranchName
}

func (s *CreateAgenticDBBranchRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *CreateAgenticDBBranchRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateAgenticDBBranchRequest) GetParentBranchId() *string {
	return s.ParentBranchId
}

func (s *CreateAgenticDBBranchRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *CreateAgenticDBBranchRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateAgenticDBBranchRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateAgenticDBBranchRequest) SetBranchName(v string) *CreateAgenticDBBranchRequest {
	s.BranchName = &v
	return s
}

func (s *CreateAgenticDBBranchRequest) SetDBClusterId(v string) *CreateAgenticDBBranchRequest {
	s.DBClusterId = &v
	return s
}

func (s *CreateAgenticDBBranchRequest) SetDescription(v string) *CreateAgenticDBBranchRequest {
	s.Description = &v
	return s
}

func (s *CreateAgenticDBBranchRequest) SetParentBranchId(v string) *CreateAgenticDBBranchRequest {
	s.ParentBranchId = &v
	return s
}

func (s *CreateAgenticDBBranchRequest) SetProjectId(v string) *CreateAgenticDBBranchRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateAgenticDBBranchRequest) SetRegionId(v string) *CreateAgenticDBBranchRequest {
	s.RegionId = &v
	return s
}

func (s *CreateAgenticDBBranchRequest) SetTenantId(v string) *CreateAgenticDBBranchRequest {
	s.TenantId = &v
	return s
}

func (s *CreateAgenticDBBranchRequest) Validate() error {
	return dara.Validate(s)
}
