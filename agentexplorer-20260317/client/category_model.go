// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCategory interface {
	dara.Model
	String() string
	GoString() string
	SetChildren(v []*CategoryChildren) *Category
	GetChildren() []*CategoryChildren
	SetCode(v string) *Category
	GetCode() *string
	SetName(v string) *Category
	GetName() *string
}

type Category struct {
	Children []*CategoryChildren `json:"children,omitempty" xml:"children,omitempty" type:"Repeated"`
	// example:
	//
	// compute
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s Category) String() string {
	return dara.Prettify(s)
}

func (s Category) GoString() string {
	return s.String()
}

func (s *Category) GetChildren() []*CategoryChildren {
	return s.Children
}

func (s *Category) GetCode() *string {
	return s.Code
}

func (s *Category) GetName() *string {
	return s.Name
}

func (s *Category) SetChildren(v []*CategoryChildren) *Category {
	s.Children = v
	return s
}

func (s *Category) SetCode(v string) *Category {
	s.Code = &v
	return s
}

func (s *Category) SetName(v string) *Category {
	s.Name = &v
	return s
}

func (s *Category) Validate() error {
	if s.Children != nil {
		for _, item := range s.Children {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CategoryChildren struct {
	// example:
	//
	// ecs
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s CategoryChildren) String() string {
	return dara.Prettify(s)
}

func (s CategoryChildren) GoString() string {
	return s.String()
}

func (s *CategoryChildren) GetCode() *string {
	return s.Code
}

func (s *CategoryChildren) GetName() *string {
	return s.Name
}

func (s *CategoryChildren) SetCode(v string) *CategoryChildren {
	s.Code = &v
	return s
}

func (s *CategoryChildren) SetName(v string) *CategoryChildren {
	s.Name = &v
	return s
}

func (s *CategoryChildren) Validate() error {
	return dara.Validate(s)
}
