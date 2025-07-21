// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCertificatePublicKeyEncryptRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlgorithm(v string) *CertificatePublicKeyEncryptRequest
	GetAlgorithm() *string
	SetCertificateId(v string) *CertificatePublicKeyEncryptRequest
	GetCertificateId() *string
	SetPlaintext(v string) *CertificatePublicKeyEncryptRequest
	GetPlaintext() *string
}

type CertificatePublicKeyEncryptRequest struct {
	// The encryption algorithm. Valid values:
	//
	// 	- RSAES_OAEP_SHA_1
	//
	// 	- RSAES_OAEP_SHA_256
	//
	// 	- SM2PKE
	//
	// >The SM2PKE encryption algorithm is supported only in regions in mainland China. In these regions, managed hardware security modules (HSMs) are used. For more information, see [Managed HSM overview](https://help.aliyun.com/document_detail/125803.html).
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
	// The data that you want to encrypt.
	//
	// The value is encoded in Base64. For example, if the hexadecimal data that you want to encrypt is `[0x31, 0x32, 0x33, 0x34]`, the Base64-encoded data is `MTIzNA==`.
	//
	// The size of data that can be encrypted varies based on the encryption algorithm that you use:
	//
	// 	- RSAES_OAEP_SHA_1: 214 bytes
	//
	// 	- RSAES_OAEP_SHA_256: 190 bytes
	//
	// 	- SM2PKE: 6,047 bytes
	//
	// If the size of data that you want to encrypt exceeds the preceding limits, you can call the [GenerateDataKey](https://help.aliyun.com/document_detail/28948.html) operation to generate a data key to encrypt the data. Then, call the CertificatePublicKeyEncrypt operation to encrypt the data key.
	//
	// This parameter is required.
	//
	// example:
	//
	// VGhlIHF1aWNrIGJyb3duIGZveCBqdW1wcyBvdmVyIHRoZSBsYXp5IGRvZy4=
	Plaintext *string `json:"Plaintext,omitempty" xml:"Plaintext,omitempty"`
}

func (s CertificatePublicKeyEncryptRequest) String() string {
	return dara.Prettify(s)
}

func (s CertificatePublicKeyEncryptRequest) GoString() string {
	return s.String()
}

func (s *CertificatePublicKeyEncryptRequest) GetAlgorithm() *string {
	return s.Algorithm
}

func (s *CertificatePublicKeyEncryptRequest) GetCertificateId() *string {
	return s.CertificateId
}

func (s *CertificatePublicKeyEncryptRequest) GetPlaintext() *string {
	return s.Plaintext
}

func (s *CertificatePublicKeyEncryptRequest) SetAlgorithm(v string) *CertificatePublicKeyEncryptRequest {
	s.Algorithm = &v
	return s
}

func (s *CertificatePublicKeyEncryptRequest) SetCertificateId(v string) *CertificatePublicKeyEncryptRequest {
	s.CertificateId = &v
	return s
}

func (s *CertificatePublicKeyEncryptRequest) SetPlaintext(v string) *CertificatePublicKeyEncryptRequest {
	s.Plaintext = &v
	return s
}

func (s *CertificatePublicKeyEncryptRequest) Validate() error {
	return dara.Validate(s)
}
