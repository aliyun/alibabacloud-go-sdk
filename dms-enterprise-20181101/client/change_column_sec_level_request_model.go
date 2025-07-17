// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeColumnSecLevelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetColumnName(v string) *ChangeColumnSecLevelRequest
	GetColumnName() *string
	SetDbId(v int64) *ChangeColumnSecLevelRequest
	GetDbId() *int64
	SetIsLogic(v bool) *ChangeColumnSecLevelRequest
	GetIsLogic() *bool
	SetNewLevel(v string) *ChangeColumnSecLevelRequest
	GetNewLevel() *string
	SetSchemaName(v string) *ChangeColumnSecLevelRequest
	GetSchemaName() *string
	SetTableName(v string) *ChangeColumnSecLevelRequest
	GetTableName() *string
	SetTid(v int64) *ChangeColumnSecLevelRequest
	GetTid() *int64
}

type ChangeColumnSecLevelRequest struct {
	// The name of the field. You can call the [ListSensitiveColumns](https://help.aliyun.com/document_detail/188103.html) operation to obtain the name of the field.
	//
	// > You can also call the [ListColumns](https://help.aliyun.com/document_detail/141870.html) operation to obtain the name of the field.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_column
	ColumnName *string `json:"ColumnName,omitempty" xml:"ColumnName,omitempty"`
	// The ID of the database. You can call the [SearchDatabase](https://help.aliyun.com/document_detail/141876.html) operation to obtain the ID of the database.
	//
	// > You can also call the [ListDatabases](https://help.aliyun.com/document_detail/141873.html) operation to obtain the ID of a physical database and the [ListLogicDatabases](https://help.aliyun.com/document_detail/141874.html) operation to obtain the ID of a logical database.
	//
	// This parameter is required.
	//
	// example:
	//
	// 325
	DbId *int64 `json:"DbId,omitempty" xml:"DbId,omitempty"`
	// Specifies whether the database is a logical database. Valid values:
	//
	// 	- true: The database is a physical database.
	//
	// 	- false: The database is a logical database.
	//
	// This parameter is required.
	//
	// example:
	//
	// false
	IsLogic *bool `json:"IsLogic,omitempty" xml:"IsLogic,omitempty"`
	// The new sensitivity level of the field that you want to specify. Valid values:
	//
	// 	- INNER: low sensitivity level
	//
	// 	- SENSITIVE: medium sensitivity level
	//
	// 	- CONFIDENTIAL: high sensitivity level
	//
	// This parameter is required.
	//
	// example:
	//
	// SENSITIVE
	NewLevel *string `json:"NewLevel,omitempty" xml:"NewLevel,omitempty"`
	// The name of the database. You can call the [ListSensitiveColumns](https://help.aliyun.com/document_detail/188103.html) operation to obtain the name of the database.
	//
	// 	- You can also call the [SearchDatabase](https://help.aliyun.com/document_detail/141876.html) operation to obtain the name of the database.
	//
	// 	- You can also call the [ListDatabases](https://help.aliyun.com/document_detail/141873.html) operation to obtain the name of a physical database and the [ListLogicDatabases](https://help.aliyun.com/document_detail/141874.html) operation to obtain the name of a logical database.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_schema
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	// The name of the table. You can call the [ListSensitiveColumns](https://help.aliyun.com/document_detail/188103.html) operation to obtain the name of the table.
	//
	// > You can also call the [ListTables](https://help.aliyun.com/document_detail/141878.html) operation to obtain the name of the table.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_table
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// The ID of the tenant.
	//
	// > To view the ID of the tenant, go to the Data Management (DMS) console and move the pointer over the profile picture in the upper-right corner. For more information, see [View information about the current tenant](https://help.aliyun.com/document_detail/181330.html).
	//
	// example:
	//
	// 43253
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s ChangeColumnSecLevelRequest) String() string {
	return dara.Prettify(s)
}

func (s ChangeColumnSecLevelRequest) GoString() string {
	return s.String()
}

func (s *ChangeColumnSecLevelRequest) GetColumnName() *string {
	return s.ColumnName
}

func (s *ChangeColumnSecLevelRequest) GetDbId() *int64 {
	return s.DbId
}

func (s *ChangeColumnSecLevelRequest) GetIsLogic() *bool {
	return s.IsLogic
}

func (s *ChangeColumnSecLevelRequest) GetNewLevel() *string {
	return s.NewLevel
}

func (s *ChangeColumnSecLevelRequest) GetSchemaName() *string {
	return s.SchemaName
}

func (s *ChangeColumnSecLevelRequest) GetTableName() *string {
	return s.TableName
}

func (s *ChangeColumnSecLevelRequest) GetTid() *int64 {
	return s.Tid
}

func (s *ChangeColumnSecLevelRequest) SetColumnName(v string) *ChangeColumnSecLevelRequest {
	s.ColumnName = &v
	return s
}

func (s *ChangeColumnSecLevelRequest) SetDbId(v int64) *ChangeColumnSecLevelRequest {
	s.DbId = &v
	return s
}

func (s *ChangeColumnSecLevelRequest) SetIsLogic(v bool) *ChangeColumnSecLevelRequest {
	s.IsLogic = &v
	return s
}

func (s *ChangeColumnSecLevelRequest) SetNewLevel(v string) *ChangeColumnSecLevelRequest {
	s.NewLevel = &v
	return s
}

func (s *ChangeColumnSecLevelRequest) SetSchemaName(v string) *ChangeColumnSecLevelRequest {
	s.SchemaName = &v
	return s
}

func (s *ChangeColumnSecLevelRequest) SetTableName(v string) *ChangeColumnSecLevelRequest {
	s.TableName = &v
	return s
}

func (s *ChangeColumnSecLevelRequest) SetTid(v int64) *ChangeColumnSecLevelRequest {
	s.Tid = &v
	return s
}

func (s *ChangeColumnSecLevelRequest) Validate() error {
	return dara.Validate(s)
}
