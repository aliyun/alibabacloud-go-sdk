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
	// example:
	//
	// this is a comment
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// example:
	//
	// abc
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// Asc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The parent entity ID. For more information, see [Concepts related to metadata entities](https://help.aliyun.com/document_detail/2880092.html).
	//
	// Only DLF and StarRocks data sources support this parameter.
	//
	// 	- For DLF data sources, you can call this API operation to query all catalogs. In this case, you must set the `ParentMetaEntityId` parameter to `dlf`.
	//
	// 	- For StarRocks data sources, you can call this API operation to query the catalogs in a specific instance. In this case, you can configure the `ParentMetaEntityId` parameter in the `starrocks:(instance_id|encoded_jdbc_url)` format.
	//
	// > \\
	//
	// `instance_id`: the ID of an instance. If the related data source is added to DataWorks in Alibaba Cloud instance mode, you must configure this parameter.\\
	//
	// `encoded_jdbc_url`: the JDBC connection string that is URL-encoded. If the related data source is added to DataWorks in connection string mode, you must configure this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// dlf
	//
	// starrocks:c-abc123xxx
	ParentMetaEntityId *string `json:"ParentMetaEntityId,omitempty" xml:"ParentMetaEntityId,omitempty"`
	// example:
	//
	// CreateTime
	SortBy      *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
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
