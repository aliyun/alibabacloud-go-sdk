// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnBgpTrafficDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeDcdnBgpTrafficDataRequest
	GetEndTime() *string
	SetInterval(v string) *DescribeDcdnBgpTrafficDataRequest
	GetInterval() *string
	SetIsp(v string) *DescribeDcdnBgpTrafficDataRequest
	GetIsp() *string
	SetStartTime(v string) *DescribeDcdnBgpTrafficDataRequest
	GetStartTime() *string
}

type DescribeDcdnBgpTrafficDataRequest struct {
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2018-11-30T00:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The data collection interval. Unit: seconds. Valid values: 300 and 3600. Default value: 300. The default value of 300 seconds is equal to 5 minutes. The value of this parameter varies based on the time range from the specified start time to the specified end time.
	//
	// example:
	//
	// 300
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The ISP. Separate multiple ISPs with commas (,). If you specify multiple ISPs, the data for the ISPs is aggregated. If you do not specify this parameter, the operation returns the data for all the ISPs.
	//
	// Valid values:
	//
	// 	- cu: China Unicom
	//
	// 	- cmi: China Mobile
	//
	// 	- ct: China Telecom
	//
	// example:
	//
	// cu
	Isp *string `json:"Isp,omitempty" xml:"Isp,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// The minimum data collection interval is an hour.
	//
	// If you do not set this parameter, data collected in the last 24 hours is queried.
	//
	// example:
	//
	// 2018-11-29T00:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDcdnBgpTrafficDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnBgpTrafficDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnBgpTrafficDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDcdnBgpTrafficDataRequest) GetInterval() *string {
	return s.Interval
}

func (s *DescribeDcdnBgpTrafficDataRequest) GetIsp() *string {
	return s.Isp
}

func (s *DescribeDcdnBgpTrafficDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDcdnBgpTrafficDataRequest) SetEndTime(v string) *DescribeDcdnBgpTrafficDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnBgpTrafficDataRequest) SetInterval(v string) *DescribeDcdnBgpTrafficDataRequest {
	s.Interval = &v
	return s
}

func (s *DescribeDcdnBgpTrafficDataRequest) SetIsp(v string) *DescribeDcdnBgpTrafficDataRequest {
	s.Isp = &v
	return s
}

func (s *DescribeDcdnBgpTrafficDataRequest) SetStartTime(v string) *DescribeDcdnBgpTrafficDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDcdnBgpTrafficDataRequest) Validate() error {
	return dara.Validate(s)
}
