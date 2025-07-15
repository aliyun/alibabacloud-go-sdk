// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSessionStatisticResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeSessionStatisticResponseBody
	GetRequestId() *string
	SetStatistic(v []*DescribeSessionStatisticResponseBodyStatistic) *DescribeSessionStatisticResponseBody
	GetStatistic() []*DescribeSessionStatisticResponseBodyStatistic
	SetTotalCount(v string) *DescribeSessionStatisticResponseBody
	GetTotalCount() *string
}

type DescribeSessionStatisticResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// C5528624-C6ED-5CA4-A4A2-7C30DBF2****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The statistics.
	Statistic []*DescribeSessionStatisticResponseBodyStatistic `json:"Statistic,omitempty" xml:"Statistic,omitempty" type:"Repeated"`
	// The total number of sessions returned.
	//
	// example:
	//
	// 1
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeSessionStatisticResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSessionStatisticResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSessionStatisticResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSessionStatisticResponseBody) GetStatistic() []*DescribeSessionStatisticResponseBodyStatistic {
	return s.Statistic
}

func (s *DescribeSessionStatisticResponseBody) GetTotalCount() *string {
	return s.TotalCount
}

func (s *DescribeSessionStatisticResponseBody) SetRequestId(v string) *DescribeSessionStatisticResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSessionStatisticResponseBody) SetStatistic(v []*DescribeSessionStatisticResponseBodyStatistic) *DescribeSessionStatisticResponseBody {
	s.Statistic = v
	return s
}

func (s *DescribeSessionStatisticResponseBody) SetTotalCount(v string) *DescribeSessionStatisticResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeSessionStatisticResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeSessionStatisticResponseBodyStatistic struct {
	// The total number of sessions in the time range.
	//
	// example:
	//
	// 4
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The point in time.
	//
	// example:
	//
	// 1690164443508
	TimePoint *int64 `json:"TimePoint,omitempty" xml:"TimePoint,omitempty"`
}

func (s DescribeSessionStatisticResponseBodyStatistic) String() string {
	return dara.Prettify(s)
}

func (s DescribeSessionStatisticResponseBodyStatistic) GoString() string {
	return s.String()
}

func (s *DescribeSessionStatisticResponseBodyStatistic) GetCount() *int64 {
	return s.Count
}

func (s *DescribeSessionStatisticResponseBodyStatistic) GetTimePoint() *int64 {
	return s.TimePoint
}

func (s *DescribeSessionStatisticResponseBodyStatistic) SetCount(v int64) *DescribeSessionStatisticResponseBodyStatistic {
	s.Count = &v
	return s
}

func (s *DescribeSessionStatisticResponseBodyStatistic) SetTimePoint(v int64) *DescribeSessionStatisticResponseBodyStatistic {
	s.TimePoint = &v
	return s
}

func (s *DescribeSessionStatisticResponseBodyStatistic) Validate() error {
	return dara.Validate(s)
}
