// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCredentialOutput interface {
	dara.Model
	String() string
	GoString() string
	SetCreatedAt(v string) *CreateCredentialOutput
	GetCreatedAt() *string
	SetId(v string) *CreateCredentialOutput
	GetId() *string
	SetName(v string) *CreateCredentialOutput
	GetName() *string
	SetType(v string) *CreateCredentialOutput
	GetType() *string
}

type CreateCredentialOutput struct {
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	Id        *string `json:"id,omitempty" xml:"id,omitempty"`
	Name      *string `json:"name,omitempty" xml:"name,omitempty"`
	Type      *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateCredentialOutput) String() string {
	return dara.Prettify(s)
}

func (s CreateCredentialOutput) GoString() string {
	return s.String()
}

func (s *CreateCredentialOutput) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *CreateCredentialOutput) GetId() *string {
	return s.Id
}

func (s *CreateCredentialOutput) GetName() *string {
	return s.Name
}

func (s *CreateCredentialOutput) GetType() *string {
	return s.Type
}

func (s *CreateCredentialOutput) SetCreatedAt(v string) *CreateCredentialOutput {
	s.CreatedAt = &v
	return s
}

func (s *CreateCredentialOutput) SetId(v string) *CreateCredentialOutput {
	s.Id = &v
	return s
}

func (s *CreateCredentialOutput) SetName(v string) *CreateCredentialOutput {
	s.Name = &v
	return s
}

func (s *CreateCredentialOutput) SetType(v string) *CreateCredentialOutput {
	s.Type = &v
	return s
}

func (s *CreateCredentialOutput) Validate() error {
	return dara.Validate(s)
}
