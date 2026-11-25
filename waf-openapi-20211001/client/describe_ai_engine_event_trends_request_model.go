// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAiEngineEventTrendsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *DescribeAiEngineEventTrendsRequest
	GetEndTime() *int64
	SetInstanceId(v string) *DescribeAiEngineEventTrendsRequest
	GetInstanceId() *string
	SetRegionId(v string) *DescribeAiEngineEventTrendsRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeAiEngineEventTrendsRequest
	GetResourceManagerResourceGroupId() *string
	SetStartTime(v int64) *DescribeAiEngineEventTrendsRequest
	GetStartTime() *int64
}

type DescribeAiEngineEventTrendsRequest struct {
	// The end time in the UNIX timestamp format (UTC). Unit: seconds. Events are filtered based on their latest discovery time. The end time must be later than the start time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1792973400
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the WAF instance.
	//
	// > You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query the ID of the current WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf-deddd*****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the WAF instance.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	//
	// example:
	//
	// rg-dess*****
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The start time in the UNIX timestamp format (UTC). Unit: seconds. Events are filtered based on their latest discovery time. You can query data from the last 30 days.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1792973400
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeAiEngineEventTrendsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAiEngineEventTrendsRequest) GoString() string {
	return s.String()
}

func (s *DescribeAiEngineEventTrendsRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeAiEngineEventTrendsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeAiEngineEventTrendsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAiEngineEventTrendsRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeAiEngineEventTrendsRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeAiEngineEventTrendsRequest) SetEndTime(v int64) *DescribeAiEngineEventTrendsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeAiEngineEventTrendsRequest) SetInstanceId(v string) *DescribeAiEngineEventTrendsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeAiEngineEventTrendsRequest) SetRegionId(v string) *DescribeAiEngineEventTrendsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAiEngineEventTrendsRequest) SetResourceManagerResourceGroupId(v string) *DescribeAiEngineEventTrendsRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeAiEngineEventTrendsRequest) SetStartTime(v int64) *DescribeAiEngineEventTrendsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeAiEngineEventTrendsRequest) Validate() error {
	return dara.Validate(s)
}
