// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePdnsRequestStatisticRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribePdnsRequestStatisticRequest
	GetDomainName() *string
	SetEndDate(v string) *DescribePdnsRequestStatisticRequest
	GetEndDate() *string
	SetLang(v string) *DescribePdnsRequestStatisticRequest
	GetLang() *string
	SetStartDate(v string) *DescribePdnsRequestStatisticRequest
	GetStartDate() *string
	SetSubDomain(v string) *DescribePdnsRequestStatisticRequest
	GetSubDomain() *string
	SetType(v string) *DescribePdnsRequestStatisticRequest
	GetType() *string
}

type DescribePdnsRequestStatisticRequest struct {
	// The primary domain name for which you want to query statistics.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end date of the query in the **YYYY-MM-DD*	- format.
	//
	// The default value is the current day.
	//
	// example:
	//
	// 2024-07-01
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// The language of the request and response. The default value is **zh**. Valid values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The start date of the query in the **YYYY-MM-DD*	- format.
	//
	// You can query data from the last 90 days.
	//
	// example:
	//
	// 2024-06-14
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// The subdomain for which you want to query statistics.
	//
	// example:
	//
	// www.example.com
	SubDomain *string `json:"SubDomain,omitempty" xml:"SubDomain,omitempty"`
	// The dimension for statistics. Valid values:
	//
	// - **ACCOUNT**: queries statistics by account.
	//
	// - **DOMAIN**: queries statistics by domain name. The DomainName parameter is required.
	//
	// - **SUB_DOMAIN**: queries statistics by subdomain. The DomainName and SubDomain parameters are required.
	//
	// example:
	//
	// ACCOUNT
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribePdnsRequestStatisticRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePdnsRequestStatisticRequest) GoString() string {
	return s.String()
}

func (s *DescribePdnsRequestStatisticRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribePdnsRequestStatisticRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *DescribePdnsRequestStatisticRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribePdnsRequestStatisticRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *DescribePdnsRequestStatisticRequest) GetSubDomain() *string {
	return s.SubDomain
}

func (s *DescribePdnsRequestStatisticRequest) GetType() *string {
	return s.Type
}

func (s *DescribePdnsRequestStatisticRequest) SetDomainName(v string) *DescribePdnsRequestStatisticRequest {
	s.DomainName = &v
	return s
}

func (s *DescribePdnsRequestStatisticRequest) SetEndDate(v string) *DescribePdnsRequestStatisticRequest {
	s.EndDate = &v
	return s
}

func (s *DescribePdnsRequestStatisticRequest) SetLang(v string) *DescribePdnsRequestStatisticRequest {
	s.Lang = &v
	return s
}

func (s *DescribePdnsRequestStatisticRequest) SetStartDate(v string) *DescribePdnsRequestStatisticRequest {
	s.StartDate = &v
	return s
}

func (s *DescribePdnsRequestStatisticRequest) SetSubDomain(v string) *DescribePdnsRequestStatisticRequest {
	s.SubDomain = &v
	return s
}

func (s *DescribePdnsRequestStatisticRequest) SetType(v string) *DescribePdnsRequestStatisticRequest {
	s.Type = &v
	return s
}

func (s *DescribePdnsRequestStatisticRequest) Validate() error {
	return dara.Validate(s)
}
