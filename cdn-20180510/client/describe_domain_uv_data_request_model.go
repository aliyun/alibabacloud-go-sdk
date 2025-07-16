// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainUvDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeDomainUvDataRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeDomainUvDataRequest
	GetEndTime() *string
	SetStartTime(v string) *DescribeDomainUvDataRequest
	GetStartTime() *string
}

type DescribeDomainUvDataRequest struct {
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
	// 2019-11-29T04:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2019-11-29T00:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDomainUvDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainUvDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainUvDataRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDomainUvDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDomainUvDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDomainUvDataRequest) SetDomainName(v string) *DescribeDomainUvDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainUvDataRequest) SetEndTime(v string) *DescribeDomainUvDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainUvDataRequest) SetStartTime(v string) *DescribeDomainUvDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainUvDataRequest) Validate() error {
	return dara.Validate(s)
}
