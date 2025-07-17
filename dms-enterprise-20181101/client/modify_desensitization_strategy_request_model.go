// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDesensitizationStrategyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetColumnName(v string) *ModifyDesensitizationStrategyRequest
	GetColumnName() *string
	SetDbId(v int32) *ModifyDesensitizationStrategyRequest
	GetDbId() *int32
	SetIsDefault(v bool) *ModifyDesensitizationStrategyRequest
	GetIsDefault() *bool
	SetIsLogic(v bool) *ModifyDesensitizationStrategyRequest
	GetIsLogic() *bool
	SetIsReset(v bool) *ModifyDesensitizationStrategyRequest
	GetIsReset() *bool
	SetRuleId(v int32) *ModifyDesensitizationStrategyRequest
	GetRuleId() *int32
	SetSchemaName(v string) *ModifyDesensitizationStrategyRequest
	GetSchemaName() *string
	SetTableName(v string) *ModifyDesensitizationStrategyRequest
	GetTableName() *string
	SetTid(v int64) *ModifyDesensitizationStrategyRequest
	GetTid() *int64
}

type ModifyDesensitizationStrategyRequest struct {
	// The name of the field. You can call the [ListSensitiveColumns](https://help.aliyun.com/document_detail/188103.html) operation to obtain the field name.
	//
	// >  You can also call the [ListColumns](https://help.aliyun.com/document_detail/141870.html) operation to obtain the field name.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_column
	ColumnName *string `json:"ColumnName,omitempty" xml:"ColumnName,omitempty"`
	// The ID of the database. You can call the [ListDatabases](https://help.aliyun.com/document_detail/141873.html) operation to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	DbId *int32 `json:"DbId,omitempty" xml:"DbId,omitempty"`
	// The desensitization algorithm of the field setting. The default value is false. The values are as follows:
	//
	// - **true**: default desensitization algorithm.
	//
	// - **false*	- :semi-desensitization algorithm.
	//
	// example:
	//
	// false
	IsDefault *bool `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// Specifies whether the database is a logical database. Valid values:
	//
	// 	- **true**: The database is a physical database.
	//
	// 	- **false**: The database is a logical database.
	//
	// This parameter is required.
	//
	// example:
	//
	// false
	IsLogic *bool `json:"IsLogic,omitempty" xml:"IsLogic,omitempty"`
	// Specifies whether to reset the masking rule. Valid value:
	//
	// 	- **true**: Reset the masking rule.
	//
	// 	- **false**: Do not reset the masking rule. This is the default value.
	//
	// example:
	//
	// false
	IsReset *bool `json:"IsReset,omitempty" xml:"IsReset,omitempty"`
	// The ID of the masking rule.
	//
	// example:
	//
	// 53
	RuleId *int32 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The name of the database. You can call the [ListSensitiveColumns](https://help.aliyun.com/document_detail/188103.html) operation to obtain the database name.
	//
	// >
	//
	// 	- If the database is a physical database, you can call the [ListDatabases](https://help.aliyun.com/document_detail/141873.html) or [SearchDatabase](https://help.aliyun.com/document_detail/141876.html) operation to obtain the database name.
	//
	// 	- If the database is a logical database, you can call the [ListLogicDatabases](https://help.aliyun.com/document_detail/141874.html) or [SearchDatabase](https://help.aliyun.com/document_detail/141876.html) operation to obtain the name of the database.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_schema
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	// The name of the table. You can call the [ListSensitiveColumns](https://help.aliyun.com/document_detail/188103.html) operation to obtain the table name.
	//
	// >  You can also call the [ListTables](https://help.aliyun.com/document_detail/141878.html) operation to obtain the table name.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_table
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// The ID of the tenant.
	//
	// >  To view the ID of the tenant, go to the Data Management (DMS) console and move the pointer over the profile picture in the upper-right corner. For more information, see [View information about the current tenant](https://help.aliyun.com/document_detail/181330.html) in the topic "Manage DMS tenants."
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s ModifyDesensitizationStrategyRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDesensitizationStrategyRequest) GoString() string {
	return s.String()
}

func (s *ModifyDesensitizationStrategyRequest) GetColumnName() *string {
	return s.ColumnName
}

func (s *ModifyDesensitizationStrategyRequest) GetDbId() *int32 {
	return s.DbId
}

func (s *ModifyDesensitizationStrategyRequest) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *ModifyDesensitizationStrategyRequest) GetIsLogic() *bool {
	return s.IsLogic
}

func (s *ModifyDesensitizationStrategyRequest) GetIsReset() *bool {
	return s.IsReset
}

func (s *ModifyDesensitizationStrategyRequest) GetRuleId() *int32 {
	return s.RuleId
}

func (s *ModifyDesensitizationStrategyRequest) GetSchemaName() *string {
	return s.SchemaName
}

func (s *ModifyDesensitizationStrategyRequest) GetTableName() *string {
	return s.TableName
}

func (s *ModifyDesensitizationStrategyRequest) GetTid() *int64 {
	return s.Tid
}

func (s *ModifyDesensitizationStrategyRequest) SetColumnName(v string) *ModifyDesensitizationStrategyRequest {
	s.ColumnName = &v
	return s
}

func (s *ModifyDesensitizationStrategyRequest) SetDbId(v int32) *ModifyDesensitizationStrategyRequest {
	s.DbId = &v
	return s
}

func (s *ModifyDesensitizationStrategyRequest) SetIsDefault(v bool) *ModifyDesensitizationStrategyRequest {
	s.IsDefault = &v
	return s
}

func (s *ModifyDesensitizationStrategyRequest) SetIsLogic(v bool) *ModifyDesensitizationStrategyRequest {
	s.IsLogic = &v
	return s
}

func (s *ModifyDesensitizationStrategyRequest) SetIsReset(v bool) *ModifyDesensitizationStrategyRequest {
	s.IsReset = &v
	return s
}

func (s *ModifyDesensitizationStrategyRequest) SetRuleId(v int32) *ModifyDesensitizationStrategyRequest {
	s.RuleId = &v
	return s
}

func (s *ModifyDesensitizationStrategyRequest) SetSchemaName(v string) *ModifyDesensitizationStrategyRequest {
	s.SchemaName = &v
	return s
}

func (s *ModifyDesensitizationStrategyRequest) SetTableName(v string) *ModifyDesensitizationStrategyRequest {
	s.TableName = &v
	return s
}

func (s *ModifyDesensitizationStrategyRequest) SetTid(v int64) *ModifyDesensitizationStrategyRequest {
	s.Tid = &v
	return s
}

func (s *ModifyDesensitizationStrategyRequest) Validate() error {
	return dara.Validate(s)
}
