// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCustomLinesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeCustomLinesRequest
	GetDomainName() *string
	SetLang(v string) *DescribeCustomLinesRequest
	GetLang() *string
	SetPageNumber(v int64) *DescribeCustomLinesRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeCustomLinesRequest
	GetPageSize() *int64
}

type DescribeCustomLinesRequest struct {
	// The domain name that already exists in Alibaba Cloud Domain Name System (DNS). You can call the [DescribeDomains ](https://www.alibabacloud.com/help/zh/dns/api-alidns-2015-01-09-describedomains?spm=a2c63.p38356.help-menu-search-29697.d_0)operation to obtain the domain name.
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
	// Default value: en
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeCustomLinesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomLinesRequest) GoString() string {
	return s.String()
}

func (s *DescribeCustomLinesRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeCustomLinesRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeCustomLinesRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeCustomLinesRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeCustomLinesRequest) SetDomainName(v string) *DescribeCustomLinesRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeCustomLinesRequest) SetLang(v string) *DescribeCustomLinesRequest {
	s.Lang = &v
	return s
}

func (s *DescribeCustomLinesRequest) SetPageNumber(v int64) *DescribeCustomLinesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeCustomLinesRequest) SetPageSize(v int64) *DescribeCustomLinesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCustomLinesRequest) Validate() error {
	return dara.Validate(s)
}
