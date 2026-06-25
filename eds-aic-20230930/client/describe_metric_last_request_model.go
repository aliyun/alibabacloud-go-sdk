// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMetricLastRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAndroidInstanceIds(v []*string) *DescribeMetricLastRequest
	GetAndroidInstanceIds() []*string
	SetEndTime(v string) *DescribeMetricLastRequest
	GetEndTime() *string
	SetInstanceIds(v []*string) *DescribeMetricLastRequest
	GetInstanceIds() []*string
	SetLength(v string) *DescribeMetricLastRequest
	GetLength() *string
	SetMetricNames(v []*string) *DescribeMetricLastRequest
	GetMetricNames() []*string
	SetNextToken(v string) *DescribeMetricLastRequest
	GetNextToken() *string
	SetPeriod(v int32) *DescribeMetricLastRequest
	GetPeriod() *int32
	SetStartTime(v string) *DescribeMetricLastRequest
	GetStartTime() *string
}

type DescribeMetricLastRequest struct {
	// For instance-level monitoring, specify a list of Cloud Phone instance IDs.
	//
	// <props="china">
	//
	// For matrix-level monitoring, specify a list of matrix IDs.
	AndroidInstanceIds []*string `json:"AndroidInstanceIds,omitempty" xml:"AndroidInstanceIds,omitempty" type:"Repeated"`
	// The end time for the query.
	//
	// example:
	//
	// 2019-01-31 11:00:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The list of instance IDs.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The number of monitoring data records to return on each page for a paged query.
	//
	// The default value is 1000. This means that 1000 monitoring data records are returned per page.
	//
	// example:
	//
	// 1000
	Length *string `json:"Length,omitempty" xml:"Length,omitempty"`
	// The list of monitoring metrics.
	//
	// This parameter is required.
	MetricNames []*string `json:"MetricNames,omitempty" xml:"MetricNames,omitempty" type:"Repeated"`
	// The token that marks the start of the current read position. If you leave this parameter empty, the query starts from the beginning.
	//
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6l5V9uONHqPtDLM2U8s****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The statistic period for monitoring data.
	//
	// Valid values: 60, 900, and 3600.
	//
	// Unit: seconds.
	//
	// example:
	//
	// 60
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The start time for the query.
	//
	// example:
	//
	// 2019-01-31 10:00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeMetricLastRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetricLastRequest) GoString() string {
	return s.String()
}

func (s *DescribeMetricLastRequest) GetAndroidInstanceIds() []*string {
	return s.AndroidInstanceIds
}

func (s *DescribeMetricLastRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeMetricLastRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *DescribeMetricLastRequest) GetLength() *string {
	return s.Length
}

func (s *DescribeMetricLastRequest) GetMetricNames() []*string {
	return s.MetricNames
}

func (s *DescribeMetricLastRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeMetricLastRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *DescribeMetricLastRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeMetricLastRequest) SetAndroidInstanceIds(v []*string) *DescribeMetricLastRequest {
	s.AndroidInstanceIds = v
	return s
}

func (s *DescribeMetricLastRequest) SetEndTime(v string) *DescribeMetricLastRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeMetricLastRequest) SetInstanceIds(v []*string) *DescribeMetricLastRequest {
	s.InstanceIds = v
	return s
}

func (s *DescribeMetricLastRequest) SetLength(v string) *DescribeMetricLastRequest {
	s.Length = &v
	return s
}

func (s *DescribeMetricLastRequest) SetMetricNames(v []*string) *DescribeMetricLastRequest {
	s.MetricNames = v
	return s
}

func (s *DescribeMetricLastRequest) SetNextToken(v string) *DescribeMetricLastRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeMetricLastRequest) SetPeriod(v int32) *DescribeMetricLastRequest {
	s.Period = &v
	return s
}

func (s *DescribeMetricLastRequest) SetStartTime(v string) *DescribeMetricLastRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeMetricLastRequest) Validate() error {
	return dara.Validate(s)
}
