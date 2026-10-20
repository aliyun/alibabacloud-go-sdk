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
	SetIncludeExtendedProperties(v bool) *ListTablesRequest
	GetIncludeExtendedProperties() *bool
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
	// The comment. Fuzzy match is supported.
	//
	// example:
	//
	// this is a comment
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// Specifies whether to return extended properties. If this parameter is set to `true`, extended properties are returned. If this parameter is set to `false`, extended properties are not returned.
	//
	// example:
	//
	// true
	IncludeExtendedProperties *bool `json:"IncludeExtendedProperties,omitempty" xml:"IncludeExtendedProperties,omitempty"`
	// The name. Fuzzy match is supported.
	//
	// example:
	//
	// abc
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The sort order. Default value: Asc. Valid values:
	//
	// - Asc: ascending order
	//
	// - Desc: descending order
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
	// The number of entries per page. Default value: 10. Maximum value: 100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the parent metadata entity. You can refer to the response of the ListDatabases or ListSchemas operation, or see [Metadata entity concepts](https://help.aliyun.com/document_detail/2880092.html).
	//
	// - The parent entity can be the database to which the table belongs. The `ParentMetaEntityId` format is `${EntityType}:${Instance ID or URL-encoded JDBC URL}:${Data catalog identifier}:${Database name}`. Use an empty string as a placeholder for non-existent levels.
	//
	// - The parent entity can also be the database schema to which the table belongs. The `ParentMetaEntityId` format is `${EntityType}:${Instance ID or URL-encoded JDBC URL}:${Data catalog identifier}:${Database name}:${Schema name}`. Use an empty string as a placeholder for non-existent levels.
	//
	// > - The database schema can be used for `ParentMetaEntityId` only when the database type supports schemas (`maxcompute`, `holo`, `postgresql`, `sqlserver`, `hybriddb_for_postgresql`, or `oracle`). For MaxCompute, the three-level model must be enabled. Otherwise, only the database can be used.
	//
	// > - For MaxCompute and DLF types, use an empty string as a placeholder for the instance ID. For MaxCompute, the database name is the MaxCompute project name.
	//
	// > - For StarRocks, the data catalog identifier is the catalog name. For DLF, the data catalog identifier is the catalog ID. Other types do not support the catalog level, so use an empty string as a placeholder.
	//
	// Examples of `ParentMetaEntityId` formats for common types:
	//
	// - `maxcompute-project:::project_name`
	//
	// - `maxcompute-schema:::project_name:schema_name` (only when the three-level model is enabled for the project)
	//
	// - `dlf-database::catalog_id:database_name`
	//
	// - `hms-database:instance_id::database_name`
	//
	// - `holo-schema:instance_id::database_name:schema_name`
	//
	// - `mysql-database:(instance_id|encoded_jdbc_url)::database_name`
	//
	// > Where:
	//
	// > - `instance_id`: the instance ID, which is required when the data source is registered in instance mode.
	//
	// > - `encoded_jdbc_url`: the URL-encoded JDBC connection string, which is required when the data source is registered by using a connection string.
	//
	// > - `catalog_id`: the DLF catalog ID.
	//
	// > - `project_name`: the MaxCompute project name.
	//
	// > - `database_name`: the database name.
	//
	// > - `schema_name`: the schema name.
	//
	// This parameter is required.
	//
	// example:
	//
	// maxcompute-project:::project_name
	ParentMetaEntityId *string `json:"ParentMetaEntityId,omitempty" xml:"ParentMetaEntityId,omitempty"`
	// The field used for sorting. Default value: CreateTime. Valid values:
	//
	// - CreateTime: creation time
	//
	// - ModifyTime: modification time
	//
	// - Name: name
	//
	// - TableType: table type
	//
	// example:
	//
	// CreateTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The list of table types to query. If this parameter is left empty, all table types are queried.
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

func (s *ListTablesRequest) GetIncludeExtendedProperties() *bool {
	return s.IncludeExtendedProperties
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

func (s *ListTablesRequest) SetIncludeExtendedProperties(v bool) *ListTablesRequest {
	s.IncludeExtendedProperties = &v
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
