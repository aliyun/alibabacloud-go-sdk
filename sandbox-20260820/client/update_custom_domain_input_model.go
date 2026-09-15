// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCustomDomainInput interface {
	dara.Model
	String() string
	GoString() string
	SetCertConfig(v *CertConfig) *UpdateCustomDomainInput
	GetCertConfig() *CertConfig
	SetDescription(v string) *UpdateCustomDomainInput
	GetDescription() *string
	SetTlsConfig(v *TLSConfig) *UpdateCustomDomainInput
	GetTlsConfig() *TLSConfig
}

type UpdateCustomDomainInput struct {
	CertConfig  *CertConfig `json:"certConfig,omitempty" xml:"certConfig,omitempty"`
	Description *string     `json:"description,omitempty" xml:"description,omitempty"`
	TlsConfig   *TLSConfig  `json:"tlsConfig,omitempty" xml:"tlsConfig,omitempty"`
}

func (s UpdateCustomDomainInput) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomDomainInput) GoString() string {
	return s.String()
}

func (s *UpdateCustomDomainInput) GetCertConfig() *CertConfig {
	return s.CertConfig
}

func (s *UpdateCustomDomainInput) GetDescription() *string {
	return s.Description
}

func (s *UpdateCustomDomainInput) GetTlsConfig() *TLSConfig {
	return s.TlsConfig
}

func (s *UpdateCustomDomainInput) SetCertConfig(v *CertConfig) *UpdateCustomDomainInput {
	s.CertConfig = v
	return s
}

func (s *UpdateCustomDomainInput) SetDescription(v string) *UpdateCustomDomainInput {
	s.Description = &v
	return s
}

func (s *UpdateCustomDomainInput) SetTlsConfig(v *TLSConfig) *UpdateCustomDomainInput {
	s.TlsConfig = v
	return s
}

func (s *UpdateCustomDomainInput) Validate() error {
	if s.CertConfig != nil {
		if err := s.CertConfig.Validate(); err != nil {
			return err
		}
	}
	if s.TlsConfig != nil {
		if err := s.TlsConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}
