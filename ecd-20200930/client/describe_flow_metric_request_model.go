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
	// End Time. Supported formats:
	//
	// - UNIX timestamp: the number of milliseconds elapsed since January 1, 1970.
	//
	// - Format: YYYY-MM-DDThh:mm:ssZ.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1664714703743
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The instance ID, which can be either a cloud computr ID or a premium public bandwidth plan ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// Ecd-fwq23f13****ornp-6inxqsvcyv6z8****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The instance type. You can select either cloud computer type or premium public bandwidth type. If you select cloud computer type, the `InstanceId` and `MetricType` must be filled in with a cloud computer ID and its corresponding traffic type. The same applies to premium public bandwidth.
	//
	// This parameter is required.
	//
	// example:
	//
	// desktop
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The type of monitoring metric. Supports monitoring data for inbound and outbound bandwidth of cloud computers, as well as inbound and outbound bandwidth for public network access of premium public bandwidth.
	//
	// This parameter is required.
	//
	// example:
	//
	// intranetOutRate
	MetricType *string `json:"MetricType,omitempty" xml:"MetricType,omitempty"`
	// The statistic period of monitoring data. Unit: seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 60
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The Region ID. You can call [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) to obtain the list of Regions supported by Elastic Desktop Service (EDS).
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Start Time. Supported formats:
	//
	// - UNIX timestamp: the number of milliseconds elapsed since January 1, 1970.
	//
	// - Format: YYYY-MM-DDThh:mm:ssZ.
	//
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
