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
	SetIncludeExtendedProperties(v bool) *ListTablesShrinkRequest
	GetIncludeExtendedProperties() *bool
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

func (s *ListTablesShrinkRequest) GetIncludeExtendedProperties() *bool {
	return s.IncludeExtendedProperties
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

func (s *ListTablesShrinkRequest) SetIncludeExtendedProperties(v bool) *ListTablesShrinkRequest {
	s.IncludeExtendedProperties = &v
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
