// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainRealTimeReqHitRateDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeDcdnDomainRealTimeReqHitRateDataRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeDcdnDomainRealTimeReqHitRateDataRequest
	GetEndTime() *string
	SetStartTime(v string) *DescribeDcdnDomainRealTimeReqHitRateDataRequest
	GetStartTime() *string
}

type DescribeDcdnDomainRealTimeReqHitRateDataRequest struct {
	// The accelerated domain name. Separate multiple accelerated domain names with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// The end time needs to be later than the start time.
	//
	// example:
	//
	// 2019-12-10T20:03:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2019-12-10T20:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDcdnDomainRealTimeReqHitRateDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainRealTimeReqHitRateDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainRealTimeReqHitRateDataRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDcdnDomainRealTimeReqHitRateDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDcdnDomainRealTimeReqHitRateDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDcdnDomainRealTimeReqHitRateDataRequest) SetDomainName(v string) *DescribeDcdnDomainRealTimeReqHitRateDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeReqHitRateDataRequest) SetEndTime(v string) *DescribeDcdnDomainRealTimeReqHitRateDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeReqHitRateDataRequest) SetStartTime(v string) *DescribeDcdnDomainRealTimeReqHitRateDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeReqHitRateDataRequest) Validate() error {
	return dara.Validate(s)
}
