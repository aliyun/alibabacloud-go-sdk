// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceDomainsRequest interface {
	dara.Model
	String() string
	GoString() string
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
	// The ID of the Alibaba Cloud Domain Name System (DNS) instance. You can call the [DescribeDomainInfo](https://www.alibabacloud.com/help/zh/dns/api-alidns-2015-01-09-describedomaininfo?spm=a2c63.p38356.help-menu-search-29697.d_0) operation to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// weriwieru
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The language of the response. Valid values:
	//
	// 	- zh: Chinese
	//
	// 	- en: English
	//
	// Default value: English.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The number of the page to return. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: 1 to 100. Default value: 20.
	//
	// example:
	//
	// 2
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeInstanceDomainsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceDomainsRequest) GoString() string {
	return s.String()
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
