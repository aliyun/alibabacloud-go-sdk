// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFlowChartRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTimestamp(v string) *DescribeFlowChartRequest
	GetEndTimestamp() *string
	SetInstanceId(v string) *DescribeFlowChartRequest
	GetInstanceId() *string
	SetInterval(v string) *DescribeFlowChartRequest
	GetInterval() *string
	SetRegionId(v string) *DescribeFlowChartRequest
	GetRegionId() *string
	SetResource(v string) *DescribeFlowChartRequest
	GetResource() *string
	SetResourceManagerResourceGroupId(v string) *DescribeFlowChartRequest
	GetResourceManagerResourceGroupId() *string
	SetStartTimestamp(v string) *DescribeFlowChartRequest
	GetStartTimestamp() *string
}

type DescribeFlowChartRequest struct {
	// The end of the time range to query. Unit: seconds. If you do not specify this parameter, the current time is used.
	//
	// example:
	//
	// 1665386280
	EndTimestamp *string `json:"EndTimestamp,omitempty" xml:"EndTimestamp,omitempty"`
	// The ID of the WAF instance.
	//
	// >  You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to obtain the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_cdnsdf3****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The time interval. Unit: seconds. The value must be an integral multiple of 60.
	//
	// This parameter is required.
	//
	// example:
	//
	// 300
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
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
	// The protected object.
	//
	// example:
	//
	// www.aliyundoc.com
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	//
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The beginning of the time range to query. Unit: seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1665331200
	StartTimestamp *string `json:"StartTimestamp,omitempty" xml:"StartTimestamp,omitempty"`
}

func (s DescribeFlowChartRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeFlowChartRequest) GoString() string {
	return s.String()
}

func (s *DescribeFlowChartRequest) GetEndTimestamp() *string {
	return s.EndTimestamp
}

func (s *DescribeFlowChartRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeFlowChartRequest) GetInterval() *string {
	return s.Interval
}

func (s *DescribeFlowChartRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeFlowChartRequest) GetResource() *string {
	return s.Resource
}

func (s *DescribeFlowChartRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeFlowChartRequest) GetStartTimestamp() *string {
	return s.StartTimestamp
}

func (s *DescribeFlowChartRequest) SetEndTimestamp(v string) *DescribeFlowChartRequest {
	s.EndTimestamp = &v
	return s
}

func (s *DescribeFlowChartRequest) SetInstanceId(v string) *DescribeFlowChartRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeFlowChartRequest) SetInterval(v string) *DescribeFlowChartRequest {
	s.Interval = &v
	return s
}

func (s *DescribeFlowChartRequest) SetRegionId(v string) *DescribeFlowChartRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeFlowChartRequest) SetResource(v string) *DescribeFlowChartRequest {
	s.Resource = &v
	return s
}

func (s *DescribeFlowChartRequest) SetResourceManagerResourceGroupId(v string) *DescribeFlowChartRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeFlowChartRequest) SetStartTimestamp(v string) *DescribeFlowChartRequest {
	s.StartTimestamp = &v
	return s
}

func (s *DescribeFlowChartRequest) Validate() error {
	return dara.Validate(s)
}
