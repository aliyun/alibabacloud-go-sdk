// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTablesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *ListTablesRequest
	GetComment() *string
	SetName(v string) *ListTablesRequest
	GetName() *string
	SetOrder(v string) *ListTablesRequest
	GetOrder() *string
	SetPageNumber(v int32) *ListTablesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListTablesRequest
	GetPageSize() *int32
	SetParentMetaEntityId(v string) *ListTablesRequest
	GetParentMetaEntityId() *string
	SetSortBy(v string) *ListTablesRequest
	GetSortBy() *string
	SetTableTypes(v []*string) *ListTablesRequest
	GetTableTypes() []*string
}

type ListTablesRequest struct {
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
	SortBy     *string   `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	TableTypes []*string `json:"TableTypes,omitempty" xml:"TableTypes,omitempty" type:"Repeated"`
}

func (s ListTablesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTablesRequest) GoString() string {
	return s.String()
}

func (s *ListTablesRequest) GetComment() *string {
	return s.Comment
}

func (s *ListTablesRequest) GetName() *string {
	return s.Name
}

func (s *ListTablesRequest) GetOrder() *string {
	return s.Order
}

func (s *ListTablesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListTablesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListTablesRequest) GetParentMetaEntityId() *string {
	return s.ParentMetaEntityId
}

func (s *ListTablesRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListTablesRequest) GetTableTypes() []*string {
	return s.TableTypes
}

func (s *ListTablesRequest) SetComment(v string) *ListTablesRequest {
	s.Comment = &v
	return s
}

func (s *ListTablesRequest) SetName(v string) *ListTablesRequest {
	s.Name = &v
	return s
}

func (s *ListTablesRequest) SetOrder(v string) *ListTablesRequest {
	s.Order = &v
	return s
}

func (s *ListTablesRequest) SetPageNumber(v int32) *ListTablesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTablesRequest) SetPageSize(v int32) *ListTablesRequest {
	s.PageSize = &v
	return s
}

func (s *ListTablesRequest) SetParentMetaEntityId(v string) *ListTablesRequest {
	s.ParentMetaEntityId = &v
	return s
}

func (s *ListTablesRequest) SetSortBy(v string) *ListTablesRequest {
	s.SortBy = &v
	return s
}

func (s *ListTablesRequest) SetTableTypes(v []*string) *ListTablesRequest {
	s.TableTypes = v
	return s
}

func (s *ListTablesRequest) Validate() error {
	return dara.Validate(s)
}
