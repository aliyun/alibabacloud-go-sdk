// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFieldRef interface {
	dara.Model
	String() string
	GoString() string
	SetIndex(v int32) *FieldRef
	GetIndex() *int32
	SetName(v string) *FieldRef
	GetName() *string
	SetType(v *FullDataType) *FieldRef
	GetType() *FullDataType
}

type FieldRef struct {
	Index *int32        `json:"index,omitempty" xml:"index,omitempty"`
	Name  *string       `json:"name,omitempty" xml:"name,omitempty"`
	Type  *FullDataType `json:"type,omitempty" xml:"type,omitempty"`
}

func (s FieldRef) String() string {
	return dara.Prettify(s)
}

func (s FieldRef) GoString() string {
	return s.String()
}

func (s *FieldRef) GetIndex() *int32 {
	return s.Index
}

func (s *FieldRef) GetName() *string {
	return s.Name
}

func (s *FieldRef) GetType() *FullDataType {
	return s.Type
}

func (s *FieldRef) SetIndex(v int32) *FieldRef {
	s.Index = &v
	return s
}

func (s *FieldRef) SetName(v string) *FieldRef {
	s.Name = &v
	return s
}

func (s *FieldRef) SetType(v *FullDataType) *FieldRef {
	s.Type = v
	return s
}

func (s *FieldRef) Validate() error {
	if s.Type != nil {
		if err := s.Type.Validate(); err != nil {
			return err
		}
	}
	return nil
}
