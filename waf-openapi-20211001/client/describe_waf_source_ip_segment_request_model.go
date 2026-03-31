// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWafSourceIpSegmentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeWafSourceIpSegmentRequest
	GetInstanceId() *string
	SetRegionId(v string) *DescribeWafSourceIpSegmentRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeWafSourceIpSegmentRequest
	GetResourceManagerResourceGroupId() *string
}

type DescribeWafSourceIpSegmentRequest struct {
	// The ID of the WAF instance.
	//
	// >  You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to obtain the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_v2_public_cn-*****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region where the WAF instance resides. Valid values:
	//
	// 	- **cn-hangzhou:*	- the Chinese mainland.
	//
	// 	- **ap-southeast-1:*	- outside the Chinese mainland.
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

func (s DescribeWafSourceIpSegmentRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeWafSourceIpSegmentRequest) GoString() string {
	return s.String()
}

func (s *DescribeWafSourceIpSegmentRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeWafSourceIpSegmentRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeWafSourceIpSegmentRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeWafSourceIpSegmentRequest) SetInstanceId(v string) *DescribeWafSourceIpSegmentRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeWafSourceIpSegmentRequest) SetRegionId(v string) *DescribeWafSourceIpSegmentRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeWafSourceIpSegmentRequest) SetResourceManagerResourceGroupId(v string) *DescribeWafSourceIpSegmentRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeWafSourceIpSegmentRequest) Validate() error {
	return dara.Validate(s)
}
