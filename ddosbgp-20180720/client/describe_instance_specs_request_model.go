// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceSpecsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceIdList(v string) *DescribeInstanceSpecsRequest
	GetInstanceIdList() *string
	SetRegionId(v string) *DescribeInstanceSpecsRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeInstanceSpecsRequest
	GetResourceGroupId() *string
}

type DescribeInstanceSpecsRequest struct {
	// The IDs of the Anti-DDoS Origin instances to query. The value is a JSON array converted to a string. Each element in the JSON array represents an instance ID in string format. Multiple IDs are separated by commas (,).
	//
	// > You can call the [DescribeInstanceList](https://help.aliyun.com/document_detail/118698.html) operation to query the IDs of all Anti-DDoS Origin instances in a specified region.
	//
	// This parameter is required.
	//
	// example:
	//
	// ["ddosbgp-cn-n6w1r7nz****"]
	InstanceIdList *string `json:"InstanceIdList,omitempty" xml:"InstanceIdList,omitempty"`
	// The region ID of the Anti-DDoS Origin instance. Default value: **cn-hangzhou**, which indicates the China (Hangzhou) region.
	//
	// > If your instance is not in the China (Hangzhou) region, you must specify the region ID of the instance. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/118703.html) operation to query all region IDs.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the Anti-DDoS Origin instance belongs in Resource Management. Default value: empty, which indicates the default resource group.
	//
	// For more information about resource groups, see [Create a resource group](https://help.aliyun.com/document_detail/94485.html).
	//
	// example:
	//
	// rg-acfm2pz25js****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeInstanceSpecsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceSpecsRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSpecsRequest) GetInstanceIdList() *string {
	return s.InstanceIdList
}

func (s *DescribeInstanceSpecsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeInstanceSpecsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeInstanceSpecsRequest) SetInstanceIdList(v string) *DescribeInstanceSpecsRequest {
	s.InstanceIdList = &v
	return s
}

func (s *DescribeInstanceSpecsRequest) SetRegionId(v string) *DescribeInstanceSpecsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeInstanceSpecsRequest) SetResourceGroupId(v string) *DescribeInstanceSpecsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeInstanceSpecsRequest) Validate() error {
	return dara.Validate(s)
}
