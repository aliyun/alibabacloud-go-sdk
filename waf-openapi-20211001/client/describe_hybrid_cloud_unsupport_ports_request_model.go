// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHybridCloudUnsupportPortsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeHybridCloudUnsupportPortsRequest
	GetInstanceId() *string
	SetRegionId(v string) *DescribeHybridCloudUnsupportPortsRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeHybridCloudUnsupportPortsRequest
	GetResourceManagerResourceGroupId() *string
}

type DescribeHybridCloudUnsupportPortsRequest struct {
	// The ID of the WAF instance.
	//
	// >  You can call the [DescribeInstanceInfo](https://help.aliyun.com/document_detail/140857.html) operation to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_cdnsdf3****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the Web Application Firewall (WAF) instance. Valid values:
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

func (s DescribeHybridCloudUnsupportPortsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeHybridCloudUnsupportPortsRequest) GoString() string {
	return s.String()
}

func (s *DescribeHybridCloudUnsupportPortsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeHybridCloudUnsupportPortsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeHybridCloudUnsupportPortsRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeHybridCloudUnsupportPortsRequest) SetInstanceId(v string) *DescribeHybridCloudUnsupportPortsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeHybridCloudUnsupportPortsRequest) SetRegionId(v string) *DescribeHybridCloudUnsupportPortsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeHybridCloudUnsupportPortsRequest) SetResourceManagerResourceGroupId(v string) *DescribeHybridCloudUnsupportPortsRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeHybridCloudUnsupportPortsRequest) Validate() error {
	return dara.Validate(s)
}
