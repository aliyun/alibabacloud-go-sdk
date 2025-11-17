// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMetricTopRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAndroidInstanceIds(v []*string) *DescribeMetricTopRequest
	GetAndroidInstanceIds() []*string
	SetEndTime(v string) *DescribeMetricTopRequest
	GetEndTime() *string
	SetInstanceIds(v []*string) *DescribeMetricTopRequest
	GetInstanceIds() []*string
	SetLength(v string) *DescribeMetricTopRequest
	GetLength() *string
	SetMetricNames(v []*string) *DescribeMetricTopRequest
	GetMetricNames() []*string
	SetNextToken(v string) *DescribeMetricTopRequest
	GetNextToken() *string
	SetPeriod(v int32) *DescribeMetricTopRequest
	GetPeriod() *int32
	SetStartTime(v string) *DescribeMetricTopRequest
	GetStartTime() *string
}

type DescribeMetricTopRequest struct {
	AndroidInstanceIds []*string `json:"AndroidInstanceIds,omitempty" xml:"AndroidInstanceIds,omitempty" type:"Repeated"`
	// example:
	//
	// 2019-01-31 11:00:00
	EndTime     *string   `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// example:
	//
	// 100
	Length *string `json:"Length,omitempty" xml:"Length,omitempty"`
	// This parameter is required.
	MetricNames []*string `json:"MetricNames,omitempty" xml:"MetricNames,omitempty" type:"Repeated"`
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6l5V9uONHqPtDLM2U8s****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 60
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// example:
	//
	// 2019-01-31 10:00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeMetricTopRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetricTopRequest) GoString() string {
	return s.String()
}

func (s *DescribeMetricTopRequest) GetAndroidInstanceIds() []*string {
	return s.AndroidInstanceIds
}

func (s *DescribeMetricTopRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeMetricTopRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *DescribeMetricTopRequest) GetLength() *string {
	return s.Length
}

func (s *DescribeMetricTopRequest) GetMetricNames() []*string {
	return s.MetricNames
}

func (s *DescribeMetricTopRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeMetricTopRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *DescribeMetricTopRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeMetricTopRequest) SetAndroidInstanceIds(v []*string) *DescribeMetricTopRequest {
	s.AndroidInstanceIds = v
	return s
}

func (s *DescribeMetricTopRequest) SetEndTime(v string) *DescribeMetricTopRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeMetricTopRequest) SetInstanceIds(v []*string) *DescribeMetricTopRequest {
	s.InstanceIds = v
	return s
}

func (s *DescribeMetricTopRequest) SetLength(v string) *DescribeMetricTopRequest {
	s.Length = &v
	return s
}

func (s *DescribeMetricTopRequest) SetMetricNames(v []*string) *DescribeMetricTopRequest {
	s.MetricNames = v
	return s
}

func (s *DescribeMetricTopRequest) SetNextToken(v string) *DescribeMetricTopRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeMetricTopRequest) SetPeriod(v int32) *DescribeMetricTopRequest {
	s.Period = &v
	return s
}

func (s *DescribeMetricTopRequest) SetStartTime(v string) *DescribeMetricTopRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeMetricTopRequest) Validate() error {
	return dara.Validate(s)
}
