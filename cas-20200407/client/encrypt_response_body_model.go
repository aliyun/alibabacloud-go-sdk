// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEncryptResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCertIdentifier(v string) *EncryptResponseBody
  GetCertIdentifier() *string 
  SetCiphertextBlob(v string) *EncryptResponseBody
  GetCiphertextBlob() *string 
  SetRequestId(v string) *EncryptResponseBody
  GetRequestId() *string 
}

type EncryptResponseBody struct {
  // The unique identifier of the certificate.
  // 
  // example:
  // 
  // 12345678-1234-1234-1234-12345678****
  CertIdentifier *string `json:"CertIdentifier,omitempty" xml:"CertIdentifier,omitempty"`
  // The data after encryption. The value is encoded in Base64.
  // 
  // example:
  // 
  // ZOyIygCyaOW6Gj****MlNKiuyjfzw=
  CiphertextBlob *string `json:"CiphertextBlob,omitempty" xml:"CiphertextBlob,omitempty"`
  // The ID of the request.
  // 
  // example:
  // 
  // 5979d897-d69f-4fc9-87dd-f3bb73c40b80
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EncryptResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EncryptResponseBody) GoString() string {
  return s.String()
}

func (s *EncryptResponseBody) GetCertIdentifier() *string  {
  return s.CertIdentifier
}

func (s *EncryptResponseBody) GetCiphertextBlob() *string  {
  return s.CiphertextBlob
}

func (s *EncryptResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EncryptResponseBody) SetCertIdentifier(v string) *EncryptResponseBody {
  s.CertIdentifier = &v
  return s
}

func (s *EncryptResponseBody) SetCiphertextBlob(v string) *EncryptResponseBody {
  s.CiphertextBlob = &v
  return s
}

func (s *EncryptResponseBody) SetRequestId(v string) *EncryptResponseBody {
  s.RequestId = &v
  return s
}

func (s *EncryptResponseBody) Validate() error {
  return dara.Validate(s)
}

