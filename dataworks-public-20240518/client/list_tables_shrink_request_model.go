// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTablesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *ListTablesShrinkRequest
	GetComment() *string
	SetName(v string) *ListTablesShrinkRequest
	GetName() *string
	SetOrder(v string) *ListTablesShrinkRequest
	GetOrder() *string
	SetPageNumber(v int32) *ListTablesShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListTablesShrinkRequest
	GetPageSize() *int32
	SetParentMetaEntityId(v string) *ListTablesShrinkRequest
	GetParentMetaEntityId() *string
	SetSortBy(v string) *ListTablesShrinkRequest
	GetSortBy() *string
	SetTableTypesShrink(v string) *ListTablesShrinkRequest
	GetTableTypesShrink() *string
}

type ListTablesShrinkRequest struct {
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
	// The parent metadata entity ID. You can refer to the responses of the ListDatabases or ListSchemas operation and [Description of concepts related to metadata entities.](https://help.aliyun.com/document_detail/2880092.html)
	//
	// 	- The parent metadata entity is a database: The format of `ParentMetaEntityId` is `${EntityType}:${Instance ID or encoded URL}:${Catalog Identifier}:${Database Name}`. Use an empty string (\\`""\\`) as a placeholder for any non-existent level.
	//
	// 	- The parent metadata entity is a database schema:. The format of `ParentMetaEntityId` is `${EntityType}:${Instance ID or encoded URL}:${Catalog Identifier}:${Database Name}:${Schema Name}`. Use an empty string (\\`""\\`) as a placeholder for any non-existent level.
	//
	// >  The schema level in `ParentMetaEntityId` is supported only for database types that support schemas, such as MaxCompute (with schema enabled), Hologres, PostgreSQL, SQL Server, HybridDB for PostgreSQL, and Oracle.``
	//
	// >  For MaxCompute and DLF types, use empty strings as the instance ID. For MaxCompute, the database name is the same as the project name.
	//
	// >  For the StarRocks type, the catalog identifier is the catalog name. For the DLF type, it refers to the catalog ID. Other types do not support a catalog-level hierarchy and the catalog identifier must be replaced with an empty string as a placeholder.
	//
	// Examples of common ParentMetaEntityId formats
	//
	// `maxcompute-project:::project_name`
	//
	// `maxcompute-schema:::project_name:schema_name` (for MaxCompute projects with schema enabled)
	//
	// `dlf-database::catalog_id:database_name`
	//
	// `hms-database:instance_id::database_name`
	//
	// `holo-schema:instance_id::database_name:schema_name`
	//
	// `mysql-database:(instance_id|encoded_jdbc_url)::database_name`
	//
	// > \\
	//
	// `instance_id`: The instance ID, required when the data source is registered in instance mode.\\
	//
	// `encoded_jdbc_url`: The JDBC connection string that has been URL encoded, required for the data source registered via a connection string.\\
	//
	// `catalog_id`: The DLF catalog ID.\\
	//
	// `project_name`: The MaxCompute project name.\\
	//
	// `database_name`: The database name.\\
	//
	// `schema_name`: The schema name.
	//
	// This parameter is required.
	//
	// example:
	//
	// maxcompute-schema:123456XXX::test_project_with_schema:default
	//
	// maxcompute-project:123456XXX::test_project_without_schema
	//
	// dlf-database:123456XXX:test_catalog:test_db
	//
	// hms-database:c-abc123xxx::test_db
	//
	// holo-schema:h-abc123xxx::test_db:test_schema
	//
	// mysql-database:jdbc%3Amysql%3A%2F%2F127.0.0.1%3A3306%2Ftest_db::test_db
	ParentMetaEntityId *string `json:"ParentMetaEntityId,omitempty" xml:"ParentMetaEntityId,omitempty"`
	// example:
	//
	// CreateTime
	SortBy           *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	TableTypesShrink *string `json:"TableTypes,omitempty" xml:"TableTypes,omitempty"`
}

func (s ListTablesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTablesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListTablesShrinkRequest) GetComment() *string {
	return s.Comment
}

func (s *ListTablesShrinkRequest) GetName() *string {
	return s.Name
}

func (s *ListTablesShrinkRequest) GetOrder() *string {
	return s.Order
}

func (s *ListTablesShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListTablesShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListTablesShrinkRequest) GetParentMetaEntityId() *string {
	return s.ParentMetaEntityId
}

func (s *ListTablesShrinkRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListTablesShrinkRequest) GetTableTypesShrink() *string {
	return s.TableTypesShrink
}

func (s *ListTablesShrinkRequest) SetComment(v string) *ListTablesShrinkRequest {
	s.Comment = &v
	return s
}

func (s *ListTablesShrinkRequest) SetName(v string) *ListTablesShrinkRequest {
	s.Name = &v
	return s
}

func (s *ListTablesShrinkRequest) SetOrder(v string) *ListTablesShrinkRequest {
	s.Order = &v
	return s
}

func (s *ListTablesShrinkRequest) SetPageNumber(v int32) *ListTablesShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTablesShrinkRequest) SetPageSize(v int32) *ListTablesShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListTablesShrinkRequest) SetParentMetaEntityId(v string) *ListTablesShrinkRequest {
	s.ParentMetaEntityId = &v
	return s
}

func (s *ListTablesShrinkRequest) SetSortBy(v string) *ListTablesShrinkRequest {
	s.SortBy = &v
	return s
}

func (s *ListTablesShrinkRequest) SetTableTypesShrink(v string) *ListTablesShrinkRequest {
	s.TableTypesShrink = &v
	return s
}

func (s *ListTablesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
