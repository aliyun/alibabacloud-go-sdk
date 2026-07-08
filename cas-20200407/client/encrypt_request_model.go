// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEncryptRequest interface {
  dara.Model
  String() string
  GoString() string
  SetAlgorithm(v string) *EncryptRequest
  GetAlgorithm() *string 
  SetCertIdentifier(v string) *EncryptRequest
  GetCertIdentifier() *string 
  SetCustomIdentifier(v string) *EncryptRequest
  GetCustomIdentifier() *string 
  SetMessageType(v string) *EncryptRequest
  GetMessageType() *string 
  SetPlaintext(v string) *EncryptRequest
  GetPlaintext() *string 
  SetWarehouseId(v int64) *EncryptRequest
  GetWarehouseId() *int64 
}

type EncryptRequest struct {
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
  // The unique identifier of the certificate. To obtain this parameter, call the [ListCert](https://help.aliyun.com/document_detail/455806.html) operation.
  // 
  // - The identifier of an SSL certificate is usually in the {Certificate ID}-cn-hangzhou format.
  // 
  // - For a private certificate authority (PCA) certificate, this is the value of the Identifier field of the private certificate.
  // 
  // example:
  // 
  // 1ef1da5f-38ed-69b3-****-037781890265
  CertIdentifier *string `json:"CertIdentifier,omitempty" xml:"CertIdentifier,omitempty"`
  // The custom identifier, which serves as a unique key.
  // 
  // example:
  // 
  // ****6bb538d538c70c01f81dg3****
  CustomIdentifier *string `json:"CustomIdentifier,omitempty" xml:"CustomIdentifier,omitempty"`
  // The message type. Valid values:
  // 
  // - RAW (default): Directly encrypts the value of Plaintext.
  // 
  // - Base64: Decodes the Base64-encoded value of Plaintext and then encrypts the decoded data.
  // 
  // example:
  // 
  // RAW
  MessageType *string `json:"MessageType,omitempty" xml:"MessageType,omitempty"`
  // The data to encrypt. The data can be plaintext or Base64-encoded plaintext. For more information, see the MessageType parameter. If you use Base64 encoding, for example, if the hexadecimal content of the data to be encrypted is `[0x31, 0x32, 0x33, 0x34]`, the corresponding Base64-encoded string is MTIzNA==. The maximum size of Plaintext depends on the Algorithm:
  // 
  // - **RSAES_OAEP_SHA_1**: 214 bytes.
  // 
  // - **RSAES_OAEP_SHA_256**: 190 bytes.
  // 
  // - **SM2PKE**: 6047 bytes.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // 1234***
  Plaintext *string `json:"Plaintext,omitempty" xml:"Plaintext,omitempty"`
  // The repository ID.
  // 
  // > To obtain this ID, call the [ListCertWarehouse](https://help.aliyun.com/document_detail/455805.html) operation.
  // 
  // example:
  // 
  // 12
  WarehouseId *int64 `json:"WarehouseId,omitempty" xml:"WarehouseId,omitempty"`
}

func (s EncryptRequest) String() string {
  return dara.Prettify(s)
}

func (s EncryptRequest) GoString() string {
  return s.String()
}

func (s *EncryptRequest) GetAlgorithm() *string  {
  return s.Algorithm
}

func (s *EncryptRequest) GetCertIdentifier() *string  {
  return s.CertIdentifier
}

func (s *EncryptRequest) GetCustomIdentifier() *string  {
  return s.CustomIdentifier
}

func (s *EncryptRequest) GetMessageType() *string  {
  return s.MessageType
}

func (s *EncryptRequest) GetPlaintext() *string  {
  return s.Plaintext
}

func (s *EncryptRequest) GetWarehouseId() *int64  {
  return s.WarehouseId
}

func (s *EncryptRequest) SetAlgorithm(v string) *EncryptRequest {
  s.Algorithm = &v
  return s
}

func (s *EncryptRequest) SetCertIdentifier(v string) *EncryptRequest {
  s.CertIdentifier = &v
  return s
}

func (s *EncryptRequest) SetCustomIdentifier(v string) *EncryptRequest {
  s.CustomIdentifier = &v
  return s
}

func (s *EncryptRequest) SetMessageType(v string) *EncryptRequest {
  s.MessageType = &v
  return s
}

func (s *EncryptRequest) SetPlaintext(v string) *EncryptRequest {
  s.Plaintext = &v
  return s
}

func (s *EncryptRequest) SetWarehouseId(v int64) *EncryptRequest {
  s.WarehouseId = &v
  return s
}

func (s *EncryptRequest) Validate() error {
  return dara.Validate(s)
}

