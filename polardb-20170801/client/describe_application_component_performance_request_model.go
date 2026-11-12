// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApplicationComponentPerformanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *DescribeApplicationComponentPerformanceRequest
	GetApplicationId() *string
	SetComponentId(v string) *DescribeApplicationComponentPerformanceRequest
	GetComponentId() *string
	SetEndTime(v string) *DescribeApplicationComponentPerformanceRequest
	GetEndTime() *string
	SetInterval(v string) *DescribeApplicationComponentPerformanceRequest
	GetInterval() *string
	SetKey(v string) *DescribeApplicationComponentPerformanceRequest
	GetKey() *string
	SetStartTime(v string) *DescribeApplicationComponentPerformanceRequest
	GetStartTime() *string
	SetType(v string) *DescribeApplicationComponentPerformanceRequest
	GetType() *string
}

type DescribeApplicationComponentPerformanceRequest struct {
	// The ID of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// pa-**************
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The ID of the application component. You can obtain the ID from the Components parameter returned by the DescribeApplicationAttribute operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// pac-*******************
	ComponentId *string `json:"ComponentId,omitempty" xml:"ComponentId,omitempty"`
	// The end time of the query. The end time must not be earlier than the start time. Specify the time in the ISO 8601 standard in the YYYY-MM-DDTHH:mmZ format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2020-09-23T02:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The sampling granularity of the performance data. Unit: seconds.
	//
	// example:
	//
	// 60
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The performance metrics to query. Separate multiple metrics with commas. You can specify up to 5 metrics.
	//
	// This parameter is required.
	//
	// example:
	//
	// PolarDBCPU
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The start time of the query. Specify the time in the ISO 8601 standard in the YYYY-MM-DDTHH:mmZ format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2020-09-23T01:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The type of performance data. Note: Add the valid values, the meaning of each value, and the correspondence with ApplicationType to the description of the Type parameter, and explicitly specify that this parameter should be omitted in most scenarios.
	//
	// example:
	//
	// supabase
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeApplicationComponentPerformanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationComponentPerformanceRequest) GoString() string {
	return s.String()
}

func (s *DescribeApplicationComponentPerformanceRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *DescribeApplicationComponentPerformanceRequest) GetComponentId() *string {
	return s.ComponentId
}

func (s *DescribeApplicationComponentPerformanceRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeApplicationComponentPerformanceRequest) GetInterval() *string {
	return s.Interval
}

func (s *DescribeApplicationComponentPerformanceRequest) GetKey() *string {
	return s.Key
}

func (s *DescribeApplicationComponentPerformanceRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeApplicationComponentPerformanceRequest) GetType() *string {
	return s.Type
}

func (s *DescribeApplicationComponentPerformanceRequest) SetApplicationId(v string) *DescribeApplicationComponentPerformanceRequest {
	s.ApplicationId = &v
	return s
}

func (s *DescribeApplicationComponentPerformanceRequest) SetComponentId(v string) *DescribeApplicationComponentPerformanceRequest {
	s.ComponentId = &v
	return s
}

func (s *DescribeApplicationComponentPerformanceRequest) SetEndTime(v string) *DescribeApplicationComponentPerformanceRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeApplicationComponentPerformanceRequest) SetInterval(v string) *DescribeApplicationComponentPerformanceRequest {
	s.Interval = &v
	return s
}

func (s *DescribeApplicationComponentPerformanceRequest) SetKey(v string) *DescribeApplicationComponentPerformanceRequest {
	s.Key = &v
	return s
}

func (s *DescribeApplicationComponentPerformanceRequest) SetStartTime(v string) *DescribeApplicationComponentPerformanceRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeApplicationComponentPerformanceRequest) SetType(v string) *DescribeApplicationComponentPerformanceRequest {
	s.Type = &v
	return s
}

func (s *DescribeApplicationComponentPerformanceRequest) Validate() error {
	return dara.Validate(s)
}
