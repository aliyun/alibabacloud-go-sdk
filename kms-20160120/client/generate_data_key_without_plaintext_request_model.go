// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateDataKeyWithoutPlaintextRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDryRun(v string) *GenerateDataKeyWithoutPlaintextRequest
	GetDryRun() *string
	SetEncryptionContext(v map[string]interface{}) *GenerateDataKeyWithoutPlaintextRequest
	GetEncryptionContext() map[string]interface{}
	SetKeyId(v string) *GenerateDataKeyWithoutPlaintextRequest
	GetKeyId() *string
	SetKeySpec(v string) *GenerateDataKeyWithoutPlaintextRequest
	GetKeySpec() *string
	SetNumberOfBytes(v int32) *GenerateDataKeyWithoutPlaintextRequest
	GetNumberOfBytes() *int32
}

type GenerateDataKeyWithoutPlaintextRequest struct {
	// Specifies whether to enable the DryRun mode.
	//
	// - true: enables the DryRun mode.
	//
	// - false (default): disables the DryRun mode.
	//
	// The DryRun mode is used to test API calls, verify your permissions on the required resources, and check if the request parameters are valid. If you enable the DryRun mode, KMS returns a failure response with a reason. The failure reasons include the following:
	//
	// - DryRunOperationError: The request would have succeeded if the DryRun parameter was not specified.
	//
	// - ValidationError: The request parameters are invalid.
	//
	// - AccessDeniedError: You are not authorized to perform this operation on the KMS resource.
	//
	// example:
	//
	// false
	DryRun *string `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// A JSON string of key-value pairs. If you specify this parameter, you must provide the same parameter when you call the Decrypt operation. For more information, see [EncryptionContext](https://help.aliyun.com/document_detail/42975.html).
	//
	// example:
	//
	// {"Example":"Example"}
	EncryptionContext map[string]interface{} `json:"EncryptionContext,omitempty" xml:"EncryptionContext,omitempty"`
	// The globally unique identifier of the CMK. You can also specify an alias that is bound to the CMK. For more information about how to use an alias, see Alias overview.
	//
	// This parameter is required.
	//
	// example:
	//
	// 599fa825-17de-417e-9554-bb032cc6****
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The length of the data key to generate. Valid values:
	//
	// - AES_256: a 256-bit symmetric key
	//
	// - AES_128: a 128-bit symmetric key
	//
	// > Use KeySpec or NumberOfBytes to specify the length of the data key. If you do not specify either parameter, KMS generates a 256-bit data key. If you specify both parameters, KMS ignores the KeySpec parameter.
	//
	// example:
	//
	// AES_256
	KeySpec *string `json:"KeySpec,omitempty" xml:"KeySpec,omitempty"`
	// The length of the data key to generate.<br> Valid values: 1 to 1024.<br> Unit: bytes<br><br><br><br><br>
	//
	// example:
	//
	// 256
	NumberOfBytes *int32 `json:"NumberOfBytes,omitempty" xml:"NumberOfBytes,omitempty"`
}

func (s GenerateDataKeyWithoutPlaintextRequest) String() string {
	return dara.Prettify(s)
}

func (s GenerateDataKeyWithoutPlaintextRequest) GoString() string {
	return s.String()
}

func (s *GenerateDataKeyWithoutPlaintextRequest) GetDryRun() *string {
	return s.DryRun
}

func (s *GenerateDataKeyWithoutPlaintextRequest) GetEncryptionContext() map[string]interface{} {
	return s.EncryptionContext
}

func (s *GenerateDataKeyWithoutPlaintextRequest) GetKeyId() *string {
	return s.KeyId
}

func (s *GenerateDataKeyWithoutPlaintextRequest) GetKeySpec() *string {
	return s.KeySpec
}

func (s *GenerateDataKeyWithoutPlaintextRequest) GetNumberOfBytes() *int32 {
	return s.NumberOfBytes
}

func (s *GenerateDataKeyWithoutPlaintextRequest) SetDryRun(v string) *GenerateDataKeyWithoutPlaintextRequest {
	s.DryRun = &v
	return s
}

func (s *GenerateDataKeyWithoutPlaintextRequest) SetEncryptionContext(v map[string]interface{}) *GenerateDataKeyWithoutPlaintextRequest {
	s.EncryptionContext = v
	return s
}

func (s *GenerateDataKeyWithoutPlaintextRequest) SetKeyId(v string) *GenerateDataKeyWithoutPlaintextRequest {
	s.KeyId = &v
	return s
}

func (s *GenerateDataKeyWithoutPlaintextRequest) SetKeySpec(v string) *GenerateDataKeyWithoutPlaintextRequest {
	s.KeySpec = &v
	return s
}

func (s *GenerateDataKeyWithoutPlaintextRequest) SetNumberOfBytes(v int32) *GenerateDataKeyWithoutPlaintextRequest {
	s.NumberOfBytes = &v
	return s
}

func (s *GenerateDataKeyWithoutPlaintextRequest) Validate() error {
	return dara.Validate(s)
}
