// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceMonitorDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeInstanceMonitorDataRequest
	GetEndTime() *string
	SetInstanceId(v string) *DescribeInstanceMonitorDataRequest
	GetInstanceId() *string
	SetPeriod(v string) *DescribeInstanceMonitorDataRequest
	GetPeriod() *string
	SetStartTime(v string) *DescribeInstanceMonitorDataRequest
	GetStartTime() *string
}

type DescribeInstanceMonitorDataRequest struct {
	// The end time of the data to retrieve. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC. If the seconds are not 00, the time is automatically rounded up to the next minute.
	//
	// example:
	//
	// 2019-10-30T08:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The instance ID. Only a single instance ID is supported.
	//
	// example:
	//
	// yourInstance ID
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The precision of the monitoring data to retrieve. Unit: seconds. Valid values: 60 (default), 300, 1200, 3600, and 14400. This parameter is actually required. If this parameter is not specified, the service returns a MissingParameter error.
	//
	// example:
	//
	// 60
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// The start time of the data to retrieve. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC. If the seconds are not 00, the time is automatically rounded up to the next minute.
	//
	// example:
	//
	// 2019-10-29T23:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeInstanceMonitorDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceMonitorDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceMonitorDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeInstanceMonitorDataRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeInstanceMonitorDataRequest) GetPeriod() *string {
	return s.Period
}

func (s *DescribeInstanceMonitorDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeInstanceMonitorDataRequest) SetEndTime(v string) *DescribeInstanceMonitorDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeInstanceMonitorDataRequest) SetInstanceId(v string) *DescribeInstanceMonitorDataRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceMonitorDataRequest) SetPeriod(v string) *DescribeInstanceMonitorDataRequest {
	s.Period = &v
	return s
}

func (s *DescribeInstanceMonitorDataRequest) SetStartTime(v string) *DescribeInstanceMonitorDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeInstanceMonitorDataRequest) Validate() error {
	return dara.Validate(s)
}
