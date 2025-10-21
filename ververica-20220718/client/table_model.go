// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTable interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *Table
	GetComment() *string
	SetMetadata(v map[string]*string) *Table
	GetMetadata() map[string]*string
	SetName(v string) *Table
	GetName() *string
	SetPartitionKeys(v []*string) *Table
	GetPartitionKeys() []*string
	SetProperties(v map[string]interface{}) *Table
	GetProperties() map[string]interface{}
	SetSchema(v *Schema) *Table
	GetSchema() *Schema
	SetTableType(v string) *Table
	GetTableType() *string
}

type Table struct {
	Comment  *string            `json:"comment,omitempty" xml:"comment,omitempty"`
	Metadata map[string]*string `json:"metadata,omitempty" xml:"metadata,omitempty"`
	// This parameter is required.
	Name          *string                `json:"name,omitempty" xml:"name,omitempty"`
	PartitionKeys []*string              `json:"partitionKeys,omitempty" xml:"partitionKeys,omitempty" type:"Repeated"`
	Properties    map[string]interface{} `json:"properties,omitempty" xml:"properties,omitempty"`
	Schema        *Schema                `json:"schema,omitempty" xml:"schema,omitempty"`
	// This parameter is required.
	TableType *string `json:"tableType,omitempty" xml:"tableType,omitempty"`
}

func (s Table) String() string {
	return dara.Prettify(s)
}

func (s Table) GoString() string {
	return s.String()
}

func (s *Table) GetComment() *string {
	return s.Comment
}

func (s *Table) GetMetadata() map[string]*string {
	return s.Metadata
}

func (s *Table) GetName() *string {
	return s.Name
}

func (s *Table) GetPartitionKeys() []*string {
	return s.PartitionKeys
}

func (s *Table) GetProperties() map[string]interface{} {
	return s.Properties
}

func (s *Table) GetSchema() *Schema {
	return s.Schema
}

func (s *Table) GetTableType() *string {
	return s.TableType
}

func (s *Table) SetComment(v string) *Table {
	s.Comment = &v
	return s
}

func (s *Table) SetMetadata(v map[string]*string) *Table {
	s.Metadata = v
	return s
}

func (s *Table) SetName(v string) *Table {
	s.Name = &v
	return s
}

func (s *Table) SetPartitionKeys(v []*string) *Table {
	s.PartitionKeys = v
	return s
}

func (s *Table) SetProperties(v map[string]interface{}) *Table {
	s.Properties = v
	return s
}

func (s *Table) SetSchema(v *Schema) *Table {
	s.Schema = v
	return s
}

func (s *Table) SetTableType(v string) *Table {
	s.TableType = &v
	return s
}

func (s *Table) Validate() error {
	if s.Schema != nil {
		if err := s.Schema.Validate(); err != nil {
			return err
		}
	}
	return nil
}
