// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateAndExportDataKeyShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDryRun(v string) *GenerateAndExportDataKeyShrinkRequest
	GetDryRun() *string
	SetEncryptionContextShrink(v string) *GenerateAndExportDataKeyShrinkRequest
	GetEncryptionContextShrink() *string
	SetKeyId(v string) *GenerateAndExportDataKeyShrinkRequest
	GetKeyId() *string
	SetKeySpec(v string) *GenerateAndExportDataKeyShrinkRequest
	GetKeySpec() *string
	SetNumberOfBytes(v int32) *GenerateAndExportDataKeyShrinkRequest
	GetNumberOfBytes() *int32
	SetPublicKeyBlob(v string) *GenerateAndExportDataKeyShrinkRequest
	GetPublicKeyBlob() *string
	SetWrappingAlgorithm(v string) *GenerateAndExportDataKeyShrinkRequest
	GetWrappingAlgorithm() *string
	SetWrappingKeySpec(v string) *GenerateAndExportDataKeyShrinkRequest
	GetWrappingKeySpec() *string
}

type GenerateAndExportDataKeyShrinkRequest struct {
	// Specifies whether to enable the dry run feature.
	//
	// - true: enables the feature.
	//
	// - false (default): disables the feature.
	//
	// The DryRun mode is used to test API calls and verify the permissions on the resources that you have access to and the validity of the request parameters. If you enable the DryRun mode, KMS always returns a failure response and the cause of the failure. The following failure causes are included:
	//
	// - DryRunOperationError: The request would have succeeded if the DryRun parameter is not specified.
	//
	// - ValidationError: The parameters specified in the request are invalid.
	//
	// - AccessDeniedError: You are not authorized to perform this operation on the KMS resource.
	//
	// example:
	//
	// false
	DryRun *string `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// A JSON string that consists of key-value pairs. If you specify this parameter, you must specify the same parameter when you call the Decrypt operation or other operations to re-encrypt the data key. For more information, see [EncryptionContext](https://help.aliyun.com/document_detail/42975.html).
	//
	// example:
	//
	// {"Example":"Example"}
	EncryptionContextShrink *string `json:"EncryptionContext,omitempty" xml:"EncryptionContext,omitempty"`
	// The ID of the key. You can also specify the alias or Amazon Resource Name (ARN) of the key. For more information about aliases, see [Manage aliases](https://help.aliyun.com/document_detail/480655.html).
	//
	// > To access a key in another Alibaba Cloud account, you must specify the ARN of the key. The key ARN is in the format of `acs:kms:${region}:${account}:key/${keyid}`.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234abcd-12ab-34cd-56ef-12345678****
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The length of the data key that you want to generate. Valid values:
	//
	// - AES_256: a 256-bit symmetric key.
	//
	// - AES_128: a 128-bit symmetric key.
	//
	// > We recommend that you use the KeySpec or NumberOfBytes parameter to specify the length of a data key. If you do not specify either of the parameters, KMS generates a 256-bit data key. If you specify both parameters, KMS ignores the KeySpec parameter.
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
	// The public key that is encoded in Base64.
	//
	// This parameter is required.
	//
	// example:
	//
	// MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAndKfC2ReLL2+y8a0+ZBBeAft/uBYo86GZiYJuflqgUzKxpyuvlo3uQkBv6b+nx+0tz8g8v7GhpPWMSW5L9mNHYsvYFsa7jTxsYdt17yj6GlUHPuMIs8hr5qbwl38IHU1iIa7nYWwE2fb3ePOvLDACRJVgGpU0yxioW80d2QD+9aU4jF5dlAahcfgsNzo2CXzCUc1+xbmNuq7Rp+H9VJB9dyYOwqnW3RhOLBo21FzpORapf0UiRlrHRpk1V6ez+aE1dofaYh/9bh0m6ioxj7j5hpZbWccuEZTMBKd+cbuBkRhJzc6Tti6qwZbDiu4fUwbZS0Tqpuo1UadiyxMW********
	PublicKeyBlob *string `json:"PublicKeyBlob,omitempty" xml:"PublicKeyBlob,omitempty"`
	// The encryption algorithm that is used to encrypt the data key using the public key specified by PublicKeyBlob. For more information about encryption algorithms, see [AsymmetricDecrypt](https://help.aliyun.com/document_detail/148130.html).<br> Valid values:<br><br>
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
	// The type of the key specified by PublicKeyBlob. For more information about key types, see [Introduction to asymmetric keys](https://help.aliyun.com/document_detail/148147.html).<br> Valid values:<br><br>
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

func (s GenerateAndExportDataKeyShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GenerateAndExportDataKeyShrinkRequest) GoString() string {
	return s.String()
}

func (s *GenerateAndExportDataKeyShrinkRequest) GetDryRun() *string {
	return s.DryRun
}

func (s *GenerateAndExportDataKeyShrinkRequest) GetEncryptionContextShrink() *string {
	return s.EncryptionContextShrink
}

func (s *GenerateAndExportDataKeyShrinkRequest) GetKeyId() *string {
	return s.KeyId
}

func (s *GenerateAndExportDataKeyShrinkRequest) GetKeySpec() *string {
	return s.KeySpec
}

func (s *GenerateAndExportDataKeyShrinkRequest) GetNumberOfBytes() *int32 {
	return s.NumberOfBytes
}

func (s *GenerateAndExportDataKeyShrinkRequest) GetPublicKeyBlob() *string {
	return s.PublicKeyBlob
}

func (s *GenerateAndExportDataKeyShrinkRequest) GetWrappingAlgorithm() *string {
	return s.WrappingAlgorithm
}

func (s *GenerateAndExportDataKeyShrinkRequest) GetWrappingKeySpec() *string {
	return s.WrappingKeySpec
}

func (s *GenerateAndExportDataKeyShrinkRequest) SetDryRun(v string) *GenerateAndExportDataKeyShrinkRequest {
	s.DryRun = &v
	return s
}

func (s *GenerateAndExportDataKeyShrinkRequest) SetEncryptionContextShrink(v string) *GenerateAndExportDataKeyShrinkRequest {
	s.EncryptionContextShrink = &v
	return s
}

func (s *GenerateAndExportDataKeyShrinkRequest) SetKeyId(v string) *GenerateAndExportDataKeyShrinkRequest {
	s.KeyId = &v
	return s
}

func (s *GenerateAndExportDataKeyShrinkRequest) SetKeySpec(v string) *GenerateAndExportDataKeyShrinkRequest {
	s.KeySpec = &v
	return s
}

func (s *GenerateAndExportDataKeyShrinkRequest) SetNumberOfBytes(v int32) *GenerateAndExportDataKeyShrinkRequest {
	s.NumberOfBytes = &v
	return s
}

func (s *GenerateAndExportDataKeyShrinkRequest) SetPublicKeyBlob(v string) *GenerateAndExportDataKeyShrinkRequest {
	s.PublicKeyBlob = &v
	return s
}

func (s *GenerateAndExportDataKeyShrinkRequest) SetWrappingAlgorithm(v string) *GenerateAndExportDataKeyShrinkRequest {
	s.WrappingAlgorithm = &v
	return s
}

func (s *GenerateAndExportDataKeyShrinkRequest) SetWrappingKeySpec(v string) *GenerateAndExportDataKeyShrinkRequest {
	s.WrappingKeySpec = &v
	return s
}

func (s *GenerateAndExportDataKeyShrinkRequest) Validate() error {
	return dara.Validate(s)
}
