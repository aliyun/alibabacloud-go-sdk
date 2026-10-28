// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGadInstanceDbMembersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeGadInstanceDbMembersRequest
	GetInstanceId() *string
	SetOwnerId(v string) *DescribeGadInstanceDbMembersRequest
	GetOwnerId() *string
	SetRegionId(v string) *DescribeGadInstanceDbMembersRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeGadInstanceDbMembersRequest
	GetResourceGroupId() *string
}

type DescribeGadInstanceDbMembersRequest struct {
	// The instance ID.
	//
	// example:
	//
	// gad-pfx10jt8b8m***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerId    *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmzawhxxc****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeGadInstanceDbMembersRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeGadInstanceDbMembersRequest) GoString() string {
	return s.String()
}

func (s *DescribeGadInstanceDbMembersRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeGadInstanceDbMembersRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *DescribeGadInstanceDbMembersRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeGadInstanceDbMembersRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeGadInstanceDbMembersRequest) SetInstanceId(v string) *DescribeGadInstanceDbMembersRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeGadInstanceDbMembersRequest) SetOwnerId(v string) *DescribeGadInstanceDbMembersRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeGadInstanceDbMembersRequest) SetRegionId(v string) *DescribeGadInstanceDbMembersRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeGadInstanceDbMembersRequest) SetResourceGroupId(v string) *DescribeGadInstanceDbMembersRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeGadInstanceDbMembersRequest) Validate() error {
	return dara.Validate(s)
}
