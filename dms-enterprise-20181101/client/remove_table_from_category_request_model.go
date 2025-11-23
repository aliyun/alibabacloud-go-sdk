// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveTableFromCategoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategoryId(v int64) *RemoveTableFromCategoryRequest
	GetCategoryId() *int64
	SetDbId(v int64) *RemoveTableFromCategoryRequest
	GetDbId() *int64
	SetTableName(v string) *RemoveTableFromCategoryRequest
	GetTableName() *string
	SetTableSchemaName(v string) *RemoveTableFromCategoryRequest
	GetTableSchemaName() *string
	SetTid(v int64) *RemoveTableFromCategoryRequest
	GetTid() *int64
}

type RemoveTableFromCategoryRequest struct {
	// The category ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30000235594
	CategoryId *int64 `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// The database ID. You can call the [ListDatabases](https://help.aliyun.com/document_detail/141873.html) operation to query the ID of a physical database and the [ListLogicDatabases](https://help.aliyun.com/document_detail/141874.html) operation to query the ID of a logical database.
	//
	// >  The value of DatabaseId is that of DbId.
	//
	// This parameter is required.
	//
	// example:
	//
	// 43153
	DbId *int64 `json:"DbId,omitempty" xml:"DbId,omitempty"`
	// The table name.
	//
	// > You can also call the [ListTables](https://help.aliyun.com/document_detail/141878.html) operation to query the table name.
	//
	// This parameter is required.
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
	// The ID of the tenant.
	//
	// >  To view the ID of the tenant, go to the Data Management (DMS) console and move the pointer over the profile picture in the upper-right corner. For more information, see the [View information about the current tenant](https://help.aliyun.com/document_detail/181330.html) section of the "Manage DMS tenants" topic.
	//
	// example:
	//
	// 3****
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s RemoveTableFromCategoryRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveTableFromCategoryRequest) GoString() string {
	return s.String()
}

func (s *RemoveTableFromCategoryRequest) GetCategoryId() *int64 {
	return s.CategoryId
}

func (s *RemoveTableFromCategoryRequest) GetDbId() *int64 {
	return s.DbId
}

func (s *RemoveTableFromCategoryRequest) GetTableName() *string {
	return s.TableName
}

func (s *RemoveTableFromCategoryRequest) GetTableSchemaName() *string {
	return s.TableSchemaName
}

func (s *RemoveTableFromCategoryRequest) GetTid() *int64 {
	return s.Tid
}

func (s *RemoveTableFromCategoryRequest) SetCategoryId(v int64) *RemoveTableFromCategoryRequest {
	s.CategoryId = &v
	return s
}

func (s *RemoveTableFromCategoryRequest) SetDbId(v int64) *RemoveTableFromCategoryRequest {
	s.DbId = &v
	return s
}

func (s *RemoveTableFromCategoryRequest) SetTableName(v string) *RemoveTableFromCategoryRequest {
	s.TableName = &v
	return s
}

func (s *RemoveTableFromCategoryRequest) SetTableSchemaName(v string) *RemoveTableFromCategoryRequest {
	s.TableSchemaName = &v
	return s
}

func (s *RemoveTableFromCategoryRequest) SetTid(v int64) *RemoveTableFromCategoryRequest {
	s.Tid = &v
	return s
}

func (s *RemoveTableFromCategoryRequest) Validate() error {
	return dara.Validate(s)
}
