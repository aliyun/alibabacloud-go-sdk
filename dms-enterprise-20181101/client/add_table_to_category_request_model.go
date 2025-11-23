// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddTableToCategoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategoryId(v int64) *AddTableToCategoryRequest
	GetCategoryId() *int64
	SetDbId(v int64) *AddTableToCategoryRequest
	GetDbId() *int64
	SetTableName(v string) *AddTableToCategoryRequest
	GetTableName() *string
	SetTableSchemaName(v string) *AddTableToCategoryRequest
	GetTableSchemaName() *string
	SetTid(v int64) *AddTableToCategoryRequest
	GetTid() *int64
}

type AddTableToCategoryRequest struct {
	// The ID of the associated category.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30000254257
	CategoryId *int64 `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// The ID of a physical database: You can call the [ListDatabases](https://help.aliyun.com/document_detail/141873.html) or [SearchDatabase](https://help.aliyun.com/document_detail/141876.html) operation to obtain the physical database ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1930****
	DbId *int64 `json:"DbId,omitempty" xml:"DbId,omitempty"`
	// The name of the table. You can call the [ListTables](https://help.aliyun.com/document_detail/141878.html) operation to query the table name.
	//
	// This parameter is required.
	//
	// example:
	//
	// table_name
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// The schema name of the table, which is required only for SQL Server instances.
	//
	// example:
	//
	// dbo
	TableSchemaName *string `json:"TableSchemaName,omitempty" xml:"TableSchemaName,omitempty"`
	// The tenant ID. You can call the [GetUserActiveTenant](https://help.aliyun.com/document_detail/198073.html) or [ListUserTenants](https://help.aliyun.com/document_detail/198074.html) operation to obtain the tenant ID.
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s AddTableToCategoryRequest) String() string {
	return dara.Prettify(s)
}

func (s AddTableToCategoryRequest) GoString() string {
	return s.String()
}

func (s *AddTableToCategoryRequest) GetCategoryId() *int64 {
	return s.CategoryId
}

func (s *AddTableToCategoryRequest) GetDbId() *int64 {
	return s.DbId
}

func (s *AddTableToCategoryRequest) GetTableName() *string {
	return s.TableName
}

func (s *AddTableToCategoryRequest) GetTableSchemaName() *string {
	return s.TableSchemaName
}

func (s *AddTableToCategoryRequest) GetTid() *int64 {
	return s.Tid
}

func (s *AddTableToCategoryRequest) SetCategoryId(v int64) *AddTableToCategoryRequest {
	s.CategoryId = &v
	return s
}

func (s *AddTableToCategoryRequest) SetDbId(v int64) *AddTableToCategoryRequest {
	s.DbId = &v
	return s
}

func (s *AddTableToCategoryRequest) SetTableName(v string) *AddTableToCategoryRequest {
	s.TableName = &v
	return s
}

func (s *AddTableToCategoryRequest) SetTableSchemaName(v string) *AddTableToCategoryRequest {
	s.TableSchemaName = &v
	return s
}

func (s *AddTableToCategoryRequest) SetTid(v int64) *AddTableToCategoryRequest {
	s.Tid = &v
	return s
}

func (s *AddTableToCategoryRequest) Validate() error {
	return dara.Validate(s)
}
