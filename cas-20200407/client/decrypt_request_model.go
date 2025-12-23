// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDecryptRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlgorithm(v string) *DecryptRequest
	GetAlgorithm() *string
	SetCertIdentifier(v string) *DecryptRequest
	GetCertIdentifier() *string
	SetCiphertextBlob(v string) *DecryptRequest
	GetCiphertextBlob() *string
	SetCustomIdentifier(v string) *DecryptRequest
	GetCustomIdentifier() *string
	SetMessageType(v string) *DecryptRequest
	GetMessageType() *string
}

type DecryptRequest struct {
	// The encryption algorithm. Valid values:
	//
	// 	- **RSAES_OAEP_SHA_1**
	//
	// 	- **RSAES_OAEP_SHA_256**
	//
	// 	- **SM2PKE**
	//
	// This parameter is required.
	//
	// example:
	//
	// RSAESOAEPSHA_1
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// The unique identifier of the certificate. You can call the [ListCert](https://help.aliyun.com/document_detail/455806.html) operation to query the identifier.
	//
	// 	- If the certificate is an SSL certificate, the value of this parameter must be in the {Certificate ID}-cn-hangzhou format.
	//
	// 	- If the certificate is a private certificate, the value of this parameter must be the value of the Identifier field for the private certificate.
	//
	// example:
	//
	// 12345678-1234-1234-1234-12345678****
	CertIdentifier *string `json:"CertIdentifier,omitempty" xml:"CertIdentifier,omitempty"`
	// The data that you want to decrypt. The value is encoded in Base64.
	//
	// This parameter is required.
	//
	// example:
	//
	// ZOyIygCyaOW6Gj****MlNKiuyjfzw=
	CiphertextBlob   *string `json:"CiphertextBlob,omitempty" xml:"CiphertextBlob,omitempty"`
	CustomIdentifier *string `json:"CustomIdentifier,omitempty" xml:"CustomIdentifier,omitempty"`
	// The value type of the Message parameter. Valid values:
	//
	// 	- RAW: The returned result is raw data encoded in UTF-8.
	//
	// 	- Base64: The returned result is Base64-encoded data. This is the default value.
	//
	// example:
	//
	// Base64
	MessageType *string `json:"MessageType,omitempty" xml:"MessageType,omitempty"`
}

func (s DecryptRequest) String() string {
	return dara.Prettify(s)
}

func (s DecryptRequest) GoString() string {
	return s.String()
}

func (s *DecryptRequest) GetAlgorithm() *string {
	return s.Algorithm
}

func (s *DecryptRequest) GetCertIdentifier() *string {
	return s.CertIdentifier
}

func (s *DecryptRequest) GetCiphertextBlob() *string {
	return s.CiphertextBlob
}

func (s *DecryptRequest) GetCustomIdentifier() *string {
	return s.CustomIdentifier
}

func (s *DecryptRequest) GetMessageType() *string {
	return s.MessageType
}

func (s *DecryptRequest) SetAlgorithm(v string) *DecryptRequest {
	s.Algorithm = &v
	return s
}

func (s *DecryptRequest) SetCertIdentifier(v string) *DecryptRequest {
	s.CertIdentifier = &v
	return s
}

func (s *DecryptRequest) SetCiphertextBlob(v string) *DecryptRequest {
	s.CiphertextBlob = &v
	return s
}

func (s *DecryptRequest) SetCustomIdentifier(v string) *DecryptRequest {
	s.CustomIdentifier = &v
	return s
}

func (s *DecryptRequest) SetMessageType(v string) *DecryptRequest {
	s.MessageType = &v
	return s
}

func (s *DecryptRequest) Validate() error {
	return dara.Validate(s)
}
