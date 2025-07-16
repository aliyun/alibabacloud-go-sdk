// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainPvDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeDomainPvDataRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeDomainPvDataRequest
	GetEndTime() *string
	SetStartTime(v string) *DescribeDomainPvDataRequest
	GetStartTime() *string
}

type DescribeDomainPvDataRequest struct {
	// The accelerated domain name. You can specify only one domain name.
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
	// 2015-11-29T00:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2015-11-28T00:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDomainPvDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainPvDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainPvDataRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDomainPvDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDomainPvDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDomainPvDataRequest) SetDomainName(v string) *DescribeDomainPvDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainPvDataRequest) SetEndTime(v string) *DescribeDomainPvDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainPvDataRequest) SetStartTime(v string) *DescribeDomainPvDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainPvDataRequest) Validate() error {
	return dara.Validate(s)
}
