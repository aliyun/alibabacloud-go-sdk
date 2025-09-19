// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVulWhitelistRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeVulWhitelistRequest
	GetCurrentPage() *int32
	SetLang(v string) *DescribeVulWhitelistRequest
	GetLang() *string
	SetPageSize(v int32) *DescribeVulWhitelistRequest
	GetPageSize() *int32
}

type DescribeVulWhitelistRequest struct {
	// The number of the page to return. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The language of the content within the request and response. Default value: **zh**. Valid values:
	//
	// 	- **zh**: Chinese.
	//
	// 	- **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The number of entries to return on each page. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeVulWhitelistRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVulWhitelistRequest) GoString() string {
	return s.String()
}

func (s *DescribeVulWhitelistRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeVulWhitelistRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeVulWhitelistRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeVulWhitelistRequest) SetCurrentPage(v int32) *DescribeVulWhitelistRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeVulWhitelistRequest) SetLang(v string) *DescribeVulWhitelistRequest {
	s.Lang = &v
	return s
}

func (s *DescribeVulWhitelistRequest) SetPageSize(v int32) *DescribeVulWhitelistRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeVulWhitelistRequest) Validate() error {
	return dara.Validate(s)
}
