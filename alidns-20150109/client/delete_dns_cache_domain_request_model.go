// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDnsCacheDomainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DeleteDnsCacheDomainRequest
	GetDomainName() *string
	SetLang(v string) *DeleteDnsCacheDomainRequest
	GetLang() *string
}

type DeleteDnsCacheDomainRequest struct {
	// The domain name.<props="china"> Call [DescribeDomains](https://help.aliyun.com/en/dns/api-alidns-2015-01-09-describedomains?spm=a2c4g.11186623.help-menu-search-29697.d_0) to obtain this value.
	//
	// <props="intl">Call [DescribeDomains](https://www.alibabacloud.com/help/en/dns/api-alidns-2015-01-09-describedomains?spm=a2c63.p38356.help-menu-search-29697.d_0) to obtain this value.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The language of the request and response. Valid values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	//   The default value is **zh**.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DeleteDnsCacheDomainRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDnsCacheDomainRequest) GoString() string {
	return s.String()
}

func (s *DeleteDnsCacheDomainRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DeleteDnsCacheDomainRequest) GetLang() *string {
	return s.Lang
}

func (s *DeleteDnsCacheDomainRequest) SetDomainName(v string) *DeleteDnsCacheDomainRequest {
	s.DomainName = &v
	return s
}

func (s *DeleteDnsCacheDomainRequest) SetLang(v string) *DeleteDnsCacheDomainRequest {
	s.Lang = &v
	return s
}

func (s *DeleteDnsCacheDomainRequest) Validate() error {
	return dara.Validate(s)
}
