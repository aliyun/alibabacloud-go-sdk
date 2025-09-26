// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCredentialOutput interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *UpdateCredentialOutput
	GetId() *string
	SetName(v string) *UpdateCredentialOutput
	GetName() *string
	SetType(v string) *UpdateCredentialOutput
	GetType() *string
	SetUpdatedAt(v string) *UpdateCredentialOutput
	GetUpdatedAt() *string
}

type UpdateCredentialOutput struct {
	Id        *string `json:"id,omitempty" xml:"id,omitempty"`
	Name      *string `json:"name,omitempty" xml:"name,omitempty"`
	Type      *string `json:"type,omitempty" xml:"type,omitempty"`
	UpdatedAt *string `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
}

func (s UpdateCredentialOutput) String() string {
	return dara.Prettify(s)
}

func (s UpdateCredentialOutput) GoString() string {
	return s.String()
}

func (s *UpdateCredentialOutput) GetId() *string {
	return s.Id
}

func (s *UpdateCredentialOutput) GetName() *string {
	return s.Name
}

func (s *UpdateCredentialOutput) GetType() *string {
	return s.Type
}

func (s *UpdateCredentialOutput) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *UpdateCredentialOutput) SetId(v string) *UpdateCredentialOutput {
	s.Id = &v
	return s
}

func (s *UpdateCredentialOutput) SetName(v string) *UpdateCredentialOutput {
	s.Name = &v
	return s
}

func (s *UpdateCredentialOutput) SetType(v string) *UpdateCredentialOutput {
	s.Type = &v
	return s
}

func (s *UpdateCredentialOutput) SetUpdatedAt(v string) *UpdateCredentialOutput {
	s.UpdatedAt = &v
	return s
}

func (s *UpdateCredentialOutput) Validate() error {
	return dara.Validate(s)
}
