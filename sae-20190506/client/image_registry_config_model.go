// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImageRegistryConfig interface {
	dara.Model
	String() string
	GoString() string
	SetAuthConfig(v *RegistryAuthenticationConfig) *ImageRegistryConfig
	GetAuthConfig() *RegistryAuthenticationConfig
	SetCertConfig(v *RegistryCertificateConfig) *ImageRegistryConfig
	GetCertConfig() *RegistryCertificateConfig
}

type ImageRegistryConfig struct {
	AuthConfig *RegistryAuthenticationConfig `json:"AuthConfig,omitempty" xml:"AuthConfig,omitempty"`
	CertConfig *RegistryCertificateConfig    `json:"CertConfig,omitempty" xml:"CertConfig,omitempty"`
}

func (s ImageRegistryConfig) String() string {
	return dara.Prettify(s)
}

func (s ImageRegistryConfig) GoString() string {
	return s.String()
}

func (s *ImageRegistryConfig) GetAuthConfig() *RegistryAuthenticationConfig {
	return s.AuthConfig
}

func (s *ImageRegistryConfig) GetCertConfig() *RegistryCertificateConfig {
	return s.CertConfig
}

func (s *ImageRegistryConfig) SetAuthConfig(v *RegistryAuthenticationConfig) *ImageRegistryConfig {
	s.AuthConfig = v
	return s
}

func (s *ImageRegistryConfig) SetCertConfig(v *RegistryCertificateConfig) *ImageRegistryConfig {
	s.CertConfig = v
	return s
}

func (s *ImageRegistryConfig) Validate() error {
	return dara.Validate(s)
}
