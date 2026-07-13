// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDohSubDomainStatisticsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndDate(v string) *DescribeDohSubDomainStatisticsRequest
	GetEndDate() *string
	SetLang(v string) *DescribeDohSubDomainStatisticsRequest
	GetLang() *string
	SetStartDate(v string) *DescribeDohSubDomainStatisticsRequest
	GetStartDate() *string
	SetSubDomain(v string) *DescribeDohSubDomainStatisticsRequest
	GetSubDomain() *string
}

type DescribeDohSubDomainStatisticsRequest struct {
	// The end date of the query in YYYY-MM-DD format.
	//
	// The default value is the current date.
	//
	// example:
	//
	// 2019-07-04
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// The language of the request and response. Default value: **zh**. Valid values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The start date of the query in YYYY-MM-DD format.
	//
	// You can query data from the last 90 days. The `StartDate` must be greater than or equal to the date 90 days before the current date.
	//
	// example:
	//
	// 2019-07-04
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// The subdomain for which to query statistics.
	//
	// This parameter is required.
	//
	// example:
	//
	// www.example.com
	SubDomain *string `json:"SubDomain,omitempty" xml:"SubDomain,omitempty"`
}

func (s DescribeDohSubDomainStatisticsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDohSubDomainStatisticsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDohSubDomainStatisticsRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *DescribeDohSubDomainStatisticsRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeDohSubDomainStatisticsRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *DescribeDohSubDomainStatisticsRequest) GetSubDomain() *string {
	return s.SubDomain
}

func (s *DescribeDohSubDomainStatisticsRequest) SetEndDate(v string) *DescribeDohSubDomainStatisticsRequest {
	s.EndDate = &v
	return s
}

func (s *DescribeDohSubDomainStatisticsRequest) SetLang(v string) *DescribeDohSubDomainStatisticsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDohSubDomainStatisticsRequest) SetStartDate(v string) *DescribeDohSubDomainStatisticsRequest {
	s.StartDate = &v
	return s
}

func (s *DescribeDohSubDomainStatisticsRequest) SetSubDomain(v string) *DescribeDohSubDomainStatisticsRequest {
	s.SubDomain = &v
	return s
}

func (s *DescribeDohSubDomainStatisticsRequest) Validate() error {
	return dara.Validate(s)
}
