// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSchema interface {
	dara.Model
	String() string
	GoString() string
	SetIndexSortConfig(v []*SchemaIndexSortConfig) *Schema
	GetIndexSortConfig() []*SchemaIndexSortConfig
	SetIndexes(v *SchemaIndexes) *Schema
	GetIndexes() *SchemaIndexes
	SetName(v string) *Schema
	GetName() *string
	SetRouteField(v string) *Schema
	GetRouteField() *string
	SetRouteFieldValues(v []*string) *Schema
	GetRouteFieldValues() []*string
	SetSecondRouteField(v string) *Schema
	GetSecondRouteField() *string
	SetTables(v map[string]*SchemaTablesValue) *Schema
	GetTables() map[string]*SchemaTablesValue
	SetTtlField(v *SchemaTtlField) *Schema
	GetTtlField() *SchemaTtlField
}

type Schema struct {
	IndexSortConfig  []*SchemaIndexSortConfig      `json:"indexSortConfig,omitempty" xml:"indexSortConfig,omitempty" type:"Repeated"`
	Indexes          *SchemaIndexes                `json:"indexes,omitempty" xml:"indexes,omitempty" type:"Struct"`
	Name             *string                       `json:"name,omitempty" xml:"name,omitempty"`
	RouteField       *string                       `json:"routeField,omitempty" xml:"routeField,omitempty"`
	RouteFieldValues []*string                     `json:"routeFieldValues,omitempty" xml:"routeFieldValues,omitempty" type:"Repeated"`
	SecondRouteField *string                       `json:"secondRouteField,omitempty" xml:"secondRouteField,omitempty"`
	Tables           map[string]*SchemaTablesValue `json:"tables,omitempty" xml:"tables,omitempty"`
	TtlField         *SchemaTtlField               `json:"ttlField,omitempty" xml:"ttlField,omitempty" type:"Struct"`
}

func (s Schema) String() string {
	return dara.Prettify(s)
}

func (s Schema) GoString() string {
	return s.String()
}

func (s *Schema) GetIndexSortConfig() []*SchemaIndexSortConfig {
	return s.IndexSortConfig
}

func (s *Schema) GetIndexes() *SchemaIndexes {
	return s.Indexes
}

func (s *Schema) GetName() *string {
	return s.Name
}

func (s *Schema) GetRouteField() *string {
	return s.RouteField
}

func (s *Schema) GetRouteFieldValues() []*string {
	return s.RouteFieldValues
}

func (s *Schema) GetSecondRouteField() *string {
	return s.SecondRouteField
}

func (s *Schema) GetTables() map[string]*SchemaTablesValue {
	return s.Tables
}

func (s *Schema) GetTtlField() *SchemaTtlField {
	return s.TtlField
}

func (s *Schema) SetIndexSortConfig(v []*SchemaIndexSortConfig) *Schema {
	s.IndexSortConfig = v
	return s
}

func (s *Schema) SetIndexes(v *SchemaIndexes) *Schema {
	s.Indexes = v
	return s
}

func (s *Schema) SetName(v string) *Schema {
	s.Name = &v
	return s
}

func (s *Schema) SetRouteField(v string) *Schema {
	s.RouteField = &v
	return s
}

func (s *Schema) SetRouteFieldValues(v []*string) *Schema {
	s.RouteFieldValues = v
	return s
}

func (s *Schema) SetSecondRouteField(v string) *Schema {
	s.SecondRouteField = &v
	return s
}

func (s *Schema) SetTables(v map[string]*SchemaTablesValue) *Schema {
	s.Tables = v
	return s
}

func (s *Schema) SetTtlField(v *SchemaTtlField) *Schema {
	s.TtlField = v
	return s
}

func (s *Schema) Validate() error {
	if s.IndexSortConfig != nil {
		for _, item := range s.IndexSortConfig {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Indexes != nil {
		if err := s.Indexes.Validate(); err != nil {
			return err
		}
	}
	if s.TtlField != nil {
		if err := s.TtlField.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SchemaIndexSortConfig struct {
	Direction *string `json:"direction,omitempty" xml:"direction,omitempty"`
	Field     *string `json:"field,omitempty" xml:"field,omitempty"`
}

func (s SchemaIndexSortConfig) String() string {
	return dara.Prettify(s)
}

func (s SchemaIndexSortConfig) GoString() string {
	return s.String()
}

func (s *SchemaIndexSortConfig) GetDirection() *string {
	return s.Direction
}

func (s *SchemaIndexSortConfig) GetField() *string {
	return s.Field
}

func (s *SchemaIndexSortConfig) SetDirection(v string) *SchemaIndexSortConfig {
	s.Direction = &v
	return s
}

func (s *SchemaIndexSortConfig) SetField(v string) *SchemaIndexSortConfig {
	s.Field = &v
	return s
}

func (s *SchemaIndexSortConfig) Validate() error {
	return dara.Validate(s)
}

type SchemaIndexes struct {
	FilterFields []*string                                  `json:"filterFields,omitempty" xml:"filterFields,omitempty" type:"Repeated"`
	SearchFields map[string]*SchemaIndexesSearchFieldsValue `json:"searchFields,omitempty" xml:"searchFields,omitempty"`
}

func (s SchemaIndexes) String() string {
	return dara.Prettify(s)
}

func (s SchemaIndexes) GoString() string {
	return s.String()
}

func (s *SchemaIndexes) GetFilterFields() []*string {
	return s.FilterFields
}

func (s *SchemaIndexes) GetSearchFields() map[string]*SchemaIndexesSearchFieldsValue {
	return s.SearchFields
}

func (s *SchemaIndexes) SetFilterFields(v []*string) *SchemaIndexes {
	s.FilterFields = v
	return s
}

func (s *SchemaIndexes) SetSearchFields(v map[string]*SchemaIndexesSearchFieldsValue) *SchemaIndexes {
	s.SearchFields = v
	return s
}

func (s *SchemaIndexes) Validate() error {
	return dara.Validate(s)
}

type SchemaTtlField struct {
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	Ttl  *int64  `json:"ttl,omitempty" xml:"ttl,omitempty"`
}

func (s SchemaTtlField) String() string {
	return dara.Prettify(s)
}

func (s SchemaTtlField) GoString() string {
	return s.String()
}

func (s *SchemaTtlField) GetName() *string {
	return s.Name
}

func (s *SchemaTtlField) GetTtl() *int64 {
	return s.Ttl
}

func (s *SchemaTtlField) SetName(v string) *SchemaTtlField {
	s.Name = &v
	return s
}

func (s *SchemaTtlField) SetTtl(v int64) *SchemaTtlField {
	s.Ttl = &v
	return s
}

func (s *SchemaTtlField) Validate() error {
	return dara.Validate(s)
}
