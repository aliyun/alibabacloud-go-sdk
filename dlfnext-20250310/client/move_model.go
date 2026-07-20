// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMove interface {
	dara.Model
	String() string
	GoString() string
	SetFieldName(v string) *Move
	GetFieldName() *string
	SetReferenceFieldName(v string) *Move
	GetReferenceFieldName() *string
	SetType(v string) *Move
	GetType() *string
}

type Move struct {
	FieldName          *string `json:"fieldName,omitempty" xml:"fieldName,omitempty"`
	ReferenceFieldName *string `json:"referenceFieldName,omitempty" xml:"referenceFieldName,omitempty"`
	Type               *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s Move) String() string {
	return dara.Prettify(s)
}

func (s Move) GoString() string {
	return s.String()
}

func (s *Move) GetFieldName() *string {
	return s.FieldName
}

func (s *Move) GetReferenceFieldName() *string {
	return s.ReferenceFieldName
}

func (s *Move) GetType() *string {
	return s.Type
}

func (s *Move) SetFieldName(v string) *Move {
	s.FieldName = &v
	return s
}

func (s *Move) SetReferenceFieldName(v string) *Move {
	s.ReferenceFieldName = &v
	return s
}

func (s *Move) SetType(v string) *Move {
	s.Type = &v
	return s
}

func (s *Move) Validate() error {
	return dara.Validate(s)
}
