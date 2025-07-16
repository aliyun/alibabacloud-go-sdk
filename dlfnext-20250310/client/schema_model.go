// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSchema interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *Schema
	GetComment() *string
	SetFields(v []*DataField) *Schema
	GetFields() []*DataField
	SetOptions(v map[string]*string) *Schema
	GetOptions() map[string]*string
	SetPartitionKeys(v []*string) *Schema
	GetPartitionKeys() []*string
	SetPrimaryKeys(v []*string) *Schema
	GetPrimaryKeys() []*string
}

type Schema struct {
	Comment       *string            `json:"comment,omitempty" xml:"comment,omitempty"`
	Fields        []*DataField       `json:"fields,omitempty" xml:"fields,omitempty" type:"Repeated"`
	Options       map[string]*string `json:"options,omitempty" xml:"options,omitempty"`
	PartitionKeys []*string          `json:"partitionKeys,omitempty" xml:"partitionKeys,omitempty" type:"Repeated"`
	PrimaryKeys   []*string          `json:"primaryKeys,omitempty" xml:"primaryKeys,omitempty" type:"Repeated"`
}

func (s Schema) String() string {
	return dara.Prettify(s)
}

func (s Schema) GoString() string {
	return s.String()
}

func (s *Schema) GetComment() *string {
	return s.Comment
}

func (s *Schema) GetFields() []*DataField {
	return s.Fields
}

func (s *Schema) GetOptions() map[string]*string {
	return s.Options
}

func (s *Schema) GetPartitionKeys() []*string {
	return s.PartitionKeys
}

func (s *Schema) GetPrimaryKeys() []*string {
	return s.PrimaryKeys
}

func (s *Schema) SetComment(v string) *Schema {
	s.Comment = &v
	return s
}

func (s *Schema) SetFields(v []*DataField) *Schema {
	s.Fields = v
	return s
}

func (s *Schema) SetOptions(v map[string]*string) *Schema {
	s.Options = v
	return s
}

func (s *Schema) SetPartitionKeys(v []*string) *Schema {
	s.PartitionKeys = v
	return s
}

func (s *Schema) SetPrimaryKeys(v []*string) *Schema {
	s.PrimaryKeys = v
	return s
}

func (s *Schema) Validate() error {
	return dara.Validate(s)
}
