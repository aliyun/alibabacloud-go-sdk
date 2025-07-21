// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateDataKeyWithoutPlaintextShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDryRun(v string) *GenerateDataKeyWithoutPlaintextShrinkRequest
	GetDryRun() *string
	SetEncryptionContextShrink(v string) *GenerateDataKeyWithoutPlaintextShrinkRequest
	GetEncryptionContextShrink() *string
	SetKeyId(v string) *GenerateDataKeyWithoutPlaintextShrinkRequest
	GetKeyId() *string
	SetKeySpec(v string) *GenerateDataKeyWithoutPlaintextShrinkRequest
	GetKeySpec() *string
	SetNumberOfBytes(v int32) *GenerateDataKeyWithoutPlaintextShrinkRequest
	GetNumberOfBytes() *int32
}

type GenerateDataKeyWithoutPlaintextShrinkRequest struct {
	DryRun *string `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// A JSON string that consists of key-value pairs. If you specify this parameter, an equivalent value is required when you call the Decrypt operation. For more information, see [EncryptionContext](https://help.aliyun.com/document_detail/42975.html).
	//
	// example:
	//
	// {"Example":"Example"}
	EncryptionContextShrink *string `json:"EncryptionContext,omitempty" xml:"EncryptionContext,omitempty"`
	// The globally unique ID of the CMK. You can also set this parameter to an alias that is bound to the CMK. For more information, see Use aliases.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234abcd-12ab-34cd-56ef-12345678****
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The length of the data key that you want to generate. Valid values:
	//
	// 	- AES_256: 256-bit symmetric key
	//
	// 	- AES_128: 128-bit symmetric key
	//
	// >  We recommend that you use the KeySpec or NumberOfBytes parameter to specify the length of a data key. If both of them are not specified, KMS generates a 256-bit data key. If both of them are specified, KMS ignores the KeySpec parameter.
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
	// 256
	NumberOfBytes *int32 `json:"NumberOfBytes,omitempty" xml:"NumberOfBytes,omitempty"`
}

func (s GenerateDataKeyWithoutPlaintextShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GenerateDataKeyWithoutPlaintextShrinkRequest) GoString() string {
	return s.String()
}

func (s *GenerateDataKeyWithoutPlaintextShrinkRequest) GetDryRun() *string {
	return s.DryRun
}

func (s *GenerateDataKeyWithoutPlaintextShrinkRequest) GetEncryptionContextShrink() *string {
	return s.EncryptionContextShrink
}

func (s *GenerateDataKeyWithoutPlaintextShrinkRequest) GetKeyId() *string {
	return s.KeyId
}

func (s *GenerateDataKeyWithoutPlaintextShrinkRequest) GetKeySpec() *string {
	return s.KeySpec
}

func (s *GenerateDataKeyWithoutPlaintextShrinkRequest) GetNumberOfBytes() *int32 {
	return s.NumberOfBytes
}

func (s *GenerateDataKeyWithoutPlaintextShrinkRequest) SetDryRun(v string) *GenerateDataKeyWithoutPlaintextShrinkRequest {
	s.DryRun = &v
	return s
}

func (s *GenerateDataKeyWithoutPlaintextShrinkRequest) SetEncryptionContextShrink(v string) *GenerateDataKeyWithoutPlaintextShrinkRequest {
	s.EncryptionContextShrink = &v
	return s
}

func (s *GenerateDataKeyWithoutPlaintextShrinkRequest) SetKeyId(v string) *GenerateDataKeyWithoutPlaintextShrinkRequest {
	s.KeyId = &v
	return s
}

func (s *GenerateDataKeyWithoutPlaintextShrinkRequest) SetKeySpec(v string) *GenerateDataKeyWithoutPlaintextShrinkRequest {
	s.KeySpec = &v
	return s
}

func (s *GenerateDataKeyWithoutPlaintextShrinkRequest) SetNumberOfBytes(v int32) *GenerateDataKeyWithoutPlaintextShrinkRequest {
	s.NumberOfBytes = &v
	return s
}

func (s *GenerateDataKeyWithoutPlaintextShrinkRequest) Validate() error {
	return dara.Validate(s)
}
