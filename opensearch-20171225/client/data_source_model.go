// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDataSource interface {
	dara.Model
	String() string
	GoString() string
	SetFields(v []map[string]*string) *DataSource
	GetFields() []map[string]*string
	SetKeyField(v string) *DataSource
	GetKeyField() *string
	SetParameters(v map[string]interface{}) *DataSource
	GetParameters() map[string]interface{}
	SetPlugins(v map[string]*DataSourcePluginsValue) *DataSource
	GetPlugins() map[string]*DataSourcePluginsValue
	SetSchemaName(v string) *DataSource
	GetSchemaName() *string
	SetTableName(v string) *DataSource
	GetTableName() *string
	SetType(v string) *DataSource
	GetType() *string
}

type DataSource struct {
	Fields     []map[string]*string               `json:"fields,omitempty" xml:"fields,omitempty" type:"Repeated"`
	KeyField   *string                            `json:"keyField,omitempty" xml:"keyField,omitempty"`
	Parameters map[string]interface{}             `json:"parameters,omitempty" xml:"parameters,omitempty"`
	Plugins    map[string]*DataSourcePluginsValue `json:"plugins,omitempty" xml:"plugins,omitempty"`
	SchemaName *string                            `json:"schemaName,omitempty" xml:"schemaName,omitempty"`
	TableName  *string                            `json:"tableName,omitempty" xml:"tableName,omitempty"`
	Type       *string                            `json:"type,omitempty" xml:"type,omitempty"`
}

func (s DataSource) String() string {
	return dara.Prettify(s)
}

func (s DataSource) GoString() string {
	return s.String()
}

func (s *DataSource) GetFields() []map[string]*string {
	return s.Fields
}

func (s *DataSource) GetKeyField() *string {
	return s.KeyField
}

func (s *DataSource) GetParameters() map[string]interface{} {
	return s.Parameters
}

func (s *DataSource) GetPlugins() map[string]*DataSourcePluginsValue {
	return s.Plugins
}

func (s *DataSource) GetSchemaName() *string {
	return s.SchemaName
}

func (s *DataSource) GetTableName() *string {
	return s.TableName
}

func (s *DataSource) GetType() *string {
	return s.Type
}

func (s *DataSource) SetFields(v []map[string]*string) *DataSource {
	s.Fields = v
	return s
}

func (s *DataSource) SetKeyField(v string) *DataSource {
	s.KeyField = &v
	return s
}

func (s *DataSource) SetParameters(v map[string]interface{}) *DataSource {
	s.Parameters = v
	return s
}

func (s *DataSource) SetPlugins(v map[string]*DataSourcePluginsValue) *DataSource {
	s.Plugins = v
	return s
}

func (s *DataSource) SetSchemaName(v string) *DataSource {
	s.SchemaName = &v
	return s
}

func (s *DataSource) SetTableName(v string) *DataSource {
	s.TableName = &v
	return s
}

func (s *DataSource) SetType(v string) *DataSource {
	s.Type = &v
	return s
}

func (s *DataSource) Validate() error {
	return dara.Validate(s)
}
