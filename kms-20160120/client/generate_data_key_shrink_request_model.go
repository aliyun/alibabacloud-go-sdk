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
}

type GenerateDataKeyShrinkRequest struct {
	DryRun *string `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The JSON string that consists of key-value pairs.
	//
	// If you specify this parameter, an equivalent value is required when you call the [Decrypt](https://help.aliyun.com/document_detail/28950.html) operation. For more information, see [EncryptionContext](https://help.aliyun.com/document_detail/42975.html).
	//
	// example:
	//
	// {"Example":"Example"}
	EncryptionContextShrink *string `json:"EncryptionContext,omitempty" xml:"EncryptionContext,omitempty"`
	// The ID of the CMK. The ID must be globally unique.
	//
	// You can also set this parameter to an alias that is bound to the CMK. For more information, see [Alias overview](https://help.aliyun.com/document_detail/68522.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 7906979c-8e06-46a2-be2d-68e3ccbc****
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The type of the data key that you want to generate. Valid values:
	//
	// 	- AES_256: a 256-bit symmetric key
	//
	// 	- AES_128: a 128-bit symmetric key
	//
	// >  We recommend that you use the KeySpec or NumberOfBytes parameter to specify the length of a data key. If none of the parameters are specified, KMS generates a 256-bit data key. If both parameters are specified, KMS ignores the KeySpec parameter.
	//
	// example:
	//
	// AES_256
	KeySpec *string `json:"KeySpec,omitempty" xml:"KeySpec,omitempty"`
	// The length of the data key that you want to generate. Unit: bytes.
	//
	// Valid values: 1 to 1024.
	//
	// Default value:
	//
	// 	- If the KeySpec parameter is set to AES_256, set the value of the NumberOfBytes parameter to 32.
	//
	// 	- If the KeySpec parameter is set to AES_128, set the value of the NumberOfBytes parameter to 16.
	//
	// example:
	//
	// 256
	NumberOfBytes *int32 `json:"NumberOfBytes,omitempty" xml:"NumberOfBytes,omitempty"`
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

func (s *GenerateDataKeyShrinkRequest) Validate() error {
	return dara.Validate(s)
}
