// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTableSchema interface {
	dara.Model
	String() string
	GoString() string
	SetCollectSinkOperatorId(v string) *TableSchema
	GetCollectSinkOperatorId() *string
	SetSchema(v *Schema) *TableSchema
	GetSchema() *Schema
	SetTableName(v string) *TableSchema
	GetTableName() *string
}

type TableSchema struct {
	CollectSinkOperatorId *string `json:"collectSinkOperatorId,omitempty" xml:"collectSinkOperatorId,omitempty"`
	Schema                *Schema `json:"schema,omitempty" xml:"schema,omitempty"`
	TableName             *string `json:"tableName,omitempty" xml:"tableName,omitempty"`
}

func (s TableSchema) String() string {
	return dara.Prettify(s)
}

func (s TableSchema) GoString() string {
	return s.String()
}

func (s *TableSchema) GetCollectSinkOperatorId() *string {
	return s.CollectSinkOperatorId
}

func (s *TableSchema) GetSchema() *Schema {
	return s.Schema
}

func (s *TableSchema) GetTableName() *string {
	return s.TableName
}

func (s *TableSchema) SetCollectSinkOperatorId(v string) *TableSchema {
	s.CollectSinkOperatorId = &v
	return s
}

func (s *TableSchema) SetSchema(v *Schema) *TableSchema {
	s.Schema = v
	return s
}

func (s *TableSchema) SetTableName(v string) *TableSchema {
	s.TableName = &v
	return s
}

func (s *TableSchema) Validate() error {
	if s.Schema != nil {
		if err := s.Schema.Validate(); err != nil {
			return err
		}
	}
	return nil
}
