// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainNsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeDomainNsRequest
	GetDomainName() *string
	SetLang(v string) *DescribeDomainNsRequest
	GetLang() *string
}

type DescribeDomainNsRequest struct {
	// The domain name. You can call the [DescribeDomains](https://www.alibabacloud.com/help/zh/dns/api-alidns-2015-01-09-describedomains?spm=a2c63.p38356.help-menu-search-29697.d_0) operation to obtain the domain name.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The language of the content in the request and response.
	//
	// Valid values:
	//
	// zh: Chinese\\
	//
	// en: English
	//
	// Default value: **zh**
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DescribeDomainNsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainNsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainNsRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDomainNsRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeDomainNsRequest) SetDomainName(v string) *DescribeDomainNsRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainNsRequest) SetLang(v string) *DescribeDomainNsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDomainNsRequest) Validate() error {
	return dara.Validate(s)
}
