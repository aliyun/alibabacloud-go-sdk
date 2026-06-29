// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetAsDefaultBranchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBranchId(v string) *SetAsDefaultBranchRequest
	GetBranchId() *string
	SetProjectId(v string) *SetAsDefaultBranchRequest
	GetProjectId() *string
	SetRegionId(v string) *SetAsDefaultBranchRequest
	GetRegionId() *string
}

type SetAsDefaultBranchRequest struct {
	// The branch ID that uniquely identifies a Supabase branch.
	//
	// This parameter is required.
	//
	// example:
	//
	// br-xxxx
	BranchId *string `json:"BranchId,omitempty" xml:"BranchId,omitempty"`
	// The ID of the Supabase project that corresponds to the primary branch.
	//
	// This parameter is required.
	//
	// example:
	//
	// spb-xxxx
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The region ID. This parameter is required when you create a primary branch. When you create a sub-branch, this parameter inherits the region of the primary branch by default.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s SetAsDefaultBranchRequest) String() string {
	return dara.Prettify(s)
}

func (s SetAsDefaultBranchRequest) GoString() string {
	return s.String()
}

func (s *SetAsDefaultBranchRequest) GetBranchId() *string {
	return s.BranchId
}

func (s *SetAsDefaultBranchRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *SetAsDefaultBranchRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SetAsDefaultBranchRequest) SetBranchId(v string) *SetAsDefaultBranchRequest {
	s.BranchId = &v
	return s
}

func (s *SetAsDefaultBranchRequest) SetProjectId(v string) *SetAsDefaultBranchRequest {
	s.ProjectId = &v
	return s
}

func (s *SetAsDefaultBranchRequest) SetRegionId(v string) *SetAsDefaultBranchRequest {
	s.RegionId = &v
	return s
}

func (s *SetAsDefaultBranchRequest) Validate() error {
	return dara.Validate(s)
}
