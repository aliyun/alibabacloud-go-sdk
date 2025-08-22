// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainPvDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeDcdnDomainPvDataRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeDcdnDomainPvDataRequest
	GetEndTime() *string
	SetStartTime(v string) *DescribeDcdnDomainPvDataRequest
	GetStartTime() *string
}

type DescribeDcdnDomainPvDataRequest struct {
	// The accelerated domain name. You can specify only one domain name in each request.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2019-11-29T00:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The start of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2019-11-28T00:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDcdnDomainPvDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainPvDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainPvDataRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDcdnDomainPvDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDcdnDomainPvDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDcdnDomainPvDataRequest) SetDomainName(v string) *DescribeDcdnDomainPvDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainPvDataRequest) SetEndTime(v string) *DescribeDcdnDomainPvDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnDomainPvDataRequest) SetStartTime(v string) *DescribeDcdnDomainPvDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDcdnDomainPvDataRequest) Validate() error {
	return dara.Validate(s)
}
