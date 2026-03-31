// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeResourceLogDeliveryStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeResourceLogDeliveryStatusRequest
	GetInstanceId() *string
	SetRegionId(v string) *DescribeResourceLogDeliveryStatusRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeResourceLogDeliveryStatusRequest
	GetResourceManagerResourceGroupId() *string
	SetResources(v string) *DescribeResourceLogDeliveryStatusRequest
	GetResources() *string
}

type DescribeResourceLogDeliveryStatusRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// waf_v3prepaid_public_cn-g6z3z*****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// alb-wewbb23dfsetetcic1242-0****,test.waf.com-waf
	Resources *string `json:"Resources,omitempty" xml:"Resources,omitempty"`
}

func (s DescribeResourceLogDeliveryStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourceLogDeliveryStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeResourceLogDeliveryStatusRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeResourceLogDeliveryStatusRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeResourceLogDeliveryStatusRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeResourceLogDeliveryStatusRequest) GetResources() *string {
	return s.Resources
}

func (s *DescribeResourceLogDeliveryStatusRequest) SetInstanceId(v string) *DescribeResourceLogDeliveryStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeResourceLogDeliveryStatusRequest) SetRegionId(v string) *DescribeResourceLogDeliveryStatusRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeResourceLogDeliveryStatusRequest) SetResourceManagerResourceGroupId(v string) *DescribeResourceLogDeliveryStatusRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeResourceLogDeliveryStatusRequest) SetResources(v string) *DescribeResourceLogDeliveryStatusRequest {
	s.Resources = &v
	return s
}

func (s *DescribeResourceLogDeliveryStatusRequest) Validate() error {
	return dara.Validate(s)
}
