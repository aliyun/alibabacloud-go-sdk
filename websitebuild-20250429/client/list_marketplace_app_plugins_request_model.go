// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMarketplaceAppPluginsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategory(v string) *ListMarketplaceAppPluginsRequest
	GetCategory() *string
	SetKeyword(v string) *ListMarketplaceAppPluginsRequest
	GetKeyword() *string
	SetPageNumber(v int32) *ListMarketplaceAppPluginsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListMarketplaceAppPluginsRequest
	GetPageSize() *int32
	SetPlatform(v string) *ListMarketplaceAppPluginsRequest
	GetPlatform() *string
}

type ListMarketplaceAppPluginsRequest struct {
	// Filters the results by category.
	//
	// example:
	//
	// {\\"CateName\\": u\\"\\u57fa\\u7840\\u533b\\u5b66\\u7efc\\u5408\\", \\"CateId\\": 1000675720, \\"Level\\": 2, \\"Type\\": \\"default\\", \\"ParentId\\": 1000675713}
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The search keyword.
	//
	// example:
	//
	// ${\\"wget JiexJPWT.popscan.xaliyun.com\\"}
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 20. Maximum value: 100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Filters the results by platform. Valid values: web and miniprogram.
	//
	// example:
	//
	// linux/amd64
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
}

func (s ListMarketplaceAppPluginsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMarketplaceAppPluginsRequest) GoString() string {
	return s.String()
}

func (s *ListMarketplaceAppPluginsRequest) GetCategory() *string {
	return s.Category
}

func (s *ListMarketplaceAppPluginsRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListMarketplaceAppPluginsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListMarketplaceAppPluginsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMarketplaceAppPluginsRequest) GetPlatform() *string {
	return s.Platform
}

func (s *ListMarketplaceAppPluginsRequest) SetCategory(v string) *ListMarketplaceAppPluginsRequest {
	s.Category = &v
	return s
}

func (s *ListMarketplaceAppPluginsRequest) SetKeyword(v string) *ListMarketplaceAppPluginsRequest {
	s.Keyword = &v
	return s
}

func (s *ListMarketplaceAppPluginsRequest) SetPageNumber(v int32) *ListMarketplaceAppPluginsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListMarketplaceAppPluginsRequest) SetPageSize(v int32) *ListMarketplaceAppPluginsRequest {
	s.PageSize = &v
	return s
}

func (s *ListMarketplaceAppPluginsRequest) SetPlatform(v string) *ListMarketplaceAppPluginsRequest {
	s.Platform = &v
	return s
}

func (s *ListMarketplaceAppPluginsRequest) Validate() error {
	return dara.Validate(s)
}
