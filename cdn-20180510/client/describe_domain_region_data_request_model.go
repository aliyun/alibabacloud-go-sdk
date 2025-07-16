// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainRegionDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeDomainRegionDataRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeDomainRegionDataRequest
	GetEndTime() *string
	SetStartTime(v string) *DescribeDomainRegionDataRequest
	GetStartTime() *string
}

type DescribeDomainRegionDataRequest struct {
	// The accelerated domain name. Separate multiple accelerated domain names with commas (,).
	//
	// By default, this operation queries the geographic distribution of users for all accelerated domain names.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end of the time range to query.
	//
	// Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// The end time must be later than the start time.
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

func (s DescribeDomainRegionDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainRegionDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainRegionDataRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDomainRegionDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDomainRegionDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDomainRegionDataRequest) SetDomainName(v string) *DescribeDomainRegionDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainRegionDataRequest) SetEndTime(v string) *DescribeDomainRegionDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainRegionDataRequest) SetStartTime(v string) *DescribeDomainRegionDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainRegionDataRequest) Validate() error {
	return dara.Validate(s)
}
