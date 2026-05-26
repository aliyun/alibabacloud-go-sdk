// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSAMLIdentityProviderCertificatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListSAMLIdentityProviderCertificatesResponseBody
	GetRequestId() *string
	SetX509Certificates(v []*ListSAMLIdentityProviderCertificatesResponseBodyX509Certificates) *ListSAMLIdentityProviderCertificatesResponseBody
	GetX509Certificates() []*ListSAMLIdentityProviderCertificatesResponseBodyX509Certificates
}

type ListSAMLIdentityProviderCertificatesResponseBody struct {
	// example:
	//
	// AABD6E03-4B3A-5687-88FF-72232670ED0C
	RequestId        *string                                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	X509Certificates []*ListSAMLIdentityProviderCertificatesResponseBodyX509Certificates `json:"X509Certificates,omitempty" xml:"X509Certificates,omitempty" type:"Repeated"`
}

func (s ListSAMLIdentityProviderCertificatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSAMLIdentityProviderCertificatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListSAMLIdentityProviderCertificatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSAMLIdentityProviderCertificatesResponseBody) GetX509Certificates() []*ListSAMLIdentityProviderCertificatesResponseBodyX509Certificates {
	return s.X509Certificates
}

func (s *ListSAMLIdentityProviderCertificatesResponseBody) SetRequestId(v string) *ListSAMLIdentityProviderCertificatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSAMLIdentityProviderCertificatesResponseBody) SetX509Certificates(v []*ListSAMLIdentityProviderCertificatesResponseBodyX509Certificates) *ListSAMLIdentityProviderCertificatesResponseBody {
	s.X509Certificates = v
	return s
}

func (s *ListSAMLIdentityProviderCertificatesResponseBody) Validate() error {
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

type ListSAMLIdentityProviderCertificatesResponseBodyX509Certificates struct {
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

func (s ListSAMLIdentityProviderCertificatesResponseBodyX509Certificates) String() string {
	return dara.Prettify(s)
}

func (s ListSAMLIdentityProviderCertificatesResponseBodyX509Certificates) GoString() string {
	return s.String()
}

func (s *ListSAMLIdentityProviderCertificatesResponseBodyX509Certificates) GetCertificateId() *string {
	return s.CertificateId
}

func (s *ListSAMLIdentityProviderCertificatesResponseBodyX509Certificates) GetX509Certificate() *string {
	return s.X509Certificate
}

func (s *ListSAMLIdentityProviderCertificatesResponseBodyX509Certificates) SetCertificateId(v string) *ListSAMLIdentityProviderCertificatesResponseBodyX509Certificates {
	s.CertificateId = &v
	return s
}

func (s *ListSAMLIdentityProviderCertificatesResponseBodyX509Certificates) SetX509Certificate(v string) *ListSAMLIdentityProviderCertificatesResponseBodyX509Certificates {
	s.X509Certificate = &v
	return s
}

func (s *ListSAMLIdentityProviderCertificatesResponseBodyX509Certificates) Validate() error {
	return dara.Validate(s)
}
