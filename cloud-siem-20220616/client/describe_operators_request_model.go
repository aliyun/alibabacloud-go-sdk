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
	// The region in which the data management center of the threat analysis feature resides. Specify this parameter based on the regions in which your assets reside. Valid values:
	//
	// 	- cn-hangzhou: Your assets reside in regions in China.
	//
	// 	- ap-southeast-1: Your assets reside in regions outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the account that you switch from the management account.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The type of the view. Valid values:
	//
	// - 0: the current Alibaba Cloud account
	//
	// - 1: the global account
	//
	// example:
	//
	// 1
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
	// The type of the scenario in which the operator is used. Valid values:
	//
	// 	- If you do not specify this parameter, the default scenario is used.
	//
	// 	- AGGREGATE: AGGREGATE scenario.
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
