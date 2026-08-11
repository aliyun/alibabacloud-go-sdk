// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRegionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DescribeRegionsRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeRegionsRequest
	GetResourceGroupId() *string
}

type DescribeRegionsRequest struct {
	// The region ID to query. Default value: **ap-southeast-1**, which indicates that the regions of cloud assets that can be protected by the Anti-DDoS Origin instance in the China (Hangzhou) region are queried.
	//
	// To query other region IDs, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html) to obtain the corresponding **RegionId**.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the Anti-DDoS Origin instance belongs in Resource Management. By default, this parameter is empty, which indicates that the instance belongs to the default resource group.
	//
	// For more information about resource groups, see [Create a resource group](https://help.aliyun.com/document_detail/94485.html).
	//
	// example:
	//
	// rg-acfm2pz25js****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeRegionsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRegionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRegionsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeRegionsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeRegionsRequest) SetRegionId(v string) *DescribeRegionsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeRegionsRequest) SetResourceGroupId(v string) *DescribeRegionsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeRegionsRequest) Validate() error {
	return dara.Validate(s)
}
