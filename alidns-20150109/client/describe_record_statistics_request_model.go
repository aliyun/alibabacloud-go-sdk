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
	// <props="china">For more information, see [DescribeDomains](https://help.aliyun.com/document_detail/29751.html).
	//
	// <props="intl">For more information, see [DescribeDomains](https://www.alibabacloud.com/help/en/dns/api-alidns-2015-01-09-describedomains).
	//
	// This parameter is required.
	//
	// example:
	//
	// exmaple.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The type of the domain name. This parameter is not case-sensitive. Valid values:
	//
	// - PUBLIC (default): authoritative domain name
	//
	// - CACHE: authoritative proxy domain name
	//
	// example:
	//
	// PUBLIC
	DomainType *string `json:"DomainType,omitempty" xml:"DomainType,omitempty"`
	// The end date of the query. The format is **YYYY-MM-DD**.
	//
	// The default value is the current date.
	//
	// example:
	//
	// 2019-07-04
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// The language of the return value. Valid values:
	//
	// - zh (default): Chinese
	//
	// - en: English
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The host record. For example, to resolve www\\.example.com, set this parameter to www.
	//
	// This parameter is required.
	//
	// example:
	//
	// www
	Rr *string `json:"Rr,omitempty" xml:"Rr,omitempty"`
	// The start date of the query. The format is **YYYY-MM-DD**.
	//
	// The start date must be within the last 90 days.
	//
	// If the time range of the query is 7 days or less, data is returned by the hour.
	//
	// If the time range of the query is more than 7 days, data is returned by the day.
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
