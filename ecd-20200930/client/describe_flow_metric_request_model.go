// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFlowMetricRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeFlowMetricRequest
	GetEndTime() *string
	SetInstanceId(v string) *DescribeFlowMetricRequest
	GetInstanceId() *string
	SetInstanceType(v string) *DescribeFlowMetricRequest
	GetInstanceType() *string
	SetMetricType(v string) *DescribeFlowMetricRequest
	GetMetricType() *string
	SetPeriod(v int32) *DescribeFlowMetricRequest
	GetPeriod() *int32
	SetRegionId(v string) *DescribeFlowMetricRequest
	GetRegionId() *string
	SetStartTime(v string) *DescribeFlowMetricRequest
	GetStartTime() *string
}

type DescribeFlowMetricRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1664714703743
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// desktop
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// intranetOutRate
	MetricType *string `json:"MetricType,omitempty" xml:"MetricType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 60
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1651817220643
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeFlowMetricRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeFlowMetricRequest) GoString() string {
	return s.String()
}

func (s *DescribeFlowMetricRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeFlowMetricRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeFlowMetricRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeFlowMetricRequest) GetMetricType() *string {
	return s.MetricType
}

func (s *DescribeFlowMetricRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *DescribeFlowMetricRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeFlowMetricRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeFlowMetricRequest) SetEndTime(v string) *DescribeFlowMetricRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeFlowMetricRequest) SetInstanceId(v string) *DescribeFlowMetricRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeFlowMetricRequest) SetInstanceType(v string) *DescribeFlowMetricRequest {
	s.InstanceType = &v
	return s
}

func (s *DescribeFlowMetricRequest) SetMetricType(v string) *DescribeFlowMetricRequest {
	s.MetricType = &v
	return s
}

func (s *DescribeFlowMetricRequest) SetPeriod(v int32) *DescribeFlowMetricRequest {
	s.Period = &v
	return s
}

func (s *DescribeFlowMetricRequest) SetRegionId(v string) *DescribeFlowMetricRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeFlowMetricRequest) SetStartTime(v string) *DescribeFlowMetricRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeFlowMetricRequest) Validate() error {
	return dara.Validate(s)
}
