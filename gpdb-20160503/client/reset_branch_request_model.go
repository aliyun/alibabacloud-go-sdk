// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetBranchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBranchId(v string) *ResetBranchRequest
	GetBranchId() *string
	SetProjectId(v string) *ResetBranchRequest
	GetProjectId() *string
	SetRegionId(v string) *ResetBranchRequest
	GetRegionId() *string
}

type ResetBranchRequest struct {
	// The branch ID, which uniquely identifies a Supabase branch.
	//
	// This parameter is required.
	//
	// example:
	//
	// br-xxxx
	BranchId *string `json:"BranchId,omitempty" xml:"BranchId,omitempty"`
	// The ID of the Supabase project to which the main branch belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// spb-xxxx
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The region ID. This parameter is required when you create a main branch. When you create a child branch, it inherits the region of the main branch by default.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ResetBranchRequest) String() string {
	return dara.Prettify(s)
}

func (s ResetBranchRequest) GoString() string {
	return s.String()
}

func (s *ResetBranchRequest) GetBranchId() *string {
	return s.BranchId
}

func (s *ResetBranchRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *ResetBranchRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ResetBranchRequest) SetBranchId(v string) *ResetBranchRequest {
	s.BranchId = &v
	return s
}

func (s *ResetBranchRequest) SetProjectId(v string) *ResetBranchRequest {
	s.ProjectId = &v
	return s
}

func (s *ResetBranchRequest) SetRegionId(v string) *ResetBranchRequest {
	s.RegionId = &v
	return s
}

func (s *ResetBranchRequest) Validate() error {
	return dara.Validate(s)
}
