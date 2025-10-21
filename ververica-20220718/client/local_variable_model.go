// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLocalVariable interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *LocalVariable
	GetName() *string
	SetValue(v string) *LocalVariable
	GetValue() *string
}

type LocalVariable struct {
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// datagen
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s LocalVariable) String() string {
	return dara.Prettify(s)
}

func (s LocalVariable) GoString() string {
	return s.String()
}

func (s *LocalVariable) GetName() *string {
	return s.Name
}

func (s *LocalVariable) GetValue() *string {
	return s.Value
}

func (s *LocalVariable) SetName(v string) *LocalVariable {
	s.Name = &v
	return s
}

func (s *LocalVariable) SetValue(v string) *LocalVariable {
	s.Value = &v
	return s
}

func (s *LocalVariable) Validate() error {
	return dara.Validate(s)
}
