// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGadInstanceDtsMembersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeGadInstanceDtsMembersRequest
	GetInstanceId() *string
	SetOwnerId(v string) *DescribeGadInstanceDtsMembersRequest
	GetOwnerId() *string
	SetRegionId(v string) *DescribeGadInstanceDtsMembersRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeGadInstanceDtsMembersRequest
	GetResourceGroupId() *string
}

type DescribeGadInstanceDtsMembersRequest struct {
	// The instance ID of the active-active instance group.
	//
	// example:
	//
	// rm-bp1i99e8l7913****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerId    *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region where the DTS instance resides. For more information, see [Supported regions](https://help.aliyun.com/document_detail/141033.html).
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

func (s DescribeGadInstanceDtsMembersRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeGadInstanceDtsMembersRequest) GoString() string {
	return s.String()
}

func (s *DescribeGadInstanceDtsMembersRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeGadInstanceDtsMembersRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *DescribeGadInstanceDtsMembersRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeGadInstanceDtsMembersRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeGadInstanceDtsMembersRequest) SetInstanceId(v string) *DescribeGadInstanceDtsMembersRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersRequest) SetOwnerId(v string) *DescribeGadInstanceDtsMembersRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersRequest) SetRegionId(v string) *DescribeGadInstanceDtsMembersRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersRequest) SetResourceGroupId(v string) *DescribeGadInstanceDtsMembersRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersRequest) Validate() error {
	return dara.Validate(s)
}
