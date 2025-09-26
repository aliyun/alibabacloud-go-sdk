// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDomainInput interface {
	dara.Model
	String() string
	GoString() string
	SetCertIdentifier(v string) *CreateDomainInput
	GetCertIdentifier() *string
	SetName(v string) *CreateDomainInput
	GetName() *string
	SetProtocol(v string) *CreateDomainInput
	GetProtocol() *string
}

type CreateDomainInput struct {
	CertIdentifier *string `json:"certIdentifier,omitempty" xml:"certIdentifier,omitempty"`
	Name           *string `json:"name,omitempty" xml:"name,omitempty"`
	Protocol       *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
}

func (s CreateDomainInput) String() string {
	return dara.Prettify(s)
}

func (s CreateDomainInput) GoString() string {
	return s.String()
}

func (s *CreateDomainInput) GetCertIdentifier() *string {
	return s.CertIdentifier
}

func (s *CreateDomainInput) GetName() *string {
	return s.Name
}

func (s *CreateDomainInput) GetProtocol() *string {
	return s.Protocol
}

func (s *CreateDomainInput) SetCertIdentifier(v string) *CreateDomainInput {
	s.CertIdentifier = &v
	return s
}

func (s *CreateDomainInput) SetName(v string) *CreateDomainInput {
	s.Name = &v
	return s
}

func (s *CreateDomainInput) SetProtocol(v string) *CreateDomainInput {
	s.Protocol = &v
	return s
}

func (s *CreateDomainInput) Validate() error {
	return dara.Validate(s)
}
