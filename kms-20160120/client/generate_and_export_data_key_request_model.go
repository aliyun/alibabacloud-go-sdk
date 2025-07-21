// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateAndExportDataKeyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDryRun(v string) *GenerateAndExportDataKeyRequest
	GetDryRun() *string
	SetEncryptionContext(v map[string]interface{}) *GenerateAndExportDataKeyRequest
	GetEncryptionContext() map[string]interface{}
	SetKeyId(v string) *GenerateAndExportDataKeyRequest
	GetKeyId() *string
	SetKeySpec(v string) *GenerateAndExportDataKeyRequest
	GetKeySpec() *string
	SetNumberOfBytes(v int32) *GenerateAndExportDataKeyRequest
	GetNumberOfBytes() *int32
	SetPublicKeyBlob(v string) *GenerateAndExportDataKeyRequest
	GetPublicKeyBlob() *string
	SetWrappingAlgorithm(v string) *GenerateAndExportDataKeyRequest
	GetWrappingAlgorithm() *string
	SetWrappingKeySpec(v string) *GenerateAndExportDataKeyRequest
	GetWrappingKeySpec() *string
}

type GenerateAndExportDataKeyRequest struct {
	DryRun *string `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// A JSON string of key-value pairs. If you specify this parameter here, an equivalent value is required when you decrypt or re-encrypt the data key. For more information, see [EncryptionContext](https://help.aliyun.com/document_detail/42975.html).
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
	// The length of the data key that you want to generate. Valid values:
	//
	// 	- AES_256: a 256-bit symmetric key
	//
	// 	- AES_128: a 128-bit symmetric key
	//
	// >  We recommend that you use the KeySpec or NumberOfBytes parameter to specify the length of a data key. If both parameters are not specified, KMS generates a 256-bit data key. If both parameters are specified, KMS ignores the KeySpec parameter.
	//
	// example:
	//
	// AES_256
	KeySpec *string `json:"KeySpec,omitempty" xml:"KeySpec,omitempty"`
	// The length of the data key that you want to generate.
	//
	// Valid values: 1 to 1024.
	//
	// Unit: bytes.
	//
	// example:
	//
	// 32
	NumberOfBytes *int32 `json:"NumberOfBytes,omitempty" xml:"NumberOfBytes,omitempty"`
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

func (s GenerateAndExportDataKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s GenerateAndExportDataKeyRequest) GoString() string {
	return s.String()
}

func (s *GenerateAndExportDataKeyRequest) GetDryRun() *string {
	return s.DryRun
}

func (s *GenerateAndExportDataKeyRequest) GetEncryptionContext() map[string]interface{} {
	return s.EncryptionContext
}

func (s *GenerateAndExportDataKeyRequest) GetKeyId() *string {
	return s.KeyId
}

func (s *GenerateAndExportDataKeyRequest) GetKeySpec() *string {
	return s.KeySpec
}

func (s *GenerateAndExportDataKeyRequest) GetNumberOfBytes() *int32 {
	return s.NumberOfBytes
}

func (s *GenerateAndExportDataKeyRequest) GetPublicKeyBlob() *string {
	return s.PublicKeyBlob
}

func (s *GenerateAndExportDataKeyRequest) GetWrappingAlgorithm() *string {
	return s.WrappingAlgorithm
}

func (s *GenerateAndExportDataKeyRequest) GetWrappingKeySpec() *string {
	return s.WrappingKeySpec
}

func (s *GenerateAndExportDataKeyRequest) SetDryRun(v string) *GenerateAndExportDataKeyRequest {
	s.DryRun = &v
	return s
}

func (s *GenerateAndExportDataKeyRequest) SetEncryptionContext(v map[string]interface{}) *GenerateAndExportDataKeyRequest {
	s.EncryptionContext = v
	return s
}

func (s *GenerateAndExportDataKeyRequest) SetKeyId(v string) *GenerateAndExportDataKeyRequest {
	s.KeyId = &v
	return s
}

func (s *GenerateAndExportDataKeyRequest) SetKeySpec(v string) *GenerateAndExportDataKeyRequest {
	s.KeySpec = &v
	return s
}

func (s *GenerateAndExportDataKeyRequest) SetNumberOfBytes(v int32) *GenerateAndExportDataKeyRequest {
	s.NumberOfBytes = &v
	return s
}

func (s *GenerateAndExportDataKeyRequest) SetPublicKeyBlob(v string) *GenerateAndExportDataKeyRequest {
	s.PublicKeyBlob = &v
	return s
}

func (s *GenerateAndExportDataKeyRequest) SetWrappingAlgorithm(v string) *GenerateAndExportDataKeyRequest {
	s.WrappingAlgorithm = &v
	return s
}

func (s *GenerateAndExportDataKeyRequest) SetWrappingKeySpec(v string) *GenerateAndExportDataKeyRequest {
	s.WrappingKeySpec = &v
	return s
}

func (s *GenerateAndExportDataKeyRequest) Validate() error {
	return dara.Validate(s)
}
