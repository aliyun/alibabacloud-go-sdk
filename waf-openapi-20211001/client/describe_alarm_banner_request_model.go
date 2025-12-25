// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAlarmBannerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeAlarmBannerRequest
	GetInstanceId() *string
	SetRegionId(v string) *DescribeAlarmBannerRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeAlarmBannerRequest
	GetResourceManagerResourceGroupId() *string
}

type DescribeAlarmBannerRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// waf_v3prepaid_public_cn-kf74****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-aek2zerdgm****
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
}

func (s DescribeAlarmBannerRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAlarmBannerRequest) GoString() string {
	return s.String()
}

func (s *DescribeAlarmBannerRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeAlarmBannerRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAlarmBannerRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeAlarmBannerRequest) SetInstanceId(v string) *DescribeAlarmBannerRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeAlarmBannerRequest) SetRegionId(v string) *DescribeAlarmBannerRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAlarmBannerRequest) SetResourceManagerResourceGroupId(v string) *DescribeAlarmBannerRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeAlarmBannerRequest) Validate() error {
	return dara.Validate(s)
}
