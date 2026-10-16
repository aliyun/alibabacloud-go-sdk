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
	EncryptionContext map[string]interface{} `json:"EncryptionContext,omitempty" xml:"EncryptionContext,omitempty"`
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
