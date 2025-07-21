// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableKeyRequest interface {
  dara.Model
  String() string
  GoString() string
  SetKeyId(v string) *EnableKeyRequest
  GetKeyId() *string 
}

type EnableKeyRequest struct {
  // The globally unique ID of the CMK.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // 1234abcd-12ab-34cd-56ef-12345678****
  KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
}

func (s EnableKeyRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableKeyRequest) GoString() string {
  return s.String()
}

func (s *EnableKeyRequest) GetKeyId() *string  {
  return s.KeyId
}

func (s *EnableKeyRequest) SetKeyId(v string) *EnableKeyRequest {
  s.KeyId = &v
  return s
}

func (s *EnableKeyRequest) Validate() error {
  return dara.Validate(s)
}

