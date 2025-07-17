// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSensitiveColumnsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetColumnName(v string) *ListSensitiveColumnsRequest
	GetColumnName() *string
	SetDbId(v int64) *ListSensitiveColumnsRequest
	GetDbId() *int64
	SetLogic(v bool) *ListSensitiveColumnsRequest
	GetLogic() *bool
	SetPageNumber(v int32) *ListSensitiveColumnsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListSensitiveColumnsRequest
	GetPageSize() *int32
	SetSchemaName(v string) *ListSensitiveColumnsRequest
	GetSchemaName() *string
	SetSecurityLevel(v string) *ListSensitiveColumnsRequest
	GetSecurityLevel() *string
	SetTableName(v string) *ListSensitiveColumnsRequest
	GetTableName() *string
	SetTid(v int64) *ListSensitiveColumnsRequest
	GetTid() *int64
}

type ListSensitiveColumnsRequest struct {
	// The name of the field. You can call the [ListColumns](https://help.aliyun.com/document_detail/141870.html) operation to query the name of the field.
	//
	// example:
	//
	// test_column
	ColumnName *string `json:"ColumnName,omitempty" xml:"ColumnName,omitempty"`
	// The ID of the database. You can call the [SearchDatabase](https://help.aliyun.com/document_detail/141876.html) operation to query the ID of the database.
	//
	// >  You can also call the [ListDatabases](https://help.aliyun.com/document_detail/141873.html) operation to query the ID of the physical database and the [ListLogicDatabases](https://help.aliyun.com/document_detail/141874.html) operation to query the ID of a logical database.
	//
	// example:
	//
	// 1860
	DbId *int64 `json:"DbId,omitempty" xml:"DbId,omitempty"`
	// Specifies whether the database is a logical database. Valid values:
	//
	// 	- true: The database is a logical database.
	//
	// 	- false: The database is a physical database.
	//
	// example:
	//
	// false
	Logic *bool `json:"Logic,omitempty" xml:"Logic,omitempty"`
	// The number of the page to return.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The name of the database. You can call the [SearchDatabase](https://help.aliyun.com/document_detail/141876.html) operation to query the name of the database.
	//
	// >  You can also call the [ListDatabases](https://help.aliyun.com/document_detail/141873.html) operation to query the name of a physical database and the [ListLogicDatabases](https://help.aliyun.com/document_detail/141874.html) operation to query the name of a logical database.
	//
	// example:
	//
	// test_schema
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	// The sensitivity level of the field. Valid values:
	//
	// 	- SENSITIVE: medium sensitivity level
	//
	// 	- CONFIDENTIAL: high sensitivity level
	//
	// example:
	//
	// SENSITIVE
	SecurityLevel *string `json:"SecurityLevel,omitempty" xml:"SecurityLevel,omitempty"`
	// The name of the table. You can call the [ListTables](https://help.aliyun.com/document_detail/141878.html) operation to query the ID of the table.
	//
	// example:
	//
	// test_table
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// The ID of the tenant.
	//
	// >  To view the ID of the tenant, move the pointer over the profile picture in the upper-right corner of the Data Management (DMS) console. For more information, see the "View information about the current tenant" section of the [Manage DMS tenants](https://help.aliyun.com/document_detail/181330.html) topic.
	//
	// example:
	//
	// 3
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s ListSensitiveColumnsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSensitiveColumnsRequest) GoString() string {
	return s.String()
}

func (s *ListSensitiveColumnsRequest) GetColumnName() *string {
	return s.ColumnName
}

func (s *ListSensitiveColumnsRequest) GetDbId() *int64 {
	return s.DbId
}

func (s *ListSensitiveColumnsRequest) GetLogic() *bool {
	return s.Logic
}

func (s *ListSensitiveColumnsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListSensitiveColumnsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSensitiveColumnsRequest) GetSchemaName() *string {
	return s.SchemaName
}

func (s *ListSensitiveColumnsRequest) GetSecurityLevel() *string {
	return s.SecurityLevel
}

func (s *ListSensitiveColumnsRequest) GetTableName() *string {
	return s.TableName
}

func (s *ListSensitiveColumnsRequest) GetTid() *int64 {
	return s.Tid
}

func (s *ListSensitiveColumnsRequest) SetColumnName(v string) *ListSensitiveColumnsRequest {
	s.ColumnName = &v
	return s
}

func (s *ListSensitiveColumnsRequest) SetDbId(v int64) *ListSensitiveColumnsRequest {
	s.DbId = &v
	return s
}

func (s *ListSensitiveColumnsRequest) SetLogic(v bool) *ListSensitiveColumnsRequest {
	s.Logic = &v
	return s
}

func (s *ListSensitiveColumnsRequest) SetPageNumber(v int32) *ListSensitiveColumnsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListSensitiveColumnsRequest) SetPageSize(v int32) *ListSensitiveColumnsRequest {
	s.PageSize = &v
	return s
}

func (s *ListSensitiveColumnsRequest) SetSchemaName(v string) *ListSensitiveColumnsRequest {
	s.SchemaName = &v
	return s
}

func (s *ListSensitiveColumnsRequest) SetSecurityLevel(v string) *ListSensitiveColumnsRequest {
	s.SecurityLevel = &v
	return s
}

func (s *ListSensitiveColumnsRequest) SetTableName(v string) *ListSensitiveColumnsRequest {
	s.TableName = &v
	return s
}

func (s *ListSensitiveColumnsRequest) SetTid(v int64) *ListSensitiveColumnsRequest {
	s.Tid = &v
	return s
}

func (s *ListSensitiveColumnsRequest) Validate() error {
	return dara.Validate(s)
}
