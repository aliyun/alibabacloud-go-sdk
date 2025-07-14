// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCustomDomainInput interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationName(v string) *CreateCustomDomainInput
	GetApplicationName() *string
	SetCertConfig(v *CertConfig) *CreateCustomDomainInput
	GetCertConfig() *CertConfig
	SetDomainName(v string) *CreateCustomDomainInput
	GetDomainName() *string
	SetKeepFullPath(v bool) *CreateCustomDomainInput
	GetKeepFullPath() *bool
	SetNamespaceID(v string) *CreateCustomDomainInput
	GetNamespaceID() *string
	SetProtocol(v string) *CreateCustomDomainInput
	GetProtocol() *string
	SetTlsConfig(v *TLSConfig) *CreateCustomDomainInput
	GetTlsConfig() *TLSConfig
	SetWafConfig(v *WAFConfig) *CreateCustomDomainInput
	GetWafConfig() *WAFConfig
}

type CreateCustomDomainInput struct {
	ApplicationName *string     `json:"applicationName,omitempty" xml:"applicationName,omitempty"`
	CertConfig      *CertConfig `json:"certConfig,omitempty" xml:"certConfig,omitempty"`
	DomainName      *string     `json:"domainName,omitempty" xml:"domainName,omitempty"`
	KeepFullPath    *bool       `json:"keepFullPath,omitempty" xml:"keepFullPath,omitempty"`
	NamespaceID     *string     `json:"namespaceID,omitempty" xml:"namespaceID,omitempty"`
	Protocol        *string     `json:"protocol,omitempty" xml:"protocol,omitempty"`
	TlsConfig       *TLSConfig  `json:"tlsConfig,omitempty" xml:"tlsConfig,omitempty"`
	WafConfig       *WAFConfig  `json:"wafConfig,omitempty" xml:"wafConfig,omitempty"`
}

func (s CreateCustomDomainInput) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomDomainInput) GoString() string {
	return s.String()
}

func (s *CreateCustomDomainInput) GetApplicationName() *string {
	return s.ApplicationName
}

func (s *CreateCustomDomainInput) GetCertConfig() *CertConfig {
	return s.CertConfig
}

func (s *CreateCustomDomainInput) GetDomainName() *string {
	return s.DomainName
}

func (s *CreateCustomDomainInput) GetKeepFullPath() *bool {
	return s.KeepFullPath
}

func (s *CreateCustomDomainInput) GetNamespaceID() *string {
	return s.NamespaceID
}

func (s *CreateCustomDomainInput) GetProtocol() *string {
	return s.Protocol
}

func (s *CreateCustomDomainInput) GetTlsConfig() *TLSConfig {
	return s.TlsConfig
}

func (s *CreateCustomDomainInput) GetWafConfig() *WAFConfig {
	return s.WafConfig
}

func (s *CreateCustomDomainInput) SetApplicationName(v string) *CreateCustomDomainInput {
	s.ApplicationName = &v
	return s
}

func (s *CreateCustomDomainInput) SetCertConfig(v *CertConfig) *CreateCustomDomainInput {
	s.CertConfig = v
	return s
}

func (s *CreateCustomDomainInput) SetDomainName(v string) *CreateCustomDomainInput {
	s.DomainName = &v
	return s
}

func (s *CreateCustomDomainInput) SetKeepFullPath(v bool) *CreateCustomDomainInput {
	s.KeepFullPath = &v
	return s
}

func (s *CreateCustomDomainInput) SetNamespaceID(v string) *CreateCustomDomainInput {
	s.NamespaceID = &v
	return s
}

func (s *CreateCustomDomainInput) SetProtocol(v string) *CreateCustomDomainInput {
	s.Protocol = &v
	return s
}

func (s *CreateCustomDomainInput) SetTlsConfig(v *TLSConfig) *CreateCustomDomainInput {
	s.TlsConfig = v
	return s
}

func (s *CreateCustomDomainInput) SetWafConfig(v *WAFConfig) *CreateCustomDomainInput {
	s.WafConfig = v
	return s
}

func (s *CreateCustomDomainInput) Validate() error {
	return dara.Validate(s)
}
