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
  EncryptionContext map[string]interface{} `json:"EncryptionContext,omitempty" xml:"EncryptionContext,omitempty"`
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

