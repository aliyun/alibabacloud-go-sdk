// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportDataKeyRequest interface {
  dara.Model
  String() string
  GoString() string
  SetCiphertextBlob(v string) *ExportDataKeyRequest
  GetCiphertextBlob() *string 
  SetDryRun(v string) *ExportDataKeyRequest
  GetDryRun() *string 
  SetEncryptionContext(v map[string]interface{}) *ExportDataKeyRequest
  GetEncryptionContext() map[string]interface{} 
  SetPublicKeyBlob(v string) *ExportDataKeyRequest
  GetPublicKeyBlob() *string 
  SetWrappingAlgorithm(v string) *ExportDataKeyRequest
  GetWrappingAlgorithm() *string 
  SetWrappingKeySpec(v string) *ExportDataKeyRequest
  GetWrappingKeySpec() *string 
}

type ExportDataKeyRequest struct {
  // The ciphertext of the data key encrypted by using a CMK.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // ODZhOWVmZDktM2QxNi00ODk0LWJkNGYtMWZjNDNmM2YyYWJmS7FmDBBQ0BkKsQrtRnidtPwirmDcS0ZuJCU41xxAAWk4Z8qsADfbV0b+i6kQmlvj79dJdGOvtX69Uycs901q********
  CiphertextBlob *string `json:"CiphertextBlob,omitempty" xml:"CiphertextBlob,omitempty"`
  DryRun *string `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
  // A JSON string that consists of key-value pairs. If you specify this parameter when you use a CMK to encrypt the data key, an equivalent value is required here. For more information, see [EncryptionContext](https://help.aliyun.com/document_detail/42975.html).
  // 
  // example:
  // 
  // {"Example":"Example"}
  EncryptionContext map[string]interface{} `json:"EncryptionContext,omitempty" xml:"EncryptionContext,omitempty"`
  // A Base64-encoded public key.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAndKfC2ReLL2+y8a0+ZBBeAft/uBYo86GZiYJuflqgUzKxpyuvlo3uQkBv6b+nx+0tz8g8v7GhpPWMSW5L9mNHYsvYFsa7jTxsYdt17yj6GlUHPuMIs8hr5qbwl38IHU1iIa7nYWwE2fb3ePOvLDACRJVgGpU0yxioW80d2QD+9aU4jF5dlAahcfgsNzo2CXzCUc1+xbmNuq7Rp+H9VJB9dyYOwqnW3RhOLBo21FzpORapf0UiRlrHRpk1V6ez+aE1dofaYh/9bh0m6ioxj7j5hpZbWccuEZTMBKd+cbuBkRhJzc6Tti6qwZbDiu4fUwbZS0Tqpuo1UadiyxMW********
  PublicKeyBlob *string `json:"PublicKeyBlob,omitempty" xml:"PublicKeyBlob,omitempty"`
  // The encryption algorithm based on which you want to use the public key specified by PublicKeyBlob to encrypt the data key. For more information about encryption algorithms, see [AsymmetricDecrypt](https://help.aliyun.com/document_detail/148130.html).
  // 
  // Valid values:
  // 
  // 	- RSAES_OAEP_SHA_256
  // 
  // 	- RSAES_OAEP_SHA_1
  // 
  // 	- SM2PKE
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // RSAES_OAEP_SHA_256
  WrappingAlgorithm *string `json:"WrappingAlgorithm,omitempty" xml:"WrappingAlgorithm,omitempty"`
  // The key type of the public key specified by PublicKeyBlob. For more information about key types, see [Introduction to asymmetric keys](https://help.aliyun.com/document_detail/148147.html).
  // 
  // Valid values:
  // 
  // 	- RSA_2048
  // 
  // 	- EC_SM2
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // RSA_2048
  WrappingKeySpec *string `json:"WrappingKeySpec,omitempty" xml:"WrappingKeySpec,omitempty"`
}

func (s ExportDataKeyRequest) String() string {
  return dara.Prettify(s)
}

func (s ExportDataKeyRequest) GoString() string {
  return s.String()
}

func (s *ExportDataKeyRequest) GetCiphertextBlob() *string  {
  return s.CiphertextBlob
}

func (s *ExportDataKeyRequest) GetDryRun() *string  {
  return s.DryRun
}

func (s *ExportDataKeyRequest) GetEncryptionContext() map[string]interface{}  {
  return s.EncryptionContext
}

func (s *ExportDataKeyRequest) GetPublicKeyBlob() *string  {
  return s.PublicKeyBlob
}

func (s *ExportDataKeyRequest) GetWrappingAlgorithm() *string  {
  return s.WrappingAlgorithm
}

func (s *ExportDataKeyRequest) GetWrappingKeySpec() *string  {
  return s.WrappingKeySpec
}

func (s *ExportDataKeyRequest) SetCiphertextBlob(v string) *ExportDataKeyRequest {
  s.CiphertextBlob = &v
  return s
}

func (s *ExportDataKeyRequest) SetDryRun(v string) *ExportDataKeyRequest {
  s.DryRun = &v
  return s
}

func (s *ExportDataKeyRequest) SetEncryptionContext(v map[string]interface{}) *ExportDataKeyRequest {
  s.EncryptionContext = v
  return s
}

func (s *ExportDataKeyRequest) SetPublicKeyBlob(v string) *ExportDataKeyRequest {
  s.PublicKeyBlob = &v
  return s
}

func (s *ExportDataKeyRequest) SetWrappingAlgorithm(v string) *ExportDataKeyRequest {
  s.WrappingAlgorithm = &v
  return s
}

func (s *ExportDataKeyRequest) SetWrappingKeySpec(v string) *ExportDataKeyRequest {
  s.WrappingKeySpec = &v
  return s
}

func (s *ExportDataKeyRequest) Validate() error {
  return dara.Validate(s)
}

