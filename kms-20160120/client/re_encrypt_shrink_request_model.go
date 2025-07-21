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
	// The ciphertext that you want to re-encrypt.
	//
	// You can set this parameter to the ciphertext that is returned after a symmetric or asymmetric encryption operation.
	//
	// 	- Symmetric encryption: the ciphertext returned after you call the [Encrypt](https://help.aliyun.com/document_detail/28949.html), [GenerateDataKey](https://help.aliyun.com/document_detail/28948.html), [GenerateDataKeyWithoutPlaintext](https://help.aliyun.com/document_detail/134043.html), or [GenerateAndExportDataKey](https://help.aliyun.com/document_detail/176804.html) operation
	//
	// 	- Asymmetric encryption: the public key-encrypted ciphertext returned after you call the [GenerateAndExportDataKey](https://help.aliyun.com/document_detail/176804.html) operation, or the ciphertext encrypted by using the public key of an asymmetric key pair outside KMS
	//
	// This parameter is required.
	//
	// example:
	//
	// ODZhOWVmZDktM2QxNi00ODk0LWJkNGYtMWZjNDNmM2YyYWJmS7FmDBBQ0BkKsQrtRnidtPwirmDcS0ZuJCU41xxAAWk4Z8qsADfbV0b+i6kQmlvj79dJdGOvtX69Uycs901q********
	CiphertextBlob *string `json:"CiphertextBlob,omitempty" xml:"CiphertextBlob,omitempty"`
	// A JSON string that consists of key-value pairs. This parameter specifies the EncryptionContext that is used to re-encrypt the decrypted data or data key.
	//
	// example:
	//
	// {"Example":"Example"}
	DestinationEncryptionContextShrink *string `json:"DestinationEncryptionContext,omitempty" xml:"DestinationEncryptionContext,omitempty"`
	// The ID of the symmetric CMK that is used to re-encrypt the ciphertext after the ciphertext is decrypted.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234abcd-12ab-34cd-56ef-12345678****
	DestinationKeyId *string `json:"DestinationKeyId,omitempty" xml:"DestinationKeyId,omitempty"`
	DryRun           *string `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The encryption algorithm based on which the public key is used to encrypt the ciphertext specified by CiphertextBlob. For more information about encryption algorithms, see [AsymmetricDecrypt](https://help.aliyun.com/document_detail/148130.html).
	//
	// Valid values:
	//
	// 	- RSAES_OAEP_SHA_256
	//
	// 	- RSAES_OAEP_SHA_1
	//
	// 	- SM2PKE
	//
	// >  If you set CiphertextBlob to the public key-encrypted ciphertext that is returned after an asymmetric encryption operation, specify this parameter.
	//
	// example:
	//
	// RSAES_OAEP_SHA_256
	SourceEncryptionAlgorithm *string `json:"SourceEncryptionAlgorithm,omitempty" xml:"SourceEncryptionAlgorithm,omitempty"`
	// A JSON string that consists of key-value pairs. If you specify EncryptionContext when you call the [Encrypt](https://help.aliyun.com/document_detail/28949.html), [GenerateDataKey](https://help.aliyun.com/document_detail/28948.html), [GenerateDataKeyWithoutPlaintext](https://help.aliyun.com/document_detail/134043.html), or [GenerateAndExportDataKey](https://help.aliyun.com/document_detail/176804.html) operation to encrypt the data or data key, an equivalent value is required here. For more information, see [EncryptionContext](https://help.aliyun.com/document_detail/42975.html).
	//
	// >  If you set CiphertextBlob to the ciphertext that is returned after a symmetric encryption operation, specify this parameter.
	//
	// example:
	//
	// {"Example":"Example"}
	SourceEncryptionContextShrink *string `json:"SourceEncryptionContext,omitempty" xml:"SourceEncryptionContext,omitempty"`
	// The ID of the CMK that is used to decrypt the ciphertext.
	//
	// This parameter is the globally unique ID of the CMK.
	//
	// >  If you set CiphertextBlob to the public key-encrypted ciphertext that is returned after an asymmetric encryption operation, specify this parameter.
	//
	// example:
	//
	// 5c438b18-05be-40ad-b6c2-3be6752c****
	SourceKeyId *string `json:"SourceKeyId,omitempty" xml:"SourceKeyId,omitempty"`
	// The ID of the CMK version that is used to decrypt the ciphertext.
	//
	// >  If you set CiphertextBlob to the public key-encrypted ciphertext that is returned after an asymmetric encryption operation, specify this parameter.
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
