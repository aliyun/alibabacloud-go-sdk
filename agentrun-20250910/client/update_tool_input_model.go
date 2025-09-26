// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateToolInput interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateToolInput
	GetDescription() *string
	SetName(v string) *UpdateToolInput
	GetName() *string
	SetSchema(v string) *UpdateToolInput
	GetSchema() *string
}

type UpdateToolInput struct {
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	Name        *string `json:"name,omitempty" xml:"name,omitempty"`
	Schema      *string `json:"schema,omitempty" xml:"schema,omitempty"`
}

func (s UpdateToolInput) String() string {
	return dara.Prettify(s)
}

func (s UpdateToolInput) GoString() string {
	return s.String()
}

func (s *UpdateToolInput) GetDescription() *string {
	return s.Description
}

func (s *UpdateToolInput) GetName() *string {
	return s.Name
}

func (s *UpdateToolInput) GetSchema() *string {
	return s.Schema
}

func (s *UpdateToolInput) SetDescription(v string) *UpdateToolInput {
	s.Description = &v
	return s
}

func (s *UpdateToolInput) SetName(v string) *UpdateToolInput {
	s.Name = &v
	return s
}

func (s *UpdateToolInput) SetSchema(v string) *UpdateToolInput {
	s.Schema = &v
	return s
}

func (s *UpdateToolInput) Validate() error {
	return dara.Validate(s)
}
