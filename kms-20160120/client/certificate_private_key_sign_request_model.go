// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCertificatePrivateKeySignRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlgorithm(v string) *CertificatePrivateKeySignRequest
	GetAlgorithm() *string
	SetCertificateId(v string) *CertificatePrivateKeySignRequest
	GetCertificateId() *string
	SetMessage(v string) *CertificatePrivateKeySignRequest
	GetMessage() *string
	SetMessageType(v string) *CertificatePrivateKeySignRequest
	GetMessageType() *string
}

type CertificatePrivateKeySignRequest struct {
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
	// >	- The SM2DSA signature algorithm is supported only in regions where managed hardware security modules (HSMs) are used in mainland China. For more information, see [Managed HSM overview](https://help.aliyun.com/document_detail/125803.html).
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
	// The data to be signed.
	//
	// The value is encoded in Base64. For example, if the hexadecimal data that you want to sign is `[0x31, 0x32, 0x33, 0x34]`, the Base64-encoded data is `MTIzNA==`.
	//
	// If the MessageType parameter is set to RAW, the size of the data must be less than or equal to 4 KB.
	//
	// If the size of the data is greater than 4 KB, you can set the MessageType parameter to DIGEST and set the Message parameter to the digest of the data. The digest is also called hash value. You can compute the digest of the data on an on-premises machine. Certificates Manager uses the digest that you compute in your own certificate application system. The message digest algorithm that you use must match the specified signature algorithm. Comply with the following mapping between signature algorithms and message digest algorithms:
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
}

func (s CertificatePrivateKeySignRequest) String() string {
	return dara.Prettify(s)
}

func (s CertificatePrivateKeySignRequest) GoString() string {
	return s.String()
}

func (s *CertificatePrivateKeySignRequest) GetAlgorithm() *string {
	return s.Algorithm
}

func (s *CertificatePrivateKeySignRequest) GetCertificateId() *string {
	return s.CertificateId
}

func (s *CertificatePrivateKeySignRequest) GetMessage() *string {
	return s.Message
}

func (s *CertificatePrivateKeySignRequest) GetMessageType() *string {
	return s.MessageType
}

func (s *CertificatePrivateKeySignRequest) SetAlgorithm(v string) *CertificatePrivateKeySignRequest {
	s.Algorithm = &v
	return s
}

func (s *CertificatePrivateKeySignRequest) SetCertificateId(v string) *CertificatePrivateKeySignRequest {
	s.CertificateId = &v
	return s
}

func (s *CertificatePrivateKeySignRequest) SetMessage(v string) *CertificatePrivateKeySignRequest {
	s.Message = &v
	return s
}

func (s *CertificatePrivateKeySignRequest) SetMessageType(v string) *CertificatePrivateKeySignRequest {
	s.MessageType = &v
	return s
}

func (s *CertificatePrivateKeySignRequest) Validate() error {
	return dara.Validate(s)
}
