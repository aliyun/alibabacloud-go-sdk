// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTableDetailModel interface {
	dara.Model
	String() string
	GoString() string
	SetCatalog(v string) *TableDetailModel
	GetCatalog() *string
	SetColumns(v []*ColDetailModel) *TableDetailModel
	GetColumns() []*ColDetailModel
	SetCreateTime(v string) *TableDetailModel
	GetCreateTime() *string
	SetDescription(v string) *TableDetailModel
	GetDescription() *string
	SetOwner(v string) *TableDetailModel
	GetOwner() *string
	SetSchemaName(v string) *TableDetailModel
	GetSchemaName() *string
	SetTableName(v string) *TableDetailModel
	GetTableName() *string
	SetTableType(v string) *TableDetailModel
	GetTableType() *string
	SetUpdateTime(v string) *TableDetailModel
	GetUpdateTime() *string
}

type TableDetailModel struct {
	Catalog     *string           `json:"Catalog,omitempty" xml:"Catalog,omitempty"`
	Columns     []*ColDetailModel `json:"Columns,omitempty" xml:"Columns,omitempty" type:"Repeated"`
	CreateTime  *string           `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description *string           `json:"Description,omitempty" xml:"Description,omitempty"`
	Owner       *string           `json:"Owner,omitempty" xml:"Owner,omitempty"`
	SchemaName  *string           `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	TableName   *string           `json:"TableName,omitempty" xml:"TableName,omitempty"`
	TableType   *string           `json:"TableType,omitempty" xml:"TableType,omitempty"`
	UpdateTime  *string           `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s TableDetailModel) String() string {
	return dara.Prettify(s)
}

func (s TableDetailModel) GoString() string {
	return s.String()
}

func (s *TableDetailModel) GetCatalog() *string {
	return s.Catalog
}

func (s *TableDetailModel) GetColumns() []*ColDetailModel {
	return s.Columns
}

func (s *TableDetailModel) GetCreateTime() *string {
	return s.CreateTime
}

func (s *TableDetailModel) GetDescription() *string {
	return s.Description
}

func (s *TableDetailModel) GetOwner() *string {
	return s.Owner
}

func (s *TableDetailModel) GetSchemaName() *string {
	return s.SchemaName
}

func (s *TableDetailModel) GetTableName() *string {
	return s.TableName
}

func (s *TableDetailModel) GetTableType() *string {
	return s.TableType
}

func (s *TableDetailModel) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *TableDetailModel) SetCatalog(v string) *TableDetailModel {
	s.Catalog = &v
	return s
}

func (s *TableDetailModel) SetColumns(v []*ColDetailModel) *TableDetailModel {
	s.Columns = v
	return s
}

func (s *TableDetailModel) SetCreateTime(v string) *TableDetailModel {
	s.CreateTime = &v
	return s
}

func (s *TableDetailModel) SetDescription(v string) *TableDetailModel {
	s.Description = &v
	return s
}

func (s *TableDetailModel) SetOwner(v string) *TableDetailModel {
	s.Owner = &v
	return s
}

func (s *TableDetailModel) SetSchemaName(v string) *TableDetailModel {
	s.SchemaName = &v
	return s
}

func (s *TableDetailModel) SetTableName(v string) *TableDetailModel {
	s.TableName = &v
	return s
}

func (s *TableDetailModel) SetTableType(v string) *TableDetailModel {
	s.TableType = &v
	return s
}

func (s *TableDetailModel) SetUpdateTime(v string) *TableDetailModel {
	s.UpdateTime = &v
	return s
}

func (s *TableDetailModel) Validate() error {
	return dara.Validate(s)
}
