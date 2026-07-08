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
	SetWarehouseId(v int64) *DecryptRequest
	GetWarehouseId() *int64
}

type DecryptRequest struct {
	// The encryption algorithm. Valid values:
	//
	// - **RSAES_OAEP_SHA_1**
	//
	// - **RSAES_OAEP_SHA_256**
	//
	// - **SM2PKE**
	//
	// This parameter is required.
	//
	// example:
	//
	// RSAESOAEPSHA_1
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// The unique identifier of the certificate. Call [ListCert](https://help.aliyun.com/document_detail/455806.html) to obtain this parameter.
	//
	// - The identifier of an SSL certificate is typically in the format {Certificate ID}-cn-hangzhou.
	//
	// - For a private certificate authority (PCA) certificate, this is the value of the Identifier field of the private certificate.
	//
	// example:
	//
	// 1ef1da5f-38ed-69b3-****-037781890265
	CertIdentifier *string `json:"CertIdentifier,omitempty" xml:"CertIdentifier,omitempty"`
	// The Base64-encoded data to decrypt.
	//
	// This parameter is required.
	//
	// example:
	//
	// ZOyIygCyaOW6Gj****MlNKiuyjfzw=
	CiphertextBlob *string `json:"CiphertextBlob,omitempty" xml:"CiphertextBlob,omitempty"`
	// A custom identifier that serves as a unique key.
	//
	// example:
	//
	// ****6bb538d538c70c01f81jh2****
	CustomIdentifier *string `json:"CustomIdentifier,omitempty" xml:"CustomIdentifier,omitempty"`
	// The message type. Valid values:
	//
	// - RAW: The response returns the plaintext in UTF-8 encoding.
	//
	// - Base64 (default): The response returns the Base64-encoded plaintext.
	//
	// example:
	//
	// Base64
	MessageType *string `json:"MessageType,omitempty" xml:"MessageType,omitempty"`
	// The ID of the repository.
	//
	// > Call [ListCertWarehouse](https://help.aliyun.com/document_detail/455805.html) to obtain this ID.
	//
	// example:
	//
	// 1
	WarehouseId *int64 `json:"WarehouseId,omitempty" xml:"WarehouseId,omitempty"`
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

func (s *DecryptRequest) GetWarehouseId() *int64 {
	return s.WarehouseId
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

func (s *DecryptRequest) SetWarehouseId(v int64) *DecryptRequest {
	s.WarehouseId = &v
	return s
}

func (s *DecryptRequest) Validate() error {
	return dara.Validate(s)
}
