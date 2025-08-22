// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainIspDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeDcdnDomainIspDataRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeDcdnDomainIspDataRequest
	GetEndTime() *string
	SetStartTime(v string) *DescribeDcdnDomainIspDataRequest
	GetStartTime() *string
}

type DescribeDcdnDomainIspDataRequest struct {
	// The accelerated domain name. You can specify only one domain name in each request.
	//
	// If you do not specify an accelerated domain name, the data of all accelerated domain names that belong to your account is queried.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end of the time range to query. Specify the time in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// > The end time must be later than the start time.
	//
	// example:
	//
	// 2019-12-06T12:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The beginning of the time range to query. Specify the time in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2019-12-05T12:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDcdnDomainIspDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainIspDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainIspDataRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDcdnDomainIspDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDcdnDomainIspDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDcdnDomainIspDataRequest) SetDomainName(v string) *DescribeDcdnDomainIspDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainIspDataRequest) SetEndTime(v string) *DescribeDcdnDomainIspDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnDomainIspDataRequest) SetStartTime(v string) *DescribeDcdnDomainIspDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDcdnDomainIspDataRequest) Validate() error {
	return dara.Validate(s)
}
