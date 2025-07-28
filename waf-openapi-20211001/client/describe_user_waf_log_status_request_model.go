// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserWafLogStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeUserWafLogStatusRequest
	GetInstanceId() *string
	SetRegionId(v string) *DescribeUserWafLogStatusRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeUserWafLogStatusRequest
	GetResourceManagerResourceGroupId() *string
}

type DescribeUserWafLogStatusRequest struct {
	// The ID of the WAF instance.
	//
	// > You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_v3prepaid_public_cn-nwy34****0j
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
	// 阿里云资源组ID。
	//
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
}

func (s DescribeUserWafLogStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserWafLogStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserWafLogStatusRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeUserWafLogStatusRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeUserWafLogStatusRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeUserWafLogStatusRequest) SetInstanceId(v string) *DescribeUserWafLogStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeUserWafLogStatusRequest) SetRegionId(v string) *DescribeUserWafLogStatusRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeUserWafLogStatusRequest) SetResourceManagerResourceGroupId(v string) *DescribeUserWafLogStatusRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeUserWafLogStatusRequest) Validate() error {
	return dara.Validate(s)
}
