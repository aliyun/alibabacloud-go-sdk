// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGadInstanceDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeGadInstanceDetailRequest
	GetInstanceId() *string
	SetOwnerId(v string) *DescribeGadInstanceDetailRequest
	GetOwnerId() *string
	SetRegionId(v string) *DescribeGadInstanceDetailRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeGadInstanceDetailRequest
	GetResourceGroupId() *string
}

type DescribeGadInstanceDetailRequest struct {
	// The instance ID of the primary instance.
	//
	// example:
	//
	// gad-bp162d4tp0500****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerId    *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the task. For more information, see [Supported regions](https://help.aliyun.com/document_detail/141033.html).
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

func (s DescribeGadInstanceDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeGadInstanceDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeGadInstanceDetailRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeGadInstanceDetailRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *DescribeGadInstanceDetailRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeGadInstanceDetailRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeGadInstanceDetailRequest) SetInstanceId(v string) *DescribeGadInstanceDetailRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeGadInstanceDetailRequest) SetOwnerId(v string) *DescribeGadInstanceDetailRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeGadInstanceDetailRequest) SetRegionId(v string) *DescribeGadInstanceDetailRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeGadInstanceDetailRequest) SetResourceGroupId(v string) *DescribeGadInstanceDetailRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeGadInstanceDetailRequest) Validate() error {
	return dara.Validate(s)
}
