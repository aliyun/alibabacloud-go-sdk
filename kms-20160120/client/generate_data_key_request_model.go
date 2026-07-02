// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateDataKeyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDryRun(v string) *GenerateDataKeyRequest
	GetDryRun() *string
	SetEncryptionContext(v map[string]interface{}) *GenerateDataKeyRequest
	GetEncryptionContext() map[string]interface{}
	SetKeyId(v string) *GenerateDataKeyRequest
	GetKeyId() *string
	SetKeySpec(v string) *GenerateDataKeyRequest
	GetKeySpec() *string
	SetNumberOfBytes(v int32) *GenerateDataKeyRequest
	GetNumberOfBytes() *int32
	SetRecipient(v string) *GenerateDataKeyRequest
	GetRecipient() *string
}

type GenerateDataKeyRequest struct {
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
	EncryptionContext map[string]interface{} `json:"EncryptionContext,omitempty" xml:"EncryptionContext,omitempty"`
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

func (s GenerateDataKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s GenerateDataKeyRequest) GoString() string {
	return s.String()
}

func (s *GenerateDataKeyRequest) GetDryRun() *string {
	return s.DryRun
}

func (s *GenerateDataKeyRequest) GetEncryptionContext() map[string]interface{} {
	return s.EncryptionContext
}

func (s *GenerateDataKeyRequest) GetKeyId() *string {
	return s.KeyId
}

func (s *GenerateDataKeyRequest) GetKeySpec() *string {
	return s.KeySpec
}

func (s *GenerateDataKeyRequest) GetNumberOfBytes() *int32 {
	return s.NumberOfBytes
}

func (s *GenerateDataKeyRequest) GetRecipient() *string {
	return s.Recipient
}

func (s *GenerateDataKeyRequest) SetDryRun(v string) *GenerateDataKeyRequest {
	s.DryRun = &v
	return s
}

func (s *GenerateDataKeyRequest) SetEncryptionContext(v map[string]interface{}) *GenerateDataKeyRequest {
	s.EncryptionContext = v
	return s
}

func (s *GenerateDataKeyRequest) SetKeyId(v string) *GenerateDataKeyRequest {
	s.KeyId = &v
	return s
}

func (s *GenerateDataKeyRequest) SetKeySpec(v string) *GenerateDataKeyRequest {
	s.KeySpec = &v
	return s
}

func (s *GenerateDataKeyRequest) SetNumberOfBytes(v int32) *GenerateDataKeyRequest {
	s.NumberOfBytes = &v
	return s
}

func (s *GenerateDataKeyRequest) SetRecipient(v string) *GenerateDataKeyRequest {
	s.Recipient = &v
	return s
}

func (s *GenerateDataKeyRequest) Validate() error {
	return dara.Validate(s)
}
