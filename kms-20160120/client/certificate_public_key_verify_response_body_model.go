// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCertificatePublicKeyVerifyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCertificateId(v string) *CertificatePublicKeyVerifyResponseBody
	GetCertificateId() *string
	SetRequestId(v string) *CertificatePublicKeyVerifyResponseBody
	GetRequestId() *string
	SetSignatureValid(v bool) *CertificatePublicKeyVerifyResponseBody
	GetSignatureValid() *bool
}

type CertificatePublicKeyVerifyResponseBody struct {
	// The ID of the certificate.
	//
	// example:
	//
	// 12345678-1234-1234-1234-12345678****
	CertificateId *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 5979d897-d69f-4fc9-87dd-f3bb73c40b80
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The verification result. Valid values:
	//
	// 	- true: The signature is valid.
	//
	// 	- false: The signature is invalid.
	//
	// example:
	//
	// true
	SignatureValid *bool `json:"SignatureValid,omitempty" xml:"SignatureValid,omitempty"`
}

func (s CertificatePublicKeyVerifyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CertificatePublicKeyVerifyResponseBody) GoString() string {
	return s.String()
}

func (s *CertificatePublicKeyVerifyResponseBody) GetCertificateId() *string {
	return s.CertificateId
}

func (s *CertificatePublicKeyVerifyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CertificatePublicKeyVerifyResponseBody) GetSignatureValid() *bool {
	return s.SignatureValid
}

func (s *CertificatePublicKeyVerifyResponseBody) SetCertificateId(v string) *CertificatePublicKeyVerifyResponseBody {
	s.CertificateId = &v
	return s
}

func (s *CertificatePublicKeyVerifyResponseBody) SetRequestId(v string) *CertificatePublicKeyVerifyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CertificatePublicKeyVerifyResponseBody) SetSignatureValid(v bool) *CertificatePublicKeyVerifyResponseBody {
	s.SignatureValid = &v
	return s
}

func (s *CertificatePublicKeyVerifyResponseBody) Validate() error {
	return dara.Validate(s)
}
