// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserLogFieldConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeliveryType(v string) *DescribeUserLogFieldConfigRequest
	GetDeliveryType() *string
	SetInstanceId(v string) *DescribeUserLogFieldConfigRequest
	GetInstanceId() *string
	SetRegionId(v string) *DescribeUserLogFieldConfigRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeUserLogFieldConfigRequest
	GetResourceManagerResourceGroupId() *string
}

type DescribeUserLogFieldConfigRequest struct {
	// example:
	//
	// sls
	DeliveryType *string `json:"DeliveryType,omitempty" xml:"DeliveryType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// waf_v3prepaid_public_cn-0*****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-aek24******
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
}

func (s DescribeUserLogFieldConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserLogFieldConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserLogFieldConfigRequest) GetDeliveryType() *string {
	return s.DeliveryType
}

func (s *DescribeUserLogFieldConfigRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeUserLogFieldConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeUserLogFieldConfigRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeUserLogFieldConfigRequest) SetDeliveryType(v string) *DescribeUserLogFieldConfigRequest {
	s.DeliveryType = &v
	return s
}

func (s *DescribeUserLogFieldConfigRequest) SetInstanceId(v string) *DescribeUserLogFieldConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeUserLogFieldConfigRequest) SetRegionId(v string) *DescribeUserLogFieldConfigRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeUserLogFieldConfigRequest) SetResourceManagerResourceGroupId(v string) *DescribeUserLogFieldConfigRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeUserLogFieldConfigRequest) Validate() error {
	return dara.Validate(s)
}
