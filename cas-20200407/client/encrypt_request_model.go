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
}

type EncryptRequest struct {
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
  // The unique identifier of the certificate. You can call the [ListCert](https://help.aliyun.com/document_detail/455806.html) operation to obtain the identifier.
  // 
  // 	- If the certificate is an SSL certificate, the value of this parameter must be in the {Certificate ID}-cn-hangzhou format.
  // 
  // 	- If the certificate is a private certificate, the value of this parameter must be the value of the Identifier field for the private certificate.
  // 
  // example:
  // 
  // 12345678-1234-1234-1234-12345678****
  CertIdentifier *string `json:"CertIdentifier,omitempty" xml:"CertIdentifier,omitempty"`
  CustomIdentifier *string `json:"CustomIdentifier,omitempty" xml:"CustomIdentifier,omitempty"`
  // The value type of the Message parameter. Valid values:
  // 
  // 	- RAW: The value of the Plaintext parameter is directly encrypted. This is the default value.
  // 
  // 	- Base64: The value of the Plaintext parameter is Base64-encoded data. The data is decoded and then encrypted.
  // 
  // example:
  // 
  // RAW
  MessageType *string `json:"MessageType,omitempty" xml:"MessageType,omitempty"`
  // The data that you want to encrypt. The value of this parameter can be raw data or Base64-encoded data. For more information, see the description of the MessageType parameter. For example, if the hexadecimal data that you want to encrypt is `[0x31, 0x32, 0x33, 0x34]`, the Base64-encoded data is MTIzNA==. The size of data that can be encrypted varies based on the encryption algorithm that you use. The following list describes the relationship between the encryption algorithms and data sizes:
  // 
  // 	- **RSAES_OAEP_SHA_1**: 214 bytes
  // 
  // 	- **RSAES_OAEP_SHA_256**: 190 bytes
  // 
  // 	- **SM2PKE**: 6,047 bytes
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // 1234***
  Plaintext *string `json:"Plaintext,omitempty" xml:"Plaintext,omitempty"`
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

func (s *EncryptRequest) Validate() error {
  return dara.Validate(s)
}

