// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDohDomainStatisticsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeDohDomainStatisticsRequest
	GetDomainName() *string
	SetEndDate(v string) *DescribeDohDomainStatisticsRequest
	GetEndDate() *string
	SetLang(v string) *DescribeDohDomainStatisticsRequest
	GetLang() *string
	SetStartDate(v string) *DescribeDohDomainStatisticsRequest
	GetStartDate() *string
}

type DescribeDohDomainStatisticsRequest struct {
	// The domain name.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end of the time range to query. Specify the time in the YYYY-MM-DD format.
	//
	// The default value is the day when you perform the operation.
	//
	// example:
	//
	// 2019-07-04
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// The language type.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The beginning of the time range to query. Specify the time in the YYYY-MM-DD format.
	//
	// You can query only the DNS records of the latest 90 days. `The value of StartDate must be greater than or equal to the difference between the current date and 90`.
	//
	// example:
	//
	// 2019-07-04
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s DescribeDohDomainStatisticsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDohDomainStatisticsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDohDomainStatisticsRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDohDomainStatisticsRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *DescribeDohDomainStatisticsRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeDohDomainStatisticsRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *DescribeDohDomainStatisticsRequest) SetDomainName(v string) *DescribeDohDomainStatisticsRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDohDomainStatisticsRequest) SetEndDate(v string) *DescribeDohDomainStatisticsRequest {
	s.EndDate = &v
	return s
}

func (s *DescribeDohDomainStatisticsRequest) SetLang(v string) *DescribeDohDomainStatisticsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDohDomainStatisticsRequest) SetStartDate(v string) *DescribeDohDomainStatisticsRequest {
	s.StartDate = &v
	return s
}

func (s *DescribeDohDomainStatisticsRequest) Validate() error {
	return dara.Validate(s)
}
