// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddExternalSAMLIdPCertificateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirectoryId(v string) *AddExternalSAMLIdPCertificateRequest
	GetDirectoryId() *string
	SetX509Certificate(v string) *AddExternalSAMLIdPCertificateRequest
	GetX509Certificate() *string
}

type AddExternalSAMLIdPCertificateRequest struct {
	// The ID of the directory.
	//
	// example:
	//
	// d-00fc2p61****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The X.509 certificate in the PEM format.
	//
	// The certificate is provided by the SAML identity provider (IdP).
	//
	// example:
	//
	// MIIC8DCCAdigAwIBAgIQP9eomUYGeoND****
	X509Certificate *string `json:"X509Certificate,omitempty" xml:"X509Certificate,omitempty"`
}

func (s AddExternalSAMLIdPCertificateRequest) String() string {
	return dara.Prettify(s)
}

func (s AddExternalSAMLIdPCertificateRequest) GoString() string {
	return s.String()
}

func (s *AddExternalSAMLIdPCertificateRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *AddExternalSAMLIdPCertificateRequest) GetX509Certificate() *string {
	return s.X509Certificate
}

func (s *AddExternalSAMLIdPCertificateRequest) SetDirectoryId(v string) *AddExternalSAMLIdPCertificateRequest {
	s.DirectoryId = &v
	return s
}

func (s *AddExternalSAMLIdPCertificateRequest) SetX509Certificate(v string) *AddExternalSAMLIdPCertificateRequest {
	s.X509Certificate = &v
	return s
}

func (s *AddExternalSAMLIdPCertificateRequest) Validate() error {
	return dara.Validate(s)
}
