// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserAbnormalTrendResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeUserAbnormalTrendResponseBody
	GetRequestId() *string
	SetTrend(v []*DescribeUserAbnormalTrendResponseBodyTrend) *DescribeUserAbnormalTrendResponseBody
	GetTrend() []*DescribeUserAbnormalTrendResponseBodyTrend
}

type DescribeUserAbnormalTrendResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// D7861F61-5B61-46CE-A47C-6B19160D5EB0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The trends of risks.
	Trend []*DescribeUserAbnormalTrendResponseBodyTrend `json:"Trend,omitempty" xml:"Trend,omitempty" type:"Repeated"`
}

func (s DescribeUserAbnormalTrendResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserAbnormalTrendResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserAbnormalTrendResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeUserAbnormalTrendResponseBody) GetTrend() []*DescribeUserAbnormalTrendResponseBodyTrend {
	return s.Trend
}

func (s *DescribeUserAbnormalTrendResponseBody) SetRequestId(v string) *DescribeUserAbnormalTrendResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUserAbnormalTrendResponseBody) SetTrend(v []*DescribeUserAbnormalTrendResponseBodyTrend) *DescribeUserAbnormalTrendResponseBody {
	s.Trend = v
	return s
}

func (s *DescribeUserAbnormalTrendResponseBody) Validate() error {
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

type DescribeUserAbnormalTrendResponseBodyTrend struct {
	// The number of high risks.
	//
	// example:
	//
	// 12
	AbnormalHigh *int64 `json:"AbnormalHigh,omitempty" xml:"AbnormalHigh,omitempty"`
	// The number of low risks.
	//
	// example:
	//
	// 23
	AbnormalLow *int64 `json:"AbnormalLow,omitempty" xml:"AbnormalLow,omitempty"`
	// The number of medium risks.
	//
	// example:
	//
	// 14
	AbnormalMedium *int64 `json:"AbnormalMedium,omitempty" xml:"AbnormalMedium,omitempty"`
	// Deprecated
	//
	// The time at which the API was called. The value is a UNIX timestamp displayed in UTC. Unit: seconds.
	//
	// 	Notice: The parameter has been deprecated, it is recommended to use the Timestamp parameter.
	//
	// example:
	//
	// 1722268800
	TimeStamp *int64 `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	// The time at which the API was called. The value is a UNIX timestamp displayed in UTC. Unit: seconds.
	//
	// example:
	//
	// 1722268800
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s DescribeUserAbnormalTrendResponseBodyTrend) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserAbnormalTrendResponseBodyTrend) GoString() string {
	return s.String()
}

func (s *DescribeUserAbnormalTrendResponseBodyTrend) GetAbnormalHigh() *int64 {
	return s.AbnormalHigh
}

func (s *DescribeUserAbnormalTrendResponseBodyTrend) GetAbnormalLow() *int64 {
	return s.AbnormalLow
}

func (s *DescribeUserAbnormalTrendResponseBodyTrend) GetAbnormalMedium() *int64 {
	return s.AbnormalMedium
}

func (s *DescribeUserAbnormalTrendResponseBodyTrend) GetTimeStamp() *int64 {
	return s.TimeStamp
}

func (s *DescribeUserAbnormalTrendResponseBodyTrend) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *DescribeUserAbnormalTrendResponseBodyTrend) SetAbnormalHigh(v int64) *DescribeUserAbnormalTrendResponseBodyTrend {
	s.AbnormalHigh = &v
	return s
}

func (s *DescribeUserAbnormalTrendResponseBodyTrend) SetAbnormalLow(v int64) *DescribeUserAbnormalTrendResponseBodyTrend {
	s.AbnormalLow = &v
	return s
}

func (s *DescribeUserAbnormalTrendResponseBodyTrend) SetAbnormalMedium(v int64) *DescribeUserAbnormalTrendResponseBodyTrend {
	s.AbnormalMedium = &v
	return s
}

func (s *DescribeUserAbnormalTrendResponseBodyTrend) SetTimeStamp(v int64) *DescribeUserAbnormalTrendResponseBodyTrend {
	s.TimeStamp = &v
	return s
}

func (s *DescribeUserAbnormalTrendResponseBodyTrend) SetTimestamp(v int64) *DescribeUserAbnormalTrendResponseBodyTrend {
	s.Timestamp = &v
	return s
}

func (s *DescribeUserAbnormalTrendResponseBodyTrend) Validate() error {
	return dara.Validate(s)
}
