// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHybridCloudBasicMonitorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeHybridCloudBasicMonitorRequest
	GetInstanceId() *string
	SetMid(v string) *DescribeHybridCloudBasicMonitorRequest
	GetMid() *string
	SetRegionId(v string) *DescribeHybridCloudBasicMonitorRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeHybridCloudBasicMonitorRequest
	GetResourceManagerResourceGroupId() *string
}

type DescribeHybridCloudBasicMonitorRequest struct {
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// >  You can call the [DescribeInstanceInfo](https://help.aliyun.com/document_detail/140857.html) operation to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf-cn-7mz****hw0u
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the node.
	//
	// This parameter is required.
	//
	// example:
	//
	// f451e90fcb068905ab379468****db42
	Mid *string `json:"Mid,omitempty" xml:"Mid,omitempty"`
	// The region in which the WAF instance is deployed. Valid values:
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

func (s DescribeHybridCloudBasicMonitorRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeHybridCloudBasicMonitorRequest) GoString() string {
	return s.String()
}

func (s *DescribeHybridCloudBasicMonitorRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeHybridCloudBasicMonitorRequest) GetMid() *string {
	return s.Mid
}

func (s *DescribeHybridCloudBasicMonitorRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeHybridCloudBasicMonitorRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeHybridCloudBasicMonitorRequest) SetInstanceId(v string) *DescribeHybridCloudBasicMonitorRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeHybridCloudBasicMonitorRequest) SetMid(v string) *DescribeHybridCloudBasicMonitorRequest {
	s.Mid = &v
	return s
}

func (s *DescribeHybridCloudBasicMonitorRequest) SetRegionId(v string) *DescribeHybridCloudBasicMonitorRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeHybridCloudBasicMonitorRequest) SetResourceManagerResourceGroupId(v string) *DescribeHybridCloudBasicMonitorRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeHybridCloudBasicMonitorRequest) Validate() error {
	return dara.Validate(s)
}
