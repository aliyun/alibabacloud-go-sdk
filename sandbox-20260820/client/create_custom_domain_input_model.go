// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCustomDomainInput interface {
	dara.Model
	String() string
	GoString() string
	SetCertConfig(v *CertConfig) *CreateCustomDomainInput
	GetCertConfig() *CertConfig
	SetDescription(v string) *CreateCustomDomainInput
	GetDescription() *string
	SetDomainName(v string) *CreateCustomDomainInput
	GetDomainName() *string
	SetTlsConfig(v *TLSConfig) *CreateCustomDomainInput
	GetTlsConfig() *TLSConfig
}

type CreateCustomDomainInput struct {
	CertConfig *CertConfig `json:"certConfig,omitempty" xml:"certConfig,omitempty"`
	// example:
	//
	// 沙箱预览环境入口域名
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// sandbox.example.com
	DomainName *string    `json:"domainName,omitempty" xml:"domainName,omitempty"`
	TlsConfig  *TLSConfig `json:"tlsConfig,omitempty" xml:"tlsConfig,omitempty"`
}

func (s CreateCustomDomainInput) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomDomainInput) GoString() string {
	return s.String()
}

func (s *CreateCustomDomainInput) GetCertConfig() *CertConfig {
	return s.CertConfig
}

func (s *CreateCustomDomainInput) GetDescription() *string {
	return s.Description
}

func (s *CreateCustomDomainInput) GetDomainName() *string {
	return s.DomainName
}

func (s *CreateCustomDomainInput) GetTlsConfig() *TLSConfig {
	return s.TlsConfig
}

func (s *CreateCustomDomainInput) SetCertConfig(v *CertConfig) *CreateCustomDomainInput {
	s.CertConfig = v
	return s
}

func (s *CreateCustomDomainInput) SetDescription(v string) *CreateCustomDomainInput {
	s.Description = &v
	return s
}

func (s *CreateCustomDomainInput) SetDomainName(v string) *CreateCustomDomainInput {
	s.DomainName = &v
	return s
}

func (s *CreateCustomDomainInput) SetTlsConfig(v *TLSConfig) *CreateCustomDomainInput {
	s.TlsConfig = v
	return s
}

func (s *CreateCustomDomainInput) Validate() error {
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
