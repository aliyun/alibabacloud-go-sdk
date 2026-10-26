// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMyAppPluginsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategory(v string) *ListMyAppPluginsRequest
	GetCategory() *string
	SetKeyword(v string) *ListMyAppPluginsRequest
	GetKeyword() *string
	SetPageNumber(v int32) *ListMyAppPluginsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListMyAppPluginsRequest
	GetPageSize() *int32
	SetStatus(v string) *ListMyAppPluginsRequest
	GetStatus() *string
}

type ListMyAppPluginsRequest struct {
	// The category used to filter results.
	//
	// example:
	//
	// {\\"Status\\": 0, \\"Name\\": \\"JGHDDLT\\", \\"BizCode\\": \\"\\", \\"CategoryId\\": 30000952447, \\"ParentCategoryId\\": 30000952446}
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The search keyword. The keyword matches the name or description.
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
	// The status used to filter results.
	//
	// example:
	//
	// NORMAL
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListMyAppPluginsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMyAppPluginsRequest) GoString() string {
	return s.String()
}

func (s *ListMyAppPluginsRequest) GetCategory() *string {
	return s.Category
}

func (s *ListMyAppPluginsRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListMyAppPluginsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListMyAppPluginsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMyAppPluginsRequest) GetStatus() *string {
	return s.Status
}

func (s *ListMyAppPluginsRequest) SetCategory(v string) *ListMyAppPluginsRequest {
	s.Category = &v
	return s
}

func (s *ListMyAppPluginsRequest) SetKeyword(v string) *ListMyAppPluginsRequest {
	s.Keyword = &v
	return s
}

func (s *ListMyAppPluginsRequest) SetPageNumber(v int32) *ListMyAppPluginsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListMyAppPluginsRequest) SetPageSize(v int32) *ListMyAppPluginsRequest {
	s.PageSize = &v
	return s
}

func (s *ListMyAppPluginsRequest) SetStatus(v string) *ListMyAppPluginsRequest {
	s.Status = &v
	return s
}

func (s *ListMyAppPluginsRequest) Validate() error {
	return dara.Validate(s)
}
