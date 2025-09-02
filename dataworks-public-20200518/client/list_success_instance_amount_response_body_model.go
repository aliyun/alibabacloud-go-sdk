// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSuccessInstanceAmountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceStatusTrend(v *ListSuccessInstanceAmountResponseBodyInstanceStatusTrend) *ListSuccessInstanceAmountResponseBody
	GetInstanceStatusTrend() *ListSuccessInstanceAmountResponseBodyInstanceStatusTrend
	SetRequestId(v string) *ListSuccessInstanceAmountResponseBody
	GetRequestId() *string
}

type ListSuccessInstanceAmountResponseBody struct {
	// Indicates the trend of the number of auto triggered node instances that are successfully run every hour on the hour of the current day.
	InstanceStatusTrend *ListSuccessInstanceAmountResponseBodyInstanceStatusTrend `json:"InstanceStatusTrend,omitempty" xml:"InstanceStatusTrend,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 952795279527ab****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListSuccessInstanceAmountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSuccessInstanceAmountResponseBody) GoString() string {
	return s.String()
}

func (s *ListSuccessInstanceAmountResponseBody) GetInstanceStatusTrend() *ListSuccessInstanceAmountResponseBodyInstanceStatusTrend {
	return s.InstanceStatusTrend
}

func (s *ListSuccessInstanceAmountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSuccessInstanceAmountResponseBody) SetInstanceStatusTrend(v *ListSuccessInstanceAmountResponseBodyInstanceStatusTrend) *ListSuccessInstanceAmountResponseBody {
	s.InstanceStatusTrend = v
	return s
}

func (s *ListSuccessInstanceAmountResponseBody) SetRequestId(v string) *ListSuccessInstanceAmountResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSuccessInstanceAmountResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListSuccessInstanceAmountResponseBodyInstanceStatusTrend struct {
	// The average trend.
	AvgTrend []*ListSuccessInstanceAmountResponseBodyInstanceStatusTrendAvgTrend `json:"AvgTrend,omitempty" xml:"AvgTrend,omitempty" type:"Repeated"`
	// The trend of the number of auto triggered node instances that are successfully run on the current day.
	TodayTrend []*ListSuccessInstanceAmountResponseBodyInstanceStatusTrendTodayTrend `json:"TodayTrend,omitempty" xml:"TodayTrend,omitempty" type:"Repeated"`
	// The trend of the number of auto triggered node instances that are successfully run one day earlier than the current day.
	YesterdayTrend []*ListSuccessInstanceAmountResponseBodyInstanceStatusTrendYesterdayTrend `json:"YesterdayTrend,omitempty" xml:"YesterdayTrend,omitempty" type:"Repeated"`
}

func (s ListSuccessInstanceAmountResponseBodyInstanceStatusTrend) String() string {
	return dara.Prettify(s)
}

func (s ListSuccessInstanceAmountResponseBodyInstanceStatusTrend) GoString() string {
	return s.String()
}

func (s *ListSuccessInstanceAmountResponseBodyInstanceStatusTrend) GetAvgTrend() []*ListSuccessInstanceAmountResponseBodyInstanceStatusTrendAvgTrend {
	return s.AvgTrend
}

func (s *ListSuccessInstanceAmountResponseBodyInstanceStatusTrend) GetTodayTrend() []*ListSuccessInstanceAmountResponseBodyInstanceStatusTrendTodayTrend {
	return s.TodayTrend
}

func (s *ListSuccessInstanceAmountResponseBodyInstanceStatusTrend) GetYesterdayTrend() []*ListSuccessInstanceAmountResponseBodyInstanceStatusTrendYesterdayTrend {
	return s.YesterdayTrend
}

func (s *ListSuccessInstanceAmountResponseBodyInstanceStatusTrend) SetAvgTrend(v []*ListSuccessInstanceAmountResponseBodyInstanceStatusTrendAvgTrend) *ListSuccessInstanceAmountResponseBodyInstanceStatusTrend {
	s.AvgTrend = v
	return s
}

func (s *ListSuccessInstanceAmountResponseBodyInstanceStatusTrend) SetTodayTrend(v []*ListSuccessInstanceAmountResponseBodyInstanceStatusTrendTodayTrend) *ListSuccessInstanceAmountResponseBodyInstanceStatusTrend {
	s.TodayTrend = v
	return s
}

func (s *ListSuccessInstanceAmountResponseBodyInstanceStatusTrend) SetYesterdayTrend(v []*ListSuccessInstanceAmountResponseBodyInstanceStatusTrendYesterdayTrend) *ListSuccessInstanceAmountResponseBodyInstanceStatusTrend {
	s.YesterdayTrend = v
	return s
}

func (s *ListSuccessInstanceAmountResponseBodyInstanceStatusTrend) Validate() error {
	return dara.Validate(s)
}

type ListSuccessInstanceAmountResponseBodyInstanceStatusTrendAvgTrend struct {
	// The number of instances that are successfully run.
	//
	// example:
	//
	// 10
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The point in time. The value is an exact hour that ranges from 00:00 to 23:00, such as 00:00, 01:00, or 02:00.
	//
	// example:
	//
	// 01:00
	TimePoint *string `json:"TimePoint,omitempty" xml:"TimePoint,omitempty"`
}

func (s ListSuccessInstanceAmountResponseBodyInstanceStatusTrendAvgTrend) String() string {
	return dara.Prettify(s)
}

func (s ListSuccessInstanceAmountResponseBodyInstanceStatusTrendAvgTrend) GoString() string {
	return s.String()
}

func (s *ListSuccessInstanceAmountResponseBodyInstanceStatusTrendAvgTrend) GetCount() *int32 {
	return s.Count
}

func (s *ListSuccessInstanceAmountResponseBodyInstanceStatusTrendAvgTrend) GetTimePoint() *string {
	return s.TimePoint
}

func (s *ListSuccessInstanceAmountResponseBodyInstanceStatusTrendAvgTrend) SetCount(v int32) *ListSuccessInstanceAmountResponseBodyInstanceStatusTrendAvgTrend {
	s.Count = &v
	return s
}

func (s *ListSuccessInstanceAmountResponseBodyInstanceStatusTrendAvgTrend) SetTimePoint(v string) *ListSuccessInstanceAmountResponseBodyInstanceStatusTrendAvgTrend {
	s.TimePoint = &v
	return s
}

func (s *ListSuccessInstanceAmountResponseBodyInstanceStatusTrendAvgTrend) Validate() error {
	return dara.Validate(s)
}

type ListSuccessInstanceAmountResponseBodyInstanceStatusTrendTodayTrend struct {
	// The number of instances that are successfully run.
	//
	// example:
	//
	// 10
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The point in time. The value is an exact hour that ranges from 00:00 to 23:00, such as 00:00, 01:00, or 02:00.
	//
	// example:
	//
	// 01:00
	TimePoint *string `json:"TimePoint,omitempty" xml:"TimePoint,omitempty"`
}

func (s ListSuccessInstanceAmountResponseBodyInstanceStatusTrendTodayTrend) String() string {
	return dara.Prettify(s)
}

func (s ListSuccessInstanceAmountResponseBodyInstanceStatusTrendTodayTrend) GoString() string {
	return s.String()
}

func (s *ListSuccessInstanceAmountResponseBodyInstanceStatusTrendTodayTrend) GetCount() *int32 {
	return s.Count
}

func (s *ListSuccessInstanceAmountResponseBodyInstanceStatusTrendTodayTrend) GetTimePoint() *string {
	return s.TimePoint
}

func (s *ListSuccessInstanceAmountResponseBodyInstanceStatusTrendTodayTrend) SetCount(v int32) *ListSuccessInstanceAmountResponseBodyInstanceStatusTrendTodayTrend {
	s.Count = &v
	return s
}

func (s *ListSuccessInstanceAmountResponseBodyInstanceStatusTrendTodayTrend) SetTimePoint(v string) *ListSuccessInstanceAmountResponseBodyInstanceStatusTrendTodayTrend {
	s.TimePoint = &v
	return s
}

func (s *ListSuccessInstanceAmountResponseBodyInstanceStatusTrendTodayTrend) Validate() error {
	return dara.Validate(s)
}

type ListSuccessInstanceAmountResponseBodyInstanceStatusTrendYesterdayTrend struct {
	// The number of instances that are successfully run.
	//
	// example:
	//
	// 10
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The point in time. The value is an exact hour that ranges from 00:00 to 23:00, such as 00:00, 01:00, or 02:00.
	//
	// example:
	//
	// 01:00
	TimePoint *string `json:"TimePoint,omitempty" xml:"TimePoint,omitempty"`
}

func (s ListSuccessInstanceAmountResponseBodyInstanceStatusTrendYesterdayTrend) String() string {
	return dara.Prettify(s)
}

func (s ListSuccessInstanceAmountResponseBodyInstanceStatusTrendYesterdayTrend) GoString() string {
	return s.String()
}

func (s *ListSuccessInstanceAmountResponseBodyInstanceStatusTrendYesterdayTrend) GetCount() *int32 {
	return s.Count
}

func (s *ListSuccessInstanceAmountResponseBodyInstanceStatusTrendYesterdayTrend) GetTimePoint() *string {
	return s.TimePoint
}

func (s *ListSuccessInstanceAmountResponseBodyInstanceStatusTrendYesterdayTrend) SetCount(v int32) *ListSuccessInstanceAmountResponseBodyInstanceStatusTrendYesterdayTrend {
	s.Count = &v
	return s
}

func (s *ListSuccessInstanceAmountResponseBodyInstanceStatusTrendYesterdayTrend) SetTimePoint(v string) *ListSuccessInstanceAmountResponseBodyInstanceStatusTrendYesterdayTrend {
	s.TimePoint = &v
	return s
}

func (s *ListSuccessInstanceAmountResponseBodyInstanceStatusTrendYesterdayTrend) Validate() error {
	return dara.Validate(s)
}
