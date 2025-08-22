// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainRealTimeByteHitRateDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeDcdnDomainRealTimeByteHitRateDataRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeDcdnDomainRealTimeByteHitRateDataRequest
	GetEndTime() *string
	SetStartTime(v string) *DescribeDcdnDomainRealTimeByteHitRateDataRequest
	GetStartTime() *string
}

type DescribeDcdnDomainRealTimeByteHitRateDataRequest struct {
	// The domain name. Separate multiple domain names with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// The end time must be later than the start time.
	//
	// example:
	//
	// 2017-12-10T21:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2017-12-10T20:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDcdnDomainRealTimeByteHitRateDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainRealTimeByteHitRateDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainRealTimeByteHitRateDataRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDcdnDomainRealTimeByteHitRateDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDcdnDomainRealTimeByteHitRateDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDcdnDomainRealTimeByteHitRateDataRequest) SetDomainName(v string) *DescribeDcdnDomainRealTimeByteHitRateDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeByteHitRateDataRequest) SetEndTime(v string) *DescribeDcdnDomainRealTimeByteHitRateDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeByteHitRateDataRequest) SetStartTime(v string) *DescribeDcdnDomainRealTimeByteHitRateDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeByteHitRateDataRequest) Validate() error {
	return dara.Validate(s)
}
