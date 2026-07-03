// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWafScopeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEntityId(v int64) *DescribeWafScopeRequest
	GetEntityId() *int64
	SetRegionId(v string) *DescribeWafScopeRequest
	GetRegionId() *string
	SetRoleFor(v int64) *DescribeWafScopeRequest
	GetRoleFor() *int64
	SetRoleType(v int32) *DescribeWafScopeRequest
	GetRoleType() *int32
}

type DescribeWafScopeRequest struct {
	// The entity ID.
	//
	// example:
	//
	// 20617784
	EntityId *int64 `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// The region where the Data Management center of threat analysis is located. Select a region based on the location of your assets. Valid values:
	//
	// - cn-hangzhou: Your assets are in the Chinese mainland or China (Hong Kong).
	//
	// - ap-southeast-1: Your assets are outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of a member. An administrator can use this parameter to assume the permissions of the member.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The view type.
	//
	// - 0: the view of the current Alibaba Cloud account.
	//
	// - 1: the view of all accounts within the enterprise.
	//
	// example:
	//
	// 1
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s DescribeWafScopeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeWafScopeRequest) GoString() string {
	return s.String()
}

func (s *DescribeWafScopeRequest) GetEntityId() *int64 {
	return s.EntityId
}

func (s *DescribeWafScopeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeWafScopeRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *DescribeWafScopeRequest) GetRoleType() *int32 {
	return s.RoleType
}

func (s *DescribeWafScopeRequest) SetEntityId(v int64) *DescribeWafScopeRequest {
	s.EntityId = &v
	return s
}

func (s *DescribeWafScopeRequest) SetRegionId(v string) *DescribeWafScopeRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeWafScopeRequest) SetRoleFor(v int64) *DescribeWafScopeRequest {
	s.RoleFor = &v
	return s
}

func (s *DescribeWafScopeRequest) SetRoleType(v int32) *DescribeWafScopeRequest {
	s.RoleType = &v
	return s
}

func (s *DescribeWafScopeRequest) Validate() error {
	return dara.Validate(s)
}
