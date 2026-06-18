// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBranchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBranchId(v string) *DescribeBranchRequest
	GetBranchId() *string
	SetRegionId(v string) *DescribeBranchRequest
	GetRegionId() *string
}

type DescribeBranchRequest struct {
	// The branch ID, which uniquely identifies a Supabase branch.
	//
	// This parameter is required.
	//
	// example:
	//
	// br-xxxx
	BranchId *string `json:"BranchId,omitempty" xml:"BranchId,omitempty"`
	// The region ID. This parameter is required when you create a primary branch. When you create a child branch, the region of the primary branch is inherited by default.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeBranchRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeBranchRequest) GoString() string {
	return s.String()
}

func (s *DescribeBranchRequest) GetBranchId() *string {
	return s.BranchId
}

func (s *DescribeBranchRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeBranchRequest) SetBranchId(v string) *DescribeBranchRequest {
	s.BranchId = &v
	return s
}

func (s *DescribeBranchRequest) SetRegionId(v string) *DescribeBranchRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeBranchRequest) Validate() error {
	return dara.Validate(s)
}
