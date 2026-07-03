// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOperatorsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DescribeOperatorsRequest
	GetRegionId() *string
	SetRoleFor(v int64) *DescribeOperatorsRequest
	GetRoleFor() *int64
	SetRoleType(v int32) *DescribeOperatorsRequest
	GetRoleType() *int32
	SetSceneType(v string) *DescribeOperatorsRequest
	GetSceneType() *string
}

type DescribeOperatorsRequest struct {
	// The region of the Data Management center for threat analysis. Select a region based on your asset location. Valid values:
	//
	// - cn-hangzhou: Assets in the Chinese mainland and China (Hong Kong).
	//
	// - ap-southeast-1: Assets outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The user ID of the member. An administrator can use this parameter to switch to the perspective of this member.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The view type.
	//
	// - 0: The view of the current Alibaba Cloud account.
	//
	// - 1: The view of all accounts that are managed by your enterprise.
	//
	// example:
	//
	// 1
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
	// The scenario for the operator. Valid values:
	//
	// - If you leave this parameter empty, the default scenario is used.
	//
	// - AGGREGATE: The aggregate function scenario.
	//
	// example:
	//
	// AGGREGATE
	SceneType *string `json:"SceneType,omitempty" xml:"SceneType,omitempty"`
}

func (s DescribeOperatorsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeOperatorsRequest) GoString() string {
	return s.String()
}

func (s *DescribeOperatorsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeOperatorsRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *DescribeOperatorsRequest) GetRoleType() *int32 {
	return s.RoleType
}

func (s *DescribeOperatorsRequest) GetSceneType() *string {
	return s.SceneType
}

func (s *DescribeOperatorsRequest) SetRegionId(v string) *DescribeOperatorsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeOperatorsRequest) SetRoleFor(v int64) *DescribeOperatorsRequest {
	s.RoleFor = &v
	return s
}

func (s *DescribeOperatorsRequest) SetRoleType(v int32) *DescribeOperatorsRequest {
	s.RoleType = &v
	return s
}

func (s *DescribeOperatorsRequest) SetSceneType(v string) *DescribeOperatorsRequest {
	s.SceneType = &v
	return s
}

func (s *DescribeOperatorsRequest) Validate() error {
	return dara.Validate(s)
}
