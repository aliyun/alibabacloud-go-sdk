// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDomainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DeleteDomainRequest
	GetDomainName() *string
	SetLang(v string) *DeleteDomainRequest
	GetLang() *string
}

type DeleteDomainRequest struct {
	// The domain name that already exists in Alibaba Cloud DNS. You can call the [DescribeDomains](https://www.alibabacloud.com/help/zh/dns/api-alidns-2015-01-09-describedomains?spm=a2c63.p38356.help-menu-search-29697.d_0) operation to obtain the domain name.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The language of the response. Valid values:
	//
	// 	- zh: Chinese
	//
	// 	- en: English
	//
	// Default value: en.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DeleteDomainRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDomainRequest) GoString() string {
	return s.String()
}

func (s *DeleteDomainRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DeleteDomainRequest) GetLang() *string {
	return s.Lang
}

func (s *DeleteDomainRequest) SetDomainName(v string) *DeleteDomainRequest {
	s.DomainName = &v
	return s
}

func (s *DeleteDomainRequest) SetLang(v string) *DeleteDomainRequest {
	s.Lang = &v
	return s
}

func (s *DeleteDomainRequest) Validate() error {
	return dara.Validate(s)
}
