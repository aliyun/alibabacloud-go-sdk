// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeProdCountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DescribeProdCountRequest
	GetRegionId() *string
	SetRoleFor(v int64) *DescribeProdCountRequest
	GetRoleFor() *int64
	SetRoleType(v int32) *DescribeProdCountRequest
	GetRoleType() *int32
}

type DescribeProdCountRequest struct {
	// The region where the Data Management center of Threat Analysis is located. Select a region for the Management Center based on the region where your assets reside. Valid values:
	//
	// - cn-hangzhou: Your assets are in the Chinese mainland or China (Hong Kong).
	//
	// - ap-southeast-1: Your assets are in a region outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The user ID of a member. This parameter allows an administrator to switch to the perspective of the member.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The type of the view.
	//
	// - 0: the view of the current Alibaba Cloud account.
	//
	// - 1: the view of all accounts that belong to the enterprise.
	//
	// example:
	//
	// 1
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s DescribeProdCountRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeProdCountRequest) GoString() string {
	return s.String()
}

func (s *DescribeProdCountRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeProdCountRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *DescribeProdCountRequest) GetRoleType() *int32 {
	return s.RoleType
}

func (s *DescribeProdCountRequest) SetRegionId(v string) *DescribeProdCountRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeProdCountRequest) SetRoleFor(v int64) *DescribeProdCountRequest {
	s.RoleFor = &v
	return s
}

func (s *DescribeProdCountRequest) SetRoleType(v int32) *DescribeProdCountRequest {
	s.RoleType = &v
	return s
}

func (s *DescribeProdCountRequest) Validate() error {
	return dara.Validate(s)
}
