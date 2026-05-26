// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetSAMLIdentityProviderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetSAMLIdentityProviderResponseBody
	GetRequestId() *string
	SetSSOIdentityProviderConfiguration(v *SetSAMLIdentityProviderResponseBodySSOIdentityProviderConfiguration) *SetSAMLIdentityProviderResponseBody
	GetSSOIdentityProviderConfiguration() *SetSAMLIdentityProviderResponseBodySSOIdentityProviderConfiguration
}

type SetSAMLIdentityProviderResponseBody struct {
	// example:
	//
	// AABD6E03-4B3A-5687-88FF-72232670ED0C
	RequestId                        *string                                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SSOIdentityProviderConfiguration *SetSAMLIdentityProviderResponseBodySSOIdentityProviderConfiguration `json:"SSOIdentityProviderConfiguration,omitempty" xml:"SSOIdentityProviderConfiguration,omitempty" type:"Struct"`
}

func (s SetSAMLIdentityProviderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetSAMLIdentityProviderResponseBody) GoString() string {
	return s.String()
}

func (s *SetSAMLIdentityProviderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetSAMLIdentityProviderResponseBody) GetSSOIdentityProviderConfiguration() *SetSAMLIdentityProviderResponseBodySSOIdentityProviderConfiguration {
	return s.SSOIdentityProviderConfiguration
}

func (s *SetSAMLIdentityProviderResponseBody) SetRequestId(v string) *SetSAMLIdentityProviderResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetSAMLIdentityProviderResponseBody) SetSSOIdentityProviderConfiguration(v *SetSAMLIdentityProviderResponseBodySSOIdentityProviderConfiguration) *SetSAMLIdentityProviderResponseBody {
	s.SSOIdentityProviderConfiguration = v
	return s
}

func (s *SetSAMLIdentityProviderResponseBody) Validate() error {
	if s.SSOIdentityProviderConfiguration != nil {
		if err := s.SSOIdentityProviderConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SetSAMLIdentityProviderResponseBodySSOIdentityProviderConfiguration struct {
	// example:
	//
	// https://idp.example.com/entity
	EntityId *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// example:
	//
	// Enabled
	JITProvisionStatus *string `json:"JITProvisionStatus,omitempty" xml:"JITProvisionStatus,omitempty"`
	// example:
	//
	// Enabled
	JITUpdateStatus *string `json:"JITUpdateStatus,omitempty" xml:"JITUpdateStatus,omitempty"`
	// example:
	//
	// https://idp.example.com/sso/saml
	LoginURL *string `json:"LoginURL,omitempty" xml:"LoginURL,omitempty"`
	// example:
	//
	// HTTP-Redirect
	SAMLBindingType *string `json:"SAMLBindingType,omitempty" xml:"SAMLBindingType,omitempty"`
	// example:
	//
	// Enabled
	SSOStatus        *string                                                                                `json:"SSOStatus,omitempty" xml:"SSOStatus,omitempty"`
	X509Certificates []*SetSAMLIdentityProviderResponseBodySSOIdentityProviderConfigurationX509Certificates `json:"X509Certificates,omitempty" xml:"X509Certificates,omitempty" type:"Repeated"`
}

func (s SetSAMLIdentityProviderResponseBodySSOIdentityProviderConfiguration) String() string {
	return dara.Prettify(s)
}

func (s SetSAMLIdentityProviderResponseBodySSOIdentityProviderConfiguration) GoString() string {
	return s.String()
}

func (s *SetSAMLIdentityProviderResponseBodySSOIdentityProviderConfiguration) GetEntityId() *string {
	return s.EntityId
}

func (s *SetSAMLIdentityProviderResponseBodySSOIdentityProviderConfiguration) GetJITProvisionStatus() *string {
	return s.JITProvisionStatus
}

func (s *SetSAMLIdentityProviderResponseBodySSOIdentityProviderConfiguration) GetJITUpdateStatus() *string {
	return s.JITUpdateStatus
}

func (s *SetSAMLIdentityProviderResponseBodySSOIdentityProviderConfiguration) GetLoginURL() *string {
	return s.LoginURL
}

func (s *SetSAMLIdentityProviderResponseBodySSOIdentityProviderConfiguration) GetSAMLBindingType() *string {
	return s.SAMLBindingType
}

func (s *SetSAMLIdentityProviderResponseBodySSOIdentityProviderConfiguration) GetSSOStatus() *string {
	return s.SSOStatus
}

func (s *SetSAMLIdentityProviderResponseBodySSOIdentityProviderConfiguration) GetX509Certificates() []*SetSAMLIdentityProviderResponseBodySSOIdentityProviderConfigurationX509Certificates {
	return s.X509Certificates
}

func (s *SetSAMLIdentityProviderResponseBodySSOIdentityProviderConfiguration) SetEntityId(v string) *SetSAMLIdentityProviderResponseBodySSOIdentityProviderConfiguration {
	s.EntityId = &v
	return s
}

func (s *SetSAMLIdentityProviderResponseBodySSOIdentityProviderConfiguration) SetJITProvisionStatus(v string) *SetSAMLIdentityProviderResponseBodySSOIdentityProviderConfiguration {
	s.JITProvisionStatus = &v
	return s
}

func (s *SetSAMLIdentityProviderResponseBodySSOIdentityProviderConfiguration) SetJITUpdateStatus(v string) *SetSAMLIdentityProviderResponseBodySSOIdentityProviderConfiguration {
	s.JITUpdateStatus = &v
	return s
}

func (s *SetSAMLIdentityProviderResponseBodySSOIdentityProviderConfiguration) SetLoginURL(v string) *SetSAMLIdentityProviderResponseBodySSOIdentityProviderConfiguration {
	s.LoginURL = &v
	return s
}

func (s *SetSAMLIdentityProviderResponseBodySSOIdentityProviderConfiguration) SetSAMLBindingType(v string) *SetSAMLIdentityProviderResponseBodySSOIdentityProviderConfiguration {
	s.SAMLBindingType = &v
	return s
}

func (s *SetSAMLIdentityProviderResponseBodySSOIdentityProviderConfiguration) SetSSOStatus(v string) *SetSAMLIdentityProviderResponseBodySSOIdentityProviderConfiguration {
	s.SSOStatus = &v
	return s
}

func (s *SetSAMLIdentityProviderResponseBodySSOIdentityProviderConfiguration) SetX509Certificates(v []*SetSAMLIdentityProviderResponseBodySSOIdentityProviderConfigurationX509Certificates) *SetSAMLIdentityProviderResponseBodySSOIdentityProviderConfiguration {
	s.X509Certificates = v
	return s
}

func (s *SetSAMLIdentityProviderResponseBodySSOIdentityProviderConfiguration) Validate() error {
	if s.X509Certificates != nil {
		for _, item := range s.X509Certificates {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SetSAMLIdentityProviderResponseBodySSOIdentityProviderConfigurationX509Certificates struct {
	// example:
	//
	// cert-xxxxxxxxxxxxxxxxxxxx
	CertificateId *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
	// example:
	//
	// -----BEGIN CERTIFICATE-----
	//
	// MIIDdz...
	//
	// -----END CERTIFICATE-----
	X509Certificate *string `json:"X509Certificate,omitempty" xml:"X509Certificate,omitempty"`
}

func (s SetSAMLIdentityProviderResponseBodySSOIdentityProviderConfigurationX509Certificates) String() string {
	return dara.Prettify(s)
}

func (s SetSAMLIdentityProviderResponseBodySSOIdentityProviderConfigurationX509Certificates) GoString() string {
	return s.String()
}

func (s *SetSAMLIdentityProviderResponseBodySSOIdentityProviderConfigurationX509Certificates) GetCertificateId() *string {
	return s.CertificateId
}

func (s *SetSAMLIdentityProviderResponseBodySSOIdentityProviderConfigurationX509Certificates) GetX509Certificate() *string {
	return s.X509Certificate
}

func (s *SetSAMLIdentityProviderResponseBodySSOIdentityProviderConfigurationX509Certificates) SetCertificateId(v string) *SetSAMLIdentityProviderResponseBodySSOIdentityProviderConfigurationX509Certificates {
	s.CertificateId = &v
	return s
}

func (s *SetSAMLIdentityProviderResponseBodySSOIdentityProviderConfigurationX509Certificates) SetX509Certificate(v string) *SetSAMLIdentityProviderResponseBodySSOIdentityProviderConfigurationX509Certificates {
	s.X509Certificate = &v
	return s
}

func (s *SetSAMLIdentityProviderResponseBodySSOIdentityProviderConfigurationX509Certificates) Validate() error {
	return dara.Validate(s)
}
