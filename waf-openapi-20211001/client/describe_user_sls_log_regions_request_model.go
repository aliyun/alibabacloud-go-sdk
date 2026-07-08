// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserSlsLogRegionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeUserSlsLogRegionsRequest
	GetInstanceId() *string
	SetRegionId(v string) *DescribeUserSlsLogRegionsRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeUserSlsLogRegionsRequest
	GetResourceManagerResourceGroupId() *string
}

type DescribeUserSlsLogRegionsRequest struct {
	// Instance ID of the WAF instance.
	//
	// > You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query instance ID of the current WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_v2_public_cn-uqm2z****0a
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region where the WAF instance resides. Valid values:
	//
	// - **cn-hangzhou**: the Chinese mainland.
	//
	// - **ap-southeast-1**: outside the Chinese mainland.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The Alibaba Cloud resource group ID.
	//
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
}

func (s DescribeUserSlsLogRegionsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserSlsLogRegionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserSlsLogRegionsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeUserSlsLogRegionsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeUserSlsLogRegionsRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeUserSlsLogRegionsRequest) SetInstanceId(v string) *DescribeUserSlsLogRegionsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeUserSlsLogRegionsRequest) SetRegionId(v string) *DescribeUserSlsLogRegionsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeUserSlsLogRegionsRequest) SetResourceManagerResourceGroupId(v string) *DescribeUserSlsLogRegionsRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeUserSlsLogRegionsRequest) Validate() error {
	return dara.Validate(s)
}
