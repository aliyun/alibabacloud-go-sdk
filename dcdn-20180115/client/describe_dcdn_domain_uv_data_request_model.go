// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainUvDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeDcdnDomainUvDataRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeDcdnDomainUvDataRequest
	GetEndTime() *string
	SetStartTime(v string) *DescribeDcdnDomainUvDataRequest
	GetStartTime() *string
}

type DescribeDcdnDomainUvDataRequest struct {
	// The accelerated domain name. You can specify only one domain name in each request.
	//
	// If you do not specify a domain name, this operation queries UV data of all accelerated domain names in your account.
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
	// 2015-11-30T00:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The start of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2015-11-29T00:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDcdnDomainUvDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainUvDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainUvDataRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDcdnDomainUvDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDcdnDomainUvDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDcdnDomainUvDataRequest) SetDomainName(v string) *DescribeDcdnDomainUvDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainUvDataRequest) SetEndTime(v string) *DescribeDcdnDomainUvDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnDomainUvDataRequest) SetStartTime(v string) *DescribeDcdnDomainUvDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDcdnDomainUvDataRequest) Validate() error {
	return dara.Validate(s)
}
