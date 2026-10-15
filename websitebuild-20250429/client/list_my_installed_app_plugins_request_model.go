// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMyInstalledAppPluginsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategory(v string) *ListMyInstalledAppPluginsRequest
	GetCategory() *string
	SetKeyword(v string) *ListMyInstalledAppPluginsRequest
	GetKeyword() *string
	SetPageNumber(v int32) *ListMyInstalledAppPluginsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListMyInstalledAppPluginsRequest
	GetPageSize() *int32
}

type ListMyInstalledAppPluginsRequest struct {
	// Filters results by category.
	//
	// example:
	//
	// {\\"Status\\": 0, \\"Name\\": \\"UELBNUP\\", \\"BizCode\\": \\"\\", \\"CategoryId\\": 30000655368, \\"ParentCategoryId\\": 30000655367}
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The search keyword. Matches against name or description.
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
}

func (s ListMyInstalledAppPluginsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMyInstalledAppPluginsRequest) GoString() string {
	return s.String()
}

func (s *ListMyInstalledAppPluginsRequest) GetCategory() *string {
	return s.Category
}

func (s *ListMyInstalledAppPluginsRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListMyInstalledAppPluginsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListMyInstalledAppPluginsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMyInstalledAppPluginsRequest) SetCategory(v string) *ListMyInstalledAppPluginsRequest {
	s.Category = &v
	return s
}

func (s *ListMyInstalledAppPluginsRequest) SetKeyword(v string) *ListMyInstalledAppPluginsRequest {
	s.Keyword = &v
	return s
}

func (s *ListMyInstalledAppPluginsRequest) SetPageNumber(v int32) *ListMyInstalledAppPluginsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListMyInstalledAppPluginsRequest) SetPageSize(v int32) *ListMyInstalledAppPluginsRequest {
	s.PageSize = &v
	return s
}

func (s *ListMyInstalledAppPluginsRequest) Validate() error {
	return dara.Validate(s)
}
