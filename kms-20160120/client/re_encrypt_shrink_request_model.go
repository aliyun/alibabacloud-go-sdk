// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReEncryptShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCiphertextBlob(v string) *ReEncryptShrinkRequest
	GetCiphertextBlob() *string
	SetDestinationEncryptionContextShrink(v string) *ReEncryptShrinkRequest
	GetDestinationEncryptionContextShrink() *string
	SetDestinationKeyId(v string) *ReEncryptShrinkRequest
	GetDestinationKeyId() *string
	SetDryRun(v string) *ReEncryptShrinkRequest
	GetDryRun() *string
	SetSourceEncryptionAlgorithm(v string) *ReEncryptShrinkRequest
	GetSourceEncryptionAlgorithm() *string
	SetSourceEncryptionContextShrink(v string) *ReEncryptShrinkRequest
	GetSourceEncryptionContextShrink() *string
	SetSourceKeyId(v string) *ReEncryptShrinkRequest
	GetSourceKeyId() *string
	SetSourceKeyVersionId(v string) *ReEncryptShrinkRequest
	GetSourceKeyVersionId() *string
}

type ReEncryptShrinkRequest struct {
	// The ciphertext that you want to re-encrypt.<br> This parameter can be the ciphertext that is returned after symmetric or asymmetric key encryption.<br><br>
	//
	// - Symmetric encryption: the ciphertext that is returned after you call the [Encrypt](https://help.aliyun.com/document_detail/28949.html), [GenerateDataKey](https://help.aliyun.com/document_detail/28948.html), [GenerateDataKeyWithoutPlaintext](https://help.aliyun.com/document_detail/134043.html), or [GenerateAndExportDataKey](https://help.aliyun.com/document_detail/176804.html) operation.
	//
	// - Asymmetric key encryption: the data that is encrypted using a public key after you call the [GenerateAndExportDataKey](https://help.aliyun.com/document_detail/176804.html) operation, or the data that is encrypted using an asymmetric public key in an external system.
	//
	// This parameter is required.
	//
	// example:
	//
	// ODZhOWVmZDktM2QxNi00ODk0LWJkNGYtMWZjNDNmM2YyYWJmS7FmDBBQ0BkKsQrtRnidtPwirmDcS0ZuJCU41xxAAWk4Z8qsADfbV0b+i6kQmlvj79dJdGOvtX69Uycs901q********
	CiphertextBlob *string `json:"CiphertextBlob,omitempty" xml:"CiphertextBlob,omitempty"`
	// A JSON string that consists of key-value pairs. This parameter specifies the encryption context for the destination master key.
	//
	// example:
	//
	// {"Example":"Example"}
	DestinationEncryptionContextShrink *string `json:"DestinationEncryptionContext,omitempty" xml:"DestinationEncryptionContext,omitempty"`
	// The ID of the symmetric master key that is used to re-encrypt the data after the ciphertext is decrypted.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234abcd-12ab-34cd-56ef-12345678****
	DestinationKeyId *string `json:"DestinationKeyId,omitempty" xml:"DestinationKeyId,omitempty"`
	// Specifies whether to enable the DryRun mode.
	//
	// - true: enables the DryRun mode.
	//
	// - false (default): disables the DryRun mode.
	//
	// The DryRun mode is used to test API calls, verify whether you have the permissions to perform operations on the required resources, and check whether the request parameters are valid. If you enable the DryRun mode, KMS always returns a failure and a reason for the failure. The reasons for the failure include the following:
	//
	// - DryRunOperationError: The request would have succeeded if the DryRun parameter was not configured.
	//
	// - ValidationError: The parameters specified in the request are invalid.
	//
	// - AccessDeniedError: You are not authorized to perform the operation on the KMS resource.
	//
	// example:
	//
	// false
	DryRun *string `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// If CiphertextBlob is the result of public key encryption, specify the public key encryption algorithm. For more information about the algorithms, see [AsymmetricDecrypt](https://help.aliyun.com/document_detail/148130.html).<br> Valid values:<br><br>
	//
	// - RSAES_OAEP_SHA_256
	//
	// - RSAES_OAEP_SHA_1
	//
	// - SM2PKE
	//
	// > You must specify this parameter when CiphertextBlob is the data that is encrypted using a public key after asymmetric key encryption.
	//
	// example:
	//
	// RSAES_OAEP_SHA_256
	SourceEncryptionAlgorithm *string `json:"SourceEncryptionAlgorithm,omitempty" xml:"SourceEncryptionAlgorithm,omitempty"`
	// A JSON string that consists of key-value pairs. If you specify this parameter when you call the [Encrypt](https://help.aliyun.com/document_detail/28949.html), [GenerateDataKey](https://help.aliyun.com/document_detail/28948.html), [GenerateDataKeyWithoutPlaintext](https://help.aliyun.com/document_detail/134043.html), or [GenerateAndExportDataKey](https://help.aliyun.com/document_detail/176804.html) operation, you must specify the same parameter to decrypt the data. For more information, see [EncryptionContext](https://help.aliyun.com/document_detail/42975.html).
	//
	// > You must specify this parameter when CiphertextBlob is the ciphertext that is returned after symmetric encryption.
	//
	// example:
	//
	// {"Example":"Example"}
	SourceEncryptionContextShrink *string `json:"SourceEncryptionContext,omitempty" xml:"SourceEncryptionContext,omitempty"`
	// The ID of the master key that is used to decrypt the ciphertext.<br> The globally unique identifier of the master key.<br><br>
	//
	// > You must specify this parameter when CiphertextBlob is the data that is encrypted using a public key after asymmetric key encryption.
	//
	// example:
	//
	// 5c438b18-05be-40ad-b6c2-3be6752c****
	SourceKeyId *string `json:"SourceKeyId,omitempty" xml:"SourceKeyId,omitempty"`
	// The ID of the key version that is used to decrypt the ciphertext.
	//
	// > You must specify this parameter when CiphertextBlob is the data that is encrypted using a public key after asymmetric key encryption.
	//
	// example:
	//
	// 2ab1a983-7072-4bbc-a582-584b5bd8****
	SourceKeyVersionId *string `json:"SourceKeyVersionId,omitempty" xml:"SourceKeyVersionId,omitempty"`
}

func (s ReEncryptShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ReEncryptShrinkRequest) GoString() string {
	return s.String()
}

func (s *ReEncryptShrinkRequest) GetCiphertextBlob() *string {
	return s.CiphertextBlob
}

func (s *ReEncryptShrinkRequest) GetDestinationEncryptionContextShrink() *string {
	return s.DestinationEncryptionContextShrink
}

func (s *ReEncryptShrinkRequest) GetDestinationKeyId() *string {
	return s.DestinationKeyId
}

func (s *ReEncryptShrinkRequest) GetDryRun() *string {
	return s.DryRun
}

func (s *ReEncryptShrinkRequest) GetSourceEncryptionAlgorithm() *string {
	return s.SourceEncryptionAlgorithm
}

func (s *ReEncryptShrinkRequest) GetSourceEncryptionContextShrink() *string {
	return s.SourceEncryptionContextShrink
}

func (s *ReEncryptShrinkRequest) GetSourceKeyId() *string {
	return s.SourceKeyId
}

func (s *ReEncryptShrinkRequest) GetSourceKeyVersionId() *string {
	return s.SourceKeyVersionId
}

func (s *ReEncryptShrinkRequest) SetCiphertextBlob(v string) *ReEncryptShrinkRequest {
	s.CiphertextBlob = &v
	return s
}

func (s *ReEncryptShrinkRequest) SetDestinationEncryptionContextShrink(v string) *ReEncryptShrinkRequest {
	s.DestinationEncryptionContextShrink = &v
	return s
}

func (s *ReEncryptShrinkRequest) SetDestinationKeyId(v string) *ReEncryptShrinkRequest {
	s.DestinationKeyId = &v
	return s
}

func (s *ReEncryptShrinkRequest) SetDryRun(v string) *ReEncryptShrinkRequest {
	s.DryRun = &v
	return s
}

func (s *ReEncryptShrinkRequest) SetSourceEncryptionAlgorithm(v string) *ReEncryptShrinkRequest {
	s.SourceEncryptionAlgorithm = &v
	return s
}

func (s *ReEncryptShrinkRequest) SetSourceEncryptionContextShrink(v string) *ReEncryptShrinkRequest {
	s.SourceEncryptionContextShrink = &v
	return s
}

func (s *ReEncryptShrinkRequest) SetSourceKeyId(v string) *ReEncryptShrinkRequest {
	s.SourceKeyId = &v
	return s
}

func (s *ReEncryptShrinkRequest) SetSourceKeyVersionId(v string) *ReEncryptShrinkRequest {
	s.SourceKeyVersionId = &v
	return s
}

func (s *ReEncryptShrinkRequest) Validate() error {
	return dara.Validate(s)
}
