// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMetaCategoryTableEntity interface {
	dara.Model
	String() string
	GoString() string
	SetCatalogName(v string) *MetaCategoryTableEntity
	GetCatalogName() *string
	SetCategoryId(v int64) *MetaCategoryTableEntity
	GetCategoryId() *int64
	SetDatabaseSearchName(v string) *MetaCategoryTableEntity
	GetDatabaseSearchName() *string
	SetDbId(v int32) *MetaCategoryTableEntity
	GetDbId() *int32
	SetDbType(v string) *MetaCategoryTableEntity
	GetDbType() *string
	SetDescription(v string) *MetaCategoryTableEntity
	GetDescription() *string
	SetInstanceId(v int32) *MetaCategoryTableEntity
	GetInstanceId() *int32
	SetSchemaName(v string) *MetaCategoryTableEntity
	GetSchemaName() *string
	SetTableName(v string) *MetaCategoryTableEntity
	GetTableName() *string
	SetTableSchemaName(v string) *MetaCategoryTableEntity
	GetTableSchemaName() *string
}

type MetaCategoryTableEntity struct {
	// For PostgreSQL-compatible databases, specify the database name.
	//
	// example:
	//
	// hive
	CatalogName *string `json:"CatalogName,omitempty" xml:"CatalogName,omitempty"`
	// The category ID.
	//
	// example:
	//
	// FC-1D123DF554A45AAB
	CategoryId *int64 `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// The name that is used to search for the database.
	//
	// example:
	//
	// xxx@yyy.zzz
	DatabaseSearchName *string `json:"DatabaseSearchName,omitempty" xml:"DatabaseSearchName,omitempty"`
	// The database ID. You can call the [ListDatabases](https://help.aliyun.com/document_detail/141873.html) operation to query the ID of a physical database and the [ListLogicDatabases](https://help.aliyun.com/document_detail/141874.html) operation to query the ID of a logical database.
	//
	// > The value of DatabaseId is that of DbId.
	//
	// example:
	//
	// 123***
	DbId *int32 `json:"DbId,omitempty" xml:"DbId,omitempty"`
	// The type of the database. Valid values include but are not limited to:
	//
	// 	- **MySQL**
	//
	// 	- **SQLServer**
	//
	// 	- **PostgreSQL**
	//
	// 	- **Oracle**
	//
	// 	- **DRDS**
	//
	// 	- **OceanBase**
	//
	// 	- **Mongo**
	//
	// 	- **Redis**
	//
	// example:
	//
	// MySQL
	DbType *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
	// The description.
	//
	// example:
	//
	// ga_platform_alb
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the instance. You can call the [ListInstances](https://help.aliyun.com/document_detail/141936.html) or [GetInstance](https://help.aliyun.com/document_detail/141567.html) operation to query the instance ID.
	//
	// example:
	//
	// 174****
	InstanceId *int32 `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Database name (for PostgreSQL-compatible databases, specify the schema name). You can call the [SearchDatabase](https://help.aliyun.com/document_detail/141876.html) operation to query the name of the database.
	//
	// > You can also call the [ListDatabases](https://help.aliyun.com/document_detail/141873.html) operation to query the SchemaName of a physical database or call the [ListLogicDatabases](https://help.aliyun.com/document_detail/141874.html) operation to query the SchemaName of a logical database.
	//
	// example:
	//
	// test_schema
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	// The name of the table.
	//
	// example:
	//
	// test_table
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// The schema name of the table, which is required only for SQL Server instances.
	//
	// example:
	//
	// dbo
	TableSchemaName *string `json:"TableSchemaName,omitempty" xml:"TableSchemaName,omitempty"`
}

func (s MetaCategoryTableEntity) String() string {
	return dara.Prettify(s)
}

func (s MetaCategoryTableEntity) GoString() string {
	return s.String()
}

func (s *MetaCategoryTableEntity) GetCatalogName() *string {
	return s.CatalogName
}

func (s *MetaCategoryTableEntity) GetCategoryId() *int64 {
	return s.CategoryId
}

func (s *MetaCategoryTableEntity) GetDatabaseSearchName() *string {
	return s.DatabaseSearchName
}

func (s *MetaCategoryTableEntity) GetDbId() *int32 {
	return s.DbId
}

func (s *MetaCategoryTableEntity) GetDbType() *string {
	return s.DbType
}

func (s *MetaCategoryTableEntity) GetDescription() *string {
	return s.Description
}

func (s *MetaCategoryTableEntity) GetInstanceId() *int32 {
	return s.InstanceId
}

func (s *MetaCategoryTableEntity) GetSchemaName() *string {
	return s.SchemaName
}

func (s *MetaCategoryTableEntity) GetTableName() *string {
	return s.TableName
}

func (s *MetaCategoryTableEntity) GetTableSchemaName() *string {
	return s.TableSchemaName
}

func (s *MetaCategoryTableEntity) SetCatalogName(v string) *MetaCategoryTableEntity {
	s.CatalogName = &v
	return s
}

func (s *MetaCategoryTableEntity) SetCategoryId(v int64) *MetaCategoryTableEntity {
	s.CategoryId = &v
	return s
}

func (s *MetaCategoryTableEntity) SetDatabaseSearchName(v string) *MetaCategoryTableEntity {
	s.DatabaseSearchName = &v
	return s
}

func (s *MetaCategoryTableEntity) SetDbId(v int32) *MetaCategoryTableEntity {
	s.DbId = &v
	return s
}

func (s *MetaCategoryTableEntity) SetDbType(v string) *MetaCategoryTableEntity {
	s.DbType = &v
	return s
}

func (s *MetaCategoryTableEntity) SetDescription(v string) *MetaCategoryTableEntity {
	s.Description = &v
	return s
}

func (s *MetaCategoryTableEntity) SetInstanceId(v int32) *MetaCategoryTableEntity {
	s.InstanceId = &v
	return s
}

func (s *MetaCategoryTableEntity) SetSchemaName(v string) *MetaCategoryTableEntity {
	s.SchemaName = &v
	return s
}

func (s *MetaCategoryTableEntity) SetTableName(v string) *MetaCategoryTableEntity {
	s.TableName = &v
	return s
}

func (s *MetaCategoryTableEntity) SetTableSchemaName(v string) *MetaCategoryTableEntity {
	s.TableSchemaName = &v
	return s
}

func (s *MetaCategoryTableEntity) Validate() error {
	return dara.Validate(s)
}
