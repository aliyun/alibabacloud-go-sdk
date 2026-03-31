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
	// This parameter is required.
	//
	// example:
	//
	// 1661131a028f72a976703f4a4082ad87
	EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// waf_v2_public_cn-lbj*****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
