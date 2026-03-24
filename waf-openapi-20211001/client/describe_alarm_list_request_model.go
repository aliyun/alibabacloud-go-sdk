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
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// > Call [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_v2_public_cn-pe33e****
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
