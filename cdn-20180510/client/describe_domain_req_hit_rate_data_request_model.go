// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainReqHitRateDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeDomainReqHitRateDataRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeDomainReqHitRateDataRequest
	GetEndTime() *string
	SetInterval(v string) *DescribeDomainReqHitRateDataRequest
	GetInterval() *string
	SetStartTime(v string) *DescribeDomainReqHitRateDataRequest
	GetStartTime() *string
}

type DescribeDomainReqHitRateDataRequest struct {
	// The accelerated domain name. You can specify up to 500 domain names in each request. Separate multiple domain names with commas (,).
	//
	// By default, this operation queries the request hit ratio for all accelerated domain names that belong to your Alibaba Cloud account.
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
	// 2017-12-22T08:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The time granularity of the data entries. Unit: seconds.
	//
	// The time granularity varies with the maximum time range per query. Valid values: 300 (5 minutes), 3600 (1 hour), and 86400 (1 day). For more information, see **Usage notes**.
	//
	// example:
	//
	// 300
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The beginning of the time range to query.
	//
	// Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2017-12-21T08:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDomainReqHitRateDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainReqHitRateDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainReqHitRateDataRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDomainReqHitRateDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDomainReqHitRateDataRequest) GetInterval() *string {
	return s.Interval
}

func (s *DescribeDomainReqHitRateDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDomainReqHitRateDataRequest) SetDomainName(v string) *DescribeDomainReqHitRateDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainReqHitRateDataRequest) SetEndTime(v string) *DescribeDomainReqHitRateDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainReqHitRateDataRequest) SetInterval(v string) *DescribeDomainReqHitRateDataRequest {
	s.Interval = &v
	return s
}

func (s *DescribeDomainReqHitRateDataRequest) SetStartTime(v string) *DescribeDomainReqHitRateDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainReqHitRateDataRequest) Validate() error {
	return dara.Validate(s)
}
