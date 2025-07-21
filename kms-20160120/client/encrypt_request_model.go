// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEncryptRequest interface {
  dara.Model
  String() string
  GoString() string
  SetDryRun(v string) *EncryptRequest
  GetDryRun() *string 
  SetEncryptionContext(v map[string]interface{}) *EncryptRequest
  GetEncryptionContext() map[string]interface{} 
  SetKeyId(v string) *EncryptRequest
  GetKeyId() *string 
  SetPlaintext(v string) *EncryptRequest
  GetPlaintext() *string 
}

type EncryptRequest struct {
  DryRun *string `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
  // A JSON string that consists of key-value pairs. If you specify this parameter, an equivalent value is required when you call the Decrypt operation. For more information, see [EncryptionContext](https://help.aliyun.com/document_detail/42975.html).
  // 
  // example:
  // 
  // {"Example":"Example"}
  EncryptionContext map[string]interface{} `json:"EncryptionContext,omitempty" xml:"EncryptionContext,omitempty"`
  // The globally unique ID of the CMK. You can also set this parameter to an alias that is bound to the CMK. For more information, see [Use aliases](https://help.aliyun.com/document_detail/68522.html).
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // 1234abcd-12ab-34cd-56ef-12345678****
  KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
  // The plaintext to be encrypted. The plaintext must be Base64 encoded.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // SGVsbG8gd29y****
  Plaintext *string `json:"Plaintext,omitempty" xml:"Plaintext,omitempty"`
}

func (s EncryptRequest) String() string {
  return dara.Prettify(s)
}

func (s EncryptRequest) GoString() string {
  return s.String()
}

func (s *EncryptRequest) GetDryRun() *string  {
  return s.DryRun
}

func (s *EncryptRequest) GetEncryptionContext() map[string]interface{}  {
  return s.EncryptionContext
}

func (s *EncryptRequest) GetKeyId() *string  {
  return s.KeyId
}

func (s *EncryptRequest) GetPlaintext() *string  {
  return s.Plaintext
}

func (s *EncryptRequest) SetDryRun(v string) *EncryptRequest {
  s.DryRun = &v
  return s
}

func (s *EncryptRequest) SetEncryptionContext(v map[string]interface{}) *EncryptRequest {
  s.EncryptionContext = v
  return s
}

func (s *EncryptRequest) SetKeyId(v string) *EncryptRequest {
  s.KeyId = &v
  return s
}

func (s *EncryptRequest) SetPlaintext(v string) *EncryptRequest {
  s.Plaintext = &v
  return s
}

func (s *EncryptRequest) Validate() error {
  return dara.Validate(s)
}

