// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSensitiveColumnInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetColumnName(v string) *ListSensitiveColumnInfoRequest
	GetColumnName() *string
	SetInstanceId(v int32) *ListSensitiveColumnInfoRequest
	GetInstanceId() *int32
	SetPageNumber(v int32) *ListSensitiveColumnInfoRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListSensitiveColumnInfoRequest
	GetPageSize() *int32
	SetSchemaName(v string) *ListSensitiveColumnInfoRequest
	GetSchemaName() *string
	SetTableName(v string) *ListSensitiveColumnInfoRequest
	GetTableName() *string
	SetTid(v int64) *ListSensitiveColumnInfoRequest
	GetTid() *int64
}

type ListSensitiveColumnInfoRequest struct {
	// The name of the sensitive field. You can call the [ListSensitiveColumns](https://help.aliyun.com/document_detail/188103.html) operation to query the name of the sensitive field.
	//
	// > You can also call the [ListColumns](https://help.aliyun.com/document_detail/141870.html) operation to query the name of the sensitive field.
	//
	// example:
	//
	// test_column
	ColumnName *string `json:"ColumnName,omitempty" xml:"ColumnName,omitempty"`
	// The ID of the instance. You can call the [ListInstances](https://help.aliyun.com/document_detail/141936.html) or [GetInstance](https://help.aliyun.com/document_detail/141567.html) operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 183****
	InstanceId *int32 `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The page number of the returned page.
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
	// The database name. You can also call the [SearchDatabase](https://help.aliyun.com/document_detail/141876.html) operation to query the name of the database.
	//
	// > You can also call the [ListDatabases](https://help.aliyun.com/document_detail/141873.html) operation to query the value of the SchemaName parameter of a physical database, or the [ListLogicDatabases](https://help.aliyun.com/document_detail/141874.html) operation to query the value of the SchemaName parameter of a logical database.
	//
	// example:
	//
	// test_schema
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	// The name of the table. You can call the [ListSensitiveColumns](https://help.aliyun.com/document_detail/188103.html) operation to query the table name.
	//
	// > You can also call the [ListTables](https://help.aliyun.com/document_detail/141878.html) operation to query the table name.
	//
	// example:
	//
	// test_table
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// The ID of the tenant.
	//
	// > To view the ID of the tenant, go to the DMS console and move the pointer over the profile picture in the upper-right corner. For more information, see the [View information about the current tenant](https://help.aliyun.com/document_detail/181330.html) section of the "Manage DMS tenants" topic.
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s ListSensitiveColumnInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSensitiveColumnInfoRequest) GoString() string {
	return s.String()
}

func (s *ListSensitiveColumnInfoRequest) GetColumnName() *string {
	return s.ColumnName
}

func (s *ListSensitiveColumnInfoRequest) GetInstanceId() *int32 {
	return s.InstanceId
}

func (s *ListSensitiveColumnInfoRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListSensitiveColumnInfoRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSensitiveColumnInfoRequest) GetSchemaName() *string {
	return s.SchemaName
}

func (s *ListSensitiveColumnInfoRequest) GetTableName() *string {
	return s.TableName
}

func (s *ListSensitiveColumnInfoRequest) GetTid() *int64 {
	return s.Tid
}

func (s *ListSensitiveColumnInfoRequest) SetColumnName(v string) *ListSensitiveColumnInfoRequest {
	s.ColumnName = &v
	return s
}

func (s *ListSensitiveColumnInfoRequest) SetInstanceId(v int32) *ListSensitiveColumnInfoRequest {
	s.InstanceId = &v
	return s
}

func (s *ListSensitiveColumnInfoRequest) SetPageNumber(v int32) *ListSensitiveColumnInfoRequest {
	s.PageNumber = &v
	return s
}

func (s *ListSensitiveColumnInfoRequest) SetPageSize(v int32) *ListSensitiveColumnInfoRequest {
	s.PageSize = &v
	return s
}

func (s *ListSensitiveColumnInfoRequest) SetSchemaName(v string) *ListSensitiveColumnInfoRequest {
	s.SchemaName = &v
	return s
}

func (s *ListSensitiveColumnInfoRequest) SetTableName(v string) *ListSensitiveColumnInfoRequest {
	s.TableName = &v
	return s
}

func (s *ListSensitiveColumnInfoRequest) SetTid(v int64) *ListSensitiveColumnInfoRequest {
	s.Tid = &v
	return s
}

func (s *ListSensitiveColumnInfoRequest) Validate() error {
	return dara.Validate(s)
}
