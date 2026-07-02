// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateMacRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlgorithm(v string) *GenerateMacRequest
	GetAlgorithm() *string
	SetDryRun(v string) *GenerateMacRequest
	GetDryRun() *string
	SetKeyId(v string) *GenerateMacRequest
	GetKeyId() *string
	SetMessage(v string) *GenerateMacRequest
	GetMessage() *string
}

type GenerateMacRequest struct {
	// The algorithm that is used to generate the message authentication code. Valid values vary based on the key specification:
	//
	// - HMAC_SM3
	//
	// - HMAC_SHA_224
	//
	// - HMAC_SHA_256
	//
	// - HMAC_SHA_384
	//
	// - HMAC_SHA_512
	//
	// This parameter is required.
	//
	// example:
	//
	// HMAC_SHA_256
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// Specifies whether to enable DryRun mode.
	//
	// - true: enabled
	//
	// - false (default): disabled
	//
	// DryRun mode is used to test API calls and verify whether you have the required permissions on the corresponding resources and whether the request parameters are correctly configured. When DryRun mode is enabled, KMS always returns a failure and provides the failure reason. Failure reasons include:
	//
	// - DryRunOperationError: The request would succeed without the DryRun parameter.
	//
	// - ValidationError: The parameters specified in the request are invalid.
	//
	// - AccessDeniedError: You are not authorized to perform this operation on the KMS resource.
	//
	// example:
	//
	// false
	DryRun *string `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the key. You can also specify a key alias or a key Amazon Resource Name (ARN). For more information about aliases, refer to [Manage key aliases](https://help.aliyun.com/document_detail/480655.html).
	//
	// >When you access a key that belongs to a different Alibaba Cloud account, you must specify the key ARN. The format of a key ARN is `acs:kms:${region}:${account}:key/${keyid}`.
	//
	// This parameter is required.
	//
	// example:
	//
	// key-hzz630494463ejqjx****
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The data for which you want to generate a message authentication code.
	//
	// The value is Base64-encoded. For example, if the hexadecimal data is `[0x31, 0x32, 0x33, 0x34]`, the corresponding Base64-encoded value is `MTIzNA==`.
	//
	// This parameter is required.
	//
	// example:
	//
	// VGhlIHF1aWNrIGJyb3duIGZveCBqdW1wcyBvdmVyIHRoZSBsYXp5IGRvZy4=
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s GenerateMacRequest) String() string {
	return dara.Prettify(s)
}

func (s GenerateMacRequest) GoString() string {
	return s.String()
}

func (s *GenerateMacRequest) GetAlgorithm() *string {
	return s.Algorithm
}

func (s *GenerateMacRequest) GetDryRun() *string {
	return s.DryRun
}

func (s *GenerateMacRequest) GetKeyId() *string {
	return s.KeyId
}

func (s *GenerateMacRequest) GetMessage() *string {
	return s.Message
}

func (s *GenerateMacRequest) SetAlgorithm(v string) *GenerateMacRequest {
	s.Algorithm = &v
	return s
}

func (s *GenerateMacRequest) SetDryRun(v string) *GenerateMacRequest {
	s.DryRun = &v
	return s
}

func (s *GenerateMacRequest) SetKeyId(v string) *GenerateMacRequest {
	s.KeyId = &v
	return s
}

func (s *GenerateMacRequest) SetMessage(v string) *GenerateMacRequest {
	s.Message = &v
	return s
}

func (s *GenerateMacRequest) Validate() error {
	return dara.Validate(s)
}
