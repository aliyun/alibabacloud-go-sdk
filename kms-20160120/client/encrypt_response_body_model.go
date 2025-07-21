// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEncryptResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCiphertextBlob(v string) *EncryptResponseBody
  GetCiphertextBlob() *string 
  SetKeyId(v string) *EncryptResponseBody
  GetKeyId() *string 
  SetKeyVersionId(v string) *EncryptResponseBody
  GetKeyVersionId() *string 
  SetRequestId(v string) *EncryptResponseBody
  GetRequestId() *string 
}

type EncryptResponseBody struct {
  // The ciphertext of the data that is encrypted by using the primary CMK version.
  // 
  // example:
  // 
  // DZhOWVmZDktM2QxNi00ODk0LWJkNGYtMWZjNDNmM2YyYWJmaaSl+TztSIMe43nbTH/Z1Wr4XfLftKhAciUmDQXuMRl4WTvKhxjMThjK****
  CiphertextBlob *string `json:"CiphertextBlob,omitempty" xml:"CiphertextBlob,omitempty"`
  // The globally unique ID of the CMK. If you set the KeyId parameter to an alias, the ID of the CMK to which the alias is bound is returned.
  // 
  // example:
  // 
  // 1234abcd-12ab-34cd-56ef-12345678****
  KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
  // The ID of the key version that is used to encrypt the plaintext. It is the primary version of the CMK.
  // 
  // example:
  // 
  // 86a9efd9-3d16-4894-bd4f-1fc43f3f****
  KeyVersionId *string `json:"KeyVersionId,omitempty" xml:"KeyVersionId,omitempty"`
  // The ID of the request.
  // 
  // example:
  // 
  // 475f1620-b9d3-4d35-b5c6-3fbdd941423d
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EncryptResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EncryptResponseBody) GoString() string {
  return s.String()
}

func (s *EncryptResponseBody) GetCiphertextBlob() *string  {
  return s.CiphertextBlob
}

func (s *EncryptResponseBody) GetKeyId() *string  {
  return s.KeyId
}

func (s *EncryptResponseBody) GetKeyVersionId() *string  {
  return s.KeyVersionId
}

func (s *EncryptResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EncryptResponseBody) SetCiphertextBlob(v string) *EncryptResponseBody {
  s.CiphertextBlob = &v
  return s
}

func (s *EncryptResponseBody) SetKeyId(v string) *EncryptResponseBody {
  s.KeyId = &v
  return s
}

func (s *EncryptResponseBody) SetKeyVersionId(v string) *EncryptResponseBody {
  s.KeyVersionId = &v
  return s
}

func (s *EncryptResponseBody) SetRequestId(v string) *EncryptResponseBody {
  s.RequestId = &v
  return s
}

func (s *EncryptResponseBody) Validate() error {
  return dara.Validate(s)
}

