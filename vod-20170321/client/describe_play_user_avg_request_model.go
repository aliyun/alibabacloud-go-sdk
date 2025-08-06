// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePlayUserAvgRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribePlayUserAvgRequest
	GetEndTime() *string
	SetOwnerId(v int64) *DescribePlayUserAvgRequest
	GetOwnerId() *int64
	SetStartTime(v string) *DescribePlayUserAvgRequest
	GetStartTime() *string
}

type DescribePlayUserAvgRequest struct {
	// The end of the time range to query. The end time must be later than the start time. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2016-06-30T13:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2016-06-29T13:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribePlayUserAvgRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePlayUserAvgRequest) GoString() string {
	return s.String()
}

func (s *DescribePlayUserAvgRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribePlayUserAvgRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribePlayUserAvgRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribePlayUserAvgRequest) SetEndTime(v string) *DescribePlayUserAvgRequest {
	s.EndTime = &v
	return s
}

func (s *DescribePlayUserAvgRequest) SetOwnerId(v int64) *DescribePlayUserAvgRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribePlayUserAvgRequest) SetStartTime(v string) *DescribePlayUserAvgRequest {
	s.StartTime = &v
	return s
}

func (s *DescribePlayUserAvgRequest) Validate() error {
	return dara.Validate(s)
}
