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
	// Specifies whether to return extended properties. Set this parameter to `true` to return extended properties or `false` to not return them.
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
	// The page size. Default value: 10. Maximum value: 100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the parent-level metadata entity. You can obtain this value from the response of the ListDatabases or ListSchemas operation. For more information, see [Metadata entity concepts](https://help.aliyun.com/document_detail/2880092.html).
	//
	// - The value can be the database to which the table belongs. The format of `ParentMetaEntityId` is `${EntityType}:${InstanceID or encoded URL}:${DataCatalogIdentifier}:${DatabaseName}`. Use an empty string as a placeholder for levels that do not exist.
	//
	// - The value can also be the database schema to which the table belongs. The format of `ParentMetaEntityId` is `${EntityType}:${InstanceID or encoded URL}:${DataCatalogIdentifier}:${DatabaseName}:${SchemaName}`. Use an empty string as a placeholder for levels that do not exist.
	//
	// > - You can set `ParentMetaEntityId` to a database schema only when the database type supports schemas (`maxcompute/holo/postgresql/sqlserver/hybriddb_for_postgresql/oracle`, where the three-layer model must be enabled for the maxcompute type). Otherwise, you can set it only to a database.
	//
	// > - For the maxcompute and dlf types, use an empty string as a placeholder for the instance ID. For the maxcompute type, the database name is the MaxCompute project name.
	//
	// > - For the starrocks type, the data catalog identifier is the catalog name. For the dlf type, the data catalog identifier is the catalog ID. Other types do not support the catalog level, and you can use an empty string as a placeholder.
	//
	// The following examples show the format of ParentMetaEntityId for common types:
	//
	// - `maxcompute-project:::project_name`
	//
	// - `maxcompute-schema:::project_name:schema_name` (only when the three-layer model is enabled for the project)
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
	// > - `instance_id`: The instance ID. This value is required when the data source is registered in instance mode.
	//
	// > - `encoded_jdbc_url`: The URL-encoded JDBC connection string. This value is required when the data source is registered by using a connection string.
	//
	// > - `catalog_id`: The DLF catalog ID.
	//
	// > - `project_name`: The MaxCompute project name.
	//
	// > - `database_name`: The database name.
	//
	// > - `schema_name`: The schema name.
	//
	// This parameter is required.
	//
	// example:
	//
	// maxcompute-project:::project_name
	ParentMetaEntityId *string `json:"ParentMetaEntityId,omitempty" xml:"ParentMetaEntityId,omitempty"`
	// The field by which to sort the results. Default value: CreateTime. Valid values:
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
	// The list of table types to query. If this parameter is left empty, all types are queried.
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
