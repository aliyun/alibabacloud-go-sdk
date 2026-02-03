// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTransform interface {
	dara.Model
	String() string
	GoString() string
	SetCastType(v *FullDataType) *Transform
	GetCastType() *FullDataType
	SetFieldRef(v *FieldRef) *Transform
	GetFieldRef() *FieldRef
	SetInputs(v []*TransformInput) *Transform
	GetInputs() []*TransformInput
	SetName(v string) *Transform
	GetName() *string
}

type Transform struct {
	CastType *FullDataType     `json:"castType,omitempty" xml:"castType,omitempty"`
	FieldRef *FieldRef         `json:"fieldRef,omitempty" xml:"fieldRef,omitempty"`
	Inputs   []*TransformInput `json:"inputs,omitempty" xml:"inputs,omitempty" type:"Repeated"`
	Name     *string           `json:"name,omitempty" xml:"name,omitempty"`
}

func (s Transform) String() string {
	return dara.Prettify(s)
}

func (s Transform) GoString() string {
	return s.String()
}

func (s *Transform) GetCastType() *FullDataType {
	return s.CastType
}

func (s *Transform) GetFieldRef() *FieldRef {
	return s.FieldRef
}

func (s *Transform) GetInputs() []*TransformInput {
	return s.Inputs
}

func (s *Transform) GetName() *string {
	return s.Name
}

func (s *Transform) SetCastType(v *FullDataType) *Transform {
	s.CastType = v
	return s
}

func (s *Transform) SetFieldRef(v *FieldRef) *Transform {
	s.FieldRef = v
	return s
}

func (s *Transform) SetInputs(v []*TransformInput) *Transform {
	s.Inputs = v
	return s
}

func (s *Transform) SetName(v string) *Transform {
	s.Name = &v
	return s
}

func (s *Transform) Validate() error {
	if s.CastType != nil {
		if err := s.CastType.Validate(); err != nil {
			return err
		}
	}
	if s.FieldRef != nil {
		if err := s.FieldRef.Validate(); err != nil {
			return err
		}
	}
	if s.Inputs != nil {
		for _, item := range s.Inputs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
