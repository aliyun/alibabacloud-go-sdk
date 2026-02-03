// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPredicate interface {
	dara.Model
	String() string
	GoString() string
	SetChildren(v []*Predicate) *Predicate
	GetChildren() []*Predicate
	SetFunction(v string) *Predicate
	GetFunction() *string
	SetKind(v string) *Predicate
	GetKind() *string
	SetLiterals(v []interface{}) *Predicate
	GetLiterals() []interface{}
	SetTransform(v *Transform) *Predicate
	GetTransform() *Transform
}

type Predicate struct {
	Children  []*Predicate  `json:"children,omitempty" xml:"children,omitempty" type:"Repeated"`
	Function  *string       `json:"function,omitempty" xml:"function,omitempty"`
	Kind      *string       `json:"kind,omitempty" xml:"kind,omitempty"`
	Literals  []interface{} `json:"literals,omitempty" xml:"literals,omitempty" type:"Repeated"`
	Transform *Transform    `json:"transform,omitempty" xml:"transform,omitempty"`
}

func (s Predicate) String() string {
	return dara.Prettify(s)
}

func (s Predicate) GoString() string {
	return s.String()
}

func (s *Predicate) GetChildren() []*Predicate {
	return s.Children
}

func (s *Predicate) GetFunction() *string {
	return s.Function
}

func (s *Predicate) GetKind() *string {
	return s.Kind
}

func (s *Predicate) GetLiterals() []interface{} {
	return s.Literals
}

func (s *Predicate) GetTransform() *Transform {
	return s.Transform
}

func (s *Predicate) SetChildren(v []*Predicate) *Predicate {
	s.Children = v
	return s
}

func (s *Predicate) SetFunction(v string) *Predicate {
	s.Function = &v
	return s
}

func (s *Predicate) SetKind(v string) *Predicate {
	s.Kind = &v
	return s
}

func (s *Predicate) SetLiterals(v []interface{}) *Predicate {
	s.Literals = v
	return s
}

func (s *Predicate) SetTransform(v *Transform) *Predicate {
	s.Transform = v
	return s
}

func (s *Predicate) Validate() error {
	if s.Children != nil {
		for _, item := range s.Children {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Transform != nil {
		if err := s.Transform.Validate(); err != nil {
			return err
		}
	}
	return nil
}
