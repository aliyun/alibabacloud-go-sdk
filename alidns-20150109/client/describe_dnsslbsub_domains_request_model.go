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
	// The domain name. You can call the [DescribeDomains](https://www.alibabacloud.com/help/zh/dns/api-alidns-2015-01-09-describedomains?spm=a2c63.p38356.help-menu-search-29697.d_0) operation to obtain the domain name.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The language of the response. Valid values:
	//
	// 	- **zh*	- (default): Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The page number. Pages start from page **1**. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: **1 to 100**. Default value: **20**.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The hostname.
	//
	// example:
	//
	// test
	Rr *string `json:"Rr,omitempty" xml:"Rr,omitempty"`
	// The IP address of the user account.
	//
	// example:
	//
	// 1.1.1.1
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
