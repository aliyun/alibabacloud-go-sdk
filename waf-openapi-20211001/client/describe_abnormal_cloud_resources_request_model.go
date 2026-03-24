// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAbnormalCloudResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeAbnormalCloudResourcesRequest
	GetInstanceId() *string
	SetRegionId(v string) *DescribeAbnormalCloudResourcesRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeAbnormalCloudResourcesRequest
	GetResourceManagerResourceGroupId() *string
}

type DescribeAbnormalCloudResourcesRequest struct {
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// > Call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query the ID of your WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_v3prepaid_public_cn-***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region where the WAF instance resides. Valid values:
	//
	// - **cn-hangzhou**: the Chinese mainland.
	//
	// - **ap-southeast-1**: a region outside the Chinese mainland.
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

func (s DescribeAbnormalCloudResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAbnormalCloudResourcesRequest) GoString() string {
	return s.String()
}

func (s *DescribeAbnormalCloudResourcesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeAbnormalCloudResourcesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAbnormalCloudResourcesRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeAbnormalCloudResourcesRequest) SetInstanceId(v string) *DescribeAbnormalCloudResourcesRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeAbnormalCloudResourcesRequest) SetRegionId(v string) *DescribeAbnormalCloudResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAbnormalCloudResourcesRequest) SetResourceManagerResourceGroupId(v string) *DescribeAbnormalCloudResourcesRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeAbnormalCloudResourcesRequest) Validate() error {
	return dara.Validate(s)
}
