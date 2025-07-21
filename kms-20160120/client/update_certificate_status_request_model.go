// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCertificateStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCertificateId(v string) *UpdateCertificateStatusRequest
	GetCertificateId() *string
	SetStatus(v string) *UpdateCertificateStatusRequest
	GetStatus() *string
}

type UpdateCertificateStatusRequest struct {
	// The ID of the certificate. The ID must be globally unique in Certificates Manager.
	//
	// This parameter is required.
	//
	// example:
	//
	// 9a28de48-8d8b-484d-a766-dec4****
	CertificateId *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
	// The status of the certificate. Valid values:
	//
	// 	- INACTIVE: The certificate is disabled.
	//
	// 	- ACTIVE: The certificate is enabled.
	//
	// 	- REVOKED: The certificate is revoked.
	//
	// > If the certificate is in the REVOKED state, you can use the certificate only to verify a signature, but not to generate a signature.
	//
	// This parameter is required.
	//
	// example:
	//
	// INACTIVE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateCertificateStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateCertificateStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateCertificateStatusRequest) GetCertificateId() *string {
	return s.CertificateId
}

func (s *UpdateCertificateStatusRequest) GetStatus() *string {
	return s.Status
}

func (s *UpdateCertificateStatusRequest) SetCertificateId(v string) *UpdateCertificateStatusRequest {
	s.CertificateId = &v
	return s
}

func (s *UpdateCertificateStatusRequest) SetStatus(v string) *UpdateCertificateStatusRequest {
	s.Status = &v
	return s
}

func (s *UpdateCertificateStatusRequest) Validate() error {
	return dara.Validate(s)
}
