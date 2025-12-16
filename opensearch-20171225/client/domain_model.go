// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDomain interface {
	dara.Model
	String() string
	GoString() string
	SetCategory(v string) *Domain
	GetCategory() *string
	SetFunctions(v map[string][]*string) *Domain
	GetFunctions() map[string][]*string
	SetName(v string) *Domain
	GetName() *string
}

type Domain struct {
	Category  *string              `json:"category,omitempty" xml:"category,omitempty"`
	Functions map[string][]*string `json:"functions,omitempty" xml:"functions,omitempty"`
	Name      *string              `json:"name,omitempty" xml:"name,omitempty"`
}

func (s Domain) String() string {
	return dara.Prettify(s)
}

func (s Domain) GoString() string {
	return s.String()
}

func (s *Domain) GetCategory() *string {
	return s.Category
}

func (s *Domain) GetFunctions() map[string][]*string {
	return s.Functions
}

func (s *Domain) GetName() *string {
	return s.Name
}

func (s *Domain) SetCategory(v string) *Domain {
	s.Category = &v
	return s
}

func (s *Domain) SetFunctions(v map[string][]*string) *Domain {
	s.Functions = v
	return s
}

func (s *Domain) SetName(v string) *Domain {
	s.Name = &v
	return s
}

func (s *Domain) Validate() error {
	return dara.Validate(s)
}
