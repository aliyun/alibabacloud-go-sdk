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
	SetRecipient(v string) *AsymmetricDecryptRequest
	GetRecipient() *string
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
	// The ciphertext to decrypt, encoded in Base64.
	//
	// > You can call the [AsymmetricEncrypt](https://help.aliyun.com/document_detail/148131.html) operation to generate the ciphertext.
	//
	// This parameter is required.
	//
	// example:
	//
	// BQKP+1zK6+ZEMxTP5qaVzcsgXtWplYBKm0NXdSnB5FzliFxE1bSiu4dnEIlca2JpeH7yz1/S6fed630H+hIH6DoM25fTLNcKj+mFB0Xnh9m2+HN59Mn4qyTfcUeadnfCXSWcGBouhXFwcdd2rJ3n337bzTf4jm659gZu3L0i6PLuxM9p7mqdwO0cKJPfGVfhnfMz+f4alMg79WB/NNyE2lyX7/qxvV49ObNrrJbKSFiz8Djocaf0IESNLMbfYI5bXjWkJlX92DQbKhibtQW8ZOJ//ZC6t0AWcUoKL6QDm/dg5koQalcleRinpB+QadFm894sLbVZ9+N4GVsv1W****==
	CiphertextBlob *string `json:"CiphertextBlob,omitempty" xml:"CiphertextBlob,omitempty"`
	// Specifies whether to enable DryRun mode. Valid values:
	//
	// - true: Enabled.
	//
	// - false (default): Disabled.
	//
	// DryRun mode is used to test API calls, verify that you have the required permissions on the relevant resources, and check whether the request parameters are configured correctly. When DryRun mode is enabled, KMS always returns a failure with a reason. The failure reasons include the following:
	//
	// - DryRunOperationError: The request would succeed if the DryRun parameter were not specified.
	//
	// - ValidationError: A parameter specified in the request is invalid.
	//
	// - AccessDeniedError: You do not have permission to perform this operation on the KMS resource.
	//
	// example:
	//
	// false
	DryRun *string `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the key. You can also specify a key alias or a key Amazon Resource Name (ARN). For more information about aliases, see [Manage key aliases](https://help.aliyun.com/document_detail/480655.html).
	//
	// > To access a key that belongs to another Alibaba Cloud account, you must specify the key ARN. The key ARN format is `acs:kms:${region}:${account}:key/${keyid}`.
	//
	// This parameter is required.
	//
	// example:
	//
	// key-hzz630494463ejqjx****
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The key version ID. The globally unique identifier of the key version.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2ab1a983-7072-4bbc-a582-584b5bd8****
	KeyVersionId *string `json:"KeyVersionId,omitempty" xml:"KeyVersionId,omitempty"`
	// example:
	//
	// { "AttestationDocument":"base64-encoded-attestion-document",  "KeyEncryptionAlgorithm":"RSAES_OAEP_SHA_256" }
	Recipient *string `json:"Recipient,omitempty" xml:"Recipient,omitempty"`
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

func (s *AsymmetricDecryptRequest) GetRecipient() *string {
	return s.Recipient
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

func (s *AsymmetricDecryptRequest) SetRecipient(v string) *AsymmetricDecryptRequest {
	s.Recipient = &v
	return s
}

func (s *AsymmetricDecryptRequest) Validate() error {
	return dara.Validate(s)
}
