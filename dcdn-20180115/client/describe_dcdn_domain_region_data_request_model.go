// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainRegionDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeDcdnDomainRegionDataRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeDcdnDomainRegionDataRequest
	GetEndTime() *string
	SetStartTime(v string) *DescribeDcdnDomainRegionDataRequest
	GetStartTime() *string
}

type DescribeDcdnDomainRegionDataRequest struct {
	// The accelerated domain name. You can specify only one domain name.
	//
	// If you do not specify an accelerated domain name, the data of all accelerated domain names that belong to your account is queried.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end of the time range to query. The end time must be later than the start time.
	//
	// Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2015-12-07T12:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The beginning of the time range to query.
	//
	// Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2015-12-05T12:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDcdnDomainRegionDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainRegionDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainRegionDataRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDcdnDomainRegionDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDcdnDomainRegionDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDcdnDomainRegionDataRequest) SetDomainName(v string) *DescribeDcdnDomainRegionDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainRegionDataRequest) SetEndTime(v string) *DescribeDcdnDomainRegionDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnDomainRegionDataRequest) SetStartTime(v string) *DescribeDcdnDomainRegionDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDcdnDomainRegionDataRequest) Validate() error {
	return dara.Validate(s)
}
