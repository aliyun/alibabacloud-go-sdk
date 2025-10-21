// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSchema interface {
	dara.Model
	String() string
	GoString() string
	SetColumns(v []*TableColumn) *Schema
	GetColumns() []*TableColumn
	SetPrimaryKey(v *PrimaryKey) *Schema
	GetPrimaryKey() *PrimaryKey
	SetWatermarkSpecs(v []*WatermarkSpec) *Schema
	GetWatermarkSpecs() []*WatermarkSpec
}

type Schema struct {
	Columns        []*TableColumn   `json:"columns,omitempty" xml:"columns,omitempty" type:"Repeated"`
	PrimaryKey     *PrimaryKey      `json:"primaryKey,omitempty" xml:"primaryKey,omitempty"`
	WatermarkSpecs []*WatermarkSpec `json:"watermarkSpecs,omitempty" xml:"watermarkSpecs,omitempty" type:"Repeated"`
}

func (s Schema) String() string {
	return dara.Prettify(s)
}

func (s Schema) GoString() string {
	return s.String()
}

func (s *Schema) GetColumns() []*TableColumn {
	return s.Columns
}

func (s *Schema) GetPrimaryKey() *PrimaryKey {
	return s.PrimaryKey
}

func (s *Schema) GetWatermarkSpecs() []*WatermarkSpec {
	return s.WatermarkSpecs
}

func (s *Schema) SetColumns(v []*TableColumn) *Schema {
	s.Columns = v
	return s
}

func (s *Schema) SetPrimaryKey(v *PrimaryKey) *Schema {
	s.PrimaryKey = v
	return s
}

func (s *Schema) SetWatermarkSpecs(v []*WatermarkSpec) *Schema {
	s.WatermarkSpecs = v
	return s
}

func (s *Schema) Validate() error {
	if s.Columns != nil {
		for _, item := range s.Columns {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.PrimaryKey != nil {
		if err := s.PrimaryKey.Validate(); err != nil {
			return err
		}
	}
	if s.WatermarkSpecs != nil {
		for _, item := range s.WatermarkSpecs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
