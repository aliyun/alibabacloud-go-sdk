// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAsymmetricEncryptRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlgorithm(v string) *AsymmetricEncryptRequest
	GetAlgorithm() *string
	SetDryRun(v string) *AsymmetricEncryptRequest
	GetDryRun() *string
	SetKeyId(v string) *AsymmetricEncryptRequest
	GetKeyId() *string
	SetKeyVersionId(v string) *AsymmetricEncryptRequest
	GetKeyVersionId() *string
	SetPlaintext(v string) *AsymmetricEncryptRequest
	GetPlaintext() *string
}

type AsymmetricEncryptRequest struct {
	// The encryption algorithm.
	//
	// This parameter is required.
	//
	// example:
	//
	// RSAES_OAEP_SHA_1
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// Specifies whether to enable the dry run mode. Valid values:
	//
	// - true: enables the dry run mode.
	//
	// - false: disables the dry run mode. This is the default value.
	//
	// The dry run mode is used to test API calls and verify whether you have the required permissions on the resources and whether the request parameters are correctly configured. When the dry run mode is enabled, KMS always returns a failure and provides the cause of the failure. The causes include:
	//
	// - DryRunOperationError: The request would succeed if the DryRun parameter is not configured.
	//
	// - ValidationError: A specified parameter in the request is invalid.
	//
	// - AccessDeniedError: You do not have permissions to perform the operation on the KMS resource.
	//
	// example:
	//
	// false
	DryRun *string `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the key. You can also specify the alias or Alibaba Cloud Resource Name (ARN) of the key. For more information about aliases, see [Manage key aliases](https://help.aliyun.com/document_detail/480655.html).
	//
	// > To access a key that belongs to another Alibaba Cloud account, you must enter the ARN of the key. The ARN of a key is in the `acs:kms:${region}:${account}:key/${keyid}` format.
	//
	// This parameter is required.
	//
	// example:
	//
	// key-hzz630494463ejqjx****
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The ID of the key version. The globally unique identifier of the key version.
	//
	// > You can call the [ListKeyVersions](https://help.aliyun.com/document_detail/133966.html) operation to obtain the key version ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2ab1a983-7072-4bbc-a582-584b5bd8****
	KeyVersionId *string `json:"KeyVersionId,omitempty" xml:"KeyVersionId,omitempty"`
	// The plaintext to be encrypted. The plaintext must be Base64-encoded.
	//
	// This parameter is required.
	//
	// example:
	//
	// SGVsbG8gd29ybGQ=
	Plaintext *string `json:"Plaintext,omitempty" xml:"Plaintext,omitempty"`
}

func (s AsymmetricEncryptRequest) String() string {
	return dara.Prettify(s)
}

func (s AsymmetricEncryptRequest) GoString() string {
	return s.String()
}

func (s *AsymmetricEncryptRequest) GetAlgorithm() *string {
	return s.Algorithm
}

func (s *AsymmetricEncryptRequest) GetDryRun() *string {
	return s.DryRun
}

func (s *AsymmetricEncryptRequest) GetKeyId() *string {
	return s.KeyId
}

func (s *AsymmetricEncryptRequest) GetKeyVersionId() *string {
	return s.KeyVersionId
}

func (s *AsymmetricEncryptRequest) GetPlaintext() *string {
	return s.Plaintext
}

func (s *AsymmetricEncryptRequest) SetAlgorithm(v string) *AsymmetricEncryptRequest {
	s.Algorithm = &v
	return s
}

func (s *AsymmetricEncryptRequest) SetDryRun(v string) *AsymmetricEncryptRequest {
	s.DryRun = &v
	return s
}

func (s *AsymmetricEncryptRequest) SetKeyId(v string) *AsymmetricEncryptRequest {
	s.KeyId = &v
	return s
}

func (s *AsymmetricEncryptRequest) SetKeyVersionId(v string) *AsymmetricEncryptRequest {
	s.KeyVersionId = &v
	return s
}

func (s *AsymmetricEncryptRequest) SetPlaintext(v string) *AsymmetricEncryptRequest {
	s.Plaintext = &v
	return s
}

func (s *AsymmetricEncryptRequest) Validate() error {
	return dara.Validate(s)
}
