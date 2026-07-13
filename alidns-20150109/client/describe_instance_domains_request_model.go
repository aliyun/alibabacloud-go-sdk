// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceDomainsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainKeywords(v string) *DescribeInstanceDomainsRequest
	GetDomainKeywords() *string
	SetInstanceId(v string) *DescribeInstanceDomainsRequest
	GetInstanceId() *string
	SetLang(v string) *DescribeInstanceDomainsRequest
	GetLang() *string
	SetPageNumber(v int64) *DescribeInstanceDomainsRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeInstanceDomainsRequest
	GetPageSize() *int64
}

type DescribeInstanceDomainsRequest struct {
	DomainKeywords *string `json:"DomainKeywords,omitempty" xml:"DomainKeywords,omitempty"`
	// The instance ID.<props="china"> You can call [DescribeDomainInfo](https://help.aliyun.com/zh/dns/api-alidns-2015-01-09-describedomaininfo?spm=a2c4g.11186623.help-menu-search-29697.d_0) to obtain the ID.
	//
	// <props="intl">You can call [DescribeDomainInfo](https://www.alibabacloud.com/help/zh/dns/api-alidns-2015-01-09-describedomaininfo?spm=a2c63.p38356.help-menu-search-29697.d_0) to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// dns-cn-9lb38ldq9**
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
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
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The page number. The value starts from 1. Default: 1.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Maximum value: 100. Default value: 20.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeInstanceDomainsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceDomainsRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceDomainsRequest) GetDomainKeywords() *string {
	return s.DomainKeywords
}

func (s *DescribeInstanceDomainsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeInstanceDomainsRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeInstanceDomainsRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeInstanceDomainsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeInstanceDomainsRequest) SetDomainKeywords(v string) *DescribeInstanceDomainsRequest {
	s.DomainKeywords = &v
	return s
}

func (s *DescribeInstanceDomainsRequest) SetInstanceId(v string) *DescribeInstanceDomainsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceDomainsRequest) SetLang(v string) *DescribeInstanceDomainsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeInstanceDomainsRequest) SetPageNumber(v int64) *DescribeInstanceDomainsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeInstanceDomainsRequest) SetPageSize(v int64) *DescribeInstanceDomainsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeInstanceDomainsRequest) Validate() error {
	return dara.Validate(s)
}
