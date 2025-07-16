// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainRealTimeReqHitRateDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeDomainRealTimeReqHitRateDataRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeDomainRealTimeReqHitRateDataRequest
	GetEndTime() *string
	SetStartTime(v string) *DescribeDomainRealTimeReqHitRateDataRequest
	GetStartTime() *string
}

type DescribeDomainRealTimeReqHitRateDataRequest struct {
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
	// 2018-01-02T11:26:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The beginning of the time range to query.
	//
	// Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2018-01-02T11:23:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDomainRealTimeReqHitRateDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainRealTimeReqHitRateDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainRealTimeReqHitRateDataRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDomainRealTimeReqHitRateDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDomainRealTimeReqHitRateDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDomainRealTimeReqHitRateDataRequest) SetDomainName(v string) *DescribeDomainRealTimeReqHitRateDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainRealTimeReqHitRateDataRequest) SetEndTime(v string) *DescribeDomainRealTimeReqHitRateDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainRealTimeReqHitRateDataRequest) SetStartTime(v string) *DescribeDomainRealTimeReqHitRateDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainRealTimeReqHitRateDataRequest) Validate() error {
	return dara.Validate(s)
}
