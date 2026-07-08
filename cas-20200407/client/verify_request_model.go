// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCertIdentifier(v string) *VerifyRequest
	GetCertIdentifier() *string
	SetCustomIdentifier(v string) *VerifyRequest
	GetCustomIdentifier() *string
	SetMessage(v string) *VerifyRequest
	GetMessage() *string
	SetMessageType(v string) *VerifyRequest
	GetMessageType() *string
	SetSignatureValue(v string) *VerifyRequest
	GetSignatureValue() *string
	SetSigningAlgorithm(v string) *VerifyRequest
	GetSigningAlgorithm() *string
	SetWarehouseId(v string) *VerifyRequest
	GetWarehouseId() *string
}

type VerifyRequest struct {
	// The unique identifier of the certificate. To get this parameter, call the [ListCert](https://help.aliyun.com/document_detail/455806.html) operation.
	//
	// - The identifier for an SSL certificate is typically in the format \\`{Certificate ID}-cn-hangzhou\\`.
	//
	// - For a PCA certificate, this is the value of the \\`Identifier\\` field.
	//
	// example:
	//
	// 1ef1da5f-38ed-69b3-****-037781890265
	CertIdentifier *string `json:"CertIdentifier,omitempty" xml:"CertIdentifier,omitempty"`
	// The custom identifier. This key must be unique.
	//
	// example:
	//
	// ****6bb538d538c70c01f81jh2****
	CustomIdentifier *string `json:"CustomIdentifier,omitempty" xml:"CustomIdentifier,omitempty"`
	// The data to verify. The data must be Base64-encoded. For example, if the hexadecimal content of the data to sign is \\`[0x31, 0x32, 0x33, 0x34]\\`, the Base64-encoded value is \\`MTIzNA==\\`. If you set \\`MessageType\\` to \\`RAW\\`, the data size must be less than 4 KB. If the data to sign is larger than 4 KB, set \\`MessageType\\` to \\`DIGEST\\`. Then, set \\`Message\\` to the message digest, or hash, that you calculate locally. The hashing algorithm for the digest must be compatible with the signature algorithm:<br>
	//
	// - The hashing algorithm for \\`SHA256withRSA\\`, \\`SHA256withRSA/PSS\\`, and \\`SHA256withECDSA\\` is SHA-256.
	//
	// - The hashing algorithm for \\`SM3withSM2\\` is SM3.
	//
	// This parameter is required.
	//
	// example:
	//
	// MTIzNA==
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The message type. Valid values:
	//
	// - **RAW*	- (default): The raw data.
	//
	// - **DIGEST**: The message digest of the raw data.
	//
	// This parameter is required.
	//
	// example:
	//
	// RAW
	MessageType *string `json:"MessageType,omitempty" xml:"MessageType,omitempty"`
	// The signature value. The value must be Base64-encoded.
	//
	// This parameter is required.
	//
	// example:
	//
	// eyaC0w3ROK5b3QcHmUtAhMY/sQjKu2t3uBfnf6J/gn7JfZtyxwcCUjzXbw5jmqJQRbj1te670Bshg9kUdanKhtHFhJjU5jX+ZMMBr6pH0gqQDJxR0K0yHXRc0Q5OQoUZ6BfpbI4Wt4jJvJSdCstz1vSg12CfEHS8Kd5qfhItK7Y=
	SignatureValue *string `json:"SignatureValue,omitempty" xml:"SignatureValue,omitempty"`
	// The signature algorithm. Valid values:
	//
	// - **SHA256withRSA**
	//
	// - **SHA256withRSA/PSS**
	//
	// - **SHA256withECDSA**
	//
	// - **SM3withSM2**
	//
	// This parameter is required.
	//
	// example:
	//
	// SHA256withRSA
	SigningAlgorithm *string `json:"SigningAlgorithm,omitempty" xml:"SigningAlgorithm,omitempty"`
	// The ID of the repository. To get this parameter, call the [ListCertWarehouse](https://help.aliyun.com/document_detail/453246.html) operation.
	//
	// example:
	//
	// 1
	WarehouseId *string `json:"WarehouseId,omitempty" xml:"WarehouseId,omitempty"`
}

func (s VerifyRequest) String() string {
	return dara.Prettify(s)
}

func (s VerifyRequest) GoString() string {
	return s.String()
}

func (s *VerifyRequest) GetCertIdentifier() *string {
	return s.CertIdentifier
}

func (s *VerifyRequest) GetCustomIdentifier() *string {
	return s.CustomIdentifier
}

func (s *VerifyRequest) GetMessage() *string {
	return s.Message
}

func (s *VerifyRequest) GetMessageType() *string {
	return s.MessageType
}

func (s *VerifyRequest) GetSignatureValue() *string {
	return s.SignatureValue
}

func (s *VerifyRequest) GetSigningAlgorithm() *string {
	return s.SigningAlgorithm
}

func (s *VerifyRequest) GetWarehouseId() *string {
	return s.WarehouseId
}

func (s *VerifyRequest) SetCertIdentifier(v string) *VerifyRequest {
	s.CertIdentifier = &v
	return s
}

func (s *VerifyRequest) SetCustomIdentifier(v string) *VerifyRequest {
	s.CustomIdentifier = &v
	return s
}

func (s *VerifyRequest) SetMessage(v string) *VerifyRequest {
	s.Message = &v
	return s
}

func (s *VerifyRequest) SetMessageType(v string) *VerifyRequest {
	s.MessageType = &v
	return s
}

func (s *VerifyRequest) SetSignatureValue(v string) *VerifyRequest {
	s.SignatureValue = &v
	return s
}

func (s *VerifyRequest) SetSigningAlgorithm(v string) *VerifyRequest {
	s.SigningAlgorithm = &v
	return s
}

func (s *VerifyRequest) SetWarehouseId(v string) *VerifyRequest {
	s.WarehouseId = &v
	return s
}

func (s *VerifyRequest) Validate() error {
	return dara.Validate(s)
}
