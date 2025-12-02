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
	SetCreatedBySource(v string) *TableDetailModel
	GetCreatedBySource() *string
	SetCreatedByUser(v string) *TableDetailModel
	GetCreatedByUser() *string
	SetDescription(v string) *TableDetailModel
	GetDescription() *string
	SetLocation(v string) *TableDetailModel
	GetLocation() *string
	SetOwner(v string) *TableDetailModel
	GetOwner() *string
	SetParameters(v map[string]*string) *TableDetailModel
	GetParameters() map[string]*string
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
	Catalog         *string            `json:"Catalog,omitempty" xml:"Catalog,omitempty"`
	Columns         []*ColDetailModel  `json:"Columns,omitempty" xml:"Columns,omitempty" type:"Repeated"`
	CreateTime      *string            `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CreatedBySource *string            `json:"CreatedBySource,omitempty" xml:"CreatedBySource,omitempty"`
	CreatedByUser   *string            `json:"CreatedByUser,omitempty" xml:"CreatedByUser,omitempty"`
	Description     *string            `json:"Description,omitempty" xml:"Description,omitempty"`
	Location        *string            `json:"Location,omitempty" xml:"Location,omitempty"`
	Owner           *string            `json:"Owner,omitempty" xml:"Owner,omitempty"`
	Parameters      map[string]*string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	SchemaName      *string            `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	TableName       *string            `json:"TableName,omitempty" xml:"TableName,omitempty"`
	TableType       *string            `json:"TableType,omitempty" xml:"TableType,omitempty"`
	UpdateTime      *string            `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
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

func (s *TableDetailModel) GetCreatedBySource() *string {
	return s.CreatedBySource
}

func (s *TableDetailModel) GetCreatedByUser() *string {
	return s.CreatedByUser
}

func (s *TableDetailModel) GetDescription() *string {
	return s.Description
}

func (s *TableDetailModel) GetLocation() *string {
	return s.Location
}

func (s *TableDetailModel) GetOwner() *string {
	return s.Owner
}

func (s *TableDetailModel) GetParameters() map[string]*string {
	return s.Parameters
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

func (s *TableDetailModel) SetCreatedBySource(v string) *TableDetailModel {
	s.CreatedBySource = &v
	return s
}

func (s *TableDetailModel) SetCreatedByUser(v string) *TableDetailModel {
	s.CreatedByUser = &v
	return s
}

func (s *TableDetailModel) SetDescription(v string) *TableDetailModel {
	s.Description = &v
	return s
}

func (s *TableDetailModel) SetLocation(v string) *TableDetailModel {
	s.Location = &v
	return s
}

func (s *TableDetailModel) SetOwner(v string) *TableDetailModel {
	s.Owner = &v
	return s
}

func (s *TableDetailModel) SetParameters(v map[string]*string) *TableDetailModel {
	s.Parameters = v
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
	if s.Columns != nil {
		for _, item := range s.Columns {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
