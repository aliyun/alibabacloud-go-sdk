// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCertificatePublicKeyVerifyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlgorithm(v string) *CertificatePublicKeyVerifyRequest
	GetAlgorithm() *string
	SetCertificateId(v string) *CertificatePublicKeyVerifyRequest
	GetCertificateId() *string
	SetMessage(v string) *CertificatePublicKeyVerifyRequest
	GetMessage() *string
	SetMessageType(v string) *CertificatePublicKeyVerifyRequest
	GetMessageType() *string
	SetSignatureValue(v string) *CertificatePublicKeyVerifyRequest
	GetSignatureValue() *string
}

type CertificatePublicKeyVerifyRequest struct {
	// The signature algorithm. Valid values:
	//
	// 	- RSA_PKCS1_SHA_256
	//
	// 	- RSA_PSS_SHA_256
	//
	// 	- ECDSA_SHA_256
	//
	// 	- SM2DSA
	//
	// > The SM2DSA signature algorithm is supported only in regions where managed hardware security modules (HSMs) are used in the Chinese mainland. For more information, see [Managed HSM overview](https://help.aliyun.com/document_detail/125803.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// ECDSA_SHA_256
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// The ID of the certificate. The ID must be globally unique in Certificates Manager.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345678-1234-1234-1234-12345678****
	CertificateId *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
	// The raw data that is signed.
	//
	// The value is encoded in Base64. For example, if the raw data in the hexadecimal format is `[0x31, 0x32, 0x33, 0x34]`, set this parameter to the Base64-encoded value `MTIzNA==`.
	//
	// If the MessageType parameter is set to RAW, the size of the data must be less than or equal to 4 KB.
	//
	// If the size of the data is greater than 4 KB, you can set the MessageType parameter to DIGEST and set the Message parameter to the digest of the data. The digest is also called hash value. You can compute the digest of the data on an on-premises device. Certificates Manager uses the digest that you compute in your own certificate application system. The message digest algorithm that you use must match the specified signature algorithm. Comply with the following mapping between signature algorithms and message digest algorithms:
	//
	// 	- If the signature algorithm is RSA_PKCS1_SHA_256, RSA_PSS_SHA_256, or ECDSA_SHA_256, the message digest algorithm must be SHA-256.
	//
	// 	- If the signature algorithm is SM2DSA, the message digest algorithm must be SM3.
	//
	// >  If the key type of the certificate is EC_SM2 and the MessageType parameter is set to DIGEST, the value of the Message parameter is `e` that is described in GB/T 32918.2-2016 6.1.
	//
	// This parameter is required.
	//
	// example:
	//
	// VGhlIHF1aWNrIGJyb3duIGZveCBqdW1wcyBvdmVyIHRoZSBsYXp5IGRvZy4=
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The type of the message. Valid values:
	//
	// 	- RAW: the raw data. This is the default value.
	//
	// 	- DIGEST: the message digest (hash value) of the raw data.
	//
	// This parameter is required.
	//
	// example:
	//
	// RAW
	MessageType *string `json:"MessageType,omitempty" xml:"MessageType,omitempty"`
	// The signature value.
	//
	// The value is encoded in Base64.
	//
	// This parameter is required.
	//
	// example:
	//
	// ZOyIygCyaOW6Gj****MlNKiuyjfzw=
	SignatureValue *string `json:"SignatureValue,omitempty" xml:"SignatureValue,omitempty"`
}

func (s CertificatePublicKeyVerifyRequest) String() string {
	return dara.Prettify(s)
}

func (s CertificatePublicKeyVerifyRequest) GoString() string {
	return s.String()
}

func (s *CertificatePublicKeyVerifyRequest) GetAlgorithm() *string {
	return s.Algorithm
}

func (s *CertificatePublicKeyVerifyRequest) GetCertificateId() *string {
	return s.CertificateId
}

func (s *CertificatePublicKeyVerifyRequest) GetMessage() *string {
	return s.Message
}

func (s *CertificatePublicKeyVerifyRequest) GetMessageType() *string {
	return s.MessageType
}

func (s *CertificatePublicKeyVerifyRequest) GetSignatureValue() *string {
	return s.SignatureValue
}

func (s *CertificatePublicKeyVerifyRequest) SetAlgorithm(v string) *CertificatePublicKeyVerifyRequest {
	s.Algorithm = &v
	return s
}

func (s *CertificatePublicKeyVerifyRequest) SetCertificateId(v string) *CertificatePublicKeyVerifyRequest {
	s.CertificateId = &v
	return s
}

func (s *CertificatePublicKeyVerifyRequest) SetMessage(v string) *CertificatePublicKeyVerifyRequest {
	s.Message = &v
	return s
}

func (s *CertificatePublicKeyVerifyRequest) SetMessageType(v string) *CertificatePublicKeyVerifyRequest {
	s.MessageType = &v
	return s
}

func (s *CertificatePublicKeyVerifyRequest) SetSignatureValue(v string) *CertificatePublicKeyVerifyRequest {
	s.SignatureValue = &v
	return s
}

func (s *CertificatePublicKeyVerifyRequest) Validate() error {
	return dara.Validate(s)
}
