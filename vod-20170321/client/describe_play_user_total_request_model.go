// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePlayUserTotalRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribePlayUserTotalRequest
	GetEndTime() *string
	SetOwnerId(v int64) *DescribePlayUserTotalRequest
	GetOwnerId() *int64
	SetStartTime(v string) *DescribePlayUserTotalRequest
	GetStartTime() *string
}

type DescribePlayUserTotalRequest struct {
	// The end time of the query. Specify the time in the <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z format (UTC).
	//
	// > The end time must be later than the start time. The maximum time span between the start time and end time is 180 days.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2016-06-30T13:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The start time of the query. Specify the time in the yyyy-MM-ddTHH:mm:ssZ format (UTC).
	//
	// This parameter is required.
	//
	// example:
	//
	// 2016-06-29T13:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribePlayUserTotalRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePlayUserTotalRequest) GoString() string {
	return s.String()
}

func (s *DescribePlayUserTotalRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribePlayUserTotalRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribePlayUserTotalRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribePlayUserTotalRequest) SetEndTime(v string) *DescribePlayUserTotalRequest {
	s.EndTime = &v
	return s
}

func (s *DescribePlayUserTotalRequest) SetOwnerId(v int64) *DescribePlayUserTotalRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribePlayUserTotalRequest) SetStartTime(v string) *DescribePlayUserTotalRequest {
	s.StartTime = &v
	return s
}

func (s *DescribePlayUserTotalRequest) Validate() error {
	return dara.Validate(s)
}
