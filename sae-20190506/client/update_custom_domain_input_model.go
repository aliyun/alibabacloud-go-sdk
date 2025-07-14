// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCustomDomainInput interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationID(v string) *UpdateCustomDomainInput
	GetApplicationID() *string
	SetApplicationName(v string) *UpdateCustomDomainInput
	GetApplicationName() *string
	SetCertConfig(v *CertConfig) *UpdateCustomDomainInput
	GetCertConfig() *CertConfig
	SetKeepFullPath(v bool) *UpdateCustomDomainInput
	GetKeepFullPath() *bool
	SetNamespaceID(v string) *UpdateCustomDomainInput
	GetNamespaceID() *string
	SetProtocol(v string) *UpdateCustomDomainInput
	GetProtocol() *string
	SetTlsConfig(v *TLSConfig) *UpdateCustomDomainInput
	GetTlsConfig() *TLSConfig
	SetWafConfig(v *WAFConfig) *UpdateCustomDomainInput
	GetWafConfig() *WAFConfig
}

type UpdateCustomDomainInput struct {
	ApplicationID   *string     `json:"applicationID,omitempty" xml:"applicationID,omitempty"`
	ApplicationName *string     `json:"applicationName,omitempty" xml:"applicationName,omitempty"`
	CertConfig      *CertConfig `json:"certConfig,omitempty" xml:"certConfig,omitempty"`
	KeepFullPath    *bool       `json:"keepFullPath,omitempty" xml:"keepFullPath,omitempty"`
	NamespaceID     *string     `json:"namespaceID,omitempty" xml:"namespaceID,omitempty"`
	Protocol        *string     `json:"protocol,omitempty" xml:"protocol,omitempty"`
	TlsConfig       *TLSConfig  `json:"tlsConfig,omitempty" xml:"tlsConfig,omitempty"`
	WafConfig       *WAFConfig  `json:"wafConfig,omitempty" xml:"wafConfig,omitempty"`
}

func (s UpdateCustomDomainInput) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomDomainInput) GoString() string {
	return s.String()
}

func (s *UpdateCustomDomainInput) GetApplicationID() *string {
	return s.ApplicationID
}

func (s *UpdateCustomDomainInput) GetApplicationName() *string {
	return s.ApplicationName
}

func (s *UpdateCustomDomainInput) GetCertConfig() *CertConfig {
	return s.CertConfig
}

func (s *UpdateCustomDomainInput) GetKeepFullPath() *bool {
	return s.KeepFullPath
}

func (s *UpdateCustomDomainInput) GetNamespaceID() *string {
	return s.NamespaceID
}

func (s *UpdateCustomDomainInput) GetProtocol() *string {
	return s.Protocol
}

func (s *UpdateCustomDomainInput) GetTlsConfig() *TLSConfig {
	return s.TlsConfig
}

func (s *UpdateCustomDomainInput) GetWafConfig() *WAFConfig {
	return s.WafConfig
}

func (s *UpdateCustomDomainInput) SetApplicationID(v string) *UpdateCustomDomainInput {
	s.ApplicationID = &v
	return s
}

func (s *UpdateCustomDomainInput) SetApplicationName(v string) *UpdateCustomDomainInput {
	s.ApplicationName = &v
	return s
}

func (s *UpdateCustomDomainInput) SetCertConfig(v *CertConfig) *UpdateCustomDomainInput {
	s.CertConfig = v
	return s
}

func (s *UpdateCustomDomainInput) SetKeepFullPath(v bool) *UpdateCustomDomainInput {
	s.KeepFullPath = &v
	return s
}

func (s *UpdateCustomDomainInput) SetNamespaceID(v string) *UpdateCustomDomainInput {
	s.NamespaceID = &v
	return s
}

func (s *UpdateCustomDomainInput) SetProtocol(v string) *UpdateCustomDomainInput {
	s.Protocol = &v
	return s
}

func (s *UpdateCustomDomainInput) SetTlsConfig(v *TLSConfig) *UpdateCustomDomainInput {
	s.TlsConfig = v
	return s
}

func (s *UpdateCustomDomainInput) SetWafConfig(v *WAFConfig) *UpdateCustomDomainInput {
	s.WafConfig = v
	return s
}

func (s *UpdateCustomDomainInput) Validate() error {
	return dara.Validate(s)
}
