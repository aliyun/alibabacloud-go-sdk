// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserEventTrendResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeUserEventTrendResponseBody
	GetRequestId() *string
	SetTrend(v []*DescribeUserEventTrendResponseBodyTrend) *DescribeUserEventTrendResponseBody
	GetTrend() []*DescribeUserEventTrendResponseBodyTrend
}

type DescribeUserEventTrendResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// F51E6DD6-B2D2-57C9-90F1-FAFD0A19DE00
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The trends of attacks.
	Trend []*DescribeUserEventTrendResponseBodyTrend `json:"Trend,omitempty" xml:"Trend,omitempty" type:"Repeated"`
}

func (s DescribeUserEventTrendResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserEventTrendResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserEventTrendResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeUserEventTrendResponseBody) GetTrend() []*DescribeUserEventTrendResponseBodyTrend {
	return s.Trend
}

func (s *DescribeUserEventTrendResponseBody) SetRequestId(v string) *DescribeUserEventTrendResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUserEventTrendResponseBody) SetTrend(v []*DescribeUserEventTrendResponseBodyTrend) *DescribeUserEventTrendResponseBody {
	s.Trend = v
	return s
}

func (s *DescribeUserEventTrendResponseBody) Validate() error {
	if s.Trend != nil {
		for _, item := range s.Trend {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeUserEventTrendResponseBodyTrend struct {
	// The number of high-risk events.
	//
	// example:
	//
	// 9
	EventHigh *int64 `json:"EventHigh,omitempty" xml:"EventHigh,omitempty"`
	// The number of low-risk events.
	//
	// example:
	//
	// 23
	EventLow *int64 `json:"EventLow,omitempty" xml:"EventLow,omitempty"`
	// The number of medium-risk events.
	//
	// example:
	//
	// 17
	EventMedium *int64 `json:"EventMedium,omitempty" xml:"EventMedium,omitempty"`
	// The time at which the API was called. The value is a UNIX timestamp displayed in UTC. Unit: seconds.
	//
	// 	Notice: The parameter has been deprecated, it is recommended to use the Timestamp parameter.
	//
	// example:
	//
	// 1723435200
	TimeStamp *int64 `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	// The time at which the API was called. The value is a UNIX timestamp displayed in UTC. Unit: seconds.
	//
	// example:
	//
	// 1723435200
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s DescribeUserEventTrendResponseBodyTrend) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserEventTrendResponseBodyTrend) GoString() string {
	return s.String()
}

func (s *DescribeUserEventTrendResponseBodyTrend) GetEventHigh() *int64 {
	return s.EventHigh
}

func (s *DescribeUserEventTrendResponseBodyTrend) GetEventLow() *int64 {
	return s.EventLow
}

func (s *DescribeUserEventTrendResponseBodyTrend) GetEventMedium() *int64 {
	return s.EventMedium
}

func (s *DescribeUserEventTrendResponseBodyTrend) GetTimeStamp() *int64 {
	return s.TimeStamp
}

func (s *DescribeUserEventTrendResponseBodyTrend) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *DescribeUserEventTrendResponseBodyTrend) SetEventHigh(v int64) *DescribeUserEventTrendResponseBodyTrend {
	s.EventHigh = &v
	return s
}

func (s *DescribeUserEventTrendResponseBodyTrend) SetEventLow(v int64) *DescribeUserEventTrendResponseBodyTrend {
	s.EventLow = &v
	return s
}

func (s *DescribeUserEventTrendResponseBodyTrend) SetEventMedium(v int64) *DescribeUserEventTrendResponseBodyTrend {
	s.EventMedium = &v
	return s
}

func (s *DescribeUserEventTrendResponseBodyTrend) SetTimeStamp(v int64) *DescribeUserEventTrendResponseBodyTrend {
	s.TimeStamp = &v
	return s
}

func (s *DescribeUserEventTrendResponseBodyTrend) SetTimestamp(v int64) *DescribeUserEventTrendResponseBodyTrend {
	s.Timestamp = &v
	return s
}

func (s *DescribeUserEventTrendResponseBodyTrend) Validate() error {
	return dara.Validate(s)
}
