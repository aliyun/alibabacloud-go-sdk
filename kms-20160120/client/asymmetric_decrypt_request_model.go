// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAsymmetricDecryptRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlgorithm(v string) *AsymmetricDecryptRequest
	GetAlgorithm() *string
	SetCiphertextBlob(v string) *AsymmetricDecryptRequest
	GetCiphertextBlob() *string
	SetDryRun(v string) *AsymmetricDecryptRequest
	GetDryRun() *string
	SetKeyId(v string) *AsymmetricDecryptRequest
	GetKeyId() *string
	SetKeyVersionId(v string) *AsymmetricDecryptRequest
	GetKeyVersionId() *string
}

type AsymmetricDecryptRequest struct {
	// The decryption algorithm.
	//
	// This parameter is required.
	//
	// example:
	//
	// RSAES_OAEP_SHA_1
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// The ciphertext to be decrypted. The ciphertext is encoded in Base64.
	//
	// > You can generate a ciphertext by calling the [AsymmetricEncrypt](https://help.aliyun.com/document_detail/148131.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// BQKP+1zK6+ZEMxTP5qaVzcsgXtWplYBKm0NXdSnB5FzliFxE1bSiu4dnEIlca2JpeH7yz1/S6fed630H+hIH6DoM25fTLNcKj+mFB0Xnh9m2+HN59Mn4qyTfcUeadnfCXSWcGBouhXFwcdd2rJ3n337bzTf4jm659gZu3L0i6PLuxM9p7mqdwO0cKJPfGVfhnfMz+f4alMg79WB/NNyE2lyX7/qxvV49ObNrrJbKSFiz8Djocaf0IESNLMbfYI5bXjWkJlX92DQbKhibtQW8ZOJ//ZC6t0AWcUoKL6QDm/dg5koQalcleRinpB+QadFm894sLbVZ9+N4GVsv1W****==
	CiphertextBlob *string `json:"CiphertextBlob,omitempty" xml:"CiphertextBlob,omitempty"`
	// Specifies whether to enable the dry run feature.
	//
	// - true: enables the dry run feature.
	//
	// - false: disables the dry run feature. This is the default value.
	//
	// The dry run feature is used to test API calls, verify the permissions on the specified resources, and check the validity of the request parameters. If you enable the dry run feature, KMS always returns a failure response and the cause of the failure. The causes of the failure include the following:
	//
	// - DryRunOperationError: The request would have succeeded if the DryRun parameter is not specified.
	//
	// - ValidationError: The specified parameter in the request is invalid.
	//
	// - AccessDeniedError: You are not authorized to perform this operation on the KMS resource.
	//
	// example:
	//
	// false
	DryRun *string `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the key. You can also specify the alias or Amazon Resource Name (ARN) of the key. For more information about aliases, see [Manage aliases](https://help.aliyun.com/document_detail/480655.html).
	//
	// > When you access a key in another Alibaba Cloud account, you must specify the ARN of the key. The ARN of a key is in the `acs:kms:${region}:${account}:key/${keyid}` format.
	//
	// This parameter is required.
	//
	// example:
	//
	// key-hzz630494463ejqjx****
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The ID of the key version. The globally unique identifier of the key version.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2ab1a983-7072-4bbc-a582-584b5bd8****
	KeyVersionId *string `json:"KeyVersionId,omitempty" xml:"KeyVersionId,omitempty"`
}

func (s AsymmetricDecryptRequest) String() string {
	return dara.Prettify(s)
}

func (s AsymmetricDecryptRequest) GoString() string {
	return s.String()
}

func (s *AsymmetricDecryptRequest) GetAlgorithm() *string {
	return s.Algorithm
}

func (s *AsymmetricDecryptRequest) GetCiphertextBlob() *string {
	return s.CiphertextBlob
}

func (s *AsymmetricDecryptRequest) GetDryRun() *string {
	return s.DryRun
}

func (s *AsymmetricDecryptRequest) GetKeyId() *string {
	return s.KeyId
}

func (s *AsymmetricDecryptRequest) GetKeyVersionId() *string {
	return s.KeyVersionId
}

func (s *AsymmetricDecryptRequest) SetAlgorithm(v string) *AsymmetricDecryptRequest {
	s.Algorithm = &v
	return s
}

func (s *AsymmetricDecryptRequest) SetCiphertextBlob(v string) *AsymmetricDecryptRequest {
	s.CiphertextBlob = &v
	return s
}

func (s *AsymmetricDecryptRequest) SetDryRun(v string) *AsymmetricDecryptRequest {
	s.DryRun = &v
	return s
}

func (s *AsymmetricDecryptRequest) SetKeyId(v string) *AsymmetricDecryptRequest {
	s.KeyId = &v
	return s
}

func (s *AsymmetricDecryptRequest) SetKeyVersionId(v string) *AsymmetricDecryptRequest {
	s.KeyVersionId = &v
	return s
}

func (s *AsymmetricDecryptRequest) Validate() error {
	return dara.Validate(s)
}
