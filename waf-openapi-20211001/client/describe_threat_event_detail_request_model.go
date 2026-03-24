// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeThreatEventDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEventId(v string) *DescribeThreatEventDetailRequest
	GetEventId() *string
	SetInstanceId(v string) *DescribeThreatEventDetailRequest
	GetInstanceId() *string
	SetRegionId(v string) *DescribeThreatEventDetailRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeThreatEventDetailRequest
	GetResourceManagerResourceGroupId() *string
}

type DescribeThreatEventDetailRequest struct {
	// The ID of the security event.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1661131a028f72a976703f4a4082ad87
	EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// > You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_v2_public_cn-lbj*****
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
	// The ID of the Alibaba Cloud resource group.
	//
	// example:
	//
	// rg-aekzhks66****
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
}

func (s DescribeThreatEventDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeThreatEventDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeThreatEventDetailRequest) GetEventId() *string {
	return s.EventId
}

func (s *DescribeThreatEventDetailRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeThreatEventDetailRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeThreatEventDetailRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeThreatEventDetailRequest) SetEventId(v string) *DescribeThreatEventDetailRequest {
	s.EventId = &v
	return s
}

func (s *DescribeThreatEventDetailRequest) SetInstanceId(v string) *DescribeThreatEventDetailRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeThreatEventDetailRequest) SetRegionId(v string) *DescribeThreatEventDetailRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeThreatEventDetailRequest) SetResourceManagerResourceGroupId(v string) *DescribeThreatEventDetailRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeThreatEventDetailRequest) Validate() error {
	return dara.Validate(s)
}
