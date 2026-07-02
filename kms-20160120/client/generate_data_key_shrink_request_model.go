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
	// Specifies whether to enable the DryRun mode.
	//
	// - true: enables the DryRun mode.
	//
	// - false (default): disables the DryRun mode.
	//
	// The DryRun mode is used to test the API call. It verifies whether you have the permissions to access the specified resources and whether the request parameters are valid. If you enable the DryRun mode, KMS always returns a failure response and a failure reason. The failure reasons include the following:
	//
	// - DryRunOperationError: The request succeeds if the DryRun parameter is not specified.
	//
	// - ValidationError: The parameters specified in the request are invalid.
	//
	// - AccessDeniedError: You are not authorized to perform this operation on the KMS resource.
	//
	// example:
	//
	// false
	DryRun *string `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// A JSON string that consists of key-value pairs.
	//
	// If you specify this parameter, you must also specify the same parameter when you call the [Decrypt](https://help.aliyun.com/document_detail/28950.html) operation. For more information, see [EncryptionContext](https://help.aliyun.com/document_detail/42975.html).
	//
	// example:
	//
	// {"Example":"Example"}
	EncryptionContextShrink *string `json:"EncryptionContext,omitempty" xml:"EncryptionContext,omitempty"`
	// The ID of the key. You can also specify the alias or the key resource name (ARN) of the key. For more information about aliases, see [Manage aliases](https://help.aliyun.com/document_detail/480655.html).
	//
	// > When you access a key in another Alibaba Cloud account, you must enter the ARN of the key. The key ARN is in the format of `acs:kms:${region}:${account}:key/${keyid}`.
	//
	// This parameter is required.
	//
	// example:
	//
	// key-hzz630494463ejqjx****
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The length of the data key to be generated. Valid values:
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
	// The length of the data key that you want to generate. Unit: bytes.
	//
	// Valid values: 1 to 1024.
	//
	// Default values:
	//
	// - If you set KeySpec to AES_256, the default value of NumberOfBytes is 32.
	//
	// - If you set KeySpec to AES_128, the default value of NumberOfBytes is 16.
	//
	// example:
	//
	// 256
	NumberOfBytes *int32  `json:"NumberOfBytes,omitempty" xml:"NumberOfBytes,omitempty"`
	Recipient     *string `json:"Recipient,omitempty" xml:"Recipient,omitempty"`
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
