// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainStatisticsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeDomainStatisticsRequest
	GetDomainName() *string
	SetDomainType(v string) *DescribeDomainStatisticsRequest
	GetDomainType() *string
	SetEndDate(v string) *DescribeDomainStatisticsRequest
	GetEndDate() *string
	SetLang(v string) *DescribeDomainStatisticsRequest
	GetLang() *string
	SetStartDate(v string) *DescribeDomainStatisticsRequest
	GetStartDate() *string
}

type DescribeDomainStatisticsRequest struct {
	// The domain name.<props="china"> Call the [DescribeDomains](https://help.aliyun.com/zh/dns/api-alidns-2015-01-09-describedomains?spm=a2c4g.11186623.help-menu-search-29697.d_0) operation to obtain the domain name.
	//
	// <props="intl">Call the [DescribeDomains](https://www.alibabacloud.com/help/zh/dns/api-alidns-2015-01-09-describedomains?spm=a2c63.p38356.help-menu-search-29697.d_0) operation to obtain the domain name.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The type of the domain name. Valid values:
	//
	// - PUBLIC: The domain name is hosted on Alibaba Cloud DNS (default).
	//
	// - CACHE: Alibaba Cloud DNS is used as a proxy for the domain name.
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
	// The language of the request and response.
	//
	// - **zh*	- (default): Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The start date of the query. The format is **YYYY-MM-DD**.
	//
	// You can query records only from the last 90 days. This means `StartDate >= Now - 90`.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2019-07-04
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s DescribeDomainStatisticsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainStatisticsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainStatisticsRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDomainStatisticsRequest) GetDomainType() *string {
	return s.DomainType
}

func (s *DescribeDomainStatisticsRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *DescribeDomainStatisticsRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeDomainStatisticsRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *DescribeDomainStatisticsRequest) SetDomainName(v string) *DescribeDomainStatisticsRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainStatisticsRequest) SetDomainType(v string) *DescribeDomainStatisticsRequest {
	s.DomainType = &v
	return s
}

func (s *DescribeDomainStatisticsRequest) SetEndDate(v string) *DescribeDomainStatisticsRequest {
	s.EndDate = &v
	return s
}

func (s *DescribeDomainStatisticsRequest) SetLang(v string) *DescribeDomainStatisticsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDomainStatisticsRequest) SetStartDate(v string) *DescribeDomainStatisticsRequest {
	s.StartDate = &v
	return s
}

func (s *DescribeDomainStatisticsRequest) Validate() error {
	return dara.Validate(s)
}
