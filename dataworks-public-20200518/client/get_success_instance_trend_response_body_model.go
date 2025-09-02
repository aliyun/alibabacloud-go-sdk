// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSuccessInstanceTrendResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceStatusTrend(v *GetSuccessInstanceTrendResponseBodyInstanceStatusTrend) *GetSuccessInstanceTrendResponseBody
	GetInstanceStatusTrend() *GetSuccessInstanceTrendResponseBodyInstanceStatusTrend
	SetRequestId(v string) *GetSuccessInstanceTrendResponseBody
	GetRequestId() *string
}

type GetSuccessInstanceTrendResponseBody struct {
	// The trend of statistics on the instance status in different time periods.
	InstanceStatusTrend *GetSuccessInstanceTrendResponseBodyInstanceStatusTrend `json:"InstanceStatusTrend,omitempty" xml:"InstanceStatusTrend,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 952795279527ab****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetSuccessInstanceTrendResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSuccessInstanceTrendResponseBody) GoString() string {
	return s.String()
}

func (s *GetSuccessInstanceTrendResponseBody) GetInstanceStatusTrend() *GetSuccessInstanceTrendResponseBodyInstanceStatusTrend {
	return s.InstanceStatusTrend
}

func (s *GetSuccessInstanceTrendResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSuccessInstanceTrendResponseBody) SetInstanceStatusTrend(v *GetSuccessInstanceTrendResponseBodyInstanceStatusTrend) *GetSuccessInstanceTrendResponseBody {
	s.InstanceStatusTrend = v
	return s
}

func (s *GetSuccessInstanceTrendResponseBody) SetRequestId(v string) *GetSuccessInstanceTrendResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSuccessInstanceTrendResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetSuccessInstanceTrendResponseBodyInstanceStatusTrend struct {
	// The average trend.
	AvgTrend []*GetSuccessInstanceTrendResponseBodyInstanceStatusTrendAvgTrend `json:"AvgTrend,omitempty" xml:"AvgTrend,omitempty" type:"Repeated"`
	// The trend on the current day.
	TodayTrend []*GetSuccessInstanceTrendResponseBodyInstanceStatusTrendTodayTrend `json:"TodayTrend,omitempty" xml:"TodayTrend,omitempty" type:"Repeated"`
	// The trend on the previous day.
	YesterdayTrend []*GetSuccessInstanceTrendResponseBodyInstanceStatusTrendYesterdayTrend `json:"YesterdayTrend,omitempty" xml:"YesterdayTrend,omitempty" type:"Repeated"`
}

func (s GetSuccessInstanceTrendResponseBodyInstanceStatusTrend) String() string {
	return dara.Prettify(s)
}

func (s GetSuccessInstanceTrendResponseBodyInstanceStatusTrend) GoString() string {
	return s.String()
}

func (s *GetSuccessInstanceTrendResponseBodyInstanceStatusTrend) GetAvgTrend() []*GetSuccessInstanceTrendResponseBodyInstanceStatusTrendAvgTrend {
	return s.AvgTrend
}

func (s *GetSuccessInstanceTrendResponseBodyInstanceStatusTrend) GetTodayTrend() []*GetSuccessInstanceTrendResponseBodyInstanceStatusTrendTodayTrend {
	return s.TodayTrend
}

func (s *GetSuccessInstanceTrendResponseBodyInstanceStatusTrend) GetYesterdayTrend() []*GetSuccessInstanceTrendResponseBodyInstanceStatusTrendYesterdayTrend {
	return s.YesterdayTrend
}

func (s *GetSuccessInstanceTrendResponseBodyInstanceStatusTrend) SetAvgTrend(v []*GetSuccessInstanceTrendResponseBodyInstanceStatusTrendAvgTrend) *GetSuccessInstanceTrendResponseBodyInstanceStatusTrend {
	s.AvgTrend = v
	return s
}

func (s *GetSuccessInstanceTrendResponseBodyInstanceStatusTrend) SetTodayTrend(v []*GetSuccessInstanceTrendResponseBodyInstanceStatusTrendTodayTrend) *GetSuccessInstanceTrendResponseBodyInstanceStatusTrend {
	s.TodayTrend = v
	return s
}

func (s *GetSuccessInstanceTrendResponseBodyInstanceStatusTrend) SetYesterdayTrend(v []*GetSuccessInstanceTrendResponseBodyInstanceStatusTrendYesterdayTrend) *GetSuccessInstanceTrendResponseBodyInstanceStatusTrend {
	s.YesterdayTrend = v
	return s
}

func (s *GetSuccessInstanceTrendResponseBodyInstanceStatusTrend) Validate() error {
	return dara.Validate(s)
}

type GetSuccessInstanceTrendResponseBodyInstanceStatusTrendAvgTrend struct {
	// The number of instances.
	//
	// example:
	//
	// 10
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The point in time. Valid values: 00:00 to 23:00.
	//
	// example:
	//
	// 01:00
	TimePoint *string `json:"TimePoint,omitempty" xml:"TimePoint,omitempty"`
}

func (s GetSuccessInstanceTrendResponseBodyInstanceStatusTrendAvgTrend) String() string {
	return dara.Prettify(s)
}

func (s GetSuccessInstanceTrendResponseBodyInstanceStatusTrendAvgTrend) GoString() string {
	return s.String()
}

func (s *GetSuccessInstanceTrendResponseBodyInstanceStatusTrendAvgTrend) GetCount() *int32 {
	return s.Count
}

func (s *GetSuccessInstanceTrendResponseBodyInstanceStatusTrendAvgTrend) GetTimePoint() *string {
	return s.TimePoint
}

func (s *GetSuccessInstanceTrendResponseBodyInstanceStatusTrendAvgTrend) SetCount(v int32) *GetSuccessInstanceTrendResponseBodyInstanceStatusTrendAvgTrend {
	s.Count = &v
	return s
}

func (s *GetSuccessInstanceTrendResponseBodyInstanceStatusTrendAvgTrend) SetTimePoint(v string) *GetSuccessInstanceTrendResponseBodyInstanceStatusTrendAvgTrend {
	s.TimePoint = &v
	return s
}

func (s *GetSuccessInstanceTrendResponseBodyInstanceStatusTrendAvgTrend) Validate() error {
	return dara.Validate(s)
}

type GetSuccessInstanceTrendResponseBodyInstanceStatusTrendTodayTrend struct {
	// The number of instances.
	//
	// example:
	//
	// 10
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The point in time. Valid values: 00:00 to 23:00.
	//
	// example:
	//
	// 01:00
	TimePoint *string `json:"TimePoint,omitempty" xml:"TimePoint,omitempty"`
}

func (s GetSuccessInstanceTrendResponseBodyInstanceStatusTrendTodayTrend) String() string {
	return dara.Prettify(s)
}

func (s GetSuccessInstanceTrendResponseBodyInstanceStatusTrendTodayTrend) GoString() string {
	return s.String()
}

func (s *GetSuccessInstanceTrendResponseBodyInstanceStatusTrendTodayTrend) GetCount() *int32 {
	return s.Count
}

func (s *GetSuccessInstanceTrendResponseBodyInstanceStatusTrendTodayTrend) GetTimePoint() *string {
	return s.TimePoint
}

func (s *GetSuccessInstanceTrendResponseBodyInstanceStatusTrendTodayTrend) SetCount(v int32) *GetSuccessInstanceTrendResponseBodyInstanceStatusTrendTodayTrend {
	s.Count = &v
	return s
}

func (s *GetSuccessInstanceTrendResponseBodyInstanceStatusTrendTodayTrend) SetTimePoint(v string) *GetSuccessInstanceTrendResponseBodyInstanceStatusTrendTodayTrend {
	s.TimePoint = &v
	return s
}

func (s *GetSuccessInstanceTrendResponseBodyInstanceStatusTrendTodayTrend) Validate() error {
	return dara.Validate(s)
}

type GetSuccessInstanceTrendResponseBodyInstanceStatusTrendYesterdayTrend struct {
	// The number of instances.
	//
	// example:
	//
	// 10
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The point in time. Valid values: 00:00 to 23:00.
	//
	// example:
	//
	// 01:00
	TimePoint *string `json:"TimePoint,omitempty" xml:"TimePoint,omitempty"`
}

func (s GetSuccessInstanceTrendResponseBodyInstanceStatusTrendYesterdayTrend) String() string {
	return dara.Prettify(s)
}

func (s GetSuccessInstanceTrendResponseBodyInstanceStatusTrendYesterdayTrend) GoString() string {
	return s.String()
}

func (s *GetSuccessInstanceTrendResponseBodyInstanceStatusTrendYesterdayTrend) GetCount() *int32 {
	return s.Count
}

func (s *GetSuccessInstanceTrendResponseBodyInstanceStatusTrendYesterdayTrend) GetTimePoint() *string {
	return s.TimePoint
}

func (s *GetSuccessInstanceTrendResponseBodyInstanceStatusTrendYesterdayTrend) SetCount(v int32) *GetSuccessInstanceTrendResponseBodyInstanceStatusTrendYesterdayTrend {
	s.Count = &v
	return s
}

func (s *GetSuccessInstanceTrendResponseBodyInstanceStatusTrendYesterdayTrend) SetTimePoint(v string) *GetSuccessInstanceTrendResponseBodyInstanceStatusTrendYesterdayTrend {
	s.TimePoint = &v
	return s
}

func (s *GetSuccessInstanceTrendResponseBodyInstanceStatusTrendYesterdayTrend) Validate() error {
	return dara.Validate(s)
}
