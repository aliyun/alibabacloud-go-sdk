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
  // Specifies whether to enable the dry run feature.
  // 
  // - true: enables the dry run feature.
  // 
  // - false (default): disables the dry run feature.
  // 
  // The dry run feature is used to test API calls and verify the permissions on the resources that you have and the validity of the request parameters. You can view the check results in the response.
  // 
  // - DryRunOperationError: The permissions and parameters are valid. If you do not specify the DryRun parameter, the request is successful.
  // 
  // - ValidationError: The parameters in the request are invalid.
  // 
  // - AccessDeniedError: You are not authorized to perform this operation on the KMS resource.
  // 
  // example:
  // 
  // false
  DryRun *string `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
  // A JSON string that consists of key-value pairs. If you specify this parameter, you must specify the same parameter when you call the Decrypt operation. For more information, see [EncryptionContext](https://help.aliyun.com/document_detail/42975.html).
  // 
  // example:
  // 
  // {"Example":"Example"}
  EncryptionContextShrink *string `json:"EncryptionContext,omitempty" xml:"EncryptionContext,omitempty"`
  // The ID of the key. You can also specify the alias or Amazon Resource Name (ARN) of the key. For more information about aliases, see [Manage aliases](https://help.aliyun.com/document_detail/480655.html).
  // 
  // > When you access a key in another Alibaba Cloud account, you must specify the ARN of the key. The ARN of a key is in the `acs:kms:${region}:${account}:key/${keyid}` format.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // key-hzz630494463ejqjx****
  KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
  // The plaintext to be encrypted. The plaintext must be Base64-encoded.
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

