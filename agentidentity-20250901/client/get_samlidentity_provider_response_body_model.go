// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSAMLIdentityProviderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetSAMLIdentityProviderResponseBody
	GetRequestId() *string
	SetSSOIdentityProviderConfiguration(v *GetSAMLIdentityProviderResponseBodySSOIdentityProviderConfiguration) *GetSAMLIdentityProviderResponseBody
	GetSSOIdentityProviderConfiguration() *GetSAMLIdentityProviderResponseBodySSOIdentityProviderConfiguration
}

type GetSAMLIdentityProviderResponseBody struct {
	// example:
	//
	// AABD6E03-4B3A-5687-88FF-72232670ED0C
	RequestId                        *string                                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SSOIdentityProviderConfiguration *GetSAMLIdentityProviderResponseBodySSOIdentityProviderConfiguration `json:"SSOIdentityProviderConfiguration,omitempty" xml:"SSOIdentityProviderConfiguration,omitempty" type:"Struct"`
}

func (s GetSAMLIdentityProviderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSAMLIdentityProviderResponseBody) GoString() string {
	return s.String()
}

func (s *GetSAMLIdentityProviderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSAMLIdentityProviderResponseBody) GetSSOIdentityProviderConfiguration() *GetSAMLIdentityProviderResponseBodySSOIdentityProviderConfiguration {
	return s.SSOIdentityProviderConfiguration
}

func (s *GetSAMLIdentityProviderResponseBody) SetRequestId(v string) *GetSAMLIdentityProviderResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSAMLIdentityProviderResponseBody) SetSSOIdentityProviderConfiguration(v *GetSAMLIdentityProviderResponseBodySSOIdentityProviderConfiguration) *GetSAMLIdentityProviderResponseBody {
	s.SSOIdentityProviderConfiguration = v
	return s
}

func (s *GetSAMLIdentityProviderResponseBody) Validate() error {
	if s.SSOIdentityProviderConfiguration != nil {
		if err := s.SSOIdentityProviderConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetSAMLIdentityProviderResponseBodySSOIdentityProviderConfiguration struct {
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
	X509Certificates []*GetSAMLIdentityProviderResponseBodySSOIdentityProviderConfigurationX509Certificates `json:"X509Certificates,omitempty" xml:"X509Certificates,omitempty" type:"Repeated"`
}

func (s GetSAMLIdentityProviderResponseBodySSOIdentityProviderConfiguration) String() string {
	return dara.Prettify(s)
}

func (s GetSAMLIdentityProviderResponseBodySSOIdentityProviderConfiguration) GoString() string {
	return s.String()
}

func (s *GetSAMLIdentityProviderResponseBodySSOIdentityProviderConfiguration) GetEntityId() *string {
	return s.EntityId
}

func (s *GetSAMLIdentityProviderResponseBodySSOIdentityProviderConfiguration) GetJITProvisionStatus() *string {
	return s.JITProvisionStatus
}

func (s *GetSAMLIdentityProviderResponseBodySSOIdentityProviderConfiguration) GetJITUpdateStatus() *string {
	return s.JITUpdateStatus
}

func (s *GetSAMLIdentityProviderResponseBodySSOIdentityProviderConfiguration) GetLoginURL() *string {
	return s.LoginURL
}

func (s *GetSAMLIdentityProviderResponseBodySSOIdentityProviderConfiguration) GetSAMLBindingType() *string {
	return s.SAMLBindingType
}

func (s *GetSAMLIdentityProviderResponseBodySSOIdentityProviderConfiguration) GetSSOStatus() *string {
	return s.SSOStatus
}

func (s *GetSAMLIdentityProviderResponseBodySSOIdentityProviderConfiguration) GetX509Certificates() []*GetSAMLIdentityProviderResponseBodySSOIdentityProviderConfigurationX509Certificates {
	return s.X509Certificates
}

func (s *GetSAMLIdentityProviderResponseBodySSOIdentityProviderConfiguration) SetEntityId(v string) *GetSAMLIdentityProviderResponseBodySSOIdentityProviderConfiguration {
	s.EntityId = &v
	return s
}

func (s *GetSAMLIdentityProviderResponseBodySSOIdentityProviderConfiguration) SetJITProvisionStatus(v string) *GetSAMLIdentityProviderResponseBodySSOIdentityProviderConfiguration {
	s.JITProvisionStatus = &v
	return s
}

func (s *GetSAMLIdentityProviderResponseBodySSOIdentityProviderConfiguration) SetJITUpdateStatus(v string) *GetSAMLIdentityProviderResponseBodySSOIdentityProviderConfiguration {
	s.JITUpdateStatus = &v
	return s
}

func (s *GetSAMLIdentityProviderResponseBodySSOIdentityProviderConfiguration) SetLoginURL(v string) *GetSAMLIdentityProviderResponseBodySSOIdentityProviderConfiguration {
	s.LoginURL = &v
	return s
}

func (s *GetSAMLIdentityProviderResponseBodySSOIdentityProviderConfiguration) SetSAMLBindingType(v string) *GetSAMLIdentityProviderResponseBodySSOIdentityProviderConfiguration {
	s.SAMLBindingType = &v
	return s
}

func (s *GetSAMLIdentityProviderResponseBodySSOIdentityProviderConfiguration) SetSSOStatus(v string) *GetSAMLIdentityProviderResponseBodySSOIdentityProviderConfiguration {
	s.SSOStatus = &v
	return s
}

func (s *GetSAMLIdentityProviderResponseBodySSOIdentityProviderConfiguration) SetX509Certificates(v []*GetSAMLIdentityProviderResponseBodySSOIdentityProviderConfigurationX509Certificates) *GetSAMLIdentityProviderResponseBodySSOIdentityProviderConfiguration {
	s.X509Certificates = v
	return s
}

func (s *GetSAMLIdentityProviderResponseBodySSOIdentityProviderConfiguration) Validate() error {
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

type GetSAMLIdentityProviderResponseBodySSOIdentityProviderConfigurationX509Certificates struct {
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

func (s GetSAMLIdentityProviderResponseBodySSOIdentityProviderConfigurationX509Certificates) String() string {
	return dara.Prettify(s)
}

func (s GetSAMLIdentityProviderResponseBodySSOIdentityProviderConfigurationX509Certificates) GoString() string {
	return s.String()
}

func (s *GetSAMLIdentityProviderResponseBodySSOIdentityProviderConfigurationX509Certificates) GetCertificateId() *string {
	return s.CertificateId
}

func (s *GetSAMLIdentityProviderResponseBodySSOIdentityProviderConfigurationX509Certificates) GetX509Certificate() *string {
	return s.X509Certificate
}

func (s *GetSAMLIdentityProviderResponseBodySSOIdentityProviderConfigurationX509Certificates) SetCertificateId(v string) *GetSAMLIdentityProviderResponseBodySSOIdentityProviderConfigurationX509Certificates {
	s.CertificateId = &v
	return s
}

func (s *GetSAMLIdentityProviderResponseBodySSOIdentityProviderConfigurationX509Certificates) SetX509Certificate(v string) *GetSAMLIdentityProviderResponseBodySSOIdentityProviderConfigurationX509Certificates {
	s.X509Certificate = &v
	return s
}

func (s *GetSAMLIdentityProviderResponseBodySSOIdentityProviderConfigurationX509Certificates) Validate() error {
	return dara.Validate(s)
}
