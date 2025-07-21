// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEncryptShrinkRequest interface {
  dara.Model
  String() string
  GoString() string
  SetDryRun(v string) *EncryptShrinkRequest
  GetDryRun() *string 
  SetEncryptionContextShrink(v string) *EncryptShrinkRequest
  GetEncryptionContextShrink() *string 
  SetKeyId(v string) *EncryptShrinkRequest
  GetKeyId() *string 
  SetPlaintext(v string) *EncryptShrinkRequest
  GetPlaintext() *string 
}

type EncryptShrinkRequest struct {
  DryRun *string `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
  // A JSON string that consists of key-value pairs. If you specify this parameter, an equivalent value is required when you call the Decrypt operation. For more information, see [EncryptionContext](https://help.aliyun.com/document_detail/42975.html).
  // 
  // example:
  // 
  // {"Example":"Example"}
  EncryptionContextShrink *string `json:"EncryptionContext,omitempty" xml:"EncryptionContext,omitempty"`
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

func (s EncryptShrinkRequest) String() string {
  return dara.Prettify(s)
}

func (s EncryptShrinkRequest) GoString() string {
  return s.String()
}

func (s *EncryptShrinkRequest) GetDryRun() *string  {
  return s.DryRun
}

func (s *EncryptShrinkRequest) GetEncryptionContextShrink() *string  {
  return s.EncryptionContextShrink
}

func (s *EncryptShrinkRequest) GetKeyId() *string  {
  return s.KeyId
}

func (s *EncryptShrinkRequest) GetPlaintext() *string  {
  return s.Plaintext
}

func (s *EncryptShrinkRequest) SetDryRun(v string) *EncryptShrinkRequest {
  s.DryRun = &v
  return s
}

func (s *EncryptShrinkRequest) SetEncryptionContextShrink(v string) *EncryptShrinkRequest {
  s.EncryptionContextShrink = &v
  return s
}

func (s *EncryptShrinkRequest) SetKeyId(v string) *EncryptShrinkRequest {
  s.KeyId = &v
  return s
}

func (s *EncryptShrinkRequest) SetPlaintext(v string) *EncryptShrinkRequest {
  s.Plaintext = &v
  return s
}

func (s *EncryptShrinkRequest) Validate() error {
  return dara.Validate(s)
}

