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
	SetResourceDirectoryAccountId(v int64) *DescribeVulWhitelistRequest
	GetResourceDirectoryAccountId() *int64
}

type DescribeVulWhitelistRequest struct {
	// The page number when paging. Minimum value: 1. Default value: 1.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The language type for the request and response messages. Default value: **zh**. Valid values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The number of entries per page when paging. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The Alibaba Cloud account ID of the member accounts in the resource directory.
	//
	// >Invoke the [DescribeMonitorAccounts](~~DescribeMonitorAccounts~~) operation to obtain this parameter.
	ResourceDirectoryAccountId *int64 `json:"ResourceDirectoryAccountId,omitempty" xml:"ResourceDirectoryAccountId,omitempty"`
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

func (s *DescribeVulWhitelistRequest) GetResourceDirectoryAccountId() *int64 {
	return s.ResourceDirectoryAccountId
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

func (s *DescribeVulWhitelistRequest) SetResourceDirectoryAccountId(v int64) *DescribeVulWhitelistRequest {
	s.ResourceDirectoryAccountId = &v
	return s
}

func (s *DescribeVulWhitelistRequest) Validate() error {
	return dara.Validate(s)
}
