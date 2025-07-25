// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainDnssecInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeDomainDnssecInfoRequest
	GetDomainName() *string
	SetLang(v string) *DescribeDomainDnssecInfoRequest
	GetLang() *string
}

type DescribeDomainDnssecInfoRequest struct {
	// The domain name. You can call the [DescribeDomains](https://www.alibabacloud.com/help/zh/dns/api-alidns-2015-01-09-describedomains?spm=a2c63.p38356.help-menu-search-29697.d_0) operation to obtain the domain name.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The language of the content within the request and response. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// Default value: **zh**
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DescribeDomainDnssecInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainDnssecInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainDnssecInfoRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDomainDnssecInfoRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeDomainDnssecInfoRequest) SetDomainName(v string) *DescribeDomainDnssecInfoRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainDnssecInfoRequest) SetLang(v string) *DescribeDomainDnssecInfoRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDomainDnssecInfoRequest) Validate() error {
	return dara.Validate(s)
}
