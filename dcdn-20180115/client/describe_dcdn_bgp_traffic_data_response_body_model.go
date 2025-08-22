// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnBgpTrafficDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBgpDataInterval(v []*DescribeDcdnBgpTrafficDataResponseBodyBgpDataInterval) *DescribeDcdnBgpTrafficDataResponseBody
	GetBgpDataInterval() []*DescribeDcdnBgpTrafficDataResponseBodyBgpDataInterval
	SetEndTime(v string) *DescribeDcdnBgpTrafficDataResponseBody
	GetEndTime() *string
	SetRequestId(v string) *DescribeDcdnBgpTrafficDataResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeDcdnBgpTrafficDataResponseBody
	GetStartTime() *string
}

type DescribeDcdnBgpTrafficDataResponseBody struct {
	// The BGP traffic at each time interval.
	BgpDataInterval []*DescribeDcdnBgpTrafficDataResponseBodyBgpDataInterval `json:"BgpDataInterval,omitempty" xml:"BgpDataInterval,omitempty" type:"Repeated"`
	// The end of the time range during which data was queried.
	//
	// example:
	//
	// 2018-11-30T00:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// E9D3257A-1B7C-414C-90C1-8D07AC47BCAC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The start of the time range during which data was queried.
	//
	// example:
	//
	// 2018-11-29T00:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDcdnBgpTrafficDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnBgpTrafficDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnBgpTrafficDataResponseBody) GetBgpDataInterval() []*DescribeDcdnBgpTrafficDataResponseBodyBgpDataInterval {
	return s.BgpDataInterval
}

func (s *DescribeDcdnBgpTrafficDataResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDcdnBgpTrafficDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDcdnBgpTrafficDataResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDcdnBgpTrafficDataResponseBody) SetBgpDataInterval(v []*DescribeDcdnBgpTrafficDataResponseBodyBgpDataInterval) *DescribeDcdnBgpTrafficDataResponseBody {
	s.BgpDataInterval = v
	return s
}

func (s *DescribeDcdnBgpTrafficDataResponseBody) SetEndTime(v string) *DescribeDcdnBgpTrafficDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnBgpTrafficDataResponseBody) SetRequestId(v string) *DescribeDcdnBgpTrafficDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnBgpTrafficDataResponseBody) SetStartTime(v string) *DescribeDcdnBgpTrafficDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDcdnBgpTrafficDataResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnBgpTrafficDataResponseBodyBgpDataInterval struct {
	// The inbound traffic. Unit: bytes.
	//
	// example:
	//
	// 318
	In *int64 `json:"In,omitempty" xml:"In,omitempty"`
	// The outbound traffic. Unit: bytes.
	//
	// example:
	//
	// 183
	Out *int64 `json:"Out,omitempty" xml:"Out,omitempty"`
	// The timestamp of the data returned.
	//
	// example:
	//
	// 2018-11-29T20:00:00Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeDcdnBgpTrafficDataResponseBodyBgpDataInterval) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnBgpTrafficDataResponseBodyBgpDataInterval) GoString() string {
	return s.String()
}

func (s *DescribeDcdnBgpTrafficDataResponseBodyBgpDataInterval) GetIn() *int64 {
	return s.In
}

func (s *DescribeDcdnBgpTrafficDataResponseBodyBgpDataInterval) GetOut() *int64 {
	return s.Out
}

func (s *DescribeDcdnBgpTrafficDataResponseBodyBgpDataInterval) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeDcdnBgpTrafficDataResponseBodyBgpDataInterval) SetIn(v int64) *DescribeDcdnBgpTrafficDataResponseBodyBgpDataInterval {
	s.In = &v
	return s
}

func (s *DescribeDcdnBgpTrafficDataResponseBodyBgpDataInterval) SetOut(v int64) *DescribeDcdnBgpTrafficDataResponseBodyBgpDataInterval {
	s.Out = &v
	return s
}

func (s *DescribeDcdnBgpTrafficDataResponseBodyBgpDataInterval) SetTimeStamp(v string) *DescribeDcdnBgpTrafficDataResponseBodyBgpDataInterval {
	s.TimeStamp = &v
	return s
}

func (s *DescribeDcdnBgpTrafficDataResponseBodyBgpDataInterval) Validate() error {
	return dara.Validate(s)
}
