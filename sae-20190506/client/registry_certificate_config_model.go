// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRegistryCertificateConfig interface {
	dara.Model
	String() string
	GoString() string
	SetCertBase64(v string) *RegistryCertificateConfig
	GetCertBase64() *string
	SetInsecure(v bool) *RegistryCertificateConfig
	GetInsecure() *bool
}

type RegistryCertificateConfig struct {
	// example:
	//
	// LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCmZha2VDZXJ0aWZpY2F0ZQotLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0t
	CertBase64 *string `json:"CertBase64,omitempty" xml:"CertBase64,omitempty"`
	// example:
	//
	// true
	Insecure *bool `json:"Insecure,omitempty" xml:"Insecure,omitempty"`
}

func (s RegistryCertificateConfig) String() string {
	return dara.Prettify(s)
}

func (s RegistryCertificateConfig) GoString() string {
	return s.String()
}

func (s *RegistryCertificateConfig) GetCertBase64() *string {
	return s.CertBase64
}

func (s *RegistryCertificateConfig) GetInsecure() *bool {
	return s.Insecure
}

func (s *RegistryCertificateConfig) SetCertBase64(v string) *RegistryCertificateConfig {
	s.CertBase64 = &v
	return s
}

func (s *RegistryCertificateConfig) SetInsecure(v bool) *RegistryCertificateConfig {
	s.Insecure = &v
	return s
}

func (s *RegistryCertificateConfig) Validate() error {
	return dara.Validate(s)
}
