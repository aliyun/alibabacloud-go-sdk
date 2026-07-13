// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRetrieveDomainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *RetrieveDomainRequest
	GetDomainName() *string
	SetLang(v string) *RetrieveDomainRequest
	GetLang() *string
}

type RetrieveDomainRequest struct {
	// The domain name.<props="china"> Call [DescribeDomains](https://help.aliyun.com/zh/dns/api-alidns-2015-01-09-describedomains?spm=a2c4g.11186623.help-menu-search-29697.d_0) to obtain the domain name.
	//
	// <props="intl">Call [DescribeDomains](https://www.alibabacloud.com/help/zh/dns/api-alidns-2015-01-09-describedomains?spm=a2c63.p38356.help-menu-search-29697.d_0) to obtain the domain name.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The language of the return value. Valid values:
	//
	// - zh: Chinese
	//
	// - en: English
	//
	// Default value: zh.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s RetrieveDomainRequest) String() string {
	return dara.Prettify(s)
}

func (s RetrieveDomainRequest) GoString() string {
	return s.String()
}

func (s *RetrieveDomainRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *RetrieveDomainRequest) GetLang() *string {
	return s.Lang
}

func (s *RetrieveDomainRequest) SetDomainName(v string) *RetrieveDomainRequest {
	s.DomainName = &v
	return s
}

func (s *RetrieveDomainRequest) SetLang(v string) *RetrieveDomainRequest {
	s.Lang = &v
	return s
}

func (s *RetrieveDomainRequest) Validate() error {
	return dara.Validate(s)
}
