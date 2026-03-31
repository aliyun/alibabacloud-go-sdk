// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApisecLogDeliveriesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeApisecLogDeliveriesRequest
	GetInstanceId() *string
	SetRegionId(v string) *DescribeApisecLogDeliveriesRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeApisecLogDeliveriesRequest
	GetResourceManagerResourceGroupId() *string
}

type DescribeApisecLogDeliveriesRequest struct {
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// >  You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_v3_public_cn-uqm2z****0a
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region where the WAF instance is deployed. Valid values:
	//
	// 	- **cn-hangzhou**: Chinese mainland.
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

func (s DescribeApisecLogDeliveriesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeApisecLogDeliveriesRequest) GoString() string {
	return s.String()
}

func (s *DescribeApisecLogDeliveriesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeApisecLogDeliveriesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeApisecLogDeliveriesRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeApisecLogDeliveriesRequest) SetInstanceId(v string) *DescribeApisecLogDeliveriesRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeApisecLogDeliveriesRequest) SetRegionId(v string) *DescribeApisecLogDeliveriesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeApisecLogDeliveriesRequest) SetResourceManagerResourceGroupId(v string) *DescribeApisecLogDeliveriesRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeApisecLogDeliveriesRequest) Validate() error {
	return dara.Validate(s)
}
