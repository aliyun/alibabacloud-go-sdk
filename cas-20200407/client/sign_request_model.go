// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSignRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCertIdentifier(v string) *SignRequest
	GetCertIdentifier() *string
	SetCustomIdentifier(v string) *SignRequest
	GetCustomIdentifier() *string
	SetMessage(v string) *SignRequest
	GetMessage() *string
	SetMessageType(v string) *SignRequest
	GetMessageType() *string
	SetSigningAlgorithm(v string) *SignRequest
	GetSigningAlgorithm() *string
	SetWarehouseId(v int64) *SignRequest
	GetWarehouseId() *int64
}

type SignRequest struct {
	// The unique identifier of the certificate. You can get this value by calling the [ListCert](https://help.aliyun.com/document_detail/455806.html) operation.
	//
	// - The identifier of an SSL certificate is typically in the format \\"{Certificate ID}-cn-hangzhou\\".
	//
	// - For a PCA certificate, this is the Identifier from the corresponding private certificate.
	//
	// example:
	//
	// 1ef1da5f-38ed-69b3-****-037781890265
	CertIdentifier *string `json:"CertIdentifier,omitempty" xml:"CertIdentifier,omitempty"`
	// A unique, user-defined identifier.
	//
	// example:
	//
	// ***e6bb538d538c70c01f81fg3****
	CustomIdentifier *string `json:"CustomIdentifier,omitempty" xml:"CustomIdentifier,omitempty"`
	// The data to sign. The MessageType parameter determines the format of this data. If MessageType is set to RAW, Message is the raw data. If MessageType is set to BASE64, Message is the Base64-encoded raw data. If MessageType is set to DIGEST, Message is the message digest (hash value). If MessageType is set to BLIND, Message is the Base64-encoded blinded message.
	//
	// This parameter is required.
	//
	// example:
	//
	// MTIzNA==
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The message type. Valid values:
	//
	// - `RAW` (default): The raw data.
	//
	// - `DIGEST`: The message digest (hash value) of the raw data.
	//
	// - `BASE64`: The Base64-encoded raw data.
	//
	// - `BLIND`: Enables blind signing. This is supported only for certificates that use an RSA algorithm.
	//
	// This parameter is required.
	//
	// example:
	//
	// RAW
	MessageType *string `json:"MessageType,omitempty" xml:"MessageType,omitempty"`
	// The signature algorithm. Valid values:
	//
	// - `SHA256withRSA`
	//
	// - `SHA256withRSA/PSS`
	//
	// - `SHA256withECDSA`
	//
	// - `SM3withSM2`
	//
	// - `SHA256withRSA/P7`
	//
	// This parameter is required.
	//
	// example:
	//
	// SHA256withRSA
	SigningAlgorithm *string `json:"SigningAlgorithm,omitempty" xml:"SigningAlgorithm,omitempty"`
	// The repository ID.
	//
	// > You can get this ID by calling the [ListCertWarehouse](https://help.aliyun.com/document_detail/455805.html) operation.
	//
	// example:
	//
	// 1
	WarehouseId *int64 `json:"WarehouseId,omitempty" xml:"WarehouseId,omitempty"`
}

func (s SignRequest) String() string {
	return dara.Prettify(s)
}

func (s SignRequest) GoString() string {
	return s.String()
}

func (s *SignRequest) GetCertIdentifier() *string {
	return s.CertIdentifier
}

func (s *SignRequest) GetCustomIdentifier() *string {
	return s.CustomIdentifier
}

func (s *SignRequest) GetMessage() *string {
	return s.Message
}

func (s *SignRequest) GetMessageType() *string {
	return s.MessageType
}

func (s *SignRequest) GetSigningAlgorithm() *string {
	return s.SigningAlgorithm
}

func (s *SignRequest) GetWarehouseId() *int64 {
	return s.WarehouseId
}

func (s *SignRequest) SetCertIdentifier(v string) *SignRequest {
	s.CertIdentifier = &v
	return s
}

func (s *SignRequest) SetCustomIdentifier(v string) *SignRequest {
	s.CustomIdentifier = &v
	return s
}

func (s *SignRequest) SetMessage(v string) *SignRequest {
	s.Message = &v
	return s
}

func (s *SignRequest) SetMessageType(v string) *SignRequest {
	s.MessageType = &v
	return s
}

func (s *SignRequest) SetSigningAlgorithm(v string) *SignRequest {
	s.SigningAlgorithm = &v
	return s
}

func (s *SignRequest) SetWarehouseId(v int64) *SignRequest {
	s.WarehouseId = &v
	return s
}

func (s *SignRequest) Validate() error {
	return dara.Validate(s)
}
