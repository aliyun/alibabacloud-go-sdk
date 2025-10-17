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
	CreateTime    *string                  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description   *string                  `json:"Description,omitempty" xml:"Description,omitempty"`
	MvDetailModel *OpenStructMvDetailModel `json:"MvDetailModel,omitempty" xml:"MvDetailModel,omitempty"`
	Owner         *string                  `json:"Owner,omitempty" xml:"Owner,omitempty"`
	SQL           *string                  `json:"SQL,omitempty" xml:"SQL,omitempty"`
	SchemaName    *string                  `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	TableName     *string                  `json:"TableName,omitempty" xml:"TableName,omitempty"`
	TableSize     *int64                   `json:"TableSize,omitempty" xml:"TableSize,omitempty"`
	TableType     *string                  `json:"TableType,omitempty" xml:"TableType,omitempty"`
	UpdateTime    *string                  `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
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
