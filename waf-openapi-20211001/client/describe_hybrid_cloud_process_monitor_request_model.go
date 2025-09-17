// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHybridCloudProcessMonitorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeHybridCloudProcessMonitorRequest
	GetInstanceId() *string
	SetMid(v string) *DescribeHybridCloudProcessMonitorRequest
	GetMid() *string
	SetRegionId(v string) *DescribeHybridCloudProcessMonitorRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeHybridCloudProcessMonitorRequest
	GetResourceManagerResourceGroupId() *string
}

type DescribeHybridCloudProcessMonitorRequest struct {
	// The ID of the Web Application Firewall (WAF) instance that you want to query.
	//
	// >  You can call the [DescribeInstanceInfo](https://help.aliyun.com/document_detail/140857.html) operation to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf-cn-7mz****hw0u
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the SDK.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2fdb63fea03b173bc9e65c24****d7d5
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

func (s DescribeHybridCloudProcessMonitorRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeHybridCloudProcessMonitorRequest) GoString() string {
	return s.String()
}

func (s *DescribeHybridCloudProcessMonitorRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeHybridCloudProcessMonitorRequest) GetMid() *string {
	return s.Mid
}

func (s *DescribeHybridCloudProcessMonitorRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeHybridCloudProcessMonitorRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeHybridCloudProcessMonitorRequest) SetInstanceId(v string) *DescribeHybridCloudProcessMonitorRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeHybridCloudProcessMonitorRequest) SetMid(v string) *DescribeHybridCloudProcessMonitorRequest {
	s.Mid = &v
	return s
}

func (s *DescribeHybridCloudProcessMonitorRequest) SetRegionId(v string) *DescribeHybridCloudProcessMonitorRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeHybridCloudProcessMonitorRequest) SetResourceManagerResourceGroupId(v string) *DescribeHybridCloudProcessMonitorRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeHybridCloudProcessMonitorRequest) Validate() error {
	return dara.Validate(s)
}
