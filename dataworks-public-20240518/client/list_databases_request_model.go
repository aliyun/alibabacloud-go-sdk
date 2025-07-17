// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDatabasesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *ListDatabasesRequest
	GetComment() *string
	SetName(v string) *ListDatabasesRequest
	GetName() *string
	SetOrder(v string) *ListDatabasesRequest
	GetOrder() *string
	SetPageNumber(v int32) *ListDatabasesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListDatabasesRequest
	GetPageSize() *int32
	SetParentMetaEntityId(v string) *ListDatabasesRequest
	GetParentMetaEntityId() *string
	SetSortBy(v string) *ListDatabasesRequest
	GetSortBy() *string
}

type ListDatabasesRequest struct {
	// example:
	//
	// test comment
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// example:
	//
	// test_tbl
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
	// The parent entity ID. For more information, see [description of concepts related to metadata entities.](https://help.aliyun.com/document_detail/2880092.html)
	//
	// The type of the parent entity can be found in the response of the ListCrawlerTypes operation.
	//
	// 	- If the parent entity is a catalog, the format of `ParentMetaEntityId` follows the response of the ListCatalogs API.
	//
	// 	- If the parent entity is a metadata crawler, the format of `ParentMetaEntityId` is `${CrawlerType}:${Instance ID or encoded URL}.`
	//
	// ParentMetaEntityId format examples
	//
	// `dlf-catalog::catalog_id`
	//
	// `holo:instance_id`
	//
	// `mysql:(instance_id|encoded_jdbc_url)`
	//
	// > \\
	//
	// `catalog_id`: The DLF catalog ID.\\
	//
	// `instance_id`: The instance ID, required for the data source registered in instance mode.\\
	//
	// `encoded_jdbc_url`: The JDBC connection string that has been URL encoded, required for the data source registered via a connection string.
	//
	// This parameter is required.
	//
	// example:
	//
	// mysql:rm-abc123xxx
	//
	// dlf-catalog:123456XXX:test_catalog
	ParentMetaEntityId *string `json:"ParentMetaEntityId,omitempty" xml:"ParentMetaEntityId,omitempty"`
	// example:
	//
	// CreateTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
}

func (s ListDatabasesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDatabasesRequest) GoString() string {
	return s.String()
}

func (s *ListDatabasesRequest) GetComment() *string {
	return s.Comment
}

func (s *ListDatabasesRequest) GetName() *string {
	return s.Name
}

func (s *ListDatabasesRequest) GetOrder() *string {
	return s.Order
}

func (s *ListDatabasesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDatabasesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDatabasesRequest) GetParentMetaEntityId() *string {
	return s.ParentMetaEntityId
}

func (s *ListDatabasesRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListDatabasesRequest) SetComment(v string) *ListDatabasesRequest {
	s.Comment = &v
	return s
}

func (s *ListDatabasesRequest) SetName(v string) *ListDatabasesRequest {
	s.Name = &v
	return s
}

func (s *ListDatabasesRequest) SetOrder(v string) *ListDatabasesRequest {
	s.Order = &v
	return s
}

func (s *ListDatabasesRequest) SetPageNumber(v int32) *ListDatabasesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDatabasesRequest) SetPageSize(v int32) *ListDatabasesRequest {
	s.PageSize = &v
	return s
}

func (s *ListDatabasesRequest) SetParentMetaEntityId(v string) *ListDatabasesRequest {
	s.ParentMetaEntityId = &v
	return s
}

func (s *ListDatabasesRequest) SetSortBy(v string) *ListDatabasesRequest {
	s.SortBy = &v
	return s
}

func (s *ListDatabasesRequest) Validate() error {
	return dara.Validate(s)
}
