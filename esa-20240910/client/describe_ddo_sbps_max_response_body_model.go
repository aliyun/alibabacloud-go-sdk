// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDDoSBpsMaxResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeDDoSBpsMaxResponseBody
	GetEndTime() *string
	SetMaxAtkBps(v int64) *DescribeDDoSBpsMaxResponseBody
	GetMaxAtkBps() *int64
	SetMaxAtkPps(v int64) *DescribeDDoSBpsMaxResponseBody
	GetMaxAtkPps() *int64
	SetRequestId(v string) *DescribeDDoSBpsMaxResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeDDoSBpsMaxResponseBody
	GetStartTime() *string
}

type DescribeDDoSBpsMaxResponseBody struct {
	// The end of the time range that was queried. The time is in ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is in UTC+0.
	//
	// example:
	//
	// 2023-04-07T02:34:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The peak attack bandwidth. Unit: bit/s.
	//
	// example:
	//
	// 10000000000
	MaxAtkBps *int64 `json:"MaxAtkBps,omitempty" xml:"MaxAtkBps,omitempty"`
	// The peak attack PPS. Unit: pps.
	//
	// example:
	//
	// 100000000
	MaxAtkPps *int64 `json:"MaxAtkPps,omitempty" xml:"MaxAtkPps,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A3790430-3A06-535F-A424-0998BD9A6C9F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The beginning of the time range to query. Specify the time in ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC+0.
	//
	// example:
	//
	// 2023-02-14T17:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDDoSBpsMaxResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDDoSBpsMaxResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDDoSBpsMaxResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDDoSBpsMaxResponseBody) GetMaxAtkBps() *int64 {
	return s.MaxAtkBps
}

func (s *DescribeDDoSBpsMaxResponseBody) GetMaxAtkPps() *int64 {
	return s.MaxAtkPps
}

func (s *DescribeDDoSBpsMaxResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDDoSBpsMaxResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDDoSBpsMaxResponseBody) SetEndTime(v string) *DescribeDDoSBpsMaxResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDDoSBpsMaxResponseBody) SetMaxAtkBps(v int64) *DescribeDDoSBpsMaxResponseBody {
	s.MaxAtkBps = &v
	return s
}

func (s *DescribeDDoSBpsMaxResponseBody) SetMaxAtkPps(v int64) *DescribeDDoSBpsMaxResponseBody {
	s.MaxAtkPps = &v
	return s
}

func (s *DescribeDDoSBpsMaxResponseBody) SetRequestId(v string) *DescribeDDoSBpsMaxResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDDoSBpsMaxResponseBody) SetStartTime(v string) *DescribeDDoSBpsMaxResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDDoSBpsMaxResponseBody) Validate() error {
	return dara.Validate(s)
}
