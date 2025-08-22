// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnBgpBpsDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBgpDataInterval(v []*DescribeDcdnBgpBpsDataResponseBodyBgpDataInterval) *DescribeDcdnBgpBpsDataResponseBody
	GetBgpDataInterval() []*DescribeDcdnBgpBpsDataResponseBodyBgpDataInterval
	SetEndTime(v string) *DescribeDcdnBgpBpsDataResponseBody
	GetEndTime() *string
	SetRequestId(v string) *DescribeDcdnBgpBpsDataResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeDcdnBgpBpsDataResponseBody
	GetStartTime() *string
}

type DescribeDcdnBgpBpsDataResponseBody struct {
	// The BGP bandwidth data that is collected for each interval.
	BgpDataInterval []*DescribeDcdnBgpBpsDataResponseBodyBgpDataInterval `json:"BgpDataInterval,omitempty" xml:"BgpDataInterval,omitempty" type:"Repeated"`
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

func (s DescribeDcdnBgpBpsDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnBgpBpsDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnBgpBpsDataResponseBody) GetBgpDataInterval() []*DescribeDcdnBgpBpsDataResponseBodyBgpDataInterval {
	return s.BgpDataInterval
}

func (s *DescribeDcdnBgpBpsDataResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDcdnBgpBpsDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDcdnBgpBpsDataResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDcdnBgpBpsDataResponseBody) SetBgpDataInterval(v []*DescribeDcdnBgpBpsDataResponseBodyBgpDataInterval) *DescribeDcdnBgpBpsDataResponseBody {
	s.BgpDataInterval = v
	return s
}

func (s *DescribeDcdnBgpBpsDataResponseBody) SetEndTime(v string) *DescribeDcdnBgpBpsDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnBgpBpsDataResponseBody) SetRequestId(v string) *DescribeDcdnBgpBpsDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnBgpBpsDataResponseBody) SetStartTime(v string) *DescribeDcdnBgpBpsDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDcdnBgpBpsDataResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnBgpBpsDataResponseBodyBgpDataInterval struct {
	// The inbound bandwidth. Unit: bit/s.
	//
	// example:
	//
	// 318
	In *float32 `json:"In,omitempty" xml:"In,omitempty"`
	// The outbound bandwidth. Unit: bit/s.
	//
	// example:
	//
	// 183
	Out *float32 `json:"Out,omitempty" xml:"Out,omitempty"`
	// The timestamp of the returned data.
	//
	// example:
	//
	// 2018-11-29T20:00:00Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeDcdnBgpBpsDataResponseBodyBgpDataInterval) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnBgpBpsDataResponseBodyBgpDataInterval) GoString() string {
	return s.String()
}

func (s *DescribeDcdnBgpBpsDataResponseBodyBgpDataInterval) GetIn() *float32 {
	return s.In
}

func (s *DescribeDcdnBgpBpsDataResponseBodyBgpDataInterval) GetOut() *float32 {
	return s.Out
}

func (s *DescribeDcdnBgpBpsDataResponseBodyBgpDataInterval) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeDcdnBgpBpsDataResponseBodyBgpDataInterval) SetIn(v float32) *DescribeDcdnBgpBpsDataResponseBodyBgpDataInterval {
	s.In = &v
	return s
}

func (s *DescribeDcdnBgpBpsDataResponseBodyBgpDataInterval) SetOut(v float32) *DescribeDcdnBgpBpsDataResponseBodyBgpDataInterval {
	s.Out = &v
	return s
}

func (s *DescribeDcdnBgpBpsDataResponseBodyBgpDataInterval) SetTimeStamp(v string) *DescribeDcdnBgpBpsDataResponseBodyBgpDataInterval {
	s.TimeStamp = &v
	return s
}

func (s *DescribeDcdnBgpBpsDataResponseBodyBgpDataInterval) Validate() error {
	return dara.Validate(s)
}
