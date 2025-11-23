// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeColumnSecurityLevelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetColumnName(v string) *ChangeColumnSecurityLevelRequest
	GetColumnName() *string
	SetDbId(v int64) *ChangeColumnSecurityLevelRequest
	GetDbId() *int64
	SetIsLogic(v bool) *ChangeColumnSecurityLevelRequest
	GetIsLogic() *bool
	SetNewSensitivityLevel(v string) *ChangeColumnSecurityLevelRequest
	GetNewSensitivityLevel() *string
	SetSchemaName(v string) *ChangeColumnSecurityLevelRequest
	GetSchemaName() *string
	SetTableName(v string) *ChangeColumnSecurityLevelRequest
	GetTableName() *string
	SetTid(v int64) *ChangeColumnSecurityLevelRequest
	GetTid() *int64
}

type ChangeColumnSecurityLevelRequest struct {
	// The name of the field. You can call the [ListSensitiveColumns](https://help.aliyun.com/document_detail/188103.html) or [ListColumns](https://help.aliyun.com/document_detail/141870.html) operation to query the column name.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_column
	ColumnName *string `json:"ColumnName,omitempty" xml:"ColumnName,omitempty"`
	// The database ID. The database can be a physical database or a logical database.
	//
	// 	- The ID of a physical database: You can call the [ListDatabases](https://help.aliyun.com/document_detail/141873.html) or [SearchDatabase](https://help.aliyun.com/document_detail/141876.html) operation to obtain the physical database ID.
	//
	// 	- To obtain the ID of a logical database, call the [ListLogicDatabases](https://help.aliyun.com/document_detail/141874.html) or [SearchDatabase](https://help.aliyun.com/document_detail/141876.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 325**
	DbId *int64 `json:"DbId,omitempty" xml:"DbId,omitempty"`
	// Specifies whether the database is a logical database. Valid values:
	//
	// 	- **true**: The database is a logical database
	//
	// 	- **false**: The database is a physical database.
	//
	// This parameter is required.
	//
	// example:
	//
	// false
	IsLogic *bool `json:"IsLogic,omitempty" xml:"IsLogic,omitempty"`
	// The new security level of the column. The valid values are the same as the sensitivity levels of the classification template that is associated with the instance. You can call the [ListSensitivityLevel](https://help.aliyun.com/document_detail/2539519.html) operation to obtain the sensitivity levels of the classification template.
	//
	// This parameter is required.
	//
	// example:
	//
	// S2
	NewSensitivityLevel *string `json:"NewSensitivityLevel,omitempty" xml:"NewSensitivityLevel,omitempty"`
	// The database name. You can call the [ListSensitiveColumns](https://help.aliyun.com/document_detail/188103.html) or [SearchDatabase](https://help.aliyun.com/document_detail/141876.html) operation to query the database name.
	//
	// > You can call the [ListDatabases](https://help.aliyun.com/document_detail/141873.html) operation to query the name of a physical database and call the [ListLogicDatabases](https://help.aliyun.com/document_detail/141874.html) operation to query the name of a logical database.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_schema
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	// The name of the table. You can call the [ListSensitiveColumns](https://help.aliyun.com/document_detail/188103.html) or [ListTables](https://help.aliyun.com/document_detail/141878.html) operation to query the table name.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_table
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// The ID of the tenant.
	//
	// > To view the tenant ID, go to the Data Management (DMS) console and move the pointer over the profile picture in the upper-right corner. For more information, see the [View information about the current tenant](https://help.aliyun.com/document_detail/181330.html) section of the "Manage DMS tenants" topic.
	//
	// example:
	//
	// 10****
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s ChangeColumnSecurityLevelRequest) String() string {
	return dara.Prettify(s)
}

func (s ChangeColumnSecurityLevelRequest) GoString() string {
	return s.String()
}

func (s *ChangeColumnSecurityLevelRequest) GetColumnName() *string {
	return s.ColumnName
}

func (s *ChangeColumnSecurityLevelRequest) GetDbId() *int64 {
	return s.DbId
}

func (s *ChangeColumnSecurityLevelRequest) GetIsLogic() *bool {
	return s.IsLogic
}

func (s *ChangeColumnSecurityLevelRequest) GetNewSensitivityLevel() *string {
	return s.NewSensitivityLevel
}

func (s *ChangeColumnSecurityLevelRequest) GetSchemaName() *string {
	return s.SchemaName
}

func (s *ChangeColumnSecurityLevelRequest) GetTableName() *string {
	return s.TableName
}

func (s *ChangeColumnSecurityLevelRequest) GetTid() *int64 {
	return s.Tid
}

func (s *ChangeColumnSecurityLevelRequest) SetColumnName(v string) *ChangeColumnSecurityLevelRequest {
	s.ColumnName = &v
	return s
}

func (s *ChangeColumnSecurityLevelRequest) SetDbId(v int64) *ChangeColumnSecurityLevelRequest {
	s.DbId = &v
	return s
}

func (s *ChangeColumnSecurityLevelRequest) SetIsLogic(v bool) *ChangeColumnSecurityLevelRequest {
	s.IsLogic = &v
	return s
}

func (s *ChangeColumnSecurityLevelRequest) SetNewSensitivityLevel(v string) *ChangeColumnSecurityLevelRequest {
	s.NewSensitivityLevel = &v
	return s
}

func (s *ChangeColumnSecurityLevelRequest) SetSchemaName(v string) *ChangeColumnSecurityLevelRequest {
	s.SchemaName = &v
	return s
}

func (s *ChangeColumnSecurityLevelRequest) SetTableName(v string) *ChangeColumnSecurityLevelRequest {
	s.TableName = &v
	return s
}

func (s *ChangeColumnSecurityLevelRequest) SetTid(v int64) *ChangeColumnSecurityLevelRequest {
	s.Tid = &v
	return s
}

func (s *ChangeColumnSecurityLevelRequest) Validate() error {
	return dara.Validate(s)
}
