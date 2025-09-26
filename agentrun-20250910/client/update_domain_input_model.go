// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDomainInput interface {
	dara.Model
	String() string
	GoString() string
	SetCertIdentifier(v string) *UpdateDomainInput
	GetCertIdentifier() *string
	SetProtocol(v string) *UpdateDomainInput
	GetProtocol() *string
}

type UpdateDomainInput struct {
	CertIdentifier *string `json:"certIdentifier,omitempty" xml:"certIdentifier,omitempty"`
	Protocol       *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
}

func (s UpdateDomainInput) String() string {
	return dara.Prettify(s)
}

func (s UpdateDomainInput) GoString() string {
	return s.String()
}

func (s *UpdateDomainInput) GetCertIdentifier() *string {
	return s.CertIdentifier
}

func (s *UpdateDomainInput) GetProtocol() *string {
	return s.Protocol
}

func (s *UpdateDomainInput) SetCertIdentifier(v string) *UpdateDomainInput {
	s.CertIdentifier = &v
	return s
}

func (s *UpdateDomainInput) SetProtocol(v string) *UpdateDomainInput {
	s.Protocol = &v
	return s
}

func (s *UpdateDomainInput) Validate() error {
	return dara.Validate(s)
}
