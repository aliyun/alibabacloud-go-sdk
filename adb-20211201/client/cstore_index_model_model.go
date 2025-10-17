// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCstoreIndexModel interface {
	dara.Model
	String() string
	GoString() string
	SetColumnOrds(v []*string) *CstoreIndexModel
	GetColumnOrds() []*string
	SetCreateTime(v string) *CstoreIndexModel
	GetCreateTime() *string
	SetDatabaseName(v string) *CstoreIndexModel
	GetDatabaseName() *string
	SetIndexColumns(v []*FieldSchemaModel) *CstoreIndexModel
	GetIndexColumns() []*FieldSchemaModel
	SetIndexName(v string) *CstoreIndexModel
	GetIndexName() *string
	SetIndexType(v string) *CstoreIndexModel
	GetIndexType() *string
	SetOptions(v map[string]*string) *CstoreIndexModel
	GetOptions() map[string]*string
	SetPhysicalTableName(v string) *CstoreIndexModel
	GetPhysicalTableName() *string
	SetUpdateTime(v string) *CstoreIndexModel
	GetUpdateTime() *string
}

type CstoreIndexModel struct {
	ColumnOrds        []*string           `json:"ColumnOrds,omitempty" xml:"ColumnOrds,omitempty" type:"Repeated"`
	CreateTime        *string             `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DatabaseName      *string             `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	IndexColumns      []*FieldSchemaModel `json:"IndexColumns,omitempty" xml:"IndexColumns,omitempty" type:"Repeated"`
	IndexName         *string             `json:"IndexName,omitempty" xml:"IndexName,omitempty"`
	IndexType         *string             `json:"IndexType,omitempty" xml:"IndexType,omitempty"`
	Options           map[string]*string  `json:"Options,omitempty" xml:"Options,omitempty"`
	PhysicalTableName *string             `json:"PhysicalTableName,omitempty" xml:"PhysicalTableName,omitempty"`
	UpdateTime        *string             `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s CstoreIndexModel) String() string {
	return dara.Prettify(s)
}

func (s CstoreIndexModel) GoString() string {
	return s.String()
}

func (s *CstoreIndexModel) GetColumnOrds() []*string {
	return s.ColumnOrds
}

func (s *CstoreIndexModel) GetCreateTime() *string {
	return s.CreateTime
}

func (s *CstoreIndexModel) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *CstoreIndexModel) GetIndexColumns() []*FieldSchemaModel {
	return s.IndexColumns
}

func (s *CstoreIndexModel) GetIndexName() *string {
	return s.IndexName
}

func (s *CstoreIndexModel) GetIndexType() *string {
	return s.IndexType
}

func (s *CstoreIndexModel) GetOptions() map[string]*string {
	return s.Options
}

func (s *CstoreIndexModel) GetPhysicalTableName() *string {
	return s.PhysicalTableName
}

func (s *CstoreIndexModel) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *CstoreIndexModel) SetColumnOrds(v []*string) *CstoreIndexModel {
	s.ColumnOrds = v
	return s
}

func (s *CstoreIndexModel) SetCreateTime(v string) *CstoreIndexModel {
	s.CreateTime = &v
	return s
}

func (s *CstoreIndexModel) SetDatabaseName(v string) *CstoreIndexModel {
	s.DatabaseName = &v
	return s
}

func (s *CstoreIndexModel) SetIndexColumns(v []*FieldSchemaModel) *CstoreIndexModel {
	s.IndexColumns = v
	return s
}

func (s *CstoreIndexModel) SetIndexName(v string) *CstoreIndexModel {
	s.IndexName = &v
	return s
}

func (s *CstoreIndexModel) SetIndexType(v string) *CstoreIndexModel {
	s.IndexType = &v
	return s
}

func (s *CstoreIndexModel) SetOptions(v map[string]*string) *CstoreIndexModel {
	s.Options = v
	return s
}

func (s *CstoreIndexModel) SetPhysicalTableName(v string) *CstoreIndexModel {
	s.PhysicalTableName = &v
	return s
}

func (s *CstoreIndexModel) SetUpdateTime(v string) *CstoreIndexModel {
	s.UpdateTime = &v
	return s
}

func (s *CstoreIndexModel) Validate() error {
	if s.IndexColumns != nil {
		for _, item := range s.IndexColumns {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
