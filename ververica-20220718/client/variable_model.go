// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVariable interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *Variable
	GetDescription() *string
	SetKind(v string) *Variable
	GetKind() *string
	SetName(v string) *Variable
	GetName() *string
	SetValue(v string) *Variable
	GetValue() *string
}

type Variable struct {
	// example:
	//
	// This is a variable description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Plain
	Kind *string `json:"kind,omitempty" xml:"kind,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// variableName
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// variableValue
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s Variable) String() string {
	return dara.Prettify(s)
}

func (s Variable) GoString() string {
	return s.String()
}

func (s *Variable) GetDescription() *string {
	return s.Description
}

func (s *Variable) GetKind() *string {
	return s.Kind
}

func (s *Variable) GetName() *string {
	return s.Name
}

func (s *Variable) GetValue() *string {
	return s.Value
}

func (s *Variable) SetDescription(v string) *Variable {
	s.Description = &v
	return s
}

func (s *Variable) SetKind(v string) *Variable {
	s.Kind = &v
	return s
}

func (s *Variable) SetName(v string) *Variable {
	s.Name = &v
	return s
}

func (s *Variable) SetValue(v string) *Variable {
	s.Value = &v
	return s
}

func (s *Variable) Validate() error {
	return dara.Validate(s)
}
