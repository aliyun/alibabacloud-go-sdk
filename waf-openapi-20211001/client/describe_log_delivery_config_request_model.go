// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLogDeliveryConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeliveryName(v string) *DescribeLogDeliveryConfigRequest
	GetDeliveryName() *string
	SetInstanceId(v string) *DescribeLogDeliveryConfigRequest
	GetInstanceId() *string
	SetRegionId(v string) *DescribeLogDeliveryConfigRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeLogDeliveryConfigRequest
	GetResourceManagerResourceGroupId() *string
}

type DescribeLogDeliveryConfigRequest struct {
	// The name of the log delivery configuration.
	//
	// This parameter is required.
	//
	// example:
	//
	// test1
	DeliveryName *string `json:"DeliveryName,omitempty" xml:"DeliveryName,omitempty"`
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// >  You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_elasticity-*****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the WAF instance. Valid values:
	//
	// 	- **cn-hangzhou**: the Chinese mainland.
	//
	// 	- **ap-southeast-1**: outside the Chinese mainland.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	//
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
}

func (s DescribeLogDeliveryConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLogDeliveryConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeLogDeliveryConfigRequest) GetDeliveryName() *string {
	return s.DeliveryName
}

func (s *DescribeLogDeliveryConfigRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeLogDeliveryConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLogDeliveryConfigRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeLogDeliveryConfigRequest) SetDeliveryName(v string) *DescribeLogDeliveryConfigRequest {
	s.DeliveryName = &v
	return s
}

func (s *DescribeLogDeliveryConfigRequest) SetInstanceId(v string) *DescribeLogDeliveryConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeLogDeliveryConfigRequest) SetRegionId(v string) *DescribeLogDeliveryConfigRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLogDeliveryConfigRequest) SetResourceManagerResourceGroupId(v string) *DescribeLogDeliveryConfigRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeLogDeliveryConfigRequest) Validate() error {
	return dara.Validate(s)
}
