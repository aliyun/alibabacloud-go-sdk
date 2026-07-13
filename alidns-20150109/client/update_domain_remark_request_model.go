// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDomainRemarkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *UpdateDomainRemarkRequest
	GetDomainName() *string
	SetLang(v string) *UpdateDomainRemarkRequest
	GetLang() *string
	SetRemark(v string) *UpdateDomainRemarkRequest
	GetRemark() *string
}

type UpdateDomainRemarkRequest struct {
	// An existing domain name in Alibaba Cloud DNS.<props="china">For more information, see [DescribeDomains ](https://help.aliyun.com/zh/dns/api-alidns-2015-01-09-describedomains?spm=a2c4g.11186623.help-menu-search-29697.d_0).
	//
	// <props="intl">For more information, see [DescribeDomains ](https://www.alibabacloud.com/help/zh/dns/api-alidns-2015-01-09-describedomains?spm=a2c63.p38356.help-menu-search-29697.d_0).
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The language of the response. Valid values:
	//
	// - zh: Chinese
	//
	// - en: English
	//
	// Default: en
	//
	// example:
	//
	// cn
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The remarks for the domain name.
	//
	// The remarks can be up to 50 characters in length. They can contain digits, letters, Chinese characters, and the following special characters: _, -, ,, and .
	//
	// example:
	//
	// 这是我在阿里云解析的第一个域名
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s UpdateDomainRemarkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDomainRemarkRequest) GoString() string {
	return s.String()
}

func (s *UpdateDomainRemarkRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *UpdateDomainRemarkRequest) GetLang() *string {
	return s.Lang
}

func (s *UpdateDomainRemarkRequest) GetRemark() *string {
	return s.Remark
}

func (s *UpdateDomainRemarkRequest) SetDomainName(v string) *UpdateDomainRemarkRequest {
	s.DomainName = &v
	return s
}

func (s *UpdateDomainRemarkRequest) SetLang(v string) *UpdateDomainRemarkRequest {
	s.Lang = &v
	return s
}

func (s *UpdateDomainRemarkRequest) SetRemark(v string) *UpdateDomainRemarkRequest {
	s.Remark = &v
	return s
}

func (s *UpdateDomainRemarkRequest) Validate() error {
	return dara.Validate(s)
}
