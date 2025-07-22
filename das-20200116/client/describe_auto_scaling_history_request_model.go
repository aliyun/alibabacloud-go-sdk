// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAutoScalingHistoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoScalingTaskType(v string) *DescribeAutoScalingHistoryRequest
	GetAutoScalingTaskType() *string
	SetEndTime(v int64) *DescribeAutoScalingHistoryRequest
	GetEndTime() *int64
	SetInstanceId(v string) *DescribeAutoScalingHistoryRequest
	GetInstanceId() *string
	SetStartTime(v int64) *DescribeAutoScalingHistoryRequest
	GetStartTime() *int64
}

type DescribeAutoScalingHistoryRequest struct {
	// The type of the auto scaling task that you want to query. Set the value to **SPEC**, which indicates that you can query the history of only automatic performance scaling tasks.
	//
	// This parameter is required.
	//
	// example:
	//
	// SPEC
	AutoScalingTaskType *string `json:"AutoScalingTaskType,omitempty" xml:"AutoScalingTaskType,omitempty"`
	// The end of the time range to query. Set this parameter to a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// > The end time must be later than the start time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1676605305796
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The instance ID.
	//
	// > Only ApsaraDB RDS for MySQL instances are supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-2ze8g2am97624****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The beginning of the time range to query. Set this parameter to a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// > The maximum time range that can be specified is 45 days.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1675833788056
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeAutoScalingHistoryRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutoScalingHistoryRequest) GoString() string {
	return s.String()
}

func (s *DescribeAutoScalingHistoryRequest) GetAutoScalingTaskType() *string {
	return s.AutoScalingTaskType
}

func (s *DescribeAutoScalingHistoryRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeAutoScalingHistoryRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeAutoScalingHistoryRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeAutoScalingHistoryRequest) SetAutoScalingTaskType(v string) *DescribeAutoScalingHistoryRequest {
	s.AutoScalingTaskType = &v
	return s
}

func (s *DescribeAutoScalingHistoryRequest) SetEndTime(v int64) *DescribeAutoScalingHistoryRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeAutoScalingHistoryRequest) SetInstanceId(v string) *DescribeAutoScalingHistoryRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeAutoScalingHistoryRequest) SetStartTime(v int64) *DescribeAutoScalingHistoryRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeAutoScalingHistoryRequest) Validate() error {
	return dara.Validate(s)
}
