// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApplicationHealthCheckReportsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationIds(v []*string) *DescribeApplicationHealthCheckReportsRequest
	GetApplicationIds() []*string
	SetEndTime(v int64) *DescribeApplicationHealthCheckReportsRequest
	GetEndTime() *int64
	SetStartTime(v int64) *DescribeApplicationHealthCheckReportsRequest
	GetStartTime() *int64
}

type DescribeApplicationHealthCheckReportsRequest struct {
	// The IDs of the internal-facing applications. You can specify up to 100 application IDs.
	//
	// This parameter is required.
	ApplicationIds []*string `json:"ApplicationIds,omitempty" xml:"ApplicationIds,omitempty" type:"Repeated"`
	// The end time of the query. The value is a UNIX timestamp in seconds. If this parameter is not specified, the default value is the current time. If only the start time is specified, the end time is set to 5 minutes after the start time. The end time must be later than the start time, and the maximum query time range is 7 days.
	//
	// example:
	//
	// 1700000600
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The start time of the query. The value is a UNIX timestamp in seconds. If this parameter is not specified, the default value is 5 minutes before the end time. If only the end time is specified, the start time is set to 5 minutes before the end time.
	//
	// example:
	//
	// 1700000000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeApplicationHealthCheckReportsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationHealthCheckReportsRequest) GoString() string {
	return s.String()
}

func (s *DescribeApplicationHealthCheckReportsRequest) GetApplicationIds() []*string {
	return s.ApplicationIds
}

func (s *DescribeApplicationHealthCheckReportsRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeApplicationHealthCheckReportsRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeApplicationHealthCheckReportsRequest) SetApplicationIds(v []*string) *DescribeApplicationHealthCheckReportsRequest {
	s.ApplicationIds = v
	return s
}

func (s *DescribeApplicationHealthCheckReportsRequest) SetEndTime(v int64) *DescribeApplicationHealthCheckReportsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeApplicationHealthCheckReportsRequest) SetStartTime(v int64) *DescribeApplicationHealthCheckReportsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeApplicationHealthCheckReportsRequest) Validate() error {
	return dara.Validate(s)
}
