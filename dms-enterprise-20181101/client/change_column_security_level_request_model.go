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
	// This parameter is required.
	//
	// example:
	//
	// test_column
	ColumnName *string `json:"ColumnName,omitempty" xml:"ColumnName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 325**
	DbId *int64 `json:"DbId,omitempty" xml:"DbId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// false
	IsLogic *bool `json:"IsLogic,omitempty" xml:"IsLogic,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// S2
	NewSensitivityLevel *string `json:"NewSensitivityLevel,omitempty" xml:"NewSensitivityLevel,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test_schema
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test_table
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
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
