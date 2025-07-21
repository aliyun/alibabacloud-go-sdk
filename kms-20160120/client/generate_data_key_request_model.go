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
}

type GenerateDataKeyRequest struct {
	DryRun *string `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The JSON string that consists of key-value pairs.
	//
	// If you specify this parameter, an equivalent value is required when you call the [Decrypt](https://help.aliyun.com/document_detail/28950.html) operation. For more information, see [EncryptionContext](https://help.aliyun.com/document_detail/42975.html).
	//
	// example:
	//
	// {"Example":"Example"}
	EncryptionContext map[string]interface{} `json:"EncryptionContext,omitempty" xml:"EncryptionContext,omitempty"`
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

func (s *GenerateDataKeyRequest) Validate() error {
	return dara.Validate(s)
}
