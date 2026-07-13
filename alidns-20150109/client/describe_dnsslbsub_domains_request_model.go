// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDNSSLBSubDomainsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeDNSSLBSubDomainsRequest
	GetDomainName() *string
	SetLang(v string) *DescribeDNSSLBSubDomainsRequest
	GetLang() *string
	SetPageNumber(v int64) *DescribeDNSSLBSubDomainsRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeDNSSLBSubDomainsRequest
	GetPageSize() *int64
	SetRr(v string) *DescribeDNSSLBSubDomainsRequest
	GetRr() *string
	SetUserClientIp(v string) *DescribeDNSSLBSubDomainsRequest
	GetUserClientIp() *string
}

type DescribeDNSSLBSubDomainsRequest struct {
	// The domain name. Call the [DescribeDomains](https://www.alibabacloud.com/help/en/dns/api-alidns-2015-01-09-describedomains) operation to obtain the domain name.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The language of the response. Valid values are:
	//
	// - **zh**: Chinese. This is the default value.
	//
	// - **en**: English.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The page number. The value starts from **1**. The default value is **1**.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. The maximum value is **100**. The default value is **20**.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The host record.
	//
	// example:
	//
	// test
	Rr *string `json:"Rr,omitempty" xml:"Rr,omitempty"`
	// The IP address of the client.
	//
	// example:
	//
	// 1.1.XX.XX
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s DescribeDNSSLBSubDomainsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDNSSLBSubDomainsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDNSSLBSubDomainsRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDNSSLBSubDomainsRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeDNSSLBSubDomainsRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeDNSSLBSubDomainsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeDNSSLBSubDomainsRequest) GetRr() *string {
	return s.Rr
}

func (s *DescribeDNSSLBSubDomainsRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *DescribeDNSSLBSubDomainsRequest) SetDomainName(v string) *DescribeDNSSLBSubDomainsRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDNSSLBSubDomainsRequest) SetLang(v string) *DescribeDNSSLBSubDomainsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDNSSLBSubDomainsRequest) SetPageNumber(v int64) *DescribeDNSSLBSubDomainsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDNSSLBSubDomainsRequest) SetPageSize(v int64) *DescribeDNSSLBSubDomainsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDNSSLBSubDomainsRequest) SetRr(v string) *DescribeDNSSLBSubDomainsRequest {
	s.Rr = &v
	return s
}

func (s *DescribeDNSSLBSubDomainsRequest) SetUserClientIp(v string) *DescribeDNSSLBSubDomainsRequest {
	s.UserClientIp = &v
	return s
}

func (s *DescribeDNSSLBSubDomainsRequest) Validate() error {
	return dara.Validate(s)
}
