// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAlarmListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeAlarmListRequest
	GetInstanceId() *string
	SetRegionId(v string) *DescribeAlarmListRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeAlarmListRequest
	GetResourceManagerResourceGroupId() *string
}

type DescribeAlarmListRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// waf_v2_public_cn-pe33e****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-aek2ipnu***
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
}

func (s DescribeAlarmListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAlarmListRequest) GoString() string {
	return s.String()
}

func (s *DescribeAlarmListRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeAlarmListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAlarmListRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeAlarmListRequest) SetInstanceId(v string) *DescribeAlarmListRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeAlarmListRequest) SetRegionId(v string) *DescribeAlarmListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAlarmListRequest) SetResourceManagerResourceGroupId(v string) *DescribeAlarmListRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeAlarmListRequest) Validate() error {
	return dara.Validate(s)
}
