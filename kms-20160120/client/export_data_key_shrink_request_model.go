// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportDataKeyShrinkRequest interface {
  dara.Model
  String() string
  GoString() string
  SetCiphertextBlob(v string) *ExportDataKeyShrinkRequest
  GetCiphertextBlob() *string 
  SetDryRun(v string) *ExportDataKeyShrinkRequest
  GetDryRun() *string 
  SetEncryptionContextShrink(v string) *ExportDataKeyShrinkRequest
  GetEncryptionContextShrink() *string 
  SetPublicKeyBlob(v string) *ExportDataKeyShrinkRequest
  GetPublicKeyBlob() *string 
  SetWrappingAlgorithm(v string) *ExportDataKeyShrinkRequest
  GetWrappingAlgorithm() *string 
  SetWrappingKeySpec(v string) *ExportDataKeyShrinkRequest
  GetWrappingKeySpec() *string 
}

type ExportDataKeyShrinkRequest struct {
  // The ciphertext of the data key that is encrypted using a master key (CMK).
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // ODZhOWVmZDktM2QxNi00ODk0LWJkNGYtMWZjNDNmM2YyYWJmS7FmDBBQ0BkKsQrtRnidtPwirmDcS0ZuJCU41xxAAWk4Z8qsADfbV0b+i6kQmlvj79dJdGOvtX69Uycs901q********
  CiphertextBlob *string `json:"CiphertextBlob,omitempty" xml:"CiphertextBlob,omitempty"`
  // Specifies whether to enable the DryRun mode.
  // 
  // - true
  // 
  // - false (default)
  // 
  // The DryRun mode is used to test the API call and verify the permissions on the specified resources and the validity of the request parameters. If you enable the DryRun mode, KMS returns a failure response and a failure reason. The failure reasons include the following:
  // 
  // - DryRunOperationError: The request would have succeeded if the DryRun parameter was not specified.
  // 
  // - ValidationError: The specified parameters in the request are invalid.
  // 
  // - AccessDeniedError: You are not authorized to perform the operation on the KMS resource.
  // 
  // example:
  // 
  // false
  DryRun *string `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
  // A JSON string that consists of key-value pairs. EncryptionContext is the encryption context that is passed in when the data key is encrypted using a CMK. For more information, see [EncryptionContext](https://help.aliyun.com/document_detail/42975.html).
  // 
  // example:
  // 
  // {"Example":"Example"}
  EncryptionContextShrink *string `json:"EncryptionContext,omitempty" xml:"EncryptionContext,omitempty"`
  // The public key in Base64 format.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAndKfC2ReLL2+y8a0+ZBBeAft/uBYo86GZiYJuflqgUzKxpyuvlo3uQkBv6b+nx+0tz8g8v7GhpPWMSW5L9mNHYsvYFsa7jTxsYdt17yj6GlUHPuMIs8hr5qbwl38IHU1iIa7nYWwE2fb3ePOvLDACRJVgGpU0yxioW80d2QD+9aU4jF5dlAahcfgsNzo2CXzCUc1+xbmNuq7Rp+H9VJB9dyYOwqnW3RhOLBo21FzpORapf0UiRlrHRpk1V6ez+aE1dofaYh/9bh0m6ioxj7j5hpZbWccuEZTMBKd+cbuBkRhJzc6Tti6qwZbDiu4fUwbZS0Tqpuo1UadiyxMW********
  PublicKeyBlob *string `json:"PublicKeyBlob,omitempty" xml:"PublicKeyBlob,omitempty"`
  // The encryption algorithm that is used to encrypt the data key using the public key specified by PublicKeyBlob. For more information about the algorithms, see [AsymmetricDecrypt](https://help.aliyun.com/document_detail/148130.html).<br> Valid values:<br><br>
  // 
  // - RSAES_OAEP_SHA_256
  // 
  // - RSAES_OAEP_SHA_1
  // 
  // - SM2PKE
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // RSAES_OAEP_SHA_256
  WrappingAlgorithm *string `json:"WrappingAlgorithm,omitempty" xml:"WrappingAlgorithm,omitempty"`
  // The type of the public key specified by PublicKeyBlob. For more information about key types, see [Introduction to asymmetric keys](https://help.aliyun.com/document_detail/148147.html).<br> Valid values:<br><br>
  // 
  // - RSA_2048
  // 
  // - EC_SM2
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // RSA_2048
  WrappingKeySpec *string `json:"WrappingKeySpec,omitempty" xml:"WrappingKeySpec,omitempty"`
}

func (s ExportDataKeyShrinkRequest) String() string {
  return dara.Prettify(s)
}

func (s ExportDataKeyShrinkRequest) GoString() string {
  return s.String()
}

func (s *ExportDataKeyShrinkRequest) GetCiphertextBlob() *string  {
  return s.CiphertextBlob
}

func (s *ExportDataKeyShrinkRequest) GetDryRun() *string  {
  return s.DryRun
}

func (s *ExportDataKeyShrinkRequest) GetEncryptionContextShrink() *string  {
  return s.EncryptionContextShrink
}

func (s *ExportDataKeyShrinkRequest) GetPublicKeyBlob() *string  {
  return s.PublicKeyBlob
}

func (s *ExportDataKeyShrinkRequest) GetWrappingAlgorithm() *string  {
  return s.WrappingAlgorithm
}

func (s *ExportDataKeyShrinkRequest) GetWrappingKeySpec() *string  {
  return s.WrappingKeySpec
}

func (s *ExportDataKeyShrinkRequest) SetCiphertextBlob(v string) *ExportDataKeyShrinkRequest {
  s.CiphertextBlob = &v
  return s
}

func (s *ExportDataKeyShrinkRequest) SetDryRun(v string) *ExportDataKeyShrinkRequest {
  s.DryRun = &v
  return s
}

func (s *ExportDataKeyShrinkRequest) SetEncryptionContextShrink(v string) *ExportDataKeyShrinkRequest {
  s.EncryptionContextShrink = &v
  return s
}

func (s *ExportDataKeyShrinkRequest) SetPublicKeyBlob(v string) *ExportDataKeyShrinkRequest {
  s.PublicKeyBlob = &v
  return s
}

func (s *ExportDataKeyShrinkRequest) SetWrappingAlgorithm(v string) *ExportDataKeyShrinkRequest {
  s.WrappingAlgorithm = &v
  return s
}

func (s *ExportDataKeyShrinkRequest) SetWrappingKeySpec(v string) *ExportDataKeyShrinkRequest {
  s.WrappingKeySpec = &v
  return s
}

func (s *ExportDataKeyShrinkRequest) Validate() error {
  return dara.Validate(s)
}

