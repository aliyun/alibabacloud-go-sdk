// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeParameterGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *DescribeParameterGroupRequest
	GetOwnerId() *int64
	SetParameterGroupId(v string) *DescribeParameterGroupRequest
	GetParameterGroupId() *string
	SetRegionId(v string) *DescribeParameterGroupRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeParameterGroupRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeParameterGroupRequest
	GetResourceOwnerId() *int64
}

type DescribeParameterGroupRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The parameter template ID. You can call the DescribeParameterGroups operation to query the parameter template ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rpg-dp****
	ParameterGroupId *string `json:"ParameterGroupId,omitempty" xml:"ParameterGroupId,omitempty"`
	// The region ID. You can call the DescribeRegions operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeParameterGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeParameterGroupRequest) GoString() string {
	return s.String()
}

func (s *DescribeParameterGroupRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeParameterGroupRequest) GetParameterGroupId() *string {
	return s.ParameterGroupId
}

func (s *DescribeParameterGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeParameterGroupRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeParameterGroupRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeParameterGroupRequest) SetOwnerId(v int64) *DescribeParameterGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeParameterGroupRequest) SetParameterGroupId(v string) *DescribeParameterGroupRequest {
	s.ParameterGroupId = &v
	return s
}

func (s *DescribeParameterGroupRequest) SetRegionId(v string) *DescribeParameterGroupRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeParameterGroupRequest) SetResourceOwnerAccount(v string) *DescribeParameterGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeParameterGroupRequest) SetResourceOwnerId(v int64) *DescribeParameterGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeParameterGroupRequest) Validate() error {
	return dara.Validate(s)
}
