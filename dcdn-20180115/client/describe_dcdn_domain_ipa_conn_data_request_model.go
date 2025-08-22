// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainIpaConnDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeDcdnDomainIpaConnDataRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeDcdnDomainIpaConnDataRequest
	GetEndTime() *string
	SetSplitBy(v string) *DescribeDcdnDomainIpaConnDataRequest
	GetSplitBy() *string
	SetStartTime(v string) *DescribeDcdnDomainIpaConnDataRequest
	GetStartTime() *string
}

type DescribeDcdnDomainIpaConnDataRequest struct {
	// The domain names accelerated by IPA. Separate multiple domain names with commas (,).
	//
	// > If you do not specify this parameter, the merged data of all accelerated domain names is returned.
	//
	// example:
	//
	// example1.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end of the time range to query.
	//
	// Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// > The end time must be later than the start time.
	//
	// example:
	//
	// 2015-02-22T7:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Specifies how query results are grouped. By default, this parameter is empty. Valid values:
	//
	// 	- domain: Query results are grouped by accelerated domain name.
	//
	// 	- An empty string: Query results are not grouped.
	//
	// example:
	//
	// domain
	SplitBy *string `json:"SplitBy,omitempty" xml:"SplitBy,omitempty"`
	// The start of the time range to query.
	//
	// Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2015-02-21T07:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDcdnDomainIpaConnDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainIpaConnDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainIpaConnDataRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDcdnDomainIpaConnDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDcdnDomainIpaConnDataRequest) GetSplitBy() *string {
	return s.SplitBy
}

func (s *DescribeDcdnDomainIpaConnDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDcdnDomainIpaConnDataRequest) SetDomainName(v string) *DescribeDcdnDomainIpaConnDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainIpaConnDataRequest) SetEndTime(v string) *DescribeDcdnDomainIpaConnDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnDomainIpaConnDataRequest) SetSplitBy(v string) *DescribeDcdnDomainIpaConnDataRequest {
	s.SplitBy = &v
	return s
}

func (s *DescribeDcdnDomainIpaConnDataRequest) SetStartTime(v string) *DescribeDcdnDomainIpaConnDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDcdnDomainIpaConnDataRequest) Validate() error {
	return dara.Validate(s)
}
