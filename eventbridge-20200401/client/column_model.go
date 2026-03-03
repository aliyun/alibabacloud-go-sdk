// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iColumn interface {
	dara.Model
	String() string
	GoString() string
	SetIsNull(v bool) *Column
	GetIsNull() *bool
	SetName(v string) *Column
	GetName() *string
	SetType(v string) *Column
	GetType() *string
	SetValue(v string) *Column
	GetValue() *string
}

type Column struct {
	IsNull *bool   `json:"IsNull,omitempty" xml:"IsNull,omitempty"`
	Name   *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Type   *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Value  *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s Column) String() string {
	return dara.Prettify(s)
}

func (s Column) GoString() string {
	return s.String()
}

func (s *Column) GetIsNull() *bool {
	return s.IsNull
}

func (s *Column) GetName() *string {
	return s.Name
}

func (s *Column) GetType() *string {
	return s.Type
}

func (s *Column) GetValue() *string {
	return s.Value
}

func (s *Column) SetIsNull(v bool) *Column {
	s.IsNull = &v
	return s
}

func (s *Column) SetName(v string) *Column {
	s.Name = &v
	return s
}

func (s *Column) SetType(v string) *Column {
	s.Type = &v
	return s
}

func (s *Column) SetValue(v string) *Column {
	s.Value = &v
	return s
}

func (s *Column) Validate() error {
	return dara.Validate(s)
}
