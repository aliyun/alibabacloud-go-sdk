// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTableSummaryModel interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTime(v string) *TableSummaryModel
	GetCreateTime() *string
	SetCreatedBySource(v string) *TableSummaryModel
	GetCreatedBySource() *string
	SetCreatedByUser(v string) *TableSummaryModel
	GetCreatedByUser() *string
	SetDescription(v string) *TableSummaryModel
	GetDescription() *string
	SetMvDetailModel(v *OpenStructMvDetailModel) *TableSummaryModel
	GetMvDetailModel() *OpenStructMvDetailModel
	SetOwner(v string) *TableSummaryModel
	GetOwner() *string
	SetSQL(v string) *TableSummaryModel
	GetSQL() *string
	SetSchemaName(v string) *TableSummaryModel
	GetSchemaName() *string
	SetTableName(v string) *TableSummaryModel
	GetTableName() *string
	SetTableSize(v int64) *TableSummaryModel
	GetTableSize() *int64
	SetTableType(v string) *TableSummaryModel
	GetTableType() *string
	SetUpdateTime(v string) *TableSummaryModel
	GetUpdateTime() *string
}

type TableSummaryModel struct {
	// The creation time.
	//
	// example:
	//
	// 2023-01-05 13:17:55
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The creation source of the table or view.
	CreatedBySource *string `json:"CreatedBySource,omitempty" xml:"CreatedBySource,omitempty"`
	// The user who created the table or view.
	CreatedByUser *string `json:"CreatedByUser,omitempty" xml:"CreatedByUser,omitempty"`
	// The description.
	//
	// example:
	//
	// descrtiption
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The materialized view details.
	MvDetailModel *OpenStructMvDetailModel `json:"MvDetailModel,omitempty" xml:"MvDetailModel,omitempty"`
	// The owner.
	//
	// example:
	//
	// abc
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The SQL statement used to create the table or view.
	//
	// example:
	//
	// create table(a varchar)
	SQL *string `json:"SQL,omitempty" xml:"SQL,omitempty"`
	// The logical database name.
	//
	// example:
	//
	// schemaName
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	// The logical table name or logical view name.
	//
	// example:
	//
	// tableName
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// The table size in bytes.
	//
	// example:
	//
	// 1234
	TableSize *int64 `json:"TableSize,omitempty" xml:"TableSize,omitempty"`
	// The table or view type.
	//
	// example:
	//
	// PHYSICAL_SCHEMA_NAME
	TableType *string `json:"TableType,omitempty" xml:"TableType,omitempty"`
	// The update time.
	//
	// example:
	//
	// 2023-01-05 13:17:55
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s TableSummaryModel) String() string {
	return dara.Prettify(s)
}

func (s TableSummaryModel) GoString() string {
	return s.String()
}

func (s *TableSummaryModel) GetCreateTime() *string {
	return s.CreateTime
}

func (s *TableSummaryModel) GetCreatedBySource() *string {
	return s.CreatedBySource
}

func (s *TableSummaryModel) GetCreatedByUser() *string {
	return s.CreatedByUser
}

func (s *TableSummaryModel) GetDescription() *string {
	return s.Description
}

func (s *TableSummaryModel) GetMvDetailModel() *OpenStructMvDetailModel {
	return s.MvDetailModel
}

func (s *TableSummaryModel) GetOwner() *string {
	return s.Owner
}

func (s *TableSummaryModel) GetSQL() *string {
	return s.SQL
}

func (s *TableSummaryModel) GetSchemaName() *string {
	return s.SchemaName
}

func (s *TableSummaryModel) GetTableName() *string {
	return s.TableName
}

func (s *TableSummaryModel) GetTableSize() *int64 {
	return s.TableSize
}

func (s *TableSummaryModel) GetTableType() *string {
	return s.TableType
}

func (s *TableSummaryModel) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *TableSummaryModel) SetCreateTime(v string) *TableSummaryModel {
	s.CreateTime = &v
	return s
}

func (s *TableSummaryModel) SetCreatedBySource(v string) *TableSummaryModel {
	s.CreatedBySource = &v
	return s
}

func (s *TableSummaryModel) SetCreatedByUser(v string) *TableSummaryModel {
	s.CreatedByUser = &v
	return s
}

func (s *TableSummaryModel) SetDescription(v string) *TableSummaryModel {
	s.Description = &v
	return s
}

func (s *TableSummaryModel) SetMvDetailModel(v *OpenStructMvDetailModel) *TableSummaryModel {
	s.MvDetailModel = v
	return s
}

func (s *TableSummaryModel) SetOwner(v string) *TableSummaryModel {
	s.Owner = &v
	return s
}

func (s *TableSummaryModel) SetSQL(v string) *TableSummaryModel {
	s.SQL = &v
	return s
}

func (s *TableSummaryModel) SetSchemaName(v string) *TableSummaryModel {
	s.SchemaName = &v
	return s
}

func (s *TableSummaryModel) SetTableName(v string) *TableSummaryModel {
	s.TableName = &v
	return s
}

func (s *TableSummaryModel) SetTableSize(v int64) *TableSummaryModel {
	s.TableSize = &v
	return s
}

func (s *TableSummaryModel) SetTableType(v string) *TableSummaryModel {
	s.TableType = &v
	return s
}

func (s *TableSummaryModel) SetUpdateTime(v string) *TableSummaryModel {
	s.UpdateTime = &v
	return s
}

func (s *TableSummaryModel) Validate() error {
	if s.MvDetailModel != nil {
		if err := s.MvDetailModel.Validate(); err != nil {
			return err
		}
	}
	return nil
}
