// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDnsCacheDomainRemarkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *UpdateDnsCacheDomainRemarkRequest
	GetDomainName() *string
	SetLang(v string) *UpdateDnsCacheDomainRemarkRequest
	GetLang() *string
	SetRemark(v string) *UpdateDnsCacheDomainRemarkRequest
	GetRemark() *string
}

type UpdateDnsCacheDomainRemarkRequest struct {
	// The domain name. You can call the [DescribeDomains](https://www.alibabacloud.com/help/zh/dns/api-alidns-2015-01-09-describedomains?spm=a2c63.p38356.help-menu-search-29697.d_0) operation to obtain the domain name.
	//
	// This parameter is required.
	//
	// example:
	//
	// dns-example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The language of the content within the request and response. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// Default: **zh**
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The remarks. The remarks can be up to 50 characters in length and can contain only letters, digits, periods (.), underscores (_), and hyphens (-).
	//
	// Delete the original remarks when it is empty.
	//
	// example:
	//
	// test
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s UpdateDnsCacheDomainRemarkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDnsCacheDomainRemarkRequest) GoString() string {
	return s.String()
}

func (s *UpdateDnsCacheDomainRemarkRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *UpdateDnsCacheDomainRemarkRequest) GetLang() *string {
	return s.Lang
}

func (s *UpdateDnsCacheDomainRemarkRequest) GetRemark() *string {
	return s.Remark
}

func (s *UpdateDnsCacheDomainRemarkRequest) SetDomainName(v string) *UpdateDnsCacheDomainRemarkRequest {
	s.DomainName = &v
	return s
}

func (s *UpdateDnsCacheDomainRemarkRequest) SetLang(v string) *UpdateDnsCacheDomainRemarkRequest {
	s.Lang = &v
	return s
}

func (s *UpdateDnsCacheDomainRemarkRequest) SetRemark(v string) *UpdateDnsCacheDomainRemarkRequest {
	s.Remark = &v
	return s
}

func (s *UpdateDnsCacheDomainRemarkRequest) Validate() error {
	return dara.Validate(s)
}
