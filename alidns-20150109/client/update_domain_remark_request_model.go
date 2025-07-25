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
	// The domain name that already exists in Alibaba Cloud Domain Name System (DNS). You can call the [DescribeDomains ](https://www.alibabacloud.com/help/zh/dns/api-alidns-2015-01-09-describedomains?spm=a2c63.p38356.help-menu-search-29697.d_0)operation to obtain the domain name.
	//
	// This parameter is required.
	//
	// example:
	//
	// mydomain.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The language of the response. Valid values:
	//
	// 	- zh: Chinese
	//
	// 	- en: English
	//
	// Default value: en
	//
	// example:
	//
	// cn
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The description of the domain name.
	//
	// It can be up to 50 characters in length and can contain digits, letters, and the following special characters: _ - , .
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
