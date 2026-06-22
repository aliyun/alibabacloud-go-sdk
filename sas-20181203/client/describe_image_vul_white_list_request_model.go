// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImageVulWhiteListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliasName(v string) *DescribeImageVulWhiteListRequest
	GetAliasName() *string
	SetCurrentPage(v int32) *DescribeImageVulWhiteListRequest
	GetCurrentPage() *int32
	SetLang(v string) *DescribeImageVulWhiteListRequest
	GetLang() *string
	SetPageSize(v int32) *DescribeImageVulWhiteListRequest
	GetPageSize() *int32
	SetSource(v string) *DescribeImageVulWhiteListRequest
	GetSource() *string
}

type DescribeImageVulWhiteListRequest struct {
	// The alias of the vulnerability to query.
	//
	// example:
	//
	// CVE-2007-5686:rpath_linux
	AliasName *string `json:"AliasName,omitempty" xml:"AliasName,omitempty"`
	// The page number of the current page in a paged query. Paging is used to display results.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The language type for the request and response. Default value: **zh**. Valid values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The maximum number of entries per page in a paged query. Paging is used to display results. Default value: 20.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The source of the whitelist. Valid values:
	//
	// - **image**: image
	//
	// - **agentless**: agentless.
	//
	// example:
	//
	// image
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s DescribeImageVulWhiteListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageVulWhiteListRequest) GoString() string {
	return s.String()
}

func (s *DescribeImageVulWhiteListRequest) GetAliasName() *string {
	return s.AliasName
}

func (s *DescribeImageVulWhiteListRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeImageVulWhiteListRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeImageVulWhiteListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeImageVulWhiteListRequest) GetSource() *string {
	return s.Source
}

func (s *DescribeImageVulWhiteListRequest) SetAliasName(v string) *DescribeImageVulWhiteListRequest {
	s.AliasName = &v
	return s
}

func (s *DescribeImageVulWhiteListRequest) SetCurrentPage(v int32) *DescribeImageVulWhiteListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeImageVulWhiteListRequest) SetLang(v string) *DescribeImageVulWhiteListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeImageVulWhiteListRequest) SetPageSize(v int32) *DescribeImageVulWhiteListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeImageVulWhiteListRequest) SetSource(v string) *DescribeImageVulWhiteListRequest {
	s.Source = &v
	return s
}

func (s *DescribeImageVulWhiteListRequest) Validate() error {
	return dara.Validate(s)
}
