// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyMacRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlgorithm(v string) *VerifyMacRequest
	GetAlgorithm() *string
	SetDryRun(v string) *VerifyMacRequest
	GetDryRun() *string
	SetKeyId(v string) *VerifyMacRequest
	GetKeyId() *string
	SetMac(v string) *VerifyMacRequest
	GetMac() *string
	SetMessage(v string) *VerifyMacRequest
	GetMessage() *string
}

type VerifyMacRequest struct {
	// The algorithm used to generate the message authentication code. Valid values vary based on the key specification:
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
	// Specifies whether to enable DryRun mode. Valid values:
	//
	// - true: enables DryRun mode.
	//
	// - false (default): disables DryRun mode.
	//
	// DryRun mode is used to test API calls and verify whether you have the required permissions on the corresponding resources and whether the request parameters are correctly configured. When DryRun mode is enabled, KMS always returns a failure and provides the failure reason. Failure reasons include:
	//
	// - DryRunOperationError: The request would succeed if the DryRun parameter is not specified.
	//
	// - ValidationError: The parameters specified in the request are invalid.
	//
	// - AccessDeniedError: You are not authorized to perform this operation on the KMS resource.
	//
	// example:
	//
	// false
	DryRun *string `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the key. You can also specify a key alias or key Amazon Resource Name (ARN). For more information about aliases, refer to [Manage key aliases](https://help.aliyun.com/document_detail/480655.html).
	//
	// >To access a key in a different Alibaba Cloud account, you must specify the key ARN. The key ARN is in the format of `acs:kms:${region}:${account}:key/${keyid}`.
	//
	// This parameter is required.
	//
	// example:
	//
	// key-hzz630494463ejqjx****
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The Base64-encoding message authentication code to verify.
	//
	// This parameter is required.
	//
	// example:
	//
	// vz1Snp+jGJbgydCFRWVWxAwIMdyfKCSp+jnMWQ==
	Mac *string `json:"Mac,omitempty" xml:"Mac,omitempty"`
	// The original message data.
	//
	// Use Base64 encoding. For example, if the hexadecimal content of the message for which you want to generate a message authentication code is `[0x31, 0x32, 0x33, 0x34]`, the corresponding Base64-encoded value is `MTIzNA==`.
	//
	// This parameter is required.
	//
	// example:
	//
	// VGhlIHF1aWNrIGJyb3duIGZveCBqdW1wcyBvdmVyIHRoZSBsYXp5IGRvZy4=
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s VerifyMacRequest) String() string {
	return dara.Prettify(s)
}

func (s VerifyMacRequest) GoString() string {
	return s.String()
}

func (s *VerifyMacRequest) GetAlgorithm() *string {
	return s.Algorithm
}

func (s *VerifyMacRequest) GetDryRun() *string {
	return s.DryRun
}

func (s *VerifyMacRequest) GetKeyId() *string {
	return s.KeyId
}

func (s *VerifyMacRequest) GetMac() *string {
	return s.Mac
}

func (s *VerifyMacRequest) GetMessage() *string {
	return s.Message
}

func (s *VerifyMacRequest) SetAlgorithm(v string) *VerifyMacRequest {
	s.Algorithm = &v
	return s
}

func (s *VerifyMacRequest) SetDryRun(v string) *VerifyMacRequest {
	s.DryRun = &v
	return s
}

func (s *VerifyMacRequest) SetKeyId(v string) *VerifyMacRequest {
	s.KeyId = &v
	return s
}

func (s *VerifyMacRequest) SetMac(v string) *VerifyMacRequest {
	s.Mac = &v
	return s
}

func (s *VerifyMacRequest) SetMessage(v string) *VerifyMacRequest {
	s.Message = &v
	return s
}

func (s *VerifyMacRequest) Validate() error {
	return dara.Validate(s)
}
