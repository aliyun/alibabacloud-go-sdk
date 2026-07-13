// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTransferDomainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainNames(v string) *TransferDomainRequest
	GetDomainNames() *string
	SetLang(v string) *TransferDomainRequest
	GetLang() *string
	SetRemark(v string) *TransferDomainRequest
	GetRemark() *string
	SetTargetUserId(v int64) *TransferDomainRequest
	GetTargetUserId() *int64
}

type TransferDomainRequest struct {
	// The domain names to transfer. Separate multiple domain names with commas. Only domain names registered with Alibaba Cloud are supported. <props="china">To get your domain names, call [DescribeDomains](https://help.aliyun.com/en/dns/api-alidns-2015-01-09-describedomains?spm=a2c4g.11186623.help-menu-search-29697.d_0).
	//
	// <props="intl">To get your domain names, call [DescribeDomains](https://www.alibabacloud.com/help/en/dns/api-alidns-2015-01-09-describedomains?spm=a2c63.p38356.help-menu-search-29697.d_0).
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com,example.net
	DomainNames *string `json:"DomainNames,omitempty" xml:"DomainNames,omitempty"`
	// The language of the request and response. Valid values:
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
	// The remarks for the transfer.
	//
	// example:
	//
	// test domain transfer
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The ID of the target user account. The specified domain names and their DNS records are transferred to this account.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12*******
	TargetUserId *int64 `json:"TargetUserId,omitempty" xml:"TargetUserId,omitempty"`
}

func (s TransferDomainRequest) String() string {
	return dara.Prettify(s)
}

func (s TransferDomainRequest) GoString() string {
	return s.String()
}

func (s *TransferDomainRequest) GetDomainNames() *string {
	return s.DomainNames
}

func (s *TransferDomainRequest) GetLang() *string {
	return s.Lang
}

func (s *TransferDomainRequest) GetRemark() *string {
	return s.Remark
}

func (s *TransferDomainRequest) GetTargetUserId() *int64 {
	return s.TargetUserId
}

func (s *TransferDomainRequest) SetDomainNames(v string) *TransferDomainRequest {
	s.DomainNames = &v
	return s
}

func (s *TransferDomainRequest) SetLang(v string) *TransferDomainRequest {
	s.Lang = &v
	return s
}

func (s *TransferDomainRequest) SetRemark(v string) *TransferDomainRequest {
	s.Remark = &v
	return s
}

func (s *TransferDomainRequest) SetTargetUserId(v int64) *TransferDomainRequest {
	s.TargetUserId = &v
	return s
}

func (s *TransferDomainRequest) Validate() error {
	return dara.Validate(s)
}
