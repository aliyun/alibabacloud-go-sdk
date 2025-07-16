// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainRealTimeByteHitRateDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeDomainRealTimeByteHitRateDataRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeDomainRealTimeByteHitRateDataRequest
	GetEndTime() *string
	SetStartTime(v string) *DescribeDomainRealTimeByteHitRateDataRequest
	GetStartTime() *string
}

type DescribeDomainRealTimeByteHitRateDataRequest struct {
	// The accelerated domain name. You can specify up to 100 domain names in each call. Separate multiple domain names with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end of the time range to query.
	//
	// Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// > The end time must be later than the start time.
	//
	// example:
	//
	// 2020-05-15T09:15:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The beginning of the time range to query.
	//
	// Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2020-05-15T09:13:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDomainRealTimeByteHitRateDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainRealTimeByteHitRateDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainRealTimeByteHitRateDataRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDomainRealTimeByteHitRateDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDomainRealTimeByteHitRateDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDomainRealTimeByteHitRateDataRequest) SetDomainName(v string) *DescribeDomainRealTimeByteHitRateDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainRealTimeByteHitRateDataRequest) SetEndTime(v string) *DescribeDomainRealTimeByteHitRateDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainRealTimeByteHitRateDataRequest) SetStartTime(v string) *DescribeDomainRealTimeByteHitRateDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainRealTimeByteHitRateDataRequest) Validate() error {
	return dara.Validate(s)
}
