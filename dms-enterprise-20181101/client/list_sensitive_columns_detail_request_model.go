// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSensitiveColumnsDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetColumnName(v string) *ListSensitiveColumnsDetailRequest
	GetColumnName() *string
	SetDbId(v int64) *ListSensitiveColumnsDetailRequest
	GetDbId() *int64
	SetLogic(v bool) *ListSensitiveColumnsDetailRequest
	GetLogic() *bool
	SetSchemaName(v string) *ListSensitiveColumnsDetailRequest
	GetSchemaName() *string
	SetTableName(v string) *ListSensitiveColumnsDetailRequest
	GetTableName() *string
	SetTid(v int64) *ListSensitiveColumnsDetailRequest
	GetTid() *int64
}

type ListSensitiveColumnsDetailRequest struct {
	// The name of the field. You can call the [ListSensitiveColumns](https://help.aliyun.com/document_detail/188103.html) operation to obtain the name of the field.
	//
	// >  You can also call the [ListColumns](https://help.aliyun.com/document_detail/141870.html) operation to obtain the name of the field.
	//
	// This parameter is required.
	//
	// example:
	//
	// ColumnName_test
	ColumnName *string `json:"ColumnName,omitempty" xml:"ColumnName,omitempty"`
	// The ID of the database. The database can be a physical database or a logical database.
	//
	// 	- To obtain the ID of a physical database, call the [ListDatabases](https://help.aliyun.com/document_detail/141873.html) or [SearchDatabase](https://help.aliyun.com/document_detail/141876.html) operation.
	//
	// 	- To obtain the ID of a logical database, call the [ListLogicDatabases](https://help.aliyun.com/document_detail/141874.html) or [SearchDatabase](https://help.aliyun.com/document_detail/141876.html) operation.
	//
	// example:
	//
	// 1860****
	DbId *int64 `json:"DbId,omitempty" xml:"DbId,omitempty"`
	// Specifies whether the database is a logical database. Valid values:
	//
	// 	- **true**: The database is a logical database.
	//
	// 	- **false**: The database is a physical database.
	//
	// example:
	//
	// false
	Logic *bool `json:"Logic,omitempty" xml:"Logic,omitempty"`
	// The name of the database. You can call the [ListSensitiveColumns](https://help.aliyun.com/document_detail/188103.html) operation to obtain the name of the database.
	//
	// > 	- You can also call the [ListDatabases](https://help.aliyun.com/document_detail/141873.html) or [SearchDatabase](https://help.aliyun.com/document_detail/141876.html) operation to obtain the name of a physical database.
	//
	// > 	- You can also call the [ListLogicDatabases](https://help.aliyun.com/document_detail/141874.html) or [SearchDatabase](https://help.aliyun.com/document_detail/141876.html) operation to obtain the name of a logical database.
	//
	// This parameter is required.
	//
	// example:
	//
	// SchemaName_test
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	// The name of the table. You can call the [ListSensitiveColumns](https://help.aliyun.com/document_detail/188103.html) operation to obtain the name of the table.
	//
	// >  You can also call the [ListTables](https://help.aliyun.com/document_detail/141878.html) operation to obtain the name of the table.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// The ID of the tenant. You can call the [GetUserActiveTenant](https://help.aliyun.com/document_detail/198073.html) or [ListUserTenants](https://help.aliyun.com/document_detail/198074.html) operation to obtain the tenant ID.
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s ListSensitiveColumnsDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSensitiveColumnsDetailRequest) GoString() string {
	return s.String()
}

func (s *ListSensitiveColumnsDetailRequest) GetColumnName() *string {
	return s.ColumnName
}

func (s *ListSensitiveColumnsDetailRequest) GetDbId() *int64 {
	return s.DbId
}

func (s *ListSensitiveColumnsDetailRequest) GetLogic() *bool {
	return s.Logic
}

func (s *ListSensitiveColumnsDetailRequest) GetSchemaName() *string {
	return s.SchemaName
}

func (s *ListSensitiveColumnsDetailRequest) GetTableName() *string {
	return s.TableName
}

func (s *ListSensitiveColumnsDetailRequest) GetTid() *int64 {
	return s.Tid
}

func (s *ListSensitiveColumnsDetailRequest) SetColumnName(v string) *ListSensitiveColumnsDetailRequest {
	s.ColumnName = &v
	return s
}

func (s *ListSensitiveColumnsDetailRequest) SetDbId(v int64) *ListSensitiveColumnsDetailRequest {
	s.DbId = &v
	return s
}

func (s *ListSensitiveColumnsDetailRequest) SetLogic(v bool) *ListSensitiveColumnsDetailRequest {
	s.Logic = &v
	return s
}

func (s *ListSensitiveColumnsDetailRequest) SetSchemaName(v string) *ListSensitiveColumnsDetailRequest {
	s.SchemaName = &v
	return s
}

func (s *ListSensitiveColumnsDetailRequest) SetTableName(v string) *ListSensitiveColumnsDetailRequest {
	s.TableName = &v
	return s
}

func (s *ListSensitiveColumnsDetailRequest) SetTid(v int64) *ListSensitiveColumnsDetailRequest {
	s.Tid = &v
	return s
}

func (s *ListSensitiveColumnsDetailRequest) Validate() error {
	return dara.Validate(s)
}
