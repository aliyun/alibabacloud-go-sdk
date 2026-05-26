// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDefinition interface {
	dara.Model
	String() string
	GoString() string
	SetCedar(v *DefinitionCedar) *Definition
	GetCedar() *DefinitionCedar
}

type Definition struct {
	Cedar *DefinitionCedar `json:"Cedar,omitempty" xml:"Cedar,omitempty" type:"Struct"`
}

func (s Definition) String() string {
	return dara.Prettify(s)
}

func (s Definition) GoString() string {
	return s.String()
}

func (s *Definition) GetCedar() *DefinitionCedar {
	return s.Cedar
}

func (s *Definition) SetCedar(v *DefinitionCedar) *Definition {
	s.Cedar = v
	return s
}

func (s *Definition) Validate() error {
	if s.Cedar != nil {
		if err := s.Cedar.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DefinitionCedar struct {
	Statement *string `json:"Statement,omitempty" xml:"Statement,omitempty"`
}

func (s DefinitionCedar) String() string {
	return dara.Prettify(s)
}

func (s DefinitionCedar) GoString() string {
	return s.String()
}

func (s *DefinitionCedar) GetStatement() *string {
	return s.Statement
}

func (s *DefinitionCedar) SetStatement(v string) *DefinitionCedar {
	s.Statement = &v
	return s
}

func (s *DefinitionCedar) Validate() error {
	return dara.Validate(s)
}
