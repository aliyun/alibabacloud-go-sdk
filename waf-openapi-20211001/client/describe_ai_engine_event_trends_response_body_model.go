// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAiEngineEventTrendsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAggregateInterval(v string) *DescribeAiEngineEventTrendsResponseBody
	GetAggregateInterval() *string
	SetPoints(v []*DescribeAiEngineEventTrendsResponseBodyPoints) *DescribeAiEngineEventTrendsResponseBody
	GetPoints() []*DescribeAiEngineEventTrendsResponseBodyPoints
	SetRequestId(v string) *DescribeAiEngineEventTrendsResponseBody
	GetRequestId() *string
}

type DescribeAiEngineEventTrendsResponseBody struct {
	// The aggregation interval, which is automatically selected based on the query time range. Valid values:
	//
	// - **1m**: 1 minute
	//
	// - **5m**: 5 minutes
	//
	// - **15m**: 15 minutes
	//
	// - **1h**: 1 hour
	//
	// - **1d**: 1 day
	//
	// example:
	//
	// 1m
	AggregateInterval *string `json:"AggregateInterval,omitempty" xml:"AggregateInterval,omitempty"`
	// The time series data of events aggregated by first discovery time, sorted in ascending order of time.
	Points []*DescribeAiEngineEventTrendsResponseBodyPoints `json:"Points,omitempty" xml:"Points,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 01A****-1219-019A-8B44-2557C*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAiEngineEventTrendsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAiEngineEventTrendsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAiEngineEventTrendsResponseBody) GetAggregateInterval() *string {
	return s.AggregateInterval
}

func (s *DescribeAiEngineEventTrendsResponseBody) GetPoints() []*DescribeAiEngineEventTrendsResponseBodyPoints {
	return s.Points
}

func (s *DescribeAiEngineEventTrendsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAiEngineEventTrendsResponseBody) SetAggregateInterval(v string) *DescribeAiEngineEventTrendsResponseBody {
	s.AggregateInterval = &v
	return s
}

func (s *DescribeAiEngineEventTrendsResponseBody) SetPoints(v []*DescribeAiEngineEventTrendsResponseBodyPoints) *DescribeAiEngineEventTrendsResponseBody {
	s.Points = v
	return s
}

func (s *DescribeAiEngineEventTrendsResponseBody) SetRequestId(v string) *DescribeAiEngineEventTrendsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAiEngineEventTrendsResponseBody) Validate() error {
	if s.Points != nil {
		for _, item := range s.Points {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAiEngineEventTrendsResponseBodyPoints struct {
	// The number of events first discovered within the time period.
	//
	// example:
	//
	// 1
	EventCount *int64 `json:"EventCount,omitempty" xml:"EventCount,omitempty"`
	// The number of events within the time period, aggregated by risk level.
	RiskCounts *DescribeAiEngineEventTrendsResponseBodyPointsRiskCounts `json:"RiskCounts,omitempty" xml:"RiskCounts,omitempty" type:"Struct"`
	// The start time of the statistical period in the UNIX timestamp format. Unit: milliseconds.
	//
	// example:
	//
	// 1792973400000
	Time *int64 `json:"Time,omitempty" xml:"Time,omitempty"`
}

func (s DescribeAiEngineEventTrendsResponseBodyPoints) String() string {
	return dara.Prettify(s)
}

func (s DescribeAiEngineEventTrendsResponseBodyPoints) GoString() string {
	return s.String()
}

func (s *DescribeAiEngineEventTrendsResponseBodyPoints) GetEventCount() *int64 {
	return s.EventCount
}

func (s *DescribeAiEngineEventTrendsResponseBodyPoints) GetRiskCounts() *DescribeAiEngineEventTrendsResponseBodyPointsRiskCounts {
	return s.RiskCounts
}

func (s *DescribeAiEngineEventTrendsResponseBodyPoints) GetTime() *int64 {
	return s.Time
}

func (s *DescribeAiEngineEventTrendsResponseBodyPoints) SetEventCount(v int64) *DescribeAiEngineEventTrendsResponseBodyPoints {
	s.EventCount = &v
	return s
}

func (s *DescribeAiEngineEventTrendsResponseBodyPoints) SetRiskCounts(v *DescribeAiEngineEventTrendsResponseBodyPointsRiskCounts) *DescribeAiEngineEventTrendsResponseBodyPoints {
	s.RiskCounts = v
	return s
}

func (s *DescribeAiEngineEventTrendsResponseBodyPoints) SetTime(v int64) *DescribeAiEngineEventTrendsResponseBodyPoints {
	s.Time = &v
	return s
}

func (s *DescribeAiEngineEventTrendsResponseBodyPoints) Validate() error {
	if s.RiskCounts != nil {
		if err := s.RiskCounts.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAiEngineEventTrendsResponseBodyPointsRiskCounts struct {
	// The number of critical events.
	//
	// example:
	//
	// 0
	Critical *int64 `json:"Critical,omitempty" xml:"Critical,omitempty"`
	// The number of high-risk events.
	//
	// example:
	//
	// 1
	High *int64 `json:"High,omitempty" xml:"High,omitempty"`
	// The number of low-risk events.
	//
	// example:
	//
	// 0
	Low *int64 `json:"Low,omitempty" xml:"Low,omitempty"`
	// The number of medium-risk events.
	//
	// example:
	//
	// 0
	Medium *int64 `json:"Medium,omitempty" xml:"Medium,omitempty"`
}

func (s DescribeAiEngineEventTrendsResponseBodyPointsRiskCounts) String() string {
	return dara.Prettify(s)
}

func (s DescribeAiEngineEventTrendsResponseBodyPointsRiskCounts) GoString() string {
	return s.String()
}

func (s *DescribeAiEngineEventTrendsResponseBodyPointsRiskCounts) GetCritical() *int64 {
	return s.Critical
}

func (s *DescribeAiEngineEventTrendsResponseBodyPointsRiskCounts) GetHigh() *int64 {
	return s.High
}

func (s *DescribeAiEngineEventTrendsResponseBodyPointsRiskCounts) GetLow() *int64 {
	return s.Low
}

func (s *DescribeAiEngineEventTrendsResponseBodyPointsRiskCounts) GetMedium() *int64 {
	return s.Medium
}

func (s *DescribeAiEngineEventTrendsResponseBodyPointsRiskCounts) SetCritical(v int64) *DescribeAiEngineEventTrendsResponseBodyPointsRiskCounts {
	s.Critical = &v
	return s
}

func (s *DescribeAiEngineEventTrendsResponseBodyPointsRiskCounts) SetHigh(v int64) *DescribeAiEngineEventTrendsResponseBodyPointsRiskCounts {
	s.High = &v
	return s
}

func (s *DescribeAiEngineEventTrendsResponseBodyPointsRiskCounts) SetLow(v int64) *DescribeAiEngineEventTrendsResponseBodyPointsRiskCounts {
	s.Low = &v
	return s
}

func (s *DescribeAiEngineEventTrendsResponseBodyPointsRiskCounts) SetMedium(v int64) *DescribeAiEngineEventTrendsResponseBodyPointsRiskCounts {
	s.Medium = &v
	return s
}

func (s *DescribeAiEngineEventTrendsResponseBodyPointsRiskCounts) Validate() error {
	return dara.Validate(s)
}
