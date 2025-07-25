// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRecordStatisticsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeRecordStatisticsRequest
	GetDomainName() *string
	SetDomainType(v string) *DescribeRecordStatisticsRequest
	GetDomainType() *string
	SetEndDate(v string) *DescribeRecordStatisticsRequest
	GetEndDate() *string
	SetLang(v string) *DescribeRecordStatisticsRequest
	GetLang() *string
	SetRr(v string) *DescribeRecordStatisticsRequest
	GetRr() *string
	SetStartDate(v string) *DescribeRecordStatisticsRequest
	GetStartDate() *string
}

type DescribeRecordStatisticsRequest struct {
	// The primary domain name.
	//
	// This parameter is required.
	//
	// example:
	//
	// dns-example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The type of the domain name. The parameter value is not case-sensitive. Valid values:
	//
	// 	- PUBLIC (default): hosted public domain name
	//
	// 	- CACHE: cache-accelerated domain name
	//
	// example:
	//
	// PUBLIC
	DomainType *string `json:"DomainType,omitempty" xml:"DomainType,omitempty"`
	// The end date of the query. Specify the end date in the **YYYY-MM-DD*	- format.
	//
	// The default value is the day when you query the data.
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
	// The hostname. If you want to resolve www.dns-exmaple.top, set Rr to www.
	//
	// This parameter is required.
	//
	// example:
	//
	// www
	Rr *string `json:"Rr,omitempty" xml:"Rr,omitempty"`
	// The start date of the query. Specify the start date in the **YYYY-MM-DD*	- format.
	//
	// You can only query the DNS records within the last 90 days.``
	//
	// If the time range is less than or equal to seven days, data is returned on an hourly basis.````
	//
	// If the time range is greater than seven days, data is returned on a daily basis.````
	//
	// This parameter is required.
	//
	// example:
	//
	// 2019-07-04
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s DescribeRecordStatisticsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRecordStatisticsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRecordStatisticsRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeRecordStatisticsRequest) GetDomainType() *string {
	return s.DomainType
}

func (s *DescribeRecordStatisticsRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *DescribeRecordStatisticsRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeRecordStatisticsRequest) GetRr() *string {
	return s.Rr
}

func (s *DescribeRecordStatisticsRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *DescribeRecordStatisticsRequest) SetDomainName(v string) *DescribeRecordStatisticsRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeRecordStatisticsRequest) SetDomainType(v string) *DescribeRecordStatisticsRequest {
	s.DomainType = &v
	return s
}

func (s *DescribeRecordStatisticsRequest) SetEndDate(v string) *DescribeRecordStatisticsRequest {
	s.EndDate = &v
	return s
}

func (s *DescribeRecordStatisticsRequest) SetLang(v string) *DescribeRecordStatisticsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeRecordStatisticsRequest) SetRr(v string) *DescribeRecordStatisticsRequest {
	s.Rr = &v
	return s
}

func (s *DescribeRecordStatisticsRequest) SetStartDate(v string) *DescribeRecordStatisticsRequest {
	s.StartDate = &v
	return s
}

func (s *DescribeRecordStatisticsRequest) Validate() error {
	return dara.Validate(s)
}
