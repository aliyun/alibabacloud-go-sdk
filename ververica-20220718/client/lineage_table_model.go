// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLineageTable interface {
	dara.Model
	String() string
	GoString() string
	SetColumns(v []*LineageColumn) *LineageTable
	GetColumns() []*LineageColumn
	SetId(v string) *LineageTable
	GetId() *string
	SetProperties(v map[string]interface{}) *LineageTable
	GetProperties() map[string]interface{}
	SetTableName(v string) *LineageTable
	GetTableName() *string
	SetWith(v map[string]interface{}) *LineageTable
	GetWith() map[string]interface{}
}

type LineageTable struct {
	Columns    []*LineageColumn       `json:"columns,omitempty" xml:"columns,omitempty" type:"Repeated"`
	Id         *string                `json:"id,omitempty" xml:"id,omitempty"`
	Properties map[string]interface{} `json:"properties,omitempty" xml:"properties,omitempty"`
	TableName  *string                `json:"tableName,omitempty" xml:"tableName,omitempty"`
	With       map[string]interface{} `json:"with,omitempty" xml:"with,omitempty"`
}

func (s LineageTable) String() string {
	return dara.Prettify(s)
}

func (s LineageTable) GoString() string {
	return s.String()
}

func (s *LineageTable) GetColumns() []*LineageColumn {
	return s.Columns
}

func (s *LineageTable) GetId() *string {
	return s.Id
}

func (s *LineageTable) GetProperties() map[string]interface{} {
	return s.Properties
}

func (s *LineageTable) GetTableName() *string {
	return s.TableName
}

func (s *LineageTable) GetWith() map[string]interface{} {
	return s.With
}

func (s *LineageTable) SetColumns(v []*LineageColumn) *LineageTable {
	s.Columns = v
	return s
}

func (s *LineageTable) SetId(v string) *LineageTable {
	s.Id = &v
	return s
}

func (s *LineageTable) SetProperties(v map[string]interface{}) *LineageTable {
	s.Properties = v
	return s
}

func (s *LineageTable) SetTableName(v string) *LineageTable {
	s.TableName = &v
	return s
}

func (s *LineageTable) SetWith(v map[string]interface{}) *LineageTable {
	s.With = v
	return s
}

func (s *LineageTable) Validate() error {
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
