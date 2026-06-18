// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBranchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBranchId(v string) *DeleteBranchRequest
	GetBranchId() *string
	SetRegionId(v string) *DeleteBranchRequest
	GetRegionId() *string
}

type DeleteBranchRequest struct {
	// The branch ID that uniquely identifies a Supabase branch.
	//
	// This parameter is required.
	//
	// example:
	//
	// br-xxxx
	BranchId *string `json:"BranchId,omitempty" xml:"BranchId,omitempty"`
	// The region ID. This parameter is required when you create a primary branch. When you create a child branch, the region is inherited from the primary branch by default.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteBranchRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteBranchRequest) GoString() string {
	return s.String()
}

func (s *DeleteBranchRequest) GetBranchId() *string {
	return s.BranchId
}

func (s *DeleteBranchRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteBranchRequest) SetBranchId(v string) *DeleteBranchRequest {
	s.BranchId = &v
	return s
}

func (s *DeleteBranchRequest) SetRegionId(v string) *DeleteBranchRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteBranchRequest) Validate() error {
	return dara.Validate(s)
}
