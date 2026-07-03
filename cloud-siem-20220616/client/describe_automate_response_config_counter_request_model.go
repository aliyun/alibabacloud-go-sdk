// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAutomateResponseConfigCounterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DescribeAutomateResponseConfigCounterRequest
	GetRegionId() *string
	SetRoleFor(v int64) *DescribeAutomateResponseConfigCounterRequest
	GetRoleFor() *int64
	SetRoleType(v int32) *DescribeAutomateResponseConfigCounterRequest
	GetRoleType() *int32
}

type DescribeAutomateResponseConfigCounterRequest struct {
	// The region of the Management Hub. Select the region based on the location of your assets. Valid values:
	//
	// - cn-hangzhou: Your assets are in the Chinese mainland or China (Hong Kong).
	//
	// - ap-southeast-1: Your assets are in a region outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The user ID of the member whose view the administrator wants to switch to.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The view type.
	//
	// - 0: The view of the current Alibaba Cloud account.
	//
	// - 1: The view of all accounts in the enterprise.
	//
	// example:
	//
	// 1
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s DescribeAutomateResponseConfigCounterRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutomateResponseConfigCounterRequest) GoString() string {
	return s.String()
}

func (s *DescribeAutomateResponseConfigCounterRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAutomateResponseConfigCounterRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *DescribeAutomateResponseConfigCounterRequest) GetRoleType() *int32 {
	return s.RoleType
}

func (s *DescribeAutomateResponseConfigCounterRequest) SetRegionId(v string) *DescribeAutomateResponseConfigCounterRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAutomateResponseConfigCounterRequest) SetRoleFor(v int64) *DescribeAutomateResponseConfigCounterRequest {
	s.RoleFor = &v
	return s
}

func (s *DescribeAutomateResponseConfigCounterRequest) SetRoleType(v int32) *DescribeAutomateResponseConfigCounterRequest {
	s.RoleType = &v
	return s
}

func (s *DescribeAutomateResponseConfigCounterRequest) Validate() error {
	return dara.Validate(s)
}
