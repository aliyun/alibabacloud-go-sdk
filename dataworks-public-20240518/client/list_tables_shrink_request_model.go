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
	// The comment on the table. Fuzzy matching is supported.
	//
	// example:
	//
	// this is a comment
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The name of the table. Fuzzy matching is supported.
	//
	// example:
	//
	// abc
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The sort order. Default value: `Asc`. Valid values:
	//
	// - `Asc`: ascending
	//
	// - `Desc`: descending
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
	// The page size. Default value: 10. Maximum value: 100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the parent metadata entity. You can obtain this ID from the response of the ListDatabases or ListSchemas operation. For details, see [Metadata entity concepts](https://help.aliyun.com/document_detail/2880092.html).
	//
	// - The value can be the database to which the table belongs. The format is `${EntityType}:${instance ID or URL-encoded connection string}:${data catalog identifier}:${database name}`. Use an empty string as a placeholder for a hierarchy level that does not exist.
	//
	// - The value can also be the schema to which the table belongs. The format is `${EntityType}:${instance ID or URL-encoded connection string}:${data catalog identifier}:${database name}:${schema name}`. Use an empty string as a placeholder for a hierarchy level that does not exist.
	//
	// > 	- You can specify a schema in `ParentMetaEntityId` only if the database type supports schemas, such as `maxcompute/holo/postgresql/sqlserver/hybriddb_for_postgresql/oracle`. For the maxcompute type, the three-layer model must be enabled. Otherwise, you can only specify a database.
	//
	// >
	//
	// > 	- For `maxcompute` and `dlf` data types, use an empty string as a placeholder for the instance ID. For the maxcompute data type, the database name is the MaxCompute project name.
	//
	// >
	//
	// > 	- For the `starrocks` type, the data catalog identifier is the catalog name. For the `dlf` type, the data catalog identifier is the catalog ID. Other types do not support the catalog level, so you can use an empty string as a placeholder.
	//
	// The following list shows the `ParentMetaEntityId` format for several common data source types:
	//
	// - `maxcompute-project:::project_name`
	//
	// - `maxcompute-schema:::project_name:schema_name` (Only when the three-layer model is enabled for the project)
	//
	// - `dlf-database::catalog_id:database_name`
	//
	// - `hms-database:instance_id::database_name`
	//
	// - `holo-schema:instance_id::database_name:schema_name`
	//
	// - `mysql-database:(instance_id|encoded_jdbc_url)::database_name`
	//
	// > In these formats:
	//
	// >
	//
	// > - `instance_id`: The instance ID. This parameter is required if the data source is registered in instance mode.
	//
	// >
	//
	// > - `encoded_jdbc_url`: The URL-encoded JDBC connection string. This parameter is required if the data source is registered by using a connection string.
	//
	// >
	//
	// > - `catalog_id`: The ID of the DLF data catalog.
	//
	// >
	//
	// > - `project_name`: The name of the MaxCompute project.
	//
	// >
	//
	// > - `database_name`: The name of the database.
	//
	// >
	//
	// > - `schema_name`: The name of the schema.
	//
	// This parameter is required.
	//
	// example:
	//
	// maxcompute-project:::project_name
	ParentMetaEntityId *string `json:"ParentMetaEntityId,omitempty" xml:"ParentMetaEntityId,omitempty"`
	// The sort field. Default value: `CreateTime`. Valid values:
	//
	// - `CreateTime`: creation time
	//
	// - `ModifyTime`: modification time
	//
	// - `Name`: name
	//
	// - `TableType`: table type
	//
	// example:
	//
	// CreateTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// A list of table types to query. If you omit this parameter, tables of all types are returned.
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
