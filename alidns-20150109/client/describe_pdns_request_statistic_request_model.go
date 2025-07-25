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
	// The primary domain name whose statistics you want to query.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end of the time range to query. Specify the time in the **YYYY-MM-DD*	- format.
	//
	// The default value is the day when you query the data.
	//
	// example:
	//
	// 2024-7-1 00:00:00
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// The language of the content within the request and response. Default value: **zh**. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The beginning of the time range to query. Specify the time in the **YYYY-MM-DD*	- format.
	//
	// You can query only records of the last 90 days.
	//
	// example:
	//
	// 2024-06-14 00:00:00
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// The subdomain name whose statistics you want to query.
	//
	// example:
	//
	// www.example.com
	SubDomain *string `json:"SubDomain,omitempty" xml:"SubDomain,omitempty"`
	// The type of the request statistics that you want to query. Valid values:
	//
	// 	- **ACCOUNT**: queries the request statistics by account.
	//
	// 	- **DOMAIN**: queries the request statistics by domain name.
	//
	// 	- **SUB_DOMAIN**: queries the request statistics by subdomain name.
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
