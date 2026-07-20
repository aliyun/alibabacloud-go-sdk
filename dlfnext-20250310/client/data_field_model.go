// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDataField interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *DataField
	GetDescription() *string
	SetId(v int32) *DataField
	GetId() *int32
	SetName(v string) *DataField
	GetName() *string
	SetType(v *FullDataType) *DataField
	GetType() *FullDataType
}

type DataField struct {
	Description *string       `json:"description,omitempty" xml:"description,omitempty"`
	Id          *int32        `json:"id,omitempty" xml:"id,omitempty"`
	Name        *string       `json:"name,omitempty" xml:"name,omitempty"`
	Type        *FullDataType `json:"type,omitempty" xml:"type,omitempty"`
}

func (s DataField) String() string {
	return dara.Prettify(s)
}

func (s DataField) GoString() string {
	return s.String()
}

func (s *DataField) GetDescription() *string {
	return s.Description
}

func (s *DataField) GetId() *int32 {
	return s.Id
}

func (s *DataField) GetName() *string {
	return s.Name
}

func (s *DataField) GetType() *FullDataType {
	return s.Type
}

func (s *DataField) SetDescription(v string) *DataField {
	s.Description = &v
	return s
}

func (s *DataField) SetId(v int32) *DataField {
	s.Id = &v
	return s
}

func (s *DataField) SetName(v string) *DataField {
	s.Name = &v
	return s
}

func (s *DataField) SetType(v *FullDataType) *DataField {
	s.Type = v
	return s
}

func (s *DataField) Validate() error {
	if s.Type != nil {
		if err := s.Type.Validate(); err != nil {
			return err
		}
	}
	return nil
}
