// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCatalogsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *ListCatalogsShrinkRequest
	GetComment() *string
	SetName(v string) *ListCatalogsShrinkRequest
	GetName() *string
	SetOrder(v string) *ListCatalogsShrinkRequest
	GetOrder() *string
	SetPageNumber(v int32) *ListCatalogsShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListCatalogsShrinkRequest
	GetPageSize() *int32
	SetParentMetaEntityId(v string) *ListCatalogsShrinkRequest
	GetParentMetaEntityId() *string
	SetSortBy(v string) *ListCatalogsShrinkRequest
	GetSortBy() *string
	SetTypesShrink(v string) *ListCatalogsShrinkRequest
	GetTypesShrink() *string
}

type ListCatalogsShrinkRequest struct {
	// The comment. Supports token-based matching.
	//
	// example:
	//
	// this is a comment
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The name. Supports fuzzy matching.
	//
	// example:
	//
	// abc
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The order in which the tables are sorted. Default value: Asc. Valid values:
	//
	// 	- Asc: ascending order.
	//
	// 	- Desc: descending order.
	//
	// example:
	//
	// Asc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of records per page. Default value: 10. Maximum value: 100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The parent entity ID. For more information, see [Concepts related to metadata entities](https://help.aliyun.com/document_detail/2880092.html).
	//
	// Currently, only the DLF and StarRocks types are supported.
	//
	// 	- For the DLF type, you can query all catalog lists. The format of `ParentMetaEntityId` is `DLF`.
	//
	// 	- For the StarRocks type, you can query the catalogs of a specific instance. The format of `ParentMetaEntityId` `is StarRocks:(instance_id|encoded_jdbc_url)`.
	//
	// > \\
	//
	// `instance_id`: The instance ID. Required if the data source is registered in instance mode.\\
	//
	// `encoded_jdbc_url`: The JDBC connection string encoded with URL encoding. Required if the data source is registered in connection-string mode.
	//
	// This parameter is required.
	//
	// example:
	//
	// dlf
	//
	// starrocks:c-abc123xxx
	ParentMetaEntityId *string `json:"ParentMetaEntityId,omitempty" xml:"ParentMetaEntityId,omitempty"`
	// The sort field. Default value: CreateTime. Valid values:
	//
	// 	- CreateTime
	//
	// 	- ModifyTime
	//
	// 	- Name
	//
	// 	- Type
	//
	// example:
	//
	// CreateTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The type. Supports exact match. If left empty, all types are queried.
	TypesShrink *string `json:"Types,omitempty" xml:"Types,omitempty"`
}

func (s ListCatalogsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCatalogsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListCatalogsShrinkRequest) GetComment() *string {
	return s.Comment
}

func (s *ListCatalogsShrinkRequest) GetName() *string {
	return s.Name
}

func (s *ListCatalogsShrinkRequest) GetOrder() *string {
	return s.Order
}

func (s *ListCatalogsShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCatalogsShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCatalogsShrinkRequest) GetParentMetaEntityId() *string {
	return s.ParentMetaEntityId
}

func (s *ListCatalogsShrinkRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListCatalogsShrinkRequest) GetTypesShrink() *string {
	return s.TypesShrink
}

func (s *ListCatalogsShrinkRequest) SetComment(v string) *ListCatalogsShrinkRequest {
	s.Comment = &v
	return s
}

func (s *ListCatalogsShrinkRequest) SetName(v string) *ListCatalogsShrinkRequest {
	s.Name = &v
	return s
}

func (s *ListCatalogsShrinkRequest) SetOrder(v string) *ListCatalogsShrinkRequest {
	s.Order = &v
	return s
}

func (s *ListCatalogsShrinkRequest) SetPageNumber(v int32) *ListCatalogsShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListCatalogsShrinkRequest) SetPageSize(v int32) *ListCatalogsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListCatalogsShrinkRequest) SetParentMetaEntityId(v string) *ListCatalogsShrinkRequest {
	s.ParentMetaEntityId = &v
	return s
}

func (s *ListCatalogsShrinkRequest) SetSortBy(v string) *ListCatalogsShrinkRequest {
	s.SortBy = &v
	return s
}

func (s *ListCatalogsShrinkRequest) SetTypesShrink(v string) *ListCatalogsShrinkRequest {
	s.TypesShrink = &v
	return s
}

func (s *ListCatalogsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
