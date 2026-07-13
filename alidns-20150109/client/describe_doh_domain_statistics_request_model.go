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
	// The end date of the query. The format is YYYY-MM-DD.
	//
	// The default value is the current day.
	//
	// example:
	//
	// 2019-07-04
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// The language.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The start date of the query. The format is YYYY-MM-DD.
	//
	// You can query data from the last 90 days only. The value of `StartDate` must be greater than or equal to the current date minus 90 days.
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
