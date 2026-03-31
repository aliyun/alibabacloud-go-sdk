// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudResourceAccessedPortsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeCloudResourceAccessedPortsRequest
	GetInstanceId() *string
	SetRegionId(v string) *DescribeCloudResourceAccessedPortsRequest
	GetRegionId() *string
	SetResourceInstanceId(v string) *DescribeCloudResourceAccessedPortsRequest
	GetResourceInstanceId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeCloudResourceAccessedPortsRequest
	GetResourceManagerResourceGroupId() *string
}

type DescribeCloudResourceAccessedPortsRequest struct {
	// The ID of the WAF instance.
	//
	// >  You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf-cn-tl32ast****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region in which the WAF instance is deployed. Valid values:
	//
	// 	- **cn-hangzhou**: the Chinese mainland.
	//
	// 	- **ap-southeast-1**: outside the Chinese mainland.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The instance ID of the resource.
	//
	// This parameter is required.
	//
	// example:
	//
	// lb-bp1*****jqnnqk5uj2p
	ResourceInstanceId *string `json:"ResourceInstanceId,omitempty" xml:"ResourceInstanceId,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	//
	// example:
	//
	// rg-aekzwwkpn****5i
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
}

func (s DescribeCloudResourceAccessedPortsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudResourceAccessedPortsRequest) GoString() string {
	return s.String()
}

func (s *DescribeCloudResourceAccessedPortsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeCloudResourceAccessedPortsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeCloudResourceAccessedPortsRequest) GetResourceInstanceId() *string {
	return s.ResourceInstanceId
}

func (s *DescribeCloudResourceAccessedPortsRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeCloudResourceAccessedPortsRequest) SetInstanceId(v string) *DescribeCloudResourceAccessedPortsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeCloudResourceAccessedPortsRequest) SetRegionId(v string) *DescribeCloudResourceAccessedPortsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeCloudResourceAccessedPortsRequest) SetResourceInstanceId(v string) *DescribeCloudResourceAccessedPortsRequest {
	s.ResourceInstanceId = &v
	return s
}

func (s *DescribeCloudResourceAccessedPortsRequest) SetResourceManagerResourceGroupId(v string) *DescribeCloudResourceAccessedPortsRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeCloudResourceAccessedPortsRequest) Validate() error {
	return dara.Validate(s)
}
