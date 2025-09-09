// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstDbLogInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLogTimeRange(v *DescribeInstDbLogInfoResponseBodyLogTimeRange) *DescribeInstDbLogInfoResponseBody
	GetLogTimeRange() *DescribeInstDbLogInfoResponseBodyLogTimeRange
	SetRequestId(v string) *DescribeInstDbLogInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeInstDbLogInfoResponseBody
	GetSuccess() *bool
}

type DescribeInstDbLogInfoResponseBody struct {
	// The time range for log query.
	LogTimeRange *DescribeInstDbLogInfoResponseBodyLogTimeRange `json:"LogTimeRange,omitempty" xml:"LogTimeRange,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// A5438952-70EE-4FA5-87A9-080DB0ASD45F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeInstDbLogInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstDbLogInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstDbLogInfoResponseBody) GetLogTimeRange() *DescribeInstDbLogInfoResponseBodyLogTimeRange {
	return s.LogTimeRange
}

func (s *DescribeInstDbLogInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInstDbLogInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeInstDbLogInfoResponseBody) SetLogTimeRange(v *DescribeInstDbLogInfoResponseBodyLogTimeRange) *DescribeInstDbLogInfoResponseBody {
	s.LogTimeRange = v
	return s
}

func (s *DescribeInstDbLogInfoResponseBody) SetRequestId(v string) *DescribeInstDbLogInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstDbLogInfoResponseBody) SetSuccess(v bool) *DescribeInstDbLogInfoResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeInstDbLogInfoResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeInstDbLogInfoResponseBodyLogTimeRange struct {
	// The start time of the query time range.
	//
	// example:
	//
	// 1568267711
	SupportLatestTime *int64 `json:"SupportLatestTime,omitempty" xml:"SupportLatestTime,omitempty"`
	// The end time of the task.
	//
	// example:
	//
	// 1568367711
	SupportOldestTime *int64 `json:"SupportOldestTime,omitempty" xml:"SupportOldestTime,omitempty"`
}

func (s DescribeInstDbLogInfoResponseBodyLogTimeRange) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstDbLogInfoResponseBodyLogTimeRange) GoString() string {
	return s.String()
}

func (s *DescribeInstDbLogInfoResponseBodyLogTimeRange) GetSupportLatestTime() *int64 {
	return s.SupportLatestTime
}

func (s *DescribeInstDbLogInfoResponseBodyLogTimeRange) GetSupportOldestTime() *int64 {
	return s.SupportOldestTime
}

func (s *DescribeInstDbLogInfoResponseBodyLogTimeRange) SetSupportLatestTime(v int64) *DescribeInstDbLogInfoResponseBodyLogTimeRange {
	s.SupportLatestTime = &v
	return s
}

func (s *DescribeInstDbLogInfoResponseBodyLogTimeRange) SetSupportOldestTime(v int64) *DescribeInstDbLogInfoResponseBodyLogTimeRange {
	s.SupportOldestTime = &v
	return s
}

func (s *DescribeInstDbLogInfoResponseBodyLogTimeRange) Validate() error {
	return dara.Validate(s)
}
