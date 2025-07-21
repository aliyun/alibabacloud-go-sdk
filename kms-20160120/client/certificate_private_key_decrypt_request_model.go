// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCertificatePrivateKeyDecryptRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlgorithm(v string) *CertificatePrivateKeyDecryptRequest
	GetAlgorithm() *string
	SetCertificateId(v string) *CertificatePrivateKeyDecryptRequest
	GetCertificateId() *string
	SetCiphertextBlob(v string) *CertificatePrivateKeyDecryptRequest
	GetCiphertextBlob() *string
}

type CertificatePrivateKeyDecryptRequest struct {
	// The encryption algorithm. Valid values:
	//
	// 	- RSAES_OAEP_SHA_1
	//
	// 	- RSAES_OAEP_SHA_256
	//
	// 	- SM2PKE
	//
	// > The SM2PKE encryption algorithm is supported only in regions in mainland China. In these regions, managed hardware security modules (HSMs) are used. For more information, see [Managed HSM overview](https://help.aliyun.com/document_detail/125803.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// RSAES_OAEP_SHA_256
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// The ID of the certificate. The ID must be globally unique in Certificates Manager.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345678-1234-1234-1234-12345678****
	CertificateId *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
	// The data that you want to decrypt.
	//
	// The value is encoded in Base64.
	//
	// This parameter is required.
	//
	// example:
	//
	// ZOyIygCyaOW6Gj****MlNKiuyjfzw=
	CiphertextBlob *string `json:"CiphertextBlob,omitempty" xml:"CiphertextBlob,omitempty"`
}

func (s CertificatePrivateKeyDecryptRequest) String() string {
	return dara.Prettify(s)
}

func (s CertificatePrivateKeyDecryptRequest) GoString() string {
	return s.String()
}

func (s *CertificatePrivateKeyDecryptRequest) GetAlgorithm() *string {
	return s.Algorithm
}

func (s *CertificatePrivateKeyDecryptRequest) GetCertificateId() *string {
	return s.CertificateId
}

func (s *CertificatePrivateKeyDecryptRequest) GetCiphertextBlob() *string {
	return s.CiphertextBlob
}

func (s *CertificatePrivateKeyDecryptRequest) SetAlgorithm(v string) *CertificatePrivateKeyDecryptRequest {
	s.Algorithm = &v
	return s
}

func (s *CertificatePrivateKeyDecryptRequest) SetCertificateId(v string) *CertificatePrivateKeyDecryptRequest {
	s.CertificateId = &v
	return s
}

func (s *CertificatePrivateKeyDecryptRequest) SetCiphertextBlob(v string) *CertificatePrivateKeyDecryptRequest {
	s.CiphertextBlob = &v
	return s
}

func (s *CertificatePrivateKeyDecryptRequest) Validate() error {
	return dara.Validate(s)
}
