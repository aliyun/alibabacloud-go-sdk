// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateDataKeyShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDryRun(v string) *GenerateDataKeyShrinkRequest
	GetDryRun() *string
	SetEncryptionContextShrink(v string) *GenerateDataKeyShrinkRequest
	GetEncryptionContextShrink() *string
	SetKeyId(v string) *GenerateDataKeyShrinkRequest
	GetKeyId() *string
	SetKeySpec(v string) *GenerateDataKeyShrinkRequest
	GetKeySpec() *string
	SetNumberOfBytes(v int32) *GenerateDataKeyShrinkRequest
	GetNumberOfBytes() *int32
	SetRecipient(v string) *GenerateDataKeyShrinkRequest
	GetRecipient() *string
}

type GenerateDataKeyShrinkRequest struct {
	// Specifies whether to enable DryRun mode.
	//
	// - true: enabled.
	//
	// - false (default): disabled.
	//
	// DryRun mode is used to test API calls, verify that you have the required permissions on the relevant resources, and check whether request parameters are configured correctly. When DryRun mode is enabled, KMS always returns a failure with a reason. Possible failure reasons include:
	//
	// - DryRunOperationError: The request would succeed if the DryRun parameter were not set.
	//
	// - ValidationError: A parameter specified in the request is invalid.
	//
	// - AccessDeniedError: You do not have permission to perform the operation on the KMS resource.
	//
	// example:
	//
	// false
	DryRun *string `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// A JSON string of key-value pairs.
	//
	// If you specify this parameter, you must provide the same parameter when you call the [Decrypt](https://help.aliyun.com/document_detail/28950.html) operation. For more information, see [EncryptionContext](https://help.aliyun.com/document_detail/42975.html).
	//
	// example:
	//
	// {"Example":"Example"}
	EncryptionContextShrink *string `json:"EncryptionContext,omitempty" xml:"EncryptionContext,omitempty"`
	// The ID of the key. You can also specify a key alias or a key Amazon Resource Name (ARN). For more information about aliases, see [Manage key aliases](https://help.aliyun.com/document_detail/480655.html).
	//
	// >To access a key in another Alibaba Cloud account, you must specify the key ARN. The key ARN format is `acs:kms:${region}:${account}:key/${keyid}`.
	//
	// This parameter is required.
	//
	// example:
	//
	// key-hzz630494463ejqjx****
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The length of the data key (DK) to generate. Valid values:
	//
	// - AES_256: a 256-bit symmetric key.
	//
	// - AES_128: a 128-bit symmetric key.
	//
	// > Use KeySpec or NumberOfBytes to specify the DK length. If neither is specified, KMS generates a 256-bit DK. If both are specified, KMS ignores the KeySpec parameter.
	//
	// example:
	//
	// AES_256
	KeySpec *string `json:"KeySpec,omitempty" xml:"KeySpec,omitempty"`
	// The length of the data key (DK) to generate, in bytes.
	//
	// Valid values: 1 to 1024.
	//
	// Default value:
	//
	// - When KeySpec is set to AES_256, the default value of NumberOfBytes is 32.
	//
	// - When KeySpec is set to AES_128, the default value of NumberOfBytes is 16.
	//
	// example:
	//
	// 256
	NumberOfBytes *int32 `json:"NumberOfBytes,omitempty" xml:"NumberOfBytes,omitempty"`
	// example:
	//
	// {
	//
	// "AttestationDocument":"base64-encoded-attestion-document",
	//
	// "KeyEncryptionAlgorithm":"RSAES_OAEP_SHA_256"
	//
	// }
	Recipient *string `json:"Recipient,omitempty" xml:"Recipient,omitempty"`
}

func (s GenerateDataKeyShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GenerateDataKeyShrinkRequest) GoString() string {
	return s.String()
}

func (s *GenerateDataKeyShrinkRequest) GetDryRun() *string {
	return s.DryRun
}

func (s *GenerateDataKeyShrinkRequest) GetEncryptionContextShrink() *string {
	return s.EncryptionContextShrink
}

func (s *GenerateDataKeyShrinkRequest) GetKeyId() *string {
	return s.KeyId
}

func (s *GenerateDataKeyShrinkRequest) GetKeySpec() *string {
	return s.KeySpec
}

func (s *GenerateDataKeyShrinkRequest) GetNumberOfBytes() *int32 {
	return s.NumberOfBytes
}

func (s *GenerateDataKeyShrinkRequest) GetRecipient() *string {
	return s.Recipient
}

func (s *GenerateDataKeyShrinkRequest) SetDryRun(v string) *GenerateDataKeyShrinkRequest {
	s.DryRun = &v
	return s
}

func (s *GenerateDataKeyShrinkRequest) SetEncryptionContextShrink(v string) *GenerateDataKeyShrinkRequest {
	s.EncryptionContextShrink = &v
	return s
}

func (s *GenerateDataKeyShrinkRequest) SetKeyId(v string) *GenerateDataKeyShrinkRequest {
	s.KeyId = &v
	return s
}

func (s *GenerateDataKeyShrinkRequest) SetKeySpec(v string) *GenerateDataKeyShrinkRequest {
	s.KeySpec = &v
	return s
}

func (s *GenerateDataKeyShrinkRequest) SetNumberOfBytes(v int32) *GenerateDataKeyShrinkRequest {
	s.NumberOfBytes = &v
	return s
}

func (s *GenerateDataKeyShrinkRequest) SetRecipient(v string) *GenerateDataKeyShrinkRequest {
	s.Recipient = &v
	return s
}

func (s *GenerateDataKeyShrinkRequest) Validate() error {
	return dara.Validate(s)
}
